#!/usr/bin/env bash

set -euxo pipefail

readonly curlcmd_test_1="i=0; while [ \$i -lt 300 ]; do i=\$((i+1)); curl regular-nginx.kip-smoke-tests | grep 'Welcome to nginx' && exit 0; sleep 1; done; exit 1"
readonly waitcmd_test_1="phase=\"\"; echo \"Waiting for test results from pod\"; until [[ \$phase = Succeeded ]]; do sleep 1; phase=\$(kubectl get pod -n kip-smoke-tests test -ojsonpath=\"{.status.phase}\"); if [[ \$phase = Failed ]]; then echo \$phase; kubectl get pods -A; exit 1; fi; echo \$phase; done"
readonly curlcmd_test_2="i=0; while [ \$i -lt 300 ]; do i=\$((i+1)); curl kip-nginx.kip-smoke-tests | grep 'Welcome to nginx' && exit 0; sleep 1; done; exit 1"
readonly waitcmd_test_2="phase=\"\"; echo \"Waiting for test results from pod\"; until [[ \$phase = Succeeded ]]; do sleep 1; phase=\$(kubectl get pod -n kip-smoke-tests test -ojsonpath=\"{.status.phase}\"); if [[ \$phase = Failed ]]; then echo \$phase; kubectl get pods -A; exit 1; fi; echo \$phase; done"


if [ "$#" -ne 1 ] || ! [ -d "$1" ]
then
    echo "Usage: $0 path/to/yaml/"
    exit
fi
readonly yaml_dir=$1

handle_error() {
    trap - EXIT
    show_kube_info || true
    cleanup || true
}

cleanup() {
    delete_kube_resources
}

delete_kube_resources() {
    kubectl delete -f $yaml_dir/kip-nginx.yaml
    kubectl delete -f $yaml_dir/regular-node-nginx.yaml
    kubectl delete -n kip-smoke-tests pod test --ignore-not-found
    kubectl delete pod test --ignore-not-found
}

show_kube_info() {
    kubectl get pods -A
    kubectl get nodes
    kubectl -n kip-smoke-tests describe pod -l statefulset.kubernetes.io/pod-name=kip-build-kip-provider-0
    kubectl logs -n kip-smoke-tests -l statefulset.kubernetes.io/pod-name=kip-build-kip-provider-0 -c init-cert --tail=-1
    kubectl logs -n kip-smoke-tests -l statefulset.kubernetes.io/pod-name=kip-build-kip-provider-0 -c kip --tail=-1
    kubectl -n kip-smoke-tests describe pod test
    kubectl logs -n kip-smoke-tests test --tail=-1
    kubectl -n kip-smoke-tests describe po regular-nginx
    kubectl -n kip-smoke-tests describe po kip-nginx
}

update_vk() {
    local version="$(git describe)"
    local patch_init="{\"spec\":{\"template\":{\"spec\":{\"initContainers\":[{\"image\":\"elotl/init-cert:$version\",\"name\":\"init-cert\"}]}}}}"
    local patch_kip="{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"image\":\"elotl/kip:$version\",\"name\":\"kip\"}]}}}}"
    kubectl patch -n kip-smoke-tests statefulset kip-build-kip-provider -p "$patch_init"
    kubectl patch -n kip-smoke-tests statefulset kip-build-kip-provider -p "$patch_kip"
    kubectl taint node kip-build-kip-provider-0 usage=kip-smoke-tests:NoSchedule --overwrite
}

run_smoke_test_1() {
    kubectl apply -f $yaml_dir/regular-node-nginx.yaml
    kubectl run test --restart=Never --namespace=kip-smoke-tests --image=elotl/debug --command -- /bin/sh -c "$curlcmd_test_1"
    timeout 420s bash -c "$waitcmd_test_1"
    kubectl delete -n kip-smoke-tests pod test --ignore-not-found
}

run_smoke_test_2() {
    kubectl apply -f $yaml_dir/kip-nginx.yaml
    kubectl run test --restart=Never --namespace=kip-smoke-tests --image=elotl/debug --command -- /bin/sh -c "$curlcmd_test_2"
    timeout 420s bash -c "$waitcmd_test_2"
    kubectl delete -n kip-smoke-tests pod test --ignore-not-found
}

fetch_kubeconfig() {
    echo "fetch kubeconfig"
    aws eks update-kubeconfig --name elotl-ci-cd
}

test_once() {
    trap handle_error EXIT
    fetch_kubeconfig
    update_vk
    run_smoke_test_1
    run_smoke_test_2
    cleanup
    trap - EXIT
}

test_once
