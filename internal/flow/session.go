package flow

import (
	"dflow/internal/persistency/models"
	"dflow/internal/persistency/repository"
	"fmt"
	"time"
)

type Session struct {
	FlowName     string
	StartedAt    time.Time
	EndedAt      time.Time
	Logs         []SessionLog
	Objective    string
	InDatabaseID string
}

func InitSession(flowName string) *Session {
	flowDBID, err := repository.InitSession(flowName)
	if err != nil {
		panic("flow not found!")
	}

	session := &Session{FlowName: flowName, InDatabaseID: flowDBID}
	session.Start()

	return session
}

// Start marks the session as started by setting StartedAt to the current time.
func (s *Session) Start() {
	if !s.StartedAt.IsZero() {
		return // Session has already started
	}
	s.StartedAt = time.Now()

	repository.NotifySessionStarted(s.InDatabaseID, s.StartedAt)
}

// End marks the session as ended by setting EndedAt to the current time.
func (s *Session) End() {
	s.EndedAt = time.Now()
}

// IsActive checks if the session is currently active.
func (s *Session) IsActive() bool {
	return !s.StartedAt.IsZero() && s.EndedAt.IsZero()
}

// Duration returns the duration of the session.
// If the Session has not ended, it returns the time since it started.
// If the Session has ended, it returns the difference between EndedAt and StartedAt.
// If the Session has not started, it returns zero Duration.
func (s *Session) Duration() time.Duration {
	if s.EndedAt.IsZero() {
		return time.Since(s.StartedAt)
	}
	return s.EndedAt.Sub(s.StartedAt)
}

// The DurationInSeconds returns Duration as integer seconds.
func (s *Session) DurationInSeconds() int {
	if s.StartedAt.IsZero() {
		return 0 // Session has not started
	}
	return int(s.Duration().Seconds())
}

func (s *Session) DurationString() string {
	if s.StartedAt.IsZero() {
		return "--:--:--" // Session has not started
	}
	return fmt.Sprintf("%02d h %02d m %02d s", int(s.Duration().Hours()), int(s.Duration().Minutes()), int(s.Duration().Seconds()))
}

// AddSessionLog adds a new log entry to the session logs.
func (s *Session) AddSessionLog(log string) {
	s.Logs = append(s.Logs, NewSessionLog(log))
}

// IsCompleted checks if the Session has ended.
func (s *Session) IsCompleted() bool {
	return !s.StartedAt.IsZero() && !s.EndedAt.IsZero()
}

func (s *Session) StoreLog(logText string) {
	repository.StoreLog(s.InDatabaseID, s.FlowName, logText)
}

func (s *Session) GetLogs() []models.Log {
	return repository.GetAllLastLogs(s.InDatabaseID, s.FlowName)
}
