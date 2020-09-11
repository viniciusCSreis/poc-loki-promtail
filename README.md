# Go Micro-services

Test project with skaffold,kind,chaos mech and nats

# Install skaffold

[install skaffold](https://skaffold.dev/docs/install/)

# Install Kind

[Install Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)

# Run micro services

wait nats cluster to be running. 
(if the nats is not running the pods will restart until the nats is running)

```bash
    ./scripts/run-local.sh
```

# See logs

1. Open grafana `localhost:3000`
1. Login admin;admin
1. Add loki datasource `url=loki:3100`
1. Explore logs using app label