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
