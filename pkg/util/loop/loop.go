package loop

import (
	"sync"
	"time"

	"github.com/golang/glog"
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
				glog.Errorf("Error executing %s Loop: %s", loop.name, err.Error())
			}
		case <-quit:
			tick.Stop()
			glog.Infof("Exiting %s Loop", loop.name)
			return
		}
	}
}
