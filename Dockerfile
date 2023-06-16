# This dockerfile uses extends image https://hub.docker.com/bridgewwater/template-golang-lib
# VERSION 1
# Author: sinlov
# dockerfile offical document https://docs.docker.com/engine/reference/builder/
# https://hub.docker.com/_/golang
FROM golang:1.17.13-buster as builder

ARG GO_PATH_SOURCE_DIR=/go/src/
WORKDIR ${GO_PATH_SOURCE_DIR}

RUN mkdir -p ${GO_PATH_SOURCE_DIR}/github.com/bridgewwater/template-golang-lib
COPY $PWD ${GO_PATH_SOURCE_DIR}/github.com/bridgewwater/template-golang-lib

RUN cd ${GO_PATH_SOURCE_DIR}/github.com/bridgewwater/template-golang-lib && \
    go mod download -x

RUN  cd ${GO_PATH_SOURCE_DIR}/github.com/bridgewwater/template-golang-lib && \
  CGO_ENABLED=0 \
  go build \
  -a \
  -installsuffix cgo \
  -ldflags '-w -s --extldflags "-static -fpic"' \
  -tags netgo \
  -o template-golang-lib \
  main.go

# https://hub.docker.com/_/alpine
FROM alpine:3.17

ARG DOCKER_CLI_VERSION=${DOCKER_CLI_VERSION}

#RUN apk --no-cache add \
#  ca-certificates mailcap curl \
#  && rm -rf /var/cache/apk/* /tmp/*

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/bridgewwater/template-golang-lib/template-golang-lib .
ENTRYPOINT ["/app/template-golang-lib"]
# CMD ["/app/template-golang-lib", "--help"]