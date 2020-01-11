package server

import (
	"io"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/util"
)

func (s InstanceProvider) Attach(stream clientapi.Milpa_AttachServer) error {
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
