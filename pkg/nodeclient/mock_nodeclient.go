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

package nodeclient

// Fake node dispenser
// fake AppRunner
//

import (
	"io"

	"github.com/elotl/kip/pkg/api"
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

func (a *MockItzoClientFactory) Deploy(pod, name string, data io.Reader) error {
	return a.DeployPackage(pod, name, data)
}
