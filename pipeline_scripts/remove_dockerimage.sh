#!/bin/sh
source pipeline_scripts/common.sh
set -x
docker rmi "$IMAGE_TAG" $IMAGE_TAG_LATEST
