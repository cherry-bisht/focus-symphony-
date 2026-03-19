package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	fmt.Println("FOCUS-SYMPHONY v1.2.0 (Deep Work Mode)")
	fmt.Println("-------------------------------------------")

	// Verify we can write to /etc/hosts
	f, err := os.OpenFile(hostsPath, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERROR: Permission Denied!")
		fmt.Println("Run with: sudo focus-symphony")
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
		case "music":
			playMusic()
		case "help":
			showHelp()
		case "exit":
			stopSession()
			fmt.Println("Exiting. Keep Focus!")
			return
		default:
			fmt.Printf("Unknown command: %s\n", input)
		}
	}
}

func startSession() {
	fmt.Println("Activating Acoustic Shield...")
	cleanHosts() // Remove old blocks

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

	fmt.Println("SUCCESS: Sites are now blocked globally.")
}

func stopSession() {
	fmt.Println("Deactivating Acoustic Shield...")
	cleanHosts()
	fmt.Println("SUCCESS: Sites are now unblocked.")
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
	fmt.Println("Now playing: Focus Symphony #1 in G Minor...")
}

func showHelp() {
	fmt.Println("Commands: start, stop, music, exit")
}
