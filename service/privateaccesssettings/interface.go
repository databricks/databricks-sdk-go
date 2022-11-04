// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package privateaccesssettings

import (
	"context"
)

// These APIs manage private access settings for this account. A private access
// settings object specifies how your workspace is accessed using AWS
// PrivateLink. Each workspace that has any PrivateLink connections must include
// the ID for a private access settings object is in its workspace configuration
// object. Your account must be enabled for PrivateLink to use these APIs.
// Before configuring PrivateLink, it is important to read the [Databricks
// article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type PrivateAccessSettingsService interface {

	// Create private access settings
	//
	// Creates a private access settings object, which specifies how your
	// workspace is accessed over [AWS
	// PrivateLink](https://aws.amazon.com/privatelink). To use AWS PrivateLink,
	// a workspace must have a private access settings object referenced by ID
	// in the workspace's `private_access_settings_id` property.
	//
	// You can share one private access settings with multiple workspaces in a
	// single account. However, private access settings are specific to AWS
	// regions, so only workspaces in the same AWS region can use a given
	// private access settings object.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	//
	// This operation is available only if your account is on the E2 version of
	// the platform and your Databricks account is enabled for PrivateLink
	// (Public Preview). Contact your Databricks representative to enable your
	// account for PrivateLink.
	CreatePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error)

	// Delete a private access settings object
	//
	// Deletes a private access settings object, which determines how your
	// workspace is accessed over [AWS
	// PrivateLink](https://aws.amazon.com/privatelink).
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	//
	// This operation is available only if your account is on the E2 version of
	// the platform and your Databricks account is enabled for PrivateLink
	// (Public Preview). Contact your Databricks representative to enable your
	// account for PrivateLink.
	DeletePrivateAccessSettings(ctx context.Context, request DeletePrivateAccessSettingsRequest) error

	// DeletePrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId calls DeletePrivateAccessSettings, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeletePrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) error

	// Get all private access settings objects
	//
	// Gets a list of all private access settings objects for an account,
	// specified by ID.
	//
	// This operation is available only if your account is on the E2 version of
	// the platform and your Databricks account is enabled for AWS PrivateLink
	// (Public Preview). Contact your Databricks representative to enable your
	// account for PrivateLink.
	GetAllPrivateAccessSettings(ctx context.Context, request GetAllPrivateAccessSettingsRequest) ([]PrivateAccessSettings, error)

	// GetAllPrivateAccessSettingsByAccountId calls GetAllPrivateAccessSettings, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetAllPrivateAccessSettingsByAccountId(ctx context.Context, accountId string) ([]PrivateAccessSettings, error)

	// PrivateAccessSettingsPrivateAccessSettingsNameToPrivateAccessSettingsIdMap retrieves a mapping to access ID by name
	//
	// This method is generated by Databricks SDK Code Generator.
	PrivateAccessSettingsPrivateAccessSettingsNameToPrivateAccessSettingsIdMap(ctx context.Context, request GetAllPrivateAccessSettingsRequest) (map[string]string, error)

	// GetPrivateAccessSettingsByPrivateAccessSettingsName retrieves PrivateAccessSettings by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetPrivateAccessSettingsByPrivateAccessSettingsName(ctx context.Context, name string) (*PrivateAccessSettings, error)

	// Get a private access settings object
	//
	// Gets a private access settings object, which specifies how your workspace
	// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	//
	// This operation is available only if your account is on the E2 version of
	// the platform and your Databricks account is enabled for PrivateLink
	// (Public Preview). Contact your Databricks representative to enable your
	// account for PrivateLink.
	GetPrivateAccessSettings(ctx context.Context, request GetPrivateAccessSettingsRequest) (*PrivateAccessSettings, error)

	// GetPrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId calls GetPrivateAccessSettings, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetPrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) (*PrivateAccessSettings, error)

	// Replace private access settings
	//
	// Updates an existing private access settings object, which specifies how
	// your workspace is accessed over [AWS
	// PrivateLink](https://aws.amazon.com/privatelink). To use AWS PrivateLink,
	// a workspace must have a private access settings object referenced by ID
	// in the workspace's `private_access_settings_id` property.
	//
	// This operation completely overwrites your existing private access
	// settings object attached to your workspaces. All workspaces attached to
	// the private access settings are affected by any change. If
	// `public_access_enabled`, `private_access_level`, or
	// `allowed_vpc_endpoint_ids` are updated, effects of these changes might
	// take several minutes to propagate to the workspace API. You can share one
	// private access settings object with multiple workspaces in a single
	// account. However, private access settings are specific to AWS regions, so
	// only workspaces in the same AWS region can use a given private access
	// settings object.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	//
	// This operation is available only if your account is on the E2 version of
	// the platform and your Databricks account is enabled for PrivateLink
	// (Public Preview). Contact your Databricks representative to enable your
	// account for PrivateLink.
	ReplacePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error
}
