package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNodeResourceGroup(t *testing.T) {
	controllerID := "testcluster"
	nodeName := "45a41997-d1bc-414f-a9b8-8f0cfd33ce17"
	instanceID := makeInstanceID(controllerID, nodeName)
	nodeRG := instanceID
	assert.True(t, isNodeResourceGroup(nodeRG, controllerID))
	clusterRG := controllerResourceGroupName(controllerID)
	assert.False(t, isNodeResourceGroup(clusterRG, controllerID))
}
