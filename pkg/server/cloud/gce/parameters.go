package gce

import "fmt"

func (c *gceClient) AddInstanceParameter(instanceID, name, value string, isSecret bool) error {
	return fmt.Errorf("not implemented")
}

func (c *gceClient) DeleteInstanceParameter(instanceID, name string) error {
	return fmt.Errorf("not implemented")
}
