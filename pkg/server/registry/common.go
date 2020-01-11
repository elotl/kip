package registry

import "github.com/elotl/cloud-instance-provider/pkg/api"

func copyObjectMetaForUpdate(dest *api.ObjectMeta, src *api.ObjectMeta) {
	dest.Labels = src.Labels
	dest.Annotations = src.Annotations
}
