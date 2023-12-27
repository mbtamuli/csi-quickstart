package emptydirclone

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

type Config struct {
	Name          string
	Endpoint      string
	VendorVersion string
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
	scheme, address, err := parseEndpoint(e.config.Endpoint)
	if err != nil {
		return err
	}

	lis, err := net.Listen(scheme, address)
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
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", "", fmt.Errorf("could not parse endpoint: %w", err)
	}

	addr := filepath.Join(u.Host, filepath.FromSlash(u.Path))

	scheme := strings.ToLower(u.Scheme)
	switch scheme {
	case "tcp":
	case "unix":
		addr = filepath.Join("/", addr)
		if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
			return "", "", fmt.Errorf("could not remove unix domain socket %q: %w", addr, err)
		}
	default:
		return "", "", fmt.Errorf("unsupported protocol: %s", scheme)
	}

	return scheme, addr, nil
}
