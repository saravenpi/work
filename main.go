package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/saravenpi/work/internal/ui"
)

const (
	version = "1.0.0"
	appName = "work"
)

// main initializes and runs the work timer application
func main() {
	var (
		helpFlag    = flag.Bool("h", false, "Display help message")
		helpLong    = flag.Bool("help", false, "Display help message")
		versionFlag = flag.Bool("v", false, "Display version")
		historyFlag = flag.Bool("history", false, "Show work history")
	)

	flag.Parse()

	if *helpFlag || *helpLong {
		printHelp()
		return
	}

	if *versionFlag {
		fmt.Printf("%s version %s\n", appName, version)
		return
	}

	var m tea.Model
	if *historyFlag {
		m = ui.NewHistoryModel()
	} else {
		m = ui.NewTimerModel()
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}

// printHelp displays usage information for the application
func printHelp() {
	fmt.Printf("Usage: %s [options]\n", appName)
	fmt.Println("A Pomodoro timer for focused work sessions")
	fmt.Println("\nOptions:")
	fmt.Println("  -h, --help     Display this help message")
	fmt.Println("  -v             Display version")
	fmt.Println("  --history      View work session history")
	fmt.Println("\nControls:")
	fmt.Println("  space          Start/pause timer")
	fmt.Println("  r              Reset timer")
	fmt.Println("  q/ctrl+c       Quit")
}