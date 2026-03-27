package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const hostsPath = "/etc/hosts"
const blockMarker = "# FOCUS-SYMPHONY-BLOCK"

var sitesToBlock = []string{
	"www.youtube.com", "youtube.com",
	"www.reddit.com", "reddit.com",
	"www.twitter.com", "twitter.com", "x.com",
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
	fmt.Println("FOCUS-SYMPHONY v1.6.0 (yt-dlp Audio Engine)")
	fmt.Println("--------------------------------------------")
	fmt.Println()
	showHelp() // FIX 1: show help menu on startup

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
	if os.Getuid() != 0 {
		relaunchWithSudo()
		return
	}
	fmt.Println("⚡ Activating Acoustic Shield...")
	cleanHosts()
	isShieldActive = true
	sessionStart = time.Now()
	f, err := os.OpenFile(hostsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("   ❌ Error: %v\n", err)
		return
	}
	defer f.Close()
	fmt.Fprintln(f, blockMarker)
	for _, site := range sitesToBlock {
		fmt.Fprintf(f, "127.0.0.1 %s\n", site)
		fmt.Fprintf(f, "::1       %s\n", site)
	}
	fmt.Fprintln(f, blockMarker)
	fmt.Println("✅ Sites blocked. Deep work mode active.")
	fmt.Println("   Tip: type 'music' to start your focus playlist")
}

func stopSession() {
	if isShieldActive {
		if os.Getuid() != 0 {
			relaunchWithSudo()
			return
		}
		fmt.Println("⚡ Deactivating Acoustic Shield...")
		cleanHosts()
		isShieldActive = false
		fmt.Println("✅ World access restored.")
	} else {
		fmt.Println("   (No active shield to stop.)")
	}
}

func cleanHosts() {
	input, _ := os.ReadFile(hostsPath)
	lines := strings.Split(string(input), "\n")
	var newLines []string
	isBlocking := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == blockMarker {
			isBlocking = !isBlocking
			continue
		}
		if !isBlocking && trimmed != "" {
			newLines = append(newLines, line)
		}
	}
	os.WriteFile(hostsPath, []byte(strings.Join(newLines, "\n")+"\n"), 0644)
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
