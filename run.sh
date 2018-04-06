#!/bin/bash

docker stack rm pc
./build.sh
docker stack deploy pc --compose-file docker-compose.yml

echo "Getting daemony container ID"
sleep 2s

DAEMONID=`docker ps | grep daemony | grep -o "^[0-9a-zA-Z]" |  tr -d '\n' | awk '{$1=$1};1'`

echo "Running daemony: ${DAEMONID}"