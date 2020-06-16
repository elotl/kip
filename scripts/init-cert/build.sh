#!/usr/bin/env bash

set -euxo pipefail

TAG=${TAG:-$(git describe --dirty)}
DKR=${DKR:-docker}
REPO=${REPO:-elotl/init-cert}
PUSH_IMAGES=${PUSH_IMAGES:-true}
UPDATE_LATEST=${UPDATE_LATEST:-false}

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $SCRIPT_DIR

LATEST_K8S_VERSION=$(curl -fsL https://storage.googleapis.com/kubernetes-release/release/stable.txt)
LATEST_K8S_MINOR=$(echo $LATEST_K8S_VERSION | sed -r 's/^v([0-9]+)\.([0-9]+)\..*$/\2/')
K8S_VERSIONS=$(seq -f '1.%g' -s ' ' 14 $LATEST_K8S_MINOR)

docker build -t ${REPO}:${TAG} --build-arg K8S_VERSIONS="${K8S_VERSIONS}" .

if $PUSH_IMAGES; then
    docker push ${REPO}:${TAG}
    if $UPDATE_LATEST; then
        docker tag ${REPO}:${TAG} ${REPO}:latest
        docker push ${REPO}:latest
    fi
fi
