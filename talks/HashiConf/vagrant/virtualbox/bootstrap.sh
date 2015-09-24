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

echo 'deb http://us.archive.ubuntu.com/ubuntu/ trusty main universe' > /tmp/s.l
echo 'deb http://us.archive.ubuntu.com/ubuntu/ trusty-security main universe' >> /tmp/s.l
echo 'deb http://us.archive.ubuntu.com/ubuntu/ trusty-updates main universe' >> /tmp/s.l
sudo cp /tmp/s.l /etc/apt/source.list
sudo apt-get update -y
sudo apt-get install -y git vim tmux build-essential
