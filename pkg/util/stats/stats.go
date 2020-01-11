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
