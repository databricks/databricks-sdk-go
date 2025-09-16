// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lrotesting

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name string `json:"-" url:"-"`
}

type CreateTestResourceRequest struct {
	// The resource to create
	Resource TestResource `json:"resource"`
}

type GetOperationRequest struct {
	// The name of the operation resource.
	Name string `json:"-" url:"-"`
}

type GetTestResourceRequest struct {
	// Resource ID to get
	ResourceId string `json:"-" url:"-"`
}

// Test resource for LRO operations
type TestResource struct {
	// Unique identifier for the resource
	Id string `json:"id,omitempty"`
	// Name of the resource
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TestResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TestResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata for test resource operations
type TestResourceOperationMetadata struct {
	// Progress percentage (0-100)
	ProgressPercent int `json:"progress_percent,omitempty"`
	// ID of the resource being operated on
	ResourceId string `json:"resource_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TestResourceOperationMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TestResourceOperationMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
