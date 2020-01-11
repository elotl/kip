package server

import (
	"encoding/json"

	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"golang.org/x/net/context"
)

func (s InstanceProvider) GetVersion(context context.Context, request *clientapi.VersionRequest) (*clientapi.VersionReply, error) {
	v := util.GetVersionInfo()
	b, err := json.Marshal(v)
	reply := clientapi.VersionReply{
		VersionInfo: b,
	}
	return &reply, err
}
