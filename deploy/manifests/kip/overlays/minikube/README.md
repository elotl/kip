## Deploy to an existing Minikube

1. Set `region`, `accessKeyID`, `secretAccessKey`, `vpcID`, `subnetID` and `extraSecurityGroups`:
     $ vi provider.yaml

2.  Apply via: 
     $ kustomize build . | kubectl apply -f -`
