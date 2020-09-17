package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/elotl/kip/pkg/api"
	v1beta1 "github.com/elotl/kip/pkg/apis/kip/v1beta1"
	"github.com/elotl/kip/pkg/server/events"
	"github.com/elotl/kip/pkg/server/registry"
	"github.com/elotl/kip/pkg/util/controllerqueue"
	"github.com/stretchr/testify/assert"
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

func (c *fakeKipClient) Create(ctx context.Context, n *v1beta1.Cell, opts metav1.CreateOptions) (*v1beta1.Cell, error) {
	if _, exists := c.recs[n.Name]; exists {
		return nil, fmt.Errorf("Alredy exists")
	}
	c.recs[n.Name] = n
	return n, nil
}

func (c *fakeKipClient) Update(ctx context.Context, n *v1beta1.Cell, opts metav1.UpdateOptions) (*v1beta1.Cell, error) {
	if _, exists := c.recs[n.Name]; exists {
		c.recs[n.Name] = n
		return n, nil
	}
	return nil, fmt.Errorf("Can't update a record that doesn't exist")
}

func (c *fakeKipClient) Delete(ctx context.Context, name string, options metav1.DeleteOptions) error {
	delete(c.recs, name)
	return nil
}

func (c *fakeKipClient) DeleteCollection(ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return nil
}

func (c *fakeKipClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.Cell, error) {
	if _, exists := c.recs[name]; exists {
		return c.recs[name], nil
	}
	return nil, fmt.Errorf("Item does not exist")
}

func (c *fakeKipClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.CellList, error) {
	lst := &v1beta1.CellList{}
	lst.Items = make([]v1beta1.Cell, 0, len(c.recs))
	for _, v := range c.recs {
		lst.Items = append(lst.Items, *v)
	}
	return lst, nil
}

func (c *fakeKipClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return watch.NewEmptyWatch(), nil
}

func (c *fakeKipClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1beta1.Cell, err error) {
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
	ctx := context.Background()
	kn, err := c.k8sKipClient.Get(ctx, n.Name, metav1.GetOptions{})
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
	ctx := context.Background()
	kn, err := c.k8sKipClient.Get(ctx, n.Name, metav1.GetOptions{})
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
	ctx := context.Background()
	_, err = c.k8sKipClient.Get(ctx, n.Name, metav1.GetOptions{})
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

	// Update: Create a node that'll need updating
	n = api.GetFakeNode()
	n.Status.Phase = api.NodeClaimed
	n, err = c.nodeLister.(*registry.NodeRegistry).CreateNode(n)
	assert.NoError(t, err)
	kn := c.makeCell(n, nil)
	ctx := context.Background()
	_, err = c.k8sKipClient.Create(ctx, kn, metav1.CreateOptions{})
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

	// Delete: have a node cell that shouldn't be there
	kn = &v1beta1.Cell{
		Status: v1beta1.CellStatus{
			ControllerID: c.controllerID,
		},
	}
	doomedName := "deleteme"
	kn.Name = doomedName
	_, err = c.k8sKipClient.Create(ctx, kn, metav1.CreateOptions{})
	assert.NoError(t, err)

	// Other Controller: preserve cells that that we didn't create
	kn = &v1beta1.Cell{
		Status: v1beta1.CellStatus{
			ControllerID: "the-other-controller",
		},
	}
	otherName := "other-controller-cell"
	kn.Name = otherName
	_, err = c.k8sKipClient.Create(ctx, kn, metav1.CreateOptions{})
	assert.NoError(t, err)

	err = c.syncAllCells()
	assert.NoError(t, err)

	err = waitForEmptyCellQueue(c.queue)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	// Ensure node add worked
	_, err = c.k8sKipClient.Get(ctx, addNode.Name, metav1.GetOptions{})
	assert.NoError(t, err)

	// ensure node update worked
	kn, err = c.k8sKipClient.Get(ctx, updatedNodePod.Status.BoundNodeName, metav1.GetOptions{})
	assert.NoError(t, err)
	assert.Equal(t, updatedNodePod.Labels, kn.Labels)

	// Ensure node delete worked
	_, err = c.k8sKipClient.Get(ctx, doomedName, metav1.GetOptions{})
	assert.Error(t, err)

	// Ensure we don't touch other cells of other controllers
	_, err = c.k8sKipClient.Get(ctx, otherName, metav1.GetOptions{})
	assert.NoError(t, err)

}
