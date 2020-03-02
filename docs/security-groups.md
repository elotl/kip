## Security Groups

Cloud providers differ on the implementation details of their security groups. The following information applies to security groups in AWS.

Each Kip controller will create a cloud single security group in the VPC in order to allow cluster communication with cells.  The security group is named `kip-<UID>-CellSecurityGroup`. That security group opens all ports to the VPC.  Additional security groups can be attached to all nodes in the cluster by specifying security group IDs in the server.yml ConfigMap.  Security groups can be attached to individual cells by using the `pod.elotl.co/security-groups` annotation detailed in server.yml
