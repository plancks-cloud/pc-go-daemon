#!/bin/bash

docker stack rm pc
./build.sh
docker stack deploy pc --compose-file docker-compose-justin-mac.yml