#!/bin/bash

set -euo pipefail

# Input through environment variables:
#   INFLUXDB_DB: Name of the database to create.


influxql() {
    influx -host localhost \
           -port 8086 \
           -execute "$@"
}

# Turn on the database with a local listener, just so we can create users.
echo "Starting InfluxDB in initialization mode."
INFLUXDB_HTTP_BIND_ADDRESS=localhost:8086 INFLUXDB_HTTP_HTTPS_ENABLED=false influxd "$@" &
pid="$!"

for i in {30..0}; do
    echo "Waiting for DB to initialize."
    sleep 0.5
    if $(influx -host localhost \
       -port 8086 \
       -execute "CREATE DATABASE broker_simulator;"); then break; fi
done

echo "InfluxDB initialization complete."
wait "$pid"
