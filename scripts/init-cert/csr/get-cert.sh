#!/usr/bin/env bash

set -xeuo pipefail

export NODE_NAME=$NODE_NAME
export INTERNAL_IP=$(ip route get 8.8.8.8 | grep src | head -n1 | awk '{print $7}')

echo "generating certificate for DNS \"$NODE_NAME\" IP \"$INTERNAL_IP\" " \
    "in directory \"$CERT_DIR\""

[[ -d "$CERT_DIR" ]]

if [[ -f "$CERT_DIR/$NODE_NAME.key" ]] && [[ -f "$CERT_DIR/$NODE_NAME.crt" ]]; then
    echo "checking existing cert"
    openssl x509 -noout -ext subjectAltName -in "$CERT_DIR/$NODE_NAME.crt" | \
        grep -i "\<$INTERNAL_IP\>" && \
        echo "found valid certificate for $INTERNAL_IP" && exit 0 || \
        echo "invalid certificate \"$CERT_DIR/$NODE_NAME.crt\", recreating it"
fi

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

i=0
CERT_DATA=""
while [[ -z "$CERT_DATA" ]]; do
    i=$((i+1))
    if [[ $i -gt 30 ]]; then
        echo "timeout waiting for CSR $CSR_NAME"
    fi
    sleep 1
    CERT_DATA=$(kubectl get csr $CSR_NAME -ojsonpath='{.status.certificate}')
done

echo "$CERT_DATA" | base64 -d > $NODE_NAME.crt
echo "generated certificate:"
openssl x509 -text -in $NODE_NAME.crt

cp $NODE_NAME.crt $CERT_DIR/
cp $NODE_NAME.key $CERT_DIR/
