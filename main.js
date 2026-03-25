/* focus-symphony — modular JS */

function initTerminal() {
  const body = document.querySelector('.fs-term-body');
  if (!body) return;

  const lines = [
    { type: 'ascii', text: '  ____ ___ ____  _   _ _____ \n | __ )_ _/ ___|| | | |_   _|\n |  _ \\| |\\___ \\| |_| | | |  \n | |_) | | ___) |  _  | | |  \n |____/___|____/|_| |_| |_|  ' },
    { type: 'plain', text: 'Open Source project by BISHT x FLAVOURTOWN' },
    { type: 'prompt', text: 'help' },
    { type: 'plain', text: 'Available Commands:' },
    { type: 'indent', text: 'start  — Begin a focus session (blocks sites)' },
    { type: 'indent', text: 'stop   — End the focus session' },
    { type: 'indent', text: 'music  — Launch terminal music player' },
    { type: 'indent', text: 'exit   — Close Focus-Symphony' },
  ];

  body.innerHTML = '';

  function buildLine(line) {
    const el = document.createElement('div');
    el.className = 'fs-term-output';
    if (line.type === 'ascii') {
      el.style.cssText = 'color:#f97316;font-size:0.44rem;margin-bottom:0.8rem;line-height:1.3;white-space:pre;';
      el.textContent = line.text;
    } else if (line.type === 'prompt') {
      el.innerHTML = '<span class="prompt">fs &gt;</span> ' + line.text;
    } else if (line.type === 'indent') {
      el.innerHTML = '&nbsp;&nbsp;' + line.text;
    } else {
      el.textContent = line.text;
    }
    return el;
  }

  lines.forEach((line, i) => {
    const el = buildLine(line);
    body.appendChild(el);
    setTimeout(() => el.classList.add('visible'), 300 + i * 220);
  });

  const cursorEl = document.createElement('div');
  cursorEl.innerHTML = '<span class="prompt">fs &gt;</span> <span class="cursor"></span>';
  setTimeout(() => body.appendChild(cursorEl), 300 + lines.length * 220);
}

function initMusicPlayer() {
  const buttons = document.querySelectorAll('.fs-btn-sm');
  let currentPlaying = null;

  buttons.forEach((btn, idx) => {
    btn.addEventListener('click', () => {
      if (currentPlaying === idx) {
        btn.textContent = '▶ PLAY';
        btn.classList.remove('playing');
        currentPlaying = null;
        return;
      }
      buttons.forEach(b => { b.textContent = '▶ PLAY'; b.classList.remove('playing'); });
      btn.textContent = '⏸ PAUSE';
      btn.classList.add('playing');
      currentPlaying = idx;
    });
  });
}

function initStatBars() {
  const fills = document.querySelectorAll('.bar-fill');
  if (!fills.length) return;

  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        const fill = entry.target;
        const target = fill.dataset.width;
        fill.style.width = '0';
        requestAnimationFrame(() => { fill.style.width = target; });
        observer.unobserve(fill);
      }
    });
  }, { threshold: 0.3 });

  fills.forEach(fill => {
    fill.dataset.width = fill.style.width;
    fill.style.width = '0';
    observer.observe(fill);
  });
}

function initNavHighlight() {
  const links = document.querySelectorAll('.fs-nav-links a');
  const sections = document.querySelectorAll('.fs > section, .fs > footer');

  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        const idx = Array.from(sections).indexOf(entry.target);
        links.forEach((l, i) => { l.style.color = i === idx ? 'var(--accent)' : ''; });
      }
    });
  }, { rootMargin: '-40% 0px -55%' });

  sections.forEach(s => observer.observe(s));
}

function initSocials() {
  const badges = document.querySelectorAll('.fs-socials span');
  const urls = {
    'GITHUB': 'https://github.com/cherry-bisht/focus-symphony-',
    'DISCORD': '#',
    'DOCS': '#',
  };
  badges.forEach(badge => {
    badge.addEventListener('click', () => {
      const url = urls[badge.textContent.trim()];
      if (url && url !== '#') window.open(url, '_blank', 'noopener');
    });
  });
}

function initCTA() {
  const githubUrl = 'https://github.com/cherry-bisht/focus-symphony-';
  [document.querySelector('.fs-btn'), document.querySelector('.fs-footer-btn')].forEach(btn => {
    if (btn) btn.addEventListener('click', () => window.open(githubUrl, '_blank', 'noopener'));
  });
}

function initFocusSymphony() {
  initTerminal();
  initMusicPlayer();
  initStatBars();
  initNavHighlight();
  initSocials();
  initCTA();
}

document.readyState === 'loading'
  ? document.addEventListener('DOMContentLoaded', initFocusSymphony)
  : initFocusSymphony();
