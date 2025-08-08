// Package recap creates a recap data for flow
package recap

import "dflow/internal/persistency/repository"

func Calculate() Recap {
	flows := repository.GetFlowsAndSessions()

	return InitRecap(flows)
}
