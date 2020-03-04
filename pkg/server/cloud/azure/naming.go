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
