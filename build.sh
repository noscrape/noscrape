#!/usr/bin/env bash

# go tool dist list

GOARCH=arm64 GOOS=darwin CGO_ENABLED=1 go build -buildmode=c-shared -o build/noscrape_darwin_arm64

GOARCH=amd64 GOOS=linux CGO_ENABLED=1 CC=x86_64-unknown-linux-gnu-gcc go build -buildmode=c-shared -o build/noscrape_linux_amd64
GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -buildmode=c-shared -o build/noscrape_windows_amd64.exe
GOARCH=amd64 GOOS=darwin CGO_ENABLED=1 go build -buildmode=c-shared -o build/noscrape_darwin_amd64

# math.MaxUint32 (untyped int constant 4294967295) overflows int
# GOARCH=arm GOOS=linux CGO_ENABLED=1 CC=x86_64-unknown-linux-gnu-gcc go build -C noscrape -o ../build/noscrape_linux_amd64
# GOARCH=arm GOOS=linux CGO_ENABLED=1 go build -C noscrape -o ../build/noscrape_linux_arm64
# GOARCH=386 GOOS=linux CGO_ENABLED=1 go build -C noscrape -o ../build/noscrape_linux_386

