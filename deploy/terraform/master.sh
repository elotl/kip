#!/bin/bash -v

curl -fL https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF > /etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet="${k8s_version}*" kubeadm="${k8s_version}*" kubectl="${k8s_version}*" kubernetes-cni docker.io

# Ensure Docker does not block forwarded packets.
iptables -P FORWARD ACCEPT
mkdir -p /etc/docker
echo -e '{\n"iptables": false\n}' > /etc/docker/daemon.json
systemctl restart docker.service || true

name=""
while [[ -z "$name" ]]; do
    sleep 1
    name="$(hostname -f)"
done

ip=""
while [[ -z "$ip" ]]; do
    sleep 1
    ip="$(host $name | awk '{print $4}' | grep -oE '\b([0-9]{1,3}\.){3}[0-9]{1,3}\b')"
done

if [ -z ${k8s_version} ]; then
    k8s_version=$(curl -fL https://storage.googleapis.com/kubernetes-release/release/stable.txt)
else
    k8s_version=v${k8s_version}
fi

cat <<EOF > /tmp/kubeadm-config.yaml
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

# Deploy VK.
echo ${virtual_kubelet_manifest} | base64 --decode > /tmp/virtual-kubelet.yaml
kubectl apply -f /tmp/virtual-kubelet.yaml
