#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

main() {
 echo "======DOCKER COMPOSE STATUS DB======"
  echo "================================="
   echo "================================"
   printf "\n\n\n\n\n\n\n"

 docker compose ps
 printf "\n\n\n\n\n\n\n"

 echo "================================================================"
echo "================================================================"
 cd scripts
 sleep 10
 echo "========= AIR SERVICE STATUS for continuous server building =========="
 echo "================================================================"
echo "================================================================"
printf "\n\n\n\n\n\n\n"

 systemctl status --user homelab-dashboard-dev.service --no-pager
 

printf "\n\n\n\n\n\n\n"
journalctl -fu homelab-dashboard-dev.service


}

main "$@"
