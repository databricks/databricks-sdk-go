// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package commands

import (
	"context"
	"net/http"
)

type databricksClient interface {
	Do(ctx context.Context, method string, path string, request any, response any) error
	ConfiguredAccountID() string
	IsAws() bool
}

// unexported type that holds implementations of just CommandExecution API methods
type commandExecutionImpl struct {
	client databricksClient
}

func (a *commandExecutionImpl) Cancel(ctx context.Context, request CancelCommand) error {
	path := "/api/1.2/commands/cancel"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *commandExecutionImpl) CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error) {
	var commandStatusResponse CommandStatusResponse
	path := "/api/1.2/commands/status"
	err := a.client.Do(ctx, http.MethodGet, path, request, &commandStatusResponse)
	return &commandStatusResponse, err
}

func (a *commandExecutionImpl) ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error) {
	var contextStatusResponse ContextStatusResponse
	path := "/api/1.2/contexts/status"
	err := a.client.Do(ctx, http.MethodGet, path, request, &contextStatusResponse)
	return &contextStatusResponse, err
}

func (a *commandExecutionImpl) Create(ctx context.Context, request CreateContext) (*Created, error) {
	var created Created
	path := "/api/1.2/contexts/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &created)
	return &created, err
}

func (a *commandExecutionImpl) Destroy(ctx context.Context, request DestroyContext) error {
	path := "/api/1.2/contexts/destroy"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *commandExecutionImpl) Execute(ctx context.Context, request Command) (*Created, error) {
	var created Created
	path := "/api/1.2/commands/execute"
	err := a.client.Do(ctx, http.MethodPost, path, request, &created)
	return &created, err
}
