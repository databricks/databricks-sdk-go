// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingsv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AccountSettingsV2 API methods
type accountSettingsV2Impl struct {
	client *client.DatabricksClient
}

func (a *accountSettingsV2Impl) GetPublicAccountSetting(ctx context.Context, request GetPublicAccountSettingRequest) (*Setting, error) {
	var setting Setting
	path := fmt.Sprintf("/api/2.1/accounts/%v/settings/%v", a.client.ConfiguredAccountID(), request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &setting)
	return &setting, err
}

func (a *accountSettingsV2Impl) PatchPublicAccountSetting(ctx context.Context, request PatchPublicAccountSettingRequest) (*Setting, error) {
	var setting Setting
	path := fmt.Sprintf("/api/2.0/accounts/%v/settings/%v", a.client.ConfiguredAccountID(), request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Setting, &setting)
	return &setting, err
}

// unexported type that holds implementations of just WorkspaceSettingsV2 API methods
type workspaceSettingsV2Impl struct {
	client *client.DatabricksClient
}

func (a *workspaceSettingsV2Impl) GetPublicWorkspaceSetting(ctx context.Context, request GetPublicWorkspaceSettingRequest) (*Setting, error) {
	var setting Setting
	path := fmt.Sprintf("/api/2.1/settings/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &setting)
	return &setting, err
}

func (a *workspaceSettingsV2Impl) PatchPublicWorkspaceSetting(ctx context.Context, request PatchPublicWorkspaceSettingRequest) (*Setting, error) {
	var setting Setting
	path := fmt.Sprintf("/api/2.1/settings/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Setting, &setting)
	return &setting, err
}
