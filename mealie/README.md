# Steps
1. Create the namespace:
  ```sh
  k apply -f namespace.yaml
  ```
2. Run HTTP
  ```sh
  k run brad-mealie --image=nginx -n mealie
  ```
