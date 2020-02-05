package server

import (
	"context"
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"k8s.io/klog"
)

func (s InstanceProvider) deleteHelper(kind, name string, cascade bool) (api.MilpaObject, error) {
	store, exists := s.Registries[kind]
	if !exists {
		return nil, fmt.Errorf("Asked to delete unknown object kind: %s", kind)
	}
	replyObj, err := store.Delete(name)
	if err != nil {
		return nil, util.WrapError(err, "Error deleting object from registry")
	}
	return replyObj, nil
}

func (s InstanceProvider) Delete(context context.Context, request *clientapi.DeleteRequest) (*clientapi.APIReply, error) {
	if !s.controllerManager.ControllersRunning() {
		return notTheLeaderReply(), nil
	}

	kind := string(request.Kind)
	switch kind {
	case "pods", "nodes", "services", "secrets":
		kind = kind[0 : len(kind)-1]
	}
	name := string(request.Name)
	klog.V(2).Infof("Delete request for: %s - %s", kind, name)
	replyObj, err := s.deleteHelper(kind, name, request.Cascade)

	if err != nil {
		return errToAPIReply(util.WrapError(err, "Error deleting resource")), nil
	}
	replyObj = filterReplyObject(replyObj)
	body, err := s.Encoder.Marshal(replyObj)
	if err != nil {
		return errToAPIReply(util.WrapError(err, "Error serializing reply object")), nil
	}
	reply := clientapi.APIReply{
		Status: 202,
		Body:   body,
	}
	return &reply, nil
}
