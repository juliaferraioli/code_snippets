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


#Naming and locating your [persistent disk](https://cloud.google.com/compute/docs/disks/persistent-disks)

If you're using persistent disks as data volumes on Google Compute Engine, you'll need to locate them in order to use them. All volumes will show up in `/dev/disk/by-id/`, but in order to use the name you give the persistent disk, you need to pass that information to `gcloud` when you attach it to the instance. Here's how.

*This assumes that you have an instance named `hello-pd`.*

Create the disk on the development machine:

```bash
dev $ gcloud compute disks create hello-pd-disk --zone us-central1-f
```

Attach the disk to the instance:

```bash
dev $ gcloud compute instances attach-disk hello-pd 
            --disk hello-pd-disk 
            --device-name hello-pd-disk 
            --zone us-central1-f
```

List the attached disks on the instance:

```bash
hello-pd $ ls -l /dev/disk/by-id/google-*
```

```bash
lrwxrwxrwx 1 root root  9 Jul 22 21:58 /dev/disk/by-id/google-hello-pd-disk -> ../../sdb
```
