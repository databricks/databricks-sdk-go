// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aifunctions

import (
	"context"
)

// Transform and enrich data with AI on Databricks.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AiFunctionsService interface {

	// Classifies content according to a set of provided labels.
	AiClassify(ctx context.Context, request AiClassifyRequest) (*AiClassifyResponse, error)

	// Extracts structured data from text and documents according to a provided
	// schema.
	AiExtract(ctx context.Context, request AiExtractRequest) (*AiExtractResponse, error)

	// Parse structured content from unstructured documents.
	AiParseDocument(ctx context.Context, request AiParseDocumentRequest) (*AiParseDocumentResponse, error)
}
