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

# set up golang
wget -q https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz
sudo ln -s /usr/local/go/bin/go /usr/local/bin/go
sudo ln -s /usr/local/go/bin/godoc /usr/local/bin/godoc
sudo ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt

# install hyperspace
git clone https://github.com/kenpratt/hyperspace.git
cd hyperspace
export GOPATH=$(pwd)
go get github.com/gorilla/websocket
go get github.com/lucasb-eyer/go-colorful
make
sudo killall server
./server/server -port 9393

# make sure hyperspace starts on reboot
sudo touch /var/log/hyperspace.log
sudo chmod 777 /var/log/hyperspace.log
sudo chmod -R go+r $HOME
sudo chmod go+x $HOME
