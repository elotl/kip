package azure

import (
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
)

func (az *AzureClient) EnsureContainerInstanceCluster() error {
	return fmt.Errorf("EnsureContainerInstanceCluster is not implemented for azure")
}

func (az *AzureClient) ListContainerInstances() ([]cloud.ContainerInstance, error) {
	return nil, fmt.Errorf("ListContainerInstances is not implemented for azure")
}

func (az *AzureClient) ListContainerInstancesFilterID(instIDs []string) ([]cloud.ContainerInstance, error) {
	return nil, fmt.Errorf("ListContainerInstances is not implemented for azure")
}

func (az *AzureClient) GetContainerInstancesStatuses(instIDs []string) (map[string][]api.UnitStatus, error) {
	return nil, fmt.Errorf("GetContainerInstancesStatus is not implemented for azure")
}

func (az *AzureClient) StartContainerInstance(pod *api.Pod) (string, error) {
	return "", fmt.Errorf("StartContainerInstance is not implemented for azure")
}

func (az *AzureClient) StopContainerInstance(instID string) error {
	return fmt.Errorf("StopContainerInstance is not implemented for azure")
}

func (az *AzureClient) WaitForContainerInstanceRunning(pod *api.Pod) (*api.Pod, error) {
	return nil, fmt.Errorf("WaitForContainerInstanceRunning is not implemented for azure")
}
