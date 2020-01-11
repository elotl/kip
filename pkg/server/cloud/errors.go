package cloud

import "fmt"

type NoCapacityError struct {
	// If both AZ and SubnetID are empty, we have no capacity
	// for this instance in the entire region
	OriginalError string
	AZ            string
	SubnetID      string
}

func (e *NoCapacityError) Error() string {
	if e.AZ != "" {
		return fmt.Sprintf("Availability Zone %s has no capacity: %s", e.AZ, e.OriginalError)
	} else if e.SubnetID != "" {
		return fmt.Sprintf("Subnet %s has no capacity: %s", e.SubnetID, e.OriginalError)
	} else {
		return fmt.Sprintf("Region has no capacity: %s", e.OriginalError)
	}
}

type UnsupportedInstanceError struct {
	OriginalError string
}

func (e *UnsupportedInstanceError) Error() string {
	return fmt.Sprintf("Unsupported spot instance type: %s", e.OriginalError)
}
