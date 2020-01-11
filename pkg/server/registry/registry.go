package registry

import (
	"errors"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
)

const MilpaRoot = "milpa/"

type Registryer interface {
	New() api.MilpaObject
	Create(obj api.MilpaObject) (api.MilpaObject, error)
	Update(obj api.MilpaObject) (api.MilpaObject, error)
	Get(id string) (api.MilpaObject, error)
	List() (api.MilpaObject, error)
	Delete(id string) (api.MilpaObject, error)
	//Validate(obj api.MilpaObject) field.ErrorList
}

// These Lister Interfaces were added to enforce read only registries
// that prevent callers from modifying core objects they have no
// business messing with.
type PodLister interface {
	GetPod(string) (*api.Pod, error)
	ListPods(func(*api.Pod) bool) (*api.PodList, error)
}

type ServiceLister interface {
	GetService(string) (*api.Service, error)
	ListServices(func(*api.Service) bool) (*api.ServiceList, error)
}

type SecretLister interface {
	GetSecret(string) (*api.Secret, error)
}

type NodeLister interface {
	GetNode(string) (*api.Node, error)
	ListNodes(func(*api.Node) bool) (*api.NodeList, error)
}

var (
	ErrAlreadyExists = errors.New("Object already exists")
	trashTTL         = 60 * time.Second
)
