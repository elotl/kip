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
