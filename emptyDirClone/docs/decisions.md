# Decisions

## Simulate [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir) or [`hostPath`](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath)

**Verdict:** I've decided to go with [`emptyDir`](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir).

**Reason:** After skimming the documentation for both, I would like to take what seems like the simplest option. [`hostPath`](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) requires a [privileged Container](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/) and would, most probably, require the plugin to be aware of the filesystem on the node and behave differently increasing complexity.

## Logging library

**Options**
- [sirupsen/logrus](https://pkg.go.dev/github.com/sirupsen/logrus)
- [go-kit/log](https://pkg.go.dev/github.com/go-kit/log)
- [go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap)
- [log/slog](https://pkg.go.dev/log/slog)
- [golang/glog](https://pkg.go.dev/github.com/golang/glog)
- [k8s.io/klog](https://pkg.go.dev/k8s.io/klog/v2)
- [go-logr/logr](https://pkg.go.dev/github.com/go-logr/logr)

**Verdict:** I've decided to go with [go-logr/logr](https://pkg.go.dev/github.com/go-logr/logr).

**Reason:** Comfortability. I'm already comfortable with [go-logr/logr](https://pkg.go.dev/github.com/go-logr/logr). It is a logging API and allows using various other "backends"(See [Implementations](https://github.com/go-logr/logr#implementations-non-exhaustive)) - including but not limited to the other libraries mentioned in the list.
