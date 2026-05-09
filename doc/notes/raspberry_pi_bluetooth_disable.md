# Raspberry Pi Bluetooth disable and rollback

This note records how Bluetooth was disabled on the Zero and how to reverse it.

## Disable Bluetooth

Stop and disable the Bluetooth daemon:

```bash
sudo systemctl disable --now bluetooth.service
```

If the system has `hciuart.service`, stop and disable it too:

```bash
sudo systemctl disable --now hciuart.service
```

If that command reports `Unit hciuart.service does not exist`, no action is needed for that service.

To prevent onboard Bluetooth from initializing at boot, add this line to the boot config:

```ini
dtoverlay=disable-bt
```

Use the config file that exists on the installed OS:

```bash
ls /boot/firmware/config.txt /boot/config.txt
```

Newer Raspberry Pi OS usually uses `/boot/firmware/config.txt`; older installs may use `/boot/config.txt`.

Reboot after changing the boot config:

```bash
sudo reboot
```

## Restore Bluetooth

Remove or comment out this line from the boot config:

```ini
dtoverlay=disable-bt
```

Re-enable the Bluetooth daemon:

```bash
sudo systemctl enable --now bluetooth.service
```

If `hciuart.service` exists on the system, re-enable it too:

```bash
sudo systemctl enable --now hciuart.service
```

Reboot after restoring the boot config:

```bash
sudo reboot
```

## Check current state

```bash
systemctl status bluetooth.service
systemctl status hciuart.service
rfkill list
```

