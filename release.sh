#!/bin/bash

VERSION="v1.0.0"

gh release create $VERSION ./focus-symphony \
  --title "Focus Symphony $VERSION – CLI Binary Release" \
  --notes "Focus Symphony CLI

Features:
- Website blocking via /etc/hosts
- Focus music via mpv + yt-dlp
- Pomodoro mode

Run:
chmod +x focus-symphony
./focus-symphony

Requirements:
- Linux
- mpv
- yt-dlp
- sudo privileges"
