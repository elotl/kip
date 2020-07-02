#!/bin/bash

set -e

if [[ $# -lt 1 ]]; then
    echo "usage $0 <tag>"
    echo "where tag should be of the form v1.2.3"
    exit 1
fi

if [[ $(git rev-parse --abbrev-ref HEAD) != "master" ]]; then
    echo "a release must be made from master"
    exit 1
fi

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=$SCRIPT_DIR/..

version_tag=$1
version_string=${version_tag:1}  # strip off the leading 'v'
if [[ $version_tag =~ ^v[0-9].* ]]; then
    echo $version_string > $ROOT_DIR/version
    git commit --allow-empty -am "release $version_tag"

    git tag -a $version_tag -m "release $version_tag"
    git push --follow-tags origin master
else
    echo "tag must be of the form v1.2.3"
fi
