package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/elotl/kip/pkg/api"
	v1beta1 "github.com/elotl/kip/pkg/apis/kip/v1beta1"
	kv1b1 "github.com/elotl/kip/pkg/k8sclient/clientset/versioned/typed/kip/v1beta1"
	"github.com/elotl/kip/pkg/server/events"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util"
	"github.com/elotl/kip/pkg/util/controllerqueue"
	"github.com/elotl/kip/pkg/util/yaml"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
)

const (
	cellNodeCreated = iota
	cellNodeUpdated
	cellNodeDeleted
	cellsFullSync
)

type CellController struct {
	controllerID  string
	nodeName      string
	k8sRestConfig *rest.Config
	k8sKipClient  kv1b1.CellInterface
	eventsSystem  *events.EventSystem
	podLister     registry.PodLister
	nodeLister    registry.NodeLister
	queue         *controllerqueue.Queue
}

type CellOp struct {
	op   int
	node *api.Node
	pod  *api.Pod
}

func NewCellController(
	controllerID, nodeName string,
	restConfig *rest.Config,
	k8sKipClient kv1b1.CellInterface,
	eventsSystem *events.EventSystem,
	podLister registry.PodLister,
	nodeLister registry.NodeLister,
) (*CellController, error) {
	c := &CellController{
		controllerID:  controllerID,
		nodeName:      nodeName,
		k8sRestConfig: restConfig,
		k8sKipClient:  k8sKipClient,
		eventsSystem:  eventsSystem,
		podLister:     podLister,
		nodeLister:    nodeLister,
	}
	err := c.CreateCRDIfNotExists()
	if err != nil {
		return nil, util.WrapError(
			err, "Could not create new Cell Controller")
	}

	c.queue = controllerqueue.New("cells", c.processOperation, controllerqueue.NumWorkers(1), controllerqueue.MaxRetries(0))
	return c, nil
}

func (c *CellController) Start(quit <-chan struct{}, wg *sync.WaitGroup) {
	c.queue.Start(quit)
	go c.runSyncLoop(quit, wg)
	c.registerEventHandlers()
}

func (c *CellController) registerEventHandlers() {
	c.eventsSystem.RegisterHandlerFunc(
		events.NodeCreated, c.nodeCreated)
	c.eventsSystem.RegisterHandlerFunc(
		events.PodRunning, c.podRunning)
	c.eventsSystem.RegisterHandlerFunc(
		events.NodePurged, c.nodePurged)
}

func (c *CellController) nodeCreated(e events.Event) error {
	node, ok := e.Object.(*api.Node)
	if !ok {
		return fmt.Errorf("invalid event object: %v", e)
	}
	c.queue.Add(CellOp{
		op:   cellNodeCreated,
		node: node,
	})
	return nil
}

func (c *CellController) podRunning(e events.Event) error {
	pod, ok := e.Object.(*api.Pod)
	if !ok {
		return fmt.Errorf("invalid event object: %v", e)
	}
	node, err := c.nodeLister.GetNode(pod.Status.BoundNodeName)
	if err != nil {
		return fmt.Errorf("could not look up node for running pod %s: %s",
			pod.Name, err)
	}
	c.queue.Add(CellOp{
		op:   cellNodeUpdated,
		node: node,
		pod:  pod,
	})

	return nil
}

func (c *CellController) nodePurged(e events.Event) error {
	node, ok := e.Object.(*api.Node)
	if !ok {
		return fmt.Errorf("Invalid event object: %v", e)
	}
	c.queue.Add(CellOp{
		op:   cellNodeDeleted,
		node: node,
	})
	return nil
}

func (c *CellController) Dump() []byte {
	dumpStruct := struct {
		WorkQueueLength int
	}{
		WorkQueueLength: c.queue.Len(),
	}
	b, err := json.MarshalIndent(dumpStruct, "", "    ")
	if err != nil {
		klog.Errorln("Error dumping data from CellsController", err)
		return nil
	}
	return b
}

func (c *CellController) runSyncLoop(quit <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	fullSyncTicker := time.NewTicker(45 * time.Second)
	c.queue.Add(CellOp{op: cellsFullSync})
	for {
		select {
		case <-fullSyncTicker.C:
			c.queue.Add(CellOp{op: cellsFullSync})
		case <-quit:
			klog.Info("Exiting CellController Sync Loop")
			return
		}
	}
}

func (c *CellController) processOperation(item interface{}) error {
	kNodeOp, ok := item.(CellOp)
	if !ok {
		return fmt.Errorf("Incrorrect item type inserted into work queue: %v",
			item)
	}
	var err error
	switch kNodeOp.op {
	case cellNodeCreated:
		err = c.createNodeCell(kNodeOp.node)
	case cellNodeUpdated:
		err = c.updateCell(kNodeOp.node, kNodeOp.pod)
	case cellNodeDeleted:
		err = c.deleteNodeCell(kNodeOp.node)
	case cellsFullSync:
		err = c.syncAllCells()
	}
	if err != nil {
		klog.Errorf("Error processing cell operation: %s", err)
	}
	return err
}

func (c *CellController) makeNodeCell(n *api.Node, p *api.Pod) *v1beta1.Cell {
	kn := &v1beta1.Cell{}
	if n != nil {
		kn.Name = n.Name
		kn.Status.ControllerID = c.controllerID
		kn.Status.Node = c.nodeName
		kn.Status.LaunchType = "On-Demand"
		if n.Spec.Spot {
			kn.Status.LaunchType = "Spot"
		}
		kn.Status.InstanceType = n.Spec.InstanceType
		kn.Status.InstanceID = n.Status.InstanceID
		kn.Status.IP = api.GetPrivateIP(n.Status.Addresses)
		podNamespace, podName := util.SplitNamespaceAndName(n.Status.BoundPodName)
		kn.Status.PodName = podName
		kn.Status.PodNamespace = podNamespace
	}
	if p != nil {
		podNamespace, podName := util.SplitNamespaceAndName(p.Name)
		kn.Status.PodName = podName
		kn.Status.PodNamespace = podNamespace
		kn.Labels = p.Labels
	}
	return kn
}

func (c *CellController) makeCell(n *api.Node, p *api.Pod) *v1beta1.Cell {
	return c.makeNodeCell(n, p)
}

func (c *CellController) createNodeCell(n *api.Node) error {
	if n == nil {
		return fmt.Errorf("Could not create cell record: invalid kip node")
	}
	kn := c.makeNodeCell(n, nil)
	_, err := c.k8sKipClient.Create(kn)
	return err
}

func (c *CellController) updateCell(n *api.Node, p *api.Pod) error {
	cell := c.makeCell(n, p)
	err := c.updateK8sCell(cell)
	if err != nil {
		return err
	}
	return nil
}

func (c *CellController) updateK8sCell(cell *v1beta1.Cell) error {
	kn, err := c.k8sKipClient.Get(cell.Name, metav1.GetOptions{})
	if err != nil {
		return util.WrapError(err, "Error getting cell from k8s for updating")
	}
	kn.Labels = cell.Labels
	kn.Status.PodName = cell.Status.PodName
	kn.Status.PodNamespace = cell.Status.PodNamespace
	kn.Status.IP = cell.Status.IP
	_, err = c.k8sKipClient.Update(kn)
	if err != nil {
		return util.WrapError(err, "Error updating cell record in k8s")
	}
	return nil
}

func (c *CellController) deleteNodeCell(n *api.Node) error {
	if n == nil {
		return fmt.Errorf("Could not delete cell record: no corresponding kip node provided")
	}
	return c.k8sKipClient.Delete(n.Name, &metav1.DeleteOptions{})
}

func (c *CellController) syncAllCells() error {
	// pull all currently running nodes
	// I think we can just do this with nodes only
	nodes, err := c.nodeLister.ListNodes(func(n *api.Node) bool {
		switch n.Status.Phase {
		case api.NodeCreated, api.NodeAvailable, api.NodeClaimed, api.NodeCleaning, api.NodeTerminating:
			return true
		}
		return false
	})
	if err != nil {
		return util.WrapError(err, "Could not load kip nodes for full sync")
	}

	pods, err := c.podLister.ListPods(registry.MatchAllPods)
	if err != nil {
		return util.WrapError(err, "Could not load kip pods for full sync")
	}
	podMap := make(map[string]*api.Pod)
	for _, p := range pods.Items {
		podMap[p.Name] = p
	}

	// pull the current list of cells from k8s
	kCellList, err := c.k8sKipClient.List(metav1.ListOptions{})
	if err != nil {
		return util.WrapError(err, "Could not load cells from k8s for full sync")
	}
	// from the pods and nodes, create a list of cells we should have,
	// store in map. Our regular node cells can come from the nodes,
	// our container instance cells come from pods.
	specCells := make(map[string]interface{})
	for _, n := range nodes.Items {
		var pod *api.Pod
		if n.Status.BoundPodName != "" {
			pod = podMap[n.Status.BoundPodName]
		}
		cell := c.makeCell(n, pod)
		specCells[cell.Name] = cell
	}

	statusCells := make(map[string]interface{})
	for i, cell := range kCellList.Items {
		if kCellList.Items[i].Status.ControllerID == c.controllerID {
			// Bug to watch out for: Don't store pointer to local loop var
			statusCells[cell.Name] = &kCellList.Items[i]
		}
	}

	add, update, delete := util.MapUserDiff(specCells, statusCells, diffCells)
	if len(add) > 0 || len(update) > 0 || len(delete) > 0 {
		klog.V(3).Infof("reconciling cell records - add: %d, update: %d, delete: %d",
			len(add), len(update), len(delete))
	}
	errs := make([]error, 0)
	for _, cellName := range add {
		cellIface, exists := specCells[cellName]
		if !exists {
			// should never happen
			klog.Errorf("Error in diff: got unknown cell (add): %s", cellName)
			continue
		}
		cell := cellIface.(*v1beta1.Cell)
		_, err := c.k8sKipClient.Create(cell)
		if err != nil {
			errs = append(errs, util.WrapError(err, "Could not create cell in k8s"))
		}
	}

	for _, cellName := range update {
		specCellIface, exists := specCells[cellName]
		if !exists {
			// should never happen
			klog.Errorf("Error in diff: got unknown cell (update): %s", cellName)
			continue
		}
		specCell := specCellIface.(*v1beta1.Cell)
		statusCellIface, exists := statusCells[cellName]
		if !exists {
			klog.Errorf("Error in diff: got unknown cell (update): %s", cellName)
		}
		statusCell := statusCellIface.(*v1beta1.Cell)
		statusCell.Status = specCell.Status
		statusCell.Labels = specCell.Labels
		_, err := c.k8sKipClient.Update(statusCell)
		if err != nil {
			errs = append(errs, util.WrapError(err, "Could not update cell in k8s"))
		}
	}

	for _, cellName := range delete {
		err := c.k8sKipClient.Delete(cellName, &metav1.DeleteOptions{})
		if err != nil {
			errs = append(errs, util.WrapError(err, "Error deleting cell in k8s"))
		}
	}

	return errors.NewAggregate(errs)
}

func diffCells(a, b interface{}) bool {
	aNode := a.(*v1beta1.Cell)
	bNode := b.(*v1beta1.Cell)
	return (aNode.Status == bNode.Status &&
		reflect.DeepEqual(aNode.Labels, bNode.Labels))
}

func (c *CellController) CreateCRDIfNotExists() error {
	crdDef := apiextensionsv1beta1.CustomResourceDefinition{}
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(v1beta1.KipCRDDefString)), 16000)
	err := decoder.Decode(&crdDef)
	if err != nil {
		return util.WrapError(err, "Error creating k8s custom resource def: could not decode: %s", err.Error())
	}
	clientset, err := apiextensionsclientset.NewForConfig(c.k8sRestConfig)
	if err != nil {
		return util.WrapError(err, "Error creating k8s custom resource client for Cells")
	}

	crdName := crdDef.Name
	_, err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crdName, metav1.GetOptions{})
	if err != nil {
		// Create if it's not found
		if apierrors.IsNotFound(err) {
			klog.V(3).Infof("Creating cells CRD in k8s")
			_, err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(&crdDef)
			if err != nil {
				if apierrors.IsAlreadyExists(err) {
					// probably a race between 2 controllers, We're OK
					// with that.  We can probably assume that we
					// don't need to update it either.
					klog.V(3).Infof("Cells CRD is already registered in k8s")
				} else {
					return util.WrapError(err, "Error ensuring k8s cell CRD exists in controller")
				}
			}
		} else {
			// Something else went wrong, lets get out of here
			return util.WrapError(err, "Error ensuring k8s cell CRD exists in controller")
		}
	}

	err = waitForCRDToExist(clientset, crdName)
	if err != nil {
		return util.WrapError(err, "Error waiting for cell CRD to exist in k8s controller")
	}
	return nil
}

func waitForCRDToExist(clientset apiextensionsclientset.Interface, crdName string) error {
	err := wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crdName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, cond := range crd.Status.Conditions {
			switch cond.Type {
			case apiextensionsv1beta1.Established:
				if cond.Status == apiextensionsv1beta1.ConditionTrue {
					return true, err
				}
			case apiextensionsv1beta1.NamesAccepted:
				if cond.Status == apiextensionsv1beta1.ConditionFalse {
					fmt.Printf("Name conflict: %v\n", cond.Reason)
				}
			}
		}
		return false, err
	})
	return err
}
