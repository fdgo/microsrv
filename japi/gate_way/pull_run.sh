#!/usr/bin/env bash

function pull {
    IMAGE_NAME="gate_way"
    TAG=$1
    REPOSINAME="registry.cn-hangzhou.aliyuncs.com/jzbro/"
    docker pull "$REPOSINAME$IMAGE_NAME:$TAG"
    docker run -d --rm  --name gate_way  -p 8080:8080 registry.cn-hangzhou.aliyuncs.com/jzbro/gate_way:$1  --registry=consul --registry_address=127.0.0.1:8500  --api_namespace=jz.micro.jzapi.web    api -handler=http
}
set -e
pull $1
