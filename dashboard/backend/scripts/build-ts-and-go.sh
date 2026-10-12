#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

PROJECT_ROOT="$HOME/DEV/homelab-dashboard"
FRONTEND_DIR="$PROJECT_ROOT/frontend"
BACKEND_DIR="$PROJECT_ROOT/backend"
TMP_DIR="$BACKEND_DIR/tmp"
SERVER_BIN="$TMP_DIR/server"

dirs=("$PROJECT_ROOT" "$FRONTEND_DIR" "$TMP_DIR" "$SERVER_BIN")

inspect_vars() {
   for dir in "${dirs[@]}" 
   do 
      printf '%s \n' "$dir"
   done
}

build_ts() {
   npm run --prefix "$FRONTEND_DIR" build
}

build_go() {
    cd "$BACKEND_DIR"
    echo "building go"
    mkdir -p "$TMP_DIR"
    go build -o "$SERVER_BIN" "$BACKEND_DIR"
    chmod +x "$SERVER_BIN"
}
main() {

        inspect_vars
        build_ts
        build_go

}

main "$@"
