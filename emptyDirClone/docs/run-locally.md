# Running Locally

## Pre-Requisites

- [Docker](https://docs.docker.com/engine/install/) | For macOS or Windows install [Docker Desktop](https://docs.docker.com/desktop/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)

## Developing Locally

### Setup cluster

1. Once the pre-requisites are met, you can create the local Kubernetes cluster using
  ```sh
  kind create cluster
  ```
2. Verify the cluster is running using
  ```sh
  kubectl get nodes
  ```
  You should see an output similar to
  ```
  NAME                 STATUS     ROLES           AGE   VERSION
  kind-control-plane   NotReady   control-plane   9s    v1.27.3
  ```
3. If you want to easily setup and tear down the cluster from scratch, you can also use the following commands to do so
  ```sh
  make cluster-up
  make cluster-down
  ```
