package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	fmt.Println(`
  ____ ___ ____  _   _ _____ 
 | __ )_ _/ ___|| | | |_   _|
 |  _ \| |\___ \| |_| | | |  
 | |_) | | ___) |  _  | | |  
 |____/___|____/|_| |_| |_|  
                               `)
	fmt.Println("FOCUS-SYMPHONY v1.1.0")
	fmt.Println("Harmonizing Linux Performance for Deep Work")
	fmt.Println("-------------------------------------------")

	// Check if we have write access to /etc/hosts
	file, err := os.OpenFile(hostsPath, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("ERROR: Insufficient privileges to modify /etc/hosts.")
		fmt.Println("Please run with: sudo focus-symphony")
		return
	}
	file.Close()

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
		case "music":
			playMusic()
		case "help":
			showHelp()
		case "exit":
			stopSession() // Cleanup before exit
			fmt.Println("Exiting Focus-Symphony. Goodbye!")
			return
		default:
			fmt.Printf("Unknown command: %s. Type 'help' for commands.\n", input)
		}
	}
}

func startSession() {
	fmt.Println("Initializing Focus Session...")
	
	// First, clean up any existing blocks to avoid duplicates
	cleanHosts()

	f, err := os.OpenFile(hostsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open hosts file: %v\n", err)
		return
	}
	defer f.Close()

	fmt.Fprintln(f, "\n"+blockMarker)
	for _, site := range sitesToBlock {
		fmt.Fprintf(f, "127.0.0.1 %s\n", site)
	}
	fmt.Fprintln(f, blockMarker)

	time.Sleep(1 * time.Second)
	fmt.Println("Acoustic Shield ACTIVATED: Distracting sites blocked.")
	fmt.Println("Session ACTIVE. Deep work mode enabled.")
}

func stopSession() {
	fmt.Println("Ending Focus Session...")
	cleanHosts()
	time.Sleep(1 * time.Second)
	fmt.Println("Acoustic Shield DEACTIVATED: Sites unblocked.")
	fmt.Println("Session STOPPED.")
}

func cleanHosts() {
	input, err := os.ReadFile(hostsPath)
	if err != nil {
		return
	}

	lines := strings.Split(string(input), "\n")
	var newLines []string
	isBlockingSection := false

	for _, line := range lines {
		if strings.TrimSpace(line) == blockMarker {
			isBlockingSection = !isBlockingSection
			continue
		}
		if !isBlockingSection {
			newLines = append(newLines, line)
		}
	}

	// Remove trailing empty lines and write back
	output := strings.Join(newLines, "\n")
	os.WriteFile(hostsPath, []byte(output), 0644)
}

func playMusic() {
	fmt.Println("Loading Focus Playlist...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Now playing: 'Lofi Beats for Coding' (YouTube/MP3 simulation)")
}

func showHelp() {
	fmt.Println("Available Commands:")
	fmt.Println("  start  - Begin a focus session (actually blocks sites)")
	fmt.Println("  stop   - End the focus session (unblocks sites)")
	fmt.Println("  music  - Launch terminal music player")
	fmt.Println("  help   - Show this help message")
	fmt.Println("  exit   - Close Focus-Symphony")
}
