# Decisions

## Simulate [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) or [`hostPath`](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath)

**Verdict:** I've decided to go with [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir).

**Reason:** After skimming the documentation for both, I would like to take what seems like the simplest option. [`hostPath`](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) requires a [privileged Container](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/) and would, most probably, require the plugin to be aware of the filesystem on the node and behave differently increasing complexity.
