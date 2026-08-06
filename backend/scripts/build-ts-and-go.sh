#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

PROJECT_ROOT="$HOME/DEV/homelab-dashboard"
FRONTEND_DIR="$PROJECT_ROOT/frontend"
BACKEND_DIR="$PROJECT_ROOT/backend"
TMP_DIR="$BACKEND_DIR/tmp"
SERVER_BIN="$TMP_DIR/server"


build_ts() {
   npm run --prefix "$FRONTEND_DIR" build
}

build_go() {
    echo "building go"
    mkdir -p "$TMP_DIR"
    go build -o "$SERVER_BIN" "$BACKEND_DIR"
    chmod +x "$SERVER_BIN"
}
main() {


        build_ts
        build_go

}

main "$@"
