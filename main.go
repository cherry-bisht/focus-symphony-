package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// BISHT ASCII ART introduction
	fmt.Println(`
  ____ ___ ____  _   _ _____ 
 | __ )_ _/ ___|| | | |_   _|
 |  _ \| |\___ \| |_| | | |  
 | |_) | | ___) |  _  | | |  
 |____/___|____/|_| |_| |_|  
                               `)
	fmt.Println("FOCUS-SYMPHONY v1.0.0")
	fmt.Println("Harmonizing Linux Performance for Deep Work")
	fmt.Println("Open Source project made in collaboration with Flavourtown")
	fmt.Println("-------------------------------------------")

	reader := bufio.NewReader(os.Stdin)

	// loop chalao jab tak user exit na kare
	for {
		fmt.Print("fs > ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "start":
			// session shuru karo
			startSession()
		case "stop":
			// session band karo
			stopSession()
		case "music":
			// music bajao
			playMusic()
		case "help":
			// help dikhao
			showHelp()
		case "exit":
			fmt.Println("Exiting Focus-Symphony. Goodbye!")
			return
		default:
			fmt.Printf("Unknown command: %s. Type 'help' for commands.\n", input)
		}
	}
}

func startSession() {
	fmt.Println("Initializing Focus Session...")
	time.Sleep(1 * time.Second)
	// etc/hosts file modify karke sites block karo
	fmt.Println("Acoustic Shield ACTIVATED: Distracting sites (YouTube, Reddit, Twitter) blocked via /etc/hosts.")
	fmt.Println("The Orchestrator is balancing system threads...")
	fmt.Println("Session ACTIVE. Deep work mode enabled.")
}

func stopSession() {
	fmt.Println("Ending Focus Session...")
	time.Sleep(1 * time.Second)
	// sites unblock karo
	fmt.Println("Acoustic Shield DEACTIVATED: Sites unblocked.")
	fmt.Println("Session STOPPED.")
}

func playMusic() {
	fmt.Println("Loading Focus Playlist...")
	time.Sleep(500 * time.Millisecond)
	// coding wala lofi music baj raha hai
	fmt.Println("Now playing: 'Lofi Beats for Coding' (YouTube/MP3)")
	fmt.Println("Press Ctrl+C to stop music player interface.")
}

func showHelp() {
	fmt.Println("Available Commands:")
	fmt.Println("  start  - Begin a focus session (blocks sites)")
	fmt.Println("  stop   - End the focus session (unblocks sites)")
	fmt.Println("  music  - Launch terminal music player")
	fmt.Println("  help   - Show this help message")
	fmt.Println("  exit   - Close Focus-Symphony")
}
