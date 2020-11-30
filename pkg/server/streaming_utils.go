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
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/docker/libkv/store"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/clientapi"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util"
	"golang.org/x/net/context"
	"k8s.io/klog"
)

type SendRecver interface {
	Send(*clientapi.StreamMsg) error
	Recv() (*clientapi.StreamMsg, error)
	Context() context.Context
}

func GetNodeForRunningPod(podName, unitName string, podRegistry *registry.PodRegistry, nodeRegistry *registry.NodeRegistry) (*api.Node, error) {
	pod, err := podRegistry.GetPod(podName)
	if err == store.ErrKeyNotFound {
		return nil, fmt.Errorf("Could not find pod in registry")
	} else if err != nil {
		return nil, util.WrapError(
			err, "Could not get pod from storage")
	}
	// if the pod is creating or dispatching, return an error
	// otherwise if not running, look for the saved logs
	if pod.Status.Phase != api.PodRunning {
		return nil, fmt.Errorf("Pod is not running")
	}
	if unitName != "" {
		names := make([]string, 0, len(pod.Spec.InitUnits)+len(pod.Spec.Units))
		api.ForAllUnits(pod, func(unit *api.Unit) {
			names = append(names, unit.Name)
		})
		if !util.StringInSlice(unitName, names) {
			return nil, fmt.Errorf("Could not find a unit named %s. Pod unit names: %s",
				unitName, strings.Join(names, ", "))
		}
	}
	if pod.Status.BoundNodeName == "" {
		return nil, fmt.Errorf("pod is unbound")
	}
	node, err := nodeRegistry.GetNode(pod.Status.BoundNodeName)
	if err == store.ErrKeyNotFound {
		return nil, fmt.Errorf("Could not find node in registry")
	} else if err != nil {
		return nil, util.WrapError(
			err, "Could not get node from storage")
	}
	return node, nil
}

func (s InstanceProvider) GetNodeForRunningPod(podName, unitName string) (*api.Node, error) {
	reg, exists := s.Registries["Pod"]
	if !exists {
		return nil, fmt.Errorf("can't find pod registry in storage")
	}
	podRegistry := reg.(*registry.PodRegistry)
	reg, exists = s.Registries["Node"]
	if !exists {
		return nil, fmt.Errorf("can't find node registry in storage")
	}
	nodeRegistry := reg.(*registry.NodeRegistry)
	return GetNodeForRunningPod(podName, unitName, podRegistry, nodeRegistry)
}

func getInitialParams(stream SendRecver, params interface{}) (*clientapi.StreamMsg, error) {
	paramsMsg, err := stream.Recv()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(paramsMsg.Data, &params)
	if err != nil {
		return nil, util.WrapError(err, "Error getting initial parameters message")
	}
	return paramsMsg, err
}

func (s InstanceProvider) grpcToWSPump(stream SendRecver, addresses []api.NetworkAddress, itzoEndpoint string, paramsMsg *clientapi.StreamMsg) error {
	ws, err := s.ItzoClientFactory.GetWSStream(addresses, itzoEndpoint)
	if err != nil {
		return util.WrapError(err, "Could not create websocket stream")
	}
	defer ws.CloseAndCleanup() //nolint

	//forward first packet to client
	err = ws.WriteRaw(paramsMsg.Data)
	if err != nil {
		if err != io.EOF {
			wrappedErr := util.WrapError(err, "Error in websocket send")
			klog.Error(wrappedErr)
			return wrappedErr
		}
		return nil
	}

	// BUG: we aren't getting the return error value from this
	// goroutine
	go func() {
		for {
			clientData, err := stream.Recv()
			if err == io.EOF {
				return
			} else if err != nil {
				select {
				case <-stream.Context().Done():
					// yuck, need to detect context being cancelled???
					// I have a feeling I'm doing this wrong...
				default:
					klog.Errorf("Error in grpc receive: %v", err)
				}
				return
			}
			err = ws.WriteRaw(clientData.Data)
			if err != nil {
				if err != io.EOF {
					klog.Errorf("Error in websocket send: %v", err)
				}
				return
			}
		}
	}()

	for {
		select {
		case <-ws.Closed():
			return nil
		case msg := <-ws.ReadMsg():
			pb := clientapi.StreamMsg{Data: []byte(msg)}
			if err := stream.Send(&pb); err != nil {
				return err
			}
		case <-time.After(5 * time.Second):
			// In case this connection never writes again
			// (e.g. tailing a silent log), we'll send empty messages
			// every X seconds to detect if the client is still there.
			// If we don't do this, we never detect the client has
			// gone away, we just hold the go-routine open until we
			// write to the client again
			msg := clientapi.StreamMsg{}
			if err := stream.Send(&msg); err != nil {
				return err
			}
		}
	}
}
