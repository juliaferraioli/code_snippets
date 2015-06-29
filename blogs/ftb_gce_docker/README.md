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

This repository accompanies the [Running a Minecraft server on Google Compute Engine with Docker](http://www.blog.juliaferraioli.com/2015/06/running-minecraft-server-on-google.html) blog entry, published on $DATE. Please follow along in the blog entry, as it provides context and explanation for this repository's contents. Files and commands may require updating as future versions of `docker` and `gcloud` are released.

Executed on local machine:

```bash
$ curl https://sdk.cloud.google.com | bash

$ gcloud auth login

$ glcoud compute ssh minecraft-docker --zone  us-central1-f
```

Executed on development environment:
```bash
$ sudo gcloud components update

$ sudo docker ps
```
Create [`Dockerfile`](https://github.com/juliaferraioli/code_snippets/blob/master/blogs/ftb_gce_docker/Dockerfile) and [`CheckEula.sh`](https://github.com/juliaferraioli/code_snippets/blob/master/blogs/ftb_gce_docker/CheckEula.sh) on development environment.

Executed in the directory with created files.

```bash
$ sudo docker build -t <docker user>/ftb .

$ sudo docker images

$ sudo docker tag <image id> gcr.io/<project id>/ftb

$ sudo gcloud docker push gcr.io/<project id>/ftb
```
Executed on  local machine:

```bash
$ gcloud compute firewall-rules create minecraft-port \
		--description "Allow connecting over port 25565" \
        --allow tcp:25565

$ gcloud compute instances create minecraft-server \
		--image container-vm \
		--zone  us-central1-f \
		--machine-type n1-standard-1
		--scopes storage-rw

$ gcloud compute instances list

$ gcloud compute ssh minecraft-server
```

Executed on `minecraft-server`:

```bash
$ sudo gcloud components update

$ sudo gcloud docker pull gcr.io/<your project id>/ftb

$ sudo docker run -p 25565:25565 -e EULA=true -d gcr.io/<your project id>/ftb
```
