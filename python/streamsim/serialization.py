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

import datetime
import io
import json
import struct
import pkg_resources

import astropy.time
import fastavro

from lsst.alert.packet import get_latest_schema_version, SchemaRegistry


def _load_latest_schema():
    latest_version_str = ".".join(map(str, get_latest_schema_version()))
    return SchemaRegistry.from_filesystem().get_by_version(latest_version_str)


_schema_latest = _load_latest_schema()


def serialize_time_offset(timedelta):
    """Encodes a timedelta as a string.

    The encoding is the python string representation of the floating-point
    total seconds in the timedelta.

    Parameters
    ----------
    timedelta : `datetime.timedelta`
        A stdlib timedelta to serialize.

    Returns
    -------
    s : `str`
        A string serialization of the timedelta.
    """
    return repr(timedelta.total_seconds())


def get_message_time_offset(msg):
    """Get the timedelta offset measuring the age of a message relative to the
    first message in the stream.

    Parameters
    ----------
    msg : `confluent_kafka.Message`
        A message retrieved from a Kafka topic

    Returns
    -------
    timedelta : `datetime.timedelta`
        The time delay between the first message in msg's topic and msg.
    """
    headers = dict(msg.headers())
    return deserialize_time_offset(headers["alertsim-time-offset"])


def deserialize_time_offset(header_value):
    """Decode a timedelta which was encoded with `serialize_time_offset`.

    Parameters
    ----------
    header_value : `str`
        A string encoding of a timedelta

    Returns
    -------
    timedelta : `datetime.timedelta`
        The timedelta represented by `header_value`
    """
    return datetime.timedelta(seconds=float(header_value))


ConfluentWireFormatHeader = struct.Struct(">bi")


def serialize_confluent_wire_header(schema_version):
    """Returns the byte prefix for Confluent Wire Format-style Kafka messages.

    Parameters
    ----------
    schema_version : `int`
        A version number which indicates the Confluent Schema Registry ID
        number of the Avro schema used to encode the message that follows this
        header.

    Returns
    -------
    header : `bytes`
        The 5-byte encoded message prefix.

    Notes
    -----
    The Confluent Wire Format is described more fully here:
    https://docs.confluent.io/current/schema-registry/serdes-develop/index.html#wire-format
    """
    return ConfluentWireFormatHeader.pack(0, schema_version)


def deserialize_confluent_wire_header(raw):
    """Parses the byte prefix for Confluent Wire Format-style Kafka messages.

    Parameters
    ----------
    raw : `bytes`
        The 5-byte encoded message prefix.

    Returns
    -------
    schema_version : `int`
        A version number which indicates the Confluent Schema Registry ID
        number of the Avro schema used to encode the message that follows this
        header.
    """
    _, version = ConfluentWireFormatHeader.unpack(raw)
    return version


def serialize_alert(alert, schema=_schema_latest):
    """Serialize an alert to a byte sequence for sending to Kafka.

    Parameters
    ----------
    alert : `dict`
        An alert payload to be serialized.
    schema : `dict`, optional
        An Avro schema definition describing how to encode `alert`. By default,
        the most recent schema version is used.

    Returns
    -------
    serialized : `bytes`
        The byte sequence describing the alert, including the Confluent Wire
        Format prefix.
    """
    buf = io.BytesIO()
    # TODO: Use a proper schema versioning system
    buf.write(serialize_confluent_wire_header(0))
    buf.write(schema.serialize(alert))
    return buf.getvalue()


def deserialize_alert(alert_bytes, schema=_schema_latest):
    """Deserialize an alert message from Kafka.

    Paramaters
    ----------
    alert_bytes : `bytes`
        Binary-encoding serialized Avro alert, including Confluent Wire
        Format prefix.
    schema : `dict`, optional
        An Avro schema definition describing how `alert_bytes` is encoded. By
        default, the latest schema is used.

    Returns
    -------
    alert : `dict`
        An alert payload.

    """
    header_bytes = alert_bytes[:5]
    version = deserialize_confluent_wire_header(header_bytes)
    assert version == 0
    return schema.deserialize(alert_bytes[5:])


def alert_time(alert):
    """Convert an alert's midPointTai field from MJD float into a datetime.

    Parameters
    ----------
    alert : `dict`
        An alert message with a diaSource.midPointTai encoded as a Modified
        Julian Date float.

    Returns
    -------
    timestamp : `datetime.datetime`
        The converted diaSource.midPointTai field.
    """
    raw = alert["diaSource"]["midPointTai"]
    return astropy.time.Time(raw, format="mjd").to_datetime()
