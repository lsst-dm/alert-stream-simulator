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
