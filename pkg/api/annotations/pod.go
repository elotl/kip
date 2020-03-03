package annotations

// PodKiyotLaunchType is an annotation users can put on their
// kubernetes pods to tell kiyot to launch the pod on a spot instance
// or container instance node
const PodLaunchType = "pod.elotl.co/launch-type"

// PodKiyotInstanceType is an annotation users can put on their
// kubernetes pods to tell kiyot to use a specific instance type for
// the node the pod will be launched onto.  This annotation will
// override specified resource requests and limits.
const PodInstanceType = "pod.elotl.co/instance-type"

// PodSecurityGroups is an annotation users can put on their
// kubernetes pods to tell kiyot to add additional security groups
// to the instance backing their pod.
const PodSecurityGroups = "pod.elotl.co/security-groups"

// PodSecurityGroups is an annotation users can put on their
// kubernetes pods to tell kiyot to attach an instance profile
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
