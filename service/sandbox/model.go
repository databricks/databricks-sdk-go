// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sandbox

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/duration"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type ComputeSpec struct {
	// Idle duration after which the sandbox is automatically terminated.
	InactivityTimeout *duration.Duration `json:"inactivity_timeout,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComputeSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComputeSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateSandboxRequest struct {
	// The sandbox to create.
	Sandbox Sandbox `json:"sandbox"`
	// Client-supplied ID that becomes the final path segment of the resource
	// name.
	SandboxId string `json:"-" url:"sandbox_id"`
}

func (s *CreateSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteSandboxRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *DeleteSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Terminal status of a unary command execution.
type ExecuteCommandStatus string

const ExecuteCommandStatusExecuteCommandStatusCompleted ExecuteCommandStatus = `EXECUTE_COMMAND_STATUS_COMPLETED`

const ExecuteCommandStatusExecuteCommandStatusFailed ExecuteCommandStatus = `EXECUTE_COMMAND_STATUS_FAILED`

const ExecuteCommandStatusExecuteCommandStatusTimedOut ExecuteCommandStatus = `EXECUTE_COMMAND_STATUS_TIMED_OUT`

// String representation for [fmt.Print]
func (f *ExecuteCommandStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExecuteCommandStatus) Set(v string) error {
	switch v {
	case `EXECUTE_COMMAND_STATUS_COMPLETED`, `EXECUTE_COMMAND_STATUS_FAILED`, `EXECUTE_COMMAND_STATUS_TIMED_OUT`:
		*f = ExecuteCommandStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXECUTE_COMMAND_STATUS_COMPLETED", "EXECUTE_COMMAND_STATUS_FAILED", "EXECUTE_COMMAND_STATUS_TIMED_OUT"`, v)
	}
}

// Values returns all possible values for ExecuteCommandStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExecuteCommandStatus) Values() []ExecuteCommandStatus {
	return []ExecuteCommandStatus{
		ExecuteCommandStatusExecuteCommandStatusCompleted,
		ExecuteCommandStatusExecuteCommandStatusFailed,
		ExecuteCommandStatusExecuteCommandStatusTimedOut,
	}
}

// Type always returns ExecuteCommandStatus to satisfy [pflag.Value] interface
func (f *ExecuteCommandStatus) Type() string {
	return "ExecuteCommandStatus"
}

// Request to run a command in the given sandbox and wait for it to finish.
type ExecuteCommandSyncRequest struct {
	// Arguments passed to `cmd`.
	Args []string `json:"args,omitempty"`
	// Executable or command to run (e.g. `/bin/echo`, `python3`).
	Cmd string `json:"cmd"`
	// Extra environment variables for the command's process, merged over the
	// sandbox's default environment.
	Envs map[string]string `json:"envs,omitempty"`
	// Maximum time to wait for the command to finish. When it elapses the
	// command is terminated and the response carries status `TIMED_OUT`. The
	// server applies a default when unset and clamps to an upper bound;
	// negative or otherwise invalid durations are rejected with
	// `INVALID_ARGUMENT`.
	ExecutionTimeout *duration.Duration `json:"execution_timeout,omitempty"`
	// Resource name of the sandbox to run the command in, in the form
	// `sandboxes/{sandbox_id}`. Bound from the URL path.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExecuteCommandSyncRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExecuteCommandSyncRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Result of a completed unary command execution: captured output, exit code,
// and terminal status.
type ExecuteCommandSyncResponse struct {
	// Daemon-generated identifier for this command execution, for correlation
	// (for example in `ListCommands`).
	CommandId string `json:"command_id,omitempty"`
	// Process exit code. Unset when the process was terminated by a signal
	// (e.g. on `TIMED_OUT`) or never started (`FAILED`) rather than exiting
	// normally.
	ExitCode int `json:"exit_code,omitempty"`
	// Terminal status of the command execution. Always set on a successful
	// response; never `EXECUTE_COMMAND_STATUS_UNSPECIFIED`.
	Status ExecuteCommandStatus `json:"status,omitempty"`
	// Captured standard error, with the same UTF-8 semantics as `stdout`.
	Stderr string `json:"stderr,omitempty"`
	// Captured standard output as UTF-8 text. Invalid UTF-8 bytes are replaced
	// with the Unicode replacement character (U+FFFD).
	Stdout string `json:"stdout,omitempty"`
	// True when `stdout` / `stderr` were truncated because the captured output
	// exceeded the server's per-response size cap. The dropped output is not
	// included in this response and is not recoverable through this unary API.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExecuteCommandSyncResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExecuteCommandSyncResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetSandboxRequest struct {
	Name string `json:"-" url:"-"`
}

func (s *GetSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListSandboxesRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSandboxesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSandboxesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of Sandboxes.
type ListSandboxesResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Sandboxes []Sandbox `json:"sandboxes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSandboxesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSandboxesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Sandbox resource representing an execution environment.
type Sandbox struct {
	// Output only. The creation time of the sandbox.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Human-readable display label for the sandbox. At most 256 bytes.
	DisplayName string `json:"display_name,omitempty"`
	// The AIP-compliant resource name, such as "sandboxes/my-sandbox".
	Name string `json:"name,omitempty"`
	// The desired configuration of the sandbox, supplied by the caller at
	// creation time.
	Spec *SandboxSpec `json:"spec,omitempty"`
	// The observed runtime state of the sandbox, populated by the server.
	Status *SandboxStatus `json:"status,omitempty"`
	// Output only. The last update time of the sandbox metadata and spec.
	UpdateTime *time.Time `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Sandbox) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Sandbox) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SandboxSpec struct {
	// Compute configuration (size, inactivity timeout) requested for the
	// sandbox.
	Compute *ComputeSpec `json:"compute,omitempty"`
}

func (s *SandboxSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Lifecycle state of a Sandbox resource. STOPPING is the transient state while
// the sandbox is being stopped -- by a Stop request or inactivity
// auto-termination -- and settles to STOPPED once the operation completes.
type SandboxState string

const SandboxStateSandboxStatePending SandboxState = `SANDBOX_STATE_PENDING`

const SandboxStateSandboxStateRunning SandboxState = `SANDBOX_STATE_RUNNING`

const SandboxStateSandboxStateStopped SandboxState = `SANDBOX_STATE_STOPPED`

const SandboxStateSandboxStateStopping SandboxState = `SANDBOX_STATE_STOPPING`

// String representation for [fmt.Print]
func (f *SandboxState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SandboxState) Set(v string) error {
	switch v {
	case `SANDBOX_STATE_PENDING`, `SANDBOX_STATE_RUNNING`, `SANDBOX_STATE_STOPPED`, `SANDBOX_STATE_STOPPING`:
		*f = SandboxState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SANDBOX_STATE_PENDING", "SANDBOX_STATE_RUNNING", "SANDBOX_STATE_STOPPED", "SANDBOX_STATE_STOPPING"`, v)
	}
}

// Values returns all possible values for SandboxState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SandboxState) Values() []SandboxState {
	return []SandboxState{
		SandboxStateSandboxStatePending,
		SandboxStateSandboxStateRunning,
		SandboxStateSandboxStateStopped,
		SandboxStateSandboxStateStopping,
	}
}

// Type always returns SandboxState to satisfy [pflag.Value] interface
func (f *SandboxState) Type() string {
	return "SandboxState"
}

type SandboxStatus struct {
	// Lifecycle state of the sandbox.
	State SandboxState `json:"state,omitempty"`
}

func (s *SandboxStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A request to start a Sandbox.
type StartSandboxRequest struct {
	// Resource name of the sandbox to start, in the form
	// `sandboxes/{sandbox_id}`.
	Name string `json:"-" url:"-"`
}

func (s *StartSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A request to stop a Sandbox.
type StopSandboxRequest struct {
	// Resource name of the sandbox to stop, in the form
	// `sandboxes/{sandbox_id}`.
	Name string `json:"-" url:"-"`
}

func (s *StopSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateSandboxRequest struct {
	// Resource name of the sandbox to update, in the form
	// `sandboxes/{sandbox_id}`.
	Name string `json:"-" url:"-"`
	// The Sandbox resource carrying new field values. Only fields named in
	// `update_mask` are read; unmasked fields are ignored.
	Sandbox Sandbox `json:"sandbox"`
	// Field paths to update. Must be a non-empty subset of: - display_name -
	// spec.compute.inactivity_timeout Any other path returns
	// INVALID_PARAMETER_VALUE.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
