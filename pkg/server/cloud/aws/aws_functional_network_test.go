package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertDefaultVpcExists(t *testing.T) {
	if !(*runFunctional) {
		t.Skip("skipping functional tests")
	}
	e := getEC2(t, testControllerID)
	_, _, err := e.assertVPCExists("default")
	assert.Nil(t, err)
}

func TestGetSubnets(t *testing.T) {
	if !(*runFunctional) {
		t.Skip("skipping functional tests")
	}
	e := getEC2(t, testControllerID)
	sn, err := e.GetSubnets()
	assert.NoError(t, err)
	assert.True(t, len(sn) > 0)
	if len(sn) > 0 {
		// Just make sure we're pulling in all the info we need
		assert.True(t, len(sn[0].AZ) > 0)
		assert.True(t, len(sn[0].ID) > 0)
		assert.True(t, len(sn[0].CIDR) > 0)
	}
}
