#!/usr/bin/env bash
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" && adb push nanotap /data/local/tmp && adb shell chmod 755 /data/local/tmp && adb shell su -c /data/local/tmp/nanotap

