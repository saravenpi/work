package models

import (
	"time"
)

const (
	SessionLengthSeconds = 1500
)

// WorkSession represents a completed work session with start, end, and duration
type WorkSession struct {
	Start    time.Time
	End      time.Time
	Duration time.Duration
}

// NewWorkSession creates a new work session with the given start time and duration
func NewWorkSession(start time.Time, duration time.Duration) WorkSession {
	return WorkSession{
		Start:    start,
		End:      start.Add(duration),
		Duration: duration,
	}
}

// String returns a formatted string representation of the work session
func (s WorkSession) String() string {
	format := "2006/01/02|15:04:05"
	return s.Start.Format(format) + " -> " + s.End.Format(format)
}