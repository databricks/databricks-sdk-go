// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just GlobalInitScripts API methods
type globalInitScriptsImpl struct {
	client *client.DatabricksClient
}

func (a *globalInitScriptsImpl) Create(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/global-init-scripts"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createResponse)
	return &createResponse, err
}

func (a *globalInitScriptsImpl) Delete(ctx context.Context, request Delete) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *globalInitScriptsImpl) Get(ctx context.Context, request Get) (*GlobalInitScriptDetailsWithContent, error) {
	var globalInitScriptDetailsWithContent GlobalInitScriptDetailsWithContent
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &globalInitScriptDetailsWithContent)
	return &globalInitScriptDetailsWithContent, err
}

func (a *globalInitScriptsImpl) List(ctx context.Context) (*ListGlobalInitScriptsResponse, error) {
	var listGlobalInitScriptsResponse ListGlobalInitScriptsResponse
	path := "/api/2.0/global-init-scripts"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listGlobalInitScriptsResponse)
	return &listGlobalInitScriptsResponse, err
}

func (a *globalInitScriptsImpl) Update(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}
