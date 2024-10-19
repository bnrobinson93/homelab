# Requirements

1. Install Rancher Desktop
2. Install kubectl (snap or brew)
3. Install k9s (snap)

# Mealie

Follow the project [here](https://docs.mealie.io/)

# Steps
1. Create the namespace:
  ```sh
  k apply -f namespace.yaml
  ```
2. Change context

```sh
kubectl config set-context --current --namespace=mealie
```
3. Create a deployment
  ```sh
  k create deployment --image=nginx mealie -o=yaml --dry-run > deployment.yaml
  ```
4. Update the deployment.yaml file to change from nginx to mealie, add a namespace, and set the port
5. Apply the yaml
```sh
k apply -f deployment.yaml
```
