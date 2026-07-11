#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

build_ts() {
 cd ~/DEV/homelab_dashboard/frontend
 tsc --no-emit
}

build_go() {

}
main() {

}

main "$@"
