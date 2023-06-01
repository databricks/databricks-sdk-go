package compute

import "sync"

type CommandExecution struct {
	clustersAPI         *ClustersAPI
	commandExecutionAPI *CommandExecutionAPI
	language            Language
	clusterID           string

	mu sync.Mutex
}
