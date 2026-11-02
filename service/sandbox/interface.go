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

	// Creates a directory in a sandbox, including missing parents by default.
	CreateDirectory(ctx context.Context, request CreateDirectoryRequest) (*CreateDirectoryResponse, error)

	// Creates a new Sandbox.
	CreateSandbox(ctx context.Context, request CreateSandboxRequest) (*Sandbox, error)

	// Deletes a file, symlink, or empty directory, with an option for recursive
	// deletion. Succeeds if the path is already absent.
	DeletePath(ctx context.Context, request DeletePathRequest) (*DeletePathResponse, error)

	// Deletes a Sandbox.
	DeleteSandbox(ctx context.Context, request DeleteSandboxRequest) error

	// Downloads a regular file from a sandbox as an application/octet-stream
	// body. The sandbox starts automatically if it is stopped. Filesystem
	// access uses the sandbox user's permissions and follows symbolic links.
	//
	// A full download returns 200 OK. A satisfiable single byte range returns
	// 206 Partial Content. Multiple byte ranges are unsupported and return 400.
	// Invalid ranges are ignored; an unsatisfiable range returns 416 with
	// Content-Range describing the file length. Range is ignored for empty
	// files.
	//
	// Conditional requests can return 304 Not Modified or 412 Precondition
	// Failed. If-Range selects a partial response only when its strong entity
	// tag matches; otherwise, the entire file is returned. HTTP-date validators
	// in If-Range cause a full download. Reads are not snapshots: concurrent
	// in-place writes can affect the returned bytes. HEAD is unsupported.
	//
	// Range and conditional options may be supplied as query parameters or HTTP
	// headers. A parameter and one header may both be supplied when their
	// values match exactly. Conflicting values or repeated parameters return
	// 400.
	DownloadFile(ctx context.Context, request DownloadFileRequest) (*DownloadFileResponse, error)

	// Runs a command in the sandbox and blocks until it exits, returning the
	// captured stdout, stderr and exit code in a single response.
	ExecuteCommandSync(ctx context.Context, request ExecuteCommandSyncRequest) (*ExecuteCommandSyncResponse, error)

	// Returns metadata for a path without following its final symlink.
	GetMetadata(ctx context.Context, request GetMetadataRequest) (*GetMetadataResponse, error)

	// Retrieves a Sandbox by name.
	GetSandbox(ctx context.Context, request GetSandboxRequest) (*Sandbox, error)

	// Lists the tracked command executions (running and completed) in a
	// sandbox.
	ListCommands(ctx context.Context, request ListCommandsRequest) (*ListCommandsResponse, error)

	// Lists a directory's immediate children in UTF-8 name order. Later pages
	// may reflect concurrent changes.
	ListDirectoryContents(ctx context.Context, request ListDirectoryContentsRequest) (*ListDirectoryContentsResponse, error)

	// Lists all Sandboxes.
	ListSandboxes(ctx context.Context, request ListSandboxesRequest) (*ListSandboxesResponse, error)

	// Renames a path within a sandbox. The destination must be absent and on
	// the same filesystem.
	RenamePath(ctx context.Context, request RenamePathRequest) (*RenamePathResponse, error)

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

	// Uploads an application/octet-stream body to a sandbox. The sandbox starts
	// automatically if it is stopped. Filesystem access uses the sandbox user's
	// permissions and follows symbolic links.
	//
	// Creates the file or replaces its entire contents, creating parent
	// directories if missing. The completed upload becomes visible atomically.
	// A failure before commit leaves an existing destination unchanged; a lost
	// response after commit does not undo the upload. Completion does not
	// guarantee crash durability.
	//
	// Returns 201 Created for a new file or 204 No Content for a replacement.
	// Partial and conditional uploads are unsupported: Content-Range, If-Match,
	// If-None-Match, and If-Unmodified-Since return 400.
	UploadFile(ctx context.Context, request UploadFileRequest) error
}
