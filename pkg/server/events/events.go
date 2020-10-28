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

package events

import (
	"fmt"
	"github.com/elotl/kip/pkg/k8sclient/clientset/versioned/scheme"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
	"reflect"
	"sync"

	"k8s.io/klog"
)

var ArbitraryChanSize = 10000

type Event struct {
	Status  string
	Source  string
	Message string
	Object  interface{}
}

type EventHandler interface {
	Handle(Event) error
}

// Copy net/http's ability to register any function
// as a handler for an event
type HandlerFunc func(Event) error

func (f HandlerFunc) Handle(e Event) error {
	return f(e)
}

type EventSystem struct {
	sync.RWMutex
	eventHandlers map[string][]EventHandler
	eventChan     chan Event
	k8sRecorder   record.EventRecorder
}

func NewEventSystem(quit <-chan struct{}, wg *sync.WaitGroup, kubeClient clientset.Interface) *EventSystem {
	eventBroadcater := record.NewBroadcaster()
	eventBroadcater.StartRecordingToSink(&typedv1.EventSinkImpl{
		Interface: typedv1.New(kubeClient.CoreV1().RESTClient()).Events(""),
	})
	recorder := eventBroadcater.NewRecorder(
		scheme.Scheme,
		v1.EventSource{Component: "kip"})
	e := &EventSystem{
		eventHandlers: make(map[string][]EventHandler),
		eventChan:     make(chan Event, ArbitraryChanSize),
		k8sRecorder:   &recorder,
	}
	go e.Run(quit, wg)
	return e
}

// At this time, you can't unregister from an event
// Why?  Because we don't need that functionality yet
// If you need that functionality, feel free to add it
func (es *EventSystem) RegisterHandler(status string, h EventHandler) {
	es.Lock()
	defer es.Unlock()
	es.eventHandlers[status] = append(es.eventHandlers[status], h)
}

func (es *EventSystem) RegisterHandlerFunc(status string, h func(Event) error) {
	es.Lock()
	defer es.Unlock()
	es.eventHandlers[status] = append(
		es.eventHandlers[status], HandlerFunc(h))
}

func (es *EventSystem) Emit(status, source string, obj interface{}, args ...interface{}) {
	// Optional argument, which can be a string, or a printf-style format
	// string with arguments.
	msg := ""
	if len(args) == 1 {
		msg = args[0].(string)
	} else if len(args) > 1 {
		msg = fmt.Sprintf(args[0].(string), args[1:]...)
	}
	e := Event{status, source, msg, obj}
	// We copy the event here (in addition to later) since the user
	// might have passed in a pointer to an object that the call can mutate
	// after we place the event in the channel and return from Emit.
	// The race detector doesn't like it when that happens.
	eCpy := copyEvent(e)
	es.eventChan <- eCpy
	k8sECpy := copyEvent(e)
	es.k8sRecorder.Eventf()
	es.RecordK8sEvent(k8sECpy)
}

// Events are passed around the system and the resulting objects
// can end up in other goroutines.  Since our API types have
// reference types in them, we try to deep copy these objects
// to prevent data races from other mutators.
func copyEvent(e Event) Event {
	cpy := e
	// I know you shouldn't comment exactly what code is doing
	// but I can never remember how the reflect API works so:
	// if we get a pointer, use that, otherwise convert to a pointer
	// because DeepCopy has a pointer receiver
	var v reflect.Value
	if reflect.ValueOf(e.Object).Kind() == reflect.Ptr {
		v = reflect.ValueOf(e.Object)
	} else {
		val := reflect.ValueOf(e.Object)
		vp := reflect.New(val.Type())
		vp.Elem().Set(val)
		v = reflect.ValueOf(vp.Interface())
	}
	// Here we're calling DeepCopy and gathering the result if the
	// method exists otherwise we do a shallow copy of the object if
	// it's not deepcopyable
	deepCopyMethod := v.MethodByName("DeepCopy")
	if deepCopyMethod.Kind() != reflect.Invalid {
		copiedValue := deepCopyMethod.Call([]reflect.Value{})
		cpy.Object = copiedValue[0].Interface()
	} else if reflect.ValueOf(cpy.Object).Kind() == reflect.Ptr {
		val := reflect.ValueOf(cpy.Object).Elem()
		vp := reflect.New(val.Type())
		vp.Elem().Set(val)
		cpy.Object = vp.Interface()
	}

	return cpy
}

func (es *EventSystem) Run(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case event := <-es.eventChan:
			es.RLock()
			handlers := es.eventHandlers[event.Status]
			globalHandlers := es.eventHandlers[AllEvents]
			handlers = append(handlers, globalHandlers...)
			es.RUnlock()
			for _, eh := range handlers {
				eventCpy := copyEvent(event)
				if reflect.ValueOf(eventCpy.Object).Kind() != reflect.Ptr {
					klog.Errorf("Event objects must be pointers: %+v", event)
					break
				}
				err := eh.Handle(eventCpy)
				if err != nil {
					klog.Errorf("Error in %s event handler: %v", event.Status, err)
				}
			}
		case <-quit:
			klog.V(2).Info("Stopping Events System")
			return
		}
	}
}
