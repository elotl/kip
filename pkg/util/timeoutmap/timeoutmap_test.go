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

package timeoutmap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// test timeout
// test timeout fall function
// tset checkin

// func makeRunningMap() *TimeoutMap {
// 	return
// }

func TestBasic(t *testing.T) {
	t.Parallel()
	m := New(true, make(chan struct{}))
	k := "key"
	v := "value"
	m.Add(k, v, 1*time.Second, Noop)
	iface, exists := m.Get(k)
	assert.True(t, exists)
	value, worked := iface.(string)
	assert.True(t, worked)
	assert.Equal(t, v, value)
}

func TestTimeoutBasic(t *testing.T) {
	t.Parallel()
	m := New(true, make(chan struct{}))
	k := "key"
	v := "value"
	m.Add(k, v, 1*time.Millisecond, Noop)
	go m.Start(100 * time.Millisecond)
	time.Sleep(250 * time.Millisecond)
	_, exists := m.Get(k)
	assert.False(t, exists)
}

func TestTimeoutCallFunc(t *testing.T) {
	t.Parallel()
	m := New(true, make(chan struct{}))
	k := "key"
	v := "value"
	funcCalled := false
	m.Add(k, v, 1*time.Millisecond, func(obj interface{}) {
		value, worked := obj.(string)
		assert.True(t, worked)
		assert.Equal(t, v, value)
		funcCalled = true
	})
	go m.Start(100 * time.Millisecond)
	time.Sleep(250 * time.Millisecond)
	_, exists := m.Get(k)
	assert.False(t, exists)
	assert.True(t, funcCalled)
}

func TestCheckin(t *testing.T) {
	t.Parallel()
	m := New(true, make(chan struct{}))
	k := "key"
	v := "value"
	funcCalled := false
	m.Add(k, v, 200*time.Millisecond, func(obj interface{}) {
		funcCalled = true
	})
	go m.Start(100 * time.Millisecond)
	time.Sleep(100 * time.Millisecond)
	m.Checkin(k)
	time.Sleep(100 * time.Millisecond)
	m.Checkin(k)
	time.Sleep(100 * time.Millisecond)
	m.Checkin(k)
	time.Sleep(100 * time.Millisecond)
	m.Checkin(k)
	assert.False(t, funcCalled)
	_, exists := m.Get(k)
	assert.True(t, exists)
}
