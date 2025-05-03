## Amazon EKS
02-Jun-2025
<br />
Instructions for [Cloud Setup Cheat Sheet II](https://steveproxna.blogspot.com/2025/06/cloud-setup-cheat-sheet-ii.html) blog post.
<br /><br />

Amazon provides Elastic Kubernetes Service as a fully managed Kubernetes container orchestration service.
<br />
Follow all instructions below in order to provision a Kubernetes cluster and test its functionality end-to-end.

#### Master SSH Key
```
cd ~/.ssh
ssh-keygen -t rsa -b 4096 -N '' -f master_ssh_key
eval $(ssh-agent -s)
ssh-add master_ssh_key
```

#### eksctl
```
curl --silent --location "https://github.com/eksctl-io/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin
```

#### Pre-Requisites
```
aws sso login
```

#### Check Resources
```
aws ec2 describe-instances --query 'Reservations[*].Instances[*].InstanceId' --output table
aws ec2 describe-addresses --query 'Addresses[*].PublicIp' --output table
aws ec2 describe-key-pairs --query 'KeyPairs[*].KeyName' --output table
aws ec2 describe-volumes --query 'Volumes[*].VolumeId' --output table
aws ec2 describe-vpcs --query 'Vpcs[*].VpcId' --output table
aws cloudformation list-stacks --query 'StackSummaries[*].StackName' --output table  
aws cloudwatch describe-alarms --query 'MetricAlarms[*].AlarmName' --output table
aws ecr describe-repositories --query 'repositories[*].repositoryName' --output table
aws ecs list-clusters --query 'clusterArns' --output table
aws eks list-clusters --query 'clusters' --output table
aws elasticbeanstalk describe-environments --query 'Environments[*].EnvironmentName' --output table
aws elb describe-load-balancers --query 'LoadBalancerDescriptions[*].LoadBalancerName' --output table
aws elbv2 describe-load-balancers --query 'LoadBalancers[*].LoadBalancerName' --output table
aws iam list-roles --query 'Roles[*].RoleName' --output table
aws iam list-users --query 'Users[*].UserName' --output table
aws lambda list-functions --query 'Functions[*].FunctionName' --output table
aws rds describe-db-instances --query 'DBInstances[*].DBInstanceIdentifier' --output table
aws route53 list-hosted-zones --query 'HostedZones[*].Name' --output table
aws s3 ls
aws sns list-topics --query 'Topics[*].TopicArn' --output table
aws sqs list-queues --query 'QueueUrls' --output table
aws ssm describe-parameters --query 'Parameters[*].Name' --output table
```

#### Cluster YAML
```
kind: ClusterConfig
apiVersion: eksctl.io/v1alpha5

metadata:
  name: stevepro-aws-eks
  region: eu-west-1
  version: "1.27"
  tags:
    createdBy: stevepro

kubernetesNetworkConfig:
  ipFamily: IPv4

iam:
  withOIDC: true
  serviceAccounts:
  - metadata:
      name: ebs-csi-controller-sa
      namespace: kube-system
    attachPolicyARNs:
    - "arn:aws:iam::aws:policy/service-role/AmazonEBSCSIDriverPolicy"
    roleOnly: true
    roleName: stevepro-aws-eks-AmazonEKS_EBS_CSI_DriverRole

addons:
- name: aws-ebs-csi-driver
  version: v1.38.1-eksbuild.2
  serviceAccountRoleARN: arn:aws:iam::4xxxxxxxxxx8:role/stevepro-aws-eks-AmazonEKS_EBS_CSI_DriverRole
- name: vpc-cni
  version: v1.19.2-eksbuild.1
- name: coredns
  version: v1.10.1-eksbuild.18
- name: kube-proxy
  version: v1.27.16-eksbuild.14

nodeGroups:
  - name: stevepro-aws-eks
    instanceType: m5.large
    desiredCapacity: 0
    minSize: 0
    maxSize: 3
    ssh:
      allow: true
      publicKeyPath: ~/.ssh/master_ssh_key.pub   
    preBootstrapCommands:
      - "true"
```

#### Create Cluster
```
eksctl create cluster -f ~/stevepro-awseks/cluster.yaml   \
    --kubeconfig ~/stevepro-awseks/kubeconfig             \
    --verbose 5
```

#### Scale Nodegroup
```
eksctl scale nodegroup                                    \
    --cluster=stevepro-aws-eks                            \
    --name=stevepro-aws-eks                               \
    --nodes=3                                             \
    --nodes-min=0                                         \
    --nodes-max=3                                         \
    --verbose 5
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
Ref: 
~\GitHub\StevePro7\Blogger\Cloud\CloudSetupCheatSheet\CloudSetupCheatSheetI\archive\CloudSetupCheatSheetNotes
```
kubectl get po -o wide
cd ~/.ssh
ssh -i master_ssh_key ec2-user@node-ip-address
# ssh -i master_ssh_key ubuntu@node-ip-address
# ssh -i master_ssh_key root@node-ip-address
```

#### Cleanup
```
kubectl delete -f Kubernetes.yaml
kubectl delete ns test-ns
```


#### Delete Cluster
```
eksctl delete cluster                                     \
    --name=stevepro-aws-eks                               \
    --region eu-west-1                                    \
    --force
```