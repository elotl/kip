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
	"fmt"
	"strings"
	"time"

	"github.com/docker/libkv/store"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/clientapi"
	"github.com/elotl/kip/pkg/nodeclient"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util"
	"golang.org/x/net/context"
	"k8s.io/klog"
)

// Logs requests can take a couple of forms:
//
//   logs podname unitname [lines] [limitbytes]
//   logs nodename filepath [lines] [limitbytes]
//
// both of these can pull logs from currently running pods or nodes if
// the pod or node are not found in the currently running set of pods
// and nodes, we look in the logs registry for old logs.
//
// Eventually I can see multiple node logs being saved so the unitname
// parameter can be used to specify the path to the logfile on the
// node (and thus the name of the logfile)

func (s InstanceProvider) getPodLog(podName, itemName string, lines, bytes int) (*api.LogFile, error) {
	reg, exists := s.Registries["Pod"]
	if !exists {
		return nil, fmt.Errorf("Fatal error: can't find pod registry in storage")
	}
	podRegistry := reg.(*registry.PodRegistry)
	pod, err := podRegistry.GetPod(podName)
	if err == store.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, util.WrapError(
			err, "Could not get pod %s from storage", podName)
	}

	if pod.Status.Phase != api.PodRunning {
		// if the pod isn't running, return nil in case there are
		// saved logs
		return nil, nil
	}

	nodeName := pod.Status.BoundNodeName
	if nodeName == "" {
		return nil, nil
	}
	log, err := s.getNodeLog(nodeName, itemName, true, lines, bytes)
	if err != nil {
		return nil, err
	}
	// need to rewrite what we got out of here
	log.Name = podName + "/" + itemName
	log.ParentObject = api.ToObjectReference(pod)
	return log, nil
}

func (s InstanceProvider) getNodeLog(nodeName, itemName string, isUnit bool, lines, bytes int) (*api.LogFile, error) {
	reg, exists := s.Registries["Node"]
	if !exists {
		return nil, fmt.Errorf("Fatal error: can't find node registry in storage")
	}
	nodeRegistry := reg.(*registry.NodeRegistry)
	node, err := nodeRegistry.GetNode(nodeName)

	if err == store.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("Could not get node %s from storage", nodeName)

	} else if len(node.Status.Addresses) == 0 {
		return nil, fmt.Errorf("Log for node %s is not available", nodeName)
	}

	client := s.ItzoClientFactory.GetClient(node.Status.Addresses)
	var data []byte
	if isUnit {
		data, err = client.GetLogs(itemName, lines, bytes)
	} else {
		data, err = client.GetFile("/var/log/itzo/itzo.log", lines, bytes)
	}
	if err != nil {
		return nil, util.WrapError(
			err, "Error getting log from node %s", nodeName)
	}

	log := api.NewLogFile()
	log.Name = nodeName + "/" + itemName
	log.ParentObject = api.ToObjectReference(node)
	log.Content = string(data)
	return log, nil
}

func (s InstanceProvider) getLogFromRegistry(resourceName, itemName string, lines, bytes int) (*api.LogFile, error) {
	reg, exists := s.Registries["Log"]
	if !exists {
		return nil, fmt.Errorf("Fatal error: can't find log registry in storage")
	}
	logRegistry := reg.(*registry.LogRegistry)
	// we'll assume it's the node log
	var log *api.LogFile
	var logs *api.LogFileList
	var err error
	if itemName == "" {
		logs, err = logRegistry.ListLogs(resourceName, "")
		if err == nil {
			if len(logs.Items) == 1 {
				log = logs.Items[0]
			} else if len(logs.Items) > 1 {
				logKind := logs.Items[0].ParentObject.Kind
				if logKind == "Pod" {
					err = fmt.Errorf("A unit name is required when getting saved logs from a pod with multiple units")
				} else {
					// Shouldn't happen
					log = logs.Items[0]
				}
			}
		}
	} else {
		// If we have an item/unit name, return that, otherwise see if
		// we only have one unit for the pod
		log, err = logRegistry.GetLog(resourceName, itemName)
	}
	if err != nil {
		return nil, util.WrapError(
			err, "Could not get log %s/%s from storage", resourceName, itemName)
	} else if log == nil {
		return nil, fmt.Errorf("Could not find log in storage")
	}
	if lines != 0 {
		// we only save a couple KB of logs so this shouldn't be
		// all that bad
		parts := strings.Split(log.Content, "\n")
		if lines < len(parts) {
			startLine := len(parts) - lines
			parts = parts[startLine:]
		}
		log.Content = strings.Join(parts, "\n")
	}
	if bytes != 0 {
		if bytes < len(log.Content) {
			startByte := len(log.Content) - bytes
			log.Content = log.Content[startByte:]
		}
	}
	return log, nil
}

func (s InstanceProvider) findLog(resourceName, itemName string, lines, bytes int) (*api.LogFile, error) {
	log, err := s.getPodLog(resourceName, itemName, lines, bytes)
	if err != nil {
		return nil, util.WrapError(
			err, "Error getting log %s from pod", resourceName)
	}
	if log != nil {
		return log, err
	}

	log, err = s.getNodeLog(resourceName, itemName, false, lines, bytes)
	if err != nil {
		return nil, util.WrapError(
			err, "Error getting log from node %s", resourceName)
	}
	if log != nil {
		return log, err
	}

	log, err = s.getLogFromRegistry(resourceName, itemName, lines, bytes)
	if err != nil {
		return nil, util.WrapError(
			err, "Error getting older log from registry %s", resourceName)
	}
	if log != nil {
		return log, err
	}
	return nil, fmt.Errorf("Could not find the requested log")
}

func (s InstanceProvider) GetLogs(context context.Context, request *clientapi.LogsRequest) (*clientapi.APIReply, error) {
	resourceName := request.ResourceName
	itemName := request.ItemName
	lines := int(request.Lines)
	bytes := int(request.Limitbytes)

	klog.V(2).Infof("Getting logs from %s/%s (max lines %d; limitbytes %d)",
		resourceName, itemName, lines, bytes)

	foundLog, err := s.findLog(resourceName, itemName, lines, bytes)
	if err != nil {
		return errToAPIReply(err), nil
	}
	body, err := s.Encoder.Marshal(foundLog)
	if err != nil {
		return errToAPIReply(
			util.WrapError(err, "Error serializing reply object")), nil
	}

	reply := clientapi.APIReply{
		Status: 200,
		Body:   body,
	}
	return &reply, nil
}

func (s InstanceProvider) StreamLogs(slr *clientapi.StreamLogsRequest, stream clientapi.Kip_StreamLogsServer) error {
	podName := slr.Pod
	unitName := slr.Unit
	withMetadata := slr.Metadata
	node, err := s.GetNodeForRunningPod(podName, "")
	if err != nil || node == nil || len(node.Status.Addresses) == 0 {
		return util.WrapError(
			err, "Could not get logs for pod %s", podName)
	}
	logsPath := nodeclient.StreamLogsEndpoint(unitName, withMetadata)
	ws, err := s.ItzoClientFactory.GetWSStream(node.Status.Addresses, logsPath)
	if err != nil {
		return util.WrapError(
			err, "Could not get logs client for pod %s", podName)
	}
	defer ws.CloseAndCleanup()
	for {
		select {
		case <-ws.Closed():
			return nil
		case msg := <-ws.ReadMsg():
			pb := clientapi.StreamMsg{
				Data: []byte(msg),
			}
			if err := stream.Send(&pb); err != nil {
				return err
			}
		case <-time.After(5 * time.Second):
			// This is relevant because this is a one-way stream (at
			// least today it is...).  In case the itzo server never
			// writes again (e.g. tailing a silent log), we'll send
			// empty messages every X seconds to detect if the client
			// is still there. If we don't do this, we never detect
			// the client has gone away, and we just hold the
			// go-routine open until we write to the client again (and
			// that write may never happen)
			msg := clientapi.StreamMsg{}
			if err := stream.Send(&msg); err != nil {
				return err
			}
		}
	}
	// NOTREACHED
}
