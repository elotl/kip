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
