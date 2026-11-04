// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Lakewatch API methods
type lakewatchImpl struct {
	client *client.DatabricksClient
}

func (a *lakewatchImpl) CreateDatasource(ctx context.Context, request CreateDatasourceRequest) (*Datasource, error) {
	var datasource Datasource
	path := "/api/2.0/lakewatch/datasources"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Datasource, &datasource)
	return &datasource, err
}

func (a *lakewatchImpl) CreateRule(ctx context.Context, request CreateRuleRequest) (*Rule, error) {
	var rule Rule
	path := "/api/2.0/lakewatch/rules"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Rule, &rule)
	return &rule, err
}

func (a *lakewatchImpl) DeleteDatasource(ctx context.Context, request DeleteDatasourceRequest) error {
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
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

func (a *lakewatchImpl) DeleteRule(ctx context.Context, request DeleteRuleRequest) error {
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
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

func (a *lakewatchImpl) GetAlertConfig(ctx context.Context, request GetAlertConfigRequest) (*AlertConfig, error) {
	var alertConfig AlertConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &alertConfig)
	return &alertConfig, err
}

func (a *lakewatchImpl) GetCommonConfig(ctx context.Context, request GetCommonConfigRequest) (*CommonConfig, error) {
	var commonConfig CommonConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &commonConfig)
	return &commonConfig, err
}

func (a *lakewatchImpl) GetDatasource(ctx context.Context, request GetDatasourceRequest) (*Datasource, error) {
	var datasource Datasource
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &datasource)
	return &datasource, err
}

func (a *lakewatchImpl) GetDatasourceConfig(ctx context.Context, request GetDatasourceConfigRequest) (*DatasourceConfig, error) {
	var datasourceConfig DatasourceConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &datasourceConfig)
	return &datasourceConfig, err
}

func (a *lakewatchImpl) GetRule(ctx context.Context, request GetRuleRequest) (*Rule, error) {
	var rule Rule
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &rule)
	return &rule, err
}

func (a *lakewatchImpl) GetRuleConfig(ctx context.Context, request GetRuleConfigRequest) (*RuleConfig, error) {
	var ruleConfig RuleConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &ruleConfig)
	return &ruleConfig, err
}

// Lists the datasources in a Unity Catalog catalog, optionally narrowed to one
// schema or filtered by name. Results are paginated. Datasources that are still
// being created or deleted are not listed.
func (a *lakewatchImpl) ListDatasources(ctx context.Context, request ListDatasourcesRequest) listing.Iterator[Datasource] {

	getNextPage := func(ctx context.Context, req ListDatasourcesRequest) (*ListDatasourcesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDatasources(ctx, req)
	}
	getItems := func(resp *ListDatasourcesResponse) []Datasource {
		return resp.Datasources
	}
	getNextReq := func(resp *ListDatasourcesResponse) *ListDatasourcesRequest {
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

// Lists the datasources in a Unity Catalog catalog, optionally narrowed to one
// schema or filtered by name. Results are paginated. Datasources that are still
// being created or deleted are not listed.
func (a *lakewatchImpl) ListDatasourcesAll(ctx context.Context, request ListDatasourcesRequest) ([]Datasource, error) {
	iterator := a.ListDatasources(ctx, request)
	return listing.ToSlice[Datasource](ctx, iterator)
}

func (a *lakewatchImpl) internalListDatasources(ctx context.Context, request ListDatasourcesRequest) (*ListDatasourcesResponse, error) {
	var listDatasourcesResponse ListDatasourcesResponse
	path := "/api/2.0/lakewatch/datasources"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDatasourcesResponse)
	return &listDatasourcesResponse, err
}

// Lists rules, with optional filtering and pagination.
func (a *lakewatchImpl) ListRules(ctx context.Context, request ListRulesRequest) listing.Iterator[Rule] {

	getNextPage := func(ctx context.Context, req ListRulesRequest) (*ListRulesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListRules(ctx, req)
	}
	getItems := func(resp *ListRulesResponse) []Rule {
		return resp.Rules
	}
	getNextReq := func(resp *ListRulesResponse) *ListRulesRequest {
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

// Lists rules, with optional filtering and pagination.
func (a *lakewatchImpl) ListRulesAll(ctx context.Context, request ListRulesRequest) ([]Rule, error) {
	iterator := a.ListRules(ctx, request)
	return listing.ToSlice[Rule](ctx, iterator)
}

func (a *lakewatchImpl) internalListRules(ctx context.Context, request ListRulesRequest) (*ListRulesResponse, error) {
	var listRulesResponse ListRulesResponse
	path := "/api/2.0/lakewatch/rules"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRulesResponse)
	return &listRulesResponse, err
}

func (a *lakewatchImpl) UpdateAlertConfig(ctx context.Context, request UpdateAlertConfigRequest) (*AlertConfig, error) {
	var alertConfig AlertConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.AlertConfig, &alertConfig)
	return &alertConfig, err
}

func (a *lakewatchImpl) UpdateCommonConfig(ctx context.Context, request UpdateCommonConfigRequest) (*CommonConfig, error) {
	var commonConfig CommonConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.CommonConfig, &commonConfig)
	return &commonConfig, err
}

func (a *lakewatchImpl) UpdateDatasource(ctx context.Context, request UpdateDatasourceRequest) (*Datasource, error) {
	var datasource Datasource
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Datasource, &datasource)
	return &datasource, err
}

func (a *lakewatchImpl) UpdateDatasourceConfig(ctx context.Context, request UpdateDatasourceConfigRequest) (*DatasourceConfig, error) {
	var datasourceConfig DatasourceConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.DatasourceConfig, &datasourceConfig)
	return &datasourceConfig, err
}

func (a *lakewatchImpl) UpdateRule(ctx context.Context, request UpdateRuleRequest) (*Rule, error) {
	var rule Rule
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Rule, &rule)
	return &rule, err
}

func (a *lakewatchImpl) UpdateRuleConfig(ctx context.Context, request UpdateRuleConfigRequest) (*RuleConfig, error) {
	var ruleConfig RuleConfig
	path := fmt.Sprintf("/api/2.0/lakewatch/%v", request.Name)
	queryParams := make(map[string]any)

	if request.Etag != "" || slices.Contains(request.ForceSendFields, "Etag") {
		queryParams["etag"] = request.Etag
	}

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
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.RuleConfig, &ruleConfig)
	return &ruleConfig, err
}
