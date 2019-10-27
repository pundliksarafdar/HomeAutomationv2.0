docker stop homeappl
docker rmi go-docker
docker build -t go-docker .
docker run -p 8080:8080 --name homeappl go-docker