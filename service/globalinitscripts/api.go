// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewGlobalInitScripts(client *client.DatabricksClient) GlobalInitScriptsService {
	return &GlobalInitScriptsAPI{
		client: client,
	}
}

type GlobalInitScriptsAPI struct {
	client *client.DatabricksClient
}

// Create a new global init script in this workspace.
func (a *GlobalInitScriptsAPI) CreateScript(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateScriptResponse, error) {
	var createScriptResponse CreateScriptResponse
	path := "/api/2.0/global-init-scripts"
	err := a.client.Post(ctx, path, request, &createScriptResponse)
	return &createScriptResponse, err
}

// Delete a global init script.
func (a *GlobalInitScriptsAPI) DeleteScript(ctx context.Context, request DeleteScriptRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get a list of all global init scripts for this workspace. This returns all
// properties for each script but **not** the script contents. To retrieve the
// contents of a script, use the [get a global init
// script](#operation/get-script) operation.
func (a *GlobalInitScriptsAPI) GetAllScripts(ctx context.Context) (*GetAllScriptsResponse, error) {
	var getAllScriptsResponse GetAllScriptsResponse
	path := "/api/2.0/global-init-scripts"
	err := a.client.Get(ctx, path, nil, &getAllScriptsResponse)
	return &getAllScriptsResponse, err
}

// Get all the details of a script, including its Base64-encoded contents.
func (a *GlobalInitScriptsAPI) GetScript(ctx context.Context, request GetScriptRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Get(ctx, path, request, nil)
	return err
}

// Update a global init script, specifying only the fields to change. All fields
// are optional. Unspecified fields retain their current value.
func (a *GlobalInitScriptsAPI) UpdateScript(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Patch(ctx, path, request)
	return err
}


func (a *GlobalInitScriptsAPI) GetScriptByScriptId(ctx context.Context, scriptId string) error {
	return a.GetScript(ctx, GetScriptRequest{
		ScriptId: scriptId,
	})
}

func (a *GlobalInitScriptsAPI) DeleteScriptByScriptId(ctx context.Context, scriptId string) error {
	return a.DeleteScript(ctx, DeleteScriptRequest{
		ScriptId: scriptId,
	})
}
