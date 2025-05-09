#!/bin/bash

# 嵌入日志
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source ${SCRIPT_DIR}/core/log.sh

case "$2" in
    upmod)

        log INFO "开始定制Kubernetes镜像"
        source ubuntu/kubernetes.sh
        kubernetes "${arch}" "${version}" "${workspace}" "${root_password}" "${username}" "${password}" "${mirror}" "${cleanup}"
        ;;
    --)
        ;;
esac
