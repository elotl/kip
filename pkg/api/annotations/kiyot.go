package annotations

// Depreciated
// PodKiyotSpotInstance is an annotation users can put on their kubernetes
// pods to tell kiyot to launch the pod on a node backed by a spot instance
const PodKiyotSpotInstance = "pod.elotl.co/spot-instance"

// PodKiyotLaunchType is an annotation users can put on their
// kubernetes pods to tell kiyot to launch the pod on a spot instance
// or container instance node
const PodKiyotLaunchType = "pod.elotl.co/launch-type"

// PodKiyotInstanceType is an annotation users can put on their
// kubernetes pods to tell kiyot to use a specific instance type for
// the node the pod will be launched onto.  This annotation will
// override specified resource requests and limits.
const PodKiyotInstanceType = "pod.elotl.co/instance-type"

// PodSecurityGroups is an annotation users can put on their
// kubernetes pods to tell kiyot to add additional security groups
// to the instance backing their pod.
const PodSecurityGroups = "pod.elotl.co/security-groups"

// PodSecurityGroups is an annotation users can put on their
// kubernetes pods to tell kiyot to attach an instance profile
// to the instance backing their pod.
const PodInstanceProfile = "pod.elotl.co/instance-profile"
