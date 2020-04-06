package healthcheck

import (
	"time"

	"github.com/elotl/kip/pkg/api"
)

const (
	terminateChanSize = 1000
)

var (
	checkPeriod = time.Duration(15 * time.Second)
)

type HealthChecker interface {
	Start()
	TerminatePodsChan() <-chan *api.Pod
}
