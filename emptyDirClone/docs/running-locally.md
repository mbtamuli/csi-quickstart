# Running Locally

## Pre-Requisites

- [Docker](https://docs.docker.com/engine/install/) | For macOS or Windows install [Docker Desktop](https://docs.docker.com/desktop/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)

## Quickstart

Run the following to have the plugin deployed into a local cluster

```sh
make cluster-up
make deploy
```

## Developing Locally

Helpful [`Make`][1] targets are added to allow doing common operational tasks. Run `make help` to see all options.

### Setup cluster

1. Once the pre-requisites are met, you can create the local Kubernetes cluster using
    ```sh
    make cluster-up

    # Or manually
    kind create cluster
    ```

1. Using any of these methods, once the cluster is created, verify the cluster is running using
    ```sh
    kubectl get nodes
    ```

1. You should see an output similar to
    ```
    NAME                 STATUS   ROLES           AGE   VERSION
    kind-control-plane   Ready    control-plane   64s   v1.27.3
    ```

### Tear down cluster

Once the pre-requisites are met, you can create the local Kubernetes cluster using
  ```sh
  make cluster-down

  # Or manually
  kind delete cluster
  ```

### Running tests

Run the tests using
```sh
go test -v ./tests/e2e
```

### Build and Deploy

1. To build the image
    ```sh
    # with the defaults - `ghcr.io/mbtamuli/csi-quickstart/emptydirclone:latest`
    make docker-build
    ```
    _**Note:** You can change the image name or tag as follows, which also works with `docker-push`,`kind-load` and `deploy` targets_
    ```sh
    # setting the image name - `myimage:latest`
    make docker-build IMG=myimage

    # setting the image name and tag - `myimage:1.0.0`
    make docker-build IMG=myimage TAG=1.0.0
    ```

1. To push the image
    ```sh
    make docker-push
    ```

1. To load the image locally into the cluster created in [Setup Cluster](#setup-cluster)
    ```sh
    make kind-load
    ```

1. To deploy the manifests
    ```
    make deploy
    ```

[1]: https://www.gnu.org/software/make/
