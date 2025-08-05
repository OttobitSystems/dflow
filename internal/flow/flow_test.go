package flow

import (
	"fmt"
	"testing"
	"time"
)

func TestFlowSessionDuration(t *testing.T) {
	session := &Session{
		StartedAt: time.Now(),
	}

	// Simulate some time passing
	time.Sleep(2 * time.Second)

	duration := session.Duration()
	if duration < 2*time.Second {
		t.Errorf("Expected duration to be at least 2 seconds, got %v", duration)
	}

	session.EndedAt = time.Now()
	durationAfterEnd := session.Duration()
	if durationAfterEnd <= duration {
		t.Errorf("Expected duration after ending to be greater than before ending, got %v", durationAfterEnd)
	}
}

func TestFlowSessionIsActive(t *testing.T) {
	session := Session{}
	session.Start()
	if !session.IsActive() {
		t.Error("Expected session to be active after starting, but it is not")
	}
	session.End()
	if session.IsActive() {
		t.Error("Expected session to be inactive after ending, but it is still active")
	}
}

func TestFlowSessionDurationInSeconds(t *testing.T) {
	session := Session{}

	session.Start()
	time.Sleep(1 * time.Second)
	durationInSeconds := session.DurationInSeconds()
	fmt.Println(durationInSeconds)
	if durationInSeconds < 1 {
		t.Errorf("Expected duration in seconds to be at least 3, got %d", durationInSeconds)
	}
	time.Sleep(1 * time.Second)
	session.End()
	durationAfterEndInSeconds := session.DurationInSeconds()
	fmt.Println(durationAfterEndInSeconds)
	if durationAfterEndInSeconds <= durationInSeconds {
		t.Errorf("Expected duration in seconds after ending to be greater than before ending, got %d", durationAfterEndInSeconds)
	}
}

// Test FlowSessionAddLog tests adding a log entry to the session logs.
func TestFlowSessionAddLog(t *testing.T) {
	session := &Session{}
	session.Start()

	logMessage := "This is a test log entry"
	session.AddSessionLog(logMessage)

	if len(session.SessionLogs) == 0 {
		t.Error("Expected session logs to contain at least one entry, but it is empty")
	}

	if session.SessionLogs[0].Log != logMessage {
		t.Errorf("Expected first log entry to be '%s', got '%s'", logMessage, session.SessionLogs[0].Log)
	}

	session.End()
	if len(session.SessionLogs) == 0 {
		t.Error("Expected session logs to still contain entries after ending the session, but it is empty")
	}
}
