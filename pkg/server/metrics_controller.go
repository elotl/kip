package server

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util/sets"
	"github.com/golang/glog"
)

const (
	cleanMetricsInterval = 30 * time.Second
)

type MetricsController struct {
	metricsRegistry *registry.MetricsRegistry
	podLister       registry.PodLister
}

func (c *MetricsController) Dump() []byte {
	dumpStruct := struct {
		NumPods int
	}{
		NumPods: len(c.metricsRegistry.ListPods()),
	}
	b, err := json.MarshalIndent(dumpStruct, "", "    ")
	if err != nil {
		glog.Errorln("Error dumping data from metrics controller", err)
		return nil
	}
	return b
}

func (c *MetricsController) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	go c.runSyncLoop(quit, wg)
}

func (c *MetricsController) runSyncLoop(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	cleanTicker := time.NewTicker(cleanMetricsInterval)
	for {
		select {
		case <-cleanTicker.C:
			err := c.cleanOldMetrics()
			if err != nil {
				glog.Errorf("Error cleaning old metrics: %s", err)
			}
		case <-quit:
			cleanTicker.Stop()
			glog.Info("Exiting MetricsController Sync Loop")
			return
		}
	}
}

func (c *MetricsController) cleanOldMetrics() error {
	pods, err := c.podLister.ListPods(registry.MatchAllLivePods)
	if err != nil {
		return err
	}
	podNameSet := sets.NewString()
	for i := range pods.Items {
		podNameSet.Insert(pods.Items[i].Name)
	}
	metricPodNames := c.metricsRegistry.ListPods()
	metricNameSet := sets.NewString(metricPodNames...)
	oldMetrics := metricNameSet.Difference(podNameSet)
	c.metricsRegistry.DeletePods(oldMetrics.List()...)
	return nil
}
