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
import os.path
import json
import io

import pytest
import astropy.time
import fastavro

from streamsim import creator, _kafka


@pytest.mark.integration_test
class TestCreateIntegration(unittest.TestCase):

    @staticmethod
    def mock_alert(alert_id, timestamp):
        """Generate a minimal mock alert. Timestamp should be an ISO time string."""
        return {
            "alertId": alert_id,
            "diaSource": {
                "midPointTai": astropy.time.Time(timestamp).mjd,
                # Below are all the required fields. Set them to zero.
                "diaSourceId": 0,
                "ccdVisitId": 0,
                "filterName": "",
                "programId": 0,
                "ra": 0,
                "decl": 0,
                "x": 0,
                "y": 0,
                "apFlux": 0,
                "apFluxErr": 0,
                "snr": 0,
                "psFlux": 0,
                "psFluxErr": 0,
                "flags": 0,
            }
        }

    def mock_alert_file(self, alerts):
        mock_file = io.BytesIO()
        fastavro.writer(mock_file, self.schema, alerts)
        mock_file.seek(0)
        return mock_file

    def setUp(self):
        test_dir = os.path.dirname(os.path.abspath(__file__))
        fixtures_dir = os.path.join(test_dir, "fixtures")
        schema_path = os.path.join(fixtures_dir, "alert_schema.avsc")
        with open(schema_path, "r") as schema_file:
            self.schema = json.load(schema_file)

    def test_create(self):
        topic_name = "TestCreateIntegration.test_create"
        alerts = [
            self.mock_alert(1, "2020-01-01T00:00:00"),
            self.mock_alert(2, "2020-01-01T00:00:00"),
            self.mock_alert(3, "2020-01-01T00:00:01"),
            self.mock_alert(4, "2020-01-01T00:00:02"),
            self.mock_alert(5, "2020-01-01T00:00:02"),
        ]
        alert_file = self.mock_alert_file(alerts)

        n_sent = creator.create("localhost:9092", topic_name, alert_file, 5.0, True)
        self.assertEqual(n_sent, len(alerts))

        kc = _kafka._KafkaClient("localhost:9092")
        kc.subscribe(topic_name, 15)

        msgs = kc.consumer.consume(n_sent, timeout=20)
        self.assertAlmostEqual(len(msgs), n_sent)
        expected_offsets = [0.0, 0.0, 1.0, 2.0, 2.0]
        have_offsets = [creator.get_message_time_offset(m).total_seconds() for m in msgs]
        self.assertEqual(have_offsets, expected_offsets)
