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
    kafka_client = _kafka._KafkaClient(broker)
    logger.debug(f"creating topic {dst_topic}")
    kafka_client.create_topic(dst_topic, dst_topic_partitions, force)

    logger.debug(f"subscribing to topic {src_topic}")
    kafka_client.subscribe(src_topic)
    start = datetime.datetime.now()
    for msg in kafka_client.iterate():
        since_start = datetime.datetime.now() - start
        offset = serialization.get_message_time_offset(msg)
        if since_start < offset:
            time.sleep((offset - since_start).total_seconds())
        logger.debug(f"sending message of size {len(msg)}")
        kafka_client.producer.produce(dst_topic, msg.value())

    kafka_client.close()
    dur = (datetime.datetime.now() - start).total_seconds()
    logger.info(f"sent 1000 alerts in {dur}s ({1000.0/dur}/s)")
