// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingsv2

import (
	"context"
)

// APIs to manage account level settings
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountSettingsV2Service interface {

	// Get a setting value at account level
	GetPublicAccountSetting(ctx context.Context, request GetPublicAccountSettingRequest) (*Setting, error)

	// Patch a setting value at account level
	PatchPublicAccountSetting(ctx context.Context, request PatchPublicAccountSettingRequest) (*Setting, error)
}

// APIs to manage workspace level settings
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceSettingsV2Service interface {

	// Get a setting value at workspace level
	GetPublicWorkspaceSetting(ctx context.Context, request GetPublicWorkspaceSettingRequest) (*Setting, error)

	// Patch a setting value at workspace level
	PatchPublicWorkspaceSetting(ctx context.Context, request PatchPublicWorkspaceSettingRequest) (*Setting, error)
}
