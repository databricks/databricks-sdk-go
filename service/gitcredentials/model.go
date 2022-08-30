// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package gitcredentials

// all definitions in this file are in alphabetical order

type CreateCredentialRequest struct {
    // Git provider. This field is case-insensitive. The available Git providers
    // are awsCodeCommit, azureDevOpsServices, 
    GitProvider string `json:"git_provider"`
    // Git username.
    GitUsername string `json:"git_username,omitempty"`
    // The personal access token used to authenticate to the corresponding Git
    // provider.
    PersonalAccessToken string `json:"personal_access_token,omitempty"`
}


type DeleteCredentialRequest struct {
    // The ID for the corresponding credential to access.
    CredentialId string ` path:"credential_id"`
}


type GetCredentialRequest struct {
    // The ID for the corresponding credential to access.
    CredentialId string ` path:"credential_id"`
}


type GetCredentialResponse struct {
    // ID of the credential object in the workspace.
    CredentialId int64 `json:"credential_id,omitempty"`
    // Git provider. This field is case-insensitive. The available Git providers
    // are awsCodeCommit, azureDevOpsServices, 
    GitProvider string `json:"git_provider,omitempty"`
    // Git username.
    GitUsername string `json:"git_username,omitempty"`
}


type GetCredentialsResponse struct {
    
    Credentials []GetCredentialResponse `json:"credentials,omitempty"`
}


type UpdateCredentialRequest struct {
    // The ID for the corresponding credential to access.
    CredentialId string ` path:"credential_id"`
    // Git provider. This field is case-insensitive. The available Git providers
    // are awsCodeCommit, azureDevOpsServices, 
    GitProvider string `json:"git_provider,omitempty"`
    // Git username.
    GitUsername string `json:"git_username,omitempty"`
    // The personal access token used to authenticate to the corresponding Git
    // provider.
    PersonalAccessToken string `json:"personal_access_token,omitempty"`
}

