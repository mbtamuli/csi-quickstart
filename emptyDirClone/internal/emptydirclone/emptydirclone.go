package emptydirclone

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

type Config struct {
	Endpoint string
}

type emptyDirClone struct {
	config Config
}

func New(config Config) *emptyDirClone {
	return &emptyDirClone{
		config: config,
	}
}

func (e *emptyDirClone) Serve() error {
	network, address, err := parseEndpoint(e.config.Endpoint)
	if err != nil {
		return err
	}

	lis, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	csi.RegisterIdentityServer(grpcServer, e)
	csi.RegisterNodeServer(grpcServer, e)

	log.Println("Listening on", e.config.Endpoint)
	return grpcServer.Serve(lis)
}

func parseEndpoint(endpoint string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(endpoint), "unix://") {
		s := strings.SplitN(endpoint, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
		return "", "", fmt.Errorf("invalid endpoint: %v", endpoint)
	}

	// Assume everything else is a file path for a Unix Domain Socket.
	return "unix", endpoint, nil
}
