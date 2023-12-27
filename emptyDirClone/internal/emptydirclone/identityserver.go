package emptydirclone

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

func (e *emptyDirClone) GetPluginInfo(context.Context, *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return &csi.GetPluginInfoResponse{
		Name:          e.config.Name,
		VendorVersion: e.config.VendorVersion,
	}, nil
}

func (e *emptyDirClone) GetPluginCapabilities(context.Context, *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	return &csi.GetPluginCapabilitiesResponse{}, nil
}

func (e *emptyDirClone) Probe(context.Context, *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	return &csi.ProbeResponse{}, nil
}
