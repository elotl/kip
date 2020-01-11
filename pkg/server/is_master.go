package server

import (
	"github.com/elotl/cloud-instance-provider/pkg/clientapi"
	"golang.org/x/net/context"
)

func (s InstanceProvider) IsLeader(context context.Context, request *clientapi.IsLeaderRequest) (*clientapi.IsLeaderReply, error) {
	reply := clientapi.IsLeaderReply{
		IsLeader: s.controllerManager.ControllersRunning()}
	return &reply, nil
}
