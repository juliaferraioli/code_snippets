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

#!/bin/bash                                                                                                        
if [ -a /opt/ftb/eula.txt ] && grep -qi true /opt/ftb/eula.txt && echo $EULA | grep -qi true ; then
  echo "EULA already accepted!"
elif  echo $EULA | grep -qi true ; then
  echo "EULA accepted by docker run -e"
  echo "# EULA accepted on $(date) by docker run -e" > /opt/ftb/eula.txt
  echo "eula=$EULA" >> /opt/ftb/eula.txt
else
  echo ""
  echo "To accept the Minecraft EULA hosted at"
  echo "  https://account.mojang.com/documents/minecraft_eula"
  echo "please add ' -e EULA=TRUE' immediately after 'docker run'"
  echo ""
  > /opt/ftb/eula.txt
  exit 1
fi
