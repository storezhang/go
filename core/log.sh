#!/bin/bash

log() {
    timestamp=$(date +'%Y-%m-%d %H:%M:%S')
    case "$1" in
        debug|DEBUG)
            printf "\e[7;36m %s [DEBUG]   %s \e[0m {%s}\n" "${timestamp}" "$2" "$3"
          ;;
        info|INFO)
            printf "\e[7;32m %s [INFO]    %s \e[0m {%s}\n" "${timestamp}" "$2" "$3"
          ;;
        warn|WARN)
            printf "\e[7;33m %s [WARNING] %s \e[0m {%s}\n" "${timestamp}" "$2" "$3"
          ;;
        error|ERROR)
            printf "\e[7;31m %s [ERROR]   %s \e[0m {%s}\n" "${timestamp}" "$2" "$3"
          ;;
        *)
          ;;
    esac
}