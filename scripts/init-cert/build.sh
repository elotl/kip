#!/usr/bin/env bash

set -euxo pipefail

TAG=${TAG:-$(git describe --dirty)}
DKR=${DKR:-docker}
REPO=${REPO:-elotl/init-cert}
PUSH_IMAGES=${PUSH_IMAGES:-true}
UPDATE_LATEST=${UPDATE_LATEST:-false}

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $SCRIPT_DIR

docker build -t ${REPO}:${TAG} --build-arg K8S_VERSIONS="1.14 1.15 1.16 1.17 1.18" .

if $PUSH_IMAGES; then
    docker push ${REPO}:${TAG}
    if $UPDATE_LATEST; then
        docker tag ${REPO}:${TAG} ${REPO}:latest
        docker push ${REPO}:latest
    fi
fi
