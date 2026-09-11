// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aifunctions

import (
	"context"
)

// Transform and enrich data with AI on Databricks.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AiFunctionsService interface {

	// Classifies content according to a set of provided labels. For REST API
	// requests, the default rate limit is 1,200 requests per minute per
	// workspace. Contact your Databricks account team to request a higher
	// limit.
	AiClassify(ctx context.Context, request AiClassifyRequest) (*AiClassifyResponse, error)

	// Extracts structured data from text and documents according to a provided
	// schema. For REST API requests, the default rate limit is 120 requests per
	// minute per workspace. Contact your Databricks account team to request a
	// higher limit.
	AiExtract(ctx context.Context, request AiExtractRequest) (*AiExtractResponse, error)

	// Parse structured content from unstructured documents. For REST API
	// requests, the default rate limit is 120 pages per minute per workspace.
	// Contact your Databricks account team to request a higher limit.
	AiParseDocument(ctx context.Context, request AiParseDocumentRequest) (*AiParseDocumentResponse, error)
}
