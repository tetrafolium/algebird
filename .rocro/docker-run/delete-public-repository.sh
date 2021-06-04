#!/bin/sh -f

exit_with_message() {
    exit_code=$1
    shift

    echo '###' "$@"
    exit "$exit_code"
}

[ $# -gt 1 ] && [ "$1" = '--force' ] && OPTS='-force'

REPO_NAME="docker-run-test"

echo +++ "ecr-public delete-repository ${OPTS} --repository=name=${REPO_NAME}"
aws ecr-public delete-repository ${OPTS} --repository-name "${REPO_NAME}" \
    || exit_with_message $? "${REPO_NAME}:" 'cannot delete (not empty). try with "--force" option.'
