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
	"testing"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestEventHandler(t *testing.T) {
	er, closer := SetupTestEventRegistry()
	defer closer()

	p1 := api.GetFakePod()
	er.eventSystem.Emit("pod-created", "event-tests", p1, "Created pod1")
	p2 := api.GetFakePod()
	er.eventSystem.Emit("pod-created", "event-tests", p2, "Created pod2")
	er.eventSystem.Emit("pod-deleted", "event-tests", p1, "Deleted pod1")

	events, err := er.ListEvents(MatchAllEvents)
	assert.Nil(t, err)
	start := api.Now()
	for len(events.Items) < 3 {
		if api.Now().After(start.Add(3 * time.Second)) {
			break
		}
		time.Sleep(10 * time.Millisecond)
		events, err = er.ListEvents(MatchAllEvents)
	}
	assert.Nil(t, err)
	assert.Len(t, events.Items, 3)
	assert.Equal(t, events.Items[0].InvolvedObject.UID, p1.UID)
	assert.Equal(t, events.Items[0].Status, "pod-created")
	assert.Equal(t, events.Items[1].InvolvedObject.UID, p2.UID)
	assert.Equal(t, events.Items[1].Status, "pod-created")
	assert.Equal(t, events.Items[2].InvolvedObject.UID, p1.UID)
	assert.Equal(t, events.Items[2].Status, "pod-deleted")
	for i := 0; i < 2; i++ {
		assert.Equal(t, events.Items[i].Source, "event-tests")
	}
}

func TestCreateAndGet(t *testing.T) {
	er, closer := SetupTestEventRegistry()
	defer closer()

	e1 := api.GetFakeEvent()
	_, err := er.Create(e1)
	assert.Nil(t, err)

	e2, err := er.Get(e1.Name)
	assert.Nil(t, err)
	assert.NotNil(t, e2)
	assertEventsEqual(t, e1, e2.(*api.Event))
}

func TestCreateAndListEvents(t *testing.T) {
	er, closer := SetupTestEventRegistry()
	defer closer()

	e1 := api.GetFakeEvent()
	_, err := er.Create(e1)
	assert.Nil(t, err)
	e2 := api.GetFakeEvent()
	_, err = er.Create(e2)
	assert.Nil(t, err)

	events, err := er.ListEvents(MatchAllEvents)
	assert.Nil(t, err)
	assert.Len(t, events.Items, 2)
	assertEventsEqual(t, events.Items[0], e1)
	assertEventsEqual(t, events.Items[1], e2)
}

func TestCreateAndDelete(t *testing.T) {
	er, closer := SetupTestEventRegistry()
	defer closer()

	e1 := api.GetFakeEvent()
	_, err := er.Create(e1)
	assert.Nil(t, err)
	e2 := api.GetFakeEvent()
	_, err = er.Create(e2)
	assert.Nil(t, err)

	_, err = er.Delete(e1.Name)
	assert.Nil(t, err)

	events, err := er.ListEvents(MatchAllEvents)
	assert.Nil(t, err)
	assert.Len(t, events.Items, 1)
	assertEventsEqual(t, events.Items[0], e2)
}

func TestCreateAndDeleteEvent(t *testing.T) {
	er, closer := SetupTestEventRegistry()
	defer closer()

	e1 := api.GetFakeEvent()
	_, err := er.Create(e1)
	assert.Nil(t, err)
	e2 := api.GetFakeEvent()
	_, err = er.Create(e2)
	assert.Nil(t, err)

	_, err = er.DeleteEvent(e1)
	assert.Nil(t, err)

	events, err := er.ListEvents(MatchAllEvents)
	assert.Nil(t, err)
	assert.Len(t, events.Items, 1)
	assertEventsEqual(t, events.Items[0], e2)
}

func assertEventsEqual(t *testing.T, e1, e2 *api.Event) {
	// assert.Equal() uses DeepEqual() under the hood, which can't compare
	// time.Time values: https://github.com/golang/go/issues/10089.
	assert.Equal(t, e1.InvolvedObject, e2.InvolvedObject)
	assert.Equal(t, e1.Status, e2.Status)
	assert.Equal(t, e1.Source, e2.Source)
	assert.Equal(t, e1.Message, e2.Message)
}

func TestListEventsOrder(t *testing.T) {
	er, closer := SetupTestEventRegistry()
	defer closer()

	for i := 0; i < 10; i++ {
		e := api.GetFakeEvent()
		_, err := er.Create(e)
		assert.Nil(t, err)
	}

	events, err := er.ListEvents(MatchAllEvents)
	assert.Nil(t, err)
	assert.Len(t, events.Items, 10)

	for i := 1; i < 10; i++ {
		cur := events.Items[i]
		prev := events.Items[i-1]
		assert.True(t, cur.CreationTimestamp.After(prev.CreationTimestamp))
	}
}

func TestListEventsByObject(t *testing.T) {
	er, closer := SetupTestEventRegistry()
	defer closer()

	for i := 0; i < 3; i++ {
		e := api.NewEvent()
		e.InvolvedObject = api.ObjectReference{
			Kind: "object",
			Name: "myobject",
			UID:  fmt.Sprintf("my-uid-%d", i),
		}
		e.Status = "object-created"
		e.Source = "event-tests"
		e.Message = "Testing events"
		_, err := er.Create(e)
		assert.Nil(t, err)
	}

	obj := api.ObjectReference{
		Kind: "object",
		Name: "myobject",
	}
	events, err := er.ListEventsByObject(&obj)
	assert.Nil(t, err)
	assert.Len(t, events.Items, 3)
	for i := 0; i < 3; i++ {
		o := events.Items[i].InvolvedObject
		assert.Equal(t, o.Kind, "object")
		assert.Equal(t, o.Name, "myobject")
	}
}
