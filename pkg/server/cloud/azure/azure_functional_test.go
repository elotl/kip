package azure

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud/functional"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/stretchr/testify/assert"
)

const (
	testControllerID   = "azure-func-test"
	testSubscriptionID = "4e84e89a-b806-4d7d-900b-cae8cb640215"
	testVNetName       = "milpa-tests/milpa-tests-vnet"
	testSubnetName     = "default"
	testRegion         = "East US"
	instanceType       = string(compute.VirtualMachineSizeTypesBasicA0)
)

var runFunctional = flag.Bool("functional", false, "run functional system tests")

func cleanupFromTest(az *AzureClient) error {
	// delete the cluster resource group and any resources in that resource group
	fmt.Println("Tearing down azure functional test")
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	_, err := az.groups.Delete(timeoutCtx, controllerResourceGroupName(az.controllerID))
	return err
}

func syncImage(controllerID string, bootImageTags cloud.BootImageTags, az *AzureClient) {
	fmt.Println("Ensuring image is available for cluster:", controllerID)
	start := time.Now()
	ic := NewImageController(controllerID, bootImageTags, az)
	quit := make(chan struct{})
	wg := &sync.WaitGroup{}
	ic.Start(quit, wg)
	ic.WaitForAvailable()
	close(quit)
	wg.Wait()
	fmt.Printf("Image synchronization finished in %v\n", time.Now().Sub(start))
}

func TestAzureCloud(t *testing.T) {
	if !(*runFunctional) {
		t.Skip("skipping cloud functional tests")
	}
	fmt.Printf("Running Functional Tests\n")
	if !util.AzureEnvVarsSet() {
		t.Fatal("Neet to setup Azure env vars for tests")
	}
	controllerID := ""
	executorNStr := os.Getenv("EXECUTOR_NUMBER")
	if executorNStr != "" {
		controllerID = fmt.Sprintf("%s-%s", testControllerID, executorNStr)
	} else {
		controllerID = api.SimpleNameGenerator.GenerateName(testControllerID)
	}

	az, err := NewAzureClient(controllerID, controllerID, testSubscriptionID, testRegion, testVNetName, testSubnetName)
	if executorNStr == "" {
		// When not running on Jenkins, clean up resources.
		defer cleanupFromTest(az)
	}
	assert.Nil(t, err)
	if err != nil {
		return
	}

	syncImage(controllerID, cloud.BootImageTags{}, az)

	imageID, err := az.GetImageId(cloud.BootImageTags{})
	if err != nil {
		assert.Fail(t, "Azure functional test failed, could not get Image ID: %v", err.Error())
		return
	}
	fmt.Println("found image:", imageID)

	ts, err := functional.SetupCloudFunctionalTest(t, az, imageID, instanceType)
	if err != nil {
		assert.FailNow(t, "Failed to setup functional test: %s", err.Error())
	}
	defer ts.Cleanup(t)
	t.Run("RunGetVMVirtualNetworksTest", func(t *testing.T) {
		RunGetVMNetworksTest(t, ts)
	})
}

func RunGetVMNetworksTest(t *testing.T, ts *functional.TestState) {
	vmResourceGroup := ts.Node1.Status.InstanceID
	vmName := ts.Node1.Status.InstanceID
	az := ts.CloudClient.(*AzureClient)
	vNets, subnets := az.GetVMNetworks(vmResourceGroup, vmName)
	assert.Len(t, vNets, 1)
	if len(vNets) == 1 {
		expectedVNetName := strings.Split(testVNetName, "/")[1]
		assert.Equal(t, expectedVNetName, vNets[0])
	}
	assert.Len(t, subnets, 1)
	if len(subnets) == 1 {
		assert.Equal(t, testSubnetName, subnets[0])
	}
}
