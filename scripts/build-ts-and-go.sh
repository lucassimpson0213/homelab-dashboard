#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

build_ts() {
 cd ~/DEV/homelab-dashboard/frontend/
 npm run  build
}

build_go() {
    go run ~/DEV/homelab-dashboard/main.go
}
main() {


        build_ts
        build_go

}

main "$@"
