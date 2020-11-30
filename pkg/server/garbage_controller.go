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

package server

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/elotl/kip/pkg/api"
	"github.com/elotl/kip/pkg/server/cloud"
	"github.com/elotl/kip/pkg/server/cloud/azure"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util/stats"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog"
)

var lastUnknownInstances map[string]bool

func init() {
	lastUnknownInstances = make(map[string]bool)
}

type GarbageControllerConfig struct {
	CleanTerminatedInterval time.Duration
	CleanInstancesInterval  time.Duration
}

type GarbageController struct {
	config                  GarbageControllerConfig
	podRegistry             *registry.PodRegistry
	nodeRegistry            *registry.NodeRegistry
	cloudClient             cloud.CloudClient
	controllerID            string
	timer                   stats.LoopTimer
	lastOrphanedAzureGroups sets.String
}

func (c *GarbageController) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	c.lastOrphanedAzureGroups = sets.NewString()
	go c.GCLoop(quit, wg)
}

func (c *GarbageController) Dump() []byte {
	b, err := json.MarshalIndent(c.timer.Copy(), "", "    ")
	if err != nil {
		klog.Errorln("Error dumping data from GarbageController", err)
		return nil
	}
	return b
}

func (c *GarbageController) GCLoop(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	cleanTermiantedTicker := time.NewTicker(c.config.CleanTerminatedInterval)
	instancesTicker := time.NewTicker(c.config.CleanInstancesInterval)
	cleanResourceGroupsTicker := time.NewTicker(3 * time.Minute)
	defer cleanTermiantedTicker.Stop()
	defer instancesTicker.Stop()
	defer cleanResourceGroupsTicker.Stop()
	for {
		// The garbage controller takes a while to stop if we
		// are timing out talking to etcd, lets give quit priority
		select {
		case <-quit:
			klog.V(2).Info("Stopping GarbageController")
			return
		default:
		}
		select {
		case <-instancesTicker.C:
			c.timer.StartLoop()
			c.CleanInstances()
			c.CleanDanglingRoutes()
			c.timer.EndLoop()
		case <-cleanTermiantedTicker.C:
			c.CleanTerminatedNodes()
		case <-cleanResourceGroupsTicker.C:
			c.CleanAzureResourceGroups()
		case <-quit:
			klog.V(2).Info("Stopping GarbageController")
			return
		}
	}
}

func (c *GarbageController) CleanDanglingRoutes() {
	err := c.cloudClient.RemoveRoute("", "")
	if err != nil {
		klog.Warningf("cleaning up dangling routes: %v", err)
	}
}

func (c *GarbageController) CleanTerminatedNodes() {
	nodes, err := c.nodeRegistry.ListNodes(func(n *api.Node) bool {
		return n.Status.Phase == api.NodeTerminated
	})
	if err != nil {
		klog.Errorln("Couldn't list terminated nodes", err)
	}
	now := api.Now()
	for _, node := range nodes.Items {
		if node.DeletionTimestamp == nil {
			klog.Warningf("Found node with nil deletion timestamp")
			_, _ = c.nodeRegistry.SetNodeDeletionTimestamp(node)
		}
		if node.DeletionTimestamp != nil &&
			node.DeletionTimestamp.Add(30*time.Second).Before(now) {
			_, err = c.nodeRegistry.PurgeNode(node)
			if err != nil {
				klog.Errorf("Error purging terminated nodes")
			}
		}
	}
}

func (c *GarbageController) CleanInstances() {
	unknownInstances := make(map[string]bool)
	nodes, err := c.nodeRegistry.ListNodes(registry.MatchAllNodes)
	if err != nil {
		klog.Errorf("Error listing nodes in GC: %s", err.Error())
		return
	}
	nodeSet := make(map[string]bool)
	for _, node := range nodes.Items {
		nodeSet[node.Name] = true
	}
	if err != nil {
		klog.Errorf("Error listing nodes for instance cleaning: %s", err.Error())
		return
	}
	instances, err := c.cloudClient.ListInstances()
	if err != nil {
		klog.Errorf("Error listing cloud instances: %s", err.Error())
		return
	}
	for _, inst := range instances {
		if !nodeSet[inst.NodeName] {
			unknownInstances[inst.ID] = true
		}
	}
	for iid := range unknownInstances {
		if lastUnknownInstances[iid] {
			klog.Errorf("Stopping unknown cloud instance %s", iid)
			go func() {
				err := c.cloudClient.StopInstance(iid)
				if err != nil {
					klog.Error(err)
				}
			}()
		}
	}
	lastUnknownInstances = unknownInstances
}

func (c *GarbageController) CleanAzureResourceGroups() {
	az, ok := c.cloudClient.(*azure.AzureClient)
	if !ok {
		return
	}
	err := c.CleanAzureResourceGroupsHelper(az)
	if err != nil {
		klog.Error(err)
	}
}

type ResourceGrouper interface {
	ListNodeResourceGroups() ([]string, error)
	DeleteResourceGroup(string) error
}

func (c *GarbageController) CleanAzureResourceGroupsHelper(client ResourceGrouper) error {
	nodes, err := c.nodeRegistry.ListNodes(func(n *api.Node) bool {
		return (n.Status.Phase != api.NodeTerminating &&
			n.Status.Phase != api.NodeTerminated)
	})
	if err != nil {
		return fmt.Errorf("Error listing nodes for cleaning azure resource groups: %s", err.Error())
	}
	expectedGroups := sets.NewString()
	for _, n := range nodes.Items {
		if n.Status.InstanceID != "" {
			expectedGroups.Insert(n.Status.InstanceID)
		}
	}

	orphanedGroups := sets.NewString()
	groups, err := client.ListNodeResourceGroups()
	if err != nil {
		return fmt.Errorf("Error listing azure resource groups, not cleaning orphaned groups: %s", err.Error())
	}
	for _, group := range groups {
		if !expectedGroups.Has(group) {
			orphanedGroups.Insert(group)
		}
	}

	doomedGroups := c.lastOrphanedAzureGroups.Intersection(orphanedGroups)
	newOrphaned := orphanedGroups.Difference(doomedGroups)
	c.lastOrphanedAzureGroups = newOrphaned

	if doomedGroups.Len() > 0 {
		klog.Errorf("Deleting %d orphaned azure resource groups: %v", doomedGroups.Len(), doomedGroups.List())
	}
	for _, groupName := range doomedGroups.List() {
		err := client.DeleteResourceGroup(groupName)
		if err != nil {
			klog.Errorf("Error deleting orphaned resource group: %s", err.Error())
		}
	}
	return nil
}
