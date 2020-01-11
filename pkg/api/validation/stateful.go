package validation

import (
	"fmt"
	"strconv"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/validation/field"
)

// Up to this point, all of our validation is static validation of api
// types. However, it would be good to be able to validate that a
// user's inputs don't conflict with parameters that are dependent on
// things outside our system.  For example, the state of the cloud,
// the user's instanceType or the state of an any internal or external
// system.

type StatefulValidator struct {
	cloudStatus   cloud.StatusKeeper
	cloudProvider string
	vpcCIDRs      []string
}

func NewStatefulValidator(status cloud.StatusKeeper, cloudProvider string, vpcCIDRs []string) *StatefulValidator {
	return &StatefulValidator{
		cloudStatus:   status,
		cloudProvider: cloudProvider,
		vpcCIDRs:      vpcCIDRs,
	}
}

func (v *StatefulValidator) ValidatePodSpec(spec *api.PodSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if spec.Placement.AvailabilityZone != "" {
		if status, ok := v.cloudStatus.(*cloud.LinkedAZSubnetStatus); ok {
			subnets := status.GetAllAZSubnets(spec.Placement.AvailabilityZone, spec.Resources.PrivateIPOnly)
			if len(subnets) == 0 {
				addressType := "public"
				if spec.Resources.PrivateIPOnly {
					addressType = "private"
				}
				msg := fmt.Sprintf("Invalid Availability Zone. No %s address subnets found in %s", addressType, spec.Placement.AvailabilityZone)
				allErrs = append(allErrs, field.Invalid(fldPath.Child("placement.availabilityZone"), spec.Placement.AvailabilityZone, msg))
			}
		} else if status, ok := v.cloudStatus.(*cloud.AZSubnetStatus); ok {
			azs := status.GetAllAvailabilityZones()
			if !util.StringInSlice(spec.Placement.AvailabilityZone, azs) {
				msg := fmt.Sprintf("Invalid Availability Zone %s. Available zones: %v", spec.Placement.AvailabilityZone, azs)
				allErrs = append(allErrs, field.Invalid(fldPath.Child("placement.availabilityZone"), spec.Placement.AvailabilityZone, msg))
			}
		}
	}
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

func (v *StatefulValidator) ValidateService(svc *api.Service) field.ErrorList {
	allErrs := field.ErrorList{}
	portsPath := field.NewPath("spec.ports")
	if v.cloudProvider == cloud.ProviderAzure {
		for i, port := range svc.Spec.Ports {
			portPath := portsPath.Index(i)
			allErrs = append(allErrs, v.validateServicePort(&port, false, portPath)...)
		}
	}
	return allErrs
}

func (v *StatefulValidator) validateServicePort(sp *api.ServicePort, requireName bool, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if v.cloudProvider == cloud.ProviderAzure &&
		sp.Protocol == api.ProtocolICMP {
		msg := "Azure does not support ICMP protocol service ports"
		allErrs = append(allErrs, field.Invalid(fldPath, sp.Protocol, msg))
	}
	return allErrs
}

func (v *StatefulValidator) isInternalSvc(svc *api.Service) bool {
	internal := true
	for i := range svc.Spec.SourceRanges {
		if svc.Spec.SourceRanges[i] == "VPC" {
			continue
		}
		if !util.CIDRInsideCIDRs(svc.Spec.SourceRanges[i], v.vpcCIDRs) {
			internal = false
			break
		}
	}
	return internal
}

func validateInRange(aMap map[string]string, key string, floor, ceil int, fldPath *field.Path) *field.Error {
	strVal := aMap[key]
	if strVal == "" {
		return nil
	}
	val, err := strconv.Atoi(strVal)
	if err != nil {
		msg := fmt.Sprintf("%s must be an integer", key)
		return field.Invalid(fldPath.Key(key), strVal, msg)
	}
	if val < floor || val > ceil {
		msg := fmt.Sprintf("%s must be in the range [%d, %d]", key, floor, ceil)
		return field.Invalid(fldPath.Key(key), strVal, msg)
	}
	return nil
}
