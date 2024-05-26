#!/usr/bin/env bash

# go tool dist list

### ARM
GOARCH=arm64 GOOS=darwin CGO_ENABLED=1 go build -o build/noscrape_darwin_arm64
# GOARCH=arm64 GOOS=linux go build -o build/noscrape_linux_arm64

### AMD64
GOARCH=amd64 GOOS=linux CGO_ENABLED=1 CC=x86_64-unknown-linux-gnu-gcc go build -o build/noscrape_linux_amd64
GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o build/noscrape_windows_amd64.exe
GOARCH=amd64 GOOS=darwin CGO_ENABLED=1 go build -o build/noscrape_darwin_amd64


