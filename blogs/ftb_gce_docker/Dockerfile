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

FROM debian:7.8
MAINTAINER Jane Doe <jane@test.com>
RUN apt-get update && apt-get install -y \
	curl \
	openjdk-7-jre \
	unzip

ENV FTB_VERSION 1_6_2
ENV PORT 25565
ENV BASE_URL http://ftb.cursecdn.com/FTB2/modpacks/FTBInfinity
ENV SERVER_FILE FTBInfinityServer.zip
 
RUN mkdir /opt/ftb
WORKDIR /opt/ftb

COPY CheckEula.sh /opt/ftb/

RUN curl $BASE_URL/$FTB_VERSION/$SERVER_FILE -o ftb.zip && \
	unzip ftb.zip && \
	chmod u+x FTBInstall.sh ServerStart.sh CheckEula.sh && \
	rm eula.txt

RUN sed -i '2i /bin/bash /opt/ftb/CheckEula.sh' /opt/ftb/ServerStart.sh

RUN /opt/ftb/FTBInstall.sh

EXPOSE ${PORT}

CMD /opt/ftb/ServerStart.sh
