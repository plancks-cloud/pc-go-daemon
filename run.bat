set export GOOS=linux
docker stop daemony
docker rm daemony
docker rmi daemony:latest
go build
docker build -t daemony:latest .
docker stack deploy pc --compose-file docker-compose.yml