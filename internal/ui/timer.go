package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/saravenpi/work/internal/models"
	"github.com/saravenpi/work/internal/storage"
)

type tickMsg time.Time

// TimerModel represents the state and behavior of the Pomodoro timer interface
type TimerModel struct {
	startTime    time.Time
	elapsed      time.Duration
	targetTime   time.Duration
	isRunning    bool
	isFinished   bool
	progress     progress.Model
	width        int
	height       int
	sessionSaved bool
}

// NewTimerModel creates a new timer model with default settings
func NewTimerModel() TimerModel {
	p := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	p.ShowPercentage = false
	return TimerModel{
		targetTime: time.Duration(models.SessionLengthSeconds) * time.Second,
		progress:   p,
	}
}

// Init initializes the timer model and returns any initial commands
func (m TimerModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the timer model state
func (m TimerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.progress.Width = msg.Width - 4
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			if m.isRunning && !m.isFinished {
				session := models.NewWorkSession(m.startTime, m.elapsed)
				_ = storage.SaveSession(session)
			}
			return m, tea.Quit
		case " ", "space":
			if !m.isFinished {
				m.isRunning = !m.isRunning
				if m.isRunning && m.startTime.IsZero() {
					m.startTime = time.Now()
				}
				if m.isRunning {
					return m, tickCmd()
				}
			}
			return m, nil
		case "r":
			m = NewTimerModel()
			m.progress.Width = m.width - 4
			return m, nil
		case "enter":
			if m.isFinished {
				m = NewTimerModel()
				m.progress.Width = m.width - 4
				return m, nil
			}
		}

	case tickMsg:
		if m.isRunning && !m.isFinished {
			m.elapsed = time.Since(m.startTime)
			
			if m.elapsed >= m.targetTime {
				m.isFinished = true
				m.isRunning = false
				if !m.sessionSaved {
					session := models.NewWorkSession(m.startTime, m.elapsed)
					if err := storage.SaveSession(session); err == nil {
						m.sessionSaved = true
					}
				}
				return m, nil
			}
			return m, tickCmd()
		}

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

// tickCmd creates a command that sends a tick message every second
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// View renders the timer interface as a string
func (m TimerModel) View() string {
	if m.width == 0 {
		return "\n  Initializing..."
	}

	var content string

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF7CCB"))

	timeStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FDFF8C"))

	messageStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7D56F4"))

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262"))

	content += titleStyle.Render("✨ Work Session Timer") + "\n"

	if m.isFinished {
		content += messageStyle.Render("🎉 Session Complete! Great work! Take a break.") + "\n"
		if m.sessionSaved {
			content += lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00FF00")).
				Render("✓ Session saved to history") + "\n"
		}
		content += helpStyle.Render("[enter] new session • [q] quit")
	} else {
		remaining := m.targetTime - m.elapsed
		if remaining < 0 {
			remaining = 0
		}

		minutes := int(remaining.Minutes())
		seconds := int(remaining.Seconds()) % 60
		timeDisplay := fmt.Sprintf("%02d:%02d", minutes, seconds)

		content += timeStyle.Render(timeDisplay) + "\n"

		var endTimeMsg string
		if m.startTime.IsZero() {
			endTime := time.Now().Add(m.targetTime)
			endTimeMsg = fmt.Sprintf("Will end at: %s", endTime.Format("15:04:05"))
		} else if m.isRunning {
			endTime := m.startTime.Add(m.targetTime)
			endTimeMsg = fmt.Sprintf("Ends at: %s", endTime.Format("15:04:05"))
		}
		if endTimeMsg != "" {
			content += messageStyle.Render(endTimeMsg) + "\n"
		}

		percent := float64(m.elapsed) / float64(m.targetTime)
		if percent > 1.0 {
			percent = 1.0
		}
		content += m.progress.ViewAs(percent) + "\n"

		var statusMsg string
		if m.isRunning {
			statusMsg = "⏸ [space] pause • [r] reset • [q] quit"
		} else if m.startTime.IsZero() {
			statusMsg = "▶ [space] start • [r] reset • [q] quit"
		} else {
			statusMsg = "▶ [space] resume • [r] reset • [q] quit"
		}
		content += helpStyle.Render(statusMsg)
	}

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(content)
}