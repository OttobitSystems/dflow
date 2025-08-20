// Package recap creates a recap data for flow
package recap

import (
	"dflow/internal/persistency/models"
	"dflow/internal/persistency/repository"
)

func Calculate() Recap {
	flows := repository.GetFlowsAndSessions()

	return InitRecap(flows)
}

func CalculateWithFlow(flow models.Flow) Recap {
	return InitRecap([]models.Flow{flow})
}
