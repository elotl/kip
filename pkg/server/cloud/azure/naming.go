package azure

import (
	"strings"

	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/hash"
	uuid "github.com/satori/go.uuid"
)

func regionalResourceGroupName(region string) string {
	region = strings.ToLower(strings.Replace(region, " ", "", -1))
	return util.CreateResourceGroupName(region)
}

func controllerResourceGroupName(controllerID string) string {
	return util.CreateClusterResourceGroupName(controllerID)
}

func makeInstanceID(controllerID, nodeName string) string {
	nodeUUID, err := uuid.FromString(nodeName)
	var compressedID string
	if err != nil {
		compressedID = strings.Replace(nodeName, "-", "", -1)
		if len(compressedID) > 26 {
			compressedID = compressedID[:26]
		}
	} else {
		compressedID = hash.Base32EncodeNoPad(nodeUUID.Bytes())
	}

	prefix := util.CreateClusterResourcePrefix(controllerID)
	return prefix + compressedID
}
