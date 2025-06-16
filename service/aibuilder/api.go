// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Custom LLMs service manages state and powers the UI for the Custom LLM product.
package aibuilder

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type AiBuilderInterface interface {

	// Cancel a Custom LLM Optimization Run.
	CancelOptimize(ctx context.Context, request CancelCustomLlmOptimizationRunRequest) error

	// Create a Custom LLM.
	CreateCustomLlm(ctx context.Context, request CreateCustomLlmRequest) (*CustomLlm, error)

	// Delete a Custom LLM.
	DeleteCustomLlm(ctx context.Context, request DeleteCustomLlmRequest) error

	// Delete a Custom LLM.
	DeleteCustomLlmById(ctx context.Context, id string) error

	// Get a Custom LLM.
	GetCustomLlm(ctx context.Context, request GetCustomLlmRequest) (*CustomLlm, error)

	// Get a Custom LLM.
	GetCustomLlmById(ctx context.Context, id string) (*CustomLlm, error)

	// Start a Custom LLM Optimization Run.
	StartOptimize(ctx context.Context, request StartCustomLlmOptimizationRunRequest) (*CustomLlm, error)

	// Update a Custom LLM.
	UpdateCustomLlm(ctx context.Context, request UpdateCustomLlmRequest) (*CustomLlm, error)
}

func NewAiBuilder(client *client.DatabricksClient) *AiBuilderAPI {
	return &AiBuilderAPI{
		aiBuilderImpl: aiBuilderImpl{
			client: client,
		},
	}
}

// The Custom LLMs service manages state and powers the UI for the Custom LLM
// product.
type AiBuilderAPI struct {
	aiBuilderImpl
}

// Delete a Custom LLM.
func (a *AiBuilderAPI) DeleteCustomLlmById(ctx context.Context, id string) error {
	return a.aiBuilderImpl.DeleteCustomLlm(ctx, DeleteCustomLlmRequest{
		Id: id,
	})
}

// Get a Custom LLM.
func (a *AiBuilderAPI) GetCustomLlmById(ctx context.Context, id string) (*CustomLlm, error) {
	return a.aiBuilderImpl.GetCustomLlm(ctx, GetCustomLlmRequest{
		Id: id,
	})
}
