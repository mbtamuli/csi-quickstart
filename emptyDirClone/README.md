# CSI Plugin simulating Kubernetes [`emptyDir`][1]

The purpose of this plugin is to simulate the functioning of an [`emptyDir`][1] via a CSI plugin. Using CSI, third-party storage providers can write and deploy plugins exposing new storage systems in Kubernetes without ever having to touch the core Kubernetes code. There are reading resources mentioned in [Required Reading](#required-reading) and [Optional Reading](#optional-reading).

We will call the plugin `emptyDirClone`.

## Introduction to [`emptyDir`][1]

[`emptyDir`][1] volumes are temporary directories exposed to the pod. These do not persist beyond the lifetime of a pod.

    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
    name: test-pd
    spec:
    containers:
    - image: registry.k8s.io/test-webserver
        name: test-container
        volumeMounts:
        - mountPath: /cache
        name: cache-volume
    volumes:
    - name: cache-volume
        emptyDir:
        sizeLimit: 500Mi
    ```


[1]: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir

## Requirements of `emptyDirClone` CSI Plugin

- [ ] Provide a temporary directory
- [ ] Provide `ReadWriteOnce` capability. See [Access Modes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes)
- [ ] Support ephemeral inline volume requests

Have a invocation similar to the following

    ```yaml
    kind: Pod
    apiVersion: v1
    metadata:
    name: my-csi-app
    spec:
    containers:
        - name: my-frontend
        image: busybox:1.28
        volumeMounts:
        - mountPath: "/data"
            name: my-csi-inline-vol
        command: [ "sleep", "1000000" ]
    volumes:
        - name: my-csi-inline-vol
        csi:
            driver: inline.storage.kubernetes.io
            volumeAttributes:
            foo: bar
    ```

## Required Reading

- [Volumes]https://kubernetes.io/docs/concepts/storage/volumes/
- [Configure a Pod to Use a Volume for Storage](https://kubernetes.io/docs/tasks/configure-pod-container/configure-volume-storage/)
- [Volumes](https://kubernetes.io/docs/concepts/storage/ephemeral-volumes/)
- Tutorial blog post - [How to write a Container Storage Interface (CSI) plugin](https://arslan.io/2018/06/21/how-to-write-a-container-storage-interface-csi-plugin/)
- [Kubernetes Container Storage Interface (CSI) Documentation](https://kubernetes-csi.github.io/docs/)
- [Container Storage Interface (CSI) specification](https://github.com/container-storage-interface/spec/blob/v1.9.0/spec.md)
- Local Testing tool - [`csc`](https://github.com/rexray/gocsi/tree/master/csc)

## Optional Reading

- [Container Storage Interface (CSI) for Kubernetes GA](https://kubernetes.io/blog/2019/01/15/container-storage-interface-ga/)
- Example implementation - [csi-driver-host-path](https://github.com/kubernetes-csi/csi-driver-host-path)
- [Ephemeral Inline CSI volumes KEP](https://github.com/kubernetes/enhancements/blob/ad6021b3d61a49040a3f835e12c8bb5424db2bbb/keps/sig-storage/20190122-csi-inline-volumes.md).
- [enhancement tracking issue #596](https://github.com/kubernetes/enhancements/issues/596).
