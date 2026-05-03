#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUTPUT_DIR="$ROOT_DIR/bin"
OUTPUT_FILE="$OUTPUT_DIR/zero_control-linux-armv6"

mkdir -p "$OUTPUT_DIR"

export GOOS=linux
export GOARCH=arm
export GOARM=6
export CGO_ENABLED=0

go build -o "$OUTPUT_FILE" ./cmd/zero_control
ls -lh "$OUTPUT_FILE"

printf 'Built %s\n' "$OUTPUT_FILE"
