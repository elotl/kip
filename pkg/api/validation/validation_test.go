/*
Copyright 2014 The Kubernetes Authors.
Copyright 2017 Elotl Inc.

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
	"fmt"
	"strings"
	"testing"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/sets"
	"github.com/elotl/cloud-instance-provider/pkg/util/validation/field"
)

// Ensure trailing slash is allowed in generate name
func TestValidateObjectMetaTrimsTrailingSlash(t *testing.T) {
	errs := ValidateObjectMeta(
		&api.ObjectMeta{Name: "test"},
		false,
		NameIsDNSSubdomain,
		field.NewPath("field"))
	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
}

func TestValidateAnnotations(t *testing.T) {
	successCases := []map[string]string{
		{"simple": "bar"},
		{"now-with-dashes": "bar"},
		{"1-starts-with-num": "bar"},
		{"1234": "bar"},
		{"simple/simple": "bar"},
		{"now-with-dashes/simple": "bar"},
		{"now-with-dashes/now-with-dashes": "bar"},
		{"now.with.dots/simple": "bar"},
		{"now-with.dashes-and.dots/simple": "bar"},
		{"1-num.2-num/3-num": "bar"},
		{"1234/5678": "bar"},
		{"1.2.3.4/5678": "bar"},
		{"UpperCase123": "bar"},
		{"a": strings.Repeat("b", totalAnnotationSizeLimitB-1)},
		{
			"a": strings.Repeat("b", totalAnnotationSizeLimitB/2-1),
			"c": strings.Repeat("d", totalAnnotationSizeLimitB/2-1),
		},
	}
	for i := range successCases {
		errs := ValidateAnnotations(successCases[i], field.NewPath("field"))
		if len(errs) != 0 {
			t.Errorf("case[%d] expected success, got %#v", i, errs)
		}
	}

	nameErrorCases := []struct {
		annotations map[string]string
		expect      string
	}{
		{map[string]string{"nospecialchars^=@": "bar"}, "must match the regex"},
		{map[string]string{"cantendwithadash-": "bar"}, "must match the regex"},
		{map[string]string{"only/one/slash": "bar"}, "must match the regex"},
		{map[string]string{strings.Repeat("a", 254): "bar"}, "must be no more than"},
	}
	for i := range nameErrorCases {
		errs := ValidateAnnotations(nameErrorCases[i].annotations, field.NewPath("field"))
		if len(errs) != 1 {
			t.Errorf("case[%d]: expected failure", i)
		} else {
			if !strings.Contains(errs[0].Detail, nameErrorCases[i].expect) {
				t.Errorf("case[%d]: error details do not include %q: %q", i, nameErrorCases[i].expect, errs[0].Detail)
			}
		}
	}
	totalSizeErrorCases := []map[string]string{
		{"a": strings.Repeat("b", totalAnnotationSizeLimitB)},
		{
			"a": strings.Repeat("b", totalAnnotationSizeLimitB/2),
			"c": strings.Repeat("d", totalAnnotationSizeLimitB/2),
		},
	}
	for i := range totalSizeErrorCases {
		errs := ValidateAnnotations(totalSizeErrorCases[i], field.NewPath("field"))
		if len(errs) != 1 {
			t.Errorf("case[%d] expected failure", i)
		}
	}
}

func TestValidateEnv(t *testing.T) {
	successCase := []api.EnvVar{
		{Name: "abc", Value: "value"},
		{Name: "ABC", Value: "value"},
		{Name: "AbC_123", Value: "value"},
		{Name: "abc", Value: ""},
	}
	if errs := validateEnv(successCase, field.NewPath("field")); len(errs) != 0 {
		t.Errorf("expected success: %v", errs)
	}

	errorCases := []struct {
		name          string
		envs          []api.EnvVar
		expectedError string
	}{
		{
			name:          "zero-length name",
			envs:          []api.EnvVar{{Name: ""}},
			expectedError: "[0].name: Required value",
		},
		{
			name:          "name not a C identifier",
			envs:          []api.EnvVar{{Name: "a.b.c"}},
			expectedError: `[0].name: Invalid value: "a.b.c": must match the regex`,
		},
	}
	for _, tc := range errorCases {
		if errs := validateEnv(tc.envs, field.NewPath("field")); len(errs) == 0 {
			t.Errorf("expected failure for %s", tc.name)
		} else {
			for i := range errs {
				str := errs[i].Error()
				if str != "" && !strings.Contains(str, tc.expectedError) {
					t.Errorf("%s: expected error detail either empty or %q, got %q", tc.name, tc.expectedError, str)
				}
			}
		}
	}
}

func TestValidateResourceParses(t *testing.T) {
	successCases := []api.ResourceSpec{
		{CPU: "1", GPU: "1", Memory: "512Mi"},
		{VolumeSize: "200Gi", Memory: "", CPU: "500m"},
		{CPU: "1.5", GPU: "1", Memory: "512Mi"},
	}
	for _, spec := range successCases {
		if errs := validateResourceSpec(&spec, field.NewPath("field")); len(errs) != 0 {
			t.Errorf("expected success: %v", errs)
		}
	}
	errorCases := []api.ResourceSpec{
		{CPU: "1aa", GPU: "1", Memory: "512Mi"},
		{VolumeSize: "200GiB"},
	}
	for k, spec := range errorCases {
		if errs := validateResourceSpec(&spec, field.NewPath("field")); len(errs) == 0 {
			t.Errorf("expected failure for %d", k)
		}
	}
}

func TestValidateRestartPolicy(t *testing.T) {
	successCases := []api.RestartPolicy{
		api.RestartPolicyAlways,
		api.RestartPolicyOnFailure,
		api.RestartPolicyNever,
	}
	for _, policy := range successCases {
		if errs := validateRestartPolicy(&policy, field.NewPath("field")); len(errs) != 0 {
			t.Errorf("expected success: %v", errs)
		}
	}

	errorCases := []api.RestartPolicy{"", "newpolicy"}

	for k, policy := range errorCases {
		if errs := validateRestartPolicy(&policy, field.NewPath("field")); len(errs) == 0 {
			t.Errorf("expected failure for %d", k)
		}
	}
}

func TestValidateSpotPolicy(t *testing.T) {
	successCases := []api.SpotPolicy{
		api.SpotAlways,
		api.SpotNever,
	}
	for _, policy := range successCases {
		if errs := validateSpotPolicy(&policy, field.NewPath("field")); len(errs) != 0 {
			t.Errorf("expected success: %v", errs)
		}
	}

	errorCases := []api.SpotPolicy{"", "notaspotpolicy"}

	for k, policy := range errorCases {
		if errs := validateSpotPolicy(&policy, field.NewPath("field")); len(errs) == 0 {
			t.Errorf("expected failure for %d", k)
		}
	}
}

func TestValidateUnits(t *testing.T) {
	successCase := []api.Unit{
		{Name: "abc", Image: "image"},
		{Name: "123", Image: "image"},
		{Name: "abc-123", Image: "image"},
		{Name: "def", Image: "image", Ports: []api.ServicePort{
			{
				Protocol:      api.ProtocolTCP,
				Port:          80,
				PortRangeSize: 1,
			},
			{
				Protocol:      api.ProtocolUDP,
				Port:          80,
				PortRangeSize: 1,
			},
		}},
	}
	if errs := validateUnits(successCase, sets.NewString(), field.NewPath("field")); len(errs) != 0 {
		t.Errorf("expected success: %v", errs)
	}

	errorCases := map[string][]api.Unit{
		"zero-length name":     {{Name: "", Image: "image"}},
		"name > 63 characters": {{Name: strings.Repeat("a", 64), Image: "image"}},
		"name not a DNS label": {{Name: "a.b.c", Image: "image"}},
		"name not unique": {
			{Name: "abc", Image: "image"},
			{Name: "abc", Image: "image"},
		},
		"zero-length image": {{Name: "abc", Image: ""}},
		"Non-unique ports": {
			{Name: "abc", Image: "image", Ports: []api.ServicePort{{
				Protocol:      api.ProtocolTCP,
				Port:          80,
				PortRangeSize: 1,
			}}},
			{Name: "def", Image: "image", Ports: []api.ServicePort{{
				Protocol:      api.ProtocolTCP,
				Port:          80,
				PortRangeSize: 1,
			}}},
		},
	}

	for k, v := range errorCases {
		if errs := validateUnits(v, sets.NewString(), field.NewPath("field")); len(errs) == 0 {
			t.Errorf("expected failure for %s", k)
		}
	}
}

func TestValidateInitUnitNames(t *testing.T) {
	initUnits := []api.Unit{
		{Name: "duplicate", Image: "image"},
	}
	units := []api.Unit{
		{Name: "duplicate", Image: "image"},
	}
	vols := sets.NewString()
	path := field.NewPath("field")
	if errs := validateInitUnits(initUnits, units, vols, path); len(errs) == 0 {
		t.Errorf("expected duplicate unit name failure")
	}
	units[0].Name = "not-duplicate"
	if errs := validateInitUnits(initUnits, units, vols, path); len(errs) > 0 {
		t.Errorf("expected no errors: %v", errs)
	}
}

func makeValidService() api.Service {
	return api.Service{
		ObjectMeta: api.ObjectMeta{
			Name:   "valid",
			Labels: map[string]string{},
		},
		Spec: api.ServiceSpec{
			Selector: api.LabelSelector{
				MatchLabels: map[string]string{"key": "val"},
			},
			Ports: []api.ServicePort{{
				Name:          "p",
				Protocol:      "TCP",
				Port:          8675,
				PortRangeSize: 1,
			}},
			SourceRanges: []string{"100.64.0.0/16"},
		},
	}
}

func TestValidateService(t *testing.T) {
	testCases := []struct {
		name     string
		tweakSvc func(svc *api.Service) // given a basic valid service, each test case can customize it
		numErrs  int
	}{
		{
			name: "missing name",
			tweakSvc: func(s *api.Service) {
				s.Name = ""
			},
			numErrs: 1,
		},
		{
			name: "invalid name",
			tweakSvc: func(s *api.Service) {
				s.Name = "-123"
			},
			numErrs: 1,
		},
		{
			name: "too long name",
			tweakSvc: func(s *api.Service) {
				s.Name = strings.Repeat("a", 25)
			},
			numErrs: 1,
		},
		{
			name: "invalid label",
			tweakSvc: func(s *api.Service) {
				s.Labels["NoUppercaseOrSpecialCharsLike=Equals"] = "bar"
			},
			numErrs: 1,
		},
		{
			name: "nil selector",
			tweakSvc: func(s *api.Service) {
				s.Spec.Selector = api.LabelSelector{}
			},
			numErrs: 0,
		},
		{
			name: "invalid selector",
			tweakSvc: func(s *api.Service) {
				s.Spec.Selector = api.LabelSelector{
					MatchLabels: map[string]string{
						"NoSpecialCharsLike=Equals": "bar",
					}}
			},
			numErrs: 2,
		},
		{
			name: "missing ports",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports = nil
			},
			numErrs: 1,
		},
		{
			name: "empty port[0] name",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Name = ""
			},
			numErrs: 0,
		},
		{
			name: "empty port[1] name",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{
					Name:          "",
					Protocol:      "TCP",
					Port:          12345,
					PortRangeSize: 1,
				})
			},
			numErrs: 1,
		},
		{
			name: "empty multi-port port[0] name",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Name = ""
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{
					Name:          "p",
					Protocol:      "TCP",
					Port:          12345,
					PortRangeSize: 1,
				})
			},
			numErrs: 1,
		},
		{
			name: "invalid port name",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Name = "INVALID"
			},
			numErrs: 1,
		},
		{
			name: "missing protocol",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Protocol = ""
			},
			numErrs: 1,
		},
		{
			name: "invalid protocol",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Protocol = "INVALID"
			},
			numErrs: 1,
		},
		{
			name: "missing port",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Port = 0
			},
			numErrs: 1,
		},
		{
			name: "invalid port",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Port = 65536
			},
			numErrs: 1,
		},
		{
			name: "milpa port",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Port = RestAPIPort
			},
			numErrs: 1,
		},
		{
			name: "dup port name",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Name = "p"
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{
					Name:          "p",
					Port:          12345,
					Protocol:      "TCP",
					PortRangeSize: 1,
				})
			},
			numErrs: 1,
		},
		{
			name: "valid 1",
			tweakSvc: func(s *api.Service) {
				// do nothing
			},
			numErrs: 0,
		},
		{
			name: "valid 2",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports[0].Protocol = "UDP"
			},
			numErrs: 0,
		},
		{
			name: "duplicate service ports for same protocol",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{Name: "q", Port: 1, Protocol: "TCP", PortRangeSize: 1})
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{Name: "r", Port: 1, Protocol: "TCP", PortRangeSize: 1})
			},
			numErrs: 2,
		},
		{
			name: "overlapping service ports for same protocol",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{Name: "q", Port: 5, Protocol: "TCP", PortRangeSize: 10})
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{Name: "r", Port: 10, Protocol: "TCP", PortRangeSize: 1})
			},
			numErrs: 1,
		},
		{
			name: "duplicate ports (different protocols)",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{Name: "q", Port: 1, Protocol: "TCP", PortRangeSize: 1})
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{Name: "r", Port: 1, Protocol: "UDP", PortRangeSize: 1})
			},
			numErrs: 0,
		},
		{
			name: "overlapping ports (different protocols)",
			tweakSvc: func(s *api.Service) {
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{Name: "q", Port: 1, Protocol: "TCP", PortRangeSize: 10})
				s.Spec.Ports = append(s.Spec.Ports, api.ServicePort{Name: "r", Port: 5, Protocol: "UDP", PortRangeSize: 20})
			},
			numErrs: 0,
		},
		{
			name: "No source range is OK",
			tweakSvc: func(s *api.Service) {
				s.Spec.SourceRanges = []string{}
			},
			numErrs: 0,
		},
		{
			name: "invalid cidr errors",
			tweakSvc: func(s *api.Service) {
				s.Spec.SourceRanges = []string{"1.2.3.4/55"}
			},
			numErrs: 1,
		},
		{
			name: "VPC Source range OK",
			tweakSvc: func(s *api.Service) {
				s.Spec.SourceRanges = []string{"VPC"}
			},
			numErrs: 0,
		},
		{
			name: "lower case vpc Source range errors",
			tweakSvc: func(s *api.Service) {
				s.Spec.SourceRanges = []string{"vpc"}
			},
			numErrs: 1,
		},
		{
			name: "invalid cidr errors",
			tweakSvc: func(s *api.Service) {
				s.Spec.SourceRanges = []string{"1.2.3.4"}
			},
			numErrs: 1,
		},
		{
			name: "duplicate CIDRs fail",
			tweakSvc: func(s *api.Service) {
				s.Spec.SourceRanges = []string{"1.2.4.0/16", "1.2.3.0/16"}
			},
			numErrs: 1,
		},
		{
			name: "servicec is named milpa",
			tweakSvc: func(s *api.Service) {
				s.Name = util.MilpaSvcName
			},
			numErrs: 1,
		},
	}

	for _, tc := range testCases {
		svc := makeValidService()
		tc.tweakSvc(&svc)
		errs := ValidateService(&svc)
		if len(errs) != tc.numErrs {
			t.Errorf("Unexpected error list for case %q: %v", tc.name, errs.ToAggregate())
			fmt.Printf("Expected %d errors, got %d errors\n", tc.numErrs, len(errs))
		}
	}
}

func TestValidateLabels(t *testing.T) {
	successCases := []map[string]string{
		{"simple": "bar"},
		{"now-with-dashes": "bar"},
		{"1-starts-with-num": "bar"},
		{"1234": "bar"},
		{"simple/simple": "bar"},
		{"now-with-dashes/simple": "bar"},
		{"now-with-dashes/now-with-dashes": "bar"},
		{"now.with.dots/simple": "bar"},
		{"now-with.dashes-and.dots/simple": "bar"},
		{"1-num.2-num/3-num": "bar"},
		{"1234/5678": "bar"},
		{"1.2.3.4/5678": "bar"},
		{"UpperCaseAreOK123": "bar"},
		{"goodvalue": "123_-.BaR"},
	}
	for i := range successCases {
		errs := ValidateLabels(successCases[i], field.NewPath("field"))
		if len(errs) != 0 {
			t.Errorf("case[%d] expected success, got %#v", i, errs)
		}
	}

	labelNameErrorCases := []struct {
		labels map[string]string
		expect string
	}{
		{map[string]string{"nospecialchars^=@": "bar"}, "must match the regex"},
		{map[string]string{"cantendwithadash-": "bar"}, "must match the regex"},
		{map[string]string{"only/one/slash": "bar"}, "must match the regex"},
		{map[string]string{strings.Repeat("a", 254): "bar"}, "must be no more than"},
	}
	for i := range labelNameErrorCases {
		errs := ValidateLabels(labelNameErrorCases[i].labels, field.NewPath("field"))
		if len(errs) != 1 {
			t.Errorf("case[%d]: expected failure", i)
		} else {
			if !strings.Contains(errs[0].Detail, labelNameErrorCases[i].expect) {
				t.Errorf("case[%d]: error details do not include %q: %q", i, labelNameErrorCases[i].expect, errs[0].Detail)
			}
		}
	}

	labelValueErrorCases := []struct {
		labels map[string]string
		expect string
	}{
		{map[string]string{"toolongvalue": strings.Repeat("a", 64)}, "must be no more than"},
		{map[string]string{"backslashesinvalue": "some\\bad\\value"}, "must match the regex"},
		{map[string]string{"nocommasallowed": "bad,value"}, "must match the regex"},
		{map[string]string{"strangecharsinvalue": "?#$notsogood"}, "must match the regex"},
	}
	for i := range labelValueErrorCases {
		errs := ValidateLabels(labelValueErrorCases[i].labels, field.NewPath("field"))
		if len(errs) != 1 {
			t.Errorf("case[%d]: expected failure", i)
		} else {
			if !strings.Contains(errs[0].Detail, labelValueErrorCases[i].expect) {
				t.Errorf("case[%d]: error details do not include %q: %q", i, labelValueErrorCases[i].expect, errs[0].Detail)
			}
		}
	}
}

// Screw it, I took k8s code, I'm taking their damn tests too...
func TestValidateVolumes(t *testing.T) {
	successCase := []api.Volume{
		{Name: "empty", VolumeSource: api.VolumeSource{EmptyDir: &api.EmptyDir{}}},
	}
	names, errs := validateVolumes(successCase, field.NewPath("field"))
	if len(errs) != 0 {
		t.Errorf("expected success: %v", errs)
	}
	if len(names) != len(successCase) || !names.HasAll("empty") {
		t.Errorf("wrong names result: %v", names)
	}

	emptyVS := api.VolumeSource{EmptyDir: &api.EmptyDir{}}
	errorCases := map[string]struct {
		V []api.Volume
		T field.ErrorType
		F string
		D string
	}{
		"zero-length name": {
			[]api.Volume{{Name: "", VolumeSource: emptyVS}},
			field.ErrorTypeRequired,
			"name", "",
		},
		"name > 63 characters": {
			[]api.Volume{{Name: strings.Repeat("a", 64), VolumeSource: emptyVS}},
			field.ErrorTypeInvalid,
			"name", "must be no more than",
		},
		"name not a DNS label": {
			[]api.Volume{{Name: "a.b.c", VolumeSource: emptyVS}},
			field.ErrorTypeInvalid,
			"name", "must match the regex",
		},
		"name not unique": {
			[]api.Volume{{Name: "abc", VolumeSource: emptyVS}, {Name: "abc", VolumeSource: emptyVS}},
			field.ErrorTypeDuplicate,
			"[1].name", "",
		},
		"empty PackagePath": {
			[]api.Volume{{Name: "abc", VolumeSource: api.VolumeSource{PackagePath: &api.PackagePath{Path: ""}}}},
			field.ErrorTypeRequired,
			"[0].packagePath.path", "",
		},
	}
	for k, v := range errorCases {
		_, errs := validateVolumes(v.V, field.NewPath("field"))
		if len(errs) == 0 {
			t.Errorf("expected failure %s for %v", k, v.V)
			continue
		}
		for i := range errs {
			if errs[i].Type != v.T {
				t.Errorf("%s: expected error to have type %q: %q", k, v.T, errs[i].Type)
			}
			if !strings.Contains(errs[i].Field, v.F) {
				t.Errorf("%s: expected error field %q: %q", k, v.F, errs[i].Field)
			}
			if !strings.Contains(errs[i].Detail, v.D) {
				t.Errorf("%s: expected error detail %q, got %q", k, v.D, errs[i].Detail)
			}
		}
	}
}

func TestValidateVolumeMounts(t *testing.T) {
	volumes := sets.NewString("abc", "123", "abc-123")

	successCase := []api.VolumeMount{
		{Name: "abc", MountPath: "/foo"},
		{Name: "123", MountPath: "/bar"},
		{Name: "abc-123", MountPath: "/baz"},
	}
	if errs := validateVolumeMounts(successCase, volumes, field.NewPath("field")); len(errs) != 0 {
		t.Errorf("expected success: %v", errs)
	}

	errorCases := map[string][]api.VolumeMount{
		"empty name":          {{Name: "", MountPath: "/foo"}},
		"name not found":      {{Name: "", MountPath: "/foo"}},
		"empty mountpath":     {{Name: "abc", MountPath: ""}},
		"colon mountpath":     {{Name: "abc", MountPath: "foo:bar"}},
		"mountpath collision": {{Name: "foo", MountPath: "/path/a"}, {Name: "bar", MountPath: "/path/a"}},
	}
	for k, v := range errorCases {
		if errs := validateVolumeMounts(v, volumes, field.NewPath("field")); len(errs) == 0 {
			t.Errorf("expected failure for %s", k)
		}
	}
}

func mkmultiports(vals ...int) []api.ServicePort {
	ports := []api.ServicePort{}
	for i := 0; i < len(vals); i += 2 {
		ports = append(ports, api.ServicePort{
			Port:          vals[i],
			PortRangeSize: vals[i+1],
		})
	}
	return ports
}

func TestValidateServicePort(t *testing.T) {
	tests := []struct {
		port      api.ServicePort
		numErrors int
	}{
		{
			port: api.ServicePort{
				Port:     -1,
				Protocol: api.ProtocolICMP,
			},
			numErrors: 0,
		},
		{
			port: api.ServicePort{
				Port:          -1,
				Protocol:      api.ProtocolICMP,
				PortRangeSize: 258,
			},
			numErrors: 1,
		},
	}
	for i, test := range tests {
		errs := ValidateServicePort(&test.port, false, field.NewPath(""))
		if len(errs) != test.numErrors {
			t.Errorf("Expected %d errors but test %d produced %d errors: %v", test.numErrors, i, len(errs), errs)
		}
	}
}

func TestValidateServicePortRangesOverlap(t *testing.T) {
	tests := []struct {
		ports []api.ServicePort
		err   bool
	}{
		{
			ports: mkmultiports(
				100, 50,
				150, 50,
				200, 1,
			),
			err: false,
		},
		{
			ports: mkmultiports(
				100, 51,
				150, 50,
				200, 1,
			),
			err: true,
		},
		{
			ports: mkmultiports(
				100, 50,
				150, 51,
				200, 1,
			),
			err: true,
		},
		{
			ports: mkmultiports(
				150, 50,
				100, 51,
			),
			err: true,
		},
	}
	for i, test := range tests {
		errs := validateServicePortOverlap(test.ports, field.NewPath(""))
		if len(errs) > 0 && !test.err {
			t.Errorf("Expected no overlap but test %d produced overlap %v", i, errs)
		} else if len(errs) == 0 && test.err {
			t.Errorf("Expected overlap but test %d produced no overlap %v", i, errs)
		}
	}
}

func TestValidatePodSecurityContext(t *testing.T) {
	tests := []struct {
		context *api.PodSecurityContext
		errlen  int
	}{
		{
			nil,
			0,
		},
		{
			&api.PodSecurityContext{
				Sysctls: []api.Sysctl{
					{
						Name:  "sysctl.name.1",
						Value: "sysctl.value.1",
					},
					{
						Name:  "sysctl.name.2",
						Value: "sysctl.value.2",
					},
				},
			},
			0,
		},
		{
			&api.PodSecurityContext{
				Sysctls: []api.Sysctl{
					{
						Name: "sysctl.name.1",
					},
				},
			},
			1,
		},
		{
			&api.PodSecurityContext{
				Sysctls: []api.Sysctl{
					{
						Value: "sysctl.value.1",
					},
				},
			},
			1,
		},
		{
			&api.PodSecurityContext{
				Sysctls: []api.Sysctl{
					{
						Name: "sysctl.name.1",
					},
					{
						Value: "sysctl.value.2",
					},
				},
			},
			2,
		},
		{
			&api.PodSecurityContext{
				Sysctls: []api.Sysctl{
					{},
				},
			},
			1,
		},
	}
	for _, tc := range tests {
		errs := validatePodSecurityContext(tc.context, field.NewPath("Sysctls"))
		if len(errs) != tc.errlen {
			t.Errorf("Expected %d errors validating PodSecurityContext %v, but got %d: %v",
				tc.errlen, tc.context, len(errs), errs)
		}
	}
}
