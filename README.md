# CSI Quickstart

The bare minimum goals are complete! Check out the plugin at [emptyDirClone](./emptyDirClone)

## Goals

- [x] Create a quickstart project for the [CSI specification](https://github.com/container-storage-interface/spec/blob/master/spec.md)
    - [x] Easy to read guide summarizing the requirements of a basic CSI plugin simulating [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir)
    - [x] Code implementing the bare minimum plugin described the guide
    - [x] References helpful for learning the supplementary knowledge not directly related to Kubernetes/CSI.
    - [x] Document every decision, resource and changes in the project.
# Stretch Goals
- [ ] Examples describing different scenarios. Different branches/subdirectories for each example, that are as much self-contained as possible.
    - [ ] Volume created external to Kubernetes, manually.
    - [ ] Have a simple API providing a volume that can be invoked by the CSI plugin.
    - [ ] Volume from a public cloud provider.
