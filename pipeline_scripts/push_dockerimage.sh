#!/bin/bash
source pipeline_scripts/common.sh
set -x
docker push "$IMAGE_TAG" 
if [[ ! -z "$IMAGE_TAG_LATEST" ]]; then
  docker push "$IMAGE_TAG_LATEST"
fi
