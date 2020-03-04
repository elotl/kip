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
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/uber-go/atomic"
)

func getImageController() *ImageController {
	client := &AzureClient{
		region:       "US West",
		controllerID: "img-ctl-cluster",
	}
	return &ImageController{
		controllerID: client.controllerID,
		bootImageTags: cloud.BootImageTags{
			Company: "elotl",
		},
		az:                client,
		resourceGroupName: regionalResourceGroupName(client.region),
		isSynced:          atomic.NewBool(false),
	}
}

func TestImageParametersMatch(t *testing.T) {
	ctl := getImageController()
	// A bit convoluted, we create a mock status keeper and put it in our
	// az cloudStatus but the status is backed by a MockCloudClient
	c := cloud.NewMockClient()
	c.Subnets = []cloud.SubnetAttributes{
		{
			AZ: "",
		},
	}
	c.StatusKeeperGetter = func() cloud.StatusKeeper {
		s, _ := cloud.NewAZSubnetStatus(c)
		return s
	}
	ctl.az.cloudStatus = c.CloudStatusKeeper()
	img := compute.Image{
		ImageProperties: &compute.ImageProperties{
			StorageProfile: &compute.ImageStorageProfile{
				ZoneResilient: to.BoolPtr(true),
			},
		},
	}
	assert.True(t, ctl.imageParametersMatch(img))
	// Now test out in a location that supports azs
	c.Subnets = []cloud.SubnetAttributes{
		{
			AZ: "1",
		},
	}
	ctl.az.cloudStatus = c.CloudStatusKeeper()
	assert.True(t, c.CloudStatusKeeper().SupportsAvailabilityZones())
	assert.True(t, ctl.imageParametersMatch(img))

	img.StorageProfile.ZoneResilient = to.BoolPtr(false)
	assert.False(t, ctl.imageParametersMatch(img))
	img.StorageProfile.ZoneResilient = nil
	assert.False(t, ctl.imageParametersMatch(img))
}
