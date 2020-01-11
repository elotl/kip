package nodeclient

import (
	"io"

	"github.com/elotl/cloud-instance-provider/pkg/api"
)

type NodeClient interface {
	Healthcheck() error
	GetLogs(unit string, lines, bytes int) ([]byte, error)
	GetFile(filepath string, lines, bytes int) ([]byte, error)
	ResizeVolume() error
	GetStatus() (*api.PodStatusReply, error)
	UpdateUnits(api.PodParameters) error
	Deploy(pod, name string, data io.Reader) error
	RunCmd(cmd api.RunCmdParams) (string, error)
}
