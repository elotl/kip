/*
Copyright 2015 The Kubernetes Authors.
Copyright 2018 Elotl Inc.

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
	"github.com/elotl/cloud-instance-provider/pkg/api"
)

// ActivePods type allows custom sorting of pods so a controller can
// pick the best ones to delete.
// Taken from k8s
type ActivePods []*api.Pod

func (s ActivePods) Len() int      { return len(s) }
func (s ActivePods) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ActivePods) Less(i, j int) bool {
	//podWaiting < PodFailed < PodDispatching < podRunning
	m := map[api.PodPhase]int{
		api.PodWaiting:     0,
		api.PodFailed:      1,
		api.PodDispatching: 2,
		api.PodRunning:     3,
	}

	if m[s[i].Status.Phase] != m[s[j].Status.Phase] {
		return m[s[i].Status.Phase] < m[s[j].Status.Phase]
	}

	// todo, when we have readyness and ready flags, add those here

	if s[i].Status.StartFailures != s[j].Status.StartFailures {
		return s[i].Status.StartFailures > s[j].Status.StartFailures
	}

	if !s[i].CreationTimestamp.Equal(s[j].CreationTimestamp) {
		return s[i].CreationTimestamp.Before(s[j].CreationTimestamp)
	}

	return false
}
