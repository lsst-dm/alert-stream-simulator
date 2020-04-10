#!/bin/bash

set -euo pipefail

CONTAINER_NAME="broker-simulator_influxdb_1"
VOLUME_NAME=$(docker inspect $CONTAINER_NAME | jq -r '.[0]["Mounts"][] | select( .Destination == "/var/lib/influxdb").Name')

docker container rm "$CONTAINER_NAME"
docker volume rm "$VOLUME_NAME"
