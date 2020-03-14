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
	"flag"
	"testing"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/etcd"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
)

var runFunctional = flag.Bool("functional", false, "run functional tests")

func TestFilterReplyObject(t *testing.T) {
	t.Parallel()
	// create a registry
	// add an object
	// make sure we get the correct form of the object
	MaxEventListSize = 2
	er, closer := registry.SetupTestEventRegistry()
	defer closer()

	e := api.GetFakeEvent()
	_, err := er.CreateEvent(e)
	assert.NoError(t, err)
	objList, err := er.List()
	assert.NoError(t, err)
	filtered := filterReplyObject(objList)
	eList := filtered.(*api.EventList)
	assert.Len(t, eList.Items, 1)
	var lastTime api.Time
	for i := 0; i < MaxEventListSize+5; i++ {
		time.Sleep(10 * time.Millisecond)
		e := api.GetFakeEvent()
		_, err := er.CreateEvent(e)
		assert.NoError(t, err)
		e, _ = er.GetEvent(&e.InvolvedObject, e.Name)
		lastTime = e.CreationTimestamp
	}
	objList, err = er.List()
	assert.NoError(t, err)
	filtered = filterReplyObject(objList)
	eList = filtered.(*api.EventList)
	assert.Len(t, eList.Items, MaxEventListSize)
	assert.Equal(t, lastTime, eList.Items[MaxEventListSize-1].CreationTimestamp)
}

func TestEnsureRegionUnchanged(t *testing.T) {
	if !(*runFunctional) {
		t.Skip("skipping region change functional tests")
	}
	etcdClient, closer, err := etcd.SetupEmbeddedEtcdTest()
	assert.NoError(t, err)
	defer closer()
	err = ensureRegionUnchanged(etcdClient, "us-east-1")
	assert.NoError(t, err)
	err = ensureRegionUnchanged(etcdClient, "us-west-2")
	assert.Error(t, err)
	err = ensureRegionUnchanged(etcdClient, "us-east-1")
	assert.NoError(t, err)
}

//func getPortMappings(containers []v1.Container) []v1.ContainerPort
func TestGetPortMappings(t *testing.T) {
	testCases := []struct {
		containers   []v1.Container
		portMappings []v1.ContainerPort
	}{
		{
			containers:   nil,
			portMappings: nil,
		},
		{
			containers: []v1.Container{
				{
					Ports: []v1.ContainerPort{},
				},
			},
			portMappings: nil,
		},
		{
			containers: []v1.Container{
				{
					Ports: []v1.ContainerPort{
						{
							HostPort:      0,
							ContainerPort: 1111,
						},
						{
							HostPort:      12345,
							ContainerPort: 0,
						},
					},
				},
			},
			portMappings: nil,
		},
		{
			containers: []v1.Container{
				{
					Ports: []v1.ContainerPort{
						{
							HostPort:      1111,
							ContainerPort: 1111,
						},
						{
							HostPort:      2222,
							ContainerPort: 3333,
						},
						{
							HostPort:      0,
							ContainerPort: 12345,
						},
					},
				},
			},
			portMappings: []v1.ContainerPort{
				{
					HostPort:      1111,
					ContainerPort: 1111,
				},
				{
					HostPort:      2222,
					ContainerPort: 3333,
				},
			},
		},
	}
	for _, tc := range testCases {
		pms := getPortMappings(tc.containers)
		assert.Equal(t, tc.portMappings, pms)
	}
}
