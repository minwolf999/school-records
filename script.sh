docker build -t schoolfile .
docker run --rm -p 8080:8080 --name server schoolfile
docker stop server
docker rmi -f schoolfile