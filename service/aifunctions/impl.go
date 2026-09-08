// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aifunctions

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AiFunctions API methods
type aiFunctionsImpl struct {
	client *client.DatabricksClient
}

func (a *aiFunctionsImpl) AiClassify(ctx context.Context, request AiClassifyRequest) (*AiClassifyResponse, error) {
	var aiClassifyResponse AiClassifyResponse
	path := "/api/2.0/ai-functions/ai-classify"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &aiClassifyResponse)
	return &aiClassifyResponse, err
}

func (a *aiFunctionsImpl) AiExtract(ctx context.Context, request AiExtractRequest) (*AiExtractResponse, error) {
	var aiExtractResponse AiExtractResponse
	path := "/api/2.0/ai-functions/ai-extract"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &aiExtractResponse)
	return &aiExtractResponse, err
}

func (a *aiFunctionsImpl) AiParseDocument(ctx context.Context, request AiParseDocumentRequest) (*AiParseDocumentResponse, error) {
	var aiParseDocumentResponse AiParseDocumentResponse
	path := "/api/2.0/ai-functions/ai-parse-document"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &aiParseDocumentResponse)
	return &aiParseDocumentResponse, err
}
