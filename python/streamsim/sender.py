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
import json
import logging

import avro.io
import confluent_kafka


logger = logging.getLogger("rubin-alert-sim.sender")


class AlertProducer(confluent_kafka.Producer):
    """A client capable of sending alert data to a Kafka broker.

    """
    def __init__(self, broker_url, topic, schema):
        kafka_config = {
            "bootstrap.servers": broker_url,
        }
        self.topic = topic
        self.writer = avro.io.DatumWriter(writer_schema=schema)
        super().__init__(kafka_config)

    def send_alert(self, alert):
        """Send a single alert packet to the Kafka topic.

        Parameters
        ----------
        alert : `dict`
            An alert that can be serialized into avro.
        """
        buffer = io.BytesIO()
        self.writer.write(alert, avro.io.BinaryEncoder(buffer))
        alert_bytes = buffer.getvalue()
        logger.debug("sending alert (%d bytes)", len(alert_bytes))
        self.produce(topic=self.topic, value=alert_bytes)
