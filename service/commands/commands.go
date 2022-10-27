package commands

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/logger"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"github.com/databricks/databricks-sdk-go/service/clusters"
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
		clusters:  clusters.NewClusters(client),
		execution: NewCommandExecution(client),
	}
}

// CommandsHighLevelAPI exposes more friendly wrapper over command execution
type CommandsHighLevelAPI struct {
	clusters  clusters.ClustersService
	execution CommandExecutionService
}

// Execute creates a spark context and executes a command and then closes context
// Any leading whitespace is trimmed
func (a *CommandsHighLevelAPI) Execute(ctx context.Context, clusterID, language, commandStr string) Results {
	ctx = useragent.InContext(ctx, "sdk-feature", "command-execution")
	cluster, err := a.clusters.Get(ctx, clusters.GetRequest{
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
	logger.Infof("Executing %s command on %s:\n%s", language, clusterID, commandStr)
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
		logger.Warnf("Command has no results: %#v", command)
		return Results{
			ResultType: "error",
			Summary:    "Command has no results",
		}
	}
	return *command.Results
}
