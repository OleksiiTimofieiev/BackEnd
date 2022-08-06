FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN clone && build

EXPOSE 8080
ENTRYPOINT <binaryPath>