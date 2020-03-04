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

// Create a map that will timeout entries and call a callback
// when an object in the map has expired
package timeoutmap

import (
	"sync"
	"time"
)

type Node struct {
	key         string
	timeoutFunc func(interface{})
	obj         interface{}
	lastCheckin time.Time
	ttl         time.Duration
}

type TimeoutMap struct {
	sync.RWMutex
	data              map[string]*Node
	stop              chan struct{}   // stop called
	quit              <-chan struct{} // external quit called
	synchronousExpire bool
}

func New(synchronousExpire bool, quit <-chan struct{}) *TimeoutMap {
	return &TimeoutMap{
		data:              make(map[string]*Node),
		synchronousExpire: synchronousExpire,
		quit:              quit,
	}
}

func Noop(obj interface{}) {
}

func (m *TimeoutMap) Keys() []string {
	keys := make([]string, 0, len(m.data))
	for k, _ := range m.data {
		keys = append(keys, k)
	}
	return keys
}

func (m *TimeoutMap) Add(key string, obj interface{}, ttl time.Duration, timeoutFunc func(obj interface{})) {
	m.Lock()
	m.data[key] = &Node{
		key:         key,
		timeoutFunc: timeoutFunc,
		obj:         obj,
		lastCheckin: time.Now().UTC(),
		ttl:         ttl,
	}
	m.Unlock()
}

func (m *TimeoutMap) Get(key string) (interface{}, bool) {
	m.RLock()
	n, exists := m.data[key]
	m.RUnlock()
	if exists {
		return n.obj, true
	} else {
		return nil, false
	}
}

func (m *TimeoutMap) Delete(key string) {
	m.Lock()
	delete(m.data, key)
	m.Unlock()
}

func (m *TimeoutMap) Checkin(key string) bool {
	m.RLock()
	n, exists := m.data[key]
	m.RUnlock()
	if exists {
		m.Lock()
		n.lastCheckin = time.Now().UTC()
		m.Unlock()
		return true
	} else {
		return false
	}
}

func (m *TimeoutMap) check() {
	timedoutNodes := make([]*Node, 0)
	now := time.Now().UTC()
	m.RLock()
	for _, n := range m.data {
		if now.After(n.lastCheckin.Add(n.ttl)) {
			timedoutNodes = append(timedoutNodes, n)
		}
	}
	m.RUnlock()
	if len(timedoutNodes) == 0 {
		return
	}
	m.Lock()
	for _, n := range timedoutNodes {
		delete(m.data, n.key)
	}
	m.Unlock()
	for _, n := range timedoutNodes {
		if m.synchronousExpire {
			n.timeoutFunc(n.obj)
		} else {
			go n.timeoutFunc(n.obj)
		}
	}
}

// We could create a goroutine for each element in the map
// and fire timers off of that but it seems a bit excessive
// instead just clean up the map periodically
func (m *TimeoutMap) Start(period time.Duration) {
	m.stop = make(chan struct{})
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			m.check()
		case <-m.stop:
			return
		case <-m.quit:
			return
		}
	}
}

func (m *TimeoutMap) Stop() {
	if m.stop != nil {
		m.stop <- struct{}{}
	}
}
