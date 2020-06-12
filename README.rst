alert-stream-simulator
################

This repository holds a simulator for Rubin's alert stream, as described in
`DMTN-149`_.

.. _DMTN-149: https://dmtn-149.lsst.io/

Installation
============

Before starting, you'll need:
 - A Linux host (use a VM if you're on OSX or Windows)
 - Python 3.6+
 - Docker
 - `docker-compose <https://docs.docker.com/compose/>`_
 - curl
 - `libsnappy <https://github.com/google/snappy>`_ (`apt-get install libsnappy-dev` / `yum install snappy-devel`)

Clone the repository, activate a virtualenv (or whatever Python env isolation
mechanism you prefer), and then run ``make install``. Go get a cup of coffee while
datasets are downloaded, dependencies are installed, and Docker containers are
built.

Usage
=====

To run the broker infrastructure, run ``docker-compose up`` from the root of the
repo. This will spin up several containers; once the log output dies down, the
system should be up and running.

Once the broker is up, open a second terminal and run ``rubin-alert-sim --help``.
This is a CLI tool for interacting with the broker. There are two steps to
simulate an alert stream:

1. First, you **create** the stream, seeding the broker with data. Do this with
   ``rubin-alert-sim create-stream``. This step handles serialization, and sets
   the data rate for the stream.
2. Second, you **play** the stream with ``rubin-alert-stream play-stream``, which
   publishes the stream at the predefined data rate. This step publishes to a
   new topic.

A basic example
---------------

Let's publish a small alert stream from a file. First, make sure the broker
infastructure is up by running ``docker-compose ps`` - we expect to see "Up" for
the "State" of all containers::

  $ docker-compose ps
  Name                             Command               State   Ports
  -----------------------------------------------------------------------------------
  Alert-stream-simulator_grafana_1     /run.sh                          Up
  alert-stream-simulator_influxdb_1    /entrypoint.sh /etc/influx ...   Up
  alert-stream-simulator_jmxtrans_1    /bin/sh -c /usr/share/jmxt ...   Up
  alert-stream-simulator_kafka_1       /etc/confluent/docker/run        Up
  alert-stream-simulator_zookeeper_1   /etc/confluent/docker/run        Up

If the infrastructure is up, we can create a stream::

  $ rubin-alert-sim create-stream --dst-topic=rubin_example data/rubin_single_ccd_sample.avro
  successfully preloaded stream with 792 alerts

And now we can replay that stream::

  $ rubin-alert-sim --verbose play-stream \
      --src-topic=rubin_example \
      --dst-topic=rubin_example_stream
  INFO:rubin-alert-sim.play:sent 792 alerts in 1.67s (474.58/s)

This second command is worth looking at closely. We set the ``--dst-topi`` to
``rubin_example_stream``: this will create a new topic with that name, and will
pace the data into it at the same rate as we had set with ``create-stream``.
Connect your consumers to the ``--dst-topic`` to simulate receiving Rubin's
alerts.


Writing your own consumer
=========================

If you want to write your own consumer, you'll need a Kafka client library, and you'll need the `lsst-alert-stream` package, which provides serialization utilities.

For a Python Kafka client, we recommend `PyKafka`_, which is both simple and powerful.
For example, the following snippet will print every alert ID in the stream::

  import pykafka
  import lsst.alert.stream.serialization

  # Connect to the stream you have running in docker:
  client = pykafka.KafkaClient(hosts='localhost:9092')

  # Connect to the topic you created with 'rubin-alert-sim play-stream'
  topic = client.topics['rubin_example_stream'] # Or whatever you set as --dst-topic

  # Note that alert messages are quite large, so you must explicitly permit very large messages:
  consumer = topic.get_simple_consumer(fetch_message_max_bytes=10000000)

  for raw_msg in consumer:
      # Parse the contents into a dictionary:
      alert = lsst.alert.stream.serialization.deserialize_alert(raw_msg)

      # Do whatever you like here:
      print(alert['alertId'])


Troubleshooting
===============

KafkaException: Topic already exists
------------------------------------

While working, you might frequently find yourself re-creating and re-running
streams. Each invocation of the ``rubin-alert-sim`` creates fresh new topics, and
by default they won't overwrite existing topics. You can pass ``--force`` to
overwrite an existing topic. For example, ``rubin-alert-sim
create-stream --dst-topic=rubin_example --force data/rubin_sample.avro``.


Networking and OSX
-------------------

The provided ``docker-compose.yml`` will run all service on the host network.
This simplifies connections to the Kafka broker from the local host (and matches
`Confluent's recommendations`_), but it means that you'll need permissions to
open ports and run listeners on the host network.

Unfortunately, Docker for Mac does not support this. To run this stack on Mac,
you'll need to run a Linux Virtual Machine.

The listeners are:

 - Kafka: ``localhost:9092`` (for the stream) and ``localhost:9292`` (for JMX metrics)
 - Zookeeper: ``localhost:2181``
 - Grafana: ``localhost:3000``
 - InfluxDB: ``localhost:8086``


This will only support connections to the Kafka broker from the same host that's
running the Kafka container. If you want to connect to the broker from another
host, you'll need to make a change to these listeners. Edit the
`docker-compose.yml` file, changing all references to "``localhost``" to the IP
address of the broker. If you'd like a lot of background on this subject, `try
this blog post <https://rmoff.net/2018/08/02/kafka-listeners-explained/>`_.

.. _Confluent's Recommendations: https://docs.confluent.io/current/installation/docker/installation/index.html#considerations
.. _PyKafka: https://pykafka.readthedocs.io/en/latest/index.html
