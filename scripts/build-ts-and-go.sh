#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

build_ts() {

 npm run --prefix frontend build
}

build_go() {
    go build -o ./tmp/server .
}
main() {


        build_ts
        build_go

}

main "$@"
