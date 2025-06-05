// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aibuilder

import (
	"context"
)

// The Custom LLMs service manages state and powers the UI for the Custom LLM
// product.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CustomLlmsService interface {

	// Cancel a Custom LLM Optimization Run.
	Cancel(ctx context.Context, request CancelCustomLlmOptimizationRunRequest) error

	// Start a Custom LLM Optimization Run.
	Create(ctx context.Context, request StartCustomLlmOptimizationRunRequest) (*CustomLlm, error)

	// Get a Custom LLM.
	Get(ctx context.Context, request GetCustomLlmRequest) (*CustomLlm, error)

	// Update a Custom LLM.
	Update(ctx context.Context, request UpdateCustomLlmRequest) (*CustomLlm, error)
}
