#!/usr/bin/env bash

set -euo pipefail

echo "generating certificate for \"$NODE_NAME\""
export NODE_NAME=$NODE_NAME

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $SCRIPT_DIR

KUBECONFIG=${KUBECONFIG:-""}
if [[ -z $KUBECONFIG ]]; then
    token2kubeconfig > kubeconfig
    export KUBECONFIG=$(pwd)/kubeconfig
fi

export CSR_NAME="virtual-kubelet-$(date +%s)"

./create-csr.sh | kubectl apply -f -

i=0
while [[ $i -lt 30 ]]; do
    i=$((i+1))
    kubectl get csr $CSR_NAME || {
        sleep 1
        continue
    }
    kubectl certificate approve $CSR_NAME
    break
done

kubectl get csr $CSR_NAME -ojsonpath='{.status.certificate}' | base64 -d > $NODE_NAME.crt
openssl x509 -text -in $NODE_NAME.crt
