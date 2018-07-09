#!/usr/bin/env bash
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" && upx -q nanotap && adb push nanotap /data/local/tmp && adb shell chmod 755 /data/loca
l/tmp && adb shell su -c /data/local/tmp/nanotap

