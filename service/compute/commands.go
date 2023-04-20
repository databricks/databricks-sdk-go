package compute

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// CommandExecutor creates a spark context and executes a command and then closes context
type CommandExecutor interface {
	Execute(ctx context.Context, clusterID, language, commandStr string) Results
}

// CommandMock mocks the execution of command
type CommandMock func(commandStr string) Results

func (m CommandMock) Execute(_ context.Context, _, _, commandStr string) Results {
	return m(commandStr)
}

func NewCommandExecutor(client *client.DatabricksClient) CommandExecutor {
	return &CommandsHighLevelAPI{
		clusters:  NewClusters(client),
		execution: NewCommandExecution(client),
	}
}

// CommandsHighLevelAPI exposes more friendly wrapper over command execution
type CommandsHighLevelAPI struct {
	clusters  *ClustersAPI
	execution *CommandExecutionAPI
}

// Execute creates a spark context and executes a command and then closes context
// Any leading whitespace is trimmed
func (a *CommandsHighLevelAPI) Execute(ctx context.Context, clusterID, language, commandStr string) Results {
	ctx = useragent.InContext(ctx, "sdk-feature", "command-execution")
	cluster, err := a.clusters.Get(ctx, GetClusterRequest{
		ClusterId: clusterID,
	})
	if err != nil {
		return Results{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	if !cluster.IsRunningOrResizing() {
		return Results{
			ResultType: "error",
			Summary:    fmt.Sprintf("Cluster %s has to be running or resizing, but is %s", clusterID, cluster.State),
		}
	}
	commandStr = TrimLeadingWhitespace(commandStr)
	logger.Debugf(ctx, "Executing %s command on %s:\n%s", language, clusterID, commandStr)
	context, err := a.execution.CreateAndWait(ctx, CreateContext{
		ClusterId: clusterID,
		Language:  Language(language),
	})
	if err != nil {
		return Results{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	defer a.execution.Destroy(ctx, DestroyContext{
		ClusterId: clusterID,
		ContextId: context.Id,
	})
	command, err := a.execution.ExecuteAndWait(ctx, Command{
		ClusterId: clusterID,
		ContextId: context.Id,
		Language:  Language(language),
		Command:   commandStr,
	})
	if err != nil {
		return Results{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	if command.Results == nil {
		logger.Warnf(ctx, "Command has no results: %#v", command)
		return Results{
			ResultType: "error",
			Summary:    "Command has no results",
		}
	}
	return *command.Results
}
