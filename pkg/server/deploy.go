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
	"io"
	"io/ioutil"
	"os"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"k8s.io/klog"
)

func (s InstanceProvider) deploy(podName, pkgName string, pkgData io.Reader) error {
	reg, exists := s.Registries["Pod"]
	if !exists {
		return fmt.Errorf("Fatal error: can't find pod registry in storage")
	}
	podRegistry := reg.(*registry.PodRegistry)
	pod, err := podRegistry.GetPod(podName)
	if err != nil {
		return util.WrapError(
			err, "Could not get pod %s from storage", podName)
	}
	if pod.Status.BoundNodeName == "" {
		return fmt.Errorf("Pod %s is not bound to any node", podName)
	}
	reg, exists = s.Registries["Node"]
	if !exists {
		return fmt.Errorf("Fatal error: can't find node registry in storage")
	}
	nodeRegistry := reg.(*registry.NodeRegistry)
	node, err := nodeRegistry.GetNode(pod.Status.BoundNodeName)
	if err != nil {
		return util.WrapError(
			err, "Could not get node %s from storage", pod.Status.BoundNodeName)
	}
	client := s.ItzoClientFactory.GetClient(node.Status.Addresses)
	err = client.Deploy(podName, pkgName, pkgData)
	if err != nil {
		return util.WrapError(
			err, "Error deploying package %s for %s: %v", pkgName, podName, err)
	}
	return nil
}

func (s InstanceProvider) Deploy(stream clientapi.Kip_DeployServer) error {
	pod := ""
	name := ""
	tmpfile, err := ioutil.TempFile("", "kipdeploy")
	if err != nil {
		reply := clientapi.APIReply{
			Status: 500,
			Body: []byte(
				fmt.Sprintf("{\"error\": \"saving package data: %v\"}", err)),
		}
		return stream.SendAndClose(&reply)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.Errorf("Failed to receive deploy request: %v", err)
			return util.WrapError(err, "Failed to receive deploy request")
		}
		pod = req.ResourceName
		name = req.ItemName
		_, err = tmpfile.Write(req.PackageData)
		if err != nil {
			reply := clientapi.APIReply{
				Status: 500,
				Body: []byte(
					fmt.Sprintf("{\"error\": \"writing package data: %v\"}", err)),
			}
			return stream.SendAndClose(&reply)
		}
	}
	_, err = tmpfile.Seek(0, 0)
	if err != nil {
		reply := clientapi.APIReply{
			Status: 500,
			Body: []byte(
				fmt.Sprintf("{\"error\": \"seeking in tempfile: %v\"}", err)),
		}
		return stream.SendAndClose(&reply)
	}
	err = s.deploy(pod, name, tmpfile)
	if err != nil {
		reply := clientapi.APIReply{
			Status: 500,
			Body: []byte(
				fmt.Sprintf("{\"error\": \"deploying package: %v\"}", err)),
		}
		return stream.SendAndClose(&reply)
	}
	reply := clientapi.APIReply{
		Status: 200,
		Body:   []byte("{}"),
	}
	klog.V(2).Infof("Deployed package %s for %s", name, pod)
	return stream.SendAndClose(&reply)
}
