#!/bin/bash
set -e

# Running openapi-generator-cli in a docker container to save local PC from dependencies installation

OPENAPI_GENERATOR_CLI="openapitools/openapi-generator-cli:v7.3.0"
SPEC_FILE="/local/docs/api-spec-oapi3.yaml"
GENERATOR_NAME="go-server"
OUT_DIR="/local/internal"

GIT_USER_ID="alewkinr"
GIT_REPO_ID="eth-trx-manager"


export GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w"

docker run --rm \
  -v "${PWD}":/local ${OPENAPI_GENERATOR_CLI} generate \
    -i "${SPEC_FILE}" \
    -g "${GENERATOR_NAME}" \
    --global-property=verbose=true,git-user-id="${GIT_USER_ID}",git-repo-id="${GIT_REPO_ID}" \
    --additional-properties=enablePostProcessFile=true,hideGenerationTimestamp=true,packageName=http,outputAsLibrary=true,router=chi,sourceFolder=http \
    -o "${OUT_DIR}"
#    todo: add onlyInterfaces=true add prop