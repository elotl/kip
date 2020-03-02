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

package stats

import (
	"sync"
	"time"
)

type LoopTimer struct {
	Start    time.Time
	End      time.Time
	LastLoop time.Duration
	Average  time.Duration
	Count    int64
	sync.Mutex
}

func (t *LoopTimer) Copy() *LoopTimer {
	t.Lock()
	defer t.Unlock()
	lt := &LoopTimer{
		Start:    t.Start,
		End:      t.End,
		LastLoop: t.LastLoop,
		Average:  t.Average,
		Count:    t.Count,
	}
	return lt
}

func (t *LoopTimer) StartLoop() {
	t.Lock()
	defer t.Unlock()

	t.Start = time.Now()
}

func (t *LoopTimer) EndLoop() {
	t.Lock()
	defer t.Unlock()

	t.End = time.Now()
	t.LastLoop = t.End.Sub(t.Start)
	// Oh boy!  Here I go casting again!
	t.Average = time.Duration((int64(t.Average)*t.Count + int64(t.LastLoop)) / (t.Count + 1))
	t.Count += 1
}
