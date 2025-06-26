// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Settings V2, Workspace Settings V2, etc.
package settingv2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type AccountSettingsV2Interface interface {
	GetPublicAccountSetting(ctx context.Context, request GetPublicAccountSettingRequest) (*Setting, error)

	GetPublicAccountSettingByName(ctx context.Context, name string) (*Setting, error)

	PatchPublicAccountSetting(ctx context.Context, request PatchPublicAccountSettingRequest) (*Setting, error)
}

func NewAccountSettingsV2(client *client.DatabricksClient) *AccountSettingsV2API {
	return &AccountSettingsV2API{
		accountSettingsV2Impl: accountSettingsV2Impl{
			client: client,
		},
	}
}

// Settings API allows users to manage settings.
type AccountSettingsV2API struct {
	accountSettingsV2Impl
}

func (a *AccountSettingsV2API) GetPublicAccountSettingByName(ctx context.Context, name string) (*Setting, error) {
	return a.accountSettingsV2Impl.GetPublicAccountSetting(ctx, GetPublicAccountSettingRequest{
		Name: name,
	})
}

type WorkspaceSettingsV2Interface interface {
	GetPublicWorkspaceSetting(ctx context.Context, request GetPublicWorkspaceSettingRequest) (*Setting, error)

	GetPublicWorkspaceSettingByName(ctx context.Context, name string) (*Setting, error)

	PatchPublicWorkspaceSetting(ctx context.Context, request PatchPublicWorkspaceSettingRequest) (*Setting, error)
}

func NewWorkspaceSettingsV2(client *client.DatabricksClient) *WorkspaceSettingsV2API {
	return &WorkspaceSettingsV2API{
		workspaceSettingsV2Impl: workspaceSettingsV2Impl{
			client: client,
		},
	}
}

// Settings API allows users to manage settings.
type WorkspaceSettingsV2API struct {
	workspaceSettingsV2Impl
}

func (a *WorkspaceSettingsV2API) GetPublicWorkspaceSettingByName(ctx context.Context, name string) (*Setting, error) {
	return a.workspaceSettingsV2Impl.GetPublicWorkspaceSetting(ctx, GetPublicWorkspaceSettingRequest{
		Name: name,
	})
}
