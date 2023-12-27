package main

import (
	"flag"

	"github.com/mbtamuli/emptyDirClone/internal/emptydirclone"
)

var version string

func main() {
	const (
		pluginname = "emptydirclone.mriyam.dev"
	)
	var endpoint string
	flag.StringVar(&endpoint, "endpoint", "unix:/csi/csi.sock", "Endpoint for the gRPC server to listen on. Default: \"unix:/csi/csi.sock\"")
	flag.Parse()

	if version == "" {
		version = "0.0.0"
	}

	cfg := emptydirclone.Config{
		Name:          pluginname,
		Endpoint:      endpoint,
		VendorVersion: version,
	}

	emptydirclone := emptydirclone.New(cfg)
	if err := emptydirclone.Serve(); err != nil {
		panic(err)
	}
}
