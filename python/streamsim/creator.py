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
import itertools

import fastavro

from streamsim import _kafka, serialization


logger = logging.getLogger("rubin-alert-sim.prepare")


def create(broker, topic, alert_file, timeout, force=False):
    """Creates a new alert stream in the broker. Returns the number of alerts
    in the stream.

    Alert streams exist in the simulator as Kafka topics with all alerts
    pre-serialized and stored in Kafka as individual messages. Each Kafka
    message has headers, and we use those headers to mark the "time offset" of
    the alert packet.

    The time offset is the number of seconds (as a float) since the first alert
    in the stream. Later, when we play the stream back, we can obey the header
    time offsets, sleeping to let data get emitted at a realistic rate.

    Parameters
    ----------
    broker : `str`
        The URL of a Kafka Broker to connect to.
    topic : `str`
        The name of a Kafka topic to create.
    alert_file : derivative of `IOBase`
        A file-like stream which holds alert data, serialized in Avro container
        format.
    timeout : `float`
        How long, in seconds, to wait for Kafka operations to return
    force : `bool`
        If true, overwrite the topic if it already exists
    """
    reader = fastavro.reader(alert_file)
    kafka_client = _kafka._KafkaClient(broker)

    # Create a new topic for us to write to.
    kafka_client.create_topic(topic, num_partitions=1, delete_if_exists=force)

    # Peak at the first alert in the file. This becomes our reference point for
    # measuring timestamps.
    first_alert = next(reader)
    first_timestamp = serialization.alert_time(first_alert)
    logger.debug(f"first timestamp: {first_timestamp}")
    n = 0
    for alert in itertools.chain([first_alert], reader):
        alert_bytes = serialization.serialize_alert(alert)
        time_offset = (serialization.alert_time(alert) - first_timestamp)
        logger.debug(f"producing alert id={alert['alertId']} with time_offset={time_offset}")
        kafka_client.producer.produce(
            topic=topic,
            value=alert_bytes,
            headers={
                "alertsim-time-offset": serialization.serialize_time_offset(time_offset),
            },
        )
        n += 1

    kafka_client.close()

    return n
