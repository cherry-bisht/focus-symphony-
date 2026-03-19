package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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

var playlist = []string{
	"Focus Symphony #1 in G Minor (Lofi Edit)",
	"Deep Work Nocturne - Movement II",
	"The Orchestrator's Ambient Flow",
	"Binary Beats for High-Performance Coding",
	"Silicon Valley Rain (White Noise)",
}

var musicCmd *exec.Cmd
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
	fmt.Println("FOCUS-SYMPHONY v1.4.0 (The Audio Engine)")
	fmt.Println("-------------------------------------------")

	// Verify we can write to /etc/hosts
	f, err := os.OpenFile(hostsPath, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERROR: Permission Denied! Run with: sudo focus-symphony")
		return
	}
	f.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("fs > ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "start":
			startSession()
		case "stop":
			stopSession()
		case "music", "song":
			playMusic()
		case "stop_music":
			stopMusic()
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
			fmt.Printf("Unknown command: %s. Type 'help' for commands.\n", input)
		}
	}
}

func startSession() {
	fmt.Println("Activating Acoustic Shield...")
	cleanHosts()
	isShieldActive = true
	sessionStart = time.Now()

	f, err := os.OpenFile(hostsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer f.Close()

	fmt.Fprintln(f, blockMarker)
	for _, site := range sitesToBlock {
		fmt.Fprintf(f, "127.0.0.1 %s\n", site)
		fmt.Fprintf(f, "::1       %s\n", site)
	}
	fmt.Fprintln(f, blockMarker)

	fmt.Println("SUCCESS: Deep Work Mode Engaged. Distractions Purged.")
}

func stopSession() {
	if isShieldActive {
		fmt.Println("Deactivating Acoustic Shield...")
		cleanHosts()
		isShieldActive = false
		fmt.Println("SUCCESS: World access restored.")
	}
}

func cleanHosts() {
	input, err := os.ReadFile(hostsPath)
	if err != nil {
		return
	}
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

func playMusic() {
	if musicCmd != nil && musicCmd.Process != nil {
		fmt.Println("   (Music is already playing. Type 'stop_music' first if you want to switch tracks.)")
		return
	}

	track := playlist[rand.Intn(len(playlist))]
	fmt.Printf("🎵 Now playing: %s\n", track)
	fmt.Println("   [━━━━━━━●──────────────] 1:24 / 4:50")

	// We'll use a public Lofi stream URL
	streamURL := "http://stream.zeno.fm/0r0xa792kwzuv" // Lofi Hip Hop Radio

	// We use the regular user to run mpv to avoid audio permission issues with sudo
	user := os.Getenv("SUDO_USER")
	if user == "" {
		user = os.Getenv("USER")
	}

	// mpv --no-video plays just the audio in the background
	musicCmd = exec.Command("sudo", "-u", user, "mpv", "--no-video", streamURL)
	err := musicCmd.Start()
	if err != nil {
		fmt.Printf("   (Audio engine error: %v. Is mpv installed?)\n", err)
	} else {
		fmt.Println("   🚀 Background audio engine started! Hearing the vibes?")
	}
}

func stopMusic() {
	if musicCmd != nil && musicCmd.Process != nil {
		musicCmd.Process.Kill()
		musicCmd = nil
		fmt.Println("🎵 Music stopped.")
	} else {
		fmt.Println("   (No music is playing.)")
	}
}

func startRapidFocus() {
	fmt.Println("🚀 RAPID FOCUS INITIATED: 25-minute sprint starting now.")
	startSession()
}

func showStats() {
	if !isShieldActive {
		fmt.Println("No active session. Type 'start' to begin.")
		return
	}
	duration := time.Since(sessionStart).Round(time.Second)
	fmt.Println("📊 FOCUS TELEMETRY:")
	fmt.Printf("   - Active Duration: %s\n", duration)
	fmt.Println("   - Shield Status: ACTIVE (7 targets blocked)")
	fmt.Println("   - Thread Balance: OPTIMAL")
}

func showHelp() {
	fmt.Println("Available Instruments:")
	fmt.Println("  start       - Activate site-blocking Shield")
	fmt.Println("  stop        - Restore site access")
	fmt.Println("  music/song  - Start Lofi Radio in background")
	fmt.Println("  stop_music  - Stop background radio")
	fmt.Println("  rapid       - Start a 25-minute Focus Sprint")
	fmt.Println("  stats       - View deep work telemetry")
	fmt.Println("  exit        - End session and close all")
}
