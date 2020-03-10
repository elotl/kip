package azure

import (
	"fmt"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/stretchr/testify/assert"
)

//func matchSpec(name, value string, spec cloud.BootImageSpec) bool
func TestMatchSpec(t *testing.T) {
	testCases := []struct {
		Properties map[string]string
		Spec       cloud.BootImageSpec
		Matches    bool
	}{
		{
			Properties: map[string]string{
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec:    cloud.BootImageSpec{},
			Matches: true,
		},
		{
			Properties: map[string]string{
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"name": "elotl-kipdev-*",
			},
			Matches: true,
		},
		{
			Properties: map[string]string{
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Matches: true,
		},
		{
			Properties: map[string]string{
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"name": "elotl-milpadev-*",
			},
			Matches: false,
		},
		{
			Properties: map[string]string{
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"id": "myid123",
			},
			Matches: false,
		},
		{
			Properties: map[string]string{
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"name": "*-kipdev-*",
			},
			Matches: true,
		},
		{
			Properties: map[string]string{
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"name": "*",
			},
			Matches: true,
		},
		{
			Properties: map[string]string{
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"id":   "myid9999",
				"name": "elotl-kipdev-*",
			},
			Matches: false,
		},
		{
			Properties: map[string]string{
				"id":   "myid9999",
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"id":   "myid1111",
				"name": "elotl-kipdev-*",
			},
			Matches: false,
		},
		{
			Properties: map[string]string{
				"id":   "myid9999",
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"id":   "myid9999",
				"name": "elotl-kipdev-*",
			},
			Matches: true,
		},
		{
			Properties: map[string]string{
				"id":   "myid9999",
				"name": "elotl-kipdev-1234-20200222-010203",
			},
			Spec: cloud.BootImageSpec{
				"id":   "*",
				"name": "elotl-kipdev-*",
			},
			Matches: true,
		},
	}
	for i, tc := range testCases {
		matches := matchSpec(tc.Properties, tc.Spec)
		assert.Equal(t, tc.Matches, matches, fmt.Sprintf(
			"test #%d %+v failed", i+1, tc))
	}
}
