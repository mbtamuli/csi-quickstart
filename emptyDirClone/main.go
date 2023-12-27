package main

import (
	"flag"
	"fmt"

	"github.com/mbtamuli/emptyDirClone/internal/emptydirclone"
)

var version string

func main() {
	var endpoint string

	flag.StringVar(&endpoint, "endpoint", "unix:/csi/csi.sock", "Endpoint for the gRPC server to listen on. Default: \"unix:/csi/csi.sock\"")
	flag.Parse()

	if version == "" {
		version = "0.0.0"
	}

	fmt.Printf("Version: %s\n", version)

	cfg := emptydirclone.Config{
		Endpoint: endpoint,
	}

	emptydirclone := emptydirclone.New(cfg)
	if err := emptydirclone.Serve(); err != nil {
		panic(err)
	}
}
