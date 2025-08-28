package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/saravenpi/work/internal/storage"
)

// HistoryModel represents the state and behavior of the history viewing interface
type HistoryModel struct {
	viewport viewport.Model
	content  string
	ready    bool
	width    int
	height   int
}

// NewHistoryModel creates a new history viewing model
func NewHistoryModel() HistoryModel {
	return HistoryModel{}
}

// Init initializes the history model and starts loading history data
func (m HistoryModel) Init() tea.Cmd {
	return m.loadHistory
}

// loadHistory loads and formats work session history for display
func (m HistoryModel) loadHistory() tea.Msg {
	sessions, err := storage.LoadHistory()
	if err != nil {
		return fmt.Sprintf("Error loading history: %v", err)
	}

	if len(sessions) == 0 {
		return "No work sessions recorded yet.\n\nStart a new session with: work"
	}

	var content strings.Builder
	content.WriteString("📚 Work Session History\n")
	content.WriteString("═══════════════════════════════════════════\n\n")

	totalTime := 0.0
	for i := len(sessions) - 1; i >= 0; i-- {
		session := sessions[i]
		duration := session.Duration
		minutes := int(duration.Minutes())
		
		content.WriteString(fmt.Sprintf("Session %d:\n", len(sessions)-i))
		content.WriteString(fmt.Sprintf("  Started: %s\n", session.Start.Format("2006-01-02 15:04:05")))
		content.WriteString(fmt.Sprintf("  Ended:   %s\n", session.End.Format("2006-01-02 15:04:05")))
		content.WriteString(fmt.Sprintf("  Duration: %d minutes\n", minutes))
		content.WriteString("\n")
		
		totalTime += duration.Minutes()
	}

	content.WriteString("═══════════════════════════════════════════\n")
	content.WriteString(fmt.Sprintf("Total Sessions: %d\n", len(sessions)))
	content.WriteString(fmt.Sprintf("Total Time: %.0f minutes (%.1f hours)\n", totalTime, totalTime/60))

	return content.String()
}

// Update handles messages and updates the history model state
func (m HistoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerHeight := 2
		footerHeight := 2
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width-4, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.ready = true
		} else {
			m.viewport.Width = msg.Width - 4
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		if m.content != "" {
			m.viewport.SetContent(m.content)
		}

	case string:
		m.content = msg
		if m.ready {
			m.viewport.SetContent(m.content)
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	if m.ready {
		m.viewport, cmd = m.viewport.Update(msg)
	}

	return m, cmd
}

// View renders the history interface as a string
func (m HistoryModel) View() string {
	if !m.ready {
		return "\n  Loading history..."
	}

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF7CCB"))

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262"))

	header := headerStyle.Render("Work History")
	footer := footerStyle.Render("↑/↓ scroll • q quit")
	
	scrollInfo := ""
	if m.viewport.TotalLineCount() > m.viewport.Height {
		scrollPercent := fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100)
		scrollInfo = footerStyle.Render(fmt.Sprintf(" • %s", scrollPercent))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		m.viewport.View(),
		footer+scrollInfo,
	)

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Padding(0, 2).
		Render(content)
}