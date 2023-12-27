# Running Locally

## Pre-Requisites

- [Docker](https://docs.docker.com/engine/install/) | For macOS or Windows install [Docker Desktop](https://docs.docker.com/desktop/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)

- \(Optional\) [stern](https://github.com/stern/stern#installation)

## Quickstart

Run the following to have the plugin deployed into a local cluster

```sh
make cluster-up
make deploy
```

## Developing Locally

[`Make`][1] targets are added to allow doing common operational tasks. Run `make help` to see all options.

### Setup cluster

1. Once the pre-requisites are met, you can create the local Kubernetes cluster using
    ```sh
    make cluster-up
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

To delete the local cluster, run
  ```sh
  make cluster-down
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
    ```sh
    make deploy
    ```

### Running tests and checking logs

1. Check plugin logs. Run the following in a separate terminal as it will "follow" the logs
    ```sh
    make logs
    ```

1. Run [`csc`](https://github.com/rexray/gocsi/tree/master/csc) tests
    ```sh
    make csc-tests
    ```


1. Run the e2e tests using. See [`e2e`](../tests/e2e/)
    ```sh
    make e2e
    ```

[1]: https://www.gnu.org/software/make/
