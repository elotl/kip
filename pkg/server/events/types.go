package events

const (
	AllEvents       = "all-events"
	NodeCleaning    = "node-cleaning"
	NodeCreated     = "node-created"
	NodePurged      = "node-purged"
	NodeRunning     = "node-running"
	PodCreated      = "pod-created"
	PodEjected      = "pod-ejected" // We found a lost node with a bound pod
	PodRunning      = "pod-running"
	PodShouldDelete = "pod-should-delete"
	PodTerminated   = "pod-terminated"
	PodUpdated      = "pod-updated"
	SecretCreated   = "secret-created"
	SecretDeleted   = "secret-deleted"
	SecretUpdated   = "secret-updated"
	ServiceCreated  = "service-created"
	ServiceDeleted  = "service-deleted"
	ServiceUpdated  = "service-updated"
	StartSpotFailed = "start-spot-failed"
	UsageCreated    = "usage-created"
)
