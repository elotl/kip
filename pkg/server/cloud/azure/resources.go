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
	"regexp"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/elotl/kip/pkg/util"
)

//const uuidFmt string = "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
const shortUUIDFmt string = "[0-9a-z]{20,28}"

var shortUUIDRegexp = regexp.MustCompile("^" + shortUUIDFmt + "$")

func (az *AzureClient) createResourceGroup(groupName string) error {
	// Todo, consider adding cluster tags to the group
	parameters := resources.Group{
		Location: to.StringPtr(az.region),
	}
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	_, err := az.groups.CreateOrUpdate(timeoutCtx, groupName, parameters)
	return err
}

func (az *AzureClient) ensureResourceGroups() error {
	err := az.ensureResourceGroup(regionalResourceGroupName(az.region))
	if err != nil {
		return err
	}
	return az.ensureResourceGroup(controllerResourceGroupName(az.controllerID))
}

func (az *AzureClient) ensureResourceGroup(rgName string) error {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	_, err := az.groups.Get(timeoutCtx, rgName)
	if isNotFoundError(err) {
		err = az.createResourceGroup(rgName)
		if err != nil {
			return util.WrapError(err, "Error creating cluster group %s", rgName)
		}
	} else if err != nil {
		return util.WrapError(err, "Error checking for existence of resource group %s", rgName)
	}
	return nil
}

func isNodeResourceGroup(name, controllerID string) bool {
	clusterGroupPrefix := util.CreateClusterResourcePrefix(controllerID)
	if strings.HasPrefix(name, clusterGroupPrefix) {
		maybeUUID := name[len(clusterGroupPrefix):]
		return shortUUIDRegexp.MatchString(maybeUUID)
	}
	return false
}

func (az *AzureClient) ListNodeResourceGroups() ([]string, error) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	page, err := az.groups.List(timeoutCtx, "", nil)
	if err != nil {
		return nil, err
	}
	groupNames := []string{}
	for page.NotDone() {
		rgs := page.Values()
		for _, rg := range rgs {
			name := to.String(rg.Name)
			// if the resource name starts with a cluster prefix and
			// ends with a UUID, then that's a milpa node resource
			// group
			if isNodeResourceGroup(name, az.controllerID) {
				groupNames = append(groupNames, name)
			}
		}
		err := page.Next()
		if err != nil {
			return groupNames, err
		}
	}
	return groupNames, nil
}

func (az *AzureClient) DeleteResourceGroup(groupID string) error {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	_, err := az.groups.Delete(timeoutCtx, groupID)
	return err
}
