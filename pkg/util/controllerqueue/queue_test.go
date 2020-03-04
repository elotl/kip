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

package controllerqueue

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func waitUntilEmpty(wq *Queue) error {
	// wait 5s for an empty queue
	for i := 0; i < 100; i++ {
		if wq.Len() == 0 {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("Timeout waiting for empty queue")
}

func TestControllerQueue(t *testing.T) {
	itemsProcessed := 0
	processItem := func(iface interface{}) error {
		itemsProcessed += 1
		return nil
	}
	wq := New("testqueue", processItem, NumWorkers(1), MaxRetries(0), Period(50*time.Millisecond))
	quit := make(chan struct{})
	wq.Start(quit)
	assert.Equal(t, 0, wq.Len())
	wq.Add("item1")
	assert.Equal(t, 1, wq.Len())
	err := waitUntilEmpty(wq)
	assert.NoError(t, err)
	assert.Equal(t, 0, wq.Len())
	assert.Equal(t, 1, itemsProcessed)
	close(quit)
	for i := 0; i < 60; i++ {
		time.Sleep(50 * time.Millisecond)
		if wq.queue.ShuttingDown() {
			break
		}
	}
	assert.True(t, wq.queue.ShuttingDown())
}
