#!/bin/bash

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=$SCRIPT_DIR/..
cd $ROOT_DIR

function run_with_iam_keys {
    # IAM_ keys are set in the jenkins UI. Those keys are linked to
    # the milpaci user in AWS and come with the minimum set of
    # permissions needed to run milpa that we publish with our docs.
    # this helps ensure we keep our publised IAM permissions up to
    # date with our build.
    AWS_ACCESS_KEY_ID=$IAM_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY=$IAM_SECRET_ACCESS_KEY $@
}

echo "Running unit tests"

go test ./...

# azure functional tests disabled 6/4/19, they're a bit flakey
# they're time consuming and we aren't targeting azure now

declare -a FUNCTIONAL_TEST_DIRS=("pkg/server/cloud/aws")
                                 # "pkg/server/cloud/azure")

# Run functional tests in parallel.
i=0
declare -a fntest_pids
for dirname in "${FUNCTIONAL_TEST_DIRS[@]}"; do
    echo "Running functional tests in $dirname"
    (set -e
    cd ${ROOT_DIR}/$dirname
    run_with_iam_keys go test -v -functional -timeout 800s) &
    fntest_pids[${i}]=$!
    i=$((i+1))
done

# Wait for functional tests to finish.
echo "Functional tests running, pids: ${fntest_pids[*]}"
for pid in ${fntest_pids[*]}; do
    echo "Waiting for pid $pid"
    wait $pid
    echo "Pid $pid finished"
done
