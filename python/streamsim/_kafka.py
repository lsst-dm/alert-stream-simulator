# This file is part of alert-stream-simulator.
#
# Developed for the LSST Data Management System.
# This product includes software developed by the LSST Project
# (https://www.lsst.org).
# See the COPYRIGHT file at the top-level directory of this distribution
# for details of code ownership.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

import logging
import time

import confluent_kafka
import confluent_kafka.admin


logger = logging.getLogger("rubin-alert-sim.kafka")


class _KafkaClient(object):
    """Combined client for Kafka producing, consuming, and administration.

    """
    def __init__(self, broker_url, id="rubin-alert-sim"):
        logger.debug(f"creating client to connect to broker url={broker_url} id={id}")
        admin_config = {
            "bootstrap.servers": broker_url,
            "socket.timeout.ms": 5_000,
            "error_cb": logger.error,
            "throttle_cb": logger.warn,
        }
        self.admin = confluent_kafka.admin.AdminClient(admin_config)
        producer_config = {
            "bootstrap.servers": broker_url,
            "socket.timeout.ms": 5_000,
            "error_cb": logger.error,
            "throttle_cb": logger.warn,
            "socket.keepalive.enable": True,
            "message.max.bytes": 10_000_000,
            "queue.buffering.max.ms": 100,
        }
        self.producer = confluent_kafka.Producer(producer_config)
        consumer_config = {
            "bootstrap.servers": broker_url,
            "socket.timeout.ms": 5_000,
            "error_cb": logger.error,
            "throttle_cb": logger.warn,
            "group.id": id,
            "auto.offset.reset": "earliest",
            "enable.partition.eof": True,
        }
        self.consumer = confluent_kafka.Consumer(consumer_config)

    def delete_topic(self, topic):
        """ Delete a topic from the Kafka cluster. """
        logger.debug(f"deleting topic name={topic}")
        response = self.admin.delete_topics([topic], operation_timeout=5.0)
        return response[topic].result()

    def create_topic(self, topic, num_partitions=1, delete_if_exists=False):
        """ Create a topic in the Kafka cluster. """
        logger.debug(f"creating topic name={topic} num_partitions={num_partitions}")
        if delete_if_exists:
            return self._create_topic_force(topic, num_partitions)
        new_topic = confluent_kafka.admin.NewTopic(
            topic=topic,
            num_partitions=num_partitions,
            replication_factor=1,
        )
        response = self.admin.create_topics([new_topic], operation_timeout=5.0)
        return response[topic].result()

    def _create_topic_force(self, topic, num_partitions):
        """ Create a topic. Delete it if it already exists. """
        new_topic = confluent_kafka.admin.NewTopic(
            topic=topic,
            num_partitions=num_partitions,
            replication_factor=1,
        )

        # Looping and retrying appears to be necessary to deal with timing
        # inconsistencies in the broker's internal state. The broker will
        # return and confirm the delete before actually doing the delete by
        # removing metadata from Zookeeper. The only resolution is to sleep,
        # here. See https://github.com/apache/kafka/pull/7343.
        for iteration in range(10):
            try:
                logger.debug(f"attempting topic creation name={topic}")
                response = self.admin.create_topics(
                    [new_topic], operation_timeout=5.0,
                )
                res = response[topic].result()
                logger.debug(f"topic created name={topic}")
                return res
            except confluent_kafka.KafkaException as e:
                if _is_topic_exists_error(e):
                    logger.debug(f"topic exists, deleting and retrying")
                    try:
                        self.delete_topic(topic)
                    except confluent_kafka.KafkaException as delete_exc:
                        if _is_unknown_topic_error(delete_exc):
                            logger.debug(delete_exc)
                            pass
                        else:
                            raise
                else:
                    raise
                time.sleep(0.05 * (2**iteration))

        # Last try.
        response = self.admin.create_topics(
            [new_topic], operation_timeout=5.0,
        )
        return response[topic].result()

    def describe_topic(self, topic, timeout=5.0):
        """ Fetch confluent_kafka.TopicMetadata describing a topic. """
        logger.debug(f"fetching cluster metadata to describe topic name={topic}")
        cluster_meta = self.consumer.list_topics(timeout=timeout)
        return cluster_meta.topics[topic]

    def close(self):
        """ Shut down the client's underlying consumer and producer. """
        logger.debug("shutting down client")
        self.consumer.close()
        self.producer.flush()

    def subscribe(self, topic, timeout=10.0):
        """Subscribes to a topic for consuming. This method doesn't use Kafka's
        Consumer Groups; it assigns all partitions manually to this process.

        """
        logger.debug(f"subscribing to topic {topic}")
        topic_meta = self.describe_topic(topic, timeout)
        assignment = []
        for partition_id in topic_meta.partitions.keys():
            logger.debug(f"adding subscription to topic partition={partition_id}")
            tp = confluent_kafka.TopicPartition(
                topic=topic,
                partition=partition_id,
                offset=confluent_kafka.OFFSET_BEGINNING,
            )
            assignment.append(tp)

        logger.debug("registering topic assignment")
        self.consumer.assign(assignment)

    def iterate(self, batch_size=100, batch_timeout=1.0):
        """Returns a generator which iterates over the messages in the topics
        to which the client is subscribed.

        The generator stops when the client has hit the last message in all
        partitions. This set of partitions is calculated just once when
        iterate() is first called; calling subscribe() after iterate() may
        cause inconsistent behavior.
        """
        assignment = self.consumer.assignment()

        # Make a map of topic-name -> set of partition IDs we're assigned to.
        # When we hit a partition EOF, remove that partition from the map.
        active_partitions = {}
        for tp in assignment:
            if tp.topic not in active_partitions:
                active_partitions[tp.topic] = set()
            active_partitions[tp.topic].add(tp.partition)

        while len(active_partitions) > 0:
            messages = self.consumer.consume(batch_size, batch_timeout)
            for m in messages:
                err = m.error()
                if err is None:
                    yield m
                elif err.code() == confluent_kafka.KafkaError._PARTITION_EOF:
                    logger.debug(f"eof for topic={m.topic()} partition={m.partition()}")
                    # Done with this partition, remove it
                    partition_set = active_partitions[m.topic()]
                    partition_set.remove(m.partition())
                    if len(partition_set) == 0:
                        # Done with all partitions for the topic, remove it
                        del active_partitions[m.topic()]
                else:
                    raise(err)


def _is_topic_exists_error(kafka_exception):
    """Returns True iff kafka_exception has the code TOPIC_ALREADY_EXISTS.

    """
    code = kafka_exception.args[0].code()
    return code == confluent_kafka.KafkaError.TOPIC_ALREADY_EXISTS


def _is_unknown_topic_error(kafka_exception):
    """Returns True iff kafka_exception has the code UNKNOWN_TOPIC_OR_PART.

    """
    code = kafka_exception.args[0].code()
    return code == confluent_kafka.KafkaError.UNKNOWN_TOPIC_OR_PART
