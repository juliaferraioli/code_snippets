<!--- Copyright 2015 Google
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.--->

#Containers & Compute Engine: creating Minecraft Roulette with Kubernetes

This repository accompanies the [Containers & Compute Engine: creating Minecraft Roulette with Kubernetes](http://www.blog.juliaferraioli.com/2015/11/containerized-minecraft-roulette.html) blog entry, published on 9 Nov 2015. Please follow along in the blog entry, as it provides context and explanation for this repository's contents. Files and commands may require updating as future versions of `docker`, `gcloud`, and `Kubernetes` are released.

All commands in this file are executed on your local or development machine.

## Set up

Update and configure `gcloud`:

```bash
local $ gcloud components update
local $ gcloud config list
local $ gcloud config set project &lt;project id&gt;
```

Install Kubernetes and bring up a cluster:

```bash
local $ curl https://goo.gl/6y01p1 -o k8s.tar.gz
local $ tar -xvf k8s.tar.gz
local $ cd kubernetes
local $ cluster/kube-up.sh
```

Verify cluster information:

```bash
local $ kubectl cluster-info
local $ kubectl get nodes
```

## Create Minecraft server pod

Define the `minecraft-server.yaml` file:

```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    name: minecraft-server
  name: minecraft-server
spec:
  containers:
  - name: minecraft-server
    image: gcr.io/&lt;project id&gt;/ftb:v3
    env:
    - name: EULA
      value: "true"
    ports:
    - containerPort: 25565
      hostPort: 25565
```

Tell `kubectl` to create it from the file:

```bash
local $ kubectl create -f minecraft-server.yaml
```

Verify that it is running:

```bash
local $ kubectl get pods
```

Describe the pod:

```bash
local $ kubectl describe pod minecraft-server
```

Find the IP address of the pod using `gcloud`:

```bash
local $ gcloud compute instances list &lt;node name&gt;
```

## Create load balancer service

Define the `minecraft-service.yaml` file:

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    name: minecraft-lb
  name: minecraft-lb
spec:
  ports:
  - port: 25565
    targetPort: 25565
  selector:
    name: minecraft-server
  type: LoadBalancer
```

Tell `kubectl` to create it from the file:

```bash
local $ kubectl create -f minecraft-service.yaml
```

Get information about the services running:

```bash
local $ kubectl get services
```

## Create replication controller

Delete the running `minecraft-server` pod:

```bash
local $ kubectl delete -f minecraft-server.yaml
```

Define the `minecraft-rc.yaml` file:

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: minecraft-rc
  labels:
    name: minecraft-rc
spec:
  replicas: 3
  selector:
    name: minecraft-server
  template:
    metadata:
      labels:
        name: minecraft-server
      name: minecraft-server
    spec:
      containers:
      - name: minecraft-server
        image: gcr.io/&lt;project id&gt;/ftb:v3
        env:
        - name: EULA
          value: "true"
        ports:
        - containerPort: 25565
          hostPort: 25565
```

Tell `kubectl` to create it from the file:

```bash
local $ kubectl create -f minecraft-rc.yaml
```

Get information about the running pods:

```bash
local $ kubectl get pods
```

## Stress testing the cluster

Try deleting a pod from the cluster:

```bash
local $ kubectl delete pod &lt;pod name&gt;
```

See that Kubernetes has restarted a pod for you:

```bash
local $ kubectl get pods
```

Scale up the number of pods:

```bash
local $ kubectl scale --replicas=4 replicationcontrollers minecraft-rc
local $ kubectl get pods
```

Scale down the number of pods:
```bash
local $ kubectl scale --replicas=2 replicationcontrollers minecraft-rc
local $ kubectl get pods
```
