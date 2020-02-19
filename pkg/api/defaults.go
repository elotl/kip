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
