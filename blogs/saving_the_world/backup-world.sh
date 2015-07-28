# Copyright 2015 Google
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#      http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#!/bin/sh
c_id=`sudo docker ps | grep "ftb:latest" | awk '{print $1}'`
if [ "$c_id" ]; then
  sudo docker stop $c_id
fi
sudo sync
sudo fsfreeze -f /minecraft-world
sudo gcloud compute disks snapshot minecraft-world-disk1 --snapshot-name "s`date +%y-%m-%d-%H%M`-minecraft-world-disk1" --zone us-central1-f
sudo fsfreeze -u /minecraft-world
if [ "$c_id" ]; then
  sudo docker start $c_id
fi
