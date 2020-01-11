// Need to start DRYing up the registry code, this is terrible
package registry

import (
	"fmt"

	"github.com/docker/libkv/store"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/api/validation"
	"github.com/elotl/cloud-instance-provider/pkg/etcd"
	"github.com/elotl/cloud-instance-provider/pkg/server/events"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/golang/glog"
)

const (
	NodePath                      string = "milpa/nodes"
	NodeTrashPath                 string = "milpa/trash/nodes"
	NodeDirectoryPlaceholder      string = "milpa/nodes/."
	NodeTrashDirectoryPlaceholder string = "milpa/trash/nodes/."
)

type NodeRegistry struct {
	etcd.Storer
	Codec       api.MilpaCodec
	eventSystem *events.EventSystem
}

func makeNodeKey(id string) string {
	return NodePath + "/" + id
}

func makeDeletedNodeKey(id string) string {
	return NodeTrashPath + "/" + id
}

func NewNodeRegistry(kvstore etcd.Storer, codec api.MilpaCodec, es *events.EventSystem) *NodeRegistry {
	// empty directories create problems and pain the butt errors
	// lets avoid them
	reg := &NodeRegistry{kvstore, codec, es}
	_ = reg.Put(NodeDirectoryPlaceholder, []byte("."), &store.WriteOptions{IsDir: true})
	_ = reg.Put(NodeTrashDirectoryPlaceholder, []byte("."), &store.WriteOptions{IsDir: true})
	return reg
}

func (reg *NodeRegistry) New() api.MilpaObject {
	return api.NewNode()
}

func (reg *NodeRegistry) Validate(obj api.MilpaObject) error {
	node := obj.(*api.Node)
	errs := validation.ValidateNode(node)
	if len(errs) > 0 {
		return validation.NewError("node", node.Name, errs)
	}
	return nil
}

func (reg *NodeRegistry) Create(obj api.MilpaObject) (api.MilpaObject, error) {
	node := obj.(*api.Node)
	return reg.CreateNode(node)
}

func (reg *NodeRegistry) Update(obj api.MilpaObject) (api.MilpaObject, error) {
	node := obj.(*api.Node)
	return reg.UpdateNode(node)
}

func (reg *NodeRegistry) Get(id string) (api.MilpaObject, error) {
	return reg.GetNode(id)
}

func MatchAllNodes(p *api.Node) bool {
	return true
}

func (reg *NodeRegistry) List() (api.MilpaObject, error) {
	return reg.ListNodes(MatchAllNodes)
}

func (reg *NodeRegistry) Delete(name string) (api.MilpaObject, error) {
	// p, err := reg.AtomicUpdate(id, func(in *api.Node) error {
	// 	// todo, how do we handle this one?  Should
	// 	// we allow users to delete nodes?
	// 	in.Status.Phase = api.NodeTerminating
	// 	now := time.Now().UTC()
	// 	glog.Infof("Setting deletion time")
	// 	in.DeletionTimestamp = &now
	// 	return nil
	// })
	// return p, err
	node, err := reg.GetNode(name)
	if err == store.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, util.WrapError(err, "Could not get node from registry")
	}
	node, err = reg.MarkForTermination(node)
	if err != nil {
		return nil, util.WrapError(err, "Could not mark node for termination and deletion")
	}
	return node, err
}

func (reg *NodeRegistry) CreateNode(n *api.Node) (*api.Node, error) {
	if err := reg.Validate(n); err != nil {
		return nil, err
	}

	key := makeNodeKey(n.Name)
	exists, err := reg.Storer.Exists(key)
	if err != nil {
		return nil, err
	} else if exists {
		return nil, fmt.Errorf("Could not create node. A node named %s already exists", n.Name)
	}

	data, err := reg.Codec.Marshal(n)
	if err != nil {
		return nil, err
	}
	_, _, err = reg.Storer.AtomicPut(key, data, nil, nil)
	if err != nil {
		return nil, util.WrapError(err, "Could not create node in registry")
	}
	newNode, err := reg.GetNode(n.Name)
	if err != nil {
		return nil, util.WrapError(err, "Could not get node after creation")
	}
	return newNode, err
}

func (reg *NodeRegistry) UpdateNode(n *api.Node) (*api.Node, error) {
	return nil, fmt.Errorf("Updating a node is not implemented at this time")
}

func (reg *NodeRegistry) GetNode(name string) (*api.Node, error) {
	key := makeNodeKey(name)
	pair, err := reg.Storer.Get(key)
	if err == store.ErrKeyNotFound {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("Error retrieving node from storage: %v", err)
	}
	node := api.NewNode()
	err = reg.Codec.Unmarshal(pair.Value, node)
	if err != nil {
		return nil, util.WrapError(err, "Error unmarshaling node from storage")
	}
	return node, nil
}

func (reg *NodeRegistry) ListNodes(filter func(*api.Node) bool) (*api.NodeList, error) {
	return reg.listNodes(NodePath, filter)
}

// This is a terrible function but is helpful for tests in the node controller
// it's possible listnodes should just do this functionality...
func (reg *NodeRegistry) ListAllNodes(filter func(*api.Node) bool) (*api.NodeList, error) {
	nodes, err := reg.listNodes(NodePath, filter)
	if err != nil {
		return nil, err
	}
	deletedNodes, err := reg.listNodes(NodeTrashPath, filter)
	if err != nil {
		return nil, err
	}
	nodes.Items = append(nodes.Items, deletedNodes.Items...)
	return nodes, nil
}

func (reg *NodeRegistry) listNodes(nodePath string, filter func(*api.Node) bool) (*api.NodeList, error) {
	pairs, err := reg.Storer.List(nodePath)
	nodelist := api.NewNodeList()
	if err != nil {
		glog.Errorf("Error listing nodes in storage: %v", err)
		return nodelist, err
	}
	nodelist.Items = make([]*api.Node, 0, len(pairs))
	for _, pair := range pairs {
		// we create a blank key because dealing with "key does not
		// exist across different DBs is a road we dont want to go
		// down yet
		if pair.Key == NodeDirectoryPlaceholder ||
			pair.Key == NodeTrashDirectoryPlaceholder {
			continue
		}
		node := api.NewNode()
		err = reg.Codec.Unmarshal(pair.Value, node)
		if err != nil {
			glog.Errorf("Error unmarshalling single node in list operation: %v", err)
			continue
		}
		if filter(node) {
			nodelist.Items = append(nodelist.Items, node)
		}
	}
	return nodelist, nil
}

func (reg *NodeRegistry) PurgeNode(node *api.Node) (*api.Node, error) {
	glog.Infof("Purging node %v", node)
	reg.eventSystem.Emit(events.NodePurged, "node-registry", node)

	node.Status.Phase = api.NodeTerminated
	if node.DeletionTimestamp == nil {
		now := api.Now()
		node.DeletionTimestamp = &now
	}
	err := reg.Storer.Delete(makeNodeKey(node.Name))
	if err != nil {
		msg := fmt.Sprintf("Error deleting %s from node registry", node.Name)
		return nil, util.WrapError(err, msg)
	}

	// write to trash.  We don't care if it already exists, overwrite if it does
	key := makeDeletedNodeKey(node.Name)
	data, err := reg.Codec.Marshal(node)
	if err != nil {
		return nil, err
	}
	err = reg.Storer.Put(
		key,
		data,
		&store.WriteOptions{
			IsDir: false,
			TTL:   trashTTL,
		})
	if err != nil {
		glog.Warningf("Could not create deleted node %s in registry: %s",
			node.Name, err.Error())
	}
	return node, nil
}

func validStateChange(old, new api.NodePhase) bool {
	switch old {
	case api.NodeCreating:
		switch new {
		case api.NodeCreating, api.NodeCreated, api.NodeAvailable, api.NodeTerminating, api.NodeTerminated:
			if new == api.NodeTerminating {
				// These are scary things to see, we allow them to happen
				// because they're likely caused by the best of intentions:
				// a node failed to come up, somehow our initial update failed
				// so we are shutting them down.  If the cloud is OK with it,
				// we can let it pass.  Log it and carry on
				glog.Warningf("Racy termination: attempting to terminate Creating node")
			}
			return true
		default:
			return false
		}
	case api.NodeCreated:
		switch new {
		case api.NodeCreated, api.NodeAvailable, api.NodeCleaning, api.NodeTerminating, api.NodeTerminated:
			return true
		default:
			return false
		}
	case api.NodeAvailable:
		switch new {
		case api.NodeAvailable, api.NodeClaimed, api.NodeTerminating, api.NodeTerminated:
			return true
		default:
			return false
		}
	case api.NodeClaimed:
		switch new {
		case api.NodeClaimed, api.NodeCleaning, api.NodeTerminating, api.NodeTerminated:
			// I will assume that a garbage collector will someday
			// need to terminate claimed nodes
			return true
		default:
			return false
		}
	case api.NodeCleaning:
		switch new {
		case api.NodeAvailable, api.NodeCleaning, api.NodeTerminating, api.NodeTerminated:
			return true
		default:
			return false
		}
	case api.NodeTerminating:
		switch new {
		case api.NodeTerminating, api.NodeTerminated:
			return true
		default:
			return false
		}
	case api.NodeTerminated:
		switch new {
		case api.NodeTerminated:
			return true
		default:
			return false
		}
	}
	glog.Fatalf("Programming error: Reached end of state transition table")
	return false
}

func (reg *NodeRegistry) UpdateStatus(node *api.Node) (*api.Node, error) {
	n, err := reg.AtomicUpdate(node.Name, func(in *api.Node) error {
		if !validStateChange(in.Status.Phase, node.Status.Phase) {
			return fmt.Errorf(
				"Invalid State Change: %s -> %s", in.Status.Phase, node.Status.Phase)
		}
		in.Status = node.Status
		if node.Status.Phase == api.NodeTerminating {
			// Setting this allows the node reaper to do its job
			node.Spec.Terminate = true
		}
		if node.Status.Phase == api.NodeTerminated &&
			node.DeletionTimestamp == nil {
			now := api.Now()
			in.DeletionTimestamp = &now
		}
		return nil
	})
	return n, err
}

func (reg *NodeRegistry) SetNodeDeletionTimestamp(node *api.Node) (*api.Node, error) {
	n, err := reg.AtomicUpdate(node.Name, func(in *api.Node) error {
		now := api.Now()
		in.DeletionTimestamp = &now
		return nil
	})
	return n, err
}

func (reg *NodeRegistry) MarkForTermination(node *api.Node) (*api.Node, error) {
	n, err := reg.AtomicUpdate(node.Name, func(in *api.Node) error {
		in.Spec.Terminate = true
		return nil
	})
	return n, err
}

type modifyNodeFunc func(*api.Node) error

func (reg *NodeRegistry) AtomicUpdate(id string, modifier modifyNodeFunc) (*api.Node, error) {
	updatedNode := api.NewNode()
	key := makeNodeKey(id)
	for {
		pair, err := reg.Storer.Get(key)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving node from storage: %v", err)
		}

		err = reg.Codec.Unmarshal(pair.Value, updatedNode)
		if err != nil {
			return nil, util.WrapError(err, "Error unmarshaling node from storage")
		}
		err = modifier(updatedNode)
		if err != nil {
			return nil, util.WrapError(err, "Error modifying node for update")
		}
		updatedValue, err := reg.Codec.Marshal(updatedNode)
		if err != nil {
			return nil, err
		}
		_, _, err = reg.Storer.AtomicPut(key, updatedValue, pair, nil)
		if err == store.ErrKeyModified {
			continue
		} else if err != nil {
			msg := fmt.Sprintf("Atomic Update of node %s failed", key)
			return nil, util.WrapError(err, msg)
		}
		break
	}
	return updatedNode, nil
}
