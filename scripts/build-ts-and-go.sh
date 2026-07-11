#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

build_ts() {
 tsc --no-emit
}
main() {

}

main "$@"
