package api

import (
	uuid "github.com/satori/go.uuid"
)

func GetFakeNode() *Node {
	n1 := NewNode()
	n1.Name = uuid.NewV4().String()
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
