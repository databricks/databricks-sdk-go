// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

import (
	"context"
	"fmt"
	"net/http"
)

type databricksClient interface {
	Do(ctx context.Context, method string, path string, request any, response any) error
	ConfiguredAccountID() string
	IsAws() bool
}

// unexported type that holds implementations of just GlobalInitScripts API methods
type globalInitScriptsImpl struct {
	client databricksClient
}

func (a *globalInitScriptsImpl) CreateScript(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateScriptResponse, error) {
	var createScriptResponse CreateScriptResponse
	path := "/api/2.0/global-init-scripts"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createScriptResponse)
	return &createScriptResponse, err
}

func (a *globalInitScriptsImpl) DeleteScript(ctx context.Context, request DeleteScriptRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *globalInitScriptsImpl) GetScript(ctx context.Context, request GetScriptRequest) (*GlobalInitScriptDetailsWithContent, error) {
	var globalInitScriptDetailsWithContent GlobalInitScriptDetailsWithContent
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &globalInitScriptDetailsWithContent)
	return &globalInitScriptDetailsWithContent, err
}

func (a *globalInitScriptsImpl) ListScripts(ctx context.Context) ([]GlobalInitScriptDetails, error) {
	var globalInitScriptDetailsList []GlobalInitScriptDetails
	path := "/api/2.0/global-init-scripts"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &globalInitScriptDetailsList)
	return globalInitScriptDetailsList, err
}

func (a *globalInitScriptsImpl) UpdateScript(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}
