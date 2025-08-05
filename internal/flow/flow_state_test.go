package flow

import (
	"testing"
)

func TestFlowStateInit(t *testing.T) {
	flow := new(FlowState)
	// Test correct initialization
	// The duration should be zero at the start
	if flow.Duration != 0 {
		t.Errorf("Expected initial duration to be 0, got %v", flow.Duration)
	}
	// FlowName should be empty at the start
	if flow.FlowName != "" {
		t.Errorf("Expected initial flow name to be empty, got %s", flow.FlowName)
	}
	// StartTime should be zero at the start
	if !flow.StartTime.IsZero() {
		t.Errorf("Expected initial start time to be zero, got %v", flow.StartTime)
	}
}
