// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Settings V2, Workspace Settings V2, etc.
package settingsv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type AccountSettingsV2Interface interface {

	// Get a setting value at account level
	GetPublicAccountSetting(ctx context.Context, request GetPublicAccountSettingRequest) (*Setting, error)

	// Patch a setting value at account level
	PatchPublicAccountSetting(ctx context.Context, request PatchPublicAccountSettingRequest) (*Setting, error)
}

func NewAccountSettingsV2(client *client.DatabricksClient) *AccountSettingsV2API {
	return &AccountSettingsV2API{
		accountSettingsV2Impl: accountSettingsV2Impl{
			client: client,
		},
	}
}

// APIs to manage account level settings
type AccountSettingsV2API struct {
	accountSettingsV2Impl
}

type WorkspaceSettingsV2Interface interface {

	// Get a setting value at workspace level
	GetPublicWorkspaceSetting(ctx context.Context, request GetPublicWorkspaceSettingRequest) (*Setting, error)

	// Patch a setting value at workspace level
	PatchPublicWorkspaceSetting(ctx context.Context, request PatchPublicWorkspaceSettingRequest) (*Setting, error)
}

func NewWorkspaceSettingsV2(client *client.DatabricksClient) *WorkspaceSettingsV2API {
	return &WorkspaceSettingsV2API{
		workspaceSettingsV2Impl: workspaceSettingsV2Impl{
			client: client,
		},
	}
}

// APIs to manage workspace level settings
type WorkspaceSettingsV2API struct {
	workspaceSettingsV2Impl
}
