package cmd

var (
	createTypes []string
	getTypes    []string
	deleteTypes []string
)

func SetupKnownTypes() {
	createTypes = []string{
		"Pod",
		"Service",
		"Metric",
		"Secret",
	}
	getTypes = make([]string, len(createTypes))
	copy(getTypes, createTypes)
	getTypes = append(getTypes, "Node", "Event")

	deleteTypes = make([]string, len(createTypes))
	copy(deleteTypes, createTypes)
	deleteTypes = append(deleteTypes, "Node")
}
