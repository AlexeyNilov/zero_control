# Raspberry Pi Zero 2 WH device profile

This note records the hardware and operating system used for the `zero_control` device.

## Confirmed details

- Device: `Raspberry Pi Zero 2 WH`
- Device-tree model: `Raspberry Pi Zero 2 W Rev 1.0`
- Operating system: `Raspberry Pi OS Lite (64-bit)`
- OS base: `Debian GNU/Linux 13 (trixie)`
- Debian version: `13.4`
- Kernel: `6.12.75+rpt-rpi-v8`
- Machine architecture: `aarch64`
- Debian package architecture: `arm64`
- CPU: `ARM Cortex-A53`
- CPU cores: `4`
- CPU frequency range: `600 MHz` to `1000 MHz`
- Memory: `416 MiB`
- Swap: `415 MiB`
- Boot partition: `mmcblk0p1`, `vfat`, mounted at `/boot/firmware`
- Root partition: `mmcblk0p2`, `ext4`, mounted at `/`
- Boot filesystem: `505M` total, `432M` available, `15%` used
- Root filesystem: `58G` total, `54G` available, `5%` used

## MQTT broker

- Broker: `Mosquitto`
- Service: `mosquitto`
- Status: installed and running
- Listener: `127.0.0.1:1883`
- Scope: local-only MQTT broker for integrating the bot and other services on the device

The broker currently listens on loopback only, so clients must run on the Raspberry Pi itself or connect through a local tunnel/proxy. Exposing MQTT to the LAN or other networks should be a separate configuration decision with authentication enabled.

### OS release

Captured from `/etc/os-release`:

```text
PRETTY_NAME="Debian GNU/Linux 13 (trixie)"
NAME="Debian GNU/Linux"
VERSION_ID="13"
VERSION="13 (trixie)"
VERSION_CODENAME=trixie
DEBIAN_VERSION_FULL=13.4
ID=debian
```

### Kernel

Captured from `uname -a`:

```text
Linux zero 6.12.75+rpt-rpi-v8 #1 SMP PREEMPT Debian 1:6.12.75-1+rpt1 (2026-03-11) aarch64 GNU/Linux
```

### Package architecture

Captured from `dpkg --print-architecture`:

```text
arm64
```

### Hardware model

Captured from `/proc/device-tree/model`:

```text
Raspberry Pi Zero 2 W Rev 1.0
```

### CPU details

Captured from `lscpu`:

```text
Architecture:                aarch64
  CPU op-mode(s):            32-bit, 64-bit
  Byte Order:                Little Endian
CPU(s):                      4
  On-line CPU(s) list:       0-3
Vendor ID:                   ARM
  Model name:                Cortex-A53
    Model:                   4
    Thread(s) per core:      1
    Core(s) per cluster:     4
    Socket(s):               -
    Cluster(s):              1
    Stepping:                r0p4
    CPU(s) scaling MHz:      60%
    CPU max MHz:             1000.0000
    CPU min MHz:             600.0000
    BogoMIPS:                38.40
    Flags:                   fp asimd evtstrm crc32 cpuid
Caches (sum of all):
  L1d:                       128 KiB (4 instances)
  L1i:                       128 KiB (4 instances)
  L2:                        512 KiB (1 instance)
NUMA:
  NUMA node(s):              1
  NUMA node0 CPU(s):         0-3
Vulnerabilities:
  Gather data sampling:      Not affected
  Indirect target selection: Not affected
  Itlb multihit:             Not affected
  L1tf:                      Not affected
  Mds:                       Not affected
  Meltdown:                  Not affected
  Mmio stale data:           Not affected
  Reg file data sampling:    Not affected
  Retbleed:                  Not affected
  Spec rstack overflow:      Not affected
  Spec store bypass:         Not affected
  Spectre v1:                Mitigation; __user pointer sanitization
  Spectre v2:                Not affected
  Srbds:                     Not affected
  Tsa:                       Not affected
  Tsx async abort:           Not affected
  Vmscape:                   Not affected
```

### Memory

Captured from `free -h`:

```text
               total        used        free      shared  buff/cache   available
Mem:           416Mi       143Mi       184Mi       2.2Mi       142Mi       272Mi
Swap:          415Mi          0B       415Mi
```

### Block devices

Captured from `lsblk -f`:

```text
NAME        FSTYPE FSVER LABEL  UUID                                 FSAVAIL FSUSE% MOUNTPOINTS
loop0       swap   1
mmcblk0
|-mmcblk0p1 vfat   FAT32 bootfs 0D58-6978                             431.1M    14% /boot/firmware
`-mmcblk0p2 ext4   1.0   rootfs e634e0a4-a958-46cb-abad-862d2102573f   53.1G     4% /
zram0       swap   1     zram0  06a890a0-bc5e-4fb6-a3be-2fa1ac9feab3                [SWAP]
```

### Filesystem usage

Captured from `df -h`:

```text
Filesystem      Size  Used Avail Use% Mounted on
udev             72M     0   72M   0% /dev
tmpfs            84M  2.2M   82M   3% /run
/dev/mmcblk0p2   58G  2.5G   54G   5% /
tmpfs           209M     0  209M   0% /dev/shm
tmpfs           5.0M  8.0K  5.0M   1% /run/lock
tmpfs           1.0M     0  1.0M   0% /run/credentials/systemd-journald.service
tmpfs           209M     0  209M   0% /tmp
/dev/mmcblk0p1  505M   74M  432M  15% /boot/firmware
tmpfs           1.0M     0  1.0M   0% /run/credentials/getty@tty1.service
tmpfs           1.0M     0  1.0M   0% /run/credentials/serial-getty@ttyAMA0.service
tmpfs            42M  8.0K   42M   1% /run/user/1000
```

## Notes

- The `WH` variant includes pre-soldered GPIO headers.
- Keep credentials, Wi-Fi passwords, tokens, and private keys out of this file.
