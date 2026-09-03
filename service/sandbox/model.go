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
	// Customer-supplied display label. Mutable via UpdateSandbox. Bounds
	// enforced at the RPC boundary (<=256 bytes, mirrors lakebox
	// MAX_SANDBOX_NAME_LEN).
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

// Lifecycle state of a Sandbox resource. STOPPING is the transient state
// surfaced while a teardown (Stop, DeleteSandbox, auto-terminate, provisioning
// failure) is in flight but the sandbox row still exists; the row settles to
// STOPPED once the workflow finishes.
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
	// Field paths to update. Must be a non-empty subset of: -
	// metadata.display_name - spec.compute.inactivity_timeout Any other path
	// returns INVALID_PARAMETER_VALUE.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

func (s *UpdateSandboxRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
