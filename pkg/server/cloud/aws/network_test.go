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

package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/stretchr/testify/assert"
)

func createTestSubnet(id string) *ec2.Subnet {
	sn := &ec2.Subnet{
		SubnetId: aws.String(id),
	}
	return sn
}

func createTestRouteTable(id string, assoc []string, routeGW []string) *ec2.RouteTable {
	associations := make([]*ec2.RouteTableAssociation, 0, len(assoc))
	for _, a := range assoc {
		associations = append(associations, &ec2.RouteTableAssociation{
			RouteTableId: aws.String(id),
			SubnetId:     aws.String(a),
		})
	}
	routes := make([]*ec2.Route, 0, len(routeGW))
	for _, gw := range routeGW {
		routes = append(routes, &ec2.Route{GatewayId: aws.String(gw)})
	}
	rt := &ec2.RouteTable{
		Associations: associations,
		Routes:       routes,
	}
	return rt
}

func TestMakeMilpaSubnetsAddressType(t *testing.T) {
	s1ID := "sub-1"
	s2ID := "sub-2"

	rt1ID := "rt-1"
	rt2ID := "rt-2"

	// single subnet using the default route table comes back as AnyAddress
	s1 := createTestSubnet(s1ID)
	s2 := createTestSubnet(s2ID)
	mainRT := createTestRouteTable(rt1ID, []string{"nope"}, []string{"igw-1234"})
	mainRT.Associations[0].Main = aws.Bool(true)
	sns, err := makeMilpaSubnets([]*ec2.Subnet{s1}, []*ec2.RouteTable{mainRT})
	assert.NoError(t, err)
	assert.Len(t, sns, 1)
	assert.Equal(t, cloud.AnyAddress, sns[0].AddressAffinity)

	// test that private and public subnets are differentiated
	pubRT := createTestRouteTable(rt1ID, []string{s1ID}, []string{"igw-1234"})
	privRT := createTestRouteTable(rt2ID, []string{s2ID}, []string{"aaa"})
	sns, err = makeMilpaSubnets([]*ec2.Subnet{s1, s2}, []*ec2.RouteTable{pubRT, privRT})
	assert.NoError(t, err)
	assert.Len(t, sns, 2)
	addressing := []cloud.SubnetAddressAffinity{}
	for _, sn := range sns {
		addressing = append(addressing, sn.AddressAffinity)
	}
	expected := []cloud.SubnetAddressAffinity{cloud.PublicAddress, cloud.PrivateAddress}
	assert.ElementsMatch(t, expected, addressing)
}
