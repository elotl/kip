#!/usr/bin/env bash

set -xeuo pipefail

export NODE_NAME=$NODE_NAME
export INTERNAL_IP=$(ip route get 8.8.8.8 | grep src | head -n1 | awk '{print $7}')

echo "generating certificate for DNS \"$NODE_NAME\" IP \"$INTERNAL_IP\" " \
    "in directory \"$CERT_DIR\""

[[ -d "$CERT_DIR" ]]

if [[ -f "$CERT_DIR/$NODE_NAME.key" ]] && [[ -f "$CERT_DIR/$NODE_NAME.crt" ]]; then
    echo "checking existing cert"
    openssl x509 -checkip $INTERNAL_IP -checkhost $NODE_NAME -noout -in "$CERT_DIR/$NODE_NAME.crt" | \
        grep "does NOT match certificate" && \
        echo "certificate \"$CERT_DIR/$NODE_NAME.crt\" not valid, recreating it" || \
        exit 0
fi

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $SCRIPT_DIR

KUBECONFIG=${KUBECONFIG:-""}
if [[ -z $KUBECONFIG ]]; then
    token2kubeconfig > kubeconfig
    export KUBECONFIG=$(pwd)/kubeconfig
fi

export CSR_NAME="${NODE_NAME}-$(date +%s)"

./create-csr.sh | kubectl apply -f -

approve_csr() {
    while true; do
        kubectl get csr $CSR_NAME || {
            sleep 1
            continue
        }
        kubectl certificate approve $CSR_NAME
        break
    done
}
export -f approve_csr
timeout 30s bash -c approve_csr

fetch_cert() {
    CERT_DATA=""
    while [[ -z "$CERT_DATA" ]]; do
        CERT_DATA=$(kubectl get csr $CSR_NAME -ojsonpath='{.status.certificate}')
        sleep 1
    done
    echo "$CERT_DATA" | base64 -d > $NODE_NAME.crt
}
export -f fetch_cert
timeout 30s bash -c fetch_cert

echo "generated certificate:"
openssl x509 -text -in $NODE_NAME.crt

cp $NODE_NAME.crt $CERT_DIR/
cp $NODE_NAME.key $CERT_DIR/
