// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sandbox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Sandbox API methods
type sandboxImpl struct {
	client *client.DatabricksClient
}

func (a *sandboxImpl) CreateSandbox(ctx context.Context, request CreateSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
	path := "/api/2.0/sandboxes"
	queryParams := make(map[string]any)

	if request.SandboxId != "" {
		queryParams["sandbox_id"] = request.SandboxId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Sandbox, &sandbox)
	return &sandbox, err
}

func (a *sandboxImpl) DeleteSandbox(ctx context.Context, request DeleteSandboxRequest) error {
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *sandboxImpl) ExecuteCommandSync(ctx context.Context, request ExecuteCommandSyncRequest) (*ExecuteCommandSyncResponse, error) {
	var executeCommandSyncResponse ExecuteCommandSyncResponse
	path := fmt.Sprintf("/api/2.0/sandbox-exec/%v/exec-sync", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &executeCommandSyncResponse)
	return &executeCommandSyncResponse, err
}

func (a *sandboxImpl) GetSandbox(ctx context.Context, request GetSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &sandbox)
	return &sandbox, err
}

// Lists all Sandboxes.
func (a *sandboxImpl) ListSandboxes(ctx context.Context, request ListSandboxesRequest) listing.Iterator[Sandbox] {

	getNextPage := func(ctx context.Context, req ListSandboxesRequest) (*ListSandboxesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSandboxes(ctx, req)
	}
	getItems := func(resp *ListSandboxesResponse) []Sandbox {
		return resp.Sandboxes
	}
	getNextReq := func(resp *ListSandboxesResponse) *ListSandboxesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists all Sandboxes.
func (a *sandboxImpl) ListSandboxesAll(ctx context.Context, request ListSandboxesRequest) ([]Sandbox, error) {
	iterator := a.ListSandboxes(ctx, request)
	return listing.ToSlice[Sandbox](ctx, iterator)
}

func (a *sandboxImpl) internalListSandboxes(ctx context.Context, request ListSandboxesRequest) (*ListSandboxesResponse, error) {
	var listSandboxesResponse ListSandboxesResponse
	path := "/api/2.0/sandboxes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSandboxesResponse)
	return &listSandboxesResponse, err
}

func (a *sandboxImpl) StartSandbox(ctx context.Context, request StartSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
	path := fmt.Sprintf("/api/2.0/%v/start", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &sandbox)
	return &sandbox, err
}

func (a *sandboxImpl) StopSandbox(ctx context.Context, request StopSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
	path := fmt.Sprintf("/api/2.0/%v/stop", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &sandbox)
	return &sandbox, err
}

func (a *sandboxImpl) UpdateSandbox(ctx context.Context, request UpdateSandboxRequest) (*Sandbox, error) {
	var sandbox Sandbox
	path := fmt.Sprintf("/api/2.0/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Sandbox, &sandbox)
	return &sandbox, err
}
