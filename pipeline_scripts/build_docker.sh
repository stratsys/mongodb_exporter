#!/bin/bash
source pipeline_scripts/common.sh
LATEST_TAG_ARG=""
[[ ! -z "$IMAGE_TAG_LATEST" ]] && LATEST_TAG_ARG="-t "$IMAGE_TAG_LATEST""
set -x
if [[ -z "$DOCKER_BUILD_ARGS" ]] ; then
  docker build --rm -t "$IMAGE_TAG" $LATEST_TAG_ARG .
else
  docker build --rm -t "$IMAGE_TAG" $LATEST_TAG_ARG "${DOCKER_BUILD_ARGS[@]}" .
fi
