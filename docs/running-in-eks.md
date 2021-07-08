# Running KIP in EKS cluster

It is common use case to use Elastic Kubernetes Cluster by AWS and run KIP inside of it.
There are few prerequisites:
1. API server endpoint access has to be set as `Private` or `Public and Private` (you can configure this in AWS console under EKS clusters -> your cluster -> Networking -> Manage networking)
2. You have to specify `vpcID`, `subnetID` as well as `extraSecurityGroups` in [provider-config](../deploy/manifests/kip/overlays/aws-eks/provider.yaml). vpcID should match your cluster VPC, and cluster control-plane and worker nodes security groups has to be listed under `cells.extraSecurityGroups`.
3. Once you fill everything what's need in [aws-eks-kip-provider-config](../deploy/manifests/kip/overlays/aws-eks/provider.yaml) you can deploy KIP using kustomize: `kustomize build deploy/manifests/kip/overlays/aws-eks | kubectl apply -f  -`

Once KIP statefulset pods are running, you should see kip-provider virtual nodes listed as Ready `kubectl get nodes -o wide -l type=virtual-kubelet`.

