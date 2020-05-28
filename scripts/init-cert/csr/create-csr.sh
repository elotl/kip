#!/usr/bin/env bash

set -xeuo pipefail

export CSR_NAME=${CSR_NAME:-$NODE_NAME-$(date +%s)}

rm -f $NODE_NAME.key $NODE_NAME.csr $NODE_NAME.crt

openssl genrsa -out $NODE_NAME.key 2048
openssl req -new -key $NODE_NAME.key -out $NODE_NAME.csr -config <(envsubst < csr.conf)

VK_CSR="$(base64 -w0 < $NODE_NAME.csr)" envsubst < csr.yaml
