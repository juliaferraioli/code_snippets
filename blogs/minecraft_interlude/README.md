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

#Minecraft, Docker, Compute Engine: an interlude

This repository accompanies the [Minecraft, Docker, Compute Engine: an interlude](http://www.blog.juliaferraioli.com/2015/08/minecraft-docker-compute-engine.html) blog entry, published on $DAY August 2015. It builds upon the [prior entries in the series](http://www.blog.juliaferraioli.com/search/label/Minecraft), which covers building, running, and maintaining the Minecraft server container. Please follow along in the blog entry, as it provides context and explanation for this repository's contents. Files and commands may require updating as future versions of `docker` and `gcloud` are released.

## Customizing the server.properties

Look at the `server.properties` file on `minecraft-server`:

```bash
minecraft-server $ sudo docker exec -ti &lt;container name&gt; /bin/bash

root@baea2dfbd18f:/opt/ftb&#35; more server.properties
```

```bash
#Minecraft server properties
&lt; snip! &gt;
server-port=25565
level-type=DEFAULT
enable-rcon=false
level-seed=
force-gamemode=false
server-ip=
max-build-height=256
spawn-npcs=true
white-list=false
spawn-animals=true
hardcore=false
snooper-enabled=true
online-mode=true
resource-pack=
pvp=true     
difficulty=1 
&lt; snip! &gt;
```

Modify the file inside the container:
```bash
root@baea2dfbd18f:/opt/ftb# echo "&lt;full modified content of server.properties&gt;" > server.properties
```

Exit the container and commit the changes to the image:

```bash
root@baea2dfbd18f:/opt/ftb# exit
minecraft-server $ sudo docker commit -m "customized the server.properties" \
                                -a "Jane Doe" \
                                &lt;container id&gt; &lt;docker user&gt;/ftb:v2
```

Verify that the image has been updated:

```bash
minecraft-server $ sudo docker images
```

```bash
REPOSITORY    TAG IMAGE ID  CREATED   VIRTUAL SIZE
&lt;docker user&gt;/ftb v2  9ead6660604b  7 seconds ago 1.013 GB
```

Push the changes to Google Container Registry:

```bash
minecraft-server $ sudo docker tag &lt;docker user&gt;/ftb:v2 gcr.io/&lt;project id&gt;/ftb-v2
minecraft-server $ sudo gcloud docker push gcr.io/&lt;project id&gt;/ftb-v2
```
