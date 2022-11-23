// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just GlobalInitScripts API methods
type globalInitScriptsImpl struct {
	client *client.DatabricksClient
}

func (a *globalInitScriptsImpl) CreateScript(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateScriptResponse, error) {
	var createScriptResponse CreateScriptResponse
	path := "/api/2.0/global-init-scripts"
	err := a.client.Post(ctx, path, request, &createScriptResponse)
	return &createScriptResponse, err
}

func (a *globalInitScriptsImpl) DeleteScript(ctx context.Context, request DeleteScriptRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *globalInitScriptsImpl) GetScript(ctx context.Context, request GetScriptRequest) (*GlobalInitScriptDetailsWithContent, error) {
	var globalInitScriptDetailsWithContent GlobalInitScriptDetailsWithContent
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Get(ctx, path, request, &globalInitScriptDetailsWithContent)
	return &globalInitScriptDetailsWithContent, err
}

func (a *globalInitScriptsImpl) ListScripts(ctx context.Context) ([]GlobalInitScriptDetails, error) {
	var globalInitScriptDetailsList []GlobalInitScriptDetails
	path := "/api/2.0/global-init-scripts"
	err := a.client.Get(ctx, path, nil, &globalInitScriptDetailsList)
	return globalInitScriptDetailsList, err
}

func (a *globalInitScriptsImpl) UpdateScript(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Patch(ctx, path, request)
	return err
}
