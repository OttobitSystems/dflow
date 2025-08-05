package flow

import "time"

type SessionLog struct {
	Log       string
	Timestamp time.Time
}

// NewSessionLog creates a new session log entry with the current timestamp.
func NewSessionLog(log string) SessionLog {
	return SessionLog{
		Log:       log,
		Timestamp: time.Now(),
	}
}
