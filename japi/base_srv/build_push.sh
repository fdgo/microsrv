#!/usr/bin/env bash

function build_image {
    IMAGE_NAME=$1
    TAG=$2
    IMAGE_FULL_NAME="$IMAGE_NAME:$TAG"
    REPOSINAME="registry.cn-hangzhou.aliyuncs.com/jzbro/"
    echo "building docker image..."
    docker build -t $IMAGE_FULL_NAME .
    docker tag $IMAGE_FULL_NAME "$REPOSINAME$IMAGE_NAME:$TAG"
    docker push  "$REPOSINAME$IMAGE_NAME:$TAG"
    docker rmi $IMAGE_FULL_NAME
    docker rmi "$REPOSINAME$IMAGE_NAME:$TAG"
}
set -e

APPNAME="base_srv"
DATETAG=$(date +"%m-%d_%H-%M-%S")

echo "building application begin..."
CGO_ENABLED=0 GOOS=linux go build  -o $APPNAME main.go plugin.go
echo "building application end..."



build_image $APPNAME $DATETAG