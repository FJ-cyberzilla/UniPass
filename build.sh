#!/usr/bin/env bash
set -euo pipefail

APP_NAME="unipass"
BUILD_DIR="dist"
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "v1.0.0")
BUILD_TIME=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
LDFLAGS="-s -w -X 'main.version=${VERSION}' -X 'main.buildTime=${BUILD_TIME}'"

echo "🚀 Building UniPass cross-platform binaries [${VERSION}]..."
rm -rf "${BUILD_DIR}"
mkdir -p "${BUILD_DIR}"

# Target OS / ARCH Matrix (Excludes Apple)
TARGETS=(
  "linux/amd64"
  "linux/arm64"
  "windows/amd64/.exe"
  "windows/386/.exe"
  "windows/arm64/.exe"
  "android/arm64"
)

for TARGET in "${TARGETS[@]}"; do
  IFS="/" read -r GOOS GOARCH EXT <<< "${TARGET}"
  EXT=${EXT:-""}
  OUTPUT_NAME="${APP_NAME}-${GOOS}-${GOARCH}${EXT}"
  
  echo "  🔨 Compiling for ${GOOS}/${GOARCH}..."
  CGO_ENABLED=0 GOOS="${GOOS}" GOARCH="${GOARCH}" \
    go build -trimpath -ldflags="${LDFLAGS}" -o "${BUILD_DIR}/${OUTPUT_NAME}" ./cmd/unipass
done

echo "✅ All cross-platform binaries successfully created in ./${BUILD_DIR}/"
ls -lh "${BUILD_DIR}"
