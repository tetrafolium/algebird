#!/bin/sh -f

exit_with_message() {
    exit_code=$1
    shift

    echo '###' "$@"
    exit "$exit_code"
}

REPO_NAME="docker-run-test"

echo +++ "ecr-public create-repository --repository=name=${REPO_NAME}"
aws ecr-public create-repository --repository-name "${REPO_NAME}" \
    || exit_with_message $? "${REPO_NAME}:" 'cannot create (probably already exists).'
