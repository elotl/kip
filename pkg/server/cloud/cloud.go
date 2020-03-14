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
	"encoding/json"
	"sort"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"k8s.io/apimachinery/pkg/util/sets"
)

const MilpaAPISGName = "CellSecurityGroup"
const PublicCIDR = "0.0.0.0/0"
const RestAPIPort = 6421

const ProviderAWS = "aws"
const ProviderGCE = "gce"
const ProviderAzure = "azure"

const ControllerTagKey = "KipControllerID"
const NameTagKey = "Name"
const NamespaceTagKey = "KipNamespace"
const NametagTagKey = "KipNametag"
const PodNameTagKey = "KipPodName"

type CloudClient interface {
	SetBootSecurityGroupIDs([]string)
	GetBootSecurityGroupIDs() []string
	StartNode(*api.Node, string) (*StartNodeResult, error)
	StartSpotNode(*api.Node, string) (*StartNodeResult, error)
	// This should always be called from a goroutine as it can take a while
	StopInstance(instanceID string) error
	WaitForRunning(node *api.Node) ([]api.NetworkAddress, error)
	EnsureMilpaSecurityGroups([]string, []string) error
	AttachSecurityGroups(node *api.Node, groups []string) error
	AssignInstanceProfile(node *api.Node, instanceProfile string) error
	ListInstancesFilterID([]string) ([]CloudInstance, error)
	ListInstances() ([]CloudInstance, error)
	ResizeVolume(node *api.Node, size int64) (error, bool)
	GetRegistryAuth() (string, string, error)
	GetImageID(spec BootImageSpec) (string, error)
	SetSustainedCPU(*api.Node, bool) error
	AddInstanceTags(string, map[string]string) error
	ControllerInsideVPC() bool
	ModifySourceDestinationCheck(string, bool) error
	RemoveRoute(string) error
	AddRoute(string, string) error
	GetVPCCIDRs() []string
	// Address spaces used by cloud-internal services that might initiate
	// connections to instances in the VPC.
	CloudStatusKeeper() StatusKeeper
	GetSubnets() ([]SubnetAttributes, error)
	GetAvailabilityZones() ([]string, error)
	GetAttributes() CloudAttributes
}

type CloudAttributes struct {
	DiskProductName           api.StorageType
	FixedSizeVolume           bool
	MaxInstanceSecurityGroups int
	Provider                  string
	Region                    string
}

type StartNodeResult struct {
	InstanceID       string
	AvailabilityZone string
}

type SubnetAddressAffinity string

const (
	PublicAddress  SubnetAddressAffinity = "Public"
	PrivateAddress SubnetAddressAffinity = "Private"
	AnyAddress     SubnetAddressAffinity = "Any"
)

type SubnetAttributes struct {
	Name string
	ID   string
	CIDR string
	AZ   string
	// In AWS subnets we use the subnets private/public address by default
	// flag to decide where to launch public and private nodes.  We store
	// that info in AddressAffinity.  In Azure, it's likely we don't have
	// that type of affinity (there's some interesting NAT options in azure)
	// so we don't really care what subnet our public and private addresses
	// go in.  Also, this is half baked so if you have an idea of a better
	// way to specify placement, knock yourself out.
	AddressAffinity SubnetAddressAffinity
	// In AWS and Azure (pretty sure...), we can get availability
	// stats However, they're harder to come by in GCP.  That said, in
	// GCP you can resize your subnets and we can always query
	// instances and bucket them.
	AvailableAddresses int
	//Capacity            int
}

type Image struct {
	ID           string
	Name         string
	CreationTime *time.Time
}

func SortImagesByCreationTime(images []Image) {
	sort.Slice(images, func(i, j int) bool {
		creationI := images[i].CreationTime
		creationJ := images[j].CreationTime
		if creationI == nil {
			return true
		}
		if creationJ == nil {
			return false
		}
		return creationI.Before(*creationJ)
	})
}

type BootImageSpec map[string]string

func (bis *BootImageSpec) String() string {
	data, err := json.Marshal(bis)
	if err != nil {
		return ""
	}
	return string(data)
}

type CloudInstance struct {
	ID       string
	NodeName string
}

type ContainerInstance struct {
	ID string
}

// List instances only gives us security identifier
type SecurityGroupIdentifier struct {
	ID   string
	Name string
}

type SecurityGroup struct {
	ID           string
	Name         string
	Ports        []InstancePort
	SourceRanges []string
}

func NewSecurityGroup(id, name string, ports []InstancePort, sources []string) SecurityGroup {
	sort.Sort(SortableSliceOfPorts(ports))
	sort.Strings(sources)
	return SecurityGroup{
		ID:           id,
		Name:         name,
		Ports:        ports,
		SourceRanges: sources,
	}
}

type LoadBalancer struct {
	Type             string
	ServiceName      string
	LoadBalancerName string
	Instances        sets.String
	Ports            []InstancePort
	SecurityGroupID  string
	Internal         bool
	Annotations      map[string]string
	DNSName          string
	IPAddress        string
}
