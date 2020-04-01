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

package registry

import (
	"fmt"

	"github.com/elotl/kip/pkg/api"
)

func checkObjectMetaForUpdate(dest *api.ObjectMeta, src *api.ObjectMeta) error {
	if !src.CreationTimestamp.IsZero() && !src.CreationTimestamp.Equal(dest.CreationTimestamp) {
		return fmt.Errorf("CreationTimestamp is an immutable field %v != %v",
			src.CreationTimestamp, dest.CreationTimestamp)
	}
	if src.UID != "" && src.UID != dest.UID {
		return fmt.Errorf("UID is an immutable field %v != %v",
			src.UID, dest.UID)
	}
	return nil
}

func copyObjectMetaForUpdate(dest *api.ObjectMeta, src *api.ObjectMeta) {
	dest.Labels = src.Labels
	dest.Annotations = src.Annotations
}
