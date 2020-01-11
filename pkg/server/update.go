package server

import (
	"bytes"
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/yaml"
	"github.com/golang/glog"
	"golang.org/x/net/context"
)

func (s InstanceProvider) Update(context context.Context, request *clientapi.UpdateRequest) (*clientapi.APIReply, error) {
	if !s.controllerManager.ControllersRunning() {
		return notTheLeaderReply(), nil
	}
	_, objectKind, err := VersionAndKind(request.Manifest)
	glog.Infof("Update request for: %s", objectKind)
	if err != nil {
		return errToAPIReply(
			util.WrapError(err, "Error determining manifest kind")), nil
	}
	store, exists := s.KV[objectKind]
	if !exists {
		return errToAPIReply(
			fmt.Errorf("Asked to update unknown object kind: %s", objectKind)), nil
	}
	milpaObj := store.New()
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(request.Manifest), bufferSize)
	err = decoder.Decode(&milpaObj)
	if err != nil {
		return errToAPIReply(
			util.WrapError(err, "Error loading manifest")), nil
	}

	checkObj := store.New()
	unknownKeysErr := yaml.DetectUnknownKeys(bytes.NewReader(request.Manifest), &checkObj, bufferSize)

	if err != nil {
		return errToAPIReply(util.WrapError(err, "Error creating resource")), nil
	}

	replyObj, err := store.Update(milpaObj)
	if err != nil {
		return errToAPIReply(util.WrapError(err, "Error updating resource")), nil
	}
	body, err := s.Encoder.Marshal(replyObj)
	if err != nil {
		return errToAPIReply(util.WrapError(err, "Error updating reply object")), nil
	}
	warningMsg := []byte{}
	if unknownKeysErr != nil {
		warningMsg = []byte(unknownKeysErr.Error())
	}
	reply := clientapi.APIReply{
		Status:  201,
		Body:    body,
		Warning: warningMsg,
	}
	return &reply, nil
}
