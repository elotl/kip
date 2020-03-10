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

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util/rand"
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

//func SortImagesByCreationTime(images []Image)
func TestSortImagesByCreationTime(t *testing.T) {
	images := []Image{}
	SortImagesByCreationTime(images)
	assert.Len(t, images, 0)
	images = make([]Image, 1000)
	for i := range images {
		ts := rand.Timestamp()
		images[i] = Image{
			Name:         fmt.Sprintf("image-%010d", ts.Unix()),
			CreationTime: &ts,
		}
	}
	SortImagesByCreationTime(images)
	prev := images[0]
	for i := 1; i < len(images)-1; i++ {
		prevCT := *prev.CreationTime
		prevName := prev.Name
		ct := *images[i].CreationTime
		name := images[i].Name
		assert.True(
			t,
			prevCT.Before(ct),
			fmt.Sprintf("%d !< %d", prevCT.Unix(), ct.Unix()))
		assert.True(
			t,
			prevName < name,
			fmt.Sprintf("%s !< %s", prevName, name))
		prev = images[i]
	}
}
