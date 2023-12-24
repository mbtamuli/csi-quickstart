# Running Locally

## Pre-Requisites

- [Docker](https://docs.docker.com/engine/install/) | For macOS or Windows install [Docker Desktop](https://docs.docker.com/desktop/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)

## Developing Locally

### Setup cluster

#### [`Make`][1]

1. Helpful [`Make`][1] targets are added to allow doing common operational tasks
    ```sh
    make cluster-up
    ```

#### Manually
1. Once the pre-requisites are met, you can create the local Kubernetes cluster using
    ```sh
    kind create cluster
    ```

Using any of these methods, once the cluster is created, verify the cluster is running using
```sh
kubectl get nodes
```
You should see an output similar to
```
NAME                 STATUS   ROLES           AGE   VERSION
kind-control-plane   Ready    control-plane   64s   v1.27.3
```

### Tear down cluster

#### `Make`

1. Helpful [`Make`] targets are added to allow doing common operational tasks
    ```sh
    make cluster-down
    ```

#### Manually
1. Once the pre-requisites are met, you can create the local Kubernetes cluster using
    ```sh
    kind delete cluster
    ```

[1]: https://www.gnu.org/software/make/
