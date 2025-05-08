#!/bin/bash

# 嵌入日志
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${SCRIPT_DIR}/core/log.sh"

for dir in ./*/; do
    if [ -f "${dir}/go.mod" ]; then
        log INFO 正在处理目录 "directory=${dir}"
        (cd "$dir" && go mod tidy && go get -u && cd ..)
    else
        log WARN 非模块化项目 "directory=${dir}"
    fi
done
