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

package nodemanager

import (
	"fmt"
	"time"

	"github.com/elotl/kip/pkg/util/timeoutmap"
	"k8s.io/klog"
)

var (
	bootLimiterPeriod   = 33 * time.Second
	unavailableDuration = 3 * time.Minute
)

type InstanceBootLimiter struct {
	unavailableInstances *timeoutmap.TimeoutMap
}

func NewInstanceBootLimiter() *InstanceBootLimiter {
	b := &InstanceBootLimiter{
		unavailableInstances: timeoutmap.New(true, nil),
	}
	return b
}

func (s *InstanceBootLimiter) Start() {
	go s.unavailableInstances.Start(bootLimiterPeriod)
}

func makeUnavailableKey(instanceType string, spot bool) string {
	return fmt.Sprintf("%s/%t", instanceType, spot)
}

func (s *InstanceBootLimiter) AddUnavailableInstance(instanceType string, spot bool) {
	key := makeUnavailableKey(instanceType, spot)
	_, exists := s.unavailableInstances.Get(key)
	if !exists {
		launchType := "On-Demand"
		if spot {
			launchType = "Spot"
		}
		klog.Warningf("%s instance type %s is unavailable, boots of that instance type will be suspended for %d seconds", launchType, instanceType, int64(unavailableDuration.Seconds()))
		s.unavailableInstances.Add(key, struct{}{}, unavailableDuration, timeoutmap.Noop)
	}
}

func (s *InstanceBootLimiter) IsUnavailableInstance(instanceType string, spot bool) bool {
	key := makeUnavailableKey(instanceType, spot)
	_, exists := s.unavailableInstances.Get(key)
	return exists
}
