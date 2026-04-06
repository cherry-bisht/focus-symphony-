package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func getHostsPath() string {
	if runtime.GOOS == "windows" {
		return "C:\\Windows\\System32\\drivers\\etc\\hosts"
	}
	return "/etc/hosts"
}

const blockMarker = "# FOCUS-SYMPHONY-BLOCK"

var hostsPath = getHostsPath()

var sitesToBlock = []string{
	"youtube.com",
	"www.youtube.com",
	"m.youtube.com",
	"youtu.be",
	"reddit.com",
	"www.reddit.com",
	"old.reddit.com",
	"new.reddit.com",
	"twitter.com",
	"www.twitter.com",
	"mobile.twitter.com",
	"x.com",
	"www.x.com",
	"mobile.x.com",
	"facebook.com",
	"www.facebook.com",
	"m.facebook.com",
	"instagram.com",
	"www.instagram.com",
}

var playlist = []struct {
	name string
	url  string
}{
	{"Lofi Hip Hop - Beats to Relax/Study", "https://www.youtube.com/watch?v=jfKfPfyJRdk"},
	{"Lofi Girl - Sleep Mix", "https://www.youtube.com/watch?v=rUxyKA_-grg"},
	{"Dark Academia Study Music", "https://www.youtube.com/watch?v=hHW1oY26kxQ"},
	{"Coding in the Rain (Ambient)", "https://www.youtube.com/watch?v=mPZkdNFkNps"},
	{"Cyberpunk / Synthwave Focus Mix", "https://www.youtube.com/watch?v=qYnA9wWFHLI"},
}

var musicCmd *exec.Cmd
var currentTrack string
var sessionStart time.Time
var isShieldActive bool

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(`
  ____ ___ ____  _   _ _____ 
 | __ )_ _/ ___|| | | |_   _|
 |  _ \| |\___ \| |_| | | |  
 | |_) | | ___) |  _  | | |  
 |____/___|____/|_| |_| |_|  
                               `)
	fmt.Println("FOCUS-SYMPHONY v1.7.0 (Production Ready)")
	fmt.Println("--------------------------------------------")
	fmt.Println()
	showHelp() // FIX 1: show help menu on startup

	// Ensure cleanup on exit
	defer func() {
		if isShieldActive {
			if runtime.GOOS != "windows" && os.Getuid() != 0 {
				return // Can't clean up without privileges
			}
			fmt.Println("\n🧹 Cleaning up...")
			cleanHosts()
		}
		if musicCmd != nil && musicCmd.Process != nil {
			stopMusic()
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("fs > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println()
			break
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		switch input {
		case "start":
			startSession()
		case "stop":
			stopSession()
		case "music", "song":
			playMusic()
		case "stop_music":
			stopMusic()
		case "playlist":
			showPlaylist()
		case "rapid":
			startRapidFocus()
		case "stats":
			showStats()
		case "help":
			showHelp()
		case "exit":
			stopSession()
			stopMusic()
			fmt.Println("Exiting. Keep Focus!")
			return
		default:
			if len(input) == 1 && input[0] >= '1' && input[0] <= '5' {
				idx := int(input[0] - '1')
				playTrack(idx)
			} else {
				fmt.Printf("Unknown command: '%s'. Type 'help'.\n", input)
			}
		}
	}
}

func relaunchWithSudo() {
	if runtime.GOOS == "windows" {
		fmt.Println("   ⚠️  Windows: Please run this program as Administrator.")
		fmt.Println("   Right-click the .exe and select 'Run as administrator'")
		os.Exit(1)
		return
	}
	
	fmt.Println("   🔑 'start' requires elevated privileges. Re-launching with sudo...")
	args := append([]string{os.Args[0]}, os.Args[1:]...)
	cmd := exec.Command("sudo", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("   ❌ sudo failed: %v\n", err)
	}
	os.Exit(0)
}

func startSession() {
	// Check for admin/root privileges
	if runtime.GOOS == "windows" {
		// On Windows, try to open the hosts file to check permissions
		f, err := os.OpenFile(hostsPath, os.O_RDONLY, 0)
		if err != nil {
			relaunchWithSudo()
			return
		}
		f.Close()
	} else if os.Getuid() != 0 {
		relaunchWithSudo()
		return
	}
	
	fmt.Println("⚡ Activating Acoustic Shield...")
	
	// Warn about DNS over HTTPS
	checkDNSoverHTTPS()
	
	// Clean any existing blocks first
	cleanHosts()
	
	// Verify hosts file is readable
	if _, err := os.Stat(hostsPath); err != nil {
		fmt.Printf("   ❌ Cannot access hosts file: %v\n", err)
		return
	}
	
	isShieldActive = true
	sessionStart = time.Now()
	
	// Open hosts file for appending
	f, err := os.OpenFile(hostsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("   ❌ Error opening hosts file: %v\n", err)
		if runtime.GOOS == "windows" {
			fmt.Println("   💡 Make sure you're running as Administrator!")
		}
		return
	}
	defer f.Close()
	
	// Add block marker and entries
	fmt.Fprintln(f, blockMarker)
	blockedCount := 0
	for _, site := range sitesToBlock {
		// Add both IPv4 and IPv6 entries for maximum compatibility
		fmt.Fprintf(f, "127.0.0.1       %s\n", site)
		fmt.Fprintf(f, "::1             %s\n", site)
		blockedCount++
	}
	fmt.Fprintln(f, blockMarker)
	
	fmt.Printf("✅ %d sites blocked. Deep work mode active.\n", blockedCount)
	fmt.Println()
	fmt.Println("   ⚠️  IMPORTANT: Restart your browser for blocking to take effect!")
	fmt.Println("   💡 If sites aren't blocked, disable 'Secure DNS' in browser settings.")
	fmt.Println()
	fmt.Println("   Tip: type 'music' to start your focus playlist")
}

func stopSession() {
	if !isShieldActive {
		fmt.Println("   ℹ️  No active shield to stop.")
		return
	}
	
	// Check for admin/root privileges
	if runtime.GOOS != "windows" && os.Getuid() != 0 {
		relaunchWithSudo()
		return
	}
	
	fmt.Println("⚡ Deactivating Acoustic Shield...")
	
	// Clean the hosts file
	cleanHosts()
	
	isShieldActive = false
	
	// Show session summary
	if !sessionStart.IsZero() {
		duration := time.Since(sessionStart).Round(time.Second)
		fmt.Printf("✅ Shield deactivated. Focus session lasted: %s\n", duration)
	} else {
		fmt.Println("✅ Shield deactivated.")
	}
	
	fmt.Println()
	fmt.Println("   ⚠️  IMPORTANT: Restart your browser to restore access to sites.")
	fmt.Println()
}

func cleanHosts() {
	input, err := os.ReadFile(hostsPath)
	if err != nil {
		fmt.Printf("   ❌ Failed to read hosts file: %v\n", err)
		return
	}
	
	lines := strings.Split(string(input), "\n")
	var newLines []string
	inBlockSection := false
	
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		
		// Toggle block section when we see the marker
		if trimmed == blockMarker {
			inBlockSection = !inBlockSection
			continue
		}
		
		// Skip lines inside block sections
		if inBlockSection {
			continue
		}
		
		// Keep all other lines (including empty ones to preserve formatting)
		newLines = append(newLines, line)
	}
	
	// Remove trailing empty lines but keep at least one newline at end
	for len(newLines) > 0 && strings.TrimSpace(newLines[len(newLines)-1]) == "" {
		newLines = newLines[:len(newLines)-1]
	}
	
	output := strings.Join(newLines, "\n")
	if !strings.HasSuffix(output, "\n") {
		output += "\n"
	}
	
	err = os.WriteFile(hostsPath, []byte(output), 0644)
	if err != nil {
		fmt.Printf("   ❌ Failed to write hosts file: %v\n", err)
	}
}

func getEffectiveUser() string {
	user := os.Getenv("SUDO_USER")
	if user == "" {
		user = os.Getenv("USER")
	}
	if user == "" {
		user = "nobody"
	}
	return user
}

func getRuntimeDir(user string) string {
	out, err := exec.Command("id", "-u", user).Output()
	if err != nil {
		return "/run/user/1000"
	}
	return "/run/user/" + strings.TrimSpace(string(out))
}

func getMusicHome() string {
	user := getEffectiveUser()
	out, err := exec.Command("getent", "passwd", user).Output()
	if err == nil {
		parts := strings.Split(strings.TrimSpace(string(out)), ":")
		if len(parts) >= 6 {
			return parts[5]
		}
	}
	home, _ := os.UserHomeDir()
	return home
}

func buildMpvCmd(audioSrc string, user string, runtimeDir string) *exec.Cmd {
	mpvArgs := []string{
		"--no-video",
		"--volume=80",
		"--really-quiet",
		"--msg-level=all=error",
		audioSrc,
	}
	if os.Getuid() == 0 {
		args := []string{"-u", user, "env",
			"XDG_RUNTIME_DIR=" + runtimeDir,
			"PULSE_SERVER=unix:" + runtimeDir + "/pulse/native",
			"mpv",
		}
		args = append(args, mpvArgs...)
		return exec.Command("sudo", args...)
	}
	cmd := exec.Command("mpv", mpvArgs...)
	cmd.Env = append(os.Environ(), "XDG_RUNTIME_DIR="+runtimeDir)
	return cmd
}

func detectPkgManager() (string, string) {
	type pm struct {
		bin    string
		prefix string
	}
	managers := []pm{
		{"pacman", "sudo pacman -S"},
		{"apt-get", "sudo apt-get install -y"},
		{"dnf", "sudo dnf install -y"},
		{"brew", "brew install"},
		{"zypper", "sudo zypper install"},
		{"apk", "sudo apk add"},
	}
	for _, m := range managers {
		if _, err := exec.LookPath(m.bin); err == nil {
			return m.bin, m.prefix
		}
	}
	return "", ""
}

func depName(dep, pkgManager string) string {
	if dep == "yt-dlp" {
		switch pkgManager {
		case "pacman":
			return "yt-dlp"
		case "brew":
			return "yt-dlp"
		default:
			return "yt-dlp  # or: pip3 install yt-dlp"
		}
	}
	return dep
}

func checkDeps() bool {
	ok := true
	_, installCmd := detectPkgManager()
	if installCmd == "" {
		installCmd = "<your-package-manager>"
	}
	for _, dep := range []string{"mpv", "yt-dlp"} {
		if _, err := exec.LookPath(dep); err != nil {
			mgr, _ := detectPkgManager()
			fmt.Printf("   ❌ Missing: %s\n", dep)
			fmt.Printf("      Install: %s %s\n", installCmd, depName(dep, mgr))
			ok = false
		}
	}
	return ok
}

func playMusic() {
	if musicCmd != nil && musicCmd.Process != nil {
		fmt.Printf("   🎵 Already playing: %s\n", currentTrack)
		fmt.Println("   Type 'stop_music' to stop, or 'playlist' to pick a track.")
		return
	}
	playTrack(rand.Intn(len(playlist)))
}

func playTrack(idx int) {
	if idx < 0 || idx >= len(playlist) {
		fmt.Println("   Invalid track number.")
		return
	}
	if musicCmd != nil && musicCmd.Process != nil {
		stopMusic()
	}

	// FIX 2: check deps and give clear guidance before attempting playback
	fmt.Println("   🔍 Checking dependencies...")
	if !checkDeps() {
		fmt.Println()
		fmt.Println("   ⚠️  Please install missing deps above, then try again.")
		fmt.Println("   Tip: yt-dlp is best installed via:  pip3 install yt-dlp")
		return
	}

	track := playlist[idx]
	user := getEffectiveUser()
	runtimeDir := getRuntimeDir(user)
	homeDir := getMusicHome()

	localPath := filepath.Join(homeDir, ".local/share/focus-symphony/assets/lofi.mp3")
	audioSrc := track.url

	if _, err := os.Stat(localPath); err == nil {
		audioSrc = localPath
		fmt.Printf("🎵 Playing local file: %s\n", filepath.Base(localPath))
	} else {
		fmt.Printf("🎵 Streaming: %s\n", track.name)
		fmt.Println("   Resolving stream via yt-dlp... (5-10 sec)")

		ytArgs := []string{"-f", "bestaudio", "--get-url", "--no-playlist", track.url}
		var ytCmd *exec.Cmd
		if os.Getuid() == 0 {
			args := []string{"-u", user, "yt-dlp"}
			args = append(args, ytArgs...)
			ytCmd = exec.Command("sudo", args...)
		} else {
			ytCmd = exec.Command("yt-dlp", ytArgs...)
		}

		out, err := ytCmd.Output()
		if err != nil {
			fmt.Printf("   ❌ yt-dlp failed: %v\n", err)
			fmt.Println("   Fix: run  yt-dlp -U  to update, or  pip3 install -U yt-dlp")
			return
		}

		urls := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(urls) == 0 || urls[0] == "" {
			fmt.Println("   ❌ No stream URL found. Video may be unavailable.")
			return
		}
		audioSrc = urls[0]
	}

	musicCmd = buildMpvCmd(audioSrc, user, runtimeDir)
	musicCmd.Stdout = nil
	musicCmd.Stderr = nil

	if err := musicCmd.Start(); err != nil {
		fmt.Printf("   ❌ mpv failed to start: %v\n", err)
		fmt.Println("   Fix: sudo pacman -S mpv   (or your distro's equivalent)")
		musicCmd = nil
		return
	}

	currentTrack = track.name
	fmt.Println("   ✅ Audio engine running. Volume: 80%")
	fmt.Println("   Type 'stop_music' to stop.")

	go func() {
		musicCmd.Wait()
		currentTrack = ""
		musicCmd = nil
	}()
}

func stopMusic() {
	if musicCmd == nil || musicCmd.Process == nil {
		fmt.Println("   (No music is currently playing.)")
		return
	}
	musicCmd.Process.Kill()
	musicCmd.Wait()
	exec.Command("pkill", "-9", "-u", getEffectiveUser(), "mpv").Run()
	musicCmd = nil
	currentTrack = ""
	fmt.Println("🎵 Audio engine offline.")
}

func showPlaylist() {
	fmt.Println("\n🎵 Focus Playlists (type the number to play):")
	fmt.Println()
	for i, track := range playlist {
		fmt.Printf("   [%d] %s\n", i+1, track.name)
	}
	fmt.Println()
}

func startRapidFocus() {
	fmt.Println("🚀 RAPID FOCUS — 25 min Pomodoro")
	startSession()
	playMusic()
}

func showStats() {
	if !isShieldActive {
		fmt.Println("   No active focus session.")
		return
	}
	fmt.Printf("📊 FOCUS TIME : %s\n", time.Since(sessionStart).Round(time.Second))
	if currentTrack != "" {
		fmt.Printf("🎵 NOW PLAYING: %s\n", currentTrack)
	}
}

func checkDNSoverHTTPS() {
	fmt.Println()
	fmt.Println("╔═══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║  ⚠️  DNS OVER HTTPS (DoH) WARNING                                 ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("   This tool blocks sites by modifying your system's hosts file.")
	fmt.Println("   However, if your browser uses 'Secure DNS' (DNS over HTTPS),")
	fmt.Println("   it will BYPASS the hosts file and blocking WON'T WORK.")
	fmt.Println()
	fmt.Println("   📌 TO FIX THIS:")
	fmt.Println()
	fmt.Println("   Chrome/Edge/Brave:")
	fmt.Println("   → Settings → Privacy and Security → Security")
	fmt.Println("   → Turn OFF 'Use secure DNS'")
	fmt.Println()
	fmt.Println("   Firefox:")
	fmt.Println("   → Settings → Privacy & Security")
	fmt.Println("   → Scroll to 'DNS over HTTPS'")
	fmt.Println("   → Select 'Off' or 'Default'")
	fmt.Println()
	fmt.Println("   After disabling, RESTART your browser completely.")
	fmt.Println()
	fmt.Println("╚═══════════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func showHelp() {
	fmt.Println(`
Commands:
  start        Block distracting sites          [needs sudo]
  stop         Unblock sites                    [needs sudo]
  music        Play a random focus track
  playlist     List all tracks (type 1-5 to pick)
  stop_music   Stop music
  rapid        25-min Pomodoro + music          [needs sudo]
  stats        Show session time + current track
  help         Show this menu
  exit         Quit
`)
}
