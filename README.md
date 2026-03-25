# Focus-Symphony 🎻

**Focus-Symphony** is an open-source productivity suite designed for developers who want to achieve deep work without digital noise. It combines a powerful Go-based terminal orchestrator with a sleek, responsive landing page.

  ____ ___ ____  _   _ _____ 
 | __ )_ _/ ___|| | | |_   _|
 |  _ \| |\___ \| |_| | | |  
 | |_) | | ___) |  _  | | |  
 |____/___|____/|_| |_| |_|  

---

## 🚀 Key Instruments (Features)

### 🎼 The Orchestrator (CLI)
A high-performance Go engine that balances your session:
*   **Acoustic Shield**: Automatically blocks distracting websites (YouTube, Reddit, Twitter) by managing `/etc/hosts`.
*   **Terminal Music Player**: Stream focus-optimized playlists directly from your terminal.
*   **System Telemetry**: Real-time monitoring of your focus session and system performance.

### 🎨 Visual Rhythm (Web)
A modern landing page designed with **Vanilla CSS** and **ScrollReveal**:
*   **Terminal Simulator**: Interactive demonstration of CLI commands.
*   **Responsive Design**: Optimized for all devices, ensuring your focus stats are always visible.
*   **Modern Aesthetics**: Clean typography and interactive animations.

---

## 🛠️ Global Access: Use Anywhere

To use **Focus-Symphony** from any directory in your terminal, follow these steps:

1.  **Build and Install**:
    ```bash
    go build -o focus-symphony main.go
    mv focus-symphony ~/.local/bin/
    
    # Install Assets (REQUIRED for Local Music)
    mkdir -p ~/.local/share/focus-symphony/assets
    cp assets/lofi.mp3 ~/.local/share/focus-symphony/assets/
    ```
2.  **Run Command**:
    Now simply type `focus-symphony` in any terminal window.

---

## 🎮 How to Conduct Your Session

Available commands in the Orchestrator:
*   `start`  - Activates the Acoustic Shield and begins your focus session.
*   `music`  - Launches the terminal music player with focus playlists.
*   `stop`   - Deactivates site blocking and restores system defaults.
*   `help`   - Displays available session controls.
*   `exit`   - Gracefully closes the Orchestrator.

---

## 🛠️ Technical Stack

- **Backend**: Go (Orchestrator Engine)
- **Frontend**: HTML5, Vanilla CSS, JavaScript (Visual Rhythm)
- **Animations**: ScrollReveal.js
- **Icons**: Remix Icon

---

## 🤝 Collaboration
This project is an Open Source initiative by **BISHT** in collaboration with **Flavourtown**.

---

*Built with passion for the developer community. Harmonizing Linux Performance for Deep Work.*
