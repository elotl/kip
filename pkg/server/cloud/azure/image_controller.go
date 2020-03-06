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

package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2017-06-01/storage"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util/controllerqueue"
	"github.com/elotl/cloud-instance-provider/pkg/util/jitter"
	"github.com/uber-go/atomic"
	"k8s.io/klog"
)

const (
	elotlImages      = "elotlimages"
	containerName    = "itzodisks"
	accountName      = "milpastorage"
	blobFormatString = `https://%s.blob.core.windows.net`
)

type ImageController struct {
	az                *AzureClient
	bootImageTags     cloud.BootImageTags
	controllerID      string
	resourceGroupName string
	isSynced          *atomic.Bool
	queue             *controllerqueue.Queue
}

func NewImageController(controllerID string, bootImageTags cloud.BootImageTags, azureClient *AzureClient) *ImageController {
	ic := &ImageController{
		controllerID:      controllerID,
		bootImageTags:     bootImageTags,
		az:                azureClient,
		resourceGroupName: regionalResourceGroupName(azureClient.region),
		isSynced:          atomic.NewBool(false),
	}
	ic.queue = controllerqueue.New("image", ic.syncSingleBlobFromQueue)
	return ic
}

func (ic *ImageController) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	ic.queue.Start(quit)
	go ic.FullSyncLoop(quit, wg)
}

func (ic *ImageController) FullSyncLoop(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	// We might have multiple controllers running in the same
	// region, prevent a thundering herd of controllers trying
	// to download and create images
	fullSyncTicker := jitter.NewTicker(20*time.Minute, 10*time.Minute)
	defer fullSyncTicker.Stop()
	if err := ic.syncBestBlob(); err != nil {
		klog.Errorln("Error doing full sync of image controller:", err)
	}
	for {
		select {
		case <-fullSyncTicker.C:
			if err := ic.syncBestBlob(); err != nil {
				klog.Errorln("Error doing full sync of image controller", err)
			}
		case <-quit:
			klog.V(2).Info("Exiting Azure Image Controller Sync Loop")
			return
		}
	}
}

func (ic *ImageController) CreateStorageAccount() (storage.Account, error) {
	var s storage.Account
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	result, err := ic.az.storage.CheckNameAvailability(
		timeoutCtx,
		storage.AccountCheckNameAvailabilityParameters{
			Name: to.StringPtr(accountName),
			Type: to.StringPtr("Microsoft.Storage/storageAccounts"),
		})
	if err != nil {
		return s, fmt.Errorf(
			"Storage account check-name-availability failed: %v\n", err)
	}
	if *result.NameAvailable != true {
		return s, fmt.Errorf(
			"Storage account name %s not available: %v\nserver message: %v\n",
			accountName, err, *result.Message)
	}
	timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	future, err := ic.az.storage.Create(
		timeoutCtx,
		ic.resourceGroupName,
		accountName,
		storage.AccountCreateParameters{
			Sku: &storage.Sku{
				Name: storage.StandardLRS,
			},
			Kind:     storage.Storage,
			Location: to.StringPtr(ic.az.region),
			AccountPropertiesCreateParameters: &storage.AccountPropertiesCreateParameters{},
			Tags: map[string]*string{
				"Created By":           to.StringPtr("milpa"),
				cloud.ControllerTagKey: to.StringPtr(ic.controllerID),
			},
		})
	if err != nil {
		return s, fmt.Errorf("Failed to start creating storage account: %v\n", err)
	}
	timeoutCtx, cancel = context.WithTimeout(ctx, azureWaitTimeout)
	defer cancel()
	err = future.WaitForCompletionRef(timeoutCtx, ic.az.storage.Client)
	if err != nil {
		return s, fmt.Errorf("Failed to finish creating storage account: %v\n", err)
	}
	return future.Result(ic.az.storage)
}

func (ic *ImageController) syncBestBlob() error {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	_, err := ic.az.storage.GetProperties(
		timeoutCtx, ic.resourceGroupName, accountName)
	if err != nil {
		if !isNotFoundError(err) {
			return err
		}
		_, err = ic.CreateStorageAccount()
		if err != nil {
			return err
		}
	}
	container := ic.getPublicContainerURL(elotlImages, containerName)
	images := make([]cloud.Image, 0, 1000)
	for marker := (azblob.Marker{}); marker.NotDone(); {
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		blobSegment, err := container.ListBlobsFlatSegment(
			timeoutCtx,
			marker,
			azblob.ListBlobsSegmentOptions{
				Details: azblob.BlobListingDetails{
					Snapshots: false,
				},
			})
		if err != nil {
			return err
		}
		marker = blobSegment.NextMarker
		for _, item := range blobSegment.Segment.BlobItems {
			name := getImageBasename(item.Name)
			images = append(images, cloud.Image{
				Id:   item.Name,
				Name: name,
			})
		}
	}
	bestBlobName, err := cloud.GetBestImage(images, ic.bootImageTags)
	if err != nil {
		return err
	}
	ic.queue.Add(bestBlobName)
	return nil
}

func getImageBasename(blobName string) string {
	return strings.Split(blobName, ".")[0]
}

func (ic *ImageController) Dump() []byte {
	dumpStruct := struct {
		WorkQueueLength   int
		ControllerID      string
		StorageAccount    string
		BootImageTags     cloud.BootImageTags
		ResourceGroupName string
	}{
		WorkQueueLength:   ic.queue.Len(),
		ControllerID:      ic.controllerID,
		StorageAccount:    accountName,
		BootImageTags:     ic.bootImageTags,
		ResourceGroupName: ic.resourceGroupName,
	}
	b, err := json.MarshalIndent(dumpStruct, "", "    ")
	if err != nil {
		klog.Errorln("Error dumping data from Azure Image Controller", err)
		return nil
	}
	return b
}

func (ic *ImageController) syncSingleBlobFromQueue(ikey interface{}) error {
	err := ic.syncSingleBlob(ikey.(string))
	ic.isSynced.Store(true)
	return err
}

func (ic *ImageController) getAccountPrimaryKey(accountName string) string {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	response, err := ic.az.storage.ListKeys(
		timeoutCtx, ic.resourceGroupName, accountName)
	if err != nil {
		klog.Errorf("Failed to list account keys: %v", err)
		return ""
	}
	return *(((*response.Keys)[0]).Value)
}

func (ic *ImageController) getContainerURL(accountName, containerName string) azblob.ContainerURL {
	key := ic.getAccountPrimaryKey(accountName)
	c, _ := azblob.NewSharedKeyCredential(accountName, key)
	p := azblob.NewPipeline(c, azblob.PipelineOptions{})
	u, _ := url.Parse(fmt.Sprintf(blobFormatString, accountName))
	service := azblob.NewServiceURL(*u, p)
	container := service.NewContainerURL(containerName)
	return container
}

func (ic *ImageController) getPublicContainerURL(accountName, containerName string) azblob.ContainerURL {
	cred := azblob.NewAnonymousCredential()
	p := azblob.NewPipeline(cred, azblob.PipelineOptions{})
	u, _ := url.Parse(fmt.Sprintf(blobFormatString, accountName))
	service := azblob.NewServiceURL(*u, p)
	container := service.NewContainerURL(containerName)
	return container
}

func (ic *ImageController) getBlobURL(accountName, containerName, blobName string) azblob.BlobURL {
	container := ic.getContainerURL(accountName, containerName)
	blob := container.NewBlobURL(blobName)
	return blob
}

func (ic *ImageController) getPublicBlobURL(accountName, containerName, blobName string) azblob.BlobURL {
	container := ic.getPublicContainerURL(accountName, containerName)
	blob := container.NewBlobURL(blobName)
	return blob
}

func isContainerNotFoundError(err error) bool {
	if detailedError, ok := err.(azblob.StorageError); ok {
		sc := detailedError.ServiceCode()
		return sc == azblob.ServiceCodeContainerNotFound
	}
	return false
}

func isBlobNotFoundError(err error) bool {
	if detailedError, ok := err.(azblob.StorageError); ok {
		sc := detailedError.ServiceCode()
		return sc == azblob.ServiceCodeBlobNotFound
	}
	return false
}

func (ic *ImageController) ensureContainer(accountName, containername string) error {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	container := ic.getContainerURL(accountName, containerName)
	_, err := container.GetProperties(
		timeoutCtx, azblob.LeaseAccessConditions{})
	if err != nil {
		if !isContainerNotFoundError(err) {
			klog.Errorf("Checking container %s failed: %v", containerName, err)
			return err
		}
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		_, err = container.Create(
			timeoutCtx, azblob.Metadata{}, azblob.PublicAccessNone)
		if err != nil {
			klog.Errorf("Creating container %s failed: %v", containerName, err)
			return err
		}
	}
	return nil
}

func (ic *ImageController) copyBlob(accountName, containerName, blobName string) (string, error) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	srcBlob := ic.getPublicBlobURL(elotlImages, containerName, blobName)
	dstBlob := ic.getBlobURL(accountName, containerName, blobName)
	_, err := dstBlob.StartCopyFromURL(
		timeoutCtx,
		srcBlob.URL(),
		azblob.Metadata{},
		azblob.ModifiedAccessConditions{},
		azblob.BlobAccessConditions{})
	if err != nil {
		return "", err
	}
	for copyStatus := azblob.CopyStatusPending; copyStatus != azblob.CopyStatusSuccess; {
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		props, err := dstBlob.GetProperties(
			timeoutCtx, azblob.BlobAccessConditions{})
		if err != nil {
			return "", err
		}
		copyStatus = props.CopyStatus()
		if copyStatus == azblob.CopyStatusAborted || copyStatus == azblob.CopyStatusFailed {
			return "", fmt.Errorf("Copying blob %s failed", blobName)
		}
		time.Sleep(3 * time.Second)
	}
	klog.V(2).Infof("Copying blob %s finished", blobName)
	return dstBlob.String(), nil
}

// Azure is currently rolling out AZs to locations.  If we create
// an image without zone resiliancy and then the location enables
// AZs then we need to re-create the image.
func (ic *ImageController) imageParametersMatch(img compute.Image) bool {
	if ic.az.CloudStatusKeeper().SupportsAvailabilityZones() {
		if img.ImageProperties == nil ||
			img.ImageProperties.StorageProfile == nil ||
			!to.Bool(img.ImageProperties.StorageProfile.ZoneResilient) {
			return false
		}
	}
	return true
}

func (ic *ImageController) syncSingleBlob(blobName string) error {
	ctx := context.Background()
	err := ic.ensureContainer(accountName, containerName)
	if err != nil {
		klog.Errorf("Error checking container %s: %v", containerName, err)
		return err
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	url := ""
	blob := ic.getBlobURL(accountName, containerName, blobName)
	_, err = blob.GetProperties(timeoutCtx, azblob.BlobAccessConditions{})
	if err != nil {
		if !isBlobNotFoundError(err) {
			klog.Errorf("Error checking blob %s: %v", blobName, err)
			return err
		}
		url, err = ic.copyBlob(accountName, containerName, blobName)
		if err != nil {
			klog.Errorf("Error copying blob %s: %v", blobName, err)
			return err
		}
	} else {
		url = blob.String()
	}
	imageName := getImageBasename(blobName)
	timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	img, err := ic.az.images.Get(timeoutCtx, ic.resourceGroupName, imageName, "")
	if err != nil && !isNotFoundError(err) {
		klog.Errorf("Error checking image %s: %v", imageName, err)
		return err
	} else if err == nil {
		if ic.imageParametersMatch(img) {
			// Image already exists in the right region with correct
			// params, everything is a-ok.
			return nil
		}
		// Old image parameters don't match, warn the user that we
		// are recreating the image
		klog.Warningln("Image parameters are out of sync, recreating image")
		timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
		defer cancel()
		future, err := ic.az.images.Delete(
			timeoutCtx, ic.resourceGroupName, imageName)
		if err != nil {
			return err
		}
		timeoutCtx, cancel = context.WithTimeout(ctx, azureWaitTimeout)
		defer cancel()
		err = future.WaitForCompletionRef(timeoutCtx, ic.az.images.Client)
		if err != nil {
			return fmt.Errorf("Failed to finish removing image %s: %v\n", imageName, err)
		}
		_, err = future.Result(ic.az.images)
		if err != nil {
			return fmt.Errorf("Failed to finish removing image %s: %v\n", imageName, err)
		}
	}

	cloudStatus := ic.az.CloudStatusKeeper()
	locationSupportsAZs := cloudStatus.SupportsAvailabilityZones()

	timeoutCtx, cancel = context.WithTimeout(ctx, azureDefaultTimeout)
	defer cancel()
	future, err := ic.az.images.CreateOrUpdate(
		timeoutCtx,
		ic.resourceGroupName,
		imageName,
		compute.Image{
			Name:     to.StringPtr(imageName),
			Location: to.StringPtr(ic.az.region),
			ImageProperties: &compute.ImageProperties{
				StorageProfile: &compute.ImageStorageProfile{
					OsDisk: &compute.ImageOSDisk{
						OsType:  compute.Linux,
						BlobURI: to.StringPtr(url),
					},
					ZoneResilient: to.BoolPtr(locationSupportsAZs),
				},
			},
			Tags: map[string]*string{
				cloud.ControllerTagKey: to.StringPtr(ic.controllerID),
			},
		})
	if err != nil {
		return fmt.Errorf("Failed to start creating image %s: %v\n",
			imageName, err)
	}
	timeoutCtx, cancel = context.WithTimeout(ctx, azureWaitTimeout)
	defer cancel()
	err = future.WaitForCompletionRef(timeoutCtx, ic.az.images.Client)
	if err != nil {
		return fmt.Errorf("Failed to finish creating image %s: %v\n",
			imageName, err)
	}
	_, err = future.Result(ic.az.images)
	if err != nil {
		return fmt.Errorf("Failed to finish creating image %s: %v\n",
			imageName, err)
	}
	klog.V(2).Infof("Created image %s from blob %s", imageName, url)
	return nil
}

func (ic *ImageController) WaitForAvailable() {
	for !ic.isSynced.Load() {
		klog.V(2).Infoln("Waiting for azure disk image to sync")
		time.Sleep(3 * time.Second)
	}
	klog.V(2).Infoln("Image synced")
}
