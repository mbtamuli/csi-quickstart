# `emptyDirClone` CSI plugin

This document introduces the `emptyDirClone` Container Storage Interface (CSI) plugin, aiming to simulate the behavior of a Kubernetes [`emptyDir`][1] for educational purposes.

CSI was developed as a standard for exposing arbitrary block and file storage storage systems to containerized workloads on Container Orchestration Systems (COs) like Kubernetes. With the adoption of the Container Storage Interface, the Kubernetes volume layer becomes truly extensible. The advantage of using CSI is that external storage providers can create and deploy plugins to introduce new storage systems in Kubernetes without having to modify the core Kubernetes code directly.

For more information, you can refer to the resources provided in the [Required Reading](#required) and [Optional Reading](#optional) sections. The CSI specification itself is included in the [Optional Reading](#optional) section. Keep in mind that this guide is designed to be beginner-friendly and serve as a quick start for CSI plugins. The additional tutorials and resources listed in the [Required Reading](#required) section are considered sufficient for understanding the topic.

We will call the plugin `emptyDirClone`.

## Quickstart

See the instructions for [running locally](./docs/running-locally.md).

## Understanding [`emptyDir`][1] Volumes

[`emptyDir`][1] volumes are temporary directories exposed to the pod. These do not persist beyond the lifetime of a pod. This is implemented by Kubernetes itself(as opposed to a separate plugin).

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

## Goals of the `emptyDirClone` CSI Plugin

- Provide a temporary directory
- Provide `ReadWriteOnce` capability. See [Access Modes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes)
- Support ephemeral inline volume requests
- It should have a invocation similar to the following
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
        - mountPath: "/data"                    # Mount path for the volume
          name: my-csi-inline-vol
        command: [ "sleep", "1000000" ]
    volumes:
      - name: my-csi-inline-vol
        csi:
          driver: inline.storage.kubernetes.io  # Name of the CSI driver
          volumeAttributes:
            foo: bar                            # Attributes passed to the driver to configure the volume
  ```

## Implementation steps

- Implement the CSI `Identity` service.
  ```proto
  service Identity {
    rpc GetPluginInfo(GetPluginInfoRequest)
      returns (GetPluginInfoResponse) {}

    rpc GetPluginCapabilities(GetPluginCapabilitiesRequest)
      returns (GetPluginCapabilitiesResponse) {}

    rpc Probe (ProbeRequest)
      returns (ProbeResponse) {}
  }
  ```
- Implement the CSI `Node` service. For ephemeral inline volume request, we only need the `NodePublishVolume`(read _mount volume_), `NodeUnpublishVolume`(read _unmount volume_)
  ```mermaid
  stateDiagram-v2
      direction TB
      [*] --> Published: NodePublishVolume
      Published --> Unpublished: NodeUnpublishVolume
      Unpublished--> [*]
  ```

  ```proto
  service Node {
    rpc NodeStageVolume (NodeStageVolumeRequest)
      returns (NodeStageVolumeResponse) {}

    rpc NodeUnstageVolume (NodeUnstageVolumeRequest)
      returns (NodeUnstageVolumeResponse) {}

    rpc NodePublishVolume (NodePublishVolumeRequest)
      returns (NodePublishVolumeResponse) {}

    rpc NodeUnpublishVolume (NodeUnpublishVolumeRequest)
      returns (NodeUnpublishVolumeResponse) {}

    rpc NodeGetVolumeStats (NodeGetVolumeStatsRequest)
      returns (NodeGetVolumeStatsResponse) {}


    rpc NodeExpandVolume(NodeExpandVolumeRequest)
      returns (NodeExpandVolumeResponse) {}


    rpc NodeGetCapabilities (NodeGetCapabilitiesRequest)
      returns (NodeGetCapabilitiesResponse) {}

    rpc NodeGetInfo (NodeGetInfoRequest)
      returns (NodeGetInfoResponse) {}
  }
  ```

## Reading/References

### Required

- [Volumes](https://kubernetes.io/docs/concepts/storage/volumes/)
- [Configure a Pod to Use a Volume for Storage](https://kubernetes.io/docs/tasks/configure-pod-container/configure-volume-storage/)
- [Volumes](https://kubernetes.io/docs/concepts/storage/ephemeral-volumes/)
- Tutorial blog post - [How to write a Container Storage Interface (CSI) plugin](https://arslan.io/2018/06/21/how-to-write-a-container-storage-interface-csi-plugin/)
- [Kubernetes Container Storage Interface (CSI) Documentation](https://kubernetes-csi.github.io/docs/)
- [Recommended Mechanism for Deploying CSI Drivers on Kubernetes](https://github.com/kubernetes/design-proposals-archive/blob/main/storage/container-storage-interface.md#recommended-mechanism-for-deploying-csi-drivers-on-kubernetes)

### Tools/Examples
- Example implementation - [csi-driver-host-path](https://github.com/kubernetes-csi/csi-driver-host-path)
- Local Testing tool - [`csc`](https://github.com/rexray/gocsi/tree/master/csc)

### Optional

- [CSI Volume Plugins in Kubernetes Design Doc](https://github.com/kubernetes/design-proposals-archive/blob/main/storage/container-storage-interface.md)
- [Container Storage Interface (CSI) for Kubernetes GA](https://kubernetes.io/blog/2019/01/15/container-storage-interface-ga/)
- [Container Storage Interface (CSI) specification](https://github.com/container-storage-interface/spec/blob/v1.9.0/spec.md)
- [Ephemeral Inline CSI volumes KEP](https://github.com/kubernetes/enhancements/blob/ad6021b3d61a49040a3f835e12c8bb5424db2bbb/keps/sig-storage/20190122-csi-inline-volumes.md).
- [enhancement tracking issue #596](https://github.com/kubernetes/enhancements/issues/596).
- Kubernetes code implementing [`emptyDir`][1] - [`pkg/volume/emptydir/empty_dir.go`](https://github.com/kubernetes/kubernetes/blob/master/pkg/volume/emptydir/empty_dir.go)

[1]: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir
