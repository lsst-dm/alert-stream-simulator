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
import astropy.time

from streamsim import creator, player, _kafka
from test import testutils


@pytest.mark.integration_test
class TestPlayIntegration(unittest.TestCase):
    @staticmethod
    def mock_alert(alert_id, timestamp):
        """Generate a minimal mock alert.
        Timestamp should be an ISO time string.

        """
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

    def test_play(self):
        src_topic_name = "TestPlayIntegration.test_play_src"
        dst_topic_name = "TestPlayIntegration.test_play_dst"
        broker_url = "localhost:9092"
        alerts = [
            testutils.mock_alert(1, "2020-01-01T00:00:00"),
            testutils.mock_alert(2, "2020-01-01T00:00:00"),
            testutils.mock_alert(3, "2020-01-01T00:00:01"),
            testutils.mock_alert(4, "2020-01-01T00:00:02"),
            testutils.mock_alert(5, "2020-01-01T00:00:02"),
        ]
        alert_file = testutils.mock_alert_file(alerts)

        # Create a stream and populate it.
        n_sent = creator.create(broker_url, src_topic_name, alert_file, 5.0, True)
        self.assertEqual(n_sent, len(alerts))

        # Replay the stream into a new topic.
        player.play(broker_url, src_topic_name, dst_topic_name, 1, True)

        # Check that the new topic contains all the messages from the stream.
        kc = _kafka._KafkaClient(broker_url)
        kc.subscribe(dst_topic_name, 15)
        msgs = kc.consumer.consume(n_sent, timeout=20)
        self.assertEqual(len(msgs), n_sent)
