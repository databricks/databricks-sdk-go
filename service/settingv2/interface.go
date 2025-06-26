// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingv2

import (
	"context"
)

// Settings API allows users to manage settings.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountSettingsV2Service interface {
	GetPublicAccountSetting(ctx context.Context, request GetPublicAccountSettingRequest) (*Setting, error)

	PatchPublicAccountSetting(ctx context.Context, request PatchPublicAccountSettingRequest) (*Setting, error)
}

// Settings API allows users to manage settings.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceSettingsV2Service interface {
	GetPublicWorkspaceSetting(ctx context.Context, request GetPublicWorkspaceSettingRequest) (*Setting, error)

	PatchPublicWorkspaceSetting(ctx context.Context, request PatchPublicWorkspaceSettingRequest) (*Setting, error)
}
