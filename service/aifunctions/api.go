// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Transform and enrich data with AI on Databricks.
package aifunctions

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type AiFunctionsInterface interface {

	// Classifies content according to a set of provided labels. For REST API
	// requests, the default rate limit is 1,200 requests per minute per workspace.
	// Contact your Databricks account team to request a higher limit.
	AiClassify(ctx context.Context, request AiClassifyRequest) (*AiClassifyResponse, error)

	// Extracts structured data from text and documents according to a provided
	// schema. For REST API requests, the default rate limit is 120 requests per
	// minute per workspace. Contact your Databricks account team to request a
	// higher limit.
	AiExtract(ctx context.Context, request AiExtractRequest) (*AiExtractResponse, error)

	// Parse structured content from unstructured documents.
	AiParseDocument(ctx context.Context, request AiParseDocumentRequest) (*AiParseDocumentResponse, error)
}

func NewAiFunctions(client *client.DatabricksClient) *AiFunctionsAPI {
	return &AiFunctionsAPI{
		aiFunctionsImpl: aiFunctionsImpl{
			client: client,
		},
	}
}

// Transform and enrich data with AI on Databricks.
type AiFunctionsAPI struct {
	aiFunctionsImpl
}
