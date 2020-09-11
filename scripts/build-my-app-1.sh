#!/bin/bash

mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o ./bin/my-app-1 -v cmd/my-app-1/main.go

docker build -t my-app-1 -f my-app-1.Dockerfile .