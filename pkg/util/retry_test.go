package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var retryVar string

func setVar(value string) {
	retryVar = value
}

func alwaysError() error {
	return fmt.Errorf("That failed")
}

func TestNoRetryNeeded(t *testing.T) {
	retryVar = ""
	val := "Got Set"
	err := Retry(1,
		func() error { setVar(val); return nil },
		func(error) bool { return false })
	assert.Nil(t, err)
	assert.Equal(t, val, retryVar)
}

func TestTimesout(t *testing.T) {
	sleepDelay = time.Duration(1 * time.Microsecond)
	retryVar = ""
	timesCalled := 0
	err := Retry(
		50*time.Millisecond,
		func() error {
			timesCalled += 1
			return fmt.Errorf("That failed")
		},
		func(error) bool { return true })
	assert.NotNil(t, err)
	assert.True(t, timesCalled > 1)
	assert.Contains(t, err.Error(), "Timed out")
}

func TestRetryWorks(t *testing.T) {
	sleepDelay = time.Duration(1 * time.Microsecond)
	retryVar = ""
	timesCalled := 0
	err := Retry(
		250*time.Millisecond,
		func() error {
			timesCalled += 1
			if timesCalled == 3 {
				return nil
			}
			return alwaysError()
		},
		func(error) bool { return true })
	assert.Nil(t, err)
	assert.Equal(t, 3, timesCalled)
}

func TestIsRetryableFalse(t *testing.T) {
	sleepDelay = time.Duration(1 * time.Microsecond)
	retryVar = ""
	timesCalled := 0
	errMsg := "That failed"
	err := Retry(
		250*time.Millisecond,
		func() error {
			timesCalled += 1
			return alwaysError()
		},
		func(error) bool { return false })
	assert.NotNil(t, err)
	assert.Equal(t, 1, timesCalled)
	assert.Contains(t, err.Error(), errMsg)
}
