package gce

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/compute/v1"
)

func getGCE(t *testing.T, controllerID string) *gceClient {
	// to use this, set the env var GOOGLE_APPLICATION_CREDENTIALS
	// to the path to your credentials file
	// e.g.
	// export GOOGLE_APPLICATION_CREDENTIALS = "/home/myuser/google/gce-project-fb308cd448d5.json"
	c, err := NewGCEClient(controllerID, "manualTestCase", "milpa-207719", WithVPCName("default"), WithZone("us-central1-a"), WithSubnetName("default"))
	assert.NoError(t, err)
	return c
}

func TestWaitForBackoff(t *testing.T) {
	tests := []struct {
		i   int
		exp time.Duration
	}{
		{i: 0, exp: 1},
		{i: 1, exp: 1},
		{i: 3, exp: 3},
		{i: 4, exp: 5},
		{i: 5, exp: 5},
		{i: 6, exp: 5},
	}
	for _, tc := range tests {
		res := waitBackoff(tc.i)
		assert.Equal(t, tc.exp*time.Second, res)
	}
}

func TestWaitForOperation(t *testing.T) {
	tests := []struct {
		opVals     []string
		opErr      bool
		returnsErr bool
	}{
		{
			opVals:     []string{statusOperationDone},
			opErr:      false,
			returnsErr: false,
		},
		{
			opVals:     []string{""},
			opErr:      true,
			returnsErr: true,
		},
		{
			opVals:     []string{"PENDING", statusOperationDone},
			opErr:      false,
			returnsErr: false,
		},
	}
	for _, tc := range tests {
		calledCount := 0
		f := func(s string) (*compute.Operation, error) {
			defer func() { calledCount++ }()
			if tc.opErr {

				return nil, fmt.Errorf("operation failed")
			}
			return &compute.Operation{Status: tc.opVals[calledCount]}, nil
		}
		err := waitOnOperation("testop", f)
		if tc.returnsErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
