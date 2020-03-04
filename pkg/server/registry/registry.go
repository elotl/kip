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

type NodeLister interface {
	GetNode(string) (*api.Node, error)
	ListNodes(func(*api.Node) bool) (*api.NodeList, error)
}

var (
	ErrAlreadyExists = errors.New("Object already exists")
	trashTTL         = 60 * time.Second
)
