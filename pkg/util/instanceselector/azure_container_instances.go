package instanceselector

import (
	"fmt"

	"github.com/elotl/cloud-instance-provider/pkg/api"
)

func AzureContainenrInstanceSelector(rs *api.ResourceSpec) (int64, int64, error) {
	return 0, 0, fmt.Errorf("Azure container instances is not implemented on milpa")
}
