#!/bin/bash

set -e

REGISTRY="ghcr.io/alewkinr"
REPOSITORY="eth-trx-manager"
if [ -z "$IMAGE_TAG" ]; then
  IMAGE_TAG=$(git rev-parse --short HEAD)
fi

docker build -t ${REGISTRY}/${REPOSITORY}:${IMAGE_TAG} -t ${REGISTRY}/${REPOSITORY}:latest -f deploy/Dockerfile .