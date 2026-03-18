const menuBtn = document.getElementById("menu-btn");
const navLinks = document.getElementById("nav-links");
const menuBtnIcon = menuBtn.querySelector("i");

// 1. Mobile Menu Logic
// menu kholo ya band karo click pe
menuBtn.addEventListener("click", () => {
  navLinks.classList.toggle("open");

  const isOpen = navLinks.classList.contains("open");
  menuBtnIcon.setAttribute(
    "class",
    isOpen ? "ri-close-line" : "ri-menu-3-line"
  );
});

// Link click hone pe menu apne aap band ho jayega
navLinks.addEventListener("click", () => {
  navLinks.classList.remove("open");
  menuBtnIcon.setAttribute("class", "ri-menu-3-line");
});

// 2. ScrollReveal Animations
// animations thoda fast rakhte hai modern tech feel ke liye
const scrollRevealOption = {
  distance: "60px",
  origin: "bottom",
  duration: 1200, 
};

// Title ko reveal karo
ScrollReveal().reveal(".header__container h1", {
  ...scrollRevealOption,
  delay: 1500, 
});

// Description ko reveal karo
ScrollReveal().reveal(".header__container p", {
  ...scrollRevealOption,
  delay: 2000,
});

// Button ko reveal karo
ScrollReveal().reveal(".header__container .header__btn", {
  ...scrollRevealOption,
  delay: 2500,
});


// Social icons ek ek karke aayenge
ScrollReveal().reveal(".socials li", {
  ...scrollRevealOption,
  delay: 3000,
  interval: 300, 
});

// 3. Terminal Simulation Logic
// terminal mein typing wala effect dalte hai
const terminalOutput = document.getElementById("terminal-output");
const openTerminalBtn = document.getElementById("open-terminal");

// Commands aur unke responses, bina kisi emoji ke
const commands = [
  { cmd: "start", res: ["Initializing Focus Session...", "Acoustic Shield ACTIVATED: Distracting sites blocked.", "The Orchestrator is balancing threads...", "Session ACTIVE. Deep work mode enabled."] },
  { cmd: "music", res: ["Loading Focus Playlist...", "Now playing: 'Lofi Beats for Coding'"] },
  { cmd: "stop", res: ["Ending Focus Session...", "Acoustic Shield DEACTIVATED.", "Session STOPPED."] }
];

let commandIndex = 0;

function simulateTerminal() {
  if (commandIndex >= commands.length) {
    commandIndex = 0;
    terminalOutput.innerHTML = `
            <pre style="color: #6366f1; font-size: 0.7rem; margin-bottom: 1rem; font-family: monospace;">
  ____ ___ ____  _   _ _____ 
 | __ )_ _/ ___|| | | |_   _|
 |  _ \| |\___ \| |_| | | |  
 | |_) | | ___) |  _  | | |  
 |____/___|____/|_| |_| |_|  
            </pre>
            <div class="line">Open Source project by BISHT x FLAVOURTOWN</div>
            <div class="line"><span class="prompt">fs ></span> help</div>
            <div class="line">Available Commands: start, stop, music, exit</div>`;
  }

  const current = commands[commandIndex];
  
  // Naya command line create karo
  const cmdLine = document.createElement("div");
  cmdLine.className = "line";
  cmdLine.innerHTML = `<span class="prompt">fs ></span> `;
  terminalOutput.appendChild(cmdLine);

  let charIndex = 0;
  const typeInterval = setInterval(() => {
    cmdLine.innerHTML += current.cmd[charIndex];
    charIndex++;
    if (charIndex >= current.cmd.length) {
      clearInterval(typeInterval);
      
      // Thoda wait karke response dikhao
      setTimeout(() => {
        current.res.forEach((line, i) => {
          setTimeout(() => {
            const resLine = document.createElement("div");
            resLine.className = "line";
            resLine.textContent = line;
            terminalOutput.appendChild(resLine);
            terminalOutput.scrollTop = terminalOutput.scrollHeight;
          }, i * 500);
        });
        
        // Saare response aane ke baad agla command line
        setTimeout(() => {
          commandIndex++;
          simulateTerminal();
        }, current.res.length * 500 + 2000);
      }, 500);
    }
  }, 100);
}

// Button click pe terminal simulation start hogi
openTerminalBtn.addEventListener("click", () => {
  document.getElementById("terminal").scrollIntoView({ behavior: "smooth" });
  if (commandIndex === 0 && terminalOutput.children.length <= 8) {
    simulateTerminal();
  }
});

// Baaki sections ko bhi reveal karo
ScrollReveal().reveal(".terminal-demo", scrollRevealOption);
ScrollReveal().reveal(".music-player", scrollRevealOption);
ScrollReveal().reveal(".playlist-card", { ...scrollRevealOption, interval: 200 });
