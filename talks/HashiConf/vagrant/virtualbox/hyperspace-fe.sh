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

# exit on exception
set -e

# install hyperspace
#git clone https://github.com/kenpratt/hyperspace.git
cd hyperspace
sed -i "s|/home/hyperspace|$HOME|" etc/nginx.conf
sed -i '3d' etc/nginx.conf
sed -i '28,33d' etc/nginx.conf

# install nginx
sudo apt-get install -y nginx
sudo rm /etc/nginx/sites-enabled/default
sudo cp $(pwd)/etc/nginx.conf /etc/nginx/sites-enabled/default
sudo service nginx restart

# install consul template
wget -q https://github.com/hashicorp/consul-template/releases/download/v0.6.5/consul-template_0.6.5_linux_amd64.tar.gz
tar xzf consul-template_0.6.5_linux_amd64.tar.gz
sudo mv consul-template_0.6.5_linux_amd64/consul-template /usr/bin
sudo rmdir consul-template_0.6.5_linux_amd64
