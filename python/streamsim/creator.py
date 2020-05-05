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
import io
import logging
import itertools
import time
import json
import datetime

import astropy.time
import confluent_kafka
import confluent_kafka.admin
import fastavro

from streamsim import _kafka


logger = logging.getLogger("rubin-alert-sim.prepare")


def create(broker, topic, alert_file, timeout, force=False):
    reader = fastavro.reader(alert_file)
    kafka_client = _kafka._KafkaClient(broker)

    # Create a new topic for us to write to.
    kafka_client.create_topic(topic, num_partitions=1, delete_if_exists=force)

    # Peak at the first alert in the file. This becomes our reference point for
    # measuring timestamps.
    first_alert = next(reader)
    first_timestamp = alert_time(first_alert)
    n = 0
    for alert in itertools.chain([first_alert], reader):
        ser = serialize(reader.writer_schema, alert)
        time_offset = (alert_time(alert) - first_timestamp)

        kafka_client.producer.produce(topic, ser, headers={
            "alertsim-time-offset": serialize_time_offset(time_offset),
        })
        n += 1

    kafka_client.close()
    return n


def serialize_time_offset(timedelta):
    """Encodes a timedelta as a string.

    The encoding is the python string representation of the floating-point total
    seconds in the timedelta.

    """
    return repr(timedelta.total_seconds())


def get_message_time_offset(msg):
    """Get the timedelta offset measuring the age of a message relative to the first
    message in the stream.

    """
    headers = dict(msg.headers())
    return deserialize_time_offset(headers["alertsim-time-offset"])


def deserialize_time_offset(header_value):
    """ Decode a timedeltas which was encoded with `serialize_time_offset`. """
    return datetime.timedelta(seconds=float(header_value))


def serialize(schema, alert):
    """ Serialize an alert to a byte sequence. """
    buf = io.BytesIO()
    fastavro.schemaless_writer(buf, schema, alert)
    return buf.read()


def alert_time(alert):
    """Convert an alert's midPointTai field from MJD float into a datetime.

    """
    raw = alert["diaSource"]["midPointTai"]
    return astropy.time.Time(raw, format="mjd").to_datetime()
