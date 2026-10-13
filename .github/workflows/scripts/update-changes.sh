#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

main() {
   cd ~/dev/homepage/       
   git checkout main
   git pull origin main
   docker compose up -d
}

main "$@"
