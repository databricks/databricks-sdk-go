// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewCommandExecution(client *client.DatabricksClient) CommandExecutionService {
	return &CommandExecutionAPI{
		client: client,
	}
}

type CommandExecutionAPI struct {
	client *client.DatabricksClient
}

// Cancel a command
func (a *CommandExecutionAPI) Cancel(ctx context.Context, request CancelCommand) error {
	path := "/api/1.2/commands/cancel"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// CancelTimeout overrides the default timeout of 20 minutes to reach Cancelled state
func CancelTimeout(dur time.Duration) retries.Option[CommandStatusResponse] {
	return retries.Timeout[CommandStatusResponse](dur)
}

// Cancel and wait to reach Cancelled state
func (a *CommandExecutionAPI) CancelAndWait(ctx context.Context, cancelCommand CancelCommand, options ...retries.Option[CommandStatusResponse]) (*CommandStatusResponse, error) {
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
			CommandId: cancelCommand.CommandId,
		})
		for _, o := range options {
			o(&retries.Info[CommandStatusResponse]{
				Info:    *commandStatusResponse,
				Timeout: i.Timeout,
			})
		}
		if err != nil {
			return nil, retries.Halt(err)
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

// Get information about a command
func (a *CommandExecutionAPI) CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error) {
	var commandStatusResponse CommandStatusResponse
	path := "/api/1.2/commands/status"
	err := a.client.Get(ctx, path, request, &commandStatusResponse)
	return &commandStatusResponse, err
}

// Get information about an execution context
func (a *CommandExecutionAPI) ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error) {
	var contextStatusResponse ContextStatusResponse
	path := "/api/1.2/contexts/status"
	err := a.client.Get(ctx, path, request, &contextStatusResponse)
	return &contextStatusResponse, err
}

// Create an execution context
func (a *CommandExecutionAPI) Create(ctx context.Context, request CreateContext) (*Created, error) {
	var created Created
	path := "/api/1.2/contexts/create"
	err := a.client.Post(ctx, path, request, &created)
	return &created, err
}

// CreateTimeout overrides the default timeout of 20 minutes to reach Running state
func CreateTimeout(dur time.Duration) retries.Option[ContextStatusResponse] {
	return retries.Timeout[ContextStatusResponse](dur)
}

// Create and wait to reach Running state
func (a *CommandExecutionAPI) CreateAndWait(ctx context.Context, createContext CreateContext, options ...retries.Option[ContextStatusResponse]) (*ContextStatusResponse, error) {
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
		for _, o := range options {
			o(&retries.Info[ContextStatusResponse]{
				Info:    *contextStatusResponse,
				Timeout: i.Timeout,
			})
		}
		if err != nil {
			return nil, retries.Halt(err)
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
func (a *CommandExecutionAPI) Destroy(ctx context.Context, request DestroyContext) error {
	path := "/api/1.2/contexts/destroy"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Run a command
func (a *CommandExecutionAPI) Execute(ctx context.Context, request Command) (*Created, error) {
	var created Created
	path := "/api/1.2/commands/execute"
	err := a.client.Post(ctx, path, request, &created)
	return &created, err
}

// ExecuteTimeout overrides the default timeout of 20 minutes to reach Finished or Error state
func ExecuteTimeout(dur time.Duration) retries.Option[CommandStatusResponse] {
	return retries.Timeout[CommandStatusResponse](dur)
}

// Execute and wait to reach Finished or Error state
func (a *CommandExecutionAPI) ExecuteAndWait(ctx context.Context, command Command, options ...retries.Option[CommandStatusResponse]) (*CommandStatusResponse, error) {
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
		for _, o := range options {
			o(&retries.Info[CommandStatusResponse]{
				Info:    *commandStatusResponse,
				Timeout: i.Timeout,
			})
		}
		if err != nil {
			return nil, retries.Halt(err)
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
