#!/bin/bash
set -e

echo "🎼 FOCUS-SYMPHONY INSTALLER"
echo "---------------------------"

# 1. Detect Package Manager
if [ -x "$(command -v pacman)" ]; then
    PM_BIN="pacman"
    INSTALL_CMD="sudo pacman -S --noconfirm"
elif [ -x "$(command -v apt-get)" ]; then
    PM_BIN="apt-get"
    INSTALL_CMD="sudo apt-get install -y"
elif [ -x "$(command -v dnf)" ]; then
    PM_BIN="dnf"
    INSTALL_CMD="sudo dnf install -y"
elif [ -x "$(command -v brew)" ]; then
    PM_BIN="brew"
    INSTALL_CMD="brew install"
elif [ -x "$(command -v zypper)" ]; then
    PM_BIN="zypper"
    INSTALL_CMD="sudo zypper install -y"
elif [ -x "$(command -v apk)" ]; then
    PM_BIN="apk"
    INSTALL_CMD="sudo apk add"
else
    echo "❌ No supported package manager found. Please install mpv and yt-dlp manually."
fi

# 2. Install Dependencies
if [ ! -z "$INSTALL_CMD" ]; then
    echo "[+] Installing dependencies (mpv, yt-dlp)..."
    $INSTALL_CMD mpv yt-dlp || echo "⚠️  Failed to install via $PM_BIN. Trying pip3 for yt-dlp..."
    
    if [ ! -x "$(command -v yt-dlp)" ] && [ -x "$(command -v pip3)" ]; then
        pip3 install yt-dlp
    fi
fi

# 3. Build Focus Symphony
echo "[+] Building Focus Symphony..."
go build -o focus-symphony main.go

# 4. Install Binary
echo "[+] Installing binary to ~/.local/bin/focus-symphony"
mkdir -p ~/.local/bin
mv focus-symphony ~/.local/bin/

# 5. Install Assets
echo "[+] Installing assets to ~/.local/share/focus-symphony/assets/"
mkdir -p ~/.local/share/focus-symphony/assets
cp assets/lofi.mp3 ~/.local/share/focus-symphony/assets/

echo "---------------------------"
echo "✅ Done! You can now run 'focus-symphony' from your terminal."
echo "   (Make sure ~/.local/bin is in your PATH)"
