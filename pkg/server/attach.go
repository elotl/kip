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
	"io"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/clientapi"
	"github.com/elotl/kip/pkg/nodeclient"
	"github.com/elotl/kip/pkg/util"
)

func (s InstanceProvider) Attach(stream clientapi.Kip_AttachServer) error {
	var params api.AttachParams
	paramsMsg, err := getInitialParams(stream, &params)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	podName := params.PodName
	node, err := s.GetNodeForRunningPod(podName, params.UnitName)
	if err != nil {
		return util.WrapError(err, "Could not get running pod %s", podName)
	}
	itzoPath := nodeclient.AttachEndpoint()
	err = s.grpcToWSPump(stream, node.Status.Addresses, itzoPath, paramsMsg)
	if err != nil {
		err = util.WrapError(err, "Could not run attach on pod %s", podName)
	}
	return err
}
