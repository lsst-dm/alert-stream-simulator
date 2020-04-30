import asyncio
import io
import json
import logging
import functools

import avro.io
import confluent_kafka
import fastavro.write


logger = logging.getLogger("rubin-alert-sim.sender")

class FastavroAlertProducer(confluent_kafka.Producer):
    def __init__(self, broker_url, topic, schema, compression="deflate", pool=None):
        kafka_config = {
            "bootstrap.servers": broker_url,
            "message.max.bytes": 10_000_000,
        }
        self.topic = topic
        self.schema = dict(schema)
        self.compression = compression
        self.pool = pool
        super().__init__(kafka_config)

    def serialize(self, alert):
        buffer = io.BytesIO()
        fastavro.write.writer(fo=buffer, codec=self.compression, schema=self.schema, records=[alert])
        alert_bytes = buffer.getvalue()
        return alert_bytes

    async def pool_serialize(self, alert):
        loop = asyncio.get_running_loop()
        return await loop.run_in_executor(self.pool, functools.partial(Serializer(self.compression, self.schema).serialize, alert=alert))

    async def send_alert(self, alert):
        """Send a single alert packet to the Kafka topic.

        Parameters
        ----------
        alert : `dict`
            An alert that can be serialized into avro.
        """
        if self.pool is not None:
            alert_bytes = await self.pool_serialize(alert)
        else:
            alert_bytes = self.serialize(alert)
        logger.debug("sending alert (%d bytes)", len(alert_bytes))
        attempt_number = 0
        while attempt_number < 5:
            try:
                self.produce(topic=self.topic, value=alert_bytes)
                break
            except BufferError:
                self.poll(1)
                attempt_number += 1

class Serializer(object):
    def __init__(self, compression, schema):
        self.compression = compression
        self.schema = schema

    def serialize(self, alert):
        buffer = io.BytesIO()
        fastavro.write.writer(fo=buffer, codec=self.compression, schema=self.schema, records=[alert])
        alert_bytes = buffer.getvalue()
        return alert_bytes


class AlertProducer(confluent_kafka.Producer):
    """A client capable of sending alert data to a Kafka broker.

    """
    def __init__(self, broker_url, topic, schema):
        kafka_config = {
            "bootstrap.servers": broker_url,
            "message.max.bytes": 10_000_000,
        }
        self.topic = topic
        self.writer = avro.io.DatumWriter(writer_schema=schema)
        super().__init__(kafka_config)

    async def send_alert(self, alert):
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
        attempt_number = 0
        while attempt_number < 5:
            try:
                self.produce(topic=self.topic, value=alert_bytes)
                break
            except BufferError:
                self.poll(1)
                attempt_number += 1
