#!/bin/bash

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=$SCRIPT_DIR/..
cd $ROOT_DIR

SSH_KEY_FILE=$(mktemp)
chmod 0600 $SSH_KEY_FILE

PORTFW_PID=""

BUILD=${TRAVIS_BUILD_NUMBER:-local}
CLUSTER_NAME="build-$BUILD"

export KUBECONFIG=$(mktemp)

cleanup() {
    kubectl get pods -A || true
    kubectl get nodes || true
    kubectl -n kube-system describe pod -l app=kip-provider || true
    kubectl logs -n kube-system -l app=kip-provider -c init-cert --tail=-1 || true
    kubectl logs -n kube-system -l app=kip-provider -c kip --tail=-1 || true
    kubectl delete pod nginx > /dev/null 2>&1 || true
    kubectl delete svc nginx > /dev/null 2>&1 || true
    kubectl delete pod test > /dev/null 2>&1 || true
    rm -rf $SSH_KEY_FILE
    rm -rf $KUBECONFIG
    if [[ -n "$PORTFW_PID" ]]; then
        kill -9 $PORTFW_PID
    fi
    cd $ROOT_DIR/deploy/terraform-aws
    terraform destroy -var cluster_name=${CLUSTER_NAME} -auto-approve
}

update_vk() {
    local version="$(git describe)"
    local patch_init="{\"spec\":{\"template\":{\"spec\":{\"initContainers\":[{\"image\":\"elotl/init-cert:$version\",\"name\":\"init-cert\"}]}}}}"
    local patch_kip="{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"image\":\"elotl/kip:$version\",\"name\":\"kip\"}]}}}}"
    kubectl patch -n kube-system statefulset kip-provider -p "$patch_init"
    kubectl patch -n kube-system statefulset kip-provider -p "$patch_kip"
}

run_smoke_test() {
    local curlcmd="i=0; while [ \$i -lt 300 ]; do i=\$((i+1)); curl nginx | grep 'Welcome to nginx' && exit 0; sleep 1; done; exit 1"
    local waitcmd="phase=\"\"; echo \"Waiting for test results from pod\"; until [[ \$phase = Succeeded ]]; do sleep 1; phase=\$(kubectl get pod test -ojsonpath=\"{.status.phase}\"); if [[ \$phase = Failed ]]; then echo \$phase; kubectl get pods -A; exit 1; fi; echo \$phase; done"
    kubectl run nginx --image=nginx --port=80
    kubectl expose pod nginx
    kubectl run test --restart=Never --image=elotl/debug --command -- /bin/sh -c "$curlcmd"
    timeout 300s bash -c "$waitcmd"
}

fetch_kubeconfig() {
    local port=$(((RANDOM%999)+30000))
    scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i $SSH_KEY_FILE ubuntu@$SSH_HOST:.kube/config $KUBECONFIG
    sed -i "s#https://.*:.*#https://127.0.0.1:$port#g" $KUBECONFIG
    ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i $SSH_KEY_FILE -N -L $port:localhost:6443 ubuntu@$SSH_HOST &
    PORTFW_PID="$!"
    # Wait for connection.
    timeout 60s sh -c "while true; do kubectl get sa default > /dev/null 2>&1 && exit 0; done"
}

trap cleanup EXIT

# Create a test cluster.
cd $ROOT_DIR/deploy/terraform-aws
terraform init
terraform apply -var cluster_name=${CLUSTER_NAME} -auto-approve
terraform show -json | jq -r '.values.root_module.resources | .[] | select(.address=="tls_private_key.ssh_key") | .values.private_key_pem' > $SSH_KEY_FILE
SSH_HOST=$(terraform show -json | jq -r '.values.root_module.resources | .[] | select(.address=="aws_instance.k8s_node") | .values.public_ip')

fetch_kubeconfig

update_vk

run_smoke_test
