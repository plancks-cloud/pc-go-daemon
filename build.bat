
docker stop daemony
docker rm daemony
docker rmi daemony:latest
set GOOS=linux
go build
docker build -t daemony:latest .

