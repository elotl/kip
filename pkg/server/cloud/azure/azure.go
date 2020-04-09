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
	"time"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2017-06-01/storage"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/apimachinery/pkg/util/errors"
)

var (
	azureDefaultTimeout       = 45 * time.Second
	azureWaitTimeout          = 90 * time.Second
	azureIPAddressWaitTimeout = 150 * time.Second
	azureStartNodeTimeout     = 300 * time.Second
	maxAzureUserTags          = 10
	maxAzureTagKeyLength      = 512
	maxAzureTagValueLength    = 255
	reservedTagCharacters     = "<>%&\\?/"
)

type AzureClient struct {
	asgs           network.ApplicationSecurityGroupsClient
	disks          compute.DisksClient
	groups         resources.GroupsClient
	images         compute.ImagesClient
	ips            network.PublicIPAddressesClient
	lbs            network.LoadBalancersClient
	nics           network.InterfacesClient
	nsgs           network.SecurityGroupsClient
	routes         network.RoutesClient
	routetables    network.RouteTablesClient
	rules          network.SecurityRulesClient
	skus           compute.ResourceSkusClient
	storage        storage.AccountsClient
	subnets        network.SubnetsClient
	vms            compute.VirtualMachinesClient
	vnets          network.VirtualNetworksClient
	controllerID   string
	nametag        string
	subscriptionID string
	region         string
	virtualNetwork VirtualNetworkAttributes
	subnet         cloud.SubnetAttributes
	nsgID          string
	nsgName        string
	bootASGNames   []string // Names, not the full azure IDs
	cloudStatus    cloud.StatusKeeper
}

func getAzureConnection(subscriptionID string) (*AzureClient, error) {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return nil, err
	}
	az := &AzureClient{
		asgs:        network.NewApplicationSecurityGroupsClient(subscriptionID),
		disks:       compute.NewDisksClient(subscriptionID),
		groups:      resources.NewGroupsClient(subscriptionID),
		images:      compute.NewImagesClient(subscriptionID),
		ips:         network.NewPublicIPAddressesClient(subscriptionID),
		lbs:         network.NewLoadBalancersClient(subscriptionID),
		nics:        network.NewInterfacesClient(subscriptionID),
		nsgs:        network.NewSecurityGroupsClient(subscriptionID),
		routes:      network.NewRoutesClient(subscriptionID),
		routetables: network.NewRouteTablesClient(subscriptionID),
		rules:       network.NewSecurityRulesClient(subscriptionID),
		skus:        compute.NewResourceSkusClient(subscriptionID),
		storage:     storage.NewAccountsClient(subscriptionID),
		subnets:     network.NewSubnetsClient(subscriptionID),
		vms:         compute.NewVirtualMachinesClient(subscriptionID),
		vnets:       network.NewVirtualNetworksClient(subscriptionID),
	}
	az.asgs.Authorizer = authorizer
	az.disks.Authorizer = authorizer
	az.groups.Authorizer = authorizer
	az.images.Authorizer = authorizer
	az.ips.Authorizer = authorizer
	az.lbs.Authorizer = authorizer
	az.nics.Authorizer = authorizer
	az.nsgs.Authorizer = authorizer
	az.routes.Authorizer = authorizer
	az.routetables.Authorizer = authorizer
	az.rules.Authorizer = authorizer
	az.skus.Authorizer = authorizer
	az.storage.Authorizer = authorizer
	az.subnets.Authorizer = authorizer
	az.vms.Authorizer = authorizer
	az.vnets.Authorizer = authorizer

	return az, nil
}

func NewAzureClient(controllerID, nametag, subscriptionID, region, vNetName, subnetName string) (*AzureClient, error) {
	az, err := getAzureConnection(subscriptionID)
	if err != nil {
		return nil, util.WrapError(err, "Could not create Azure API client")
	}
	az.controllerID = controllerID
	az.nametag = nametag
	az.region = region
	az.subscriptionID = subscriptionID

	err = az.ensureResourceGroups()
	if err != nil {
		return az, util.WrapError(err, "Error setting up cluster resource group")
	}

	// if the user specified a vNet, use that, otherwise create the vNet
	err = az.setupClusterVNet(vNetName, subnetName)
	if err != nil {
		return az, util.WrapError(err, "Error setting up cluster virtual network")
	}
	az.cloudStatus, err = cloud.NewAZSubnetStatus(az)
	if err != nil {
		return az, util.WrapError(err, "Error creating azure cloud status keeper")
	}
	return az, err
}

func CheckConnection(subscriptionID string) error {
	az, err := getAzureConnection(subscriptionID)
	if err != nil {
		return util.WrapError(err, "Error testing Azure API connection")
	}
	top := int32(1)
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	_, err = az.groups.List(timeoutCtx, "", &top)
	if err != nil {
		return fmt.Errorf("Could not access azure API")
	}
	return nil
}

func (az *AzureClient) CloudStatusKeeper() cloud.StatusKeeper {
	return az.cloudStatus
}

func (az *AzureClient) GetVPCCIDRs() []string {
	return az.virtualNetwork.CIDRs
}

func (az *AzureClient) GetAttributes() cloud.CloudAttributes {
	return cloud.CloudAttributes{
		DiskProductName:           api.StorageStandardSSD,
		FixedSizeVolume:           true,
		MaxInstanceSecurityGroups: 1,
		Provider:                  cloud.ProviderAzure,
		Region:                    az.region,
	}
}

func (az *AzureClient) IsAvailable() (bool, error) {
	// TODO
	return true, nil
}

func filterLabelsForTags(resource string, labels map[string]string) (map[string]*string, error) {
	illegalKeys := []string{"Node", cloud.ControllerTagKey}
	allErrs := []error{}
	filteredLabels := make(map[string]*string)
	labelKeys := make([]string, len(labels))
	i := 0
	for k := range labels {
		labelKeys[i] = k
		i++
	}
	sort.Strings(labelKeys)
	for i, k := range labelKeys {
		v := labels[k]
		// constraints:
		// <= 15 tags (reserve 5 for milpa)
		// key - 512 characters
		// value - 255 chars
		// Key can't be one of our internal Milpa tag keys
		if i >= maxAzureUserTags {
			e := fmt.Errorf("Error tagging resource %s: Users are limited to %d tags", resource, maxAzureUserTags)
			allErrs = append(allErrs, e)
			break
		}
		if util.StringInSlice(k, illegalKeys) {
			allErrs = append(allErrs,
				fmt.Errorf("Error tagging instance %s, %s in illegal keys: %v",
					resource, k, illegalKeys))
			continue
		}
		if len(k) > maxAzureTagKeyLength {
			allErrs = append(allErrs,
				fmt.Errorf("Error tagging instance %s, keys are limited to 127 chars", resource))
			continue
		}
		if len(v) > maxAzureTagValueLength {
			allErrs = append(allErrs,
				fmt.Errorf("Error tagging instance %s, values are limited to 255 chars", resource))
			continue
		}
		k = replaceReservedTagChars(k)
		filteredLabels[k] = &v
	}
	var err error
	if len(allErrs) > 0 {
		err = errors.NewAggregate(allErrs)
	}
	return filteredLabels, err
}

func replaceReservedTagChars(s string) string {
	return strings.Map(func(c rune) rune {
		for _, r := range reservedTagCharacters {
			if c == r {
				return '-'
			}
		}
		return c
	}, s)
}

func (az *AzureClient) GetRegistryAuth() (string, string, error) {
	return "", "", fmt.Errorf("Azure registry not implemented (ACR is being phased out at the start of 2019")
}

// Display name is like "Central India", "East US 2" or "Brazil South"
// name is westindia, westus2, brazilsouth
// seems to be always concat and lowercase.
// to see all locations run: az account list-locations
func (az *AzureClient) locationName() string {
	return strings.ToLower(strings.Replace(az.region, " ", "", -1))
}
