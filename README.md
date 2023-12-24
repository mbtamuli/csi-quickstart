# CSI Quickstart

The goals of this project are

- [ ] Create a quickstart project for the [CSI specification](https://github.com/container-storage-interface/spec/blob/master/spec.md)
    - [x] Easy to read guide summarizing the requirements of a basic CSI plugin simulating [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir).
    - [ ] Code implementing the bare minimum plugin described the guide.
    - [x] References helpful for learning the supplementary knowledge not directly related to Kubernetes/CSI.
    - [ ] Document every decision, resource and changes in the project.

### Stretch Goals
- [ ] Examples describing different scenarios. Different branches/subdirectories for each example, that are as much self-contained as possible.
    - [ ] Volume created external to Kubernetes, manually.
    - [ ] Have a simple API providing a volume that can be invoked by the CSI plugin.
    - [ ] Volume from a public cloud provider.

## Decisions

### Simulate [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) or [`hostPath`](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath)

**Verdict:** I've decided to go with [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir).

**Reason:** After skimming the documentation for both, I would like to take what seems like the simplest option. [`hostPath`](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) requires a [privileged Container](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/) and would, most probably, require the plugin to be aware of the filesystem on the node and behave differently increasing complexity.
