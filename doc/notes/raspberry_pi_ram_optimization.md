# Raspberry Pi RAM and power optimization

This note records conservative RAM and power optimization options for the Zero.

## Context

- Device: `Raspberry Pi Zero 2 WH`
- OS: `Raspberry Pi OS Lite (64-bit)`, based on Debian 13 (`trixie`)
- RAM: about `416 MiB`
- Primary workload: Telegram bot
- Network: Wi-Fi

The goal is to reduce unnecessary services and hardware initialization without breaking remote access, Wi-Fi, or Telegram HTTPS connectivity.

## Good candidates

### Bluetooth

Bluetooth is not needed for the Telegram bot. See:

```text
doc/notes/raspberry_pi_bluetooth_disable.md
```

### Audio

If no audio output is needed, disable onboard audio in `/boot/firmware/config.txt`.

Comment out:

```ini
dtparam=audio=on
```

Or add:

```ini
dtparam=audio=off
```

Reboot after changing boot config:

```bash
sudo reboot
```

Rollback: remove `dtparam=audio=off` or restore `dtparam=audio=on`.

### Camera and display auto-detect

If no CSI camera or DSI display is attached, disable auto-detection in `/boot/firmware/config.txt`:

```ini
camera_auto_detect=0
display_auto_detect=0
```

Rollback:

```ini
camera_auto_detect=1
display_auto_detect=1
```

Reboot after changing boot config.

### Activity LED

This saves very little power, but is low risk if status LEDs are not useful.

Add to `/boot/firmware/config.txt`:

```ini
dtparam=act_led_trigger=none
dtparam=act_led_activelow=off
```

Rollback: remove those lines.

Power LED behavior varies by Raspberry Pi model, so verify support before relying on a power LED setting.

### Serial login console

If UART login is not used, disable the serial getty:

```bash
sudo systemctl disable --now serial-getty@ttyAMA0.service
```

Rollback:

```bash
sudo systemctl enable --now serial-getty@ttyAMA0.service
```

### Avahi / mDNS

Only disable this if `.local` hostname discovery is not needed.

Disable:

```bash
sudo systemctl disable --now avahi-daemon.service
```

Rollback:

```bash
sudo systemctl enable --now avahi-daemon.service
```

## Swap

Disabling swap may make failures more abrupt on a device with about `416 MiB` RAM. `zram` swap consumes little memory until it is used.

If swap should still be disabled, use the dedicated note:

```text
doc/notes/raspberry_pi_swap_disable.md
```

## Measure before disabling more

Use these commands to find services and processes that actually consume memory:

```bash
systemctl --type=service --state=running --no-pager
systemctl list-unit-files --state=enabled --type=service --no-pager
ps -eo pid,comm,rss --sort=-rss | head -30
free -h
```

The process RSS list is usually the most useful signal. Services that are enabled but idle and small are not worth optimizing.

## Avoid disabling

Do not disable these for the current use case:

- Wi-Fi, NetworkManager, or `wpa_supplicant`: required for network access.
- SSH, unless another access path exists.
- Time sync: Telegram HTTPS depends on correct time, and the Pi has no battery-backed RTC.
- Journald/logging: needed to debug bot failures.
- DNS resolver or other network services unless their role is understood.

## Reference

Raspberry Pi boot configuration options are documented in the official `config.txt` documentation:

```text
https://www.raspberrypi.com/documentation/computers/config_txt.html
```

