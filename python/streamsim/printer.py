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

from streamsim import _kafka
from lsst.alert.stream import serialization


logger = logging.getLogger("rubin-alert-sim.printer")


def print_stream(broker, topic, tls_config=None):
    """Prints the contents of a stream in real time.

    Parameters
    ----------
    broker : `str`
        The URL of a Kafka broker to connect to.
    topic : `str`
        The name of a Kafka topic to read.
    tls_config : `streamsim._kafka.TLSConfig` or `None`
        If not None, then a configuration bundle for TLS auth when connecting
        to the broker.
    """
    kafka_client = _kafka._KafkaClient(
        broker, enable_eof=False, tls_config=tls_config,
    )
    kafka_client.consumer.subscribe([topic])
    time.sleep(1)
    while True:
        messages = kafka_client.consumer.consume(100, 0.1)
        for msg in messages:
            alert = serialization.deserialize_alert(msg.value())
            now = str(datetime.datetime.now())
            alertID = alert['alertId']
            historySize = len(alert.get('prvDiaSources') or [])
            size = len(msg.value())
            print(f"{now}  id={alertID}  len(history)={historySize}  size={size}")
