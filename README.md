# CSI Quickstart

The goals of this project are

- [x] Create a quickstart project for the [CSI specification](https://github.com/container-storage-interface/spec/blob/master/spec.md) - [emptyDirClone](./emptyDirClone/)
    - [x] Easy to read guide summarizing the requirements of a basic CSI plugin simulating [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir). - [emptyDirClone Guide](./emptyDirClone/)
    - [x] Code implementing the bare minimum plugin described the guide. [emptyDirClone](./emptyDirClone/)
    - [x] References helpful for learning the supplementary knowledge not directly related to Kubernetes/CSI. - [References](./emptyDirClone/README.md#readingreferences)
    - [x] Document every decision, resource and changes in the project. - [Decisions](./emptyDirClone/docs/decisions.md)

### Stretch Goals
- [ ] Examples describing different scenarios. Different branches/subdirectories for each example, that are as much self-contained as possible.
    - [ ] Volume created external to Kubernetes, manually.
    - [ ] Have a simple API providing a volume that can be invoked by the CSI plugin.
    - [ ] Volume from a public cloud provider.
