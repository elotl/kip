package server

import (
	"sync"

	"github.com/golang/glog"
	"github.com/uber-go/atomic"
)

// The ControllerManager was created to make the interaction
// between leader election and controllers easier.  It takes
// care of starting and stopping controllers based on the
// leader elector.
type ControllerManager struct {
	controllers         map[string]Controller
	controllersRunning  *atomic.Bool
	startChan           chan struct{}
	stopChan            chan struct{}
	controllerQuit      chan struct{}
	controllerWaitGroup *sync.WaitGroup
}

func NewControllerManager(controllers map[string]Controller) *ControllerManager {
	return &ControllerManager{
		controllers:         controllers,
		startChan:           make(chan struct{}),
		stopChan:            make(chan struct{}),
		controllersRunning:  atomic.NewBool(false),
		controllerQuit:      nil,
		controllerWaitGroup: nil,
	}
}

func (cm *ControllerManager) StartControllers() {
	cm.startChan <- struct{}{}
}

func (cm *ControllerManager) StopControllers() {
	cm.stopChan <- struct{}{}
}

func (cm *ControllerManager) ControllersRunning() bool {
	return cm.controllersRunning.Load()
}

func (cm *ControllerManager) GetController(name string) (Controller, bool) {
	c, exists := cm.controllers[name]
	return c, exists
}

func (cm *ControllerManager) GetAllControllers() map[string]Controller {
	cpy := map[string]Controller{}
	for k, v := range cm.controllers {
		cpy[k] = v
	}
	return cpy
}

// This doesn't take a quit channel on purpose. This is because if you
// start listening for quit, it becomes difficult to handle starting
// and stopping controllers through the channel as well so we just let
// this goroutine run until the end of the milpa process
func (cm *ControllerManager) Start() {
	for {
		select {
		case <-cm.startChan:
			cm.startControllersHelper()
		case <-cm.stopChan:
			cm.stopControllersHelper()
		}
	}
}

// Our leader election used to tell our controller manager to shutdown
// now this takes care of that. We could simplify the controller
// manager a fair amount...  I'm concerned we'll need a clustered
// system again eventually.
func (cm *ControllerManager) WaitForShutdown(systemShutdown <-chan struct{}, systemWG *sync.WaitGroup) {
	systemWG.Add(1)
	defer systemWG.Done()

	select {
	case <-systemShutdown:
		glog.Infof("Shutting down controllers")
		cm.StopControllers()
		return
	}
}

func (cm *ControllerManager) startControllersHelper() {
	if cm.ControllersRunning() {
		glog.Warning("Asked to start controllers but they are already running")
		return
	}
	glog.Info("Starting controllers")
	cm.controllerQuit = make(chan struct{})
	cm.controllerWaitGroup = &sync.WaitGroup{}
	cm.controllersRunning.Store(true)
	for name, controller := range cm.controllers {
		glog.Infof("Starting %s", name)
		go controller.Start(cm.controllerQuit, cm.controllerWaitGroup)
	}
	glog.Info("Finished starting controllers")
}

func (cm *ControllerManager) stopControllersHelper() {
	if !cm.ControllersRunning() {
		glog.Warning("Asked to stop controllers but they are not running")
		return
	}
	glog.Info("Starting to stop controllers")
	close(cm.controllerQuit)
	cm.controllerWaitGroup.Wait()
	cm.controllersRunning.Store(false)
	glog.Info("All controllers stopped")
}
