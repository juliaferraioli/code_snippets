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

#Saving the world: using persistent storage with a containerized Minecraft server

This repository accompanies the [Running a Minecraft server on Google Compute Engine with Docker](http://www.blog.juliaferraioli.com/2015/07/saving-world-using-persistent-storage.html) blog entry, published on $DATE. It builds upon the [first entry in the series](http://www.blog.juliaferraioli.com/2015/06/running-minecraft-server-on-google.html), which covers building and running the Minecraft server container. Please follow along in the blog entry, as it provides context and explanation for this repository's contents. Files and commands may require updating as future versions of `docker` and `gcloud` are released.

## Saving the world on a persistent disk

Executed on local machine:

```bash
$ gcloud compute disks create minecraft-world-disk1 --zone us-central1-f

$ gcloud compute disks list
```

Create the disk, executed on local machine:

```bash
local $ gcloud compute instances attach-disk minecraft-server \
                       --disk minecraft-world-disk1 \
                       --device-name minecraft-world-disk1 \
                       --zone us-central1-f
```

Executed on `minecraft-server`:

```bash
$ sudo mkdir -p /minecraft-world

$ sudo /usr/share/google/safe_format_and_mount \
                            -m "mkfs.ext4 -F" \
                            /dev/disk/by-id/google-minecraft-world-disk1 \
                            /minecraft-world
```

Run container, executed on `minecraft-server`:

```bash
$ sudo docker run -p 25565:25565 \
                  -v /minecraft-world:/opt/ftb/world \
                  -e EULA=true -d gcr.io/<your project id>/ftb
```

Look inside the persistent disk on `minecraft-server`:

```bash
$ ls -alh /minecraft-world
```

Test the persistence of the world data, executed on `minecraft-server`:

```bash
$ sudo docker ps

$ sudo docker stop <container name/ID>

$ sudo docker rm <container name/ID>

$ sudo docker run -p 25565:25565 \
                  -v /minecraft-world:/opt/ftb/world \
                  -e EULA=true \
                  -d gcr.io/<your project id>/ftb
```

Mount a disk with existing world data, executed on `minecraft-server`:

```bash
$ sudo mkdir -p /minecraft-world

$ sudo mount /dev/disk/by-id/google-minecraft-world-disk1 \
             /minecraft-world
```

## Snapshotting

Manually create snapshot, executed on `minecraft-server`:

```bash
$ sudo docker stop <container ID>

$ sudo sync

$ sudo fsfreeze -f /minecraft-world

$ sudo gcloud compute disks snapshot minecraft-world-disk1 \
                      --snapshot-name \
                      "minecraft-world-disk1-`date +%m-%d-%y`"

$ sudo fsfreeze -u /minecraft-world                      
```

Recreate `minecraft-server` with new scope, executed on local machine:

```bash
$ gcloud compute instances create minecraft-server \
                 --image container-vm \
                 --zone  us-central1-f \
                 --machine-type n1-standard-1 \
                 --scopes storage-rw,compute-rw
```

Create automated backups with [backup script](), executed on `minecraft-server`:

```bash
$ crontab -e
```

Add line to file:

```
00 04 * * * /<path to backup script>/backup-world.sh
```

Restoring from snapshot, executed on local machine:

```bash
$ gcloud compute disks create minecraft-world-restored \
                 --source-snapshot=<snapshot name> \
                 --zone us-central1-f
```

Attach disk to instance, executed on local machine:

```bash
$ gcloud compute instances attach-disk minecraft-server \
                 --disk minecraft-world-restored \
                 --device-name minecraft-world-disk1 \
                 --zone us-central1-f
```
