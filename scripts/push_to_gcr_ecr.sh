#!/usr/bin/env bash

set -exuo pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=$SCRIPT_DIR/..
cd $ROOT_DIR

if [[ $(git describe --dirty) =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)$ ]]; then
    gcloud auth configure-docker --quiet
    `aws ecr get-login`
    make img push-img DKR=docker REGISTRY_REPO=gcr.io/elotl-kip/virtual-kubelet
    make img push-img DKR=docker REGISTRY_REPO=689494258501.dkr.ecr.us-east-1.amazonaws.com/elotl-kip/virtual-kubelet
fi
