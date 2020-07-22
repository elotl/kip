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
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/util/validation/field"
)

// Up to this point, all of our validation is static validation of api
// types. However, it would be good to be able to validate that a
// user's inputs don't conflict with parameters that are dependent on
// things outside our system.  For example, the state of the cloud,
// the user's instanceType or the state of an any internal or external
// system.

type StatefulValidator struct {
	cloudProvider string
	vpcCIDRs      []string
}

func NewStatefulValidator(cloudProvider string, vpcCIDRs []string) *StatefulValidator {
	return &StatefulValidator{
		cloudProvider: cloudProvider,
		vpcCIDRs:      vpcCIDRs,
	}
}

func (v *StatefulValidator) ValidatePodSpec(spec *api.PodSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if v.cloudProvider == cloud.ProviderAzure {
		if spec.Resources.SustainedCPU != nil && *spec.Resources.SustainedCPU {
			msg := "Azure does not support burstable instances with sustained CPUs"
			allErrs = append(allErrs, field.Invalid(fldPath.Child("resources.sustainedCPU"), *spec.Resources.SustainedCPU, msg))
		}
		if spec.Spot.Policy != api.SpotNever {
			msg := "Spot instances/Low Priority VMs are not supported (yet) with Azure, only spec.spot.policy: Never is supported"
			allErrs = append(allErrs, field.Invalid(fldPath.Child("spot.policy"), spec.Spot.Policy, msg))
		}
	}
	return allErrs
}
