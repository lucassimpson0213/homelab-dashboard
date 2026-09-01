# danmwallace.fedora.hyprland

Install and configure the Hyprland tiling Wayland compositor on Fedora, along with a
curated set of supporting desktop tools (SDDM, Waybar, Alacritty, Wofi, nwg-drawer,
nwg-dock-hyprland) and a themed set of user dotfiles.

## Requirements

- Ansible >= 2.16
- Target host running Fedora (41, 42, or 43)
- `community.general` collection (for the `copr` module)
- A pre-existing local user account on the target — the role does not create users
- Privilege escalation (`become: true`) — installs system packages, enables `sddm`,
  and switches the default systemd target to `graphical.target`

The role enables the `lionheartp/Hyprland` COPR repository and installs from it.

## Role Variables

| Name             | Type | Required | Default                                 | Description                                                                 |
|------------------|------|----------|-----------------------------------------|-----------------------------------------------------------------------------|
| `hyprland_user`  | str  | no       | `{{ ansible_facts['env']['USER'] }}`    | Local user that owns the rendered dotfiles under `/home/<user>/.config/`.   |
| `hyprland_theme` | str  | no       | `nord`                                  | Theme palette. One of: `monochrome`, `nord`, `tokyo-night`.                 |

The `hyprland_theme` value selects a palette file under `vars/themes/<theme>.yml`,
which exposes the `hyprland_palette` dict consumed by the role's Jinja templates.

The `hyprland_user` default reads `$USER` on the controller, which is rarely the
right value for a remote host. Override it explicitly in production.

## Dependencies

None.

## Example Playbook

```yaml
- name: Set up a Hyprland workstation
  hosts: workstations
  become: true
  roles:
    - role: danmwallace.fedora.hyprland
      vars:
        hyprland_user: dwallace
        hyprland_theme: tokyo-night
```

## Testing

A Molecule scenario lives under `molecule/default/` using the ansible-native
delegated driver with a Fedora podman container. Note that this role installs
a large desktop package set (Firefox, LibreOffice, Thunderbird, etc.) and enables
SDDM as a systemd service — a fully passing `converge` and `verify` in a
container may require additional tuning; for high-fidelity validation run the
role against a real Fedora VM.

```bash
molecule test -s default
```

## License

MIT
