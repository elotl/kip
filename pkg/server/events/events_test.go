package events

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// It's halloween in a few days.  Might as well get in the spirit!
const VampireHunterEntersCastleEventType = "vampire-hunter-enters-castle"

type Vampire struct {
	Name   string
	Status string
}

type VampireHunter struct {
	Name   string
	Status string
}

func (vh *VampireHunter) KillVampire(v *Vampire) {
	fmt.Printf("My name is %s and I will kill %s\n", vh.Name, v.Name)
	v.Status = "Extinguished"
}

func (v *Vampire) Handle(e Event) error {
	fmt.Printf("Vampire %s is Handling %v\n", v.Name, e)
	switch e.Status {
	case VampireHunterEntersCastleEventType:
		vh := e.Object.(*VampireHunter)
		vh.KillVampire(v)
	default:
		return fmt.Errorf("Unknown event %s", e.Status)
	}
	return nil
}

func TestBasicHandler(t *testing.T) {
	t.Parallel()
	quitChan := make(chan struct{})
	wg := &sync.WaitGroup{}
	es := NewEventSystem(quitChan, wg)
	dracula := Vampire{"Dracula", "Undead"}
	buffy := VampireHunter{"Buffy", "Alive"}
	es.RegisterHandler(VampireHunterEntersCastleEventType, &dracula)
	es.Emit(VampireHunterEntersCastleEventType, "events-test", &buffy, "testing events")
	time.Sleep(1 * time.Second) // Todo, we can get around this sleep if needed
	assert.Equal(t, "Extinguished", dracula.Status)
}

type DeepCopyable struct {
	Name  string
	Slice []string
}

func (in *DeepCopyable) DeepCopy() *DeepCopyable {
	if in == nil {
		return nil
	}
	out := new(DeepCopyable)
	*out = *in
	if in.Slice != nil {
		in, out := &in.Slice, &out.Slice
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return out
}

type NotDeepCopyable struct {
	Name  string
	Slice []string
}

func TestCopyDeepCopyableEvents(t *testing.T) {
	// create an event with an object that can be DeepCopied
	original := DeepCopyable{
		Name:  "myobj",
		Slice: []string{"foo"},
	}
	e := Event{
		Object: &original,
	}
	cpy := copyEvent(e)
	o1 := e.Object.(*DeepCopyable)
	o2 := cpy.Object.(*DeepCopyable)
	p1 := *(*uint64)(unsafe.Pointer(&o1.Slice))
	p2 := *(*uint64)(unsafe.Pointer(&o2.Slice))
	assert.NotEqual(t, p1, p2)
	o1.Slice[0] = "bar"
	assert.Equal(t, "foo", o2.Slice[0])

	// Test copying a
	o1.Slice[0] = "foo"
	e.Object = original
	cpy = copyEvent(e)
	o3 := e.Object.(DeepCopyable)
	o4 := cpy.Object.(*DeepCopyable)
	p1 = *(*uint64)(unsafe.Pointer(&o3.Slice))
	p2 = *(*uint64)(unsafe.Pointer(&o4.Slice))
	assert.NotEqual(t, p1, p2)
}

func TestShallowCopyOtherObjects(t *testing.T) {
	original := NotDeepCopyable{
		Name:  "myobj",
		Slice: []string{"foo"},
	}
	e := Event{
		Object: &original,
	}
	cpy := copyEvent(e)
	o1 := e.Object.(*NotDeepCopyable)
	o2 := cpy.Object.(*NotDeepCopyable)
	p1 := *(*uint64)(unsafe.Pointer(&o1))
	p2 := *(*uint64)(unsafe.Pointer(&o2))
	p3 := *(*uint64)(unsafe.Pointer(&o1.Slice))
	p4 := *(*uint64)(unsafe.Pointer(&o2.Slice))
	assert.NotEqual(t, p1, p2)
	assert.Equal(t, p3, p4)
	o1.Slice[0] = "bar"
	assert.Equal(t, "bar", o2.Slice[0])
}
