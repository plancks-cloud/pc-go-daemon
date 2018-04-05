#!/bin/bash

docker stack rm pc


docker ps
echo "Waiting..."
sleep 6s
docker ps
sleep 1s

./build.sh
docker stack deploy pc --compose-file docker-compose-kevin-mac.yml