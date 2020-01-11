package server

import (
	"io"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/nodeclient"
	"github.com/elotl/cloud-instance-provider/pkg/util"
)

func (s InstanceProvider) Exec(stream clientapi.Milpa_ExecServer) error {
	var params api.ExecParams
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
		return util.WrapError(err, "Could not exec on pod %s", podName)
	}

	itzoPath := nodeclient.ExecEndpoint()
	err = s.grpcToWSPump(stream, node.Status.Addresses, itzoPath, paramsMsg)
	if err != nil {
		err = util.WrapError(err, "Could not run exec on pod %s", podName)
	}
	return err
}
