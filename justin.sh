#!/bin/bash

docker stack rm pc

docker ps
echo "Waiting..."
sleep 6s
docker ps
sleep 1s

./build.sh
docker stack deploy pc --compose-file docker-compose-justin-mac.yml

echo "Getting daemony container ID"
sleep 2s

DAEMONID=`docker ps | grep daemony | grep -o "^[0-9a-zA-Z]" |  tr -d '\n' | awk '{$1=$1};1'`

echo "Running daemony: ${DAEMONID}"
echo "Connecting to logs:"
docker logs -f ${DAEMONID}