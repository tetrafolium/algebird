#!/bin/sh -f

exit_with_message() {
    exit_code=$1
    shift

    echo '###' "$@"
    exit "$exit_code"
}

REPO_NAME="docker-run-test"

echo +++ "ecr-public describe-reposotiries --repository-name=${REPO_NAME}"
aws ecr-public describe-repositories --repository-name "${REPO_NAME}" \
    || exit_with_message $? 'failed unexpected.'

echo +++ "ecr-public describe-images --repository-name=${REPO_NAME}"
aws ecr-public describe-images --repository-name "${REPO_NAME}" \
    || exit_with_message $? 'failed unexpected.'
