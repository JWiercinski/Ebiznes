FROM node:18.16.0 AS reactapp
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Warsaw"
COPY . /reactclient
WORKDIR /reactclient
RUN npm ci
EXPOSE 3000
