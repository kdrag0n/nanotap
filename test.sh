#!/usr/bin/env bash
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" && goupx nanotap > /dev/null 2>&1 && adb push nanotap config.toml test.stage2.sh /data/local/tmp > /dev/null && adb shell chmod 755 /data/local/tmp/test.stage2.sh && adb shell su -c /data/local/tmp/test.stage2.sh

