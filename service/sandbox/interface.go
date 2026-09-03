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

	// Retrieves a Sandbox by name.
	GetSandbox(ctx context.Context, request GetSandboxRequest) (*Sandbox, error)

	// Lists all Sandboxes.
	ListSandboxes(ctx context.Context, request ListSandboxesRequest) (*ListSandboxesResponse, error)

	// Starts a stopped Sandbox by atomically restoring the TerminatedSandbox
	// tombstone to the active table in PENDING and re-running the provisioning
	// workflow. The provisioning workflow's claimWarmPoolSandbox step re-mints
	// the app_instance_name (deterministic from sandbox_id, so equal to the
	// prior life's name). The tombstone's volume_id is preserved so the new
	// AppInstance binds to the same backing device file. The restored sandbox
	// gets a fresh uid and create_time. Returns NOT_FOUND if no tombstone
	// exists for the given (workspace_id, sandbox_id) — the sandbox may not
	// exist or may currently be active; clients can disambiguate via Get.
	StartSandbox(ctx context.Context, request StartSandboxRequest) (*Sandbox, error)

	// Stops a Sandbox, terminating the sandbox while allowing future use of
	// StartSandbox to re-provision the same Sandbox without re-creating a brand
	// new one. Transitions the active row to TERMINATING with
	// USER_REQUEST_STOP; the termination workflow settles to a
	// TerminatedSandbox tombstone (no row drop) so the sandbox can later be
	// restarted via StartSandbox.
	StopSandbox(ctx context.Context, request StopSandboxRequest) (*Sandbox, error)

	// Updates mutable fields on an existing Sandbox. Allowlisted update_mask
	// paths today: metadata.display_name, spec.compute.inactivity_timeout.
	// Returns INVALID_PARAMETER_VALUE for empty masks or unknown paths;
	// NOT_FOUND if no active row or tombstone exists for the given sandbox.
	// Concurrent-update conflicts surface as ABORTED via EStore OccConflict.
	UpdateSandbox(ctx context.Context, request UpdateSandboxRequest) (*Sandbox, error)
}
