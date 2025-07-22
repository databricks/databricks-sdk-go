// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingsv2

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type BooleanMessage struct {
	Value bool `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BooleanMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BooleanMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPublicAccountSettingRequest struct {
	Name string `json:"-" url:"-"`
}

type GetPublicWorkspaceSettingRequest struct {
	Name string `json:"-" url:"-"`
}

type IntegerMessage struct {
	Value int `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IntegerMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IntegerMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PatchPublicAccountSettingRequest struct {
	Name string `json:"-" url:"-"`

	Setting Setting `json:"setting"`
}

type PatchPublicWorkspaceSettingRequest struct {
	Name string `json:"-" url:"-"`

	Setting Setting `json:"setting"`
}

type Setting struct {
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`

	EffectiveBooleanVal *BooleanMessage `json:"effective_boolean_val,omitempty"`

	EffectiveIntegerVal *IntegerMessage `json:"effective_integer_val,omitempty"`

	EffectiveStringVal *StringMessage `json:"effective_string_val,omitempty"`

	IntegerVal *IntegerMessage `json:"integer_val,omitempty"`
	// Name of the setting.
	Name string `json:"name,omitempty"`

	StringVal *StringMessage `json:"string_val,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Setting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Setting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StringMessage struct {
	// Represents a generic string value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StringMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StringMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
