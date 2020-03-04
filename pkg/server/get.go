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
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/util"
)

func (s InstanceProvider) Get(context context.Context, request *clientapi.GetRequest) (*clientapi.APIReply, error) {
	kind := string(request.Kind)
	switch kind {
	case "pods", "nodes", "services", "replicasets", "deployments", "secrets", "events":
		kind = kind[0 : len(kind)-1]
	}
	name := string(request.Name)
	store, exists := s.Registries[kind]
	if !exists {
		return errToAPIReply(
			fmt.Errorf("Asked to get unknown object kind: %s", kind)), nil
	}
	var replyObj api.MilpaObject
	var err error
	if name == "" {
		replyObj, err = store.List()
	} else {
		replyObj, err = store.Get(name)
	}
	replyObj = filterReplyObject(replyObj)
	if err != nil {
		return errToAPIReply(util.WrapError(err, "Error getting resource")), nil
	}
	body, err := s.Encoder.Marshal(replyObj)
	if err != nil {
		return errToAPIReply(util.WrapError(err, "Error serializing reply object")), nil
	}
	reply := clientapi.APIReply{
		Status: 200,
		Body:   body,
	}
	return &reply, nil
}
