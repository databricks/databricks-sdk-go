// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

func NewCommandExecution(client *client.DatabricksClient) *CommandExecutionAPI {
	return &CommandExecutionAPI{
		CommandExecutionService: &commandExecutionAPI{
			client: client,
		},
	}
}

// This API allows execution of Python, Scala, SQL, or R commands on running
// Databricks Clusters.
type CommandExecutionAPI struct {
	// CommandExecutionService contains low-level REST API interface.
	CommandExecutionService
}

// Cancel a command
//
// Cancels a currently running command within an execution context.
//
// The command ID is obtained from a prior successful call to __execute__.
func (a *CommandExecutionAPI) Cancel(ctx context.Context, request CancelCommand) error {
	return a.CommandExecutionService.Cancel(ctx, request)
}

// Calls [CommandExecutionAPI.Cancel] and waits to reach Cancelled state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[CommandStatusResponse](60*time.Minute) functional option.
func (a *CommandExecutionAPI) CancelAndWait(ctx context.Context, cancelCommand CancelCommand, options ...retries.Option[CommandStatusResponse]) (*CommandStatusResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Cancel(ctx, cancelCommand)
	if err != nil {
		return nil, err
	}
	i := retries.Info[CommandStatusResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[CommandStatusResponse](ctx, i.Timeout, func() (*CommandStatusResponse, *retries.Err) {
		commandStatusResponse, err := a.CommandStatus(ctx, CommandStatusRequest{
			ClusterId: cancelCommand.ClusterId,
			CommandId: cancelCommand.CommandId,
			ContextId: cancelCommand.ContextId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[CommandStatusResponse]{
				Info:    *commandStatusResponse,
				Timeout: i.Timeout,
			})
		}
		status := commandStatusResponse.Status
		statusMessage := commandStatusResponse.Results.Cause
		switch status {
		case CommandStatusCancelled: // target state
			return commandStatusResponse, nil
		case CommandStatusError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				CommandStatusCancelled, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Get command info
//
// Gets the status of and, if available, the results from a currently executing
// command.
//
// The command ID is obtained from a prior successful call to __execute__.
func (a *CommandExecutionAPI) CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error) {
	return a.CommandExecutionService.CommandStatus(ctx, request)
}

// Get status
//
// Gets the status for an execution context.
func (a *CommandExecutionAPI) ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error) {
	return a.CommandExecutionService.ContextStatus(ctx, request)
}

// Create an execution context
//
// Creates an execution context for running cluster commands.
//
// If successful, this method returns the ID of the new execution context.
func (a *CommandExecutionAPI) Create(ctx context.Context, request CreateContext) (*Created, error) {
	return a.CommandExecutionService.Create(ctx, request)
}

// Calls [CommandExecutionAPI.Create] and waits to reach Running state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ContextStatusResponse](60*time.Minute) functional option.
func (a *CommandExecutionAPI) CreateAndWait(ctx context.Context, createContext CreateContext, options ...retries.Option[ContextStatusResponse]) (*ContextStatusResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	created, err := a.Create(ctx, createContext)
	if err != nil {
		return nil, err
	}
	i := retries.Info[ContextStatusResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[ContextStatusResponse](ctx, i.Timeout, func() (*ContextStatusResponse, *retries.Err) {
		contextStatusResponse, err := a.ContextStatus(ctx, ContextStatusRequest{
			ClusterId: createContext.ClusterId,
			ContextId: created.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[ContextStatusResponse]{
				Info:    *contextStatusResponse,
				Timeout: i.Timeout,
			})
		}
		status := contextStatusResponse.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case ContextStatusRunning: // target state
			return contextStatusResponse, nil
		case ContextStatusError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ContextStatusRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Delete an execution context
//
// Deletes an execution context.
func (a *CommandExecutionAPI) Destroy(ctx context.Context, request DestroyContext) error {
	return a.CommandExecutionService.Destroy(ctx, request)
}

// Run a command
//
// Runs a cluster command in the given execution context, using the provided
// language.
//
// If successful, it returns an ID for tracking the status of the command's
// execution.
func (a *CommandExecutionAPI) Execute(ctx context.Context, request Command) (*Created, error) {
	return a.CommandExecutionService.Execute(ctx, request)
}

// Calls [CommandExecutionAPI.Execute] and waits to reach Finished or Error state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[CommandStatusResponse](60*time.Minute) functional option.
func (a *CommandExecutionAPI) ExecuteAndWait(ctx context.Context, command Command, options ...retries.Option[CommandStatusResponse]) (*CommandStatusResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	created, err := a.Execute(ctx, command)
	if err != nil {
		return nil, err
	}
	i := retries.Info[CommandStatusResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[CommandStatusResponse](ctx, i.Timeout, func() (*CommandStatusResponse, *retries.Err) {
		commandStatusResponse, err := a.CommandStatus(ctx, CommandStatusRequest{
			ClusterId: command.ClusterId,
			CommandId: created.Id,
			ContextId: command.ContextId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[CommandStatusResponse]{
				Info:    *commandStatusResponse,
				Timeout: i.Timeout,
			})
		}
		status := commandStatusResponse.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case CommandStatusFinished, CommandStatusError: // target state
			return commandStatusResponse, nil
		case CommandStatusCancelled, CommandStatusCancelling:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				CommandStatusFinished, CommandStatusError, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// unexported type that holds implementations of just CommandExecution API methods
type commandExecutionAPI struct {
	client *client.DatabricksClient
}

func (a *commandExecutionAPI) Cancel(ctx context.Context, request CancelCommand) error {
	path := "/api/1.2/commands/cancel"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *commandExecutionAPI) CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error) {
	var commandStatusResponse CommandStatusResponse
	path := "/api/1.2/commands/status"
	err := a.client.Get(ctx, path, request, &commandStatusResponse)
	return &commandStatusResponse, err
}

func (a *commandExecutionAPI) ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error) {
	var contextStatusResponse ContextStatusResponse
	path := "/api/1.2/contexts/status"
	err := a.client.Get(ctx, path, request, &contextStatusResponse)
	return &contextStatusResponse, err
}

func (a *commandExecutionAPI) Create(ctx context.Context, request CreateContext) (*Created, error) {
	var created Created
	path := "/api/1.2/contexts/create"
	err := a.client.Post(ctx, path, request, &created)
	return &created, err
}

func (a *commandExecutionAPI) Destroy(ctx context.Context, request DestroyContext) error {
	path := "/api/1.2/contexts/destroy"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *commandExecutionAPI) Execute(ctx context.Context, request Command) (*Created, error) {
	var created Created
	path := "/api/1.2/commands/execute"
	err := a.client.Post(ctx, path, request, &created)
	return &created, err
}
