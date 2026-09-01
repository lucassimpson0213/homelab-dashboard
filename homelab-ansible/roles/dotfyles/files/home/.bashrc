# ~/.bashrc
# ============================================================

# Source global definitions
if [ -f /etc/bashrc ]; then
    . /etc/bashrc
fi

# ------------------------------------------------------------
# PATH setup (safe, idempotent)
# ------------------------------------------------------------
if [[ ":$PATH:" != *":$HOME/.local/bin:"* ]]; then
    PATH="$HOME/.local/bin:$PATH"
fi

if [[ ":$PATH:" != *":$HOME/bin:"* ]]; then
    PATH="$HOME/bin:$PATH"
fi

export PATH

# Cargo
[ -f "$HOME/.cargo/env" ] && . "$HOME/.cargo/env"
export VISUAL=nvim
export EDITOR=nvim

# Go
export PATH="/usr/local/go/bin:$PATH"

# ------------------------------------------------------------
# Only continue for interactive shells
# ------------------------------------------------------------
[[ -n $PS1 ]] || return

# ------------------------------------------------------------
# Environment
# ------------------------------------------------------------
export EDITOR='nvim'
export VISUAL='vim'
export PAGER='less'
export TZ='America/New_York'
export HISTCONTROL='ignoredups'
export HISTSIZE=5000
export HISTFILESIZE=5000
export GREP_COLOR='1;36'
export LSCOLORS='ExGxbEaECxxEhEhBaDaCaD'

# ------------------------------------------------------------
# Less colors
# ------------------------------------------------------------
export LESS_TERMCAP_mb=$(
    tput bold
    tput setaf 1
)
export LESS_TERMCAP_md=$(
    tput bold
    tput setaf 1
)
export LESS_TERMCAP_me=$(tput sgr0)
export LESS_TERMCAP_se=$(tput sgr0)
export LESS_TERMCAP_so=$(
    tput bold
    tput setaf 3
    tput setab 4
)
export LESS_TERMCAP_ue=$(tput sgr0)
export LESS_TERMCAP_us=$(
    tput smul
    tput bold
    tput setaf 2
)

# ------------------------------------------------------------
# Shell options
# ------------------------------------------------------------
shopt -s cdspell checkwinsize extglob
shopt -s autocd 2>/dev/null || true
shopt -s dirspell 2>/dev/null || true

# ------------------------------------------------------------
# Aliases
# ------------------------------------------------------------
alias ..='cd ..'
alias bnix='nix-shell -p \
  pkgsCross.i686-embedded.buildPackages.gccWithoutTargetLibc \
  gdb \
  qemu'
alias l='ls'
alias ll='ls -lha'
alias ag='rg'
alias hl='rg --passthru'
alias gerp='grep'
alias chomd='chmod'
alias suod='sudo'
alias externalip='curl -sS https://ysap.sh/ip'

grep --color=auto </dev/null &>/dev/null && alias grep='grep --color=auto'
xdg-open --version &>/dev/null && alias open='xdg-open'

# ------------------------------------------------------------
# Git aliases
# ------------------------------------------------------------
alias ga='git add .'
alias gb='git branch'
alias gco='git checkout'
alias gcm='git checkout main && git pull'
alias gd='git diff'
alias gs='git status'
alias gp='git push origin HEAD'

# ------------------------------------------------------------
# Git helper functions
# ------------------------------------------------------------
gmb() {
    local main
    main=$(git symbolic-ref --short refs/remotes/origin/HEAD 2>/dev/null)
    main=${main#origin/}
    [[ -n $main ]] && echo "$main"
}

# ------------------------------------------------------------
# Prompt (clean + fast)
# ------------------------------------------------------------
COLOR256=()
COLOR256[0]=$(tput setaf 1)
COLOR256[256]=$(tput sgr0)
COLOR256[257]=$(tput bold)

PROMPT_COLORS=()

set_prompt_colors() {
    local h=${1:-0} i j=0
    for i in {22..231}; do
        ((i % 30 == h)) || continue
        PROMPT_COLORS[$j]=$(tput setaf "$i")
        ((j++))
    done
}

PS1='\u@\h \w $(git rev-parse --abbrev-ref HEAD 2>/dev/null | sed "s/.*/(git:&)/")\$ '
set_prompt_colors 24

PROMPT_COMMAND='printf "\033]0;%s@%s:%s\007" "$USER" "${HOSTNAME%%.*}" "${PWD/#$HOME/~}"'

PROMPT_DIRTRIM=6

# ------------------------------------------------------------
# VM display (INTENTIONALLY kept)
# ------------------------------------------------------------

# ------------------------------------------------------------
# Completions
# ------------------------------------------------------------
[ -f /etc/bash/bash_completion ] && . /etc/bash/bash_completion
[ -f ~/.bash_completion ] && . ~/.bash_completion

# ------------------------------------------------------------
# Local overrides (optional)
# ------------------------------------------------------------
[ -f ~/.bash_aliases ] && . ~/.bash_aliases
[ -f ~/.bashrc.local ] && . ~/.bashrc.local
alias alacritty='env LIBGL_ALWAYS_SOFTWARE=1 alacritty'
