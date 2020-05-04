import time
import logging

import confluent_kafka


logger = logging.getLogger("rubin-alert-sim.replayer")


class Replayer(object):
    def __init__(self, broker, src_topic, dst_topic, timeout, block_size):
        producer_config = {
            "bootstrap.servers": broker,
            "message.max.bytes": 10_000_000,
            "socket.timeout.ms": 10000,
            "error_cb": logger.error,
            "throttle_cb": logger.warn,
            "on_delivery": lambda err, msg: logger.error(err) if err is not None else None,
            "auto.offset.reset": "earliest",
        }
        self.producer = confluent_kafka.Producer(producer_config)

        consumer_config = {
            "bootstrap.servers": broker,
            "session.timeout.ms": 6000,
            "socket.timeout.ms": 10000,
            "fetch.message.max.bytes": block_size*500_000,
            "error_cb": logger.error,
            "throttle_cb": logger.warn,
            "on_commit": lambda err, commits: logger.error(err) if err is not None else None,
            'default.topic.config': {
                'auto.offset.reset': "earliest",
            },
            "message.max.bytes": 10_000_000,
            "enable.partition.eof": True,
            "group.id": f"{src_topic}__to__{dst_topic}",
            "auto.offset.reset": "earliest"
        }
        self.consumer = confluent_kafka.Consumer(consumer_config)
        logger.debug(f"subscribing to {src_topic}")
        tp = confluent_kafka.TopicPartition(
            topic=src_topic,
            partition=0,
            offset=confluent_kafka.OFFSET_BEGINNING,
        )
        self.consumer.assign([tp])

        self.dst_topic = dst_topic
        self.timeout = timeout

    def replay_loop(self, block_size):
        more = True
        total_sent = 0
        start = time.perf_counter()
        while more:
            n_sent, more = self.replay_n(block_size)
            total_sent += n_sent
        dur = time.perf_counter() - start
        rate = total_sent / dur
        logger.info(f"replayed {total_sent} total alerts in {dur}s ({rate}/s)")

    def flush_all(self, timeout=10):
        start = time.time()
        n_unsent = -1
        while n_unsent != 0 and (time.time() - start) < timeout:
            logger.debug(f"n_unsent={n_unsent}")
            n_unsent = self.producer.flush(self.timeout)
        if n_unsent > 0:
            raise ValueError(f"unable to flush all, n_unsent={n_unsent}")

    def replay_n(self, n):
        start = time.perf_counter()

        messages = self.consumer.consume(num_messages=n, timeout=self.timeout)
        logger.debug(f"got {len(messages)} messages")
        more = True
        for m in messages:
            logger.debug(f"sending {len(m.value())} size msg")
            err = m.error()
            if err is not None:
                if err.code() == confluent_kafka.KafkaError._PARTITION_EOF:
                    more = False
                    return 0, False
                else:
                    raise(err)
            self.producer.produce(
                topic=self.dst_topic,
                partition=m.offset() % 2,
                value=m.value(),
                key=m.key(),
                headers=m.headers(),
            )

        self.flush_all()
        # Track stats
        n_sent = len(messages)
        dur = time.perf_counter() - start
        rate = n_sent / dur
        logger.info(f"replayed {n_sent} alerts in {dur}s ({rate}/s)")

        return n_sent, more

    def close(self):
        self.consumer.close()
