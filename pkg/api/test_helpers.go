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

package api

import (
	uuid "github.com/satori/go.uuid"
)

func GetFakeNode() *Node {
	n1 := NewNode()
	n1.Name = uuid.NewV4().String()
	n1.Spec.Architecture = ArchX8664
	n1.Spec.InstanceType = "t2.nano"
	n1.Spec.BootImage = "ami-12345"
	return n1
}

func GetFakePod() *Pod {
	lab := map[string]string{"Label1": "Value1"}
	units := []Unit{Unit{
		Name:    "unit-name",
		Image:   "UnitImage",
		Command: []string{"UnitCommand"},
	}}
	name := "deleteme-" + uuid.NewV4().String()
	p1 := NewPod()
	p1.Name = name
	p1.Labels = lab
	p1.Spec.Units = units
	p1.Spec.Architecture = ArchX8664
	p1.Spec.InstanceType = "t2.nano"
	return p1
}

func GetFakeEvent() *Event {
	e := NewEvent()
	e.InvolvedObject = ObjectReference{
		Kind: "object",
		Name: "myobject",
		UID:  "my-uid-12345",
	}
	e.Status = "object-created"
	e.Source = "event-tests"
	e.Message = "Testing events"
	return e
}

func GetFakeLog() *LogFile {
	log := NewLogFile()
	log.Name = "test-log-file-" + uuid.NewV4().String()
	node := GetFakeNode()
	log.ParentObject = ToObjectReference(node)
	log.Content = "log data"
	return log
}
