package e2e

import (
	"context"
	"testing"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/e2e-framework/klient/k8s"
	"sigs.k8s.io/e2e-framework/klient/wait"
	"sigs.k8s.io/e2e-framework/klient/wait/conditions"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

func TestEmptyDirClone(t *testing.T) {
	emptyDirCloneFeature := features.New("emptyDirClone").
		Setup(func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			client, err := cfg.NewClient()
			if err != nil {
				t.Fatal(err)
			}

			// create a pod
			pod := newPod(cfg.Namespace(), "test-pod")
			if err != nil {
				t.Fatal(err)
			}
			if err := client.Resources().Create(ctx, pod); err != nil {
				t.Fatal(err)
			}

			// create a job with one container
			writeJob := newJobWithReadWriteCmd(newPod(cfg.Namespace(), "test-job-write"), 1)
			if err != nil {
				t.Fatal(err)
			}
			if err := client.Resources().Create(ctx, writeJob); err != nil {
				t.Fatal(err)
			}

			// create a job with two containers
			readJob := newJobWithReadWriteCmd(newPod(cfg.Namespace(), "test-job-read"), 2)
			if err := client.Resources().Create(ctx, readJob); err != nil {
				t.Fatal(err)
			}
			return ctx
		}).
		Assess("pod is running", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			client, err := cfg.NewClient()
			if err != nil {
				t.Fatal(err)
			}

			// check for the pod
			found := v1.Pod{
				ObjectMeta: metav1.ObjectMeta{Name: "test-pod", Namespace: cfg.Namespace()},
			}

			// wait for the pod to be in 'Running' phase
			err = wait.For(conditions.New(client.Resources()).ResourceMatch(&found, func(object k8s.Object) bool {
				p := object.(*v1.Pod)
				t.Log(p.Status.Phase)
				return p.Status.Phase == v1.PodRunning
			}), wait.WithTimeout(time.Second*30))
			if err != nil {
				t.Fatal(err)
			}

			return ctx
		}).
		Assess("single container is able to read/write", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			client, err := cfg.NewClient()
			if err != nil {
				t.Fatal(err)
			}

			// check for the job
			found := batchv1.Job{
				ObjectMeta: metav1.ObjectMeta{Name: "test-job-write", Namespace: cfg.Namespace()},
			}
			// wait for the job to succeed
			err = wait.For(conditions.New(client.Resources()).ResourceMatch(&found, func(object k8s.Object) bool {
				p := object.(*batchv1.Job)
				return p.Status.Succeeded == 1
			}), wait.WithTimeout(time.Second*30))
			if err != nil {
				t.Fatal(err)
			}

			return ctx
		}).
		Assess("two containers are able to read/write", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			client, err := cfg.NewClient()
			if err != nil {
				t.Fatal(err)
			}

			// check for the job
			found := batchv1.Job{
				ObjectMeta: metav1.ObjectMeta{Name: "test-job-read", Namespace: cfg.Namespace()},
			}
			// wait for the job to succeed
			err = wait.For(conditions.New(client.Resources()).ResourceMatch(&found, func(object k8s.Object) bool {
				p := object.(*batchv1.Job)
				return p.Status.Succeeded == 1
			}), wait.WithTimeout(time.Second*30))
			if err != nil {
				t.Fatal(err)
			}

			return ctx
		}).Feature()

	testEnv.Test(t, emptyDirCloneFeature)
}

func newPod(namespace string, name string) *v1.Pod {
	gracePeriod := int64(0)
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				"app": "emptyDirClone-csi-volume-test",
			},
		},
		Spec: v1.PodSpec{
			TerminationGracePeriodSeconds: &gracePeriod,
			Containers: []v1.Container{
				{
					Name:  "csi-volume-test",
					Image: "busybox:1.28",
					Command: []string{
						"sleep",
						"1000000",
					},
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      "csi-inline-vol",
							MountPath: "/data",
						},
					},
				},
			},
			Volumes: []v1.Volume{
				{
					Name: "csi-inline-vol",
					VolumeSource: v1.VolumeSource{
						CSI: &v1.CSIVolumeSource{
							Driver: "emptydirclone.mriyam.dev",
						},
					},
				},
			},
		},
	}
}

func newJobWithReadWriteCmd(pod *v1.Pod, count int) *batchv1.Job {
	job := &batchv1.Job{
		ObjectMeta: pod.ObjectMeta,
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: *pod.Spec.DeepCopy(),
			},
		},
	}
	job.Spec.Template.Spec.RestartPolicy = v1.RestartPolicyNever

	job.Spec.Template.Spec.Containers[0] = v1.Container{
		Name:    "csi-volume-test-write",
		Image:   "busybox:1.28",
		Command: []string{"/bin/sh"},
		Args: []string{
			"-c",
			"touch /data/hello.txt && touch /data/world.txt && stat /data/hello.txt",
		},
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "csi-inline-vol",
				MountPath: "/data",
			},
		},
	}

	if count == 1 {
		return job
	}

	job.Spec.Template.Spec.Containers = append(job.Spec.Template.Spec.Containers, v1.Container{
		Name:    "csi-volume-test-read",
		Image:   "busybox:1.28",
		Command: []string{"/bin/sh"},
		Args: []string{
			"-c",
			"stat /data/hello.txt && stat /data/world.txt",
		},
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "csi-inline-vol",
				MountPath: "/data",
			},
		},
	})

	if count == 2 {
		return job
	}

	return nil
}
