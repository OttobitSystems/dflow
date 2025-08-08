// Package recap creates a recap data for flow
package recap

import (
	"dflow/internal/persistency/models"
	"time"
)

type Recap struct {
	FlowsRecap []FlowRecap
}

type FlowRecap struct {
	Name       string
	LastEnter  time.Time
	TimeInFlow time.Duration
}

func InitRecap(flows []models.Flow) Recap {
	recap := &Recap{
		make([]FlowRecap, len(flows)),
	}

	i := 0
	for _, flow := range flows {
		recap.FlowsRecap[i] = InitFlowRecap(flow.Name, flow.Sessions)
		i++
	}

	return *recap
}

func InitFlowRecap(flowName string, sessions []models.Session) FlowRecap {
	if len(sessions) == 0 {
		return FlowRecap{
			Name: flowName,
		}
	}

	flowRecap := &FlowRecap{
		Name:      flowName,
		LastEnter: sessions[len(sessions)-1].StartedAt,
	}

	var totalDuration time.Duration

	for _, session := range sessions {
		if session.CompletedAt.After(session.StartedAt) {
			totalDuration += session.CompletedAt.Sub(session.StartedAt)
		}
	}

	flowRecap.TimeInFlow = totalDuration

	return *flowRecap
}
