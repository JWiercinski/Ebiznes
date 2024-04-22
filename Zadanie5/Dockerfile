FROM golang:1.22 AS goapp
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Warsaw"
COPY GoServer /GoServer
WORKDIR /GoServer
RUN go mod download
EXPOSE 22222
FROM node:18.16.0 AS reactapp
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Warsaw"
COPY reactclient /reactclient
WORKDIR /reactclient
RUN npm ci
EXPOSE 3000
