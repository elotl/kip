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

func GetFakeService() *Service {
	s := NewService()
	s.Name = "test-service-" + uuid.NewV4().String()
	s.Name = s.Name[:24]
	s.Spec.Selector = LabelSelector{
		MatchLabels: map[string]string{"app": "myapp"},
	}
	s.Spec.Ports = []ServicePort{{
		Name:          "port",
		Protocol:      "TCP",
		Port:          5432,
		PortRangeSize: 1,
	}}
	s.Spec.SourceRanges = []string{"10.0.0.0/8"}
	return s
}

func GetFakeSecret(name string, keyvalues ...string) *Secret {
	s := NewSecret()
	s.Name = name
	for i := 0; i < len(keyvalues)-1; i += 2 {
		s.Data[keyvalues[i]] = []byte(keyvalues[i+1])
	}
	return s
}

func GetFakeLog() *LogFile {
	log := NewLogFile()
	log.Name = "test-log-file-" + uuid.NewV4().String()
	node := GetFakeNode()
	log.ParentObject = ToObjectReference(node)
	log.Content = "log data"
	return log
}

func GetFakeUsage() *Usage {
	usage := NewUsage()
	usage.Name = uuid.NewV4().String()
	usage.Provider = "AWS"
	usage.Instance = &InstanceUsage{
		Type: "t2.nano",
		Spot: false,
	}
	return usage
}
