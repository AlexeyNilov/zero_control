#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
LOCAL_FILE="${LOCAL_FILE:-$ROOT_DIR/bin/zero_control-linux-arm64}"
PI_HOST="${1:-${PI_HOST:-}}"
PI_USER="${2:-${PI_USER:-pi}}"
PI_PATH="${3:-${PI_PATH:-/home/pi/zero_control/zero_control}}"

if [[ -z "$PI_HOST" ]]; then
  printf 'Usage: %s <pi-host> [pi-user] [remote-path]\n' "$(basename "$0")" >&2
  printf 'Example: %s raspberrypi.local pi /home/pi/zero_control/zero_control\n' "$(basename "$0")" >&2
  exit 1
fi

if [[ ! -f "$LOCAL_FILE" ]]; then
  printf 'Build artifact not found: %s\n' "$LOCAL_FILE" >&2
  printf 'Run scripts/build_arm.sh first or set LOCAL_FILE.\n' >&2
  exit 1
fi

scp "$LOCAL_FILE" "${PI_USER}@${PI_HOST}:${PI_PATH}"

printf 'Copied %s to %s@%s:%s\n' "$LOCAL_FILE" "$PI_USER" "$PI_HOST" "$PI_PATH"
