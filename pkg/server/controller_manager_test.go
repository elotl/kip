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
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type FakeController struct {
	StartCalled int
}

func (c *FakeController) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	c.StartCalled += 1
	wg.Add(1)
	defer wg.Done()
	select {//nolint
	case <-quit:
		return
	}
}

func (c *FakeController) Dump() []byte {
	return nil
}

func setupControllerManagerTest() *ControllerManager {
	// create a fake controller that can be started and stopped
	ctls := map[string]Controller{
		"FakeController": &FakeController{},
	}
	return NewControllerManager(ctls)
}

func TestControllerManagerGet(t *testing.T) {
	cm := setupControllerManagerTest()
	_, exists := cm.GetController("FakeController")
	assert.True(t, exists)
	_, exists = cm.GetController("NonExistantController")
	assert.False(t, exists)
}

func waitForControllersRunningState(t *testing.T, cm *ControllerManager, isRunning bool) {
	deadline := time.Now().Add(5 * time.Second)
	for {
		if cm.ControllersRunning() == isRunning {
			break
		}
		if time.Now().After(deadline) {
			t.Errorf("Controller not running but should be running")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func TestControllerManagerStartStop(t *testing.T) {
	cm := setupControllerManagerTest()
	assert.False(t, cm.ControllersRunning())

	go cm.Start()

	iterations := 2
	for i := 0; i < iterations; i++ {
		cm.StartControllers()
		waitForControllersRunningState(t, cm, true)
		cm.StartControllers()

		cm.StopControllers()
		waitForControllersRunningState(t, cm, false)
		// double close() on a chan will panic if we don't guard against
		// double stopping
		cm.StopControllers()
	}
	ctl, exists := cm.GetController("FakeController")
	assert.True(t, exists)
	fakeCtl := ctl.(*FakeController)
	assert.Equal(t, iterations, fakeCtl.StartCalled)
}
