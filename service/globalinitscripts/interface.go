// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package globalinitscripts

import (
	"context"
)

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
type GlobalInitScriptsService interface {

	// Create init script.
	//
	// Creates a new global init script in this workspace.
	Create(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateResponse, error)

	// Delete init script.
	//
	// Deletes a global init script.
	Delete(ctx context.Context, request Delete) error

	// Get an init script.
	//
	// Gets all the details of a script, including its Base64-encoded contents.
	Get(ctx context.Context, request Get) (*GlobalInitScriptDetailsWithContent, error)

	// Get init scripts.
	//
	// "Get a list of all global init scripts for this workspace. This returns
	// all properties for each script but **not** the script contents. To
	// retrieve the contents of a script, use the [get a global init
	// script](#operation/get-script) operation.
	//
	// Use ListAll() to get all GlobalInitScriptDetails instances
	List(ctx context.Context) (*ListGlobalInitScriptsResponse, error)

	// Update init script.
	//
	// Updates a global init script, specifying only the fields to change. All
	// fields are optional. Unspecified fields retain their current value.
	Update(ctx context.Context, request GlobalInitScriptUpdateRequest) error
}
