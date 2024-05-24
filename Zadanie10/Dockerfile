FROM golang:1.22 AS goapp
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Warsaw"
COPY GoServer /GoServer
WORKDIR /GoServer
RUN go mod download
RUN go build .
EXPOSE 22222
CMD go run .
FROM node:18.16.0 AS reactapp
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Warsaw"
COPY reactclient /reactclient
WORKDIR /reactclient
RUN npm ci
EXPOSE 3000
CMD npm start
