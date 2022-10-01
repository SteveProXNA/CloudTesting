##### README.md
###### 15/11/2022

Reference: https://medium.com/lightbaseio/web-application-firewall-in-go-feat-owasp-modsecurity-core-rule-set-3f97a26e3311
```
01. Locally
go mod init testwebapi
go get "github.com/gorilla/mux"
go mod tidy
touch main.go
go run main.go
curl http://localhost:8081/
curl http://localhost:8081/test/artists.php
curl http://localhost:8081/test/artists.php?artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user
```
```
02. Docker
Ctrl + Shift + P | Add Docker Files to Workspace
Go | 8081 | No
Right click Dockerfile | Build image... | 02example:latest
docker build --rm -f "Dockerfile" -t 02example:latest "."
Right click 01example:latest | Run interactive
docker run -d -p 8081:8081 02example:latest
curl http://localhost:8081/
curl http://localhost:8081/test/artists.php
curl http://localhost:8081/test/artists.php?artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user
docker exec -it [containerID] bash
```
```
03. Kubernetes
kind create cluster
docker build -t stevepro/testwebapi:4.0 .
kind load docker-image stevepro/testwebapi:4.0
kubectl apply -f Kubernetes.yaml
kubectl get nodes -o wide
kubectl get services
curl http://172.18.0.2:32277/
curl http://172.18.0.2:32277/test/artists.php
curl http://172.18.0.2:32277/test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user
kubectl logs -f testwebapi-787c48fd87-shc6k
kubectl exec -it testwebapi-787c48fd87-s2p6z -- bash
kubectl delete -f Kubernetes.yaml
kind delete cluster
```
