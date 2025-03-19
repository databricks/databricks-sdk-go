// TODO : Add the missing methods and implement the methods
// This file has not been completely shifted from the SDK-Beta
// as we still don't have the wait for state methods in the SDK-mod
package compute

import (
	"context"
)

type CommandExecutorV2 struct {
	executionAPI *CommandExecutionAPI
	clusterID    string
	contextID    string
}

type commandExecutionAPIUtilities interface {
	// Start(ctx context.Context, clusterID string, language Language) (*CommandExecutorV2, error)
}

// Start the command execution context on a cluster and ensure it transitions to a running state
func (c *CommandExecutorV2) Destroy(ctx context.Context) error {
	_, err := c.executionAPI.Destroy(ctx, DestroyContext{
		ClusterId: c.clusterID,
		ContextId: c.contextID,
	})
	return err
}

// CommandExecutor creates a spark context and executes a command and then closes context
type CommandExecutor interface {
	Execute(ctx context.Context, clusterID, language, commandStr string) Results
}

// CommandMock mocks the execution of command
type CommandMock func(commandStr string) Results

func (m CommandMock) Execute(_ context.Context, _, _, commandStr string) Results {
	return m(commandStr)
}

// CommandsHighLevelAPI exposes more friendly wrapper over command execution
type CommandsHighLevelAPI struct {
}
