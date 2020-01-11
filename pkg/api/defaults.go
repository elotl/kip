package api

func NewNode() *Node {
	node := Node{
		TypeMeta: TypeMeta{Kind: "Node"},
		Status: NodeStatus{
			Phase: NodeCreating,
		},
	}
	node.ObjectMeta.Create()
	node.TypeMeta.Create()
	return &node
}

func NewNodeList() *NodeList {
	list := NodeList{
		TypeMeta: TypeMeta{Kind: "NodeList"},
		Items:    make([]*Node, 0),
	}
	list.TypeMeta.Create()
	return &list
}

func NewPod() *Pod {
	p := Pod{
		TypeMeta: TypeMeta{Kind: "Pod"},
		Spec: PodSpec{
			Phase:         PodRunning,
			RestartPolicy: RestartPolicyAlways,
			Spot: PodSpot{
				Policy: SpotNever,
			},
		},
		Status: PodStatus{
			Phase:           PodWaiting,
			LastPhaseChange: Now(),
		}}
	p.ObjectMeta.Create()
	p.TypeMeta.Create()
	return &p
}

func NewPodList() *PodList {
	list := PodList{
		TypeMeta: TypeMeta{Kind: "PodList"},
		Items:    make([]*Pod, 0),
	}
	list.TypeMeta.Create()
	return &list
}

func (p *ServicePort) UnmarshalJSON(data []byte) error {
	// prevent recursive calls
	type servicePortAlias ServicePort
	s := &servicePortAlias{
		Protocol:      ProtocolTCP,
		PortRangeSize: 1,
	}
	err := json.Unmarshal(data, s)
	if err != nil {
		return err
	}
	*p = ServicePort(*s)
	return nil
}

func NewService() *Service {
	svc := Service{
		TypeMeta: TypeMeta{Kind: "Service"},
	}
	svc.ObjectMeta.Create()
	svc.TypeMeta.Create()
	return &svc
}

func NewServiceList() *ServiceList {
	list := ServiceList{
		TypeMeta: TypeMeta{Kind: "ServiceList"},
		Items:    make([]*Service, 0),
	}
	list.TypeMeta.Create()
	return &list
}

func NewSecret() *Secret {
	s := Secret{
		TypeMeta: TypeMeta{Kind: "Secret"},
		Data:     map[string][]byte{},
	}
	s.ObjectMeta.Create()
	s.TypeMeta.Create()
	return &s
}

func NewSecretList() *SecretList {
	list := SecretList{
		TypeMeta: TypeMeta{Kind: "SecretList"},
		Items:    make([]*Secret, 0),
	}
	list.TypeMeta.Create()
	return &list
}

func NewEvent() *Event {
	e := Event{
		TypeMeta: TypeMeta{Kind: "Event"},
	}
	e.TypeMeta.Create()
	e.ObjectMeta.Create()
	return &e
}

func NewEventList() *EventList {
	list := EventList{
		TypeMeta: TypeMeta{Kind: "EventList"},
		Items:    make([]*Event, 0),
	}
	list.TypeMeta.Create()
	return &list
}

func NewLogFile() *LogFile {
	log := LogFile{
		TypeMeta: TypeMeta{Kind: "LogFile"},
	}
	log.TypeMeta.Create()
	log.ObjectMeta.Create()
	return &log
}

func NewLogFileList() *LogFileList {
	list := LogFileList{
		TypeMeta: TypeMeta{Kind: "LogFileList"},
		Items:    make([]*LogFile, 0),
	}
	list.TypeMeta.Create()
	return &list
}

func NewUsage() *Usage {
	u := Usage{
		TypeMeta: TypeMeta{Kind: "Usage"},
	}
	u.TypeMeta.Create()
	u.ObjectMeta.Create()
	return &u
}

func NewUsageList() *UsageList {
	list := UsageList{
		TypeMeta: TypeMeta{Kind: "UsageList"},
		Items:    make([]*Usage, 0),
	}
	list.TypeMeta.Create()
	return &list
}

func NewMetrics() *Metrics {
	m := Metrics{
		TypeMeta: TypeMeta{Kind: "Metrics"},
	}
	m.TypeMeta.Create()
	m.ObjectMeta.Create()
	return &m
}

func NewMetricsList() *MetricsList {
	list := MetricsList{
		TypeMeta: TypeMeta{Kind: "MetricsList"},
		Items:    make([]*Metrics, 0),
	}
	list.TypeMeta.Create()
	return &list
}
