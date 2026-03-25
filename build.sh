#!/bin/bash
set -e

echo "[+] Building Focus Symphony..."
go build -o focus-symphony main.go

echo "[+] Done. Binary: ./focus-symphony"
