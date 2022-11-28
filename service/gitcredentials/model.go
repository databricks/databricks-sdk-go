// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package gitcredentials

// all definitions in this file are in alphabetical order

type CreateCredentials struct {
	// Git provider. This field is case-insensitive. The available Git providers
	// are awsCodeCommit, azureDevOpsServices,
	GitProvider string `json:"git_provider"`
	// Git username.
	GitUsername string `json:"git_username,omitempty"`
	// The personal access token used to authenticate to the corresponding Git
	// provider.
	PersonalAccessToken string `json:"personal_access_token,omitempty"`
}

type CreateCredentialsResponse struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are awsCodeCommit, azureDevOpsServices,
	GitProvider string `json:"git_provider,omitempty"`
	// Git username.
	GitUsername string `json:"git_username,omitempty"`
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are awsCodeCommit, azureDevOpsServices,
	GitProvider string `json:"git_provider,omitempty"`
	// Git username.
	GitUsername string `json:"git_username,omitempty"`
}

// Delete a credential
type Delete struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" path:"credential_id"`
}

// Get a credential entry
type Get struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" path:"credential_id"`
}

type GetCredentialsResponse struct {
	Credentials []CredentialInfo `json:"credentials,omitempty"`
}

type UpdateCredentials struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" path:"credential_id"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are awsCodeCommit, azureDevOpsServices,
	GitProvider string `json:"git_provider,omitempty"`
	// Git username.
	GitUsername string `json:"git_username,omitempty"`
	// The personal access token used to authenticate to the corresponding Git
	// provider.
	PersonalAccessToken string `json:"personal_access_token,omitempty"`
}
