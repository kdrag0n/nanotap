#!/usr/bin/env bash
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" && upx -q nanotap > /dev/null 2>&1 && adb push nanotap /data/local/tmp && adb shell chmod 755 /data/local/tmp && adb shell su -c /data/local/tmp/nanotap

