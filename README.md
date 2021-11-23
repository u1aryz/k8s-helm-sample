# k8s-helm-sample
Simple app running on k8s.

## Requirements
* Docker
* kubectl
* helm

## Usage on minikube
- Build
    ```
    minikube -p helm-sample start
    minikube -p helm-sample docker-env | source
    docker build -t u1aryz/api:v1.0 api
    docker build -t u1aryz/frontend:v1.0 frontend
    ```

- Start
    ```
    helm install helm-sample k8s -f k8s/values-minikube.yaml
    minikube -p helm-sample addons enable ingress
    minikube -p helm-sample tunnel
    ```
    You can visit [page](http://localhost) from your browser.

- Clean up
    ```
    helm delete helm-sample
    minikube -p helm-sample docker-env -u | source
    minikube -p helm-sample stop
    minikube -p helm-sample delete
    ```
