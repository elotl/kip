#!/bin/bash -v

curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
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

cat <<EOF > /tmp/kubeadm-config.yaml
apiVersion: kubeadm.k8s.io/v1beta1
kind: JoinConfiguration
discovery:
  bootstrapToken:
    token: ${k8stoken}
    unsafeSkipCAVerification: true
    apiServerEndpoint: ${masterIP}:6443
nodeRegistration:
  name: $name
  kubeletExtraArgs:
    node-ip: $ip
    cloud-provider: aws
EOF

kubeadm join --config=/tmp/kubeadm-config.yaml
