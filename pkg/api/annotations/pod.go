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

// PodLaunchType is an annotation users can put on their
// kubernetes pods to tell kip to launch the pod on a spot instance
// or container instance node
const PodLaunchType = "pod.elotl.co/launch-type"

// PodInstanceType is an annotation users can put on their
// kubernetes pods to tell kip to use a specific instance type for
// the node the pod will be launched onto.  This annotation will
// override specified resource requests and limits.
const PodInstanceType = "pod.elotl.co/instance-type"

// PodResourcesPrivateIPOnly is an annotation users can put on their
// kubernetes pods to tell kip to run this pod on a node that only
// has a private IP address and no public IP address. Setting this
// value to false will not override the cloud subnet settings.
const PodResourcesPrivateIPOnly = "pod.elotl.co/private-ip-only"

// PodSecurityGroups is an annotation users can put on their
// kubernetes pods to tell kip to add additional security groups
// to the instance backing their pod.
const PodSecurityGroups = "pod.elotl.co/security-groups"

// PodSecurityGroups is an annotation users can put on their
// kubernetes pods to tell kip to attach an instance profile
// to the instance backing their pod.
const PodInstanceProfile = "pod.elotl.co/instance-profile"

// PodTaskExecutionRole is the ARN of the task execution role that
// the fargate docker daemon can assume. This role is used for
// pulling images from ECR
const PodTaskExecutionRole = "pod.elotl.co/task-execution-role"

// The TaskRoleArn is the short name or full ARN of the IAM role
// that containers in a fargate task can assume. All containers
// in the task assume this role.
const PodTaskRole = "pod.elotl.co/task-role"

// The PodHealthcheckHealthyTimeout annotation is used to customize the
// healthcheck timeout for pods. If a pod doesn't have a healthy
// response to healthcheck probes for greater than healthcheck-timeout
// the pod will be terminated and restarted according to the pod's
// restartPolicy.  A healthcheck-timeout equal to zero means the pod
// will not be terminated due to failing healthchecks.
const PodHealthcheckHealthyTimeout = "pod.elotl.co/healthcheck-healthy-timeout"
