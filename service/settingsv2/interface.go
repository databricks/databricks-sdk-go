// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingsv2

import (
	"context"
)

// APIs to manage account level settings
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountSettingsV2Service interface {

	// Get a setting value at account level. See
	// :method:settingsv2/listaccountsettingsmetadata for list of setting
	// available via public APIs at account level.
	GetPublicAccountSetting(ctx context.Context, request GetPublicAccountSettingRequest) (*Setting, error)

	// List valid setting keys and metadata. These settings are available to be
	// referenced via GET :method:settingsv2/getpublicaccountsetting and PATCH
	// :method:settingsv2/patchpublicaccountsetting APIs
	ListAccountSettingsMetadata(ctx context.Context, request ListAccountSettingsMetadataRequest) (*ListAccountSettingsMetadataResponse, error)

	// Patch a setting value at account level. See
	// :method:settingsv2/listaccountsettingsmetadata for list of setting
	// available via public APIs at account level.
	PatchPublicAccountSetting(ctx context.Context, request PatchPublicAccountSettingRequest) (*Setting, error)
}

// APIs to manage workspace level settings
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceSettingsV2Service interface {

	// Get a setting value at workspace level. See
	// :method:settingsv2/listworkspacesettingsmetadata for list of setting
	// available via public APIs.
	GetPublicWorkspaceSetting(ctx context.Context, request GetPublicWorkspaceSettingRequest) (*Setting, error)

	// List valid setting keys and metadata. These settings are available to be
	// referenced via GET :method:settingsv2/getpublicworkspacesetting and PATCH
	// :method:settingsv2/patchpublicworkspacesetting APIs
	ListWorkspaceSettingsMetadata(ctx context.Context, request ListWorkspaceSettingsMetadataRequest) (*ListWorkspaceSettingsMetadataResponse, error)

	// Patch a setting value at workspace level. See
	// :method:settingsv2/listworkspacesettingsmetadata for list of setting
	// available via public APIs at workspace level.
	PatchPublicWorkspaceSetting(ctx context.Context, request PatchPublicWorkspaceSettingRequest) (*Setting, error)
}
