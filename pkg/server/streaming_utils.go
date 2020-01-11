package server

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/docker/libkv/store"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/golang/glog"
	"golang.org/x/net/context"
)

type SendRecver interface {
	Send(*clientapi.StreamMsg) error
	Recv() (*clientapi.StreamMsg, error)
	Context() context.Context
}

func (s InstanceProvider) GetNodeForRunningPod(podName, unitName string) (*api.Node, error) {
	reg, exists := s.KV["Pod"]
	if !exists {
		return nil, fmt.Errorf("Fatal error: can't find pod registry in storage")
	}
	podRegistry := reg.(*registry.PodRegistry)
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
		allUnits := api.AllPodUnits(pod)
		names := make([]string, len(allUnits))
		for i := 0; i < len(allUnits); i++ {
			names[i] = allUnits[i].Name
		}
		if !util.StringInSlice(unitName, names) {
			return nil, fmt.Errorf("Could not find a unit named %s. Pod unit names: %s",
				unitName, strings.Join(names, ", "))
		}
	}
	if pod.Status.BoundNodeName == "" {
		return nil, fmt.Errorf("pod is unbound")
	}
	reg, exists = s.KV["Node"]
	if !exists {
		return nil, fmt.Errorf("can't find node registry in storage")
	}
	nodeRegistry := reg.(*registry.NodeRegistry)
	node, err := nodeRegistry.GetNode(pod.Status.BoundNodeName)
	if err == store.ErrKeyNotFound {
		return nil, fmt.Errorf("Could not find node in registry")
	} else if err != nil {
		return nil, util.WrapError(
			err, "Could not get node from storage")
	}
	return node, nil
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
	defer ws.CloseAndCleanup()

	//forward first packet to client
	err = ws.WriteRaw(paramsMsg.Data)
	if err != nil {
		if err != io.EOF {
			wrappedErr := util.WrapError(err, "Error in websocket send")
			glog.Error(wrappedErr)
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
					glog.Errorf("Error in grpc receive: %v", err)
				}
				return
			}
			err = ws.WriteRaw(clientData.Data)
			if err != nil {
				if err != io.EOF {
					glog.Errorf("Error in websocket send: %v", err)
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
