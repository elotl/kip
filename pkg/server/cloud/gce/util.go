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

package gce

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/api/googleapi"
)

const (
	gceComputeAPIEndpoint = "https://www.googleapis.com/compute/v1/"
)

func isNotFoundError(err error) bool {
	return isHTTPErrorCode(err, http.StatusNotFound)
}

func isHTTPErrorCode(err error, code int) bool {
	apiErr, ok := err.(*googleapi.Error)
	return ok && apiErr.Code == code
}

func CreateKipCellNetworkTag(controllerID string) string {
	return fmt.Sprintf("kip-%s", controllerID)
}

func (c *gceClient) getNetworkURL() string {
	return gceComputeAPIEndpoint + strings.Join([]string{"projects", c.projectID, "global", "networks", c.vpcName}, "/")
}

func (c *gceClient) getProjectURL() string {
	return gceComputeAPIEndpoint + strings.Join([]string{"projects", c.projectID}, "/")
}
