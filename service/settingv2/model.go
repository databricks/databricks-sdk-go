// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingv2

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type GetPublicAccountSettingRequest struct {
	Name string `json:"-" url:"-"`
}

type GetPublicWorkspaceSettingRequest struct {
	Name string `json:"-" url:"-"`
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
	// Name of the setting.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Setting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Setting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
