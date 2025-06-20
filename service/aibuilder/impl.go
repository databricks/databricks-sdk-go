// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aibuilder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AiBuilder API methods
type aiBuilderImpl struct {
	client *client.DatabricksClient
}

func (a *aiBuilderImpl) CancelOptimize(ctx context.Context, request CancelCustomLlmOptimizationRunRequest) error {
	var cancelOptimizeResponse CancelOptimizeResponse
	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize/cancel", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelOptimizeResponse)
	return err
}

func (a *aiBuilderImpl) CreateCustomLlm(ctx context.Context, request CreateCustomLlmRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := "/api/2.0/custom-llms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}

func (a *aiBuilderImpl) DeleteCustomLlm(ctx context.Context, request DeleteCustomLlmRequest) error {
	var deleteCustomLlmResponse DeleteCustomLlmResponse
	path := fmt.Sprintf("/api/2.0/custom-lms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteCustomLlmResponse)
	return err
}

func (a *aiBuilderImpl) GetCustomLlm(ctx context.Context, request GetCustomLlmRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}

func (a *aiBuilderImpl) StartOptimize(ctx context.Context, request StartCustomLlmOptimizationRunRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}

func (a *aiBuilderImpl) UpdateCustomLlm(ctx context.Context, request UpdateCustomLlmRequest) (*CustomLlm, error) {
	var customLlm CustomLlm
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &customLlm)
	return &customLlm, err
}
