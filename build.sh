#!/bin/bash

docker rmi planckscloud/plancks-cloud:0.1
docker rmi planckscloud/plancks-cloud:latest

export GOOS=linux
go build
docker build -t planckscloud/plancks-cloud:0.1 .
docker tag planckscloud/plancks-cloud:0.1 planckscloud/plancks-cloud:latest

docker push planckscloud/plancks-cloud:0.1
docker push planckscloud/plancks-cloud:latest