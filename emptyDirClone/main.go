package main

import (
	"flag"

	"github.com/mbtamuli/emptyDirClone/internal/emptydirclone"
)

func main() {
	var endpoint string

	flag.StringVar(&endpoint, "endpoint", "unix:///csi/csi.sock", "Endpoint for the gRPC server to listen on. Default: \"unix:///csi/csi.sock\"")
	flag.Parse()

	cfg := emptydirclone.Config{
		Endpoint: endpoint,
	}

	emptydirclone := emptydirclone.New(cfg)
	if err := emptydirclone.Serve(); err != nil {
		panic(err)
	}
}
