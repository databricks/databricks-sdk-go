// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package privateaccesssettings

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewPrivateAccessSettings(client *client.DatabricksClient) PrivateAccessSettingsService {
	return &PrivateAccessSettingsAPI{
		client: client,
	}
}

type PrivateAccessSettingsAPI struct {
	client *client.DatabricksClient
}

// Create private access settings
//
// Creates a private access settings object, which specifies how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink). To
// use AWS PrivateLink, a workspace must have a private access settings object
// referenced by ID in the workspace's `private_access_settings_id` property.
//
// You can share one private access settings with multiple workspaces in a
// single account. However, private access settings are specific to AWS regions,
// so only workspaces in the same AWS region can use a given private access
// settings object.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) CreatePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", request.AccountId)
	err := a.client.Post(ctx, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

// Delete a private access settings object
//
// Deletes a private access settings object, which determines how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) DeletePrivateAccessSettings(ctx context.Context, request DeletePrivateAccessSettingsRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", request.AccountId, request.PrivateAccessSettingsId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a private access settings object
//
// Deletes a private access settings object, which determines how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) DeletePrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) error {
	return a.DeletePrivateAccessSettings(ctx, DeletePrivateAccessSettingsRequest{
		AccountId:               accountId,
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}

// Get all private access settings objects
//
// Gets a list of all private access settings objects for an account, specified
// by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for AWS PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetAllPrivateAccessSettings(ctx context.Context, request GetAllPrivateAccessSettingsRequest) ([]PrivateAccessSettings, error) {
	var privateAccessSettingsList []PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", request.AccountId)
	err := a.client.Get(ctx, path, request, &privateAccessSettingsList)
	return privateAccessSettingsList, err
}

func (a *PrivateAccessSettingsAPI) PrivateAccessSettingsPrivateAccessSettingsNameToPrivateAccessSettingsIdMap(ctx context.Context, request GetAllPrivateAccessSettingsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.GetAllPrivateAccessSettings(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.PrivateAccessSettingsName] = v.PrivateAccessSettingsId
	}
	return mapping, nil
}

func (a *PrivateAccessSettingsAPI) GetPrivateAccessSettingsByPrivateAccessSettingsName(ctx context.Context, name string) (*PrivateAccessSettings, error) {
	result, err := a.GetAllPrivateAccessSettings(ctx, GetAllPrivateAccessSettingsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.PrivateAccessSettingsName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("PrivateAccessSettings named '%s' does not exist", name)
}

// Get all private access settings objects
//
// Gets a list of all private access settings objects for an account, specified
// by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for AWS PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetAllPrivateAccessSettingsByAccountId(ctx context.Context, accountId string) ([]PrivateAccessSettings, error) {
	return a.GetAllPrivateAccessSettings(ctx, GetAllPrivateAccessSettingsRequest{
		AccountId: accountId,
	})
}

// Get a private access settings object
//
// Gets a private access settings object, which specifies how your workspace is
// accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetPrivateAccessSettings(ctx context.Context, request GetPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", request.AccountId, request.PrivateAccessSettingsId)
	err := a.client.Get(ctx, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

// Get a private access settings object
//
// Gets a private access settings object, which specifies how your workspace is
// accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetPrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) (*PrivateAccessSettings, error) {
	return a.GetPrivateAccessSettings(ctx, GetPrivateAccessSettingsRequest{
		AccountId:               accountId,
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}

// Replace private access settings
//
// Updates an existing private access settings object, which specifies how your
// workspace is accessed over [AWS
// PrivateLink](https://aws.amazon.com/privatelink). To use AWS PrivateLink, a
// workspace must have a private access settings object referenced by ID in the
// workspace's `private_access_settings_id` property.
//
// This operation completely overwrites your existing private access settings
// object attached to your workspaces. All workspaces attached to the private
// access settings are affected by any change. If `public_access_enabled`,
// `private_access_level`, or `allowed_vpc_endpoint_ids` are updated, effects of
// these changes might take several minutes to propagate to the workspace API.
// You can share one private access settings object with multiple workspaces in
// a single account. However, private access settings are specific to AWS
// regions, so only workspaces in the same AWS region can use a given private
// access settings object.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) ReplacePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", request.AccountId, request.PrivateAccessSettingsId)
	err := a.client.Put(ctx, path, request)
	return err
}
