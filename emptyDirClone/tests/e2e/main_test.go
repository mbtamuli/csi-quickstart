package e2e

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/e2e-framework/klient/decoder"
	"sigs.k8s.io/e2e-framework/klient/k8s"
	"sigs.k8s.io/e2e-framework/klient/k8s/resources"
	"sigs.k8s.io/e2e-framework/klient/wait"
	"sigs.k8s.io/e2e-framework/klient/wait/conditions"
	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/envfuncs"
	"sigs.k8s.io/e2e-framework/support/kind"
)

var (
	testEnv         env.Environment
	kindClusterName string
	namespace       string
)

func TestMain(m *testing.M) {
	const (
		csiNamespace string = "emptydirclone"
	)
	csiPluginPodLabels := map[string]string{"app": "emptydirclone-plugin"}
	cfg, _ := envconf.NewFromFlags()
	testEnv = env.NewWithConfig(cfg)
	kindClusterName = envconf.RandomName("csi", 10)
	namespace = envconf.RandomName("emptydirclone", 20)
	testEnv.Setup(
		envfuncs.CreateCluster(kind.NewProvider(), kindClusterName),
		envfuncs.CreateNamespace(csiNamespace),
		deployEmptyDirClone(csiNamespace, csiPluginPodLabels),
		envfuncs.CreateNamespace(namespace),
	)

	testEnv.Finish(
		envfuncs.ExportClusterLogs(kindClusterName, fmt.Sprintf("logs/cluster-logs-%s", kindClusterName)),
		envfuncs.DestroyCluster(kindClusterName),
	)

	os.Exit(testEnv.Run(m))
}

func deployEmptyDirClone(namespace string, labels map[string]string) env.Func {
	return func(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
		decoder.ApplyWithManifestDir(
			ctx,
			cfg.Client().Resources(),
			"../../deploy",
			"*",
			[]resources.CreateOption{},
		)

		pods := &v1.PodList{
			Items: []v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: namespace,
						Labels:    labels,
					},
				},
			},
		}
		// wait for the pods to be Ready
		if err := wait.For(
			conditions.New(cfg.Client().Resources()).
				ResourcesMatch(pods,
					func(object k8s.Object) bool {
						// phase := object.(*v1.Pod).Status.Phase
						// fmt.Println("Current Phase of the pod resource", "phase", phase)
						// return phase == v1.PodRunning

						status := object.(*v1.Pod).Status
						fmt.Println("Current Status of the pod resource", "status", status)
						for _, cond := range status.Conditions {
							if cond.Type == v1.PodReady && cond.Status == v1.ConditionTrue {
								return true
							}
						}
						return false
					},
				),
			wait.WithInterval(time.Second*30),
			wait.WithTimeout(time.Minute*5)); err != nil {
			return ctx, err
		}

		return ctx, nil
	}
}
