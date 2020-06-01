#!/usr/bin/env bash

set -euxo pipefail

DKR=${DKR:-docker}
PUSH_IMAGES=${PUSH_IMAGES:-false}

major=1
for minor in $(seq 12 18); do
    tag="k8s-${major}.${minor}"
    docker build -t elotl/init-cert:${tag} --build-arg K8S_VERSION_MAJOR=${major} --build-arg K8S_VERSION_MINOR=${minor} .
    $PUSH_IMAGES && docker push elotl/init-cert:${tag} || true
done
docker tag elotl/init-cert:k8s-1.18 elotl/init-cert:latest
$PUSH_IMAGES && docker push elotl/init-cert:latest || true
