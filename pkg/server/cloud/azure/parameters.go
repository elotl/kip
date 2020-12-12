package azure

import "fmt"

func (az *AzureClient) AddInstanceParameter(instanceID, name, value string, isSecret bool) error {
	return fmt.Errorf("not implemented")
}

func (az *AzureClient) DeleteInstanceParameter(instanceID, name string) error {
	return fmt.Errorf("not implemented")
}
