#!/usr/bin/env bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=$SCRIPT_DIR/..

BUILD=${TRAVIS_BUILD_NUMBER:-local}
CLUSTER_NAME="build-$BUILD"
USE_REGION=${USE_REGION:-us-east-1}
STATE_BUCKET=${STATE_BUCKET:-elotl-tf-state}
STATE_PATH=${STATE_PATH:-build/terraform-${BUILD}.tfstate}

setup() {
    SSH_KEY_FILE=$(mktemp)
    chmod 0600 $SSH_KEY_FILE
    export KUBECONFIG=$(mktemp)
    TFDIR=$(mktemp -d)
    PORTFW_PID=""
}

handle_error() {
    trap - EXIT
    show_kube_info || true
    cleanup || true
}

cleanup() {
    delete_kube_resources
    destroy_cluster
    delete_temp_files
    stop_port_forwarding
}

stop_port_forwarding() {
    [[ -n "$PORTFW_PID" ]] && kill -9 $PORTFW_PID || true
}

delete_temp_files() {
    [[ -n "$SSH_KEY_FILE" ]] && rm -rf $SSH_KEY_FILE
    [[ -n "$KUBECONFIG" ]] && rm -rf $KUBECONFIG
    [[ -n "$TFDIR" ]] && rm -rf $TFDIR
}

delete_kube_resources() {
    kubectl delete pod nginx
    kubectl delete svc nginx
    kubectl delete pod test
}

show_kube_info() {
    kubectl get pods -A
    kubectl get nodes
    kubectl -n kube-system describe pod -l app=kip-provider
    kubectl logs -n kube-system -l app=kip-provider -c init-cert --tail=-1
    kubectl logs -n kube-system -l app=kip-provider -c kip --tail=-1
}

create_cluster() {
    # Create a test cluster.
    pushd $TFDIR
    cat > main.tf <<EOF
terraform {
  required_version = "~> 0.12"
  backend "s3" {
    region  = "${USE_REGION}"
    bucket  = "${STATE_BUCKET}"
    key     = "${STATE_PATH}"
    encrypt = "true"
  }
}
module "kip-aws" {
  source        = "${ROOT_DIR}/deploy/terraform-aws"
  cluster_name  = "${CLUSTER_NAME}"
}
EOF
    terraform init
    terraform apply -auto-approve
    terraform show -json | jq -r '.values.root_module.child_modules[] | select(.address=="module.kip-aws") | .resources[] | select(.address=="tls_private_key.ssh_key") | .values.private_key_pem' > $SSH_KEY_FILE
    SSH_HOST=$(terraform show -json | jq -r '.values.root_module.child_modules[] | select(.address=="module.kip-aws") | .resources[] | select(.address=="aws_instance.k8s_node") | .values.public_ip')
    popd
}

destroy_cluster() {
    pushd $TFDIR
    terraform destroy -auto-approve
    popd
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

test_once() {
    set -euxo pipefail
    trap handle_error EXIT
    setup
    create_cluster
    fetch_kubeconfig
    update_vk
    run_smoke_test
    cleanup
    trap - EXIT
}

# Run test_once() if running as a shell script, and return if sourced.
(return 0 2>/dev/null) || test_once
