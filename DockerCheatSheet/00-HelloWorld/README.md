##### README.md
###### 31/08/2021
```
01. Hello World
docker run hello-world
```
```
02. Getting Started
docker run -d -p 80:80 docker/getting-started
docker images
docker ps
docker stop $(docker ps -q)
docker rmi $(docker images -qa) --force
docker system prune -a
```
