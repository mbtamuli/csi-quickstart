package emptydirclone

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

type Config struct {
	Endpoint      string
	Name          string
	NodeID        string
	VendorVersion string
}

type emptyDirClone struct {
	config Config
	logger logr.Logger
}

type AccessType int

const (
	MountAccess AccessType = iota
	BlockAccess
)

func New(config Config, logger logr.Logger) (*emptyDirClone, error) {
	if config.Name == "" {
		return nil, errors.New("no driver name provided")
	}

	if config.NodeID == "" {
		return nil, errors.New("no node id provided")
	}

	if config.Endpoint == "" {
		return nil, errors.New("no driver endpoint provided")
	}

	return &emptyDirClone{
		config: config,
		logger: logger.WithName("emptydirclone"),
	}, nil
}

func (e *emptyDirClone) Serve() error {
	debugLogger := e.logger.WithName("Serve").V(2)

	debugLogger.Info("gRPC server starting", "config", e.config)

	scheme, address, err := parseEndpoint(e.config.Endpoint)
	if err != nil {
		e.logger.Error(err, "failed to parse address")

		return err
	}

	debugLogger.Info("parsed endpoints", "scheme", scheme, "address", address)

	lis, err := net.Listen(scheme, address)
	if err != nil {
		e.logger.Error(err, "failed to listen")
		return err
	}
	defer lis.Close()

	e.logger.Info("gRPC Server listening", "scheme", scheme, "address", address)

	logger := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		grpcLogger := debugLogger.WithName("gRPC")
		if grpcLogger.Enabled() {
			grpcLogger.Info("method", "method", info.FullMethod, "request", req)
		}
		resp, err := handler(ctx, req)
		if err != nil {
			e.logger.Error(err, "method failed", "FullMethod", info.FullMethod)
		}
		if grpcLogger.Enabled() && err == nil {
			grpcLogger.Info("method", "method", info.FullMethod, "response", resp)
		}
		return resp, err
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(logger))
	csi.RegisterIdentityServer(grpcServer, e)
	csi.RegisterNodeServer(grpcServer, e)

	debugLogger.Info("registering servers and listening")

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

// getVolumePath returns the canonical path for hostpath volume
func getVolumePath(volName string) string {
	// return filepath.Join("/csi-data-dir", volName)
	return filepath.Join("/csi/data", volName)
}
