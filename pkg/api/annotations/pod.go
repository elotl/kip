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

package annotations

// PodMilpactlVolumeName is used to specify the name of a volume that
// the controller will create containing milpactl and certs for
// connecting to the milpa controller from the pod.
const PodMilpactlVolumeName = "pod.elotl.co/milpactl-volume-name"

// PodTaskExecutionRole is the ARN of the task execution role that
// the fargate docker daemon can assume. This role is used for
// pulling images from ECR
const PodTaskExecutionRole = "pod.elotl.co/task-execution-role"

// The TaskRoleArn is the short name or full ARN of the IAM role
// that containers in a fargate task can assume. All containers
// in the task assume this role.
const PodTaskRole = "pod.elotl.co/task-role"
