package azure

import (
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/elotl/cloud-instance-provider/pkg/util"
)

func RetryDelete(timeout time.Duration, f func() error) error {
	isRetryable := func(err error) bool {
		return strings.Contains(err.Error(), "still allocated")
	}
	return util.Retry(timeout, f, isRetryable)
}

func makeZoneParam(zone string) *[]string {
	if zone == "" {
		return nil
	}
	return to.StringSlicePtr([]string{zone})
}
