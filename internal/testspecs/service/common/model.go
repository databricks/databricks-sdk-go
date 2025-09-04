// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package common

import (
	"encoding/json"

	"github.com/databricks/databricks-sdk-go/internal/testspecs/service/lrotesting"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// Serialization format for DatabricksServiceException with error details. This
// message doesn't work for ScalaPB-04 as google.protobuf.Any is only available
// to ScalaPB-09. Note the definition of this message should be in sync with
// DatabricksServiceExceptionProto defined in
// /api-base/proto/legacy/databricks.proto except the later one doesn't have the
// error details field defined.
type DatabricksServiceExceptionWithDetailsProto struct {
	Details []json.RawMessage `json:"details,omitempty"`

	ErrorCode lrotesting.ErrorCode `json:"error_code,omitempty"`

	Message string `json:"message,omitempty"`

	StackTrace string `json:"stack_trace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksServiceExceptionWithDetailsProto) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksServiceExceptionWithDetailsProto) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// If the value is `false`, it means the operation is still in progress. If
	// `true`, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `json:"done,omitempty"`
	// The error result of the operation in case of failure or cancellation.
	Error *DatabricksServiceExceptionWithDetailsProto `json:"error,omitempty"`
	// Service-specific metadata associated with the operation. It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata. Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// The server-assigned name, which is only unique within the same service
	// that originally returns it. If you use the default HTTP mapping, the
	// `name` should be a resource name ending with `operations/{unique_id}`.
	//
	// Note: multi-segment resource names are not yet supported in the RPC
	// framework and SDK/TF. Until that support is added, `name` must be string
	// without internal `/` separators.
	Name string `json:"name,omitempty"`
	// The normal, successful response of the operation. If the original method
	// returns no data on success, such as `Delete`, the response is
	// `google.protobuf.Empty`. If the original method is standard
	// `Get`/`Create`/`Update`, the response should be the resource. For other
	// methods, the response should have the type `XxxResponse`, where `Xxx` is
	// the original method name. For example, if the original method name is
	// `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
	Response json.RawMessage `json:"response,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Operation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Operation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
