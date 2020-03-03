package server

import (
	"fmt"
	"testing"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	v1beta1 "github.com/elotl/cloud-instance-provider/pkg/apis/kip/v1beta1"
	"github.com/elotl/cloud-instance-provider/pkg/server/events"
	"github.com/elotl/cloud-instance-provider/pkg/server/registry"
	"github.com/elotl/cloud-instance-provider/pkg/util/controllerqueue"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
)

type fakeKipClient struct {
	recs map[string]*v1beta1.Cell
}

func newFakeKipClient() *fakeKipClient {
	return &fakeKipClient{
		recs: make(map[string]*v1beta1.Cell),
	}
}

func (c *fakeKipClient) Create(n *v1beta1.Cell) (*v1beta1.Cell, error) {
	if _, exists := c.recs[n.Name]; exists {
		return nil, fmt.Errorf("Alredy exists")
	}
	c.recs[n.Name] = n
	return n, nil
}

func (c *fakeKipClient) Update(n *v1beta1.Cell) (*v1beta1.Cell, error) {
	if _, exists := c.recs[n.Name]; exists {
		c.recs[n.Name] = n
		return n, nil
	}
	return nil, fmt.Errorf("Can't update a record that doesn't exist")
}

func (c *fakeKipClient) Delete(name string, options *v1.DeleteOptions) error {
	delete(c.recs, name)
	return nil
}

func (c *fakeKipClient) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return nil
}

func (c *fakeKipClient) Get(name string, options v1.GetOptions) (*v1beta1.Cell, error) {
	if _, exists := c.recs[name]; exists {
		return c.recs[name], nil
	}
	return nil, fmt.Errorf("Item does not exist")
}

func (c *fakeKipClient) List(opts v1.ListOptions) (*v1beta1.CellList, error) {
	lst := &v1beta1.CellList{}
	lst.Items = make([]v1beta1.Cell, 0, len(c.recs))
	for _, v := range c.recs {
		lst.Items = append(lst.Items, *v)
	}
	return lst, nil
}

func (c *fakeKipClient) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return watch.NewEmptyWatch(), nil
}

func (c *fakeKipClient) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Cell, err error) {
	return nil, nil
}

func createTestCellController() (*CellController, func()) {
	podRegistry, closer1 := registry.SetupTestPodRegistry()
	nodeRegistry, closer2 := registry.SetupTestNodeRegistry()

	c := &CellController{
		controllerID:  "test1",
		k8sRestConfig: nil,
		k8sKipClient:  newFakeKipClient(),
		eventsSystem:  nil,
		podLister:     podRegistry,
		nodeLister:    nodeRegistry,
	}
	quit := make(chan struct{})
	queueCloser := func() { close(quit) }
	closer := func() { queueCloser(); closer1(); closer2() }
	c.queue = controllerqueue.New("cells", c.processOperation, controllerqueue.NumWorkers(1))
	c.queue.Start(quit)
	return c, closer
}

func waitForEmptyCellQueue(q *controllerqueue.Queue) error {
	start := time.Now()
	timeout := 3 * time.Second
	for time.Since(start) < timeout {
		if q.Len() == 0 {
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("Timed out waiting for an empty queue")
}

func TestCellCtlNodeCreated(t *testing.T) {
	c, closer := createTestCellController()
	defer closer()

	n := api.GetFakeNode()
	n.Status.Phase = api.NodeClaimed
	n, err := c.nodeLister.(*registry.NodeRegistry).CreateNode(n)
	assert.NoError(t, err)
	err = c.nodeCreated(events.Event{
		Object: n,
	})
	assert.NoError(t, err)
	err = waitForEmptyCellQueue(c.queue)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	kn, err := c.k8sKipClient.Get(n.Name, metav1.GetOptions{})
	assert.NoError(t, err)
	assert.Equal(t, n.Name, kn.Name)
	assert.Equal(t, c.controllerID, kn.Status.ControllerID)
	assert.Equal(t, n.Spec.InstanceType, kn.Status.InstanceType)
	assert.Equal(t, n.Status.InstanceID, kn.Status.InstanceID)
}

func TestCellCtlPodRunning(t *testing.T) {
	c, closer := createTestCellController()
	defer closer()

	p := api.GetFakePod()
	p.Labels = map[string]string{"app": "tester"}
	n := api.GetFakeNode()
	n.Status.InstanceID = "i-123"
	p.Status.BoundNodeName = n.Name
	p.Status.BoundInstanceID = n.Status.InstanceID
	n.Status.BoundPodName = p.Name
	n.Status.Phase = api.NodeClaimed
	p, err := c.podLister.(*registry.PodRegistry).CreatePod(p)
	assert.NoError(t, err)
	n, err = c.nodeLister.(*registry.NodeRegistry).CreateNode(n)
	assert.NoError(t, err)
	err = c.nodeCreated(events.Event{
		Object: n,
	})
	assert.NoError(t, err)

	err = c.podRunning(events.Event{
		Object: p,
	})
	assert.NoError(t, err)
	err = waitForEmptyCellQueue(c.queue)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	kn, err := c.k8sKipClient.Get(n.Name, metav1.GetOptions{})
	assert.NoError(t, err)
	assert.Equal(t, n.Name, kn.Name)
	assert.Equal(t, p.Labels, kn.Labels)
	assert.Equal(t, p.Name, kn.Status.PodName)
	assert.Equal(t, n.Spec.InstanceType, kn.Status.InstanceType)
	assert.Equal(t, n.Status.InstanceID, kn.Status.InstanceID)
}

func TestCellCtlNodePurged(t *testing.T) {
	c, closer := createTestCellController()
	defer closer()

	p := api.GetFakePod()
	p.Labels = map[string]string{"app": "tester"}
	n := api.GetFakeNode()
	n.Status.Phase = api.NodeClaimed
	n, err := c.nodeLister.(*registry.NodeRegistry).CreateNode(n)
	assert.NoError(t, err)
	err = c.nodePurged(events.Event{
		Object: n,
	})
	assert.NoError(t, err)
	err = waitForEmptyCellQueue(c.queue)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	_, err = c.k8sKipClient.Get(n.Name, metav1.GetOptions{})
	assert.Error(t, err)
}

func TestCellCtlSyncAll(t *testing.T) {
	c, closer := createTestCellController()
	defer closer()

	// Add: create a node that doesn't have a cell
	n := api.GetFakeNode()
	n.Status.Phase = api.NodeClaimed
	_, err := c.nodeLister.(*registry.NodeRegistry).CreateNode(n)
	assert.NoError(t, err)
	addNode := n

	// // Add: create a container instance that doesn't have a cell
	// p := api.GetFakePod()
	// p.Status.Phase = api.PodRunning
	// p.Status.BoundInstanceID = "milpa-1234-foo"
	// truth := true
	// p.Spec.Resources.ContainerInstance = &truth
	// p.Spec.InstanceType = ""
	// p, err = c.podLister.(*registry.PodRegistry).CreatePod(p)
	// assert.NoError(t, err)
	// addPod := p

	// Update: Create a node that'll need updating
	n = api.GetFakeNode()
	n.Status.Phase = api.NodeClaimed
	n, err = c.nodeLister.(*registry.NodeRegistry).CreateNode(n)
	assert.NoError(t, err)
	kn := c.makeCell(n, nil)
	_, err = c.k8sKipClient.Create(kn)
	assert.NoError(t, err)

	p := api.GetFakePod()
	p.Labels = map[string]string{"app": "tester"}
	p.Status.BoundNodeName = n.Name
	n.Status.BoundPodName = p.Name
	p, err = c.podLister.(*registry.PodRegistry).CreatePod(p)
	assert.NoError(t, err)
	n, err = c.nodeLister.(*registry.NodeRegistry).UpdateStatus(n)
	assert.NoError(t, err)
	updatedNodePod := p

	// // Update: create a container instance that needs updating
	// p = api.GetFakePod()
	// p.Status.Phase = api.PodRunning
	// p.Status.BoundInstanceID = "milpa-5678-bar"
	// p.Spec.Resources.ContainerInstance = &truth
	// p.Spec.InstanceType = ""
	// p, err = c.podLister.(*registry.PodRegistry).CreatePod(p)
	// assert.NoError(t, err)
	// kn = c.makeCell(nil, p)
	// _, err = c.k8sKipClient.Create(kn)
	// assert.NoError(t, err)
	// p.Status.Addresses = api.NewNetworkAddresses("1.2.3.4", "ec2-1-2-3-4.aws.co")
	// _, err = c.podLister.(*registry.PodRegistry).UpdatePodStatus(p, "")
	// assert.NoError(t, err)
	// updatedContainerInstancePod := p

	// Delete: have a node cell that shouldn't be there
	kn = &v1beta1.Cell{
		Status: v1beta1.CellStatus{
			ControllerID: c.controllerID,
		},
	}
	doomedName := "deleteme"
	kn.Name = doomedName
	_, err = c.k8sKipClient.Create(kn)
	assert.NoError(t, err)

	// Other Controller: preserve cells that that we didn't create
	kn = &v1beta1.Cell{
		Status: v1beta1.CellStatus{
			ControllerID: "the-other-controller",
		},
	}
	otherName := "other-controller-cell"
	kn.Name = otherName
	_, err = c.k8sKipClient.Create(kn)
	assert.NoError(t, err)

	err = c.syncAllCells()
	assert.NoError(t, err)

	err = waitForEmptyCellQueue(c.queue)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	// Ensure node add worked
	_, err = c.k8sKipClient.Get(addNode.Name, metav1.GetOptions{})
	assert.NoError(t, err)

	// // Ensure container instance add worked
	// _, err = c.k8sKipClient.Get(addPod.Status.BoundInstanceID, metav1.GetOptions{})
	// assert.NoError(t, err)

	// ensure node update worked
	kn, err = c.k8sKipClient.Get(updatedNodePod.Status.BoundNodeName, metav1.GetOptions{})
	assert.NoError(t, err)
	assert.Equal(t, updatedNodePod.Labels, kn.Labels)

	// // ensure container instance update worked
	// kn, err = c.k8sKipClient.Get(updatedContainerInstancePod.Status.BoundInstanceID, metav1.GetOptions{})
	// assert.NoError(t, err)
	// assert.NotEmpty(t, kn.Status.IP)

	// Ensure node delete worked
	_, err = c.k8sKipClient.Get(doomedName, metav1.GetOptions{})
	assert.Error(t, err)

	// Ensure we don't touch other cells of other controllers
	_, err = c.k8sKipClient.Get(otherName, metav1.GetOptions{})
	assert.NoError(t, err)

}
