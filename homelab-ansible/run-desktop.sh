#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

main() {
    ansible-playbook \
  -i inventory.ansible.yaml \
  prometheus.ansible.yaml \
  --ask-become-pass \
  --ask-vault-pass  
}

main "$@"
