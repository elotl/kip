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

package registry

import (
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/docker/libkv/store"
	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/etcd"
	"github.com/elotl/kip/pkg/server/events"
	"github.com/elotl/kip/pkg/util"
	"k8s.io/klog"
)

const (
	EventDirectory                 string        = "milpa/events/"
	EventDirectoryPlaceholder      string        = "milpa/events/."
	EventTrashDirectory            string        = "milpa/trash/events/"
	EventTrashDirectoryPlaceholder string        = "milpa/trash/events/."
	DefaultTTL                     time.Duration = 1 * time.Hour
)

type EventRegistry struct {
	etcd.Storer
	codec       api.MilpaCodec
	eventSystem *events.EventSystem
	ttl         time.Duration // Warning: TTLs don't work in BoltDB.
}

func makeKey(prefix string, obj *api.ObjectReference, name string) string {
	key := prefix
	if obj != nil && obj.Kind != "" {
		key = key + "/" + obj.Kind
		if obj.Name != "" {
			key = key + "/" + obj.Name
			if obj.UID != "" {
				key = key + "/" + obj.UID
				if name != "" {
					key = key + "/" + name
				}
			}
		}
	}
	return key
}

func makeEventKey(obj *api.ObjectReference, name string) string {
	return makeKey(EventDirectory, obj, name)
}

func makeDeletedEventKey(obj *api.ObjectReference, name string) string {
	return makeKey(EventTrashDirectory, obj, name)
}

func NewEventRegistry(kvstore etcd.Storer, codec api.MilpaCodec, es *events.EventSystem) *EventRegistry {
	reg := &EventRegistry{kvstore, codec, es, DefaultTTL}
	_ = reg.Put(EventDirectoryPlaceholder, []byte("."), &store.WriteOptions{IsDir: true})
	_ = reg.Put(EventTrashDirectoryPlaceholder, []byte("."), &store.WriteOptions{IsDir: true})
	es.RegisterHandler(events.AllEvents, reg)
	return reg
}

func (reg *EventRegistry) Handle(e events.Event) error {
	obj := e.Object
	value := reflect.ValueOf(obj).Elem()
	// Event.Object is always a first-class Milpa object, i.e. a struct type
	// with TypeMeta and ObjectMeta embedded into it.
	kind := value.FieldByName("Kind").String()
	name := value.FieldByName("Name").String()
	uid := value.FieldByName("UID").String()
	ev := api.NewEvent()
	ev.InvolvedObject = api.ObjectReference{
		Kind: kind,
		Name: name,
		UID:  uid,
	}
	ev.Status = string(e.Status)
	ev.Source = e.Source
	ev.Message = e.Message
	_, err := reg.CreateEvent(ev)
	if err != nil {
		klog.Errorf("Error creating event %v in storage: %v", ev, err)
		return err
	}
	return nil
}

func (reg *EventRegistry) New() api.MilpaObject {
	return api.NewEvent()
}

func (reg *EventRegistry) Create(obj api.MilpaObject) (api.MilpaObject, error) {
	e := obj.(*api.Event)
	return reg.CreateEvent(e)
}

func (reg *EventRegistry) CreateEvent(e *api.Event) (*api.Event, error) {
	key := makeEventKey(&e.InvolvedObject, e.Name)
	exists, err := reg.Storer.Exists(key)
	if err != nil {
		return nil, err
	} else if exists {
		return nil, fmt.Errorf("Could not create event %v. %s",
			e, ErrAlreadyExists.Error())
	}
	data, err := reg.codec.Marshal(e)
	if err != nil {
		return nil, err
	}
	wo := store.WriteOptions{
		IsDir: false,
		TTL:   reg.ttl,
	}
	err = reg.Storer.Put(key, data, &wo)
	if err != nil {
		return nil, util.WrapError(err, "Could not create event in registry")
	}
	newEvent, err := reg.GetEvent(&e.InvolvedObject, e.Name)
	if err != nil {
		return nil, util.WrapError(err, "Could not get event after creation")
	}
	return newEvent, nil
}

func MatchAllEvents(e *api.Event) bool {
	return true
}

func (reg *EventRegistry) List() (api.MilpaObject, error) {
	return reg.ListEvents(MatchAllEvents)
}

func (reg *EventRegistry) ListEventsWithPrefix(prefix string, filter func(*api.Event) bool) (*api.EventList, error) {
	// TODO: test if this works with etcd.
	pairs, err := reg.Storer.List(prefix)
	eventList := api.NewEventList()
	if err != nil {
		klog.Errorf("Error listing events in storage with prefix %s: %v",
			prefix, err)
		return eventList, err
	}
	eventList.Items = make([]*api.Event, 0, len(pairs))
	for _, pair := range pairs {
		if pair.Key == EventDirectoryPlaceholder {
			continue
		}
		e := api.NewEvent()
		err = reg.codec.Unmarshal(pair.Value, e)
		if err != nil {
			klog.Errorf("Error unmarshalling single event in list operation: %v", err)
			continue
		}
		if filter(e) {
			eventList.Items = append(eventList.Items, e)
		}
	}
	sort.Slice(eventList.Items, func(i, j int) bool {
		return eventList.Items[i].CreationTimestamp.Before(eventList.Items[j].CreationTimestamp)
	})
	return eventList, nil
}

func (reg *EventRegistry) ListEvents(filter func(*api.Event) bool) (*api.EventList, error) {
	return reg.ListEventsWithPrefix(EventDirectory, filter)
}

func (reg *EventRegistry) ListEventsByObject(obj *api.ObjectReference) (*api.EventList, error) {
	return reg.ListEventsWithPrefix(makeEventKey(obj, ""), MatchAllEvents)
}

func (reg *EventRegistry) Get(name string) (api.MilpaObject, error) {
	evs, err := reg.ListEvents(func(e *api.Event) bool {
		return e.Name == name
	})
	if err != nil {
		return nil, err
	}
	if len(evs.Items) == 0 {
		return nil, fmt.Errorf("Event with id %s not found", name)
	}
	return evs.Items[0], nil
}

func (reg *EventRegistry) GetEvent(obj *api.ObjectReference, name string) (*api.Event, error) {
	key := makeEventKey(obj, name)
	pair, err := reg.Storer.Get(key)
	if err == store.ErrKeyNotFound {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("Error retrieving event from storage: %v", err)
	}
	event := api.NewEvent()
	err = reg.codec.Unmarshal(pair.Value, event)
	if err != nil {
		return nil, util.WrapError(err, "Error unmarshaling event from storage")
	}
	return event, nil
}

func (reg *EventRegistry) Delete(name string) (api.MilpaObject, error) {
	e, err := reg.Get(name)
	if err != nil {
		return nil, err
	}
	return reg.DeleteEvent(e.(*api.Event))
}

func (reg *EventRegistry) DeleteEvent(e *api.Event) (*api.Event, error) {
	now := api.Now()
	e.DeletionTimestamp = &now
	err := reg.Storer.Delete(makeEventKey(&e.InvolvedObject, e.Name))
	if err != nil {
		return nil, util.WrapError(err,
			"Error deleting %s from event registry", e.Name)
	}
	key := makeDeletedEventKey(&e.InvolvedObject, e.Name)
	data, err := reg.codec.Marshal(e)
	if err != nil {
		return nil, err
	}
	err = reg.Storer.Put(
		key,
		data,
		&store.WriteOptions{
			IsDir: false,
			TTL:   reg.ttl,
		})
	if err != nil {
		klog.Warningf("Could not create deleted event %s in registry: %v",
			e.Name, err)
	}
	return e, nil
}

func (reg *EventRegistry) Update(obj api.MilpaObject) (api.MilpaObject, error) {
	return nil, fmt.Errorf("Updating events is not supported")
}
