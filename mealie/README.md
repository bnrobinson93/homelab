# Requirements

1. Install Rancher Desktop
2. Install kubectl (snap or brew)
3. Install k9s (snap)

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
