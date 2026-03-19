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
	"Midnight Coding Session (Lofi)",
	"The Conductor's Deep Space Mix",
	"Focus Symphony - Bass Boosted",
	"Ambient Rain & Code",
	"Cyberpunk Chill Engine",
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
	fmt.Println("FOCUS-SYMPHONY v1.5.0 (The Audio Overhaul)")
	fmt.Println("-------------------------------------------")

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
			fmt.Printf("Unknown command: %s. Type 'help'.\n", input)
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
	fmt.Println("SUCCESS: Sites blocked. Deep work mode active.")
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

func playMusic() {
	if musicCmd != nil && musicCmd.Process != nil {
		fmt.Println("   (Music is already playing. Type 'stop_music' first.)")
		return
	}

	track := playlist[rand.Intn(len(playlist))]
	fmt.Printf("🎵 Injecting Audio: %s\n", track)
	
	// Ultra-stable stream
	streamURL := "http://usa9.fastcast4u.com:8014/1" 

	user := os.Getenv("SUDO_USER")
	if user == "" {
		user = os.Getenv("USER")
	}

	// CRITICAL: We need to pass the audio environment variables to sudo
	// This usually fixes the "no sound" issue in Linux
	runtimeDir := "/run/user/1000" // Default for primary user
	if u, err := exec.Command("id", "-u", user).Output(); err == nil {
		runtimeDir = "/run/user/" + strings.TrimSpace(string(u))
	}

	// Try mpv with environment variables
	musicCmd = exec.Command("sudo", "-u", user, "env", "XDG_RUNTIME_DIR="+runtimeDir, "mpv", "--no-video", "--volume=80", streamURL)
	err := musicCmd.Start()
	
	if err != nil {
		fmt.Printf("   ❌ Audio Engine Error: %v\n", err)
	} else {
		fmt.Println("   🚀 SOUND WAVE ACTIVATED. Hearing the beats now?")
		fmt.Println("   (It may take 5 seconds to buffer...)")
	}
}

func stopMusic() {
	if musicCmd != nil && musicCmd.Process != nil {
		musicCmd.Process.Kill()
		musicCmd = nil
		fmt.Println("🎵 Audio Engine Offline.")
	}
}

func startRapidFocus() {
	fmt.Println("🚀 RAPID FOCUS INITIATED (25m)")
	startSession()
}

func showStats() {
	if !isShieldActive {
		fmt.Println("No active session.")
		return
	}
	fmt.Printf("📊 FOCUS TIME: %s\n", time.Since(sessionStart).Round(time.Second))
}

func showHelp() {
	fmt.Println("Commands: start, stop, music, stop_music, rapid, stats, exit")
}
