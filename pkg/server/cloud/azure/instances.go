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

package azure

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	glob "github.com/ryanuber/go-glob"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
)

var (
	milpaIPConfig    = "milpaIPConfig"
	milpaPodIPConfig = "milpaPodIPConfig"
)

func (az *AzureClient) StopInstance(instanceID string) error {
	// Deleting the instance's resource group should delete the VM and
	// everything else associated with the VM (NIC, Disks, Public IP,
	// etc.)
	return az.DeleteResourceGroup(instanceID)
}

func (az *AzureClient) createIPAddress(instanceID string, zone string) (*network.PublicIPAddress, error) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	future, err := az.ips.CreateOrUpdate(timeoutCtx, instanceID, instanceID,
		network.PublicIPAddress{
			Sku: &network.PublicIPAddressSku{
				Name: network.PublicIPAddressSkuNameStandard,
			},
			Name:     to.StringPtr(instanceID),
			Location: to.StringPtr(az.region),
			Zones:    makeZoneParam(zone),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAddressVersion:   network.IPv4,
				PublicIPAllocationMethod: network.Static,
				DNSSettings: &network.PublicIPAddressDNSSettings{
					DomainNameLabel: to.StringPtr(instanceID),
				},
			},
		},
	)

	if err != nil {
		return nil, util.WrapError(err, "cannot create public ip address for instance %s", instanceID)
	}

	timeoutCtx, cancel = context.WithTimeout(ctx, azureIPAddressWaitTimeout)
	defer cancel()
	err = future.WaitForCompletionRef(timeoutCtx, az.ips.Client)
	if err != nil {
		return nil, util.WrapError(err, "Error waiting for IP address creation for instance %s", instanceID)
	}
	ip, err := future.Result(az.ips)
	if err != nil {
		return nil, err
	}
	return &ip, nil
}

func (az *AzureClient) createNIC(instanceID string, ipID string) (string, error) {
	nicName := instanceID
	nicParams := network.Interface{
		Name:     to.StringPtr(nicName),
		Location: to.StringPtr(az.region),
		InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
			NetworkSecurityGroup: &network.SecurityGroup{
				ID: to.StringPtr(az.nsgID),
			},
			// Todo: look into accelerated networking.
			// https://azure.microsoft.com/en-us/blog/maximize-your-vm-s-performance-with-accelerated-networking-now-generally-available-for-both-windows-and-linux/
			// EnableAcceleratedNetworking: to.BoolPtr(true),
			IPConfigurations: &[]network.InterfaceIPConfiguration{
				{
					// This name is used to find the IP configuration later on
					Name: to.StringPtr(milpaIPConfig),
					InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
						Subnet: &network.Subnet{
							Name: to.StringPtr(az.subnet.Name),
							ID:   to.StringPtr(az.subnet.ID),
						},
						PrivateIPAllocationMethod: network.Dynamic,
						PrivateIPAddressVersion:   network.IPv4,
						PublicIPAddress: &network.PublicIPAddress{
							ID: to.StringPtr(ipID),
						},
						Primary: to.BoolPtr(true),
					},
				},
				{
					// Pod IP address.
					Name: to.StringPtr(milpaPodIPConfig),
					InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
						Subnet: &network.Subnet{
							Name: to.StringPtr(az.subnet.Name),
							ID:   to.StringPtr(az.subnet.ID),
						},
						PrivateIPAllocationMethod: network.Dynamic,
						PrivateIPAddressVersion:   network.IPv4,
						Primary:                   to.BoolPtr(false),
					},
				},
			},
		},
	}
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	future, err := az.nics.CreateOrUpdate(timeoutCtx, instanceID, nicName, nicParams)
	if err != nil {
		return "", fmt.Errorf("cannot create nic: %v", err)
	}
	timeoutCtx, cancel = context.WithTimeout(ctx, azureWaitTimeout)
	defer cancel()
	err = future.WaitForCompletionRef(timeoutCtx, az.nics.Client)
	if err != nil {
		return "", fmt.Errorf("cannot get nic create or update future response: %v", err)
	}
	nic, err := future.Result(az.nics)
	if err != nil {
		return "", err
	}
	return to.String(nic.ID), nil
}

func (az *AzureClient) StartNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	klog.V(2).Infof("Starting instance for node: %v", node)
	instanceID := makeInstanceID(az.controllerID, node.Name)
	err := az.createResourceGroup(instanceID)
	if err != nil {
		return nil, util.WrapError(err, "Error creating new Azure Resource Group for Node %s", node.Name)
	}
	cleanup := func() {
		err := az.DeleteResourceGroup(instanceID)
		if err != nil {
			klog.Errorln(
				"Error deleting azure resource group after start failure",
				err,
			)
		}
	}

	zone := node.Spec.Placement.AvailabilityZone
	ipID := ""
	if !node.Spec.Resources.PrivateIPOnly {
		ip, err := az.createIPAddress(instanceID, zone)
		if err != nil {
			cleanup()
			// todo, see if our subnet is constrained, return that error instead
			return nil, util.WrapError(err, "Error creating new IP for Node %s", node.Name)
		}
		ipID = to.String(ip.ID)
	}
	nicID, err := az.createNIC(instanceID, ipID)
	if err != nil {
		// Todo, see if our subnet is constrained on private addresses
		cleanup()
		return nil, util.WrapError(err, "Error creating new NIC for Node %s", node.Name)
	}
	metadataptr := to.StringPtr(metadata)
	// Azure doesn't like emptystring metadata
	if len(metadata) == 0 {
		metadataptr = nil
	}

	volSizeGiB := cloud.ToSaneVolumeSize(node.Spec.Resources.VolumeSize)
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	future, err := az.vms.CreateOrUpdate(timeoutCtx, instanceID, instanceID,
		compute.VirtualMachine{
			Location: to.StringPtr(az.region),
			Zones:    makeZoneParam(zone),
			VirtualMachineProperties: &compute.VirtualMachineProperties{
				HardwareProfile: &compute.HardwareProfile{
					VMSize: compute.VirtualMachineSizeTypes(node.Spec.InstanceType),
				},

				StorageProfile: &compute.StorageProfile{
					ImageReference: &compute.ImageReference{
						ID: to.StringPtr(node.Spec.BootImage),
					},
					OsDisk: &compute.OSDisk{
						OsType:       compute.Linux,
						Name:         to.StringPtr(instanceID),
						CreateOption: compute.DiskCreateOptionTypesFromImage,
						DiskSizeGB:   to.Int32Ptr(volSizeGiB),
					},
				},

				OsProfile: &compute.OSProfile{
					ComputerName:  to.StringPtr(instanceID),
					AdminUsername: to.StringPtr("milpa"),
					// We can't _not_ include this so we'll add it here
					// but it wont be used by our image
					AdminPassword: to.StringPtr("thisisunused-AFdsj483.fd8r37r"),
					LinuxConfiguration: &compute.LinuxConfiguration{
						ProvisionVMAgent: to.BoolPtr(false),
					},
					CustomData: metadataptr,
				},
				NetworkProfile: &compute.NetworkProfile{
					NetworkInterfaces: &[]compute.NetworkInterfaceReference{
						{
							ID: to.StringPtr(nicID),
							NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						},
					},
				},
			},
			Tags: az.getNodeTags(node),
		},
	)
	if err != nil {
		cleanup()
		return nil, getStartVMError(err, az.subnet.ID, node.Spec.Placement.AvailabilityZone)
	}
	timeoutCtx, cancel = context.WithTimeout(ctx, azureStartNodeTimeout)
	defer cancel()
	err = future.WaitForCompletionRef(timeoutCtx, az.vms.Client)
	if err != nil {
		cleanup()
		return nil, util.WrapError(err, "Error waiting for VM %s to boot", instanceID)
	}
	_, err = future.Result(az.vms)
	if err != nil {
		cleanup()
		return nil, util.WrapError(err, "Error getting VM %s after booting", instanceID)
	}
	startResult := &cloud.StartNodeResult{
		InstanceID:       instanceID,
		AvailabilityZone: zone,
	}
	return startResult, nil
}

func (az *AzureClient) StartSpotNode(node *api.Node, metadata string) (*cloud.StartNodeResult, error) {
	return az.StartNode(node, metadata)
}

func (az *AzureClient) getNodeTags(node *api.Node) map[string]*string {
	nametag := util.CreateUnboundNodeNameTag(az.nametag)
	tags := map[string]*string{
		"Name":                 to.StringPtr(nametag),
		"Node":                 to.StringPtr(node.Name),
		cloud.ControllerTagKey: to.StringPtr(az.controllerID),
		cloud.NametagTagKey:    to.StringPtr(az.nametag),
	}
	return tags
}

func (az *AzureClient) WaitForRunning(node *api.Node) ([]api.NetworkAddress, error) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	nic, err := az.nics.Get(timeoutCtx, node.Status.InstanceID, node.Status.InstanceID, "")
	if err != nil {
		return nil, util.WrapError(err, "Error getting recently started VM %s from Azure", node.Status.InstanceID)
	}

	addresses := []api.NetworkAddress{}
	if nic.IPConfigurations != nil {
		for _, ipconfig := range *nic.IPConfigurations {
			if to.String(ipconfig.Name) == milpaIPConfig {
				ipProperties := ipconfig.InterfaceIPConfigurationPropertiesFormat
				if ipProperties == nil {
					return addresses, fmt.Errorf("invalid response from Azure when getting %s interface parameters", node.Status.InstanceID)
				}
				ip := to.String(ipconfig.PrivateIPAddress)
				addresses = api.NewNetworkAddresses(ip, "")
				if ipconfig.PublicIPAddress != nil && ipconfig.PublicIPAddress.ID != nil {
					ip, err := az.getPublicIP(node.Status.InstanceID)
					if err == nil && ip != "" {
						addresses = api.SetPublicAddresses(ip, "", addresses)
					}
				}
			}
			if to.String(ipconfig.Name) == milpaPodIPConfig {
				ipProperties := ipconfig.InterfaceIPConfigurationPropertiesFormat
				if ipProperties == nil {
					return addresses, fmt.Errorf("invalid response from Azure when getting %s pod IP address", node.Status.InstanceID)
				}
				addresses = api.SetPodIP(to.String(ipconfig.PrivateIPAddress), addresses)
			}
		}
	}
	return addresses, nil
}

func (az *AzureClient) SetSustainedCPU(node *api.Node, enabled bool) error {
	klog.V(2).Infoln("Setting sustained CPU in Azure has no impact")
	return nil
}

// This is a bit of an overkill right now, since on Azure we only filter image
// list results via matching the name of the image to a glob, e.g.
// elotl-kipdev-*.
func matchSpec(properties map[string]string, spec cloud.BootImageSpec) bool {
	if len(spec) < 1 {
		// No spec.
		return true
	}
	if len(properties) < len(spec) {
		// The spec contains more properties.
		return false
	}
	for name, value := range spec {
		if propValue, ok := properties[name]; ok {
			if glob.Glob(value, propValue) {
				continue
			}
		}
		return false
	}
	return true
}

func (az *AzureClient) GetImageID(spec cloud.BootImageSpec) (string, error) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	rgName := regionalResourceGroupName(az.region)
	resultPage, err := az.images.ListByResourceGroup(timeoutCtx, rgName)
	if err != nil {
		return "", err
	}
	images := make(map[string]compute.Image)
	imageNames := make([]string, 0)
	locationName := az.locationName()
	for resultPage.NotDone() {
		azImages := resultPage.Values()
		for _, azImage := range azImages {
			if to.String(azImage.Location) != locationName {
				continue
			}
			// We have to filter results ourselves, unlike on AWS. For now,
			// filtering based on name is supported (and we filter
			// unconditionally on location above).
			name := to.String(azImage.Name)
			properties := map[string]string{
				"name": name,
			}
			if !matchSpec(properties, spec) {
				continue
			}
			images[name] = azImage
			imageNames = append(imageNames, name)
		}
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		err := resultPage.NextWithContext(timeoutCtx)
		if err != nil {
			return "", err
		}
	}
	if len(images) == 0 {
		msg := fmt.Sprintf("no images found for spec %+v", spec)
		klog.Errorf("%s", msg)
		return "", fmt.Errorf("%s", msg)
	}
	// compute.Image has no creation timestamp, so we rely on naming convention
	// here: the name has a timestamp in it, so we can get the latest one via
	// lexicographical sorting.
	sort.Strings(imageNames)
	latestImage := imageNames[len(imageNames)-1]
	return to.String(images[latestImage].ID), nil
}

func (az *AzureClient) ListInstancesFilterID(ids []string) ([]cloud.CloudInstance, error) {
	idsSet := sets.NewString(ids...)
	return az.listInstancesHelper(func(inst compute.VirtualMachine) bool {
		return idsSet.Has(to.String(inst.Name))
	})
}

func (az *AzureClient) ListInstances() (insts []cloud.CloudInstance, err error) {
	return az.listInstancesHelper(func(inst compute.VirtualMachine) bool {
		tags := inst.Tags
		return tags != nil &&
			to.String(tags[cloud.ControllerTagKey]) == az.controllerID
	})
}

func (az *AzureClient) listInstancesHelper(filter func(inst compute.VirtualMachine) bool) (insts []cloud.CloudInstance, err error) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	resultsPage, err := az.vms.ListAll(timeoutCtx)
	if err != nil {
		return
	}
	insts = make([]cloud.CloudInstance, 0, 100) // Randomly chosen size...
	instNames := sets.NewString()
	for resultsPage.NotDone() {
		instances := resultsPage.Values()
		for _, inst := range instances {
			instName := to.String(inst.Name)
			if filter(inst) {
				instNames.Insert(instName)
				insts = append(insts, cloud.CloudInstance{
					ID:       instName,
					NodeName: to.String(inst.Tags["Node"]),
				})
			}
		}
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		err = resultsPage.NextWithContext(timeoutCtx)
		if err != nil {
			return
		}
	}

	return insts, nil
}

func getSecurityGroupsFromInterface(iface network.Interface) []cloud.SecurityGroupIdentifier {
	ipConfig, err := getMilpaIPConfiguration(iface)
	if err != nil || ipConfig.ApplicationSecurityGroups == nil {
		return []cloud.SecurityGroupIdentifier{}
	}

	asgs := *ipConfig.ApplicationSecurityGroups
	groups := make([]cloud.SecurityGroupIdentifier, len(asgs))

	for i := range asgs {
		details, err := azure.ParseResourceID(to.String(asgs[i].ID))
		if err != nil {
			continue
		}
		groups[i] = cloud.SecurityGroupIdentifier{
			ID:   details.ResourceName,
			Name: details.ResourceName,
		}
	}
	return groups
}

func (az *AzureClient) AddInstanceTags(iid string, labels map[string]string) error {
	newTags, err := filterLabelsForTags(iid, labels)
	if err != nil {
		klog.Warning(err)
	}
	if len(newTags) > 0 {
		ctx := context.Background()
		timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		vm, err := az.vms.Get(timeoutCtx, iid, iid, "")
		if err != nil {
			return util.WrapError(err, "Error looking up VM %s in azure", iid)
		}
		tags := vm.Tags
		for k, v := range newTags {
			tags[k] = v
		}
		vm.Tags = tags
		vm.VirtualMachineProperties = nil
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		future, err := az.vms.CreateOrUpdate(timeoutCtx, iid, iid, vm)
		if err != nil {
			return util.WrapError(err, "Error adding tags to vm %s", iid)
		}
		timeoutCtx, cancel = context.WithTimeout(ctx, azureWaitTimeout)
		defer cancel()
		err = future.WaitForCompletionRef(timeoutCtx, az.vms.Client)
		if err != nil {
			return util.WrapError(
				err, "Error waiting for VM %s", iid)
		}
		_, err = future.Result(az.vms)
		if err != nil {
			return util.WrapError(
				err, "Error getting VM %s after adding tags", iid)
		}
	}
	return nil
}

// Todo: need to figure out what to do with this
func (az *AzureClient) ResizeVolume(node *api.Node, size int64) (error, bool) {
	return nil, true
}

func getStartVMError(err error, subnetID, az string) error {
	if isSubnetConstrainedError(err) {
		return &cloud.NoCapacityError{
			OriginalError: err.Error(),
			SubnetID:      subnetID,
		}
	} else if isAZConstrainedError(err, az) {
		return &cloud.NoCapacityError{
			OriginalError: err.Error(),
			AZ:            az,
		}
	} else if isInstanceConstrainedError(err) {
		return &cloud.NoCapacityError{
			OriginalError: err.Error(),
		}
	} else if isUnsupportedInstanceError(err) {
		return &cloud.UnsupportedInstanceError{err.Error()}
	}
	return util.WrapError(err, "Could not run instance")
}

// Todo, need to see what we get back when we have no more addresses
func isSubnetConstrainedError(err error) bool {
	return false
}

func isAZConstrainedError(err error, az string) bool {
	if strings.Contains(err.Error(), "ZonalAllocationFailed") ||
		// Hack: this should impact all AZs but not an empty AZ...
		strings.Contains(err.Error(), "ResourceTypeNotSupportAvailabilityZones") {
		return true
	} else if az != "" {
		if strings.Contains(err.Error(), "AllocationFailed") ||
			strings.Contains(err.Error(), "SkuNotAvailable") {
			return true
		}
	}
	return false
}

func isInstanceConstrainedError(err error) bool {
	if strings.Contains(err.Error(), "AllocationFailed") ||
		strings.Contains(err.Error(), "quota limit") {
		return true
	}
	return false
}

func isUnsupportedInstanceError(err error) bool {
	return strings.Contains(err.Error(), "SkuNotAvailable")
}

func (az *AzureClient) AssignInstanceProfile(node *api.Node, instanceProfile string) error {
	return nil
}
