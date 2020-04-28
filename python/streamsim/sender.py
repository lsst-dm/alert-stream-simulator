import logging
import confluent_kafka


class AlertProducer(confluent_kafka.Producer):
    """A client capable of sending alert data to a Kafka broker.

    """
    def __init__(self, broker_url, topic):
        kafka_config = {
            "bootstrap.servers": broker_url,
        }
        self.topic = topic
        super().__init__(kafka_config)

    def send_alert(self, alert):
        """Send a single alert packet to the Kafka topic.

        Parameters
        ----------
        alert : `bytes`
            A sequence of bytes that should encode an alert.
        """
        logging.debug("sending alert (%d bytes)", len(alert))
        self.produce(topic=self.topic, value=alert)
