#FROM ubuntu
FROM ubuntu:22.04
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Warsaw"
RUN apt-get update
RUN apt-get install -y tzdata
#TASK1
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:deadsnakes/ppa -y
RUN apt-get update
#TASK2
RUN apt-get install -y python3.8 
RUN apt-get install -y openjdk-8-jdk 
#snapd
#RUN snap install --classic kotlin
RUN apt update
RUN apt install -y kotlin
#RUN curl -s "https://get.sdkman.io" | bash
#RUN sdk install gradle
RUN useradd -ms /bin/bash jackson
WORKDIR /home/jackson
USER jackson