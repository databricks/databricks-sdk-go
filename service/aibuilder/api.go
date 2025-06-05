// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// The Custom LLMs service manages state and powers the UI for the Custom LLM product.
package aibuilder

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type CustomLlmsInterface interface {

	// Cancel a Custom LLM Optimization Run.
	Cancel(ctx context.Context, request CancelCustomLlmOptimizationRunRequest) error

	// Start a Custom LLM Optimization Run.
	Create(ctx context.Context, request StartCustomLlmOptimizationRunRequest) (*CustomLlm, error)

	// Get a Custom LLM.
	Get(ctx context.Context, request GetCustomLlmRequest) (*CustomLlm, error)

	// Get a Custom LLM.
	GetById(ctx context.Context, id string) (*CustomLlm, error)

	// Update a Custom LLM.
	Update(ctx context.Context, request UpdateCustomLlmRequest) (*CustomLlm, error)
}

func NewCustomLlms(client *client.DatabricksClient) *CustomLlmsAPI {
	return &CustomLlmsAPI{
		customLlmsImpl: customLlmsImpl{
			client: client,
		},
	}
}

// The Custom LLMs service manages state and powers the UI for the Custom LLM
// product.
type CustomLlmsAPI struct {
	customLlmsImpl
}

// Get a Custom LLM.
func (a *CustomLlmsAPI) GetById(ctx context.Context, id string) (*CustomLlm, error) {
	return a.customLlmsImpl.Get(ctx, GetCustomLlmRequest{
		Id: id,
	})
}
