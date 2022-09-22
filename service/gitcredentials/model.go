// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package gitcredentials

// all definitions in this file are in alphabetical order

type CreateCredentials struct {
	GitProvider string `json:"git_provider"`

	GitUsername string `json:"git_username,omitempty"`

	PersonalAccessToken string `json:"personal_access_token,omitempty"`
}

type CreateCredentialsResponse struct {
	CredentialId int64 `json:"credential_id,omitempty"`

	GitProvider string `json:"git_provider,omitempty"`

	GitUsername string `json:"git_username,omitempty"`
}

type GetCredentialResponse struct {
	CredentialId int64 `json:"credential_id,omitempty"`

	GitProvider string `json:"git_provider,omitempty"`

	GitUsername string `json:"git_username,omitempty"`
}

type GetCredentialsResponse struct {
	Credentials []GetCredentialResponse `json:"credentials,omitempty"`
}

type UpdateCredentials struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" path:"credential_id"`

	GitProvider string `json:"git_provider,omitempty"`

	GitUsername string `json:"git_username,omitempty"`

	PersonalAccessToken string `json:"personal_access_token,omitempty"`
}

// ID of the credential object in the workspace.

type DeleteRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" path:"credential_id"`
}

type GetRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" path:"credential_id"`
}

// Git provider. This field is case-insensitive. The available Git providers are
// awsCodeCommit, azureDevOpsServices,

// Git username.

// The personal access token used to authenticate to the corresponding Git
// provider.
