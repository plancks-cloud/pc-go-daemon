#!/bin/bash

docker stop daemony
docker rm daemony
docker rmi daemony:latest

export GOOS=linux

go build

docker build -t daemony:latest .

