// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package privateaccesssettings

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// These APIs manage private access settings for this account. A private access
// settings object specifies how your workspace is accessed using AWS
// PrivateLink. Each workspace that has any PrivateLink connections must include
// the ID for a private access settings object is in its workspace configuration
// object. Your account must be enabled for PrivateLink to use these APIs.
// Before configuring PrivateLink, it is important to read the [Databricks
// article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
type PrivateaccesssettingsService interface {
	// Create a private access settings object, which specifies how your
	// workspace is accessed over [AWS
	// PrivateLink](https://aws.amazon.com/privatelink). To use AWS PrivateLink,
	// a workspace must have a private access settings object referenced by ID
	// in the workspace&#39;s `private_access_settings_id` property. You can share
	// one private access settings with multiple workspaces in a single
	// account.However, private access settings are region specific, so only
	// workspaces in the same region may use a given private access settings
	// object. Before configuring PrivateLink, it is important to read the
	// [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	// This operation is available only if your account is on the E2 version of
	// the platform and your Databricks account is enabled for PrivateLink
	// (Public Preview). Contact your Databricks representative to enable your
	// account for PrivateLink.
	CreatePrivateAccessSettings(ctx context.Context, upsertPrivateAccessSettingsRequest UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error)
	// Delete a private access settings object, which determines how your
	// workspace is accessed over [AWS
	// PrivateLink](https://aws.amazon.com/privatelink). Before configuring
	// PrivateLink, it is important to read the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	// This operation is available only if your account is on the E2 version of
	// the platform and your Databricks account is enabled for PrivateLink
	// (Public Preview). Contact your Databricks representative to enable your
	// account for PrivateLink.
	DeletePrivateAccessSettings(ctx context.Context, deletePrivateAccessSettingsRequest DeletePrivateAccessSettingsRequest) error
	// Get a list of all private access settings objects for an account,
	// specified by ID. This operation is available only if your account is on
	// the E2 version of the platform and your Databricks account is enabled for
	// PrivateLink (Public Preview). Contact your Databricks representative to
	// enable your account for PrivateLink.
	GetAllPrivateAccessSettings(ctx context.Context, getAllPrivateAccessSettingsRequest GetAllPrivateAccessSettingsRequest) (*ListPrivateAccessSettingsResponse, error)
	// Get a private access settings object, which specifies how your workspace
	// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
	// Before configuring PrivateLink, it is important to read the [Databricks
	// article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	// This operation is available only if your account is on the E2 version of
	// the platform and your Databricks account is enabled for PrivateLink
	// (Public Preview). Contact your Databricks representative to enable your
	// account for PrivateLink.
	GetPrivateAccessSettings(ctx context.Context, getPrivateAccessSettingsRequest GetPrivateAccessSettingsRequest) (*PrivateAccessSettings, error)
	// Update an existing private access settings object, which specifies how
	// your workspace is accessed over [AWS
	// PrivateLink](https://aws.amazon.com/privatelink). To use AWS PrivateLink,
	// a workspace must have a private access settings object referenced by ID
	// in the workspace&#39;s `private_access_settings_id` property. This operation
	// fully overwrites your existing private access settings object attached to
	// your workspaces. All workspaces attached to the private access settings
	// will see the effects of any change. If updating `public_access_enabled`,
	// `private_access_level`, or `allowed_vpc_endpoint_ids`, effects of the
	// change may take a couple minutes to propagate to the workspace API. You
	// can share one private access settings with multiple workspaces in a
	// single account. However, private access settings are region specific, so
	// only workspaces in the same region may use a given private access
	// settings object. Before configuring PrivateLink, it is important to read
	// the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	// This operation is available only if your account is on the E2 version of
	// the platform and your Databricks account is enabled for PrivateLink
	// (Public Preview). Contact your Databricks representative to enable your
	// account for PrivateLink.
	ReplacePrivateAccessSettings(ctx context.Context, upsertPrivateAccessSettingsRequest UpsertPrivateAccessSettingsRequest) error
	GetAllPrivateAccessSettingsByAccountId(ctx context.Context, accountId string) (*ListPrivateAccessSettingsResponse, error)
	DeletePrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) error
	GetPrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) (*PrivateAccessSettings, error)
}

func New(client *client.DatabricksClient) PrivateaccesssettingsService {
	return &PrivateaccesssettingsAPI{
		client: client,
	}
}

type PrivateaccesssettingsAPI struct {
	client *client.DatabricksClient
}

// Create a private access settings object, which specifies how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink). To
// use AWS PrivateLink, a workspace must have a private access settings object
// referenced by ID in the workspace&#39;s `private_access_settings_id` property.
// You can share one private access settings with multiple workspaces in a
// single account.However, private access settings are region specific, so only
// workspaces in the same region may use a given private access settings object.
// Before configuring PrivateLink, it is important to read the [Databricks
// article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateaccesssettingsAPI) CreatePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/accounts/%v/private-access-settings", request.AccountId)
	err := a.client.Post(ctx, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

// Delete a private access settings object, which determines how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
// Before configuring PrivateLink, it is important to read the [Databricks
// article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateaccesssettingsAPI) DeletePrivateAccessSettings(ctx context.Context, request DeletePrivateAccessSettingsRequest) error {
	path := fmt.Sprintf("/accounts/%v/private-access-settings/%v", request.AccountId, request.PrivateAccessSettingsId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get a list of all private access settings objects for an account, specified
// by ID. This operation is available only if your account is on the E2 version
// of the platform and your Databricks account is enabled for PrivateLink
// (Public Preview). Contact your Databricks representative to enable your
// account for PrivateLink.
func (a *PrivateaccesssettingsAPI) GetAllPrivateAccessSettings(ctx context.Context, request GetAllPrivateAccessSettingsRequest) (*ListPrivateAccessSettingsResponse, error) {
	var listPrivateAccessSettingsResponse ListPrivateAccessSettingsResponse
	path := fmt.Sprintf("/accounts/%v/private-access-settings", request.AccountId)
	err := a.client.Get(ctx, path, request, &listPrivateAccessSettingsResponse)
	return &listPrivateAccessSettingsResponse, err
}

// Get a private access settings object, which specifies how your workspace is
// accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink). Before
// configuring PrivateLink, it is important to read the [Databricks article
// about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateaccesssettingsAPI) GetPrivateAccessSettings(ctx context.Context, request GetPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/accounts/%v/private-access-settings/%v", request.AccountId, request.PrivateAccessSettingsId)
	err := a.client.Get(ctx, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

// Update an existing private access settings object, which specifies how your
// workspace is accessed over [AWS
// PrivateLink](https://aws.amazon.com/privatelink). To use AWS PrivateLink, a
// workspace must have a private access settings object referenced by ID in the
// workspace&#39;s `private_access_settings_id` property. This operation fully
// overwrites your existing private access settings object attached to your
// workspaces. All workspaces attached to the private access settings will see
// the effects of any change. If updating `public_access_enabled`,
// `private_access_level`, or `allowed_vpc_endpoint_ids`, effects of the change
// may take a couple minutes to propagate to the workspace API. You can share
// one private access settings with multiple workspaces in a single account.
// However, private access settings are region specific, so only workspaces in
// the same region may use a given private access settings object. Before
// configuring PrivateLink, it is important to read the [Databricks article
// about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateaccesssettingsAPI) ReplacePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {
	path := fmt.Sprintf("/accounts/%v/private-access-settings/%v", request.AccountId, request.PrivateAccessSettingsId)
	err := a.client.Put(ctx, path, request)
	return err
}

func (a *PrivateaccesssettingsAPI) GetAllPrivateAccessSettingsByAccountId(ctx context.Context, accountId string) (*ListPrivateAccessSettingsResponse, error) {
	return a.GetAllPrivateAccessSettings(ctx, GetAllPrivateAccessSettingsRequest{
		AccountId: accountId,
	})
}

func (a *PrivateaccesssettingsAPI) DeletePrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) error {
	return a.DeletePrivateAccessSettings(ctx, DeletePrivateAccessSettingsRequest{
		AccountId:               accountId,
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}

func (a *PrivateaccesssettingsAPI) GetPrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) (*PrivateAccessSettings, error) {
	return a.GetPrivateAccessSettings(ctx, GetPrivateAccessSettingsRequest{
		AccountId:               accountId,
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}
