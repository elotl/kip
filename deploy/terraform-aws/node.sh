#!/bin/bash

set -e

apt-get update
apt-get install -y apt-transport-https ca-certificates curl gnupg lsb-release
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF > /etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
apt-get update
apt-get remove -y docker docker-engine docker.io containerd runc
apt-get install -y kubelet="${k8s_version}*" kubeadm="${k8s_version}*" kubectl="${k8s_version}*" docker-ce docker-ce-cli containerd.io iproute2

# Ensure Docker does not block forwarded packets.
iptables -P FORWARD ACCEPT
mkdir -p /etc/docker
echo -e '{\n"iptables": false\n}' > /etc/docker/daemon.json
modprobe br_netfilter
systemctl enable docker.service || true
systemctl restart docker.service || true

name="$(hostname).ec2.internal"
while [[ -z "$name" ]]; do
    echo "waiting for IP address"
    sleep 1
done

ip=""
while [[ -z "$ip" ]]; do
    echo "waiting for IP address"
    sleep 1
    ip="$(ip route get 8.8.8.8 | grep '\<src\>' | head -1 | awk '{print $7}')"
done

echo "hostname: $name IP address: $ip"

if [ -z ${k8s_version} ]; then
    k8s_version=$(curl -fL https://storage.googleapis.com/kubernetes-release/release/stable.txt)
else
    k8s_version=v${k8s_version}
fi

cat <<EOF > /tmp/old-kubeadm-config.yaml
apiVersion: kubeadm.k8s.io/v1beta1
kind: InitConfiguration
bootstrapTokens:
- groups:
  - system:bootstrappers:kubeadm:default-node-token
  token: ${k8stoken}
nodeRegistration:
  name: $name
  kubeletExtraArgs:
    node-ip: $ip
    cloud-provider: aws
---
apiVersion: kubeadm.k8s.io/v1beta1
kind: ClusterConfiguration
networking:
  podSubnet: ${pod_cidr}
  serviceSubnet: ${service_cidr}
apiServer:
  certSANs:
  - 127.0.0.1
  - localhost
  extraArgs:
    enable-admission-plugins: DefaultStorageClass,NodeRestriction
    cloud-provider: aws
controllerManager:
  extraArgs:
    cloud-provider: aws
    configure-cloud-routes: "false"
    address: 0.0.0.0
kubernetesVersion: "$k8s_version"
EOF
kubeadm config migrate --old-config /tmp/old-kubeadm-config.yaml --new-config /tmp/kubeadm-config.yaml
kubeadm init --config=/tmp/kubeadm-config.yaml

export KUBECONFIG=/etc/kubernetes/admin.conf

# Configure kubectl.
mkdir -p /home/ubuntu/.kube
chown ubuntu: /home/ubuntu/.kube
cp -i $KUBECONFIG /home/ubuntu/.kube/config
chown ubuntu: /home/ubuntu/.kube/config

# Create a default storage class, backed by EBS.
curl -fL https://raw.githubusercontent.com/elotl/milpa-deploy/master/deploy/storageclass-ebs.yaml | kubectl apply -f -

# Deploy CNI plugin.
curl -fL https://raw.githubusercontent.com/elotl/milpa-deploy/master/deploy/cni/aws-k8s-cni.yaml | kubectl apply -f -

# Don't run kube-proxy on kip.
kubectl patch -p '{"spec":{"template":{"spec":{"affinity":{"nodeAffinity":{"requiredDuringSchedulingIgnoredDuringExecution":{"nodeSelectorTerms":[{"matchExpressions":[{"key":"type","operator":"NotIn","values":["virtual-kubelet"]}]}]}}}}}}}' -n kube-system ds kube-proxy

# Deploy main manifest.
echo "${kip_manifest}" | base64 --decode > /tmp/kip.yaml
kubectl apply -f /tmp/kip.yaml
