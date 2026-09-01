#!/bin/sh
swayidle -w \
    timeout 900 'swaylock -f -i ~/Pictures/Wallpapers/portal.png' \
    timeout 1800 'systemctl suspend' \
    before-sleep 'swaylock -f -c 000000' &
