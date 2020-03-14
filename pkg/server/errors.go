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

	"github.com/docker/libkv/store"
	"github.com/elotl/kip/pkg/api/validation"
	"github.com/elotl/kip/pkg/clientapi"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/klog"
)

const (
	StatusNotFound            = 404
	StatusAlreadyExists       = 409
	MisdirectedRequest        = 421
	StatusUnprocessableEntity = 422
	StatusServerError         = 500
)

// Special helper to assist in the unwrapping of util.WrappedError type
// In that case, the wrapped error doesn't have the right type and the
// correct type doesn't have the wrapped error message.  So we create
// this helper and do the unwrapping of type and message in the caller
func errToAPIReplyHelper(origErr error, errMsg string) *clientapi.APIReply {
	switch origErr.(type) {
	case validation.ValidationError:
		return &clientapi.APIReply{
			Status: StatusUnprocessableEntity,
			Body:   []byte(errMsg),
		}
	default:
		output, marshallErr := json.Marshal(errMsg)
		if marshallErr != nil {
			klog.Errorf("Could not marshal error: %v", marshallErr)
			// ain't json but thats ok, dont' call marshall recursively
			output = []byte(marshallErr.Error())
		}

		var status int32 = StatusServerError
		if origErr == store.ErrKeyNotFound {
			status = StatusNotFound
		} else if origErr == registry.ErrAlreadyExists {
			status = StatusAlreadyExists
		}
		return &clientapi.APIReply{
			Status: status,
			Body:   output,
		}
	}
}

func errToAPIReply(e error) *clientapi.APIReply {
	if we, ok := e.(*util.WrappedError); ok {
		return errToAPIReplyHelper(we.Cause(), we.Error())
	} else {
		return errToAPIReplyHelper(e, e.Error())
	}
}

func notTheLeaderReply() *clientapi.APIReply {
	// send back a redirection
	msg := "Server is not leader, please try another endpoint"
	output, marshallErr := json.Marshal(msg)
	if marshallErr != nil {
		output = []byte(marshallErr.Error())
	}
	return &clientapi.APIReply{
		Status: MisdirectedRequest,
		Body:   output,
	}
}
