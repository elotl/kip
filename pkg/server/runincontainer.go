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
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/nodeclient"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/wsstream"
	vkapi "github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	"k8s.io/klog"
)

type WinSize struct {
	Rows uint16
	Cols uint16
	X    uint16
	Y    uint16
}

func (p *InstanceProvider) RunInContainer(ctx context.Context, namespace, podName, containerName string, cmd []string, attach vkapi.AttachIO) error {
	ctx, span := trace.StartSpan(ctx, "RunInContainer")
	defer span.End()
	_ = addAttributes(ctx, span, namespaceKey, namespace, nameKey, podName, containerNameKey, containerName)
	klog.V(2).Infof("RunInContainer %q %v", podName, cmd)
	tty := attach.TTY()
	stdin := attach.Stdin()
	stdout := attach.Stdout()
	stderr := attach.Stderr()
	resize := attach.Resize()
	params := api.ExecParams{
		PodName:     util.WithNamespace(namespace, podName),
		UnitName:    containerName,
		Command:     cmd,
		Interactive: stdin != nil,
		TTY:         tty,
	}
	ws, err := p.getWS(params.PodName, params.UnitName)
	if err != nil {
		return err
	}
	defer ws.CloseAndCleanup()
	err = sendParams(ws, params)
	if err != nil {
		return err
	}
	sendWinSize(ws, WinSize{
		Cols: 80,
		Rows: 25,
	})
	if resize != nil {
		// Send tty resize messages to the other side.
		go func() {
			for termsize := range resize {
				klog.V(2).Infof("exec requesting window resize %+v", termsize)
				err = sendWinSize(ws, WinSize{
					Cols: termsize.Width,
					Rows: termsize.Height,
				})
				if err != nil {
					klog.Errorf("exec sending window resize: %v", err)
					continue
				}
			}
		}()
	}
	return p.muxToWS(ws, stdin, stdout, stderr, tty)
}

func sendWinSize(ws *wsstream.WSStream, winsize WinSize) error {
	b, err := json.Marshal(winsize)
	if err != nil {
		return util.WrapError(err, "error resizing pty")
	}
	f := wsstream.PackMessage(4, b)
	return ws.WriteRaw(f)
}

func sendParams(ws *wsstream.WSStream, params api.ExecParams) error {
	b, err := json.Marshal(params)
	if err != nil {
		return util.WrapError(err, "marshaling initial exec params")
	}
	return ws.WriteRaw(b)
}

func (p *InstanceProvider) getWS(podName, unitName string) (*wsstream.WSStream, error) {
	node, err := GetNodeForRunningPod(podName, unitName, p.getPodRegistry(), p.getNodeRegistry())
	if err != nil {
		return nil, util.WrapError(err, "could not exec on pod %s", podName)
	}
	itzoEndpoint := nodeclient.ExecEndpoint()
	ws, err := p.ItzoClientFactory.GetWSStream(node.Status.Addresses, itzoEndpoint)
	if err != nil {
		return nil, util.WrapError(err, "could not create websocket stream")
	}
	return ws, nil
}

func (p *InstanceProvider) muxToWS(ws *wsstream.WSStream, stdin io.Reader, stdout, stderr io.WriteCloser, tty bool) error {
	go func() {
		if stdin == nil {
			return
		}
		b := make([]byte, 1)
		for {
			n, err := stdin.Read(b)
			eof := err == io.EOF
			if err != nil && !eof {
				klog.Errorf("exec reading stdin: %v", err)
				return
			}
			if eof && n == 0 {
				klog.V(2).Infof("exec stdin EOF")
				return
			}
			// CRLF conversion if a tty is used.
			if tty && b[0] == '\r' {
				b[0] = '\n'
			}
			f := wsstream.PackMessage(wsstream.StdinChan, b)
			err = ws.WriteRaw(f)
			if err != nil {
				klog.Errorf("exec ws send: %v", err)
				return
			}
			if eof {
				klog.V(2).Infof("exec stdin EOF")
				return
			}
		}
	}()
	for {
		select {
		case <-ws.Closed():
			return nil
		case msg := <-ws.ReadMsg():
			ch, data, err := wsstream.UnpackMessage(msg)
			if err != nil {
				klog.Errorf("exec reading from ws: %v", err)
				continue
			}
			var writer io.WriteCloser
			switch ch {
			case wsstream.StdoutChan:
				writer = stdout
				if writer == nil {
					continue
				}
			case wsstream.StderrChan:
				writer = stderr
				if writer == nil {
					continue
				}
			case wsstream.ExitCodeChan:
				exitCode, err := strconv.Atoi(string(data))
				if err != nil {
					klog.Errorf("exec invalid exit code %v", data)
					continue
				}
				klog.V(2).Infof("exec got exit code %d", exitCode)
				continue
			default:
				klog.Errorf("exec unknown channel %d from ws", ch)
				continue
			}
			_, err = writer.Write(data)
			if err != nil {
				klog.Errorf("exec writing to output %d: %v", ch, err)
				break
			}
		}
	}
}
