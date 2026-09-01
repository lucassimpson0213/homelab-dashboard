#!/usr/bin/env bash

set -u

OUT="${HOME}/migration-audit-$(date +%Y%m%d-%H%M%S).txt"

section() {
    printf '\n\n========== %s ==========\n' "$1" | tee -a "$OUT"
}

run() {
    "$@" 2>&1 | tee -a "$OUT"
}

: > "$OUT"

section "SYSTEM"
run hostnamectl
run uname -a

section "USER-INSTALLED RPM PACKAGES"
if command -v dnf >/dev/null; then
    run dnf repoquery --userinstalled --qf '%{name}'
fi

section "FLATPAKS"
if command -v flatpak >/dev/null; then
    run flatpak list --app
fi

section "ENABLED SYSTEM SERVICES"
run systemctl list-unit-files --state=enabled --no-pager

section "ENABLED USER SERVICES"
run systemctl --user list-unit-files --state=enabled --no-pager

section "SYSTEM TIMERS"
run systemctl list-timers --all --no-pager

section "USER TIMERS"
run systemctl --user list-timers --all --no-pager

section "CRON"
crontab -l 2>&1 | tee -a "$OUT"

section "IMPORTANT CONFIG DIRECTORIES"
for path in \
    "$HOME/.config" \
    "$HOME/.local/bin" \
    "$HOME/.ssh" \
    "$HOME/.gnupg"
do
    if [[ -e "$path" ]]; then
        du -sh "$path" 2>/dev/null | tee -a "$OUT"
        find "$path" -maxdepth 2 -mindepth 1 2>/dev/null | tee -a "$OUT"
    fi
done

section "COMMON DOTFILES"
find "$HOME" -maxdepth 1 -type f -name '.*' -print 2>/dev/null |
    sort |
    tee -a "$OUT"

section "GIT REPOSITORIES"
while IFS= read -r gitdir; do
    repo="${gitdir%/.git}"

    printf '\n--- %s ---\n' "$repo" | tee -a "$OUT"

    git -C "$repo" status --short 2>&1 | tee -a "$OUT"

    upstream="$(git -C "$repo" rev-parse --abbrev-ref '@{upstream}' 2>/dev/null || true)"

    if [[ -n "$upstream" ]]; then
        git -C "$repo" rev-list --left-right --count \
            "${upstream}...HEAD" 2>&1 | tee -a "$OUT"
    else
        echo "NO UPSTREAM CONFIGURED" | tee -a "$OUT"
    fi
done < <(
    find "$HOME" \
        -path "$HOME/.cache" -prune -o \
        -path "$HOME/.local/share/Trash" -prune -o \
        -type d -name .git -print 2>/dev/null
)

section "PODMAN CONTAINERS"
if command -v podman >/dev/null; then
    run podman ps -a
fi

section "PODMAN VOLUMES"
if command -v podman >/dev/null; then
    run podman volume ls
fi

section "DOCKER CONTAINERS"
if command -v docker >/dev/null; then
    run docker ps -a
fi

section "DOCKER VOLUMES"
if command -v docker >/dev/null; then
    run docker volume ls
fi

section "LARGE FILES IN HOME (>500 MB)"
find "$HOME" \
    -path "$HOME/.cache" -prune -o \
    -type f -size +500M \
    -printf '%s %p\n' 2>/dev/null |
    sort -nr |
    numfmt --field=1 --to=iec |
    tee -a "$OUT"

section "HOME DIRECTORY SIZE"
run du -sh "$HOME"

printf '\n\nAudit written to:\n%s\n' "$OUT"
