package gce

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getGCE(t *testing.T, controllerID string) *gceClient {
	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/bcox/google/milpa-207719-fb309bd448d5.json")
	assert.NoError(t, err)
	c, err := NewGCEClient(controllerID, "bcoxtest", "milpa-207719", WithVPCName("default"), WithZone("us-central1-a"), WithSubnetName("default"))
	assert.NoError(t, err)
	return c
}

func TestGCECloud(t *testing.T) {
	fmt.Printf("Running Cloud Test\n")
	controllerID := "bcoxtestcontroller"
	cloudClient := getGCE(t, controllerID)
	fmt.Println(cloudClient)
	err := cloudClient.EnsureMilpaSecurityGroups(
		[]string{},
		[]string{},
	)
	assert.NoError(t, err)

}
