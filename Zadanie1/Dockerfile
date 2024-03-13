#FROM ubuntu
FROM ubuntu:22.04
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Warsaw"
#VOLUME /home/jackson
#RUN useradd -ms /bin/bash jackson
#RUN adduser jackson
RUN mkdir -p /data/app
WORKDIR /data/app
RUN apt-get update
RUN apt-get install -y tzdata
#TASK1
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:deadsnakes/ppa -y
RUN apt-get update
RUN apt-get install -y python3.8 
#TASK2
RUN apt-get install -y openjdk-8-jdk  openjdk-8-jre
RUN apt update
RUN apt install -y kotlin
#TASK3
#RUN add-apt-repository ppa:cwchien/gradle
#RUN apt-get update
#RUN apt-get install -y gradle
RUN apt-get update
RUN apt-get -y install curl unzip
RUN \
    cd /usr/local && \
    curl -L https://services.gradle.org/distributions/gradle-2.5-bin.zip -o gradle-2.5-bin.zip && \
    unzip gradle-2.5-bin.zip && \
    rm gradle-2.5-bin.zip
ENV GRADLE_HOME=/usr/local/gradle-2.5
ENV PATH=$PATH:$GRADLE_HOME/bin JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
#USER jackson
WORKDIR /home/jackson
#RUN mkdir jdbcsqliteapp
COPY jdbcsqliteapp /home/jackson/jdbcsqliteapp
#COPY jdbcsqliteapp/* /home/jackson/jdbcsqliteapp/
WORKDIR /home/jackson/jdbcsqliteapp
#ENV GRADLE_USER_HOME=~/.gradle
#RUN gradle --stop
#RUN gradle init
#RUN ./gradlew --version
#RUN chmod +rx /home/jackson/.gradle/wrapper/dists/gradle-4.4.1-bin/46gopw3g8i1v3zqqx4q949t2x/gradle-4.4.1/bin/gradle
#RUN ./gradlew wrapper --gradle-version latest
RUN gradle build
WORKDIR /home/jackson