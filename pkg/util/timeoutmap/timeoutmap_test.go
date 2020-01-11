package timeoutmap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// test timeout
// test timeout fall function
// tset checkin

// func makeRunningMap() *TimeoutMap {
// 	return
// }

func TestBasic(t *testing.T) {
	t.Parallel()
	m := New(true, make(chan struct{}))
	k := "key"
	v := "value"
	m.Add(k, v, 1*time.Second, Noop)
	iface, exists := m.Get(k)
	assert.True(t, exists)
	value, worked := iface.(string)
	assert.True(t, worked)
	assert.Equal(t, v, value)
}

func TestTimeoutBasic(t *testing.T) {
	t.Parallel()
	m := New(true, make(chan struct{}))
	k := "key"
	v := "value"
	m.Add(k, v, 1*time.Millisecond, Noop)
	go m.Start(100 * time.Millisecond)
	time.Sleep(250 * time.Millisecond)
	_, exists := m.Get(k)
	assert.False(t, exists)
}

func TestTimeoutCallFunc(t *testing.T) {
	t.Parallel()
	m := New(true, make(chan struct{}))
	k := "key"
	v := "value"
	funcCalled := false
	m.Add(k, v, 1*time.Millisecond, func(obj interface{}) {
		value, worked := obj.(string)
		assert.True(t, worked)
		assert.Equal(t, v, value)
		funcCalled = true
	})
	go m.Start(100 * time.Millisecond)
	time.Sleep(250 * time.Millisecond)
	_, exists := m.Get(k)
	assert.False(t, exists)
	assert.True(t, funcCalled)
}

func TestCheckin(t *testing.T) {
	t.Parallel()
	m := New(true, make(chan struct{}))
	k := "key"
	v := "value"
	funcCalled := false
	m.Add(k, v, 200*time.Millisecond, func(obj interface{}) {
		funcCalled = true
	})
	go m.Start(100 * time.Millisecond)
	time.Sleep(100 * time.Millisecond)
	m.Checkin(k)
	time.Sleep(100 * time.Millisecond)
	m.Checkin(k)
	time.Sleep(100 * time.Millisecond)
	m.Checkin(k)
	time.Sleep(100 * time.Millisecond)
	m.Checkin(k)
	assert.False(t, funcCalled)
	_, exists := m.Get(k)
	assert.True(t, exists)
}
