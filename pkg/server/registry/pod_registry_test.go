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

package registry

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestPodCreateGet(t *testing.T) {
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()

	p1 := api.GetFakePod()
	_, err := podRegistry.CreatePod(p1)
	if err != nil {
		t.Error(err)
	}
	p2, err := podRegistry.GetPod(p1.Name)
	if err != nil {
		t.Error(err)
	}
	p3, err := podRegistry.GetPod(p2.Name)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(p3, p2) {
		t.Errorf("did not get the same pod back:\n%#v\n%#v", p1, p2)
	}
}

func TestPodDoubleCreate(t *testing.T) {
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()

	p1 := api.GetFakePod()
	_, err := podRegistry.CreatePod(p1)
	assert.Nil(t, err)
	_, err = podRegistry.CreatePod(p1)
	assert.NotNil(t, err)
}

func TestTerminatePodAndReCreate(t *testing.T) {
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()

	origPod := api.GetFakePod()
	p, err := podRegistry.CreatePod(origPod)
	assert.NoError(t, err)
	err = podRegistry.TerminatePod(p, api.PodTerminated, "")
	assert.NoError(t, err)
	allPods, err := podRegistry.ListPods(MatchAllPods)
	assert.NoError(t, err)
	assert.Len(t, allPods.Items, 0)
	allLivePods, err := podRegistry.ListPods(MatchAllLivePods)
	assert.NoError(t, err)
	assert.Len(t, allLivePods.Items, 0)
	_, err = podRegistry.CreatePod(origPod)
	assert.NoError(t, err)
	allLivePods, err = podRegistry.ListPods(MatchAllLivePods)
	assert.NoError(t, err)
	assert.Len(t, allLivePods.Items, 1)
}

func TestMarkPodForTermination(t *testing.T) {
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()
	p1 := api.GetFakePod()
	_, err := podRegistry.CreatePod(p1)
	if err != nil {
		t.Error(err)
	}
	p2, err := podRegistry.MarkPodForTermination(p1)
	if err != nil {
		t.Error(err)
	}
	if p2.Spec.Phase != api.PodTerminated ||
		p2.Status.Phase != p1.Status.Phase {
		t.Errorf("Delete should set spec to terminated but not set status")
	}
}

func TestListPods(t *testing.T) {
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()
	p1 := api.GetFakePod()
	_, err := podRegistry.CreatePod(p1)
	if err != nil {
		t.Error(err)
	}
	p2 := api.GetFakePod()
	_, err = podRegistry.CreatePod(p2)
	if err != nil {
		t.Error(err)
	}

	apilist, err := podRegistry.List()
	pods := apilist.(*api.PodList)
	if err != nil {
		t.Error(err)
	} else if len(pods.Items) != 2 {
		t.Errorf("Expected to get 2 pods back, got %d", len(pods.Items))
	}

	pods, err = podRegistry.ListPods(func(p *api.Pod) bool {
		return p.Name == p1.Name
	})
	if err != nil {
		t.Error(err)
	} else if len(pods.Items) != 1 {
		t.Errorf("Error listing pods with predicate, should have filtered down to 1 pod, got %d", len(pods.Items))
	}
}

func TestCreatePodInstanceFromResources(t *testing.T) {
	// setup and configure the instanceselector
	// try creating a pod from resources
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()
	p1 := api.GetFakePod()
	p1.Spec.InstanceType = ""
	p1.Spec.Resources.CPU = "1"
	p1.Spec.Resources.Memory = "1Gi"
	p1.Spec.Resources.DedicatedCPU = true
	_, err := podRegistry.CreatePod(p1)
	assert.Nil(t, err)
	p2, err := podRegistry.GetPod(p1.Name)
	assert.Nil(t, err)
	assert.Equal(t, "c5.large", p2.Spec.InstanceType)
}

func TestPodPhaseUpdateUpdatesTime(t *testing.T) {
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()
	p := api.GetFakePod()
	p, err := podRegistry.CreatePod(p)
	assert.NoError(t, err)
	t1 := p.Status.LastPhaseChange
	// Change to running will fail without going through dispatching
	p.Status.Phase = api.PodRunning
	time.Sleep(10 * time.Millisecond)
	newPod, err := podRegistry.UpdatePodStatus(p, "")
	assert.Error(t, err)
	assert.Equal(t, t1, newPod.Status.LastPhaseChange)
	p.Status.Phase = api.PodDispatching
	time.Sleep(10 * time.Millisecond)
	newPod, err = podRegistry.UpdatePodStatus(p, "")
	assert.NoError(t, err)
	assert.True(t, newPod.Status.LastPhaseChange.After(t1))
}

func TestAtomicUpdateHappyPath(t *testing.T) {
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()
	p1 := api.GetFakePod()
	_, err := podRegistry.CreatePod(p1)
	if err != nil {
		t.Error(err)
	}
	_, err = podRegistry.AtomicUpdate(p1.Name, func(p *api.Pod) error {
		p.Status.Phase = api.PodTerminated
		return nil
	})
	p2, err := podRegistry.GetPod(p1.Name)
	if err != nil {
		t.Error(err)
	} else if p2.Status.Phase != api.PodTerminated ||
		p2.Name != p1.Name {
		t.Errorf("Error in atomic update of pod")
	}
}

func TestAtomicUpdateWithAfterUpdate(t *testing.T) {
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()
	// Make sure that we can do
	p1 := api.GetFakePod()
	_, err := podRegistry.CreatePod(p1)
	if err != nil {
		t.Error(err)
	}
	// To test a collision, we create a modifier function with a side
	// effect that will run only once.  That side effect is an update
	// to the same pod.  Make sure they both updates run correctly and
	// that our main update runs twice and the inner update runs once.
	modifyCallCount := 0
	var once sync.Once
	p2, err := podRegistry.AtomicUpdate(p1.Name, func(p *api.Pod) error {
		once.Do(func() {
			innerPod, err := podRegistry.GetPod(p1.Name)
			if err != nil {
				t.Error(err)
			}
			_, err = podRegistry.AtomicUpdate(innerPod.Name, func(innerP *api.Pod) error {
				modifyCallCount += 1
				innerP.Spec.Phase = api.PodTerminated
				return nil
			})
			if err != nil {
				t.Error(err)
			}
		})
		modifyCallCount += 1
		p.Status.Phase = api.PodTerminated
		return nil
	})

	if err != nil {
		t.Error(err)
	} else if p2.Status.Phase != api.PodTerminated ||
		p2.Spec.Phase != api.PodTerminated ||
		p2.Name != p1.Name {
		t.Errorf("Error in atomic update of pod")
	}
	p3, err := podRegistry.GetPod(p1.Name)
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(p3, p2) {
		t.Errorf("did not get the same pod back after atomic update:\n%#v\n%#v", p1, p2)
	}
	if modifyCallCount != 3 {
		t.Errorf("Should have seen modify get called 3 times, saw %d:", modifyCallCount)
	}
}

func TestAtomicUpdateWithModifyFailure(t *testing.T) {
	podRegistry, closer := SetupTestPodRegistry()
	defer closer()
	p1 := api.GetFakePod()
	_, err := podRegistry.CreatePod(p1)
	if err != nil {
		t.Error(err)
	}
	// Create a collision (see TestAtomicUpdateWithAfterUpdate)
	// but have our modify function fail.  Make sure
	// we return an error
	var once sync.Once
	modifyCallCount := 0
	_, err = podRegistry.AtomicUpdate(p1.Name, func(p *api.Pod) error {
		once.Do(func() {
			innerPod, err := podRegistry.GetPod(p1.Name)
			if err != nil {
				t.Error(err)
			}
			_, err = podRegistry.AtomicUpdate(innerPod.Name, func(innerP *api.Pod) error {
				modifyCallCount += 1
				innerP.Status.Phase = api.PodTerminated
				return nil
			})
			if err != nil {
				t.Error(err)
			}
		})
		modifyCallCount += 1
		if p.Status.Phase != api.PodWaiting {
			return fmt.Errorf("Expected to find pod in creating state")
		}
		p.Status.Phase = api.PodRunning
		return nil
	})

	if err == nil {
		t.Errorf("Atomic update should have returned an error:")
	}
	finalPod, err := podRegistry.GetPod(p1.Name)
	if err != nil {
		t.Error(err)
	}
	if finalPod.Status.Phase != api.PodTerminated {
		t.Errorf("pod should have terminated phase, has: %s", finalPod.Status.Phase)
	}
}
