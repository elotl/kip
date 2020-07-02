#!/usr/bin/env bash
#
# This script is run by the build when all tests have passed. If we are on a
# release tag, then it will update the :latest tags for elotl/kip and
# elotl/init-cert.
#

set -exuo pipefail

TAG=${TAG:-$(git describe --tags --dirty)}
DKR=${DKR:-docker}

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=$SCRIPT_DIR/..
cd $ROOT_DIR

if [[ $TAG =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)$ ]]; then
    $DKR tag elotl/kip:$TAG elotl/kip:latest
    $DKR tag elotl/init-cert:$TAG elotl/init-cert:latest
fi
