#!/bin/bash
docker run --rm -v "$PWD":/go/src/app -w /go/src/app \
    amazonlinux:2 \
    /bin/bash -c "yum install -y gcc go zip && \
                  GOOS=linux GOARCH=amd64 go build -o bin/bootstrap handler.go && \
                  cd bin && zip bootstrap.zip bootstrap"
