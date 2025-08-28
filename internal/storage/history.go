package storage

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/saravenpi/work/internal/models"
)

const (
	historyFileName = ".workhistory"
)

// GetHistoryPath returns the full path to the work history file
func GetHistoryPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find home directory: %w", err)
	}
	return filepath.Join(home, historyFileName), nil
}

// SaveSession appends a work session to the history file
func SaveSession(session models.WorkSession) error {
	path, err := GetHistoryPath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening history file: %w", err)
	}
	defer file.Close()

	entry := fmt.Sprintf("- %s\n", session.String())
	if _, err := file.WriteString(entry); err != nil {
		return fmt.Errorf("error writing to history file: %w", err)
	}

	return nil
}

// LoadHistory reads and parses all work sessions from the history file
func LoadHistory() ([]models.WorkSession, error) {
	path, err := GetHistoryPath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.WorkSession{}, nil
		}
		return nil, fmt.Errorf("error opening history file: %w", err)
	}
	defer file.Close()

	var sessions []models.WorkSession
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		session, err := parseSessionLine(line)
		if err == nil {
			sessions = append(sessions, session)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading history file: %w", err)
	}

	return sessions, nil
}

// parseSessionLine converts a history file line into a WorkSession struct
func parseSessionLine(line string) (models.WorkSession, error) {
	line = strings.TrimPrefix(line, "- ")
	parts := strings.Split(line, " -> ")
	if len(parts) != 2 {
		return models.WorkSession{}, fmt.Errorf("invalid session format")
	}

	format := "2006/01/02|15:04:05"
	start, err := time.Parse(format, parts[0])
	if err != nil {
		return models.WorkSession{}, err
	}

	end, err := time.Parse(format, parts[1])
	if err != nil {
		return models.WorkSession{}, err
	}

	return models.WorkSession{
		Start:    start,
		End:      end,
		Duration: end.Sub(start),
	}, nil
}