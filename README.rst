broker-simulator
################

This repository holds a simulator for Rubin's alert stream, as described in
`DMTN-149`_.
.. _DMTN-149: https://dmtn-149.lsst.io/


Networking and OSX
==================

The provided ``docker-compose.yml`` will run all service on the host network.
This simplifies connections to the Kafka broker from the local host (and matches
`Confluent's recommendations`_), but it means that you'll need permissions to
open ports and run listeners on the host network.

Unfortunately, Docker for Mac does not support this. To run this stack on Mac,
you'll need to run a Linux Virtual Machine.

The listeners are:

 - Kafka: ``localhost:9092``
 - Zookeeper: ``localhost:32181``

This will only support connections to the Kafka broker from the same host that's
running the Kafka container. If you want to connect to the broker from another
host, you'll need to make a change to these listeners. Edit the
`docker-compose.yml` file, changing all references to "``localhost``" to the IP
address of the broker. If you'd like a lot of background on this subject, `try
this blog post <https://rmoff.net/2018/08/02/kafka-listeners-explained/>`.

.. _Confluent's Recommendations: https://docs.confluent.io/current/installation/docker/installation/index.html#considerations
