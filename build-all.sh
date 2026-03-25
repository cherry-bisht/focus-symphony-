#!/bin/bash

echo "[+] Building Linux"
GOOS=linux GOARCH=amd64 go build -o focus-symphony-linux main.go

echo "[+] Building Windows"
GOOS=windows GOARCH=amd64 go build -o focus-symphony.exe main.go

echo "[+] Building Mac"
GOOS=darwin GOARCH=amd64 go build -o focus-symphony-mac main.go

echo "[+] Done."
