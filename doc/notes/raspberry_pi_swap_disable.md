# Raspberry Pi swap disable and rollback

This note records how to disable swap on the Zero and how to reverse it.

## Context

The device runs Raspberry Pi OS Lite based on Debian 13 (`trixie`).

On this system, swap is managed by `rpi-swap` as a `zram+file` setup. The active swap device appears as:

```text
/dev/zram0
```

The device has about `416 MiB` of RAM, so disabling swap can make the system terminate processes sooner under memory pressure.

## Disable swap

Create a local `rpi-swap` override:

```bash
sudo mkdir -p /etc/rpi/swap.conf.d
sudo nano /etc/rpi/swap.conf.d/99-disable-swap.conf
```

Add this content:

```ini
[Main]
Mechanism=none
```

Stop the current swap unit immediately:

```bash
sudo systemctl stop dev-zram0.swap
```

Reboot to verify the persistent configuration:

```bash
sudo reboot
```

## Check current state

After reboot, verify that no swap is active:

```bash
swapon --show
free -h
systemctl list-units --quiet --type=swap
```

Expected result:

- `swapon --show` prints nothing.
- `free -h` shows `Swap: 0B`.
- `systemctl list-units --quiet --type=swap` prints no active swap unit.

## Restore swap

Remove the local override:

```bash
sudo rm /etc/rpi/swap.conf.d/99-disable-swap.conf
```

Reboot:

```bash
sudo reboot
```

Verify that swap is active again:

```bash
swapon --show
free -h
systemctl list-units --quiet --type=swap
```

## Notes

- Do not delete `/var/swap` until swap has been disabled, rebooted, and verified.
- On Raspberry Pi OS `trixie`, old `dphys-swapfile` instructions may not apply.

