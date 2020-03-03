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

package loop

import (
	"sync"
	"time"

	"k8s.io/klog"
)

type LoopFunc func() error

type Loop struct {
	name   string
	period time.Duration
	f      LoopFunc
}

func New(name string, period time.Duration, f LoopFunc) *Loop {
	return &Loop{
		name:   name,
		f:      f,
		period: period,
	}
}

func (loop *Loop) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	go loop.run(quit, wg)
}

func (loop *Loop) run(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	tick := time.NewTicker(loop.period)
	for {
		select {
		case <-tick.C:
			err := loop.f()
			if err != nil {
				klog.Errorf("Error executing %s Loop: %s", loop.name, err.Error())
			}
		case <-quit:
			tick.Stop()
			klog.V(2).Infof("Exiting %s Loop", loop.name)
			return
		}
	}
}
