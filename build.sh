#!/bin/bash

docker rmi daemony:latest
export GOOS=linux
go build
docker build -t daemony:latest .

