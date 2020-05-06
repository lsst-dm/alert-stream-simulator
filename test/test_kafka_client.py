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
import time
import unittest
import pytest
import logging

import confluent_kafka

from streamsim import _kafka


@pytest.mark.integration_test
class TestKafkaClientIntegration(unittest.TestCase):
    BROKER_URL = "localhost:9092"
    CLIENT_ID = "test"

    def test_create_topic(self):
        test_topic_name = "test_topic"

        client = _kafka._KafkaClient(self.BROKER_URL, self.CLIENT_ID)
        try:
            client.create_topic(test_topic_name, 12)
            time.sleep(0.5)
            topic_description = client.describe_topic(test_topic_name)
            self.assertEqual(topic_description.topic, test_topic_name)
            self.assertEqual(len(topic_description.partitions), 12)
        finally:
            client.delete_topic(test_topic_name)
            client.close()

    def test_force_create(self):
        client = _kafka._KafkaClient(self.BROKER_URL, self.CLIENT_ID)
        try:
            preexisting_topic = "test_topic_existing"
            new_topic = "test_topic_new"
            # Create a new topic
            client.create_topic(preexisting_topic)
            # Try to create without forcing recreation: it should fail
            with self.assertRaises(confluent_kafka.KafkaException):
                client.create_topic(preexisting_topic)

            # Now try to overwrite: it should work
            client.create_topic(preexisting_topic, delete_if_exists=True)

            # Finally, try to create a totally new topic with delete_if_exists
            client.create_topic(new_topic, delete_if_exists=True)
        finally:
            try:
                client.delete_topic(preexisting_topic)
            except Exception as e:
                logging.warn(f"unable to delete topic {preexisting_topic}: {e}")
            try:
                client.delete_topic(new_topic)
            except Exception as e:
                logging.warn(f"unable to delete topic {new_topic}: {e}")
            client.close()
