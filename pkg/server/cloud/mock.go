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

package cloud

import (
	"fmt"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util"
)

type MockCloudClient struct {
	Instances          map[string]CloudInstance
	ContainerInstances map[string]ContainerInstance

	ControllerID string
	InsideVPC    bool
	VPCCIDRs     []string
	Subnet       SubnetAttributes

	Starter             func(node *api.Node, image Image, metadata, iamPermissions string) (string, error)
	SpotStarter         func(node *api.Node, image Image, metadata, iamPermissions string) (string, error)
	Stopper             func(instanceID string) error
	Waiter              func(node *api.Node) ([]api.NetworkAddress, error)
	Lister              func() ([]CloudInstance, error)
	Resizer             func(node *api.Node, size int64) (error, bool)
	ContainerAuthorizer func(string) (string, string, error)
	ImageGetter         func(BootImageSpec) (Image, error)

	InstanceListerFilter func([]string) ([]CloudInstance, error)
	InstanceLister       func() ([]CloudInstance, error)

	DNSInfoGetter func() ([]string, []string, error)

	RouteRemover func(string, string) error
	RouteAdder   func(string, string) error

	AvailabilityChecker func() (bool, error)

	// Container Instance Funcs
	ContainerClusterEnsurer          func() error
	ContainerInstanceLister          func() ([]ContainerInstance, error)
	ContainerInstanceListerFilter    func(instIDs []string) ([]ContainerInstance, error)
	ContainerInstancesStatusesGetter func(instIDs []string) (map[string][]api.UnitStatus, error)
	ContainerInstanceRunner          func(*api.Pod) (string, error)
	ContainerInstanceStopper         func(string) error
	ContainerInstanceWaiter          func(*api.Pod) (*api.Pod, error)

	InstanceParameterAdder   func(instanceID, key, value string, isSecret bool) error
	InstanceParameterRemover func(instanceID, key string) error
}

func (m *MockCloudClient) AddInstanceParameter(instanceID, key, value string, isSecret bool) error {
	return m.InstanceParameterAdder(instanceID, key, value, isSecret)
}

func (m *MockCloudClient) DeleteInstanceParameter(instanceID, key string) error {
	return m.InstanceParameterRemover(instanceID, key)
}

func (m *MockCloudClient) SetBootSecurityGroupIDs([]string) {
}

func (m *MockCloudClient) GetBootSecurityGroupIDs() []string {
	return nil
}

func (m *MockCloudClient) StartNode(node *api.Node, image Image, metadata, iamPermissions string) (string, error) {
	return m.Starter(node, image, metadata, iamPermissions)
}

func (m *MockCloudClient) StartSpotNode(node *api.Node, image Image, metadata, iamPermissions string) (string, error) {
	return m.SpotStarter(node, image, metadata, iamPermissions)
}

func (m *MockCloudClient) StopInstance(instanceID string) error {
	return m.Stopper(instanceID)
}

func (m *MockCloudClient) WaitForRunning(node *api.Node) ([]api.NetworkAddress, error) {
	return m.Waiter(node)
}

func (m *MockCloudClient) ResizeVolume(node *api.Node, size int64) (error, bool) {
	return m.Resizer(node, size)
}

func (m *MockCloudClient) GetRegistryAuth(image string) (string, string, error) {
	return m.ContainerAuthorizer(image)
}

func (m *MockCloudClient) GetImage(spec BootImageSpec) (Image, error) {
	return m.ImageGetter(spec)
}

func (m *MockCloudClient) SetSustainedCPU(n *api.Node, enabled bool) error {
	return nil
}

func (m *MockCloudClient) AddInstanceTags(string, map[string]string) error {
	return nil
}

func (c *MockCloudClient) IsAvailable() (bool, error) {
	return c.AvailabilityChecker()
}

func (c *MockCloudClient) EnsureMilpaSecurityGroups([]string, []string) error {
	return nil
}

func (c *MockCloudClient) ListInstancesFilterID(iid []string) ([]CloudInstance, error) {
	return c.InstanceListerFilter(iid)
}

func (c *MockCloudClient) ListInstances() ([]CloudInstance, error) {
	return c.InstanceLister()
}

func (e *MockCloudClient) CreateSGName(svcName string) string {
	return fmt.Sprintf("%s.%s.%s", e.ControllerID, "default", svcName)
}

func (e *MockCloudClient) ConnectWithPublicIPs() bool {
	return e.InsideVPC
}

func (e *MockCloudClient) ModifySourceDestinationCheck(iid string, enable bool) error {
	return nil
}

func (e *MockCloudClient) GetDNSInfo() ([]string, []string, error) {
	return e.DNSInfoGetter()
}

func (e *MockCloudClient) RemoveRoute(destinationCIDR, nextHop string) error {
	return e.RouteRemover(destinationCIDR, nextHop)
}

func (e *MockCloudClient) AddRoute(destinationCIDR, instanceID string) error {
	return e.RouteAdder(destinationCIDR, instanceID)
}

func (e *MockCloudClient) GetVPCCIDRs() []string {
	return e.VPCCIDRs
}

func (e *MockCloudClient) AddInstances(insts ...CloudInstance) {
	for _, inst := range insts {
		e.Instances[inst.ID] = inst
	}
}

func (m *MockCloudClient) GetAttributes() CloudAttributes {
	return CloudAttributes{
		DiskProductName: api.StorageGP2,
		FixedSizeVolume: false,
		Provider:        ProviderAWS,
		Region:          "us-east-1",
		Zone:            m.Subnet.AZ,
	}
}

func (m *MockCloudClient) EnsureContainerInstanceCluster() error {
	return m.ContainerClusterEnsurer()
}

func (m *MockCloudClient) ListContainerInstances() ([]ContainerInstance, error) {
	return m.ContainerInstanceLister()
}

func (m *MockCloudClient) ListContainerInstancesFilterID(insts []string) ([]ContainerInstance, error) {
	return m.ContainerInstanceListerFilter(insts)
}

func (m *MockCloudClient) GetContainerInstancesStatuses(instIDs []string) (map[string][]api.UnitStatus, error) {
	return m.ContainerInstancesStatusesGetter(instIDs)
}

func (m *MockCloudClient) StartContainerInstance(pod *api.Pod) (string, error) {
	return m.ContainerInstanceRunner(pod)

}

func (m *MockCloudClient) StopContainerInstance(instID string) error {
	return m.ContainerInstanceStopper(instID)
}

func (m *MockCloudClient) WaitForContainerInstanceRunning(pod *api.Pod) (*api.Pod, error) {
	return m.ContainerInstanceWaiter(pod)
}

func (m *MockCloudClient) AttachSecurityGroups(node *api.Node, groups []string) error {
	return nil
}

func (m *MockCloudClient) AddIAMPermissions(node *api.Node, permissions string) error {
	return nil
}

func (_ *MockCloudClient) SplitBootImageSpec(spec BootImageSpec) []BootImageSpec {
	return []BootImageSpec{spec}
}

func (_ *MockCloudClient) GetArchitecture(_ string) api.Architecture {
	return api.ArchX8664
}

func NewMockClient() *MockCloudClient {
	net := &MockCloudClient{
		Instances:          make(map[string]CloudInstance),
		ContainerInstances: make(map[string]ContainerInstance),
		ControllerID:       "test_cluster",
		InsideVPC:          true,
		VPCCIDRs:           []string{"172.20.0.0/16"},
		Subnet: SubnetAttributes{
			ID:                 "sub-1111",
			CIDR:               "172.16.0.0/16",
			AZ:                 "us-east-1a",
			AddressAffinity:    PublicAddress,
			AvailableAddresses: 250,
		},
	}

	net.InstanceListerFilter = func(iid []string) ([]CloudInstance, error) {
		insts := make([]CloudInstance, 0, len(iid))
		for _, inst := range net.Instances {
			if util.StringInSlice(inst.ID, iid) {
				insts = append(insts, inst)
			}
		}
		return insts, nil
	}

	net.InstanceLister = func() ([]CloudInstance, error) {
		insts := make([]CloudInstance, 0, len(net.Instances))
		for _, inst := range net.Instances {
			insts = append(insts, inst)
		}
		return insts, nil
	}

	net.ContainerInstanceRunner = func(p *api.Pod) (string, error) {
		id := p.Status.BoundInstanceID
		inst := ContainerInstance{ID: id}
		net.ContainerInstances[id] = inst
		return id, nil
	}

	net.ContainerInstanceStopper = func(id string) error {
		if _, exists := net.ContainerInstances[id]; !exists {
			return fmt.Errorf("Container Instance %s does not exist", id)
		}
		delete(net.ContainerInstances, id)
		return nil
	}

	net.ContainerInstanceListerFilter = func(ids []string) ([]ContainerInstance, error) {
		insts := make([]ContainerInstance, 0, len(ids))
		for _, inst := range net.ContainerInstances {
			if util.StringInSlice(inst.ID, ids) {
				insts = append(insts, inst)
			}
		}
		return insts, nil
	}

	net.ContainerInstanceLister = func() ([]ContainerInstance, error) {
		insts := make([]ContainerInstance, 0, len(net.ContainerInstances))
		for _, inst := range net.ContainerInstances {
			insts = append(insts, inst)
		}

		return insts, nil
	}

	net.ContainerInstanceWaiter = func(p *api.Pod) (*api.Pod, error) {
		return p, nil
	}

	net.RouteRemover = func(destinationCIDR, nextHop string) error {
		return nil
	}

	net.RouteAdder = func(destinationCIDR, nextHop string) error {
		return nil
	}

	net.AvailabilityChecker = func() (bool, error) {
		return true, nil
	}

	net.ImageGetter = func(BootImageSpec) (Image, error) {
		t := time.Now()
		img := Image{
			ID:           "1234",
			Name:         "MockImage",
			RootDevice:   "/dev/xvda",
			CreationTime: &t,
		}
		return img, nil
	}

	net.DNSInfoGetter = func() ([]string, []string, error) {
		return []string{"cloud.internal"}, []string{"1.1.1.1"}, nil
	}

	net.Starter = func(node *api.Node, image Image, metadata, iamPermissions string) (string, error) {
		inst := CloudInstance{
			ID:       node.Status.InstanceID,
			NodeName: node.Name,
		}
		net.Instances[node.Status.InstanceID] = inst
		return node.Status.InstanceID, nil
	}

	net.Stopper = func(instID string) error {
		if _, exists := net.Instances[instID]; !exists {
			return fmt.Errorf("Instance %s does not exist", instID)
		}
		delete(net.Instances, instID)
		return nil
	}

	instanceParameters := make(map[string]string)

	net.InstanceParameterAdder = func(instanceID, key, value string, isSecret bool) error {
		instanceParameters[fmt.Sprintf("%s/%s", instanceID, key)] = value
		return nil
	}

	net.InstanceParameterRemover = func(instanceID, key string) error {
		fullKey := fmt.Sprintf("%s/%s", instanceID, key)
		_, ok := instanceParameters[fullKey]
		if !ok {
			return fmt.Errorf("mock InstanceParameterRemove() no such parameter %s", fullKey)
		}
		delete(instanceParameters, fullKey)
		return nil
	}

	return net
}
