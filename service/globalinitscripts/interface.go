// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

import (
	"context"
)

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type GlobalInitScriptsService interface {
	// Create a new global init script in this workspace.
	CreateScript(ctx context.Context, globalInitScriptCreateRequest GlobalInitScriptCreateRequest) (*CreateScriptResponse, error)

	// Delete a global init script.
	DeleteScript(ctx context.Context, deleteScriptRequest DeleteScriptRequest) error

	DeleteScriptByScriptId(ctx context.Context, scriptId string) error
	// Get a list of all global init scripts for this workspace. This returns
	// all properties for each script but **not** the script contents. To
	// retrieve the contents of a script, use the [get a global init
	// script](#operation/get-script) operation.
	GetAllScripts(ctx context.Context) (*GetAllScriptsResponse, error)

	// Get all the details of a script, including its Base64-encoded contents.
	GetScript(ctx context.Context, getScriptRequest GetScriptRequest) error

	GetScriptByScriptId(ctx context.Context, scriptId string) error
	// Update a global init script, specifying only the fields to change. All
	// fields are optional. Unspecified fields retain their current value.
	UpdateScript(ctx context.Context, globalInitScriptUpdateRequest GlobalInitScriptUpdateRequest) error
}
