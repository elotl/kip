#!/bin/bash

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=$SCRIPT_DIR/..
cd $ROOT_DIR

function cleanup() {
    cd $ROOT_DIR/deploy/terraform
    terraform destroy -var ssh-key-name=vilmos -var cluster-name=build-${TRAVIS_BUILD_NUMBER} -auto-approve
}

wait_for_vk() {
    # TODO
    return
}

trap cleanup EXIT

cd $ROOT_DIR/deploy/terraform
terraform init
terraform apply -var ssh-key-name=vilmos -var cluster-name=build-${TRAVIS_BUILD_NUMBER} -auto-approve
wait_for_vk
