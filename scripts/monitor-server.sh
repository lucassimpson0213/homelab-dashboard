#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

main() {
 systemctl status --user homelab-dashboard-dev.service
 journalctl --user -fu homelab-dashboard-dev.service
}

main "$@"
