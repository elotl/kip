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
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/yaml"
	"golang.org/x/net/context"
	"k8s.io/klog"
)

func (s InstanceProvider) Update(context context.Context, request *clientapi.UpdateRequest) (*clientapi.APIReply, error) {
	if !s.controllerManager.ControllersRunning() {
		return notTheLeaderReply(), nil
	}
	_, objectKind, err := VersionAndKind(request.Manifest)
	klog.V(2).Infof("Update request for: %s", objectKind)
	if err != nil {
		return errToAPIReply(
			util.WrapError(err, "Error determining manifest kind")), nil
	}
	store, exists := s.Registries[objectKind]
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
