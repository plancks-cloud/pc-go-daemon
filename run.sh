#!/bin/bash

# docker stop pc-go-daemon
# docker rm pc-go-daemon
# docker rmi pc-go-daemon:latest

# export GOOS=linux

# go build

# docker build -t pc-go-daemon:latest .
# docker run --name pc-go-daemon -v /var/run/docker.sock:/var/run/docker.sock --publish 8080:8080 pc-go-daemon:latest

./build.sh
docker stack deploy pc --compose-file docker-compose.yml