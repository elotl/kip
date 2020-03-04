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

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/assert"
)

func TestGetMilpaIPConfiguration(t *testing.T) {
	cases := []struct {
		iface   network.Interface
		isError bool
	}{
		{
			iface:   network.Interface{},
			isError: true,
		},
		{
			iface: network.Interface{
				InterfacePropertiesFormat: &network.InterfacePropertiesFormat{},
			},
			isError: true,
		},
		{
			iface: network.Interface{
				InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
					IPConfigurations: &[]network.InterfaceIPConfiguration{
						{
							Name: to.StringPtr(milpaIPConfig),
						},
					},
				},
			},
			isError: false,
		},
		{
			iface: network.Interface{
				InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
					IPConfigurations: &[]network.InterfaceIPConfiguration{
						{
							Name: to.StringPtr("something else"),
						},
						{
							Name: to.StringPtr(milpaIPConfig),
						},
					},
				},
			},
			isError: false,
		},
	}
	for _, tc := range cases {
		ipconfig, err := getMilpaIPConfiguration(tc.iface)
		if tc.isError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, ipconfig)
			assert.Equal(t, milpaIPConfig, to.String(ipconfig.Name))
		}
	}
}

func TestGetSecurityGroupsFromInterface(t *testing.T) {
	cases := []struct {
		iface network.Interface
		names []string
	}{
		{
			iface: network.Interface{
				InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
					IPConfigurations: &[]network.InterfaceIPConfiguration{
						{
							Name: to.StringPtr(milpaIPConfig),
							InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
								ApplicationSecurityGroups: &[]network.ApplicationSecurityGroup{
									{ID: to.StringPtr("/subscriptions/4e84e89a-b806-4d7d-900b-cae8cb640215/resourceGroups/milpa-testcluster/providers/Microsoft.Network/applicationSecurityGroups/my-svc1")},
									{ID: to.StringPtr("/subscriptions/4e84e89a-b806-4d7d-900b-cae8cb640215/resourceGroups/milpa-testcluster/providers/Microsoft.Network/applicationSecurityGroups/my-svc2")},
								},
							},
						},
					},
				},
			},
			names: []string{"my-svc1", "my-svc2"},
		},
	}
	for _, tc := range cases {
		ids := getSecurityGroupsFromInterface(tc.iface)
		assert.Len(t, ids, len(tc.names))
		sgnames := make([]string, len(ids))
		for i := range ids {
			sgnames[i] = ids[i].Name
		}
		assert.Equal(t, tc.names, sgnames)
	}
}
