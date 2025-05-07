#!/usr/bin/sudo /bin/bash

# 嵌入日志
source core/log.sh

chroot() {
    name=$1
    basedir=$2
    commands=${*:2+1}

    log DEBUG 开始执行定制命令 "name=${name}, basedir=${basedir}, commands=${commands}"
    sudo /sbin/chroot "${basedir}" sh -c "${commands}" > /dev/null 2>&1 || exit 1
    log INFO 定制命令成功 "name=${name}, basedir=${basedir}, commands=${commands}"
}
