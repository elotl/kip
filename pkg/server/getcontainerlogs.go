package server

import (
	"context"
	"fmt"
	"io"

	"github.com/containerd/containerd/log"
	vkapi "github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
)

func (p *InstanceProvider) GetContainerLogs(ctx context.Context, namespace, podName, containerName string, opts vkapi.ContainerLogOpts) (io.ReadCloser, error) {
	ctx, span := trace.StartSpan(ctx, "GetContainerLogs")
	defer span.End()
	ctx = addAttributes(ctx, span, namespaceKey, namespace, nameKey, podName, containerNameKey, containerName)
	log.G(ctx).Infof("GetContainerLogs %q", podName)
	return nil, fmt.Errorf("not implemented")
}
