// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package v2

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type MethodOptions struct {
	Query string `json:"query,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MethodOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MethodOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
