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
3. Run HTTP
  ```sh
  k run brad-mealie --image=nginx -n mealie
  ```
