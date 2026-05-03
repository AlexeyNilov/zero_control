#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUTPUT_DIR="$ROOT_DIR/bin"
OUTPUT_FILE="$OUTPUT_DIR/zero_control.exe"

mkdir -p "$OUTPUT_DIR"

go build -o "$OUTPUT_FILE" ./cmd/zero_control

printf 'Built %s\n' "$OUTPUT_FILE"
