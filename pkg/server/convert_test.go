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

package server

import (
	"fmt"
	"math"
	"testing"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/kip/pkg/util/rand"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func fakeInstanceProvider() (string, string) {
	ipStr := fmt.Sprintf(
		"%d.%d.%d.%d",
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256))
	return rand.String(8), ipStr
}

//func getStatus(milpaPod *api.Pod, pod *v1.Pod) v1.PodStatus
func TestGetStatus(t *testing.T) {
	_, ip := fakeInstanceProvider()
	pod := &v1.Pod{}
	testCases := []struct {
		milpaPodPhase api.PodPhase
		k8sPodPhase   v1.PodPhase
		modPod        func(*api.Pod)
	}{
		{
			milpaPodPhase: api.PodDispatching,
			k8sPodPhase:   v1.PodPending,
		},
		{
			milpaPodPhase: api.PodFailed,
			k8sPodPhase:   v1.PodPending,
		},
		{
			milpaPodPhase: api.PodFailed,
			k8sPodPhase:   v1.PodFailed,
			modPod: func(p *api.Pod) {
				p.Spec.RestartPolicy = api.RestartPolicyNever
			},
		},
		{
			milpaPodPhase: api.PodFailed,
			k8sPodPhase:   v1.PodFailed,
			modPod: func(p *api.Pod) {
				p.Status.StartFailures = allowedStartFailures + 1
			},
		},
		{
			milpaPodPhase: api.PodRunning,
			k8sPodPhase:   v1.PodRunning,
		},
		{
			milpaPodPhase: api.PodSucceeded,
			k8sPodPhase:   v1.PodSucceeded,
		},
		{
			milpaPodPhase: api.PodTerminated,
			k8sPodPhase:   v1.PodFailed,
		},
		{
			milpaPodPhase: api.PodWaiting,
			k8sPodPhase:   v1.PodPending,
		},
	}
	for _, tc := range testCases {
		milpaPod := api.GetFakePod()
		milpaPod.Status.Phase = tc.milpaPodPhase
		if tc.modPod != nil {
			tc.modPod(milpaPod)
		}
		podStatus := getStatus(ip, milpaPod, pod)
		assert.Equal(t, podStatus.Phase, tc.k8sPodPhase)
	}
}

//func unitToContainerStatus(st api.UnitStatus) v1.ContainerStatus
func TestUnitToContainerStatus(t *testing.T) {
	testCases := []struct {
		unitState api.UnitState
	}{
		{
			unitState: api.UnitState{
				Waiting: &api.UnitStateWaiting{
					Reason: "waiting to start",
				},
			},
		},
		{
			unitState: api.UnitState{
				Running: &api.UnitStateRunning{
					StartedAt: api.Now(),
				},
			},
		},
		{
			unitState: api.UnitState{
				Terminated: &api.UnitStateTerminated{
					ExitCode:   int32(rand.Intn(256)),
					FinishedAt: api.Now(),
				},
			},
		},
	}
	for _, tc := range testCases {
		us := api.UnitStatus{
			Name:         "myunit",
			RestartCount: 0,
			Image:        "elotl/myimage",
			State:        tc.unitState,
		}
		cs := unitToContainerStatus(us)
		if us.State.Waiting != nil {
			assert.NotNil(t, cs.State.Waiting)
			assert.Nil(t, cs.State.Running)
			assert.Nil(t, cs.State.Terminated)
			assert.Equal(t, us.State.Waiting.Reason, cs.State.Waiting.Reason)
		}
		if us.State.Running != nil {
			assert.NotNil(t, cs.State.Running)
			assert.Nil(t, cs.State.Waiting)
			assert.Nil(t, cs.State.Terminated)
			assert.Equal(
				t,
				us.State.Running.StartedAt.Time,
				cs.State.Running.StartedAt.Time)
		}
		if us.State.Terminated != nil {
			assert.NotNil(t, cs.State.Terminated)
			assert.Nil(t, cs.State.Running)
			assert.Nil(t, cs.State.Waiting)
			assert.Equal(
				t,
				us.State.Terminated.ExitCode,
				cs.State.Terminated.ExitCode)
			assert.Equal(
				t,
				us.State.Terminated.FinishedAt.Time,
				cs.State.Terminated.FinishedAt.Time)
		}
	}
}

//func containerToUnit(container v1.Container) api.Unit
//func unitToContainer(unit api.Unit, container *v1.Container) v1.Container
func TestUnitToContainer(t *testing.T) {
	user := int64(rand.Intn(65536))
	group := int64(rand.Intn(65536))
	testCases := []struct {
		unit api.Unit
	}{
		{
			unit: api.Unit{
				Name:  rand.String(16),
				Image: fmt.Sprintf("elotl/%s:latest", rand.String(8)),
			},
		},
		{
			unit: api.Unit{
				Name:  rand.String(16),
				Image: fmt.Sprintf("elotl/%s:latest", rand.String(8)),
				Command: []string{
					"/bin/bash",
				},
				Args: []string{
					"-c",
					"sleep 1",
				},
			},
		},
		{
			unit: api.Unit{
				Name:  rand.String(16),
				Image: fmt.Sprintf("elotl/%s:latest", rand.String(8)),
				Env: []api.EnvVar{
					{
						Name:  "env1",
						Value: "value1",
					},
					{
						Name:  "foo",
						Value: "bar",
					},
				},
			},
		},
		{
			unit: api.Unit{
				Name:  rand.String(16),
				Image: fmt.Sprintf("elotl/%s:latest", rand.String(8)),
				VolumeMounts: []api.VolumeMount{
					{
						Name:      "myvolume",
						MountPath: "/my/path",
					},
					{
						Name:      "foo-volume",
						MountPath: "/var/run/bar",
					},
				},
			},
		},
		{
			unit: api.Unit{
				Name:  rand.String(16),
				Image: fmt.Sprintf("elotl/%s:latest", rand.String(8)),
				Ports: []api.ContainerPort{
					{
						Name:          "my-tcp-port",
						Protocol:      api.ProtocolTCP,
						ContainerPort: 80,
						HostPort:      8880,
					},
					{
						Name:          "my-udp-port",
						Protocol:      api.ProtocolUDP,
						ContainerPort: 53,
						HostPort:      5353,
					},
				},
			},
		},
		{
			unit: api.Unit{
				Name:       rand.String(16),
				Image:      fmt.Sprintf("elotl/%s:latest", rand.String(8)),
				WorkingDir: "/home/nobody",
			},
		},
		{
			unit: api.Unit{
				Name:  rand.String(16),
				Image: fmt.Sprintf("elotl/%s:latest", rand.String(8)),
				SecurityContext: &api.SecurityContext{
					RunAsUser:  &user,
					RunAsGroup: &group,
					Capabilities: &api.Capabilities{
						Add: []string{
							"add-cap-1",
							"add-cap-2",
						},
						Drop: []string{
							"drop-cap-1",
							"drop-cap-2",
							"drop-cap-3",
						},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		container := unitToContainer(tc.unit, nil)
		assert.Equal(t, tc.unit.Name, container.Name)
		assert.Equal(t, tc.unit.Image, container.Image)
		assert.Equal(t, tc.unit.WorkingDir, container.WorkingDir)
		assert.ElementsMatch(t, tc.unit.Command, container.Command)
		assert.ElementsMatch(t, tc.unit.Args, container.Args)
		assert.Equal(t, len(tc.unit.Env), len(container.Env))
		for _, e := range tc.unit.Env {
			env := v1.EnvVar{
				Name:  e.Name,
				Value: e.Value,
			}
			assert.Contains(t, container.Env, env)
		}
		assert.Equal(t, len(tc.unit.VolumeMounts), len(container.VolumeMounts))
		for _, v := range tc.unit.VolumeMounts {
			vol := v1.VolumeMount{
				Name:      v.Name,
				MountPath: v.MountPath,
			}
			assert.Contains(t, container.VolumeMounts, vol)
		}
		for _, p := range tc.unit.Ports {
			port := v1.ContainerPort{
				Name:          p.Name,
				Protocol:      v1.Protocol(string(p.Protocol)),
				HostPort:      int32(p.HostPort),
				ContainerPort: int32(p.ContainerPort),
			}
			assert.Contains(t, container.Ports, port)
		}
		if tc.unit.SecurityContext != nil {
			assert.Equal(
				t,
				tc.unit.SecurityContext.RunAsUser,
				container.SecurityContext.RunAsUser)
			assert.Equal(
				t,
				tc.unit.SecurityContext.RunAsGroup,
				container.SecurityContext.RunAsGroup)
			if tc.unit.SecurityContext.Capabilities != nil {
				assert.NotNil(t, container.SecurityContext.Capabilities)
				assert.Equal(
					t,
					len(tc.unit.SecurityContext.Capabilities.Add),
					len(container.SecurityContext.Capabilities.Add))
				for _, a := range tc.unit.SecurityContext.Capabilities.Add {
					assert.Contains(
						t,
						container.SecurityContext.Capabilities.Add,
						v1.Capability(a))
				}
				assert.Equal(
					t,
					len(tc.unit.SecurityContext.Capabilities.Drop),
					len(container.SecurityContext.Capabilities.Drop))
				for _, d := range tc.unit.SecurityContext.Capabilities.Drop {
					assert.Contains(
						t,
						container.SecurityContext.Capabilities.Drop,
						v1.Capability(d))
				}
			}
		}
		unit := containerToUnit(container)
		removeVolumeMount(&unit, resolvconfVolumeName)
		removeVolumeMount(&unit, etchostsVolumeName)
		assert.Equal(t, tc.unit, unit)
	}
}

func removeVolume(pod *api.Pod, name string) {
	idx := -1
	for i, vol := range pod.Spec.Volumes {
		if vol.Name == name {
			idx = i
			break
		}
	}
	if idx != -1 {
		pod.Spec.Volumes = append(
			pod.Spec.Volumes[:idx], pod.Spec.Volumes[idx+1:]...)
		if len(pod.Spec.Volumes) == 0 {
			pod.Spec.Volumes = nil
		}
	}
	for i := range pod.Spec.InitUnits {
		removeVolumeMount(&pod.Spec.InitUnits[i], name)
	}
	for i := range pod.Spec.Units {
		removeVolumeMount(&pod.Spec.Units[i], name)
	}
}

func removeVolumeMount(unit *api.Unit, name string) {
	idx := -1
	for i, vol := range unit.VolumeMounts {
		if vol.Name == name {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}
	if len(unit.VolumeMounts) == 1 {
		unit.VolumeMounts = nil
		return
	}
	unit.VolumeMounts = append(
		unit.VolumeMounts[:idx], unit.VolumeMounts[idx+1:]...)
}

func hostPathTypePtr(hpt api.HostPathType) *api.HostPathType {
	return &hpt
}

//func k8sToMilpaVolume(vol v1.Volume) *api.Volume
//func milpaToK8sVolume(vol api.Volume) *v1.Volume
func TestMilpaToK8sVolume(t *testing.T) {
	i32 := int32(rand.Intn(math.MaxInt32))
	bTrue := true
	testCases := []struct {
		volume api.Volume
	}{
		{
			volume: api.Volume{
				Name: rand.String(16),
				VolumeSource: api.VolumeSource{
					EmptyDir: &api.EmptyDir{
						Medium:    api.StorageMediumMemory,
						SizeLimit: int64(rand.Intn(math.MaxInt64)),
					},
				},
			},
		},
		{
			volume: api.Volume{
				Name: rand.String(16),
				VolumeSource: api.VolumeSource{
					ConfigMap: &api.ConfigMapVolumeSource{
						Items: []api.KeyToPath{
							{
								Key:  rand.String(8),
								Path: fmt.Sprintf("/var/run/%s", rand.String(8)),
								Mode: &i32,
							},
						},
						DefaultMode: &i32,
						Optional:    &bTrue,
					},
				},
			},
		},
		{
			volume: api.Volume{
				Name: rand.String(16),
				VolumeSource: api.VolumeSource{
					Secret: &api.SecretVolumeSource{
						SecretName: rand.String(16),
						Items: []api.KeyToPath{
							{
								Key:  rand.String(8),
								Path: fmt.Sprintf("/var/run/%s", rand.String(8)),
								Mode: &i32,
							},
						},
						DefaultMode: &i32,
						Optional:    &bTrue,
					},
				},
			},
		},
		{
			volume: api.Volume{
				Name: rand.String(16),
				VolumeSource: api.VolumeSource{
					HostPath: &api.HostPathVolumeSource{
						Path: rand.String(16),
						Type: hostPathTypePtr(api.HostPathDirectoryOrCreate),
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		vol := milpaToK8sVolume(tc.volume)
		assert.NotNil(t, vol)
		assert.Equal(t, tc.volume.Name, vol.Name)
		if tc.volume.EmptyDir != nil {
			assert.NotNil(t, vol.EmptyDir)
			assert.Equal(
				t,
				string(tc.volume.EmptyDir.Medium),
				string(vol.EmptyDir.Medium),
			)
			assert.Equal(
				t,
				tc.volume.EmptyDir.SizeLimit,
				vol.EmptyDir.SizeLimit.Value(),
			)
		}
		if tc.volume.ConfigMap != nil {
			assert.NotNil(t, vol.ConfigMap)
			assert.Equal(
				t,
				tc.volume.ConfigMap.Name,
				vol.ConfigMap.Name,
			)
			assert.Equal(
				t,
				tc.volume.ConfigMap.DefaultMode,
				vol.ConfigMap.DefaultMode,
			)
			assert.Equal(
				t,
				tc.volume.ConfigMap.Optional,
				vol.ConfigMap.Optional,
			)
			assert.Equal(
				t,
				len(tc.volume.ConfigMap.Items),
				len(vol.ConfigMap.Items),
			)
			for _, item := range tc.volume.ConfigMap.Items {
				ktp := v1.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: item.Mode,
				}
				assert.Contains(t, vol.ConfigMap.Items, ktp)
			}
		}
		if tc.volume.Secret != nil {
			assert.NotNil(t, vol.Secret)
			assert.Equal(t, tc.volume.Secret.SecretName, vol.Secret.SecretName)
			assert.Equal(
				t,
				tc.volume.Secret.DefaultMode,
				vol.Secret.DefaultMode,
			)
			assert.Equal(
				t,
				tc.volume.Secret.Optional,
				vol.Secret.Optional,
			)
			assert.Equal(
				t,
				len(tc.volume.Secret.Items),
				len(vol.Secret.Items),
			)
			for _, item := range tc.volume.Secret.Items {
				ktp := v1.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: item.Mode,
				}
				assert.Contains(t, vol.Secret.Items, ktp)
			}
		}
		if tc.volume.HostPath != nil {
			assert.NotNil(t, vol.HostPath)
			assert.Equal(t, tc.volume.HostPath.Path, vol.HostPath.Path)
			if tc.volume.HostPath.Type != nil {
				assert.NotNil(t, vol.HostPath.Type)
				assert.Equal(
					t,
					string(*tc.volume.HostPath.Type),
					string(*vol.HostPath.Type))
			} else {
				assert.Nil(t, vol.HostPath.Type)
			}
		}
		milpaVolume := k8sToMilpaVolume(*vol)
		assert.Equal(t, tc.volume, *milpaVolume)
	}
}

func TestProjectedVolumeConversion(t *testing.T) {
	falseVal := false
	volMode := int32(0555)
	secMode := int32(0400)
	cmMode := int32(0400)
	tests := []struct {
		sources []api.VolumeProjection
	}{
		{
			sources: []api.VolumeProjection{
				{
					Secret: &api.SecretProjection{
						LocalObjectReference: api.LocalObjectReference{"sec1"},
						Items: []api.KeyToPath{
							{
								Key:  "sec1Key1",
								Path: "sec1/Key1",
								Mode: &secMode,
							},
							{
								Key:  "sec1Key2",
								Path: "sec1/Key2",
								Mode: &secMode,
							},
						},
						Optional: &falseVal,
					},
				},
				{
					Secret: &api.SecretProjection{
						LocalObjectReference: api.LocalObjectReference{"sec2"},
						Items: []api.KeyToPath{
							{
								Key:  "sec2Key1",
								Path: "sec2/Key1",
								Mode: &secMode,
							},
							{
								Key:  "sec2Key2",
								Path: "sec2/Key2",
								Mode: &secMode,
							},
						},
						Optional: &falseVal,
					},
				},
				{
					ConfigMap: &api.ConfigMapProjection{
						LocalObjectReference: api.LocalObjectReference{"cm1"},
						Items: []api.KeyToPath{
							{
								Key:  "cm1Key1",
								Path: "cm1/Key1",
								Mode: &secMode,
							},
							{
								Key:  "cm1Key2",
								Path: "cm1/Key2",
								Mode: &cmMode,
							},
						},
						Optional: nil,
					},
				},
			},
		},
	}
	for i, tc := range tests {
		name := fmt.Sprintf("testcase %d", i)
		milpaVol := api.Volume{
			Name: name,
			VolumeSource: api.VolumeSource{
				Projected: &api.ProjectedVolumeSource{
					Sources:     tc.sources,
					DefaultMode: &volMode,
				},
			},
		}
		kv := milpaToK8sVolume(milpaVol)
		mv := k8sToMilpaVolume(*kv)
		assert.Equal(t, milpaVol, *mv)
	}
}

//func k8sToMilpaPod(pod *v1.Pod) (*api.Pod, error)
//func milpaToK8sPod(milpaPod *api.Pod) (*v1.Pod, error)
func TestMilpaToK8sPod(t *testing.T) {
	i64 := int64(rand.Intn(math.MaxInt64))
	node, ip := fakeInstanceProvider()
	milpaPod := api.NewPod()
	milpaPod.Namespace = rand.String(16)
	milpaPod.Name = util.WithNamespace(milpaPod.Namespace, rand.String(16))
	milpaPod.Spec = api.PodSpec{
		Spot: api.PodSpot{
			Policy: api.SpotNever,
		},
		Phase:         api.PodRunning,
		RestartPolicy: api.RestartPolicyOnFailure,
		Units: []api.Unit{
			{
				Name:  rand.String(8),
				Image: fmt.Sprintf("elotl/%s:latest", rand.String(8)),
				Command: []string{
					"unit-1-cmd",
				},
				Args: []string{
					"-a",
					"--bb",
					"ccc",
				},
			},
		},
		InitUnits: []api.Unit{
			{
				Name:  rand.String(8),
				Image: fmt.Sprintf("elotl/%s:latest", rand.String(8)),
				Command: []string{
					"initunit-1-cmd",
				},
				Args: []string{
					"--init",
				},
			},
		},
		Volumes: []api.Volume{
			{
				Name: rand.String(16),
				VolumeSource: api.VolumeSource{
					EmptyDir: &api.EmptyDir{
						Medium:    api.StorageMediumMemory,
						SizeLimit: int64(rand.Intn(math.MaxInt64)),
					},
				},
			},
		},
		SecurityContext: &api.PodSecurityContext{
			NamespaceOptions: &api.NamespaceOption{
				Network: api.NamespaceModeNode,
				Pid:     api.NamespaceModeNode,
				Ipc:     api.NamespaceModeContainer,
			},
			RunAsUser:  &i64,
			RunAsGroup: &i64,
			SupplementalGroups: []int64{
				int64(rand.Intn(math.MaxInt64)),
			},
			Sysctls: []api.Sysctl{
				{
					Name:  rand.String(16),
					Value: rand.String(16),
				},
				{
					Name:  rand.String(16),
					Value: rand.String(16),
				},
				{
					Name:  rand.String(16),
					Value: rand.String(16),
				},
			},
		},
	}
	pod := milpaToK8sPod(node, ip, milpaPod)
	assert.NotNil(t, pod)
	assert.Equal(t, len(milpaPod.Spec.Units), len(pod.Spec.Containers))
	for _, unit := range milpaPod.Spec.Units {
		container := unitToContainer(unit, nil)
		assert.Contains(t, pod.Spec.Containers, container)
	}
	assert.Equal(t, len(milpaPod.Spec.InitUnits), len(pod.Spec.InitContainers))
	for _, unit := range milpaPod.Spec.InitUnits {
		container := unitToContainer(unit, nil)
		assert.Contains(t, pod.Spec.InitContainers, container)
	}
	assert.Equal(t, len(milpaPod.Spec.Volumes), len(pod.Spec.Volumes))
	for _, vol := range milpaPod.Spec.Volumes {
		volume := milpaToK8sVolume(vol)
		assert.Contains(t, pod.Spec.Volumes, *volume)
	}
	assert.NotNil(t, pod.Spec.SecurityContext)
	assert.Equal(
		t,
		milpaPod.Spec.SecurityContext.RunAsUser,
		pod.Spec.SecurityContext.RunAsUser)
	assert.Equal(
		t,
		milpaPod.Spec.SecurityContext.RunAsGroup,
		pod.Spec.SecurityContext.RunAsGroup)
	assert.ElementsMatch(
		t,
		milpaPod.Spec.SecurityContext.SupplementalGroups,
		pod.Spec.SecurityContext.SupplementalGroups)
	assert.Equal(
		t,
		len(milpaPod.Spec.SecurityContext.Sysctls),
		len(pod.Spec.SecurityContext.Sysctls))
	for _, sysctl := range milpaPod.Spec.SecurityContext.Sysctls {
		sc := v1.Sysctl{
			Name:  sysctl.Name,
			Value: sysctl.Value,
		}
		assert.Contains(t, pod.Spec.SecurityContext.Sysctls, sc)
	}
	assert.Equal(t, len(milpaPod.Spec.Volumes), len(pod.Spec.Volumes))
	for _, volume := range milpaPod.Spec.Volumes {
		found := false
		for _, vol := range pod.Spec.Volumes {
			if vol.Name == volume.Name {
				found = true
				break
			}
		}
		assert.True(t, found, "volume %q is missing in k8s pod", volume.Name)
	}
	mPod := k8sToMilpaPod(pod)
	assert.NotNil(t, mPod)
	removeVolume(mPod, resolvconfVolumeName)
	removeVolume(mPod, etchostsVolumeName)
	assert.Equal(t, milpaPod.TypeMeta, mPod.TypeMeta)
	assert.Equal(t, milpaPod.ObjectMeta, mPod.ObjectMeta)
	assert.Equal(t, milpaPod.Spec, mPod.Spec)
}

func TestConvertingProbes(t *testing.T) {
	mp := &api.Probe{
		Handler: api.Handler{
			HTTPGet: &api.HTTPGetAction{
				Path:   "foo",
				Port:   intstr.FromInt(2),
				Host:   "localhost",
				Scheme: api.URISchemeHTTP,
				HTTPHeaders: []api.HTTPHeader{
					{
						Name:  "x-name",
						Value: "my value",
					},
				},
			},
		},
		InitialDelaySeconds: 1,
		TimeoutSeconds:      2,
		PeriodSeconds:       3,
		SuccessThreshold:    4,
		FailureThreshold:    5,
	}
	kp := milpaProbeToK8sProbe(mp)
	mp2 := k8sProbeToMilpaProbe(kp)
	assert.Equal(t, mp, mp2)
}

//func aggregateResources(spec v1.PodSpec) api.ResourceSpec
func TestAggregateResources(t *testing.T) {
	testCases := []struct {
		requirements []v1.ResourceRequirements
		resources    api.ResourceSpec
		nodeSelector map[string]string
	}{
		{
			requirements: []v1.ResourceRequirements{},
			resources:    api.ResourceSpec{},
		},
		{
			requirements: []v1.ResourceRequirements{
				{
					Limits: v1.ResourceList{
						v1.ResourceCPU: resource.MustParse("500m"),
					},
				},
			},
			resources: api.ResourceSpec{
				CPU: "0.50",
			},
		},
		{
			requirements: []v1.ResourceRequirements{
				{
					Requests: v1.ResourceList{
						v1.ResourceMemory: resource.MustParse("512Mi"),
					},
				},
			},
			resources: api.ResourceSpec{
				Memory: "0.50Gi",
			},
		},
		{
			requirements: []v1.ResourceRequirements{
				{
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("500m"),
						v1.ResourceMemory: resource.MustParse("1Gi"),
					},
				},
			},
			resources: api.ResourceSpec{
				CPU:    "0.50",
				Memory: "1.00Gi",
			},
		},
		{
			requirements: []v1.ResourceRequirements{
				{
					Limits: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("1500m"),
						v1.ResourceMemory: resource.MustParse("1536Mi"),
					},
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("1000m"),
						v1.ResourceMemory: resource.MustParse("1024Mi"),
					},
				},
			},
			resources: api.ResourceSpec{
				CPU:    "1.50",
				Memory: "1.50Gi",
			},
		},
		{
			requirements: []v1.ResourceRequirements{
				{
					Limits: v1.ResourceList{
						ResourceLimitsGPU: resource.MustParse("1"),
					},
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("2000m"),
						v1.ResourceMemory: resource.MustParse("2Gi"),
					},
				},
			},
			resources: api.ResourceSpec{
				CPU:    "2.00",
				Memory: "2.00Gi",
				GPU:    "1",
			},
		},
		{
			requirements: []v1.ResourceRequirements{
				{
					Limits: v1.ResourceList{
						ResourceLimitsGPU: resource.MustParse("2"),
					},
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("2000m"),
						v1.ResourceMemory: resource.MustParse("2Gi"),
					},
				},
			},
			nodeSelector: map[string]string{
				"node.elotl.co/gpu-nvidia-tesla-v100": "",
			},
			resources: api.ResourceSpec{
				CPU:    "2.00",
				Memory: "2.00Gi",
				GPU:    "2 nvidia-tesla-v100",
			},
		},
		{
			requirements: []v1.ResourceRequirements{
				{
					Limits: v1.ResourceList{
						ResourceLimitsGPU: resource.MustParse("4"),
					},
					Requests: v1.ResourceList{
						v1.ResourceCPU:    resource.MustParse("2000m"),
						v1.ResourceMemory: resource.MustParse("2Gi"),
					},
				},
			},
			nodeSelector: map[string]string{
				"cloud.google.com/gke-accelerator": "nvidia-tesla-k80",
			},
			resources: api.ResourceSpec{
				CPU:    "2.00",
				Memory: "2.00Gi",
				GPU:    "4 nvidia-tesla-k80",
			},
		},
	}
	for _, tc := range testCases {
		containers := make([]v1.Container, len(tc.requirements))
		for i, req := range tc.requirements {
			containers[i].Resources = req
		}
		resources := aggregateResources(containers, tc.nodeSelector)
		assert.Equal(t, tc.resources, resources)
	}
}
