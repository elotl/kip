package controllerqueue

import (
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
)

type QueueFunc func(interface{}) error

type option func(*Queue)

func NumWorkers(n int) option {
	return func(cq *Queue) {
		cq.numWorkers = n
	}
}

func MaxRetries(n int) option {
	return func(cq *Queue) {
		cq.maxRetries = n
	}
}

func Period(d time.Duration) option {
	return func(cq *Queue) {
		cq.workerLoopPeriod = d
	}
}

type Queue struct {
	name             string
	queue            workqueue.RateLimitingInterface
	workerLoopPeriod time.Duration
	numWorkers       int
	maxRetries       int
	f                QueueFunc
}

func New(name string, f QueueFunc, opts ...option) *Queue {
	queue := &Queue{
		name:             name,
		queue:            workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), name),
		workerLoopPeriod: 1 * time.Second,
		numWorkers:       1,
		maxRetries:       10,
		f:                f,
	}
	for _, opt := range opts {
		opt(queue)
	}
	return queue
}

func (cq *Queue) Start(quit <-chan struct{}) {
	for i := 0; i < cq.numWorkers; i++ {
		// Note: The queue is shut down in the fullSyncLoop
		go wait.Until(cq.worker, cq.workerLoopPeriod, quit)
	}
	go cq.listenForQuit(quit)
}

func (cq *Queue) listenForQuit(quit <-chan struct{}) {
	<-quit
	cq.queue.ShutDown()
}

func (cq *Queue) handleErr(err error, key interface{}) {
	if err == nil {
		cq.queue.Forget(key)
		return
	}

	if cq.queue.NumRequeues(key) < cq.maxRetries {
		klog.V(2).Infof("Error syncing %s %q, retrying. Error: %v", cq.name, key, err)
		cq.queue.AddRateLimited(key)
		return
	}

	klog.Warningf("Dropping %s %q out of the queue: %v", cq.name, key, err)
	cq.queue.Forget(key)
}

func (cq *Queue) worker() {
	for cq.processNextWorkItem() {
	}
}

func (cq *Queue) processNextWorkItem() bool {
	key, quit := cq.queue.Get()
	if quit {
		return false
	}
	defer cq.queue.Done(key)

	err := cq.f(key)
	if err != nil {
		cq.handleErr(err, key)
	}

	return true
}

func (cq *Queue) Len() int {
	return cq.queue.Len()
}

func (cq *Queue) Get() (interface{}, bool) {
	return cq.queue.Get()
}

func (cq *Queue) Add(item interface{}) {
	cq.queue.Add(item)
}
