// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aibuilder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just CustomLlms API methods
type customLlmsImpl struct {
	client *client.DatabricksClient
}

func (a *customLlmsImpl) Cancel(ctx context.Context, request CancelCustomLlmOptimizationRunRequest) error {
	var cancelResponse CancelResponse
	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize/cancel", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelResponse)
	return err
}

func (a *customLlmsImpl) Create(ctx context.Context, request StartCustomLlmOptimizationRunRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}

func (a *customLlmsImpl) Get(ctx context.Context, request GetCustomLlmRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}

func (a *customLlmsImpl) Update(ctx context.Context, request UpdateCustomLlmRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}
