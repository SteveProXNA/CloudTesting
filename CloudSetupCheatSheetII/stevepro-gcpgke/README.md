## Google GKE
02-Jun-2025
<br />
Instructions for [Cloud Setup Cheat Sheet II](https://steveproxna.blogspot.com/2025/06/cloud-setup-cheat-sheet-ii.html) blog post.
<br /><br />

Google provides the Google Kubernetes Engine as fully managed Kubernetes container orchestration service.
<br />
Follow all instructions below in order to provision a Kubernetes cluster and test its functionality end-to-end.

#### Pre-Requisites
```
gcloud auth login
gcloud auth application-default login
gcloud auth configure-docker
gcloud config set project SteveProProject
```

#### Check Resources
```
gcloud compute instances list
gcloud compute disks list
gcloud compute forwarding-rules list
gcloud compute firewall-rules list
gcloud compute addresses list
gcloud container clusters list
```

#### Create Cluster
```
gcloud container clusters create stevepro-gcp-gke               \
    --project=steveproproject                                   \
    --zone=europe-west1-b                                       \
    --machine-type=e2-standard-2                                \
    --disk-type pd-standard                                     \
    --cluster-version=1.30.10-gke.1070000                       \
    --num-nodes=3                                               \
    --network=default                                           \
    --create-subnetwork=name=stevepro-gcp-gke-subnet,range=/28  \
    --enable-ip-alias                                           \
    --enable-intra-node-visibility                              \
    --logging=NONE                                              \
    --monitoring=NONE                                           \
    --enable-network-policy                                     \
    --labels=prefix=stevepro-gcp-gke,created-by=${USER}         \
    --no-enable-managed-prometheus                              \
    --quiet --verbosity debug
```

#### Get Credentials
```
gcloud container clusters get-credentials stevepro-gcp-gke      \
    --zone=europe-west1-b                                       \
    --quiet --verbosity debug
```

#### Deploy Test
```
kubectl create ns test-ns
kubectl config set-context --current --namespace=test-ns
kubectl apply -f Kubernetes.yaml
kubectl port-forward service/flask-api-service 8080:80
curl http://localhost:8080
```

#### Output
```
Hello World (Python)!
```

#### Shell into Node
```
mkdir -p ~/GitHub/luksa
cd ~/GitHub/luksa
git clone https://github.com/luksa/kubectl-plugins.git
cd kubectl-plugins
chmod +x kubectl-ssh
kubectl get nodes
./kubectl-ssh node gke-stevepro-gcp-gke-default-pool-0b4ca8ca-sjpj
```

#### Cleanup
```
kubectl delete -f Kubernetes.yaml
kubectl delete ns test-ns
```

#### Delete Cluster
```
gcloud container clusters delete stevepro-gcp-gke               \
    --zone europe-west1-b                                       \
    --quiet --verbosity debug
```