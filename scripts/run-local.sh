#!/bin/bash

CURRENT_CONTEXT=$(kubectl config current-context)

if [ "$CURRENT_CONTEXT" != 'kind-kind' ]; then
  echo "CURRENT_CONTEXT: $CURRENT_CONTEXT != kind-kind"
  exit 1
fi

kubectl get clusterrolebinding default-admin 2> /dev/null || kubectl create clusterrolebinding default-admin --clusterrole=cluster-admin --serviceaccount=default:default

skaffold run --port-forward=true --tail
skaffold delete > /dev/null &