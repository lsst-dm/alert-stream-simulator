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
import datetime
import time

from streamsim import _kafka, serialization


logger = logging.getLogger("rubin-alert-sim.play")


def play(broker, src_topic, dst_topic, dst_topic_partitions, force):
    """Replays an existing alert stream in the broker, copying it from
    src_topic to dst_topic while obeying time offset headers.

    Alert streams exist in the simulator as Kafka topics with all alerts
    pre-serialized and stored in Kafka as individual messages. Each Kafka
    message has headers, and we use those headers to mark the "time offset" of
    the alert packet.

    The time offset is the number of seconds (as a float) since the first alert
    in the stream. When we play the stream back, we obey the header time
    offsets, sleeping to let data get emitted at a realistic rate.

    Parameters
    ----------
    broker : `str`
        The URL of a Kafka Broker to connect to.
    src_topic : `str`
        The name of a Kafka topic to act as the source. All alerts will be
        read from this topic. This topic should have been created with the
        `create-stream` command.
    dst_topic : `str`
        The name of a Kafka topic to create to act as the real-time stream.
    dst_topic_partitions : `int`
        How many partitions to create for `dst_topic`
    force : `bool`
        If true, overwrite `dst_topic` if it already exists
    """
    kafka_client = _kafka._KafkaClient(broker)
    logger.debug(f"creating topic {dst_topic}")
    kafka_client.create_topic(dst_topic, dst_topic_partitions, force)

    logger.debug(f"subscribing to topic {src_topic}")
    kafka_client.subscribe(src_topic)

    start = datetime.datetime.now()
    n = 0
    for msg in kafka_client.iterate():
        since_start = datetime.datetime.now() - start
        offset = serialization.get_message_time_offset(msg)
        if since_start < offset:
            time.sleep((offset - since_start).total_seconds())

        logger.debug(f"sending message of size {len(msg)}")
        kafka_client.producer.produce(dst_topic, msg.value())

        n += 1

    kafka_client.close()

    return n
