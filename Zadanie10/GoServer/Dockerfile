FROM golang:1.22 AS goapp
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ="Europe/Warsaw"
COPY . /GoServer
WORKDIR /GoServer
RUN go mod download
RUN go build .
EXPOSE 22222
CMD go run .
