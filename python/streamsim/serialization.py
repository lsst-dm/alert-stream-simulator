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
import struct

import astropy.time
import fastavro


def serialize_time_offset(timedelta):
    """Encodes a timedelta as a string.

    The encoding is the python string representation of the floating-point
    total seconds in the timedelta.

    """
    return repr(timedelta.total_seconds())


def get_message_time_offset(msg):
    """Get the timedelta offset measuring the age of a message relative to the
    first message in the stream.

    """
    headers = dict(msg.headers())
    return deserialize_time_offset(headers["alertsim-time-offset"])


def deserialize_time_offset(header_value):
    """Decode a timedeltas which was encoded with `serialize_time_offset`.

    """
    return datetime.timedelta(seconds=float(header_value))


ConfluentWireFormatHeader = struct.Struct(">bi")


def serialize_confluent_wire_header(schema_version):
    # Use the Confluent Wire format, described here:
    # https://docs.confluent.io/current/schema-registry/serdes-develop/index.html#wire-format
    return ConfluentWireFormatHeader.pack(0, schema_version)


def deserialize_confluent_wire_header(raw):
    return ConfluentWireFormatHeader.unpack(raw)


def serialize_alert(schema, alert):
    """ Serialize an alert to a byte sequence. """
    buf = io.BytesIO()
    # TODO: Use a proper schema versioning system
    buf.write(serialize_confluent_wire_header(0))
    fastavro.schemaless_writer(buf, schema, alert)
    return buf.getvalue()


def deserialize_alert(schema, alert_bytes):
    header_bytes = alert_bytes[:5]
    magic, version = deserialize_confluent_wire_header(header_bytes)
    assert magic == 0
    assert version == 0
    content_bytes = io.BytesIO(alert_bytes[5:])
    return fastavro.schemaless_reader(content_bytes, schema)


def alert_time(alert):
    """Convert an alert's midPointTai field from MJD float into a datetime.

    """
    raw = alert["diaSource"]["midPointTai"]
    return astropy.time.Time(raw, format="mjd").to_datetime()
