/*
Copyright 2020 Elotl Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/nodeclient"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/wsstream"
	vkapi "github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	"k8s.io/klog"
)

func (p *InstanceProvider) GetContainerLogs(ctx context.Context, namespace, podName, containerName string, opts vkapi.ContainerLogOpts) (io.ReadCloser, error) {
	ctx, span := trace.StartSpan(ctx, "GetContainerLogs")
	defer span.End()
	_ = addAttributes(ctx, span, namespaceKey, namespace, nameKey, podName, containerNameKey, containerName)
	klog.V(5).Infof("GetContainerLogs %+v", opts)
	follow := opts.Follow
	podName = util.WithNamespace(namespace, podName)
	node, err := p.GetNodeForRunningPod(podName, "")
	if !follow || err != nil || node == nil || len(node.Status.Addresses) == 0 {
		klog.V(5).Infof("pulling logs for pod %+v", opts)
		return p.getContainerLogs(podName, containerName, opts)
	}
	klog.V(5).Infof("tailing logs for pod %+v", opts)
	return p.tailContainerLogs(node, podName, containerName, opts)
}

func (p *InstanceProvider) getContainerLogs(podName, containerName string, opts vkapi.ContainerLogOpts) (io.ReadCloser, error) {
	lines := opts.Tail
	limit := opts.LimitBytes
	foundLog, err := p.findLog(podName, containerName, lines, limit)
	if err != nil {
		klog.Errorf("finding logs for %q/%q: %v", podName, containerName, err)
		return nil, util.WrapError(
			err, "finding logs for %q/%q", podName, containerName)
	}
	buf := ioutil.NopCloser(bytes.NewReader([]byte(foundLog.Content)))
	return buf, nil
}

func (p *InstanceProvider) tailContainerLogs(node *api.Node, podName, containerName string, opts vkapi.ContainerLogOpts) (io.ReadCloser, error) {
	withMetadata := opts.Timestamps
	logsPath := nodeclient.StreamLogsEndpoint(containerName, withMetadata)
	ws, err := p.ItzoClientFactory.GetWSStream(node.Status.Addresses, logsPath)
	if err != nil {
		return nil, util.WrapError(
			err, "could not get logs client for pod %q", podName)
	}
	logs, err := p.findLog(podName, containerName, opts.Tail, opts.LimitBytes)
	if err != nil {
		return nil, util.WrapError(
			err, "finding logs for %q/%q", podName, containerName)
	}
	return &containerLogs{
		ws:  ws,
		buf: []byte(logs.Content),
	}, nil
}

type containerLogs struct {
	ws  *wsstream.WSStream
	buf []byte
}

func (l *containerLogs) Read(buf []byte) (int, error) {
	klog.V(5).Infof("reading logs from ws stream")
	n := 0
	if len(l.buf) > 0 {
		klog.V(5).Infof("reading %d bytes from buffer", len(l.buf))
		n = copy(buf, l.buf)
		l.buf = l.buf[n:]
		return n, nil
	}
	select {
	case <-l.ws.Closed():
		klog.V(5).Infof("ws stream closed")
		return 0, io.EOF
	case frame := <-l.ws.ReadMsg():
		n, b, err := wsstream.UnpackMessage(frame)
		if err != nil {
			klog.Errorf("reading ws stream: %v", err)
			return 0, err
		}
		klog.V(5).Infof("read %d bytes from ws stream", n)
		n = copy(buf, b)
		l.buf = append(l.buf[:], b[n:]...)
		klog.V(5).Infof("copied %d bytes from ws stream", n)
		return n, nil
	}
}

func (l *containerLogs) Close() error {
	klog.V(5).Infof("closing ws stream")
	return l.ws.CloseAndCleanup()
}
