// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package privateaccesssettings

// all definitions in this file are in alphabetical order

type DeletePrivateAccessSettingsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" path:"private_access_settings_id"`
}

type GetAllPrivateAccessSettingsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetPrivateAccessSettingsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" path:"private_access_settings_id"`
}

type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID
	// returned when registering the VPC endpoint configuration in your
	// Databricks account. This is _not_ the ID of the VPC endpoint in AWS.
	//
	// Only used when `private_access_level` is set to `ENDPOINT`. This is an
	// allow list of VPC endpoints registered in your Databricks account that
	// can connect to your workspace over AWS PrivateLink.
	//
	// **Note**: If hybrid access to your workspace is enabled by setting
	// `public_access_enabled` to `true`, this control only works for
	// PrivateLink connections. To control how your workspace is accessed via
	// public internet, see [IP access
	// lists](https://docs.databricks.com/security/network/ip-access-list.html).
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ANY` (deprecated): Any VPC endpoint can connect to your
	// workspace. * `ACCOUNT` level access (the default) allows only VPC
	// endpoints that are registered in your Databricks account connect to your
	// workspace. * `ENDPOINT` level access allows only specified VPC endpoints
	// connect to your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel PrivateAccessSettingsPrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string `json:"private_access_settings_name,omitempty"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The AWS region for workspaces attached to this private access settings
	// object.
	Region string `json:"region,omitempty"`
}

// The private access level controls which VPC endpoints can connect to the UI
// or API of any workspace that attaches this private access settings object. *
// `ANY` (deprecated): Any VPC endpoint can connect to your workspace. *
// `ACCOUNT` level access (the default) allows only VPC endpoints that are
// registered in your Databricks account connect to your workspace. * `ENDPOINT`
// level access allows only specified VPC endpoints connect to your workspace.
// For details, see `allowed_vpc_endpoint_ids`.
type PrivateAccessSettingsPrivateAccessLevel string

const PrivateAccessSettingsPrivateAccessLevelAccount PrivateAccessSettingsPrivateAccessLevel = `ACCOUNT`

const PrivateAccessSettingsPrivateAccessLevelAny PrivateAccessSettingsPrivateAccessLevel = `ANY`

const PrivateAccessSettingsPrivateAccessLevelEndpoint PrivateAccessSettingsPrivateAccessLevel = `ENDPOINT`

type UpsertPrivateAccessSettingsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID that
	// is returned when registering the VPC endpoint configuration in your
	// Databricks account. This is not the ID of the VPC endpoint in AWS.
	//
	// Only used when `private_access_level` is set to `ENDPOINT`. This is an
	// allow list of VPC endpoints that in your account that can connect to your
	// workspace over AWS PrivateLink.
	//
	// If hybrid access to your workspace is enabled by setting
	// `public_access_enabled` to `true`, this control only works for
	// PrivateLink connections. To control how your workspace is accessed via
	// public internet, see [IP access
	// lists](https://docs.databricks.com/security/network/ip-access-list.html).
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ANY` level access (deprecated) allows any VPC endpoint connect
	// to your workspace. It is unsupported to create new
	// `PrivateAccessSettings` objects with the deprecated `ANY` access level or
	// update an existing object to `ANY` access level from another value. *
	// `ACCOUNT` level access allows only VPC endpoints that are registered in
	// your Databricks account connect to your workspace. * `ENDPOINT` level
	// access allows only specified VPC endpoints connect to your workspace. For
	// details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel UpsertPrivateAccessSettingsRequestPrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" path:"private_access_settings_id"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string `json:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The AWS region for workspaces associated with this private access
	// settings object. This must be a [region that Databricks supports for
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/regions.html).
	Region string `json:"region"`
}

// The private access level controls which VPC endpoints can connect to the UI
// or API of any workspace that attaches this private access settings object. *
// `ANY` level access (deprecated) allows any VPC endpoint connect to your
// workspace. It is unsupported to create new `PrivateAccessSettings` objects
// with the deprecated `ANY` access level or update an existing object to `ANY`
// access level from another value. * `ACCOUNT` level access allows only VPC
// endpoints that are registered in your Databricks account connect to your
// workspace. * `ENDPOINT` level access allows only specified VPC endpoints
// connect to your workspace. For details, see `allowed_vpc_endpoint_ids`.
type UpsertPrivateAccessSettingsRequestPrivateAccessLevel string

const UpsertPrivateAccessSettingsRequestPrivateAccessLevelAccount UpsertPrivateAccessSettingsRequestPrivateAccessLevel = `ACCOUNT`

const UpsertPrivateAccessSettingsRequestPrivateAccessLevelAny UpsertPrivateAccessSettingsRequestPrivateAccessLevel = `ANY`

const UpsertPrivateAccessSettingsRequestPrivateAccessLevelEndpoint UpsertPrivateAccessSettingsRequestPrivateAccessLevel = `ENDPOINT`
