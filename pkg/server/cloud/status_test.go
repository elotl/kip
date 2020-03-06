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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStatusSubnetSupportsAddressType(t *testing.T) {
	assert.True(t, subnetSupportsAddressType(PublicAddress, false))
	assert.True(t, subnetSupportsAddressType(PrivateAddress, true))
	assert.True(t, subnetSupportsAddressType(AnyAddress, true))
	assert.True(t, subnetSupportsAddressType(AnyAddress, false))
	assert.False(t, subnetSupportsAddressType(PublicAddress, true))
	assert.False(t, subnetSupportsAddressType(PrivateAddress, false))
}

func GetMockMultiSubnetNetworker() *MockCloudClient {
	c := NewMockClient()
	c.Subnets = []SubnetAttributes{
		{
			ID:                 "sub-1111",
			CIDR:               "172.16.0.0/16",
			AZ:                 "us-east-1a",
			AddressAffinity:    PublicAddress,
			AvailableAddresses: 250,
		},
		{
			ID:                 "sub-2222",
			CIDR:               "172.17.0.0/16",
			AZ:                 "us-east-1a",
			AddressAffinity:    PrivateAddress,
			AvailableAddresses: 250,
		},
		{
			ID:                 "sub-3333",
			CIDR:               "172.18.0.0/16",
			AZ:                 "us-east-1b",
			AddressAffinity:    PrivateAddress,
			AvailableAddresses: 250,
		},
		{
			ID:                 "sub-4444",
			CIDR:               "172.19.0.0/16",
			AZ:                 "us-east-1c",
			AddressAffinity:    PublicAddress,
			AvailableAddresses: 250,
		},
		{
			ID:                 "sub-5555",
			CIDR:               "172.20.0.0/16",
			AZ:                 "us-east-1c",
			AddressAffinity:    PublicAddress,
			AvailableAddresses: 250,
		},
	}
	return c
}

func TestLinkedStatusGetAllAZSubnets(t *testing.T) {
	c := GetMockMultiSubnetNetworker()
	s, err := NewLinkedAZSubnetStatus(c)
	assert.NoError(t, err)
	assert.True(t, s.SupportsAvailabilityZones())
	sns := s.GetAllAZSubnets("us-east-1a", false)
	assert.Len(t, sns, 1)
	assert.Equal(t, "sub-1111", sns[0].ID)
	sns = s.GetAllAZSubnets("us-east-1a", true)
	assert.Len(t, sns, 1)
	assert.Equal(t, "sub-2222", sns[0].ID)
	sns = s.GetAllAZSubnets("us-east-1c", false)
	assert.Len(t, sns, 2)
	sns = s.GetAllAZSubnets("us-east-1b", false)
	assert.Len(t, sns, 0)
}

func TestLinkedStatusGetAvailableZones(t *testing.T) {
	c := GetMockMultiSubnetNetworker()
	s, err := NewLinkedAZSubnetStatus(c)
	assert.NoError(t, err)
	itype := "t2.nano"
	zones := s.GetAvailableZones(itype, false, false)
	allPublicZones := []string{"us-east-1a", "us-east-1c"}
	assert.ElementsMatch(t, allPublicZones, zones)
	s.AddUnavailableZone(itype, false, "us-east-1a")
	zones = s.GetAvailableZones(itype, false, false)
	assert.ElementsMatch(t, []string{"us-east-1c"}, zones)
	zones = s.GetAvailableZones("t2.micro", false, false)
	assert.ElementsMatch(t, allPublicZones, zones)
	// We treat spot availability and ondemand availability totally
	// separately
	zones = s.GetAvailableZones(itype, true, false)
	assert.ElementsMatch(t, allPublicZones, zones)
}

func TestLinkedStatusGetAvailableSubnets(t *testing.T) {
	c := GetMockMultiSubnetNetworker()
	s, err := NewLinkedAZSubnetStatus(c)
	assert.NoError(t, err)
	itype := "t2.nano"
	sns := s.GetAvailableSubnets(itype, false, false)
	assert.Len(t, sns, 3)
	sns = s.GetAvailableAZSubnets(itype, "us-east-1c", false, false)
	assert.Len(t, sns, 2)
	sns = s.GetAvailableAZSubnets(itype, "us-east-1c", false, true)
	assert.Len(t, sns, 0)
	s.AddUnavailableSubnet(itype, false, "sub-1111")
	sns = s.GetAvailableSubnets(itype, false, false)
	assert.Len(t, sns, 2)
	sns = s.GetAvailableAZSubnets(itype, "us-east-1c", false, false)
	assert.Len(t, sns, 2)
	sns = s.GetAvailableAZSubnets(itype, "us-east-1a", false, false)
	assert.Len(t, sns, 0)
}

func TestLinkedStatusAddUnavailable(t *testing.T) {
	c := GetMockMultiSubnetNetworker()
	s, err := NewLinkedAZSubnetStatus(c)
	assert.NoError(t, err)
	itype := "t2.nano"
	s.AddUnavailableInstance(itype, false)
	assert.Len(t, s.GetAvailableSubnets(itype, false, false), 0)
	assert.True(t, s.IsUnavailableZone(itype, false, false, ""))
	itype = "t2.micro"
	s.AddUnavailableZone(itype, false, "us-east-1c")
	assert.Len(t, s.GetAvailableSubnets(itype, false, false), 1)
	fmt.Println(s.GetAvailableSubnets(itype, false, false))
	assert.False(t, s.IsUnavailableZone(itype, false, false, "us-east-1a"))
	itype = "t2.small"
	s.AddUnavailableSubnet(itype, false, "sub-1111")
	assert.Len(t, s.GetAvailableSubnets(itype, false, false), 2)
	assert.True(t, s.IsUnavailableZone(itype, false, false, "us-east-1a"))
}

func TestLinkedStatusSubnetRefresh(t *testing.T) {
	c := NewMockClient()
	s, err := NewLinkedAZSubnetStatus(c)
	assert.NoError(t, err)
	assert.Len(t, s.subnets, 3)
	c.Subnets = []SubnetAttributes{
		{
			ID:                 "sub-1234",
			CIDR:               "172.16.0.0/16",
			AZ:                 "us-east-1a",
			AddressAffinity:    PublicAddress,
			AvailableAddresses: 10,
		},
		{
			ID:                 "sub-5678",
			CIDR:               "172.16.0.0/16",
			AZ:                 "us-east-1b",
			AddressAffinity:    PublicAddress,
			AvailableAddresses: 10,
		},
	}
	subnetRefreshPeriod = 50 * time.Millisecond
	s.Start()
	time.Sleep(250 * time.Millisecond)
	assert.Len(t, s.subnets, 2)
	assert.Len(t, s.GetAllSubnets(), 2)
	// Test getting subnets when there is an error
	c.SubnetGetter = func() ([]SubnetAttributes, error) {
		return nil, fmt.Errorf("Subnets shouldnt change when there is an error")
	}
	time.Sleep(250 * time.Millisecond)
	fmt.Println("getting subnets")
	assert.Len(t, s.subnets, 2)
	assert.Len(t, s.GetAllSubnets(), 2)
	assert.Equal(t, s.subnets, s.GetAllSubnets())
	fmt.Println("waiting on stop")
	s.Stop()
}

func TestUnlinkedAZStatus(t *testing.T) {
	c := GetMockMultiSubnetNetworker()
	s, err := NewAZSubnetStatus(c)
	assert.NoError(t, err)
	assert.True(t, s.SupportsAvailabilityZones())
	assert.Len(t, s.subnets, 5)
	expectedZones := []string{"us-east-1a", "us-east-1b", "us-east-1c"}
	itype := "t2.nano"
	zones := s.GetAvailableZones(itype, false, false)
	assert.ElementsMatch(t, expectedZones, zones)
	s.AddUnavailableZone(itype, false, "us-east-1a")
	zones = s.GetAvailableZones(itype, false, false)
	assert.ElementsMatch(t, expectedZones[1:], zones)
	assert.True(t, s.IsUnavailableZone(itype, false, false, "us-east-1a"))
	assert.False(t, s.IsUnavailableZone(itype, false, false, "us-east-1c"))
	s.AddUnavailableInstance(itype, false)
	zones = s.GetAvailableZones(itype, false, false)
	assert.Len(t, zones, 0)
}

func getSNIDs(sns []SubnetAttributes) []string {
	snIDs := make([]string, len(sns))
	for i := range sns {
		snIDs[i] = sns[i].ID
	}
	return snIDs
}

func TestUnlinkedSubnetStatus(t *testing.T) {
	c := GetMockMultiSubnetNetworker()
	s, err := NewAZSubnetStatus(c)
	assert.NoError(t, err)
	assert.Len(t, s.subnets, 5)
	expectedSubnetIDs := []string{"sub-1111", "sub-4444", "sub-5555"}
	itype := "t2.nano"
	sns := s.GetAvailableSubnets(itype, false, false)
	fmt.Println(s.subnets)
	fmt.Println(s.availabilityZones)
	snIDs := getSNIDs(sns)
	assert.ElementsMatch(t, expectedSubnetIDs, snIDs)
	s.AddUnavailableSubnet(itype, false, "sub-1111")
	sns = s.GetAvailableSubnets(itype, false, false)
	snIDs = getSNIDs(sns)
	assert.ElementsMatch(t, expectedSubnetIDs[1:], snIDs)
	assert.True(t, s.IsUnavailableSubnet(itype, false, "sub-1111"))
	assert.False(t, s.IsUnavailableSubnet(itype, false, "sub-4444"))
	s.AddUnavailableInstance(itype, false)
	sns = s.GetAvailableSubnets(itype, false, false)
	assert.Len(t, sns, 0)
}

func TestUnlinkedGetAllSubnetsAndAZs(t *testing.T) {
	c := GetMockMultiSubnetNetworker()
	s, err := NewAZSubnetStatus(c)
	assert.NoError(t, err)
	assert.Len(t, s.GetAllSubnets(), 5)
	assert.Len(t, s.GetAllAvailabilityZones(), 3)
}
