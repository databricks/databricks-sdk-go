package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/logger"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/clusters"
)

func New(cfg *databricks.Config) *CommandsAPI {
	return &CommandsAPI{
		client: client.New(cfg),
	}
}

// CommandsAPI exposes the Command & Context API
type CommandsAPI struct {
	client *client.DatabricksClient
}

// Command is the struct that contains what the 1.2 api returns for the commands api
type Command struct {
	ID      string          `json:"id,omitempty"`
	Status  string          `json:"status,omitempty"`
	Results *CommandResults `json:"results,omitempty"`
}

// Execute creates a spark context and executes a command and then closes context
// Any leading whitespace is trimmed
func (a *CommandsAPI) Execute(ctx context.Context, clusterID, language, commandStr string) CommandResults {
	cluster, err := clusters.New(a.client).GetCluster(ctx, clusters.GetClusterRequest{
		ClusterId: clusterID,
	})
	if err != nil {
		return CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	if !cluster.IsRunningOrResizing() {
		return CommandResults{
			ResultType: "error",
			Summary:    fmt.Sprintf("Cluster %s has to be running or resizing, but is %s", clusterID, cluster.State),
		}
	}
	commandStr = TrimLeadingWhitespace(commandStr)
	logger.Infof("Executing %s command on %s:\n%s", language, clusterID, commandStr)

	// this is the place, where API version propagation through context looks strange,
	// but makes some sense, otherwise there'll have to be full api prefix specification
	// in every client plus SCIM API could not directly be reused in the Accounts API context.
	ctx = context.WithValue(ctx, client.Api, client.API_1_2)

	context, err := a.createContext(ctx, language, clusterID)
	if err != nil {
		return CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	err = a.waitForContextReady(ctx, context, clusterID)
	if err != nil {
		return CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	commandID, err := a.createCommand(ctx, context, clusterID, language, commandStr)
	if err != nil {
		return CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	// TODO: merge getCommand and waitForCommandFinished to "waitForCommandResults"
	err = a.waitForCommandFinished(ctx, commandID, context, clusterID)
	if err != nil {
		return CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	command, err := a.getCommand(ctx, commandID, context, clusterID)
	if err != nil {
		return CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	err = a.deleteContext(ctx, context, clusterID)
	if err != nil {
		return CommandResults{
			ResultType: "error",
			Summary:    err.Error(),
		}
	}
	if command.Results == nil {
		logger.Warnf("Command has no results: %#v", command)
		return CommandResults{
			ResultType: "error",
			Summary:    "Command has no results",
		}
	}
	return *command.Results
}

type genericCommandRequest struct {
	CommandID string `json:"commandId,omitempty" url:"commandId,omitempty"`
	Language  string `json:"language,omitempty" url:"language,omitempty"`
	ClusterID string `json:"clusterId,omitempty" url:"clusterId,omitempty"`
	ContextID string `json:"contextId,omitempty" url:"contextId,omitempty"`
	Command   string `json:"command,omitempty" url:"command,omitempty"`
}

func (a *CommandsAPI) createCommand(ctx context.Context, contextID, clusterID, language, commandStr string) (string, error) {
	var command Command
	err := a.client.Post(ctx, "/api/1.2/commands/execute", genericCommandRequest{
		Language:  language,
		ClusterID: clusterID,
		ContextID: contextID,
		Command:   commandStr,
	}, &command)
	return command.ID, err
}

func (a *CommandsAPI) getCommand(ctx context.Context, commandID, contextID, clusterID string) (Command, error) {
	var commandResp Command
	err := a.client.Get(ctx, "/api/1.2/commands/status", genericCommandRequest{
		CommandID: commandID,
		ContextID: contextID,
		ClusterID: clusterID,
	}, &commandResp)
	return commandResp, err
}

func (a *CommandsAPI) waitForCommandFinished(ctx context.Context, commandID, contextID, clusterID string) error {
	return retries.Wait(ctx, 10*time.Minute, func() *retries.Err {
		commandInfo, err := a.getCommand(ctx, commandID, contextID, clusterID)
		if err != nil {
			return retries.Halt(err)
		}
		switch commandInfo.Status {
		case "Cancelling", "Cancelled", "Error":
			return retries.Halt(fmt.Errorf("Command cannot finish: %s", commandInfo.Status))
		case "Finished":
			return nil
		}
		logger.Debugf("Command is in %s state", commandInfo.Status)
		return retries.Continues(commandInfo.Status)
	})
}

func (a *CommandsAPI) createContext(ctx context.Context, language, clusterID string) (string, error) {
	var context Command // internal hack, yes
	err := a.client.Post(ctx, "/api/1.2/contexts/create", genericCommandRequest{
		Language:  language,
		ClusterID: clusterID,
	}, &context)
	return context.ID, err
}

func (a *CommandsAPI) getContext(ctx context.Context, contextID, clusterID string) (string, error) {
	var contextStatus Command // internal hack, yes
	err := a.client.Get(ctx, "/api/1.2/contexts/status", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, &contextStatus)
	return contextStatus.Status, err
}

func (a *CommandsAPI) deleteContext(ctx context.Context, contextID, clusterID string) error {
	return a.client.Post(ctx, "/api/1.2/contexts/destroy", genericCommandRequest{
		ContextID: contextID,
		ClusterID: clusterID,
	}, nil)
}

func (a *CommandsAPI) waitForContextReady(ctx context.Context, contextID, clusterID string) error {
	return retries.Wait(ctx, 10*time.Minute, func() *retries.Err {
		status, err := a.getContext(ctx, contextID, clusterID)
		if err != nil {
			return retries.Halt(err)
		}
		if status == "Error" {
			return retries.Halt(fmt.Errorf(status))
		}
		if status == "Running" {
			return nil
		}
		return retries.Continues(status)
	})
}
