package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/golang/glog"
)

func (s InstanceProvider) deploy(podName, pkgName string, pkgData io.Reader) error {
	reg, exists := s.KV["Pod"]
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
	reg, exists = s.KV["Node"]
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

func (s InstanceProvider) Deploy(stream clientapi.Milpa_DeployServer) error {
	pod := ""
	name := ""
	tmpfile, err := ioutil.TempFile("", "milpadeploy")
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
			glog.Errorf("Failed to receive deploy request: %v", err)
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
	glog.Infof("Deployed package %s for %s", name, pod)
	return stream.SendAndClose(&reply)
}
