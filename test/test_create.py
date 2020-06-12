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
import unittest

import pytest

from lsst.alert.stream import serialization
from streamsim import creator, _kafka, timestamps
from test import testutils


@pytest.mark.integration_test
class TestCreateIntegration(unittest.TestCase):
    def test_create(self):
        topic_name = "TestCreateIntegration.test_create"
        alerts = [
            testutils.mock_alert(1, "2020-01-01T00:00:00"),
            testutils.mock_alert(2, "2020-01-01T00:00:00"),
            testutils.mock_alert(3, "2020-01-01T00:00:01"),
            testutils.mock_alert(4, "2020-01-01T00:00:02"),
            testutils.mock_alert(5, "2020-01-01T00:00:02"),
        ]
        alert_file = testutils.mock_alert_file(alerts)

        n_sent = creator.create("localhost:9092", topic_name, alert_file, 5.0, True)
        self.assertEqual(n_sent, len(alerts))

        kc = _kafka._KafkaClient("localhost:9092")
        kc.subscribe(topic_name, 15)

        msgs = kc.consumer.consume(n_sent, timeout=20)
        self.assertEqual(len(msgs), n_sent)

        # Messages should have offset headers attached
        expected_offsets = [0.0, 0.0, 1.0, 2.0, 2.0]
        have_offsets = [timestamps.get_message_time_offset(m).total_seconds() for m in msgs]
        self.assertEqual(have_offsets, expected_offsets)

        # Messages should be deserializable, but not necessarily in the
        # original order
        have = [serialization.deserialize_alert(m.value()) for m in msgs]
        want_by_id = {alert["alertId"]: alert for alert in alerts}
        for alert in have:
            want_timestamp = want_by_id[alert["alertId"]]["diaSource"]["midPointTai"]
            have_timestamp = alert["diaSource"]["midPointTai"]
            self.assertEqual(have_timestamp, want_timestamp)

        kc.close()
