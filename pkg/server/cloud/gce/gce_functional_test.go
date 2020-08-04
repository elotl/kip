package gce

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud/functional"
	"github.com/elotl/kip/pkg/util"
	"github.com/stretchr/testify/assert"
)

const (
	testControllerID = "gcefunctionaltest"
	testProjectID    = "elotl-dev"
	testZone         = "us-west1-b"
	vpcID            = "default"
	subnetName       = "default"
	bootImageID      = "https://www.googleapis.com/compute/v1/projects/elotl-kip/global/images/elotl-kip-latest"
	rootDevice       = ""
	instanceType     = "f1-micro"
)

var runFunctional = flag.Bool("functional", false, "run functional system tests")

func getGCEClient(t *testing.T, controllerID string) *gceClient {
	if !util.GCEEnvVarsSet() {
		t.Fatal("Neet to setup AWS env vars for tests")
	}
	clientEmail := os.Getenv("GCE_CLIENT_EMAIL")
	privateKey := os.Getenv("GCE_PRIVATE_KEY")
	gceClient, err := NewGCEClient(controllerID, controllerID, testProjectID,
		WithZone(testZone), WithVPCName(vpcID), WithSubnetName(subnetName),
		WithCredentials(clientEmail, privateKey))
	if err != nil {
		msg := "Error setting up GCE client: " + err.Error()
		assert.FailNow(t, msg)
	}
	return gceClient
}

// Each functional test run creates a unique securty group that we need to delete
func ensureSecurityGroupDeleted(c *gceClient) error {
	apiGroupName := CreateKipCellNetworkTag(c.controllerID)
	sg, err := c.FindSecurityGroup(apiGroupName)
	if err != nil {
		return util.WrapError(err, "Error finding security group")
	}
	if sg == nil {
		fmt.Println("No security group found, not deleting group")
		return nil
	}
	err = c.DeleteSecurityGroup(apiGroupName)
	return err
}

func TestGCECloud(t *testing.T) {
	if !(*runFunctional) {
		t.Skip("skipping gce cloud functional tests")
	}
	fmt.Printf("Running GCE Functional Tests\n")
	if !util.GCEEnvVarsSet() {
		t.Fatal("Neet to setup GCE env vars for tests")
	}

	controllerID := api.SimpleNameGenerator.GenerateName(testControllerID)
	fmt.Println("GCE Controller ID:", controllerID)
	gceClient := getGCEClient(t, controllerID)

	defer func() {
		err := ensureSecurityGroupDeleted(gceClient)
		if err != nil {
			assert.FailNow(t, "Failed to delete cell security group")
		}
	}()

	ts, err := functional.SetupCloudFunctionalTest(
		t, gceClient, bootImageID, rootDevice, instanceType)
	if err != nil {
		assert.FailNow(t, "Failed to setup functional test: %s", err.Error())
	}
	defer ts.Cleanup(t)
}
