#!/usr/bin/env bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

BUILD=${TRAVIS_BUILD_NUMBER:-local}
USE_REGION=${USE_REGION:-us-east-1}

handle_error() {
    trap - EXIT
    show_kube_info || true
    cleanup || true
}

cleanup() {
    delete_kube_resources
}

delete_kube_resources() {
    kubectl delete -f $SCRIPT_DIR/eks-smoke-test/kip-nginx.yaml
    kubectl delete -f $SCRIPT_DIR/eks-smoke-test/regular-node-nginx.yaml
    kubectl delete pod test
}

show_kube_info() {
    kubectl get pods -A
    kubectl get nodes
    kubectl -n kip-smoke-tests describe pod -l statefulset.kubernetes.io/pod-name=kip-build-kip-provider-0
    kubectl logs -n kip-smoke-tests -l statefulset.kubernetes.io/pod-name=kip-build-kip-provider-0r -c init-cert --tail=-1
    kubectl logs -n kip-smoke-tests -l statefulset.kubernetes.io/pod-name=kip-build-kip-provider-0 -c kip --tail=-1
}

update_vk() {
    local version="$(git describe)"
    local patch_init="{\"spec\":{\"template\":{\"spec\":{\"initContainers\":[{\"image\":\"elotl/init-cert:$version\",\"name\":\"init-cert\"}]}}}}"
    local patch_kip="{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"image\":\"elotl/kip:$version\",\"name\":\"kip\"}]}}}}"
    kubectl patch -n kip-smoke-tests statefulset kip-build-kip-provider -p "$patch_init"
    kubectl patch -n kip-smoke-tests statefulset kip-build-kip-provider -p "$patch_kip"
}

run_smoke_test_1() {
     kubectl apply -f $SCRIPT_DIR/eks-smoke-test/kip-nginx.yaml
     kubectl apply -f $SCRIPT_DIR/eks-smoke-test/regular-node-nginx.yaml
     kubectl get nodes -o wide
     kubectl get pods -n kip-smoke-tests
#    local curlcmd="i=0; while [ \$i -lt 300 ]; do i=\$((i+1)); curl nginx | grep 'Welcome to nginx' && exit 0; sleep 1; done; exit 1"
#    local waitcmd="phase=\"\"; echo \"Waiting for test results from pod\"; until [[ \$phase = Succeeded ]]; do sleep 1; phase=\$(kubectl get pod test -ojsonpath=\"{.status.phase}\"); if [[ \$phase = Failed ]]; then echo \$phase; kubectl get pods -A; exit 1; fi; echo \$phase; done"
#    kubectl run test --restart=Never --image=elotl/debug --command -- /bin/sh -c "$curlcmd"
#    timeout 420s bash -c "$waitcmd"
}

run_smoke_test_2() {
    local curlcmd="i=0; while [ \$i -lt 300 ]; do i=\$((i+1)); curl nginx | grep 'Welcome to nginx' && exit 0; sleep 1; done; exit 1"
    local waitcmd="phase=\"\"; echo \"Waiting for test results from pod\"; until [[ \$phase = Succeeded ]]; do sleep 1; phase=\$(kubectl get pod test -ojsonpath=\"{.status.phase}\"); if [[ \$phase = Failed ]]; then echo \$phase; kubectl get pods -A; exit 1; fi; echo \$phase; done"
    kubectl run test --restart=Never --image=elotl/debug --command -- /bin/sh -c "$curlcmd"
    timeout 420s bash -c "$waitcmd"
}

fetch_kubeconfig() {
    echo "fetch kubeconfig"
    aws eks update-kubeconfig --name elotl-ci-cd
}

test_once() {
    set -euxo pipefail
    trap handle_error EXIT
    fetch_kubeconfig
    update_vk
    run_smoke_test_1
#    run_smoke_test_2
    cleanup
    trap - EXIT
}

# Run test_once() if running as a shell script, and return if sourced.
(return 0 2>/dev/null) || test_once
