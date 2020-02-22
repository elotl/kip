package validation

import (
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util/validation/field"
	"github.com/stretchr/testify/assert"
)

func TestValidateStatefulPodSpecPlacementLinkedStatus(t *testing.T) {
	cs, err := cloud.NewLinkedAZSubnetStatus(cloud.NewMockClient())
	assert.NoError(t, err)
	sv := StatefulValidator{cloudStatus: cs}
	azTests := []struct {
		az      string
		numErrs int
	}{
		{"", 0},
		{"us-east-1a", 0},
		{"us-east-2b", 1},
		{"foo", 1},
	}
	for i, testCase := range azTests {
		ps := api.PodSpec{}
		ps.Placement.AvailabilityZone = testCase.az
		errs := sv.ValidatePodSpec(&ps, field.NewPath("field"))
		assert.Len(t, errs, testCase.numErrs, "test %d failed", i)
	}
}

func TestValidateStatefulPodSpecPlacement(t *testing.T) {
	cs, err := cloud.NewAZSubnetStatus(cloud.NewMockClient())
	assert.NoError(t, err)
	sv := StatefulValidator{cloudStatus: cs}
	azTests := []struct {
		az      string
		numErrs int
	}{
		{"", 0},
		{"us-east-1a", 0},
		{"us-east-2b", 1},
		{"foo", 1},
	}
	for i, testCase := range azTests {
		ps := api.PodSpec{}
		ps.Placement.AvailabilityZone = testCase.az
		errs := sv.ValidatePodSpec(&ps, field.NewPath("field"))
		assert.Len(t, errs, testCase.numErrs, "test %d failed", i)
	}
}

func TestValidateStatefulPodSpecSustainedCPU(t *testing.T) {
	sv := StatefulValidator{}
	truth := true
	f := false
	azTests := []struct {
		provider  string
		resources api.ResourceSpec
		numErrs   int
	}{
		{
			provider: cloud.ProviderAWS,
			resources: api.ResourceSpec{
				SustainedCPU: &truth,
			},
			numErrs: 0,
		},
		{
			provider: cloud.ProviderAzure,
			resources: api.ResourceSpec{
				SustainedCPU: &f,
			},
			numErrs: 0,
		},
		{
			provider: cloud.ProviderAzure,
			resources: api.ResourceSpec{
				SustainedCPU: &truth,
			},
			numErrs: 1,
		},
	}
	for i, testCase := range azTests {
		ps := api.NewPod().Spec
		ps.Resources = testCase.resources
		sv.cloudProvider = testCase.provider
		errs := sv.ValidatePodSpec(&ps, field.NewPath("field"))
		assert.Len(t, errs, testCase.numErrs, "test %d failed", i)
	}
}
