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

package validation

import (
	"testing"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util/validation/field"
	"github.com/stretchr/testify/assert"
)

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
