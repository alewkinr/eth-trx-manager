#!/bin/bash
set -e

# Running openapi-generator-cli in a docker container to save local PC from dependencies installation

OPENAPI_GENERATOR_CLI="openapitools/openapi-generator-cli:v7.3.0"
SPEC_FILE="/local/docs/api-spec-oapi3.yaml"
LANG="go"
OUT_DIR="/local/internal/http"

GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w"

docker run --rm \
  -v "${PWD}":/local ${OPENAPI_GENERATOR_CLI} generate \
    -i "${SPEC_FILE}" \
    -g "${LANG}" \
    --global-property=verbose=true \
    --additional-properties=hideGenerationTimestamp=true,packageName=http,withGoMod=false,isGoSubmodule=true,generateInterfaces=true \
    -o "${OUT_DIR}"