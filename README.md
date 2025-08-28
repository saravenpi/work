# Work - Pomodoro Timer

A beautiful terminal-based Pomodoro timer written in Go with an interactive TUI. Built with the amazing [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework by [Charm](https://charm.sh/) 💖

## Features

- 🎯 **25-minute focused work sessions** (classic Pomodoro technique)
- ⏱️ **Interactive timer** with real-time countdown and progress bar
- 📊 **Session history tracking** with detailed statistics
- 🎨 **Beautiful TUI** with colors, animations, and responsive design
- 💾 **Automatic session saving** to `~/.workhistory`
- ⌨️ **Intuitive keyboard controls**

## Installation

### Using the install script
```bash
./install.sh
```

### Manual installation
```bash
make
sudo cp work /usr/local/bin/
```

### Build from source
```bash
go mod download
make
```

## Usage

### Start a work session
```bash
work
```

### View session history
```bash
work --history
```

### Command line options
```bash
work -h         # Show help
work -v         # Show version
work --history  # View work session history
```

## Controls

### During a timer session:
- **Space** - Start/pause timer
- **R** - Reset timer
- **Q** or **Ctrl+C** - Quit (saves partial session)

### In history view:
- **↑/↓** - Scroll through history
- **Q** - Quit

## Technical Details

Built with Go and the incredible [Bubble Tea TUI framework](https://github.com/charmbracelet/bubbletea), which makes building delightful terminal applications a joy. Special thanks to the [Charm team](https://charm.sh/) for creating such beautiful and functional tools for the terminal! 🌟

## Requirements

- Go 1.20 or higher
- Terminal with UTF-8 support and color capabilities

## History File

Work sessions are automatically saved to `~/.workhistory` in a simple, readable format:
```
- 2024/01/15|09:30:00 -> 2024/01/15|09:55:00
- 2024/01/15|10:15:00 -> 2024/01/15|10:40:00
```
