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

from streamsim import _kafka, timestamps


logger = logging.getLogger("rubin-alert-sim.play")


def play(broker, src_topic, dst_topic, dst_topic_partitions, create_topic,
         repeat_interval=-1, tls_config=None):
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
    create_topic : `bool`
        If true, create `dst_topic`, overwritig if it already exists
    tls_config : `streamsim._kafka.TLSConfig` or `None`
        If not None, then a configuration bundle for TLS auth when connecting
        to the broker.
    """
    kafka_client = _kafka._KafkaClient(broker, tls_config=tls_config)
    if create_topic:
        logger.debug(f"creating topic {dst_topic}")
        kafka_client.create_topic(dst_topic, dst_topic_partitions, True)

    logger.debug(f"subscribing to topic {src_topic}")
    kafka_client.subscribe(src_topic)

    n_total = 0
    while True:
        start = datetime.datetime.now()
        n = 0
        for msg in kafka_client.iterate():
            since_start = datetime.datetime.now() - start
            offset = timestamps.get_message_time_offset(msg)
            if since_start < offset:
                time.sleep((offset - since_start).total_seconds())

            logger.debug(f"sending message of size {len(msg)}")
            kafka_client.producer.produce(dst_topic, msg.value())

            n += 1
        kafka_client.producer.flush()
        n_total += n

        if repeat_interval < 0:
            break

        since_start_sec = (datetime.datetime.now() - start).total_seconds()
        logger.info(f"sent {n} alerts in {since_start_sec:.2f}s ({n / since_start_sec:.2f}/s)")

        if since_start_sec < repeat_interval:
            logger.info(f"going to sleep for {repeat_interval - since_start_sec:.2f}s")

            time.sleep(repeat_interval - since_start_sec)
        else:
            logger.warn(f"unable to keep up with repeat interval: took {since_start_sec}s to run"
                        + ", interval is {repeat_interval}")
        kafka_client.seek_to_beginning()

    kafka_client.close()

    return n_total
