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
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/stretchr/testify/assert"
)

func TestMakeIPPermissions(t *testing.T) {
	// make sure that port range size gets converted
	ipRange := []string{"192.168.1.0/24"}
	ec2IPRange := []*ec2.IpRange{&ec2.IpRange{CidrIp: &ipRange[0]}}
	tests := []struct {
		ir  cloud.IngressRule
		ipp ec2.IpPermission
	}{
		{
			ir: cloud.IngressRule{
				Port:          10,
				PortRangeSize: 1,
				Protocol:      api.ProtocolTCP,
				Source:        ipRange[0],
			},
			ipp: ec2.IpPermission{
				IpProtocol: aws.String("TCP"),
				FromPort:   aws.Int64(10),
				ToPort:     aws.Int64(10),
				IpRanges:   ec2IPRange,
			},
		},
		{
			ir: cloud.IngressRule{
				Port:          10,
				PortRangeSize: 0,
				Protocol:      api.ProtocolTCP,
				Source:        ipRange[0],
			},
			ipp: ec2.IpPermission{
				IpProtocol: aws.String("TCP"),
				FromPort:   aws.Int64(10),
				ToPort:     aws.Int64(10),
				IpRanges:   ec2IPRange,
			},
		},
		{
			ir: cloud.IngressRule{
				Port:          10,
				PortRangeSize: 10,
				Protocol:      api.ProtocolTCP,
				Source:        ipRange[0],
			},
			ipp: ec2.IpPermission{
				IpProtocol: aws.String("TCP"),
				FromPort:   aws.Int64(10),
				ToPort:     aws.Int64(19),
				IpRanges:   ec2IPRange,
			},
		},
	}
	for i, test := range tests {
		awsIPP := makeIPPermissions([]cloud.IngressRule{test.ir})[0]
		if *awsIPP.IpProtocol != *test.ipp.IpProtocol ||
			*awsIPP.FromPort != *test.ipp.FromPort ||
			*awsIPP.ToPort != *test.ipp.ToPort ||
			*awsIPP.IpRanges[0].CidrIp != *test.ipp.IpRanges[0].CidrIp {
			t.Errorf("Failed test %d: %v", i, test)
		}
	}
}

func TestAwsSGToMilpa(t *testing.T) {
	ipRange := []string{"192.168.1.0/24"}
	ec2IPRange := []*ec2.IpRange{&ec2.IpRange{CidrIp: &ipRange[0]}}
	tests := []struct {
		awssg *ec2.SecurityGroup
		sg    cloud.SecurityGroup
	}{
		{
			awssg: &ec2.SecurityGroup{
				GroupId:   aws.String("1"),
				GroupName: aws.String("Foo"),
				IpPermissions: []*ec2.IpPermission{{
					IpProtocol: aws.String("tcp"),
					FromPort:   aws.Int64(10),
					ToPort:     aws.Int64(10),
					IpRanges:   ec2IPRange,
				}},
			},
			sg: cloud.SecurityGroup{
				ID:   "1",
				Name: "Foo",
				Ports: []cloud.InstancePort{{
					Name:          "",
					Protocol:      api.ProtocolTCP,
					Port:          10,
					PortRangeSize: 1,
				}},
				SourceRanges: ipRange,
			},
		},
		{
			awssg: &ec2.SecurityGroup{
				GroupId:   aws.String("1"),
				GroupName: aws.String("Foo"),
				IpPermissions: []*ec2.IpPermission{{
					IpProtocol: aws.String("udp"),
					FromPort:   aws.Int64(10),
					ToPort:     aws.Int64(19),
					IpRanges:   ec2IPRange,
				}},
			},
			sg: cloud.SecurityGroup{
				ID:   "1",
				Name: "Foo",
				Ports: []cloud.InstancePort{{
					Name:          "",
					Protocol:      api.ProtocolUDP,
					Port:          10,
					PortRangeSize: 10,
				}},
				SourceRanges: ipRange,
			},
		},
	}
	for i, test := range tests {
		sg := awsSGToMilpa(test.awssg)
		assert.Equal(t, test.sg, sg, "Failed test %d: %v", i, test)
	}
}
