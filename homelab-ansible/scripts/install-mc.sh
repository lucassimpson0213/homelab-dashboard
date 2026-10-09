#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

main() {
    curl --progress-bar -L https://dl.min.io/aistor/mc/release/linux-amd64/mc -o mc
chmod +x ./mc

sudo mv ./mc /usr/local/bin/
mc --version  
}

main "$@"
