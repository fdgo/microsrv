#!/usr/bin/env bash

function pull {
    IMAGE_NAME="base_srv"
    TAG=$1
    REPOSINAME="registry.cn-hangzhou.aliyuncs.com/jzbro/"
    docker pull "$REPOSINAME$IMAGE_NAME:$TAG"
    docker run -d --rm  --name $IMAGE_NAME  -p 9600:9600  registry.cn-hangzhou.aliyuncs.com/jzbro/$IMAGE_NAME:$1
}
set -e
pull $1
