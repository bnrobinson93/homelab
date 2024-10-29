# Requirements

1. Install Rancher Desktop
2. Install kubectl (snap or brew)
3. Install k9s (snap)
4. Install helm (snap)

- Also need to run this: `sudo ln -s /snap/k9s/current/bin/k9s /snap/bin/`

# Mealie

Follow the project [here](https://docs.mealie.io/)

# Steps

## Namespaces

1. Create the namespace:

```sh
kubectl apply -f namespace.yaml
```

2. Change context

```sh
kubectl config set-context --current --namespace=mealie
```

## Deployment

3. Create a deployment

```sh
kubectl create deployment --image=nginx mealie -o=yaml --dry-run > deployment.yaml
```

4. Update the deployment.yaml file to change from nginx to mealie, add a namespace, and set the port
5. Apply the yaml

```sh
kubectl apply -f deployment.yaml
```

## Services

6. Forward the port

```sh
kubectl port-forward pods/mealie-6497fc698b-ghlht 9000
```

7. Navigate to [localhost:9000](http://localhost:9000)
8. With services, this can stick around

# Prometheus

```sh
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
# Create monitoring/prometheus-stack
helm install prometheus-stack prometheus-community/kube-prometheus-stack --namespace=monitoring --create-namespace
```

What gets installed:

- prometheus-stack-grafana - UI and metrics visualization
- prometheus-stack-kube-prom-alertmanager - Alerts based on certain metrics (e.g. if memory of pod >800M, email)
- prometheus-stack-kube-prom-operator - Simplifies the deployment of Prometheus & resources for K8s. Allows for monitoring by label
- prometheus-stack-kube-prom-prometheus - Collecting and storing metrics as a time series
- prometheus-stack-kube-state-metrics - Contains state of objects (deployments, pods, etc) and exposes to Prometheus
- prometheus-stack-prometheus-node-exporter - Agent on each worker node to provide node-level metrics
