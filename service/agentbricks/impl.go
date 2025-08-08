// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package agentbricks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/service/agentbricks/agentbrickspb"
)

// unexported type that holds implementations of just AgentBricks API methods
type agentBricksImpl struct {
	client *client.DatabricksClient
}

func (a *agentBricksImpl) CancelOptimize(ctx context.Context, request CancelCustomLlmOptimizationRunRequest) error {
	requestPb, pbErr := CancelCustomLlmOptimizationRunRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize/cancel", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *agentBricksImpl) CreateCustomLlm(ctx context.Context, request CreateCustomLlmRequest) (*CustomLlm, error) {
	requestPb, pbErr := CreateCustomLlmRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customLlmPb agentbrickspb.CustomLlmPb
	path := "/api/2.0/custom-llms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&customLlmPb,
	)
	resp, err := CustomLlmFromPb(&customLlmPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *agentBricksImpl) DeleteCustomLlm(ctx context.Context, request DeleteCustomLlmRequest) error {
	requestPb, pbErr := DeleteCustomLlmRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *agentBricksImpl) GetCustomLlm(ctx context.Context, request GetCustomLlmRequest) (*CustomLlm, error) {
	requestPb, pbErr := GetCustomLlmRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customLlmPb agentbrickspb.CustomLlmPb
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&customLlmPb,
	)
	resp, err := CustomLlmFromPb(&customLlmPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *agentBricksImpl) StartOptimize(ctx context.Context, request StartCustomLlmOptimizationRunRequest) (*CustomLlm, error) {
	requestPb, pbErr := StartCustomLlmOptimizationRunRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customLlmPb agentbrickspb.CustomLlmPb
	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&customLlmPb,
	)
	resp, err := CustomLlmFromPb(&customLlmPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *agentBricksImpl) UpdateCustomLlm(ctx context.Context, request UpdateCustomLlmRequest) (*CustomLlm, error) {
	requestPb, pbErr := UpdateCustomLlmRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customLlmPb agentbrickspb.CustomLlmPb
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&customLlmPb,
	)
	resp, err := CustomLlmFromPb(&customLlmPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
