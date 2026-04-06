# Focus-Symphony 🎻

**Focus-Symphony** is a powerful, open-source productivity CLI tool that helps developers achieve deep work by blocking distracting websites. Built with Go for speed and reliability, it works on Linux, macOS, and Windows.

```
  ____ ___ ____  _   _ _____ 
 | __ )_ _/ ___|| | | |_   _|
 |  _ \| |\___ \| |_| | | |  
 | |_) | | ___) |  _  | | |  
 |____/___|____/|_| |_| |_|  
```

---

## 🎯 Features

### 🛡️ Acoustic Shield (Website Blocker)
- Blocks distracting websites (YouTube, Reddit, Twitter/X, Facebook, Instagram) system-wide
- Works by modifying your system's `hosts` file for reliable blocking
- Automatically manages both IPv4 and IPv6 entries
- Clean start/stop with automatic cleanup

### 🎵 Terminal Music Player
- Stream focus-optimized playlists directly from your terminal
- Lofi, Dark Academia, Cyberpunk/Synthwave tracks
- Powered by `mpv` and `yt-dlp` for high-quality audio

### 📊 Session Tracking
- Track your focus session duration
- View statistics on your productivity

---

## 🚀 Quick Start

### Download Pre-built Binary (Recommended)

Go to [Releases](https://github.com/AbhishekMauryaGEEK/focus-symphony-/releases) and download the binary for your system:

- **Linux (x64)**: `focus-symphony-linux-amd64`
- **Linux (ARM64)**: `focus-symphony-linux-arm64`
- **Windows (x64)**: `focus-symphony-windows-amd64.exe`
- **macOS (Intel)**: `focus-symphony-macos-amd64`
- **macOS (Apple Silicon)**: `focus-symphony-macos-arm64`

**Linux/macOS:**
```bash
# Make executable
chmod +x focus-symphony-*

# Run (requires sudo/admin)
sudo ./focus-symphony-*
```

**Windows:**
1. Right-click `focus-symphony-windows-amd64.exe`
2. Select **"Run as Administrator"**

---

## 📦 Installation from Source

### Prerequisites

You'll need:
- **Go 1.19+** (for building)
- **mpv** (for music playback)
- **yt-dlp** (for streaming music from YouTube)

### Linux Installation

#### Ubuntu / Debian / Pop!_OS / Linux Mint
```bash
# Install dependencies
sudo apt update
sudo apt install -y golang mpv

# Install yt-dlp (via pip for latest version)
sudo apt install -y python3-pip
pip3 install yt-dlp

# Clone and build
git clone https://github.com/AbhishekMauryaGEEK/focus-symphony-.git
cd focus-symphony-
go build -o focus-symphony main.go

# Optional: Install to system
sudo mv focus-symphony /usr/local/bin/
```

#### Fedora / RHEL / CentOS
```bash
# Install dependencies
sudo dnf install -y golang mpv python3-pip
pip3 install yt-dlp

# Clone and build
git clone https://github.com/AbhishekMauryaGEEK/focus-symphony-.git
cd focus-symphony-
go build -o focus-symphony main.go

# Optional: Install to system
sudo mv focus-symphony /usr/local/bin/
```

#### Arch Linux / Manjaro
```bash
# Install dependencies
sudo pacman -S go mpv yt-dlp

# Clone and build
git clone https://github.com/AbhishekMauryaGEEK/focus-symphony-.git
cd focus-symphony-
go build -o focus-symphony main.go

# Optional: Install to system
sudo mv focus-symphony /usr/local/bin/
```

### macOS Installation

```bash
# Install dependencies with Homebrew
brew install go mpv yt-dlp

# Clone and build
git clone https://github.com/AbhishekMauryaGEEK/focus-symphony-.git
cd focus-symphony-
go build -o focus-symphony main.go

# Optional: Install to system
sudo mv focus-symphony /usr/local/bin/
```

### Windows Installation

1. Install Go from [golang.org](https://go.dev/dl/)
2. Clone the repository:
   ```powershell
   git clone https://github.com/AbhishekMauryaGEEK/focus-symphony-.git
   cd focus-symphony-
   go build -o focus-symphony.exe main.go
   ```
3. Right-click `focus-symphony.exe` → **Run as Administrator**

> **Note:** Music features require `mpv` and `yt-dlp`. On Windows, install via [Scoop](https://scoop.sh/) or [Chocolatey](https://chocolatey.org/).

---

## 🎮 Usage

Run the program with elevated privileges:

```bash
# Linux/macOS
sudo ./focus-symphony

# Windows (Right-click → Run as Administrator)
focus-symphony.exe
```

### Commands

Once running, use these commands:

| Command | Description | Requires Admin |
|---------|-------------|----------------|
| `start` | Block distracting sites | ✅ Yes |
| `stop` | Unblock sites | ✅ Yes |
| `music` | Play a random focus track | ❌ No |
| `playlist` | Show all available tracks | ❌ No |
| `1-5` | Play specific track by number | ❌ No |
| `stop_music` | Stop music playback | ❌ No |
| `rapid` | 25-min Pomodoro (blocking + music) | ✅ Yes |
| `stats` | Show session statistics | ❌ No |
| `help` | Display help menu | ❌ No |
| `exit` | Quit the program | ❌ No |

### Example Session

```
fs > start
⚡ Activating Acoustic Shield...

╔═══════════════════════════════════════════════════════════════════╗
║  ⚠️  DNS OVER HTTPS (DoH) WARNING                                 ║
╚═══════════════════════════════════════════════════════════════════╝

[... DNS warning details ...]

✅ 19 sites blocked. Deep work mode active.

   ⚠️  IMPORTANT: Restart your browser for blocking to take effect!
   💡 If sites aren't blocked, disable 'Secure DNS' in browser settings.

fs > music
🎵 Streaming: Lofi Hip Hop - Beats to Relax/Study
   ✅ Audio engine running. Volume: 80%

fs > stats
📊 FOCUS TIME : 25m30s
🎵 NOW PLAYING: Lofi Hip Hop - Beats to Relax/Study

fs > stop
⚡ Deactivating Acoustic Shield...
✅ Shield deactivated. Focus session lasted: 25m30s

fs > exit
Exiting. Keep Focus!
```

---

## 🏗️ How It Works (Architecture)

Focus-Symphony uses a simple but effective architecture:

```
┌─────────────────┐
│  Focus-Symphony │
│   CLI Program   │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  /etc/hosts     │  (Linux/macOS)
│  C:\Windows\    │  (Windows)
│  System32\...\  │
│  hosts          │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  DNS Resolution │
│  127.0.0.1      │  (Localhost)
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Sites Blocked  │
│  ✅ No Access   │
└─────────────────┘
```

**How blocking works:**

1. **Modify Hosts File**: When you run `start`, the program adds entries to your system's hosts file mapping blocked domains (e.g., `youtube.com`) to `127.0.0.1` (localhost)

2. **DNS Override**: When your browser tries to visit `youtube.com`, your system checks the hosts file first, finds the `127.0.0.1` entry, and never makes a real DNS query

3. **Site Unreachable**: The browser tries to connect to `127.0.0.1` instead of the real site, which fails → site blocked

4. **Clean Removal**: When you run `stop`, the program removes these entries, restoring normal internet access

**Why this is effective:**
- Works at the OS level (blocks across all browsers and apps)
- No browser extensions needed
- Fast and lightweight
- Deterministic behavior

**Important limitation:**
- **DNS over HTTPS (DoH)** bypasses the hosts file by using encrypted DNS directly to Google/Cloudflare servers. You **must** disable DoH in your browser for blocking to work. See troubleshooting below.

---

## ⚠️ Critical: DNS over HTTPS (DoH) Must Be Disabled

**Focus-Symphony blocks sites by modifying your hosts file, but modern browsers use "Secure DNS" (DNS over HTTPS) which bypasses the hosts file.**

### How to Disable DoH:

#### Chrome / Edge / Brave
1. Open **Settings** → **Privacy and security** → **Security**
2. Scroll to **"Use secure DNS"**
3. **Turn it OFF** or select "With your current service provider"
4. **Restart your browser**

#### Firefox
1. Open **Settings** → **Privacy & Security**
2. Scroll to **"DNS over HTTPS"**
3. Select **"Off"** or **"Default"**
4. **Restart your browser**

#### Safari (macOS)
1. Safari typically uses system DNS (no DoH by default)
2. If issues persist, check **System Preferences** → **Profiles** for custom DNS configs

**Without disabling DoH, the blocking will NOT work!**

---

## 🐛 Troubleshooting

### Sites Not Blocking

**Problem:** YouTube/Reddit still accessible after running `start`

**Solutions:**
1. ✅ **Disable Secure DNS (DoH)** in your browser (see above)
2. ✅ **Restart your browser completely** (close all windows, reopen)
3. ✅ **Clear browser DNS cache**:
   - Chrome: Visit `chrome://net-internals/#dns` → Click "Clear host cache"
   - Firefox: Restart browser
4. ✅ **Verify hosts file was modified**:
   - Linux/Mac: `cat /etc/hosts | grep FOCUS-SYMPHONY`
   - Windows: `type C:\Windows\System32\drivers\etc\hosts | findstr FOCUS-SYMPHONY`
5. ✅ **Run with admin privileges** (sudo on Linux/Mac, "Run as Administrator" on Windows)

### Music Not Playing

**Problem:** Music command fails or no audio

**Solutions:**
1. ✅ **Check dependencies**:
   ```bash
   mpv --version
   yt-dlp --version
   ```
   Install if missing (see Installation section)

2. ✅ **Update yt-dlp** (often fixes stream resolution issues):
   ```bash
   pip3 install -U yt-dlp
   ```

3. ✅ **Test mpv directly**:
   ```bash
   mpv --no-video https://www.youtube.com/watch?v=jfKfPfyJRdk
   ```

4. ✅ **Check audio output**:
   - Ensure speakers/headphones are connected
   - Check system volume isn't muted

### Permission Errors

**Problem:** "Permission denied" or "Access denied"

**Solutions:**
1. ✅ **Linux/Mac**: Run with `sudo`
   ```bash
   sudo ./focus-symphony
   ```

2. ✅ **Windows**: Right-click executable → "Run as Administrator"

3. ✅ **Check file ownership** (Linux/Mac):
   ```bash
   ls -l /etc/hosts
   # Should be owned by root
   ```

### Start/Stop Not Working

**Problem:** Running `start` or `stop` doesn't change blocking status

**Solutions:**
1. ✅ **Restart with admin/sudo** (the tool will auto-prompt if needed)
2. ✅ **Check for leftover entries**:
   - Manually inspect hosts file
   - Remove any broken `# FOCUS-SYMPHONY-BLOCK` markers
3. ✅ **Verify hosts file isn't locked** by another program
4. ✅ **On Windows**: Disable antivirus temporarily (some block hosts file modification)

### Browser Still Shows Cached Content

**Problem:** Blocked site loads from browser cache

**Solutions:**
1. ✅ **Hard refresh** the page: `Ctrl+Shift+R` (or `Cmd+Shift+R` on Mac)
2. ✅ **Clear browser cache**:
   - Chrome: Settings → Privacy → Clear browsing data
   - Firefox: Settings → Privacy → Clear Data
3. ✅ **Use incognito/private mode** to test

### Inconsistent Blocking (Works Sometimes)

**Problem:** YouTube blocks but Reddit doesn't

**Solutions:**
1. ✅ **v1.7.0 fixes this!** Update to latest version
2. ✅ **Check for mobile subdomains**:
   - Old versions didn't block `m.youtube.com`, `old.reddit.com`, etc.
   - New version blocks all variants
3. ✅ **Verify both IPv4 and IPv6 entries** exist:
   ```bash
   grep "youtube.com" /etc/hosts
   # Should show both 127.0.0.1 and ::1 entries
   ```

---

## 🤝 AI Usage Disclosure

This project was developed with assistance from AI tools:

- **Code**: GitHub Copilot and Claude AI assisted with Go development, error handling, and cross-platform compatibility
- **README**: AI helped structure and refine documentation for clarity
- **Demo Site**: Landing page HTML/CSS was partially generated with AI assistance

All AI-generated code has been reviewed, tested, and customized for this project's specific needs. The core blocking logic, architecture decisions, and feature design are human-directed.

---

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

---

## 🙏 Credits

- **BISHT** ([@cherry-bisht](https://github.com/cherry-bisht)) - Original creator
- **Flavourtown** - Collaboration
- **AbhishekMauryaGEEK** - Current maintainer

Built with passion for the developer community. **Harmonizing productivity for deep work.**

---

## 🔗 Links

- [Report Issues](https://github.com/AbhishekMauryaGEEK/focus-symphony-/issues)
- [Contribute](https://github.com/AbhishekMauryaGEEK/focus-symphony-/pulls)
- [Releases](https://github.com/AbhishekMauryaGEEK/focus-symphony-/releases)

---

**⭐ If this helps you stay focused, give it a star!**

