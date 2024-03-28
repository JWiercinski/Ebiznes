FROM ubuntu:22.04 AS mainbuild
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Warsaw"
RUN apt-get update
RUN apt-get install -y tzdata
RUN apt-get update
RUN apt-get install -y software-properties-common
RUN apt-get update
RUN apt-get install -y openjdk-8-jdk  openjdk-8-jre
RUN echo "deb https://repo.scala-sbt.org/scalasbt/debian all main" | tee /etc/apt/sources.list.d/sbt.list
RUN echo "deb https://repo.scala-sbt.org/scalasbt/debian /" | tee /etc/apt/sources.list.d/sbt_old.list
RUN apt-get install -y curl
RUN curl -sL "https://keyserver.ubuntu.com/pks/lookup?op=get&search=0x2EE0EA64E40A89B84B2DF73499E82A75642AC823" | apt-key add
RUN apt-get update
RUN apt-get install -y sbt=1.9.2
RUN apt-get install -y wget
EXPOSE 9000
FROM mainbuild AS scalaapp
RUN wget www.scala-lang.org/files/archive/scala-2.12.12.deb
RUN dpkg -i scala-2.12.12.deb
ENV SCALA_HOME /usr/share/scala
WORKDIR /scalaproject
COPY scalapplay2 /scalapplay2
WORKDIR /scalapplay2
FROM mainbuild AS ngrokapp
RUN wget -q -c -nc https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip
RUN apt-get install unzip
RUN unzip -qq -n ngrok-stable-linux-amd64.zip
