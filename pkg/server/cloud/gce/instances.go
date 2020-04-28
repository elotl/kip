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

package gce

import (
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
)

func (c *gceClient) StartNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	return nil, NI()
}

func (c *gceClient) StartSpotNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	return nil, NI()
}

func (c *gceClient) WaitForRunning(node *api.Node) ([]api.NetworkAddress, error) {
	return nil, NI()
}

func (c *gceClient) StopInstance(instanceID string) error {
	return NI()
}

func (c *gceClient) ResizeVolume(node *api.Node, size int64) (error, bool) {
	return NI(), false
}

func (c *gceClient) SetSustainedCPU(node *api.Node, enabled bool) error {
	return nil
}

func (c *gceClient) ListInstancesFilterID(ids []string) ([]cloud.CloudInstance, error) {
	return nil, NI()
}

func (c *gceClient) ListInstances() ([]cloud.CloudInstance, error) {
	return nil, NI()
}

func (c *gceClient) AddInstanceTags(iid string, labels map[string]string) error {
	return NI()
}

func (c *gceClient) GetImageID(spec cloud.BootImageSpec) (string, error) {
	return "", NI()
}

func (c *gceClient) AssignInstanceProfile(node *api.Node, instanceProfile string) error {
	return NI()
}
