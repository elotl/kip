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
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/stretchr/testify/assert"
)

func makeIR(p InstancePort, sr []string) []IngressRule {
	return []IngressRule{NewIngressRule(p, sr[0])}
}

func TestMergeSecurityGroupsSimple(t *testing.T) {
	// create table of cases
	vpcSourceRange := []string{"172.0.0.0/8"}
	publicSourceRange := []string{"0.0.0.0/0"}
	a := InstancePort{Port: 1, Protocol: api.ProtocolTCP, PortRangeSize: 1}
	b := InstancePort{Port: 2, Protocol: api.ProtocolTCP, PortRangeSize: 1}
	c := InstancePort{Port: 3, Protocol: api.ProtocolTCP, PortRangeSize: 1}
	d := InstancePort{Port: 1, Protocol: api.ProtocolUDP, PortRangeSize: 1}
	e := InstancePort{Port: 2, Protocol: api.ProtocolUDP, PortRangeSize: 1}
	f := InstancePort{Port: 3, Protocol: api.ProtocolUDP, PortRangeSize: 1}
	sgEmpty := SecurityGroup{}
	sgA := SecurityGroup{
		Ports:        []InstancePort{a},
		SourceRanges: vpcSourceRange,
	}
	sgAB := SecurityGroup{
		Ports:        []InstancePort{a, b},
		SourceRanges: vpcSourceRange,
	}
	sgABCD := SecurityGroup{
		Ports:        []InstancePort{a, b, c, d},
		SourceRanges: vpcSourceRange,
	}
	//only new
	add, delete := MergeSecurityGroups(sgEmpty, []InstancePort{a}, vpcSourceRange)
	assert.Len(t, add, 1)
	assert.Len(t, delete, 0)
	assert.ElementsMatch(t, add, makeIR(a, vpcSourceRange))
	// only old
	add, delete = MergeSecurityGroups(sgA, []InstancePort{}, vpcSourceRange)
	assert.Len(t, add, 0)
	assert.Len(t, delete, 1)
	assert.ElementsMatch(t, delete, makeIR(a, vpcSourceRange))
	//one new, one old
	add, delete = MergeSecurityGroups(sgA, []InstancePort{b}, vpcSourceRange)
	assert.Len(t, add, 1)
	assert.Len(t, delete, 1)
	assert.ElementsMatch(t, makeIR(b, vpcSourceRange), add)
	assert.ElementsMatch(t, makeIR(a, vpcSourceRange), delete)
	// equal
	add, delete = MergeSecurityGroups(sgA, []InstancePort{a}, vpcSourceRange)
	assert.Len(t, add, 0)
	assert.Len(t, delete, 0)
	// two equal
	add, delete = MergeSecurityGroups(sgAB, []InstancePort{a, b}, vpcSourceRange)
	assert.Len(t, add, 0)
	assert.Len(t, delete, 0)
	// one new, one common, one old
	add, delete = MergeSecurityGroups(sgAB, []InstancePort{b, c}, vpcSourceRange)
	assert.Len(t, add, 1)
	assert.Len(t, delete, 1)
	assert.ElementsMatch(t, makeIR(c, vpcSourceRange), add)
	assert.ElementsMatch(t, makeIR(a, vpcSourceRange), delete)

	// two new, two common, two old
	add, delete = MergeSecurityGroups(sgABCD, []InstancePort{c, d, e, f}, vpcSourceRange)
	assert.Len(t, add, 2)
	assert.Len(t, delete, 2)
	assert.ElementsMatch(t, append(makeIR(e, vpcSourceRange), makeIR(f, vpcSourceRange)...), add)
	assert.ElementsMatch(t, append(makeIR(a, vpcSourceRange), makeIR(b, vpcSourceRange)...), delete)

	//public -> private returns everything
	add, delete = MergeSecurityGroups(sgAB, []InstancePort{a, b}, publicSourceRange)
	assert.Len(t, add, 2)
	assert.Len(t, delete, 2)
}

func TestMergeSecurityGroupSourceRanges(t *testing.T) {
	// sg1 - ports 1 and 2, ranges 172.16.0.0/16, 100.64.0.0/16
	// sg2 - ports 2 and 3, ranges 172.16.0.0/16, 192.168.1.1/32, 0.0.0.0/0
	sA := "172.16.0.0/16"
	sB := "100.64.0.0/16"
	sC := "192.168.1.1/32"
	sD := "0.0.0.0/0"
	srA := []string{sA, sB}
	srB := []string{sB, sC, sD}
	a := InstancePort{Port: 1, Protocol: api.ProtocolTCP}
	b := InstancePort{Port: 2, Protocol: api.ProtocolTCP}
	c := InstancePort{Port: 3, Protocol: api.ProtocolTCP}
	sgA := SecurityGroup{
		Ports:        []InstancePort{a, b},
		SourceRanges: srA,
	}
	add, delete := MergeSecurityGroups(sgA, []InstancePort{b, c}, srB)
	assert.Len(t, add, 5)
	assert.Len(t, delete, 3)
}

func TestBootImageTags(t *testing.T) {
	bit := BootImageTags{}
	bit0 := BootImageTags{}
	assert.Equal(t, bit.String(), "----")
	bit0.Set(bit.String())
	assert.Equal(t, bit0.String(), "----")
	assert.True(t, bit.Matches(bit0))
	bit = BootImageTags{
		Company: "elotl",
		Product: "awesomenewproduct",
		Version: "123",
		Date:    "20180830",
		Time:    "1613",
	}
	assert.Equal(t, bit.String(), "elotl-awesomenewproduct-123-20180830-1613")
	bit0.Set(bit.String())
	assert.Equal(t, bit0.String(), "elotl-awesomenewproduct-123-20180830-1613")
	assert.True(t, bit.Matches(bit0))
	bit = BootImageTags{
		Company: "elotl",
		Version: "123",
		Date:    "20180830",
	}
	assert.Equal(t, bit.String(), "elotl--123-20180830-")
	bit0.Set(bit.String())
	assert.Equal(t, bit0.String(), "elotl--123-20180830-")
	assert.True(t, bit.Matches(bit0))
	bit = BootImageTags{
		Company: "elotl",
		Product: "awesomenewproduct",
		Version: "123",
		Date:    "20180830",
		Time:    "1613",
	}
	assert.True(t, bit.Matches(bit0))
}

func TestSortImages(t *testing.T) {
	images := []Image{
		Image{
			Name: "elotl-milpa-1-20180101-223344",
		},
		Image{
			Name: "elotl-milpa-2-20180101-010101",
		},
		Image{
			Name: "elotl-milpa-2-20180102-020202",
		},
		Image{
			Name: "elotl-milpa-1-20180102-010101",
		},
	}
	SortImages(images)
	assert.Equal(t, "elotl-milpa-1-20180101-223344", images[0].Name)
	assert.Equal(t, "elotl-milpa-1-20180102-010101", images[1].Name)
	assert.Equal(t, "elotl-milpa-2-20180101-010101", images[2].Name)
	assert.Equal(t, "elotl-milpa-2-20180102-020202", images[3].Name)
}

func TestFilterImages(t *testing.T) {
	images := []Image{
		Image{
			Id:   "1",
			Name: "elotl-milpa-1-20180101-010101",
		},
		Image{
			Id:   "2",
			Name: "elotl-milpa-2-20180101-010101",
		},
		Image{
			Id:   "3",
			Name: "elotl-milpadev-2-20180102-020202",
		},
		Image{
			Id:   "4",
			Name: "elotl-milpadev-1-20180102-010101",
		},
	}
	tags := BootImageTags{
		Company: "elotl",
		Product: "milpa",
	}
	res := FilterImages(images, tags)
	assert.Len(t, res, 2)
	assert.Contains(t, res, images[0])
	assert.Contains(t, res, images[1])
	tags = BootImageTags{
		Company: "elotl",
		Product: "milpadev",
	}
	res = FilterImages(images, tags)
	assert.Len(t, res, 2)
	assert.Contains(t, res, images[2])
	assert.Contains(t, res, images[3])
	tags = BootImageTags{
		Company: "elotl",
		Product: "milpa",
		Version: "1",
	}
	res = FilterImages(images, tags)
	assert.Len(t, res, 1)
	assert.Contains(t, res, images[0])
	tags = BootImageTags{
		Company: "elotl",
		Product: "milpadev",
		Version: "2",
	}
	res = FilterImages(images, tags)
	assert.Len(t, res, 1)
	assert.Contains(t, res, images[2])
}

func TestComparePorts(t *testing.T) {
	tests := []struct {
		a    InstancePort
		b    InstancePort
		less bool
	}{
		{
			a:    InstancePort{Port: 5},
			b:    InstancePort{Port: 7},
			less: true,
		},
		{
			a:    InstancePort{Port: 7},
			b:    InstancePort{Port: 5},
			less: false,
		},
		{
			a:    InstancePort{Port: 5, PortRangeSize: 1, Protocol: api.ProtocolUDP},
			b:    InstancePort{Port: 5, PortRangeSize: 3, Protocol: api.ProtocolTCP},
			less: true,
		},
		{
			a:    InstancePort{Port: 5, PortRangeSize: 4, Protocol: api.ProtocolUDP},
			b:    InstancePort{Port: 5, PortRangeSize: 3, Protocol: api.ProtocolTCP},
			less: false,
		},
		{
			a:    InstancePort{Port: 5, PortRangeSize: 3, Protocol: api.ProtocolUDP},
			b:    InstancePort{Port: 5, PortRangeSize: 3, Protocol: api.ProtocolTCP},
			less: false,
		},
	}
	for i, test := range tests {
		if lessPorts(test.a, test.b) != test.less {
			t.Errorf("Test %d failed: %v", i, test)
		}
	}
}
