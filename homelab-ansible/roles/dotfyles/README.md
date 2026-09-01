# dotfiles Ansible role

This role restores the portable dotfiles found in the supplied dotfiles backup.

Included by default:

- `.bashrc`
- `.gitconfig`
- `.profile`
- `.tmux.conf`
- `~/.config/alacritty`
- `~/.config/fish`
- `~/.config/kitty`
- `~/.config/waybar`
- `~/.config/wofi`

The uploaded archive contained empty `~/.config/hypr` and `~/.config/nvim` directories, so those are not included yet. Put those directories under `files/config/` and add `hypr` / `nvim` to `dotfiles_config_dirs` when you have the actual files.

Browser/session/cache-oriented directories such as Falkon profiles, Pulse runtime state, dconf binary state, and credential-oriented directories are intentionally not restored by default.

## Example

```yaml
- hosts: desktops
  roles:
    - role: dotfiles
      vars:
        dotfiles_user: lucassimpson
        dotfiles_home: /home/lucassimpson
```

If the username differs per machine, define `dotfiles_user` in inventory or host vars.

## Add another config directory

Copy it into:

```text
roles/dotfiles/files/config/<name>/
```

Then add the directory name to `dotfiles_config_dirs`.
