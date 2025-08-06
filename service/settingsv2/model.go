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

type ListAccountSettingsMetadataRequest struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous
	// `ListAccountSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListAccountSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountSettingsMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountSettingsMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata []SettingsMetadata `json:"settings_metadata,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountSettingsMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountSettingsMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceSettingsMetadataRequest struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous
	// `ListWorkspaceSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListWorkspaceSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceSettingsMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceSettingsMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata []SettingsMetadata `json:"settings_metadata,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceSettingsMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceSettingsMetadataResponse) MarshalJSON() ([]byte, error) {
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

type SettingsMetadata struct {
	// Setting description for what this setting controls
	Description string `json:"description,omitempty"`
	// Link to databricks documentation for the setting
	DocsLink string `json:"docs_link,omitempty"`
	// Name of the setting.
	Name string `json:"name,omitempty"`
	// Type of the setting. To set this setting, the value sent must match this
	// type.
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SettingsMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SettingsMetadata) MarshalJSON() ([]byte, error) {
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
