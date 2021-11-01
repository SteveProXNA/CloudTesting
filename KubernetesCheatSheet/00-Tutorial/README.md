##### README.md
###### 15/09/2021
```
01. Hello Minikube
Launch terminal
minikube start
minikube dashboard
minikube ip
minikube ssh
minikube stop
minikube delete
```
```
02. Create Deployment
minikube start
minikube dashboard
kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
kubectl get deployments
kubectl get pods
kubectl get events
kubectl config view
```
```
03. Create Service
kubectl expose deployment hello-node --type=LoadBalancer --port=8080
kubectl get services
minikube service hello-node
```
```
04. Enable Addons
minikube addons list
minikube addons enable metrics-server
kubectl get pod,svc -n kube-system
minikube addons disable metrics-server
```
ADDITIONAL
```
05. Exec into Pod
kubectl attach hello-node-7567d9fdc9-mblh7 -i
kubectl exec hello-node-7567d9fdc9-mblh7 -- ls /
kubectl exec --it hello-node-7567d9fdc9-mblh7 -- /bin/sh
kubectl exec --stdin --tty hello-node-7567d9fdc9-mblh7 -- /bin/sh
```
```
06. Tail logs
kubectl logs -f hello-node-7567d9fdc9-mblh7
minikube stop
minikube delete
```
