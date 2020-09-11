#!/bin/bash

CURRENT_CONTEXT=$(kubectl config current-context)

if [ "$CURRENT_CONTEXT" != 'kind-kind' ]; then
  echo "CURRENT_CONTEXT: $CURRENT_CONTEXT != kind-kind"
  exit 1
fi

kubectl get clusterrolebinding default-admin 2> /dev/null || kubectl create clusterrolebinding default-admin --clusterrole=cluster-admin --serviceaccount=default:default

mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o ./bin/my-app-1 -v cmd/my-app-1/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o ./bin/my-app-2 -v cmd/my-app-2/main.go

skaffold run --port-forward=true --tail
skaffold delete > /dev/null &