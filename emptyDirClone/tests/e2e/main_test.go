package e2e

import (
	"os"
	"testing"

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
	cfg, _ := envconf.NewFromFlags()
	testEnv = env.NewWithConfig(cfg)
	kindClusterName = envconf.RandomName("csi", 10)
	namespace = envconf.RandomName("emptydirclone", 20)

	testEnv.Setup(
		envfuncs.CreateCluster(kind.NewProvider(), kindClusterName),
		envfuncs.CreateNamespace(namespace),
	)

	testEnv.Finish(
		envfuncs.DeleteNamespace(namespace),
		envfuncs.DestroyCluster(kindClusterName),
	)

	os.Exit(testEnv.Run(m))
}
