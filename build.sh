#!/usr/bin/env bash

go build -ldflags="-s -w" -o faktor.bin cmd/all/main.go
upx --brute faktor.bin
