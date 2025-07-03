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

	requestPb, pbErr := cancelCustomLlmOptimizationRunRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize/cancel", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *aiBuilderImpl) CreateCustomLlm(ctx context.Context, request CreateCustomLlmRequest) (*CustomLlm, error) {

	requestPb, pbErr := createCustomLlmRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customLlmPb customLlmPb
	path := "/api/2.0/custom-llms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&customLlmPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := customLlmFromPb(&customLlmPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aiBuilderImpl) DeleteCustomLlm(ctx context.Context, request DeleteCustomLlmRequest) error {

	requestPb, pbErr := deleteCustomLlmRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/custom-lms/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *aiBuilderImpl) GetCustomLlm(ctx context.Context, request GetCustomLlmRequest) (*CustomLlm, error) {

	requestPb, pbErr := getCustomLlmRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customLlmPb customLlmPb
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&customLlmPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := customLlmFromPb(&customLlmPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aiBuilderImpl) StartOptimize(ctx context.Context, request StartCustomLlmOptimizationRunRequest) (*CustomLlm, error) {

	requestPb, pbErr := startCustomLlmOptimizationRunRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customLlmPb customLlmPb
	path := fmt.Sprintf("/api/2.0/custom-llms/%v/optimize", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&customLlmPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := customLlmFromPb(&customLlmPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *aiBuilderImpl) UpdateCustomLlm(ctx context.Context, request UpdateCustomLlmRequest) (*CustomLlm, error) {

	requestPb, pbErr := updateCustomLlmRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customLlmPb customLlmPb
	path := fmt.Sprintf("/api/2.0/custom-llms/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&customLlmPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := customLlmFromPb(&customLlmPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
