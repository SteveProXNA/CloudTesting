##### README.md
###### 15/09/2021
```
01. Locally
go mod init testwebapi
touch main.go
go run main.go
curl http://localhost:8081/test
```
```
02. Docker
Ctrl + Shift + P | Add Docker Files to Workspace
Go | 8081 | No
Right click Dockerfile | Build image... | 01example:latest
Right click 01example:latest | Run interactive
curl http://localhost:8081/test
```
```
03. Kubernetes
minikube start
minikube docker-env
eval $(minikube -p minikube docker-env)
docker build -t stevepro/testwebapi:1.0 .
kubectl apply -f Kubernetes.yaml
minikube service testwebapi-service --url
curl http://192.168.49.2:30799/test
kubectl delete -f Kubernetes.yaml
minikube stop
```
