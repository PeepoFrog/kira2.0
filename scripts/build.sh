#!/usr/bin/env bash
set -e
set -x

echo "INFO: Setting variables"

LOCAL_PLATFORM=$(uname | awk '{print tolower($0)}')
LOCAL_ARCH=$(uname -m | grep -q 'arm\|aarch' && echo "arm64" || echo "amd64")
LOCAL_OUT="${GOBIN}/kira2"

echo "DEBUG: LOCAL_PLATFORM: $LOCAL_PLATFORM"
echo "           LOCAL_ARCH: $LOCAL_ARCH"
echo "            LOCAL_OUT: $LOCAL_OUT"

PLATFORM="$1" 
ARCH="$2" 
OUTPUT="$3" 
[ -z "$PLATFORM" ] && PLATFORM="$LOCAL_PLATFORM"
[ -z "$ARCH" ] && ARCH="$LOCAL_ARCH"
[ -z "$OUTPUT" ] && OUTPUT="$LOCAL_OUT"

CONSTANTS_FILE=./internal/types/constants.go
VERSION=$(awk -F'Version = ' '/Version/{print $2}' $CONSTANTS_FILE | tr -d '[:space:]\"')

if [[ -z "$VERSION" ]]; then
    echo "ERROR: Version was NOT found in constants '$CONSTANTS_FILE' !" >&2
    sleep 5
    exit 1
fi

rm -fv "$OUTPUT" || echo "ERROR: Failed to wipe old kira2 binary"

#[LOCAL]
#LOCAL_OUT="/home/eugene/Code/github.com/MrLutik/kira2.0"
#OUTPUT="$LOCAL_OUT"

go mod tidy
GO111MODULE=on go mod verify
env GOOS="$PLATFORM" GOARCH="$ARCH" go build -o "${OUTPUT}_launcher" ./cmd/kira2_launcher/.
env GOOS="$PLATFORM" GOARCH="$ARCH" go build -o "$OUTPUT" ./cmd/kira2/.