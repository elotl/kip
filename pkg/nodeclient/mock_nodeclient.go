package nodeclient

// Fake node dispenser
// fake AppRunner
//

import (
	"io"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/wsstream"
)

func NewMockItzoClientFactory() *MockItzoClientFactory {
	return &MockItzoClientFactory{
		Health: func() error {
			return nil
		},
		Logs: func(unit string, lines, bytes int) ([]byte, error) {
			return []byte("logs"), nil
		},
		File: func(unit string, lines, bytes int) ([]byte, error) {
			return []byte("file"), nil
		},
		Resize: func() error {
			return nil
		},
		Status: func() (*api.PodStatusReply, error) {
			return &api.PodStatusReply{}, nil
		},
		Update: func(pp api.PodParameters) error {
			return nil
		},
		DeployPackage: func(pod, name string, data io.Reader) error {
			return nil
		},
	}
}

type MockItzoClientFactory struct {
	Health        func() error
	Logs          func(unit string, lines, bytes int) ([]byte, error)
	File          func(unit string, lines, bytes int) ([]byte, error)
	Resize        func() error
	Status        func() (*api.PodStatusReply, error)
	Update        func(pp api.PodParameters) error
	DeployPackage func(pod, name string, data io.Reader) error
}

// screw it, make the factory implement the interface as well...
func (a *MockItzoClientFactory) GetClient(addy []api.NetworkAddress) NodeClient {
	return a
}

func (a *MockItzoClientFactory) GetWSStream(addy []api.NetworkAddress, path string) (*wsstream.WSStream, error) {
	return nil, nil
}

// screw it, make the factory implement the interface as well...
func (a *MockItzoClientFactory) DeleteClient(addy []api.NetworkAddress) {
}

func (a *MockItzoClientFactory) Healthcheck() error {
	return a.Health()
}

func (a *MockItzoClientFactory) GetLogs(unit string, lines, bytes int) ([]byte, error) {
	return a.Logs(unit, lines, bytes)
}

func (a *MockItzoClientFactory) GetFile(unit string, lines, bytes int) ([]byte, error) {
	return a.Logs(unit, lines, bytes)
}

func (a *MockItzoClientFactory) ResizeVolume() error {
	return a.Resize()
}

func (a *MockItzoClientFactory) GetStatus() (*api.PodStatusReply, error) {
	return a.Status()
}

func (a *MockItzoClientFactory) UpdateUnits(pp api.PodParameters) error {
	return a.Update(pp)
}

func (a *MockItzoClientFactory) RunCmd(cmdParams api.RunCmdParams) (string, error) {
	return "", nil
}

func (a *MockItzoClientFactory) Deploy(pod, name string, data io.Reader) error {
	return a.DeployPackage(pod, name, data)
}
