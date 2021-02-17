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

type InsufficientCapacityError struct {
	InstanceType  string
	OriginalError string
}

func (e *InsufficientCapacityError) Error() string {
	return fmt.Sprintf("Insufficient capacity of %s instances, err: %s", e.InstanceType, e.OriginalError)
}
