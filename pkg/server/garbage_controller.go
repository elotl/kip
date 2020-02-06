package server

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud/aws"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud/azure"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util/sets"
	"github.com/elotl/cloud-instance-provider/pkg/util/stats"
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
	lastOldTaskDefs         sets.String
}

func (c *GarbageController) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	c.lastOrphanedAzureGroups = sets.NewString()
	go c.GCLoop(quit, wg)
}

func (c *GarbageController) Dump() []byte {
	b, err := json.MarshalIndent(c.timer, "", "    ")
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
	for iid, _ := range unknownInstances {
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

func (c *GarbageController) cleanFargateTaskDefs() {
	client, ok := c.cloudClient.(*aws.AwsEC2)
	if !ok {
		return
	}
	if c.lastOldTaskDefs == nil {
		c.lastOldTaskDefs = sets.NewString()
	}
	taskDefARNs, err := client.ListTaskDefinitions()
	if err != nil {
		klog.Errorln("Error listing ECS Fargate task definitions for cleanup:", err)
		return
	}
	if len(taskDefARNs) == 0 {
		return
	}
	pods, err := c.podRegistry.ListPods(registry.MatchAllPods)
	if err != nil {
		klog.Errorln("Error listing pods when cleaning up fargate task definitions", err)
	}
	podNames := sets.NewString()
	for i := range pods.Items {
		podNames.Insert(pods.Items[i].Name)
	}
	oldTaskDefs := getOutdatedTaskDefinitions(taskDefARNs, podNames, c.controllerID)
	doomedTaskDefARNs := c.lastOldTaskDefs.Intersection(oldTaskDefs)
	for _, taskDefARN := range doomedTaskDefARNs.List() {
		err := client.DeregisterTaskDefinition(taskDefARN)
		if err != nil {
			klog.Errorln("Error cleaning up old task definition", taskDefARN)
		}
	}
	c.lastOldTaskDefs = oldTaskDefs
}

// Task definitions have a taskName and taskRevision in the form
// taskname:revision.  We should be deactivating all but the most
// recent task definions Here we find all the older task definition
// revisions.
func getOutdatedTaskDefinitions(taskDefARNs []string, podNames sets.String, controllerID string) sets.String {
	oldTaskDefs := sets.NewString()
	latestRevision := make(map[string]int)
	for _, taskDef := range taskDefARNs {
		name, revision := aws.SplitTaskDef(taskDef, controllerID)
		if len(name) == 0 || revision == 0 {
			continue
		}
		if revision > latestRevision[name] {
			latestRevision[name] = revision
		}
	}
	for _, taskDef := range taskDefARNs {
		name, revision := aws.SplitTaskDef(taskDef, controllerID)
		if !podNames.Has(name) {
			oldTaskDefs.Insert(taskDef)
		} else {
			if revision < latestRevision[name] {
				oldTaskDefs.Insert(taskDef)
			}
		}
	}
	return oldTaskDefs
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
