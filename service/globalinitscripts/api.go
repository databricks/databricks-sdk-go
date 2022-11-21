// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewGlobalInitScripts(client *client.DatabricksClient) *GlobalInitScriptsAPI {
	return &GlobalInitScriptsAPI{
		GlobalInitScriptsService: &globalInitScriptsAPI{
			client: client,
		},
	}
}

// The Global Init Scripts API enables Workspace administrators to configure
// global initialization scripts for their workspace. These scripts run on every
// node in every cluster in the workspace.
//
// **Important:** Existing clusters must be restarted to pick up any changes
// made to global init scripts. Global init scripts are run in order. If the
// init script returns with a bad exit code, the Apache Spark container fails to
// launch and init scripts with later position are skipped. If enough containers
// fail, the entire cluster fails with a `GLOBAL_INIT_SCRIPT_FAILURE` error
// code.
type GlobalInitScriptsAPI struct {
	// GlobalInitScriptsService contains low-level REST API interface.
	GlobalInitScriptsService
}

// Create init script
//
// Creates a new global init script in this workspace.
func (a *GlobalInitScriptsAPI) CreateScript(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateScriptResponse, error) {
	return a.GlobalInitScriptsService.CreateScript(ctx, request)
}

// Delete init script
//
// Deletes a global init script.
func (a *GlobalInitScriptsAPI) DeleteScript(ctx context.Context, request DeleteScriptRequest) error {
	return a.GlobalInitScriptsService.DeleteScript(ctx, request)
}

// Delete init script
//
// Deletes a global init script.
func (a *GlobalInitScriptsAPI) DeleteScriptByScriptId(ctx context.Context, scriptId string) error {
	return a.DeleteScript(ctx, DeleteScriptRequest{
		ScriptId: scriptId,
	})
}

// Get an init script
//
// Gets all the details of a script, including its Base64-encoded contents.
func (a *GlobalInitScriptsAPI) GetScript(ctx context.Context, request GetScriptRequest) (*GlobalInitScriptDetailsWithContent, error) {
	return a.GlobalInitScriptsService.GetScript(ctx, request)
}

// Get an init script
//
// Gets all the details of a script, including its Base64-encoded contents.
func (a *GlobalInitScriptsAPI) GetScriptByScriptId(ctx context.Context, scriptId string) (*GlobalInitScriptDetailsWithContent, error) {
	return a.GetScript(ctx, GetScriptRequest{
		ScriptId: scriptId,
	})
}

// Get init scripts
//
// "Get a list of all global init scripts for this workspace. This returns all
// properties for each script but **not** the script contents. To retrieve the
// contents of a script, use the [get a global init
// script](#operation/get-script) operation.
func (a *GlobalInitScriptsAPI) ListScripts(ctx context.Context) ([]GlobalInitScriptDetails, error) {
	return a.GlobalInitScriptsService.ListScripts(ctx)
}

// Update init script
//
// Updates a global init script, specifying only the fields to change. All
// fields are optional. Unspecified fields retain their current value.
func (a *GlobalInitScriptsAPI) UpdateScript(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	return a.GlobalInitScriptsService.UpdateScript(ctx, request)
}

// unexported type that holds implementations of just GlobalInitScripts API methods
type globalInitScriptsAPI struct {
	client *client.DatabricksClient
}

func (a *globalInitScriptsAPI) CreateScript(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateScriptResponse, error) {
	var createScriptResponse CreateScriptResponse
	path := "/api/2.0/global-init-scripts"
	err := a.client.Post(ctx, path, request, &createScriptResponse)
	return &createScriptResponse, err
}

func (a *globalInitScriptsAPI) DeleteScript(ctx context.Context, request DeleteScriptRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *globalInitScriptsAPI) GetScript(ctx context.Context, request GetScriptRequest) (*GlobalInitScriptDetailsWithContent, error) {
	var globalInitScriptDetailsWithContent GlobalInitScriptDetailsWithContent
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Get(ctx, path, request, &globalInitScriptDetailsWithContent)
	return &globalInitScriptDetailsWithContent, err
}

func (a *globalInitScriptsAPI) ListScripts(ctx context.Context) ([]GlobalInitScriptDetails, error) {
	var globalInitScriptDetailsList []GlobalInitScriptDetails
	path := "/api/2.0/global-init-scripts"
	err := a.client.Get(ctx, path, nil, &globalInitScriptDetailsList)
	return globalInitScriptDetailsList, err
}

func (a *globalInitScriptsAPI) UpdateScript(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Patch(ctx, path, request)
	return err
}
