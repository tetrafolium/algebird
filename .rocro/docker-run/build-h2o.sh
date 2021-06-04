#!/bin/sh -f

exit_with_message() {
    exit_code=$1
    shift

    echo '###' "$@"
    exit $exit_code
}

REPO_NAME="docker-run-test"
VERSION_TAG="h2o"
LOCAL_TAG="${REPO_NAME}:${VERSION_TAG}"

command -v jq \
    || exit_with_message $? 'please install `jq` command.'

echo +++ "ecr-public describe-repositories --repository-names=${REPO_NAME}"
aws ecr-public describe-repositories --repository-names "${REPO_NAME}" \
    || exit_with_message $? 'please exec `.rocro/docker-run/create-public-repository.sh` first.'

PUBLIC_REG_URI="$(aws ecr-public describe-registries \
                  | jq '.registries[] | .registryUri' \
                  | sed -e 's|"\(.*\)"|\1|')"
PUBLIC_SERVER="$(echo ${PUBLIC_REG_URI} | cut -d / -f 1)"
PUBLIC_TAG="${PUBLIC_REG_URI}/${LOCAL_TAG}"
echo +++ "PUBLIC_REG_URI = ${PUBLIC_REG_URI}"
echo +++ "PUBLIC_SERVER = ${PUBLIC_SERVER}"
echo +++ "PUBLIC_TAG = ${PUBLIC_TAG}"

echo +++ "ecr-public describe-images --repository-name=${REPO_NAME} ... before docker push"
aws ecr-public describe-images --repository-name "${REPO_NAME}"

echo +++ "docker login ..."
aws ecr-public get-login-password \
    | docker login --username AWS --password-stdin "${PUBLIC_SERVER}"

echo +++ "docker build ... ${LOCAL_TAG}"
docker build -f "${VERSION_TAG}.Dockerfile" -t "${LOCAL_TAG}" .

echo +++ "docker push ... ${PUBLIC_TAG}"
docker tag "${LOCAL_TAG}" "${PUBLIC_TAG}"
docker push "${PUBLIC_TAG}"

echo +++ "ecr-public describe-images --repository-name=${REPO_NAME} ... after docker push"
aws ecr-public describe-images --repository-name "${REPO_NAME}"
