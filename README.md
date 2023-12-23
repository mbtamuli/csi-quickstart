# CSI Quickstart

The goals of this project are

- [ ] Create a quickstart project for the [CSI specification](https://github.com/container-storage-interface/spec/blob/master/spec.md)
    - [ ] Easy to read guide summarizing the requirements of a basic CSI plugin. Simulating something like [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir)/[`hostPath`](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) - to be decided.
    - [ ] Code implementing the bare minimum plugin described the guide.
    - [ ] References helpful for learning the supplementary knowledge not directly related to Kubernetes/CSI.

## Stretch Goals
- [ ] Examples describing different scenarios
    - Different branches/subdirectories for each example, that are as much self-contained as possible
    - [ ] Volume created external to Kubernetes, manually.
    - [ ] Have a simple API providing a volume that can be invoked by the CSI plugin.
    - [ ] Volume from a public cloud provider.
