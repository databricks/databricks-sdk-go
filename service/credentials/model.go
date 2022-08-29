// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package credentials

// all definitions in this file are in alphabetical order

type CreateCredentialRequest struct {
    // Databricks account ID. When you create or manage workspaces, your account
    // must be on the E2 version of the platform or on a select custom plan that
    // allows multiple workspaces per account. If you are configuring log
    // delivery, all account types are supported. For non-E2 account types, get
    // your account ID from the [Accounts
    // Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
    AccountId string ` path:"account_id"`
    
    AwsCredentials CreateCredentialRequestAwsCredentials `json:"aws_credentials"`
    // The human-readable name of the credential configuration object.
    CredentialsName string `json:"credentials_name"`
}


type CreateCredentialRequestAwsCredentials struct {
    
    StsRole *CreateCredentialRequestAwsCredentialsStsRole `json:"sts_role,omitempty"`
}


type CreateCredentialRequestAwsCredentialsStsRole struct {
    // The Amazon Resource Name (ARN) of the cross account role.
    RoleArn string `json:"role_arn,omitempty"`
}


type Credential struct {
    // The Databricks account ID that hosts the credential.
    AccountId string `json:"account_id,omitempty"`
    
    AwsCredentials *CredentialAwsCredentials `json:"aws_credentials,omitempty"`
    // Time in epoch milliseconds when the credential was created.
    CreationTime int64 `json:"creation_time,omitempty"`
    // Databricks credential configuration ID.
    CredentialsId string `json:"credentials_id,omitempty"`
    // The human-readable name of the credential configuration object.
    CredentialsName string `json:"credentials_name,omitempty"`
}


type CredentialAwsCredentials struct {
    
    StsRole *CredentialAwsCredentialsStsRole `json:"sts_role,omitempty"`
}


type CredentialAwsCredentialsStsRole struct {
    // The external ID that needs to be trusted by the cross-account role. This
    // is always your Databricks account ID.
    ExternalId string `json:"external_id,omitempty"`
    // The Amazon Resource Name (ARN) of the cross-account role.
    RoleArn string `json:"role_arn,omitempty"`
}


type DeleteCredentialConfigRequest struct {
    // Databricks account ID. When you create or manage workspaces, your account
    // must be on the E2 version of the platform or on a select custom plan that
    // allows multiple workspaces per account. If you are configuring log
    // delivery, all account types are supported. For non-E2 account types, get
    // your account ID from the [Accounts
    // Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
    AccountId string ` path:"account_id"`
    // Databricks Account API credential configuration ID
    CredentialsId string ` path:"credentials_id"`
}


type GetCredentialConfigRequest struct {
    // Databricks account ID. When you create or manage workspaces, your account
    // must be on the E2 version of the platform or on a select custom plan that
    // allows multiple workspaces per account. If you are configuring log
    // delivery, all account types are supported. For non-E2 account types, get
    // your account ID from the [Accounts
    // Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
    AccountId string ` path:"account_id"`
    // Databricks Account API credential configuration ID
    CredentialsId string ` path:"credentials_id"`
}


type ListCredentialsRequest struct {
    // Databricks account ID. When you create or manage workspaces, your account
    // must be on the E2 version of the platform or on a select custom plan that
    // allows multiple workspaces per account. If you are configuring log
    // delivery, all account types are supported. For non-E2 account types, get
    // your account ID from the [Accounts
    // Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
    AccountId string ` path:"account_id"`
}

// List of credential configuration objects.
type ListCredentialsResponse []Credential

