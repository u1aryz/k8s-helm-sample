# k8s-helm-sample
Simple app running on k8s.

## Requirements
* Docker
* kubectl
* helm

## Setup
- Start your cluster

    If you use a minikube, run it.
    ```
    minikube -p helm-sample start
    ```

- Create secret

    **NOTE**
    Please replace the following `<YOUR_REDIS_PASSWORD>`.
    ```
    kubectl create secret generic app-secret \
        --from-literal=redis-password=<YOUR_REDIS_PASSWORD>
    ```

    ex.
    ```
    kubectl create secret generic app-secret \
        --from-literal=redis-password=hogepiyo
    ```

## Usage on minikube
- Build
    ```
    minikube -p helm-sample docker-env | source
    docker build -t u1aryz/api:v1.0 api
    docker build -t u1aryz/frontend:v1.0 frontend
    ```

- Start
    ```
    helm dependency update k8s
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

## How to access to redis
- Get your redis password

    POSIX
    ```
    export REDIS_PASSWORD=$(kubectl get secret app-secret -o jsonpath="{.data.redis-password}" | base64 --decode)
    ```
    Fish shell
    ```
    set -x REDIS_PASSWORD (kubectl get secret app-secret -o jsonpath="{.data.redis-password}" | base64 --decode)
    ```

- Create pod that you can use as a client
    ```
    kubectl run redis-client --restart='Never' --env REDISCLI_AUTH=$REDIS_PASSWORD \
        --image docker.io/bitnami/redis:6.2.6-debian-10-r21 --command -- sleep infinity
    ```

- Attach to redis client
    ```
    kubectl exec --tty -i redis-client -- bash
    ```

- Access to redis on redis-client pod
    ```
    redis-cli -h helm-sample-redis-master
    ```
