# Raspberry Pi Zero 2 WH OS install

This note records the chosen base OS setup for `zero_control`.

## Target

- Device: `Raspberry Pi Zero 2 WH`
- OS: `Raspberry Pi OS Lite (64-bit)`
- Access method: `SSH` with public-key authentication only

## What you need

- Raspberry Pi Zero 2 WH
- microSD card, preferably `16 GB` or larger
- microSD card reader for your computer
- stable `5V` power supply
- Wi-Fi SSID, password, and country
- an SSH key pair on your development machine, or willingness to generate one during imaging

## Install the OS

Use `Raspberry Pi Imager` on another computer.

1. Insert the microSD card into your computer.
2. Open `Raspberry Pi Imager`.
3. Choose `Device` -> `Raspberry Pi Zero 2 W`.
4. Choose `OS` -> `Raspberry Pi OS Lite (64-bit)`.
5. Choose `Storage` -> your microSD card.
6. Click `Next`.
7. When prompted, click `Edit Settings`.

## Imager settings

Set the following before writing the image:

- `Hostname`: for example `zero-control`
- `Username`: choose the admin user name you want on the Pi
- `Wi-Fi`: enter SSID, password, and correct country
- `Time zone` and `keyboard layout`: set your locale
- `SSH`: enable it
- `Authentication`: choose `Allow public-key authentication only`

If you already have an SSH key on your machine, let Imager use the existing public key. If not, use Imager's key generation option or create one manually before imaging.

## Optional: generate an SSH key manually

### Windows PowerShell

```powershell
ssh-keygen -t ed25519 -C "zero-control"
```

### Linux or macOS

```bash
ssh-keygen -t ed25519 -C "zero-control"
```

Accept the default file location unless you have a reason to store the key elsewhere.

## Write and boot

1. Save the settings in Imager.
2. Confirm the write.
3. Wait for the image write and verification to complete.
4. Remove the card from your computer.
5. Insert the card into the Pi.
6. Power on the Pi.

First boot may take a few minutes.

## First connection

Connect from another machine on the same network:

```bash
ssh youruser@zero-control.local
```

Replace `youruser` with the username configured in Imager.

If `.local` name resolution does not work on your network, use the device IP address assigned by your router:

```bash
ssh youruser@192.168.x.x
```

## Initial update

After logging in:

```bash
sudo apt update
sudo apt full-upgrade -y
sudo reboot
```

## Why this setup

- `Raspberry Pi OS Lite` keeps the base system minimal and avoids unnecessary desktop packages.
- The `64-bit` build matches the `ARMv8` CPU in the Pi Zero 2 WH.
- SSH key authentication is safer and more appropriate than password-based remote login for a headless bot host.

