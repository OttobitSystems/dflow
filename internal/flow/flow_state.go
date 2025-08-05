package flow

import "time"

type FlowState struct {
	FlowName  string
	StartTime time.Time
	Duration  time.Duration
}

func (flowState *FlowState) Start() {
	// Logic to start the flow state
	flowState.StartTime = time.Now()
}

func (flowState *FlowState) End() {
	// Logic to end the flow state
	flowState.Duration = time.Since(flowState.StartTime)
}

func (flowState *FlowState) CheckDuration() int {
	// Compare current time with start time
	if flowState.StartTime.IsZero() {
		return 0 // Flow has not started yet
	}

	return int(time.Since(flowState.StartTime).Seconds())
}
