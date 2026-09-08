// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sandbox

import (
	"context"
)

// Create, manage, and control the lifecycle of sandboxes -- isolated,
// pre-configured, low-latency Serverless compute environments for running code.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type SandboxService interface {

	// Creates a new Sandbox.
	CreateSandbox(ctx context.Context, request CreateSandboxRequest) (*Sandbox, error)

	// Deletes a Sandbox.
	DeleteSandbox(ctx context.Context, request DeleteSandboxRequest) error

	// Runs a command in the sandbox and blocks until it exits, returning the
	// captured stdout, stderr and exit code in a single response.
	ExecuteCommandSync(ctx context.Context, request ExecuteCommandSyncRequest) (*ExecuteCommandSyncResponse, error)

	// Retrieves a Sandbox by name.
	GetSandbox(ctx context.Context, request GetSandboxRequest) (*Sandbox, error)

	// Lists all Sandboxes.
	ListSandboxes(ctx context.Context, request ListSandboxesRequest) (*ListSandboxesResponse, error)

	// Starts a previously stopped Sandbox under the same sandbox name. Returns
	// NOT_FOUND if there is no stopped sandbox to start for the given name.
	StartSandbox(ctx context.Context, request StartSandboxRequest) (*Sandbox, error)

	// Stops a running Sandbox, preserving it so it can later be restarted with
	// a Start request.
	StopSandbox(ctx context.Context, request StopSandboxRequest) (*Sandbox, error)

	// Updates mutable fields on an existing Sandbox. Allowlisted update_mask
	// paths today: display_name, spec.compute.inactivity_timeout. Returns
	// INVALID_PARAMETER_VALUE for empty masks or unknown paths; NOT_FOUND if
	// the sandbox does not exist.
	UpdateSandbox(ctx context.Context, request UpdateSandboxRequest) (*Sandbox, error)
}
