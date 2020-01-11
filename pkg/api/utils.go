package api

func AllPodUnits(pod *Pod) []Unit {
	units := make([]Unit, 0, len(pod.Spec.Units)+len(pod.Spec.InitUnits))
	for i := 0; i < len(pod.Spec.InitUnits); i++ {
		units = append(units, pod.Spec.InitUnits[i])
	}
	for i := 0; i < len(pod.Spec.Units); i++ {
		units = append(units, pod.Spec.Units[i])
	}
	return units
}
