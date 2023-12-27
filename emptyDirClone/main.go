package main

import (
	"flag"
	"os"

	"github.com/go-logr/logr"
	"github.com/mbtamuli/emptyDirClone/internal/emptydirclone"
	"github.com/mbtamuli/emptyDirClone/internal/log"
)

var version string

func main() {
	const (
		pluginname = "emptydirclone.mriyam.dev"
	)

	var (
		logger      logr.Logger
		level       int
		environment string
		endpoint    string
		nodeid      string
	)

	flag.IntVar(&level, "verbosity", 0, "Log level verbosity. Lower is verbose and less imporant. Higher is quiet and more important. Default: 0; Valid options: -128 to 128")
	flag.StringVar(&environment, "environment", "production", "Environment the application is run in. Default: \"production\"; Valid options: \"production\", \"development\"")
	flag.StringVar(&endpoint, "endpoint", "unix:///csi/csi.sock", "Endpoint for the gRPC server to listen on. Default: \"unix:///csi/csi.sock\"")
	flag.StringVar(&nodeid, "nodeid", "", "Kubernetes Node ID")
	flag.Parse()

	logger = log.SetupLogger(level, environment)

	if version == "" {
		version = "0.0.0"
	}

	cfg := emptydirclone.Config{
		Endpoint:      endpoint,
		Name:          pluginname,
		NodeID:        nodeid,
		VendorVersion: version,
	}

	emptydirclone, err := emptydirclone.New(cfg, logger)
	if err != nil {
		logger.Error(err, "unable to set up the plugin")
		os.Exit(1)
	}

	if err := emptydirclone.Serve(); err != nil {
		logger.Error(err, "unable to start the plugin")
		os.Exit(1)
	}
}
