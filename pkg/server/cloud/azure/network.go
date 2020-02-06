package azure

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/sets"
	"k8s.io/klog"
)

type VirtualNetworkAttributes struct {
	Name          string
	ID            string
	ResourceGroup string
	CIDRs         []string
}

func toVNetAttrs(azVNet *network.VirtualNetwork) (VirtualNetworkAttributes, error) {
	vnet := VirtualNetworkAttributes{}
	details, err := azure.ParseResourceID(to.String(azVNet.ID))
	if err != nil {
		return vnet, util.WrapError(err, "Error parsing virtual network information from API")
	}
	vnet.Name = to.String(azVNet.Name)
	vnet.ID = to.String(azVNet.ID)
	vnet.ResourceGroup = details.ResourceGroup
	if azVNet.VirtualNetworkPropertiesFormat == nil ||
		azVNet.AddressSpace == nil ||
		azVNet.AddressSpace.AddressPrefixes == nil ||
		len(*azVNet.AddressSpace.AddressPrefixes) == 0 {
		return vnet, fmt.Errorf("Fatal error: Could not get virtual network address space from API Server")
	}
	vnet.CIDRs = (*azVNet.AddressSpace.AddressPrefixes)
	return vnet, nil
}

// Todo, fix this to return an error, we NEED a resource group here
func splitVNetName(vNetName string) (string, string) {
	resourceGroup := ""
	if strings.Contains(vNetName, "/") {
		parts := strings.SplitN(vNetName, "/", 2)
		resourceGroup = parts[0]
		vNetName = parts[1]
	}
	return resourceGroup, vNetName
}

func (az *AzureClient) getVNet(fullVNetName string) (VirtualNetworkAttributes, error) {
	resourceGroup, vNetName := splitVNetName(fullVNetName)
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	azVNet, err := az.vnets.Get(timeoutCtx, resourceGroup, vNetName, "")
	var vNetAttrs VirtualNetworkAttributes
	if err != nil {
		return vNetAttrs, util.WrapError(err, "Error getting virtual network %s", fullVNetName)
	}
	vNetAttrs, err = toVNetAttrs(&azVNet)
	if err != nil {
		return vNetAttrs, err
	}
	return vNetAttrs, nil
}

// If the user supplied a vNetName, use that vNet
// -- query it, see if it exists, get the resource group
// otherwise use the local VNet
func (az *AzureClient) setupClusterVNet(vNetName, subnetName string) error {
	if vNetName != "" {
		if subnetName == "" {
			return fmt.Errorf("Error setting up azure networking: a subnet name (cloud.azure.subnetName) must be supplied in server.yml if a virtual network name is specified in server.yml")
		}
		vnet, err := az.getVNet(vNetName)
		if err != nil {
			return err
		}
		az.virtualNetwork = vnet
		subnet, err := az.getSubnet(vnet.ResourceGroup, vnet.Name, subnetName)
		if err != nil {
			return err
		}
		az.subnet = subnet
	} else {
		vnet, subnet, err := az.getLocalInstanceNetwork()
		if err != nil {
			return err
		}
		az.virtualNetwork = vnet
		az.subnet = subnet
	}
	return nil
}

// subnetIDs look like:
// /subscriptions/4e84e89a-b806-4d7d-900b-cae8cb640215/resourceGroups/milpa-bcox/providers/Microsoft.Network/virtualNetworks/milpa-bcox/subnets/milpa-bcox-default
func subnetToVirtualNetwork(resourceID string) string {
	const resourceIDPatternText = `resourceGroups/([^/]+)/.*/virtualNetworks/([^/]+)/subnets`
	resourceIDPattern := regexp.MustCompile(resourceIDPatternText)
	match := resourceIDPattern.FindStringSubmatch(resourceID)
	if len(match) == 3 {
		return match[1] + "/" + match[2]
	}
	return ""
}

func resourceIDToResourceName(resourceID string) string {
	parts := strings.Split(resourceID, "/")
	return parts[len(parts)-1]
}

// Well this turned into an ugly hunk of poop... Thanks Azure!
// VNetNames look like <resourceGroup>/<vNetName>
// SubnetNames look like <subnet_name> without the resource group
func (az *AzureClient) GetVMNetworks(vmResourceGroup, vmName string) (vNets, subnetNames []string) {
	vNets = []string{}
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	vm, err := az.vms.Get(timeoutCtx, vmResourceGroup, vmName, "")
	if err != nil {
		klog.V(2).Infof("Could not find controller VM %s/%s in subscription",
			vmResourceGroup, vmName)
		return
	}
	if vm.VirtualMachineProperties == nil ||
		vm.NetworkProfile == nil ||
		vm.NetworkProfile.NetworkInterfaces == nil {
		return
	}
	ifaces := *vm.NetworkProfile.NetworkInterfaces
	for i := range ifaces {
		nicID := to.String(ifaces[i].ID)
		if nicID == "" {
			continue
		}
		details, err := azure.ParseResourceID(nicID)
		if err != nil {
			klog.Errorln("Error parsing resource ID for controller NIC", err)
			continue
		}
		nicResourceGroup := details.ResourceGroup
		nicName := details.ResourceName
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		nic, err := az.nics.Get(timeoutCtx, nicResourceGroup, nicName, "")
		if err != nil {
			klog.V(2).Infof("Could not find controller NIC %s/%s in subscription",
				vmResourceGroup, vmName)
			continue
		}
		if nic.InterfacePropertiesFormat == nil ||
			nic.IPConfigurations == nil {
			continue
		}
		ipconfig := *nic.IPConfigurations
		for j := range ipconfig {
			if ipconfig[j].InterfaceIPConfigurationPropertiesFormat == nil {
				continue
			}
			if ipconfig[j].Subnet != nil {
				snid := to.String(ipconfig[j].Subnet.ID)
				snName := resourceIDToResourceName(snid)
				subnetNames = append(subnetNames, snName)
				vmVNet := subnetToVirtualNetwork(snid)
				vNets = append(vNets, vmVNet)
			}
		}
	}
	return
}

func (az *AzureClient) getLocalInstanceNetwork() (VirtualNetworkAttributes, cloud.SubnetAttributes, error) {
	var (
		vNet   VirtualNetworkAttributes
		subnet cloud.SubnetAttributes
	)
	vmResourceGroup, vmName := getMetadataInstanceName()
	if vmResourceGroup == "" || vmName == "" {
		return vNet, subnet, fmt.Errorf("could not connect to instance metadata to determine controller network properties. A virtualNetworkName and subnetName will need to be specified in the cloud.azure section of server.yml")
	}
	vNetNames, subnetNames := az.GetVMNetworks(vmResourceGroup, vmName)
	if len(vNetNames) == 0 {
		return vNet, subnet, fmt.Errorf("could not detect which virtual network the controller is attached to. A virtualNetworkName will need to be specified in server.yml")
	} else if len(vNetNames) > 1 {
		return vNet, subnet, fmt.Errorf("Multiple virtual networks are attached to this instance and it is impossible to tell which network nodes should be launched into. A virtualNetworkName will need to be specified in the cloud.azure section server.yml")
	}
	klog.V(2).Infof("local machine is connected to virtual network %s", vNetNames[0])
	if len(subnetNames) == 0 {
		return vNet, subnet, fmt.Errorf("could not detect which subnet the controller is attached to. A subnetName will need to be specified in server.yml")
	} else if len(subnetNames) > 1 {
		return vNet, subnet, fmt.Errorf("Multiple subnets are attached to this instance and it is impossible to tell which subnet nodes should be launched into. A subnetName will need to be specified in the cloud.azure section server.yml")
	}
	klog.V(2).Infof("local machine is connected to subnet %s", subnetNames[0])

	vNet, err := az.getVNet(vNetNames[0])
	if err != nil {
		return vNet, subnet, util.WrapError(err, "Error looking up local machine's vNet %s. Please specify a virtualNetworkName and subnetName in server.yml", vNetNames[0])
	}
	subnet, err = az.getSubnet(vNet.ResourceGroup, vNet.Name, subnetNames[0])
	if err != nil {
		return vNet, subnet, util.WrapError(err, "Error looking up local machine's subnet %s. Please specify a virtualNetworkName and subnetName in server.yml", subnetNames[0])
	}

	return vNet, subnet, nil
}

func (az *AzureClient) ControllerInsideVPC() bool {
	_, _, err := az.getLocalInstanceNetwork()
	return err == nil
}

func (az *AzureClient) ModifySourceDestinationCheck(instanceID string, isEnabled bool) error {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	nic, err := az.nics.Get(timeoutCtx, instanceID, instanceID, "")
	if err != nil {
		return util.WrapError(err, "looking up NIC for VM src/dst check")
	}
	nic.EnableIPForwarding = to.BoolPtr(isEnabled)
	ctx = context.Background()
	timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	future, err := az.nics.CreateOrUpdate(
		timeoutCtx, instanceID, instanceID, nic)
	if err != nil {
		return fmt.Errorf("updating NIC for VM %q: %v", instanceID, err)
	}
	timeoutCtx, cancel = context.WithTimeout(ctx, azureWaitTimeout)
	defer cancel()
	err = future.WaitForCompletionRef(timeoutCtx, az.nics.Client)
	if err != nil {
		return fmt.Errorf("getting NIC update response for VM %s: %v",
			instanceID, err)
	}
	_, err = future.Result(az.nics)
	if err != nil {
		return err
	}
	klog.V(2).Infof("enabled src/dst check on %q", instanceID)
	return nil
}

func (az *AzureClient) RemoveRoute(destinationCIDR string) error {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	vnet, err := az.vnets.Get(
		timeoutCtx, az.virtualNetwork.ResourceGroup, az.virtualNetwork.Name, "")
	if err != nil {
		return util.WrapError(err, "getting virtual network")
	}
	if vnet.VirtualNetworkPropertiesFormat == nil ||
		vnet.VirtualNetworkPropertiesFormat.Subnets == nil {
		return fmt.Errorf("error getting subnets removing %q", destinationCIDR)
	}
	for _, subnet := range *vnet.VirtualNetworkPropertiesFormat.Subnets {
		if subnet.SubnetPropertiesFormat == nil ||
			subnet.SubnetPropertiesFormat.RouteTable == nil {
			return fmt.Errorf("error getting rt removing %q", destinationCIDR)
		}
		details, err := azure.ParseResourceID(
			to.String(subnet.SubnetPropertiesFormat.RouteTable.ID))
		if err != nil {
			return util.WrapError(err, "parsing route table ID")
		}
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		rt, err := az.routetables.Get(
			timeoutCtx, details.ResourceGroup, details.ResourceName, "")
		if err != nil {
			return util.WrapError(err, "getting route table")
		}
		rtName := to.String(rt.Name)
		for _, route := range *rt.Routes {
			if route.RoutePropertiesFormat == nil {
				continue
			}
			cidr := to.String(route.RoutePropertiesFormat.AddressPrefix)
			if cidr != destinationCIDR {
				continue
			}
			details, err = azure.ParseResourceID(to.String(route.ID))
			if err != nil {
				continue
			}
			timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
			defer cancel()
			future, err := az.routes.Delete(
				timeoutCtx,
				details.ResourceGroup,
				rtName,
				details.ResourceName)
			if err != nil {
				return util.WrapError(err, "removing route")
			}
			timeoutCtx, cancel = context.WithTimeout(ctx, azureWaitTimeout)
			defer cancel()
			err = future.WaitForCompletionRef(timeoutCtx, az.routes.Client)
			if err != nil {
				return fmt.Errorf("getting RT remove response for %q: %v",
					destinationCIDR, err)
			}
			_, err = future.Result(az.routes)
			if err != nil {
				return err
			}
			klog.V(2).Infof("removed route for %q", destinationCIDR)
		}
	}
	return nil
}

func (az *AzureClient) AddRoute(destinationCIDR, instanceID string) error {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	nic, err := az.nics.Get(timeoutCtx, instanceID, instanceID, "")
	if err != nil {
		return util.WrapError(err, "looking up NIC for adding route")
	}
	ipconfig, err := getMilpaIPConfiguration(nic)
	if err != nil {
		return util.WrapError(err, "getting IP config for adding route")
	}
	details, err := azure.ParseResourceID(to.String(ipconfig.Subnet.ID))
	if err != nil {
		return util.WrapError(err, "parsing IP config for adding route")
	}
	timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	subnet, err := az.subnets.Get(
		timeoutCtx,
		details.ResourceGroup,
		az.virtualNetwork.Name,
		details.ResourceName,
		"")
	if err != nil {
		return util.WrapError(err, "getting subnet for adding route")
	}
	details, err = azure.ParseResourceID(to.String(subnet.RouteTable.ID))
	if err != nil {
		return util.WrapError(err, "parsing route table ID")
	}
	route := network.Route{
		Name: to.StringPtr(instanceID),
		RoutePropertiesFormat: &network.RoutePropertiesFormat{
			AddressPrefix:    to.StringPtr(destinationCIDR),
			NextHopType:      network.RouteNextHopTypeVirtualAppliance,
			NextHopIPAddress: ipconfig.PrivateIPAddress,
		},
	}
	timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	_, err = az.routes.CreateOrUpdate(
		timeoutCtx,
		details.ResourceGroup,
		details.ResourceName,
		to.String(route.Name),
		route)
	if err != nil {
		return util.WrapError(err, "adding route")
	}
	klog.V(2).Infof("created route %q via VM %q", destinationCIDR, instanceID)
	return nil
}

// From the docs: Azure reserves some IP addresses within each
// subnet. The first and last IP addresses of each subnet are reserved
// for protocol conformance, along with the x.x.x.1-x.x.x.3 addresses
// of each subnet, which are used for Azure services.
func addressCount(ipnet *net.IPNet) int {
	prefixLen, bits := ipnet.Mask.Size()
	numAddresses := 1 << (uint64(bits) - uint64(prefixLen))
	octets := int(numAddresses / 256)
	unavailable := octets*3 + 2
	return numAddresses - unavailable
}

func (az *AzureClient) GetSubnets() ([]cloud.SubnetAttributes, error) {
	return []cloud.SubnetAttributes{az.subnet}, nil
}

func (az *AzureClient) getSubnet(resourceGroup, virtualNetworkName, subnetName string) (cloud.SubnetAttributes, error) {
	var subnet cloud.SubnetAttributes
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	azs, err := az.subnets.Get(timeoutCtx, resourceGroup, virtualNetworkName, subnetName, "")
	if err != nil {
		return subnet, err
	}
	subnet.Name = to.String(azs.Name)
	subnet.ID = to.String(azs.ID)
	subnet.CIDR = to.String(azs.AddressPrefix)
	subnet.AddressAffinity = cloud.AnyAddress
	subnet.AZ = ""
	return subnet, nil
}

var locationsWithAZSupport = []string{
	"centralus",
	"eastus2",
	"francecentral",
	"northeurope",
	"southeastasia",
	"westeurope",
	"westus2",
}

func (az *AzureClient) GetAvailabilityZones() ([]string, error) {
	// Azure returns bad data for some locations. See:
	// https://github.com/Azure/azure-rest-api-specs/issues/5163
	// While that's open, we'll hard-code the locations that support
	// AZs
	if !util.StringInSlice(az.locationName(), locationsWithAZSupport) {
		return []string{}, nil
	}

	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	resultPage, err := az.skus.List(timeoutCtx)
	if err != nil {
		return nil, err
	}
	zones := sets.NewString()
	ourLocation := az.locationName()
	for resultPage.NotDone() {
		skus := resultPage.Values()
		for _, sku := range skus {
			if to.String(sku.ResourceType) != "virtualMachines" {
				continue
			}
			if sku.Locations == nil || len(*sku.Locations) == 0 || (*sku.Locations)[0] != ourLocation {
				continue
			}
			if sku.LocationInfo != nil {
				for _, li := range *sku.LocationInfo {
					if li.Zones != nil {
						zones.Insert(*li.Zones...)
					}
				}
			}
		}
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		err = resultPage.NextWithContext(timeoutCtx)
		if err != nil {
			return nil, err
		}
	}
	return zones.List(), nil
}

func (az *AzureClient) getPublicIP(name string) (string, error) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	ip, err := az.ips.Get(timeoutCtx, name, name, "")
	if err != nil {
		return "", err
	}
	prop := ip.PublicIPAddressPropertiesFormat
	if prop == nil {
		return "", fmt.Errorf("No ip address found")
	}
	return to.String(prop.IPAddress), nil
}

func getMilpaIPConfiguration(iface network.Interface) (*network.InterfaceIPConfiguration, error) {
	if iface.InterfacePropertiesFormat == nil || iface.InterfacePropertiesFormat.IPConfigurations == nil {
		return nil, fmt.Errorf("Could not parse nic %s info. Azure sent back a garbage reply", to.String(iface.Name))
	}
	for i := range *iface.InterfacePropertiesFormat.IPConfigurations {
		if to.String((*iface.IPConfigurations)[i].Name) == milpaIPConfig {
			return &((*iface.IPConfigurations)[i]), nil
		}
	}
	return nil, fmt.Errorf("Could not find milpa IP configuration")
}
