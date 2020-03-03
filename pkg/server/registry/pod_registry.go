package registry

import (
	"fmt"
	"time"

	"github.com/docker/libkv/store"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/api/validation"
	"github.com/elotl/cloud-instance-provider/pkg/etcd"
	"github.com/elotl/cloud-instance-provider/pkg/server/events"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/instanceselector"
	"github.com/elotl/cloud-instance-provider/pkg/util/validation/field"
	"k8s.io/klog"
)

const (
	PodPath                      string = "milpa/pods"
	PodTrashPath                 string = "milpa/trash/pods"
	PodDirectoryPlaceholder      string = "milpa/pods/."
	PodTrashDirectoryPlaceholder string = "milpa/trash/pods/."
)

var (
	terminatedPodTTL = 3 * time.Minute
)

type PodRegistry struct {
	etcd.Storer
	Codec             api.MilpaCodec
	eventSystem       *events.EventSystem
	statefulValidator *validation.StatefulValidator
}

func makePodKey(id string) string {
	return PodPath + "/" + id
}

func makeDeletedPodKey(id string) string {
	return PodTrashPath + "/" + id
}

func NewPodRegistry(kvstore etcd.Storer, codec api.MilpaCodec, es *events.EventSystem, sv *validation.StatefulValidator) *PodRegistry {
	// empty directories create problems and pain the butt errors
	// lets avoid them
	reg := &PodRegistry{kvstore, codec, es, sv}
	reg.Put(PodDirectoryPlaceholder, []byte("."), &store.WriteOptions{IsDir: true})
	reg.Put(PodTrashDirectoryPlaceholder, []byte("."), &store.WriteOptions{IsDir: true})
	return reg
}

func (reg *PodRegistry) New() api.MilpaObject {
	return api.NewPod()
}

func (reg *PodRegistry) Validate(obj api.MilpaObject) error {
	pod := obj.(*api.Pod)
	errs := validation.ValidatePod(pod)
	if len(errs) > 0 {
		return validation.NewError("pod", pod.Name, errs)
	}
	errs = reg.statefulValidator.ValidatePodSpec(&pod.Spec, field.NewPath("spec"))
	if len(errs) > 0 {
		return validation.NewError("Pod", pod.Name, errs)
	}
	return nil
}

func (reg *PodRegistry) Create(obj api.MilpaObject) (api.MilpaObject, error) {
	pod := obj.(*api.Pod)
	return reg.CreatePod(pod)
}

func (reg *PodRegistry) Update(obj api.MilpaObject) (api.MilpaObject, error) {
	pod := obj.(*api.Pod)
	return reg.UpdatePodSpecAndLabels(pod)
}

func (reg *PodRegistry) Get(name string) (api.MilpaObject, error) {
	return reg.GetPod(name)
}

func MatchAllPods(p *api.Pod) bool {
	return true
}

func MatchAllLivePods(p *api.Pod) bool {
	if api.IsTerminalPodPhase(p.Spec.Phase) &&
		api.IsTerminalPodPhase(p.Status.Phase) {
		return false
	}
	return true
}

func (reg *PodRegistry) List() (api.MilpaObject, error) {
	return reg.ListPods(MatchAllPods)
}

// While this is called delete, that's simply because it's the
// handler for the milpactl call to "Delete".  Really, it
// specifies that we should terminate this pod and then delete it
func (reg *PodRegistry) Delete(name string) (api.MilpaObject, error) {
	pod, err := reg.GetPod(name)
	if err != nil {
		msg := fmt.Sprintf("Could not delete pod %s", name)
		return nil, util.WrapError(err, msg)
	}
	// allow the user to totally delete a terminated pod
	if api.IsTerminalPodPhase(pod.Spec.Phase) &&
		api.IsTerminalPodPhase(pod.Status.Phase) {
		err := reg.Storer.Delete(makePodKey(pod.Name))
		if err != nil {
			return nil, util.WrapError(err,
				"Error deleting %s from pod registry", pod.Name)
		}
		return pod, nil
	} else {
		return reg.MarkPodForTermination(pod)
	}
}

// I see this kinda like validation.  I'm not sure of a better
// place to put it.  Basically, it makes sure that the pod is
// in proper shape before it's created
func (reg *PodRegistry) preCreatePod(p *api.Pod) (*api.Pod, error) {
	instanceType, sustainedCPU, err := instanceselector.ResourcesToInstanceType(&p.Spec)
	if err != nil {
		return nil, util.WrapError(err, "Could not create pod %s, failure to convert resources to instance type", p.Name)
	}
	p.Spec.InstanceType = instanceType
	p.Spec.Resources.SustainedCPU = sustainedCPU
	return p, nil
}

func (reg *PodRegistry) isLivePod(name string) bool {
	pod, err := reg.GetPod(name)
	if err == store.ErrKeyNotFound {
		return false
	} else if err != nil {
		klog.Errorf("Error getting pod: %s, assuming pod is alive", err)
		return true
	}

	if api.IsTerminalPodPhase(pod.Spec.Phase) &&
		api.IsTerminalPodPhase(pod.Status.Phase) {
		return false
	}
	return true
}

func (reg *PodRegistry) CreatePod(p *api.Pod) (*api.Pod, error) {
	if err := reg.Validate(p); err != nil {
		return nil, err
	}
	p, err := reg.preCreatePod(p)
	if err != nil {
		return nil, err
	}
	key := makePodKey(p.Name)
	exists, err := reg.Storer.Exists(key)
	if err != nil {
		return nil, err
	} else if exists && reg.isLivePod(p.Name) {
		// if the pod has been terminated, then we overwrite it
		return nil, fmt.Errorf("Could not create pod %s. %s",
			p.Name, ErrAlreadyExists.Error())
	}

	data, err := reg.Codec.Marshal(p)
	if err != nil {
		return nil, err
	}
	// Don't atomic put here, just overwrite (yeah, there's a bit of a race)
	err = reg.Storer.Put(key, data, nil)
	if err != nil {
		return nil, util.WrapError(err, "Could not create pod in registry")
	}
	newPod, err := reg.GetPod(p.Name)
	if err != nil {
		return nil, util.WrapError(err, "Could not get pod after creation")
	}
	reg.eventSystem.Emit(events.PodCreated, "pod-registry", newPod)
	return newPod, nil
}

func (reg *PodRegistry) UpdatePodSpecAndLabels(p *api.Pod) (*api.Pod, error) {
	if err := reg.Validate(p); err != nil {
		return nil, err
	}
	p, err := reg.preCreatePod(p)
	if err != nil {
		return nil, err
	}
	p, err = reg.AtomicUpdate(p.Name, func(in *api.Pod) error {
		copyObjectMetaForUpdate(&in.ObjectMeta, &p.ObjectMeta)
		in.Spec = p.Spec
		return nil
	})
	return p, err
}

func (reg *PodRegistry) GetPod(k string) (*api.Pod, error) {
	key := makePodKey(k)
	pair, err := reg.Storer.Get(key)
	if err == store.ErrKeyNotFound {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("Error retrieving pod from storage: %v", err)
	}
	pod := api.NewPod()
	err = reg.Codec.Unmarshal(pair.Value, pod)
	if err != nil {
		return nil, util.WrapError(err, "Error unmarshaling pod from storage")
	}
	return pod, nil
}

func (reg *PodRegistry) MarkPodForTermination(pod *api.Pod) (*api.Pod, error) {
	p, err := reg.AtomicUpdate(pod.Name, func(in *api.Pod) error {
		in.Spec.Phase = api.PodTerminated
		return nil
	})
	if err == nil {
		eventMsg := fmt.Sprintf("Pod marked for deletion")
		reg.eventSystem.Emit(events.PodShouldDelete, "pod-registry", p, eventMsg)
	}
	return p, err
}

func (reg *PodRegistry) ListPods(filter func(*api.Pod) bool) (*api.PodList, error) {
	pairs, err := reg.Storer.List(PodPath)
	podlist := api.NewPodList()
	if err != nil {
		klog.Errorf("Error listing pods in storage: %v", err)
		return podlist, err
	}
	podlist.Items = make([]*api.Pod, 0, len(pairs))
	for _, pair := range pairs {
		// we create a blank key because dealing with "key does not
		// exist across different DBs is a road we dont want to go
		// down yet
		if pair.Key == PodDirectoryPlaceholder {
			continue
		}
		pod := api.NewPod()
		err = reg.Codec.Unmarshal(pair.Value, pod)
		if err != nil {
			klog.Errorf("Error unmarshalling single pod in list operation: %v", err)
			continue
		}
		if filter(pod) {
			podlist.Items = append(podlist.Items, pod)
		}
	}
	return podlist, nil
}

func validStatusPhaseChange(old, new api.PodPhase) bool {
	switch old {
	case api.PodWaiting:
		switch new {
		case api.PodDispatching, api.PodTerminated:
			return true
		default:
			return false
		}
	case api.PodDispatching:
		switch new {
		case api.PodRunning, api.PodFailed, api.PodTerminated:
			return true
		default:
			return false
		}
	case api.PodRunning:
		switch new {
		case api.PodFailed, api.PodTerminated, api.PodSucceeded:
			return true
		default:
			return false
		}
	case api.PodFailed:
		switch new {
		case api.PodWaiting, api.PodTerminated:
			return true
		default:
			return false
		}
	case api.PodSucceeded:
		switch new {
		case api.PodTerminated:
			return true
		default:
			return false
		}
	case api.PodTerminated:
		return false
	}
	klog.Fatalf("Programming error: Reached end of state transition table")
	return false
}

func (reg *PodRegistry) UpdatePodStatus(p *api.Pod, reason string) (*api.Pod, error) {
	status := p.Status
	p1, err := reg.AtomicUpdate(p.Name, func(in *api.Pod) error {
		if in.Status.Phase != p.Status.Phase &&
			!validStatusPhaseChange(in.Status.Phase, p.Status.Phase) {
			return fmt.Errorf(
				"Invalid State Change: %s -> %s", in.Status.Phase, p.Status.Phase)
		}

		if in.Status.Phase != p.Status.Phase {
			p.Status.LastPhaseChange = api.Now()
		}
		in.Status = status
		return nil
	})
	if err == nil {
		eventMsg := fmt.Sprintf("Pod status phase: '%s'", p.Status.Phase)
		if len(reason) > 0 {
			eventMsg += ": " + reason
		}
		reg.eventSystem.Emit(events.PodUpdated, "pod-registry", p1, eventMsg)
	}
	return p, err
}

func (reg *PodRegistry) TerminatePod(pod *api.Pod, phase api.PodPhase, msg string) error {
	reg.eventSystem.Emit(events.PodTerminated, "pod-registry", pod, msg)
	return reg.Storer.Delete(makePodKey(pod.Name))
}

type modifyPodFunc func(*api.Pod) error

func (reg *PodRegistry) AtomicUpdate(name string, modifier modifyPodFunc) (*api.Pod, error) {
	updatedPod := api.NewPod()
	key := makePodKey(name)
	for {
		pair, err := reg.Storer.Get(key)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving pod from storage: %v", err)
		}

		err = reg.Codec.Unmarshal(pair.Value, updatedPod)
		if err != nil {
			return nil, util.WrapError(err, "Error unmarshaling pod from storage")
		}
		err = modifier(updatedPod)
		if err != nil {
			return nil, util.WrapError(err, "Error modifying pod for update")
		}
		updatedValue, err := reg.Codec.Marshal(updatedPod)
		if err != nil {
			return nil, err
		}
		_, _, err = reg.Storer.AtomicPut(key, updatedValue, pair, nil)
		if err == store.ErrKeyModified {
			continue
		} else if err != nil {
			msg := fmt.Sprintf("Atomic Update of pod %s failed", key)
			return nil, util.WrapError(err, msg)
		}
		break
	}
	return updatedPod, nil
}
