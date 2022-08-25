// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package secrets

// all definitions in this file are in alphabetical order

type AclItem struct {
    // The permission level applied to the principal. 
    Permission AclItemPermission `json:"permission"`
    // The principal in which the permission is applied. 
    Principal string `json:"principal"`
}

// The permission level applied to the principal. 
type AclItemPermission string


const AclItemPermissionRead AclItemPermission = `READ`

const AclItemPermissionWrite AclItemPermission = `WRITE`

const AclItemPermissionManage AclItemPermission = `MANAGE`

type AzureKeyVaultSecretScopeMetadata struct {
    // The DNS of the KeyVault 
    DnsName string `json:"dns_name"`
    // The resource id of the azure KeyVault that user wants to associate the 
    // scope with. 
    ResourceId string `json:"resource_id"`
}


type CreateScopeRequest struct {
    // The principal that is initially granted ``MANAGE`` permission to the 
    // created scope. 
    InitialManagePrincipal string `json:"initial_manage_principal,omitempty"`
    // Scope name requested by the user. Scope names are unique. 
    Scope string `json:"scope"`
    // The backend type the scope will be created with. If not specified, will 
    // default to ``DATABRICKS`` 
    ScopeBackendType CreateScopeRequestScopeBackendType `json:"scope_backend_type,omitempty"`
}

// The backend type the scope will be created with. If not specified, will 
// default to ``DATABRICKS`` 
type CreateScopeRequestScopeBackendType string


const CreateScopeRequestScopeBackendTypeDatabricks CreateScopeRequestScopeBackendType = `DATABRICKS`

const CreateScopeRequestScopeBackendTypeAzureKeyvault CreateScopeRequestScopeBackendType = `AZURE_KEYVAULT`

type DeleteAclRequest struct {
    // The principal to remove an existing ACL from. 
    Principal string `json:"principal"`
    // The name of the scope to remove permissions from. 
    Scope string `json:"scope"`
}


type DeleteScopeRequest struct {
    // Name of the scope to delete. 
    Scope string `json:"scope"`
}


type DeleteSecretRequest struct {
    // Name of the secret to delete. 
    Key string `json:"key"`
    // The name of the scope that contains the secret to delete. 
    Scope string `json:"scope"`
}


type GetAclRequest struct {
    // The principal to fetch ACL information for. 
    Principal string ` url:"principal,omitempty"`
    // The name of the scope to fetch ACL information from. 
    Scope string ` url:"scope,omitempty"`
}


type GetAclResponse struct {
    // The permission level applied to the principal. 
    Permission GetAclResponsePermission `json:"permission"`
    // The principal in which the permission is applied. 
    Principal string `json:"principal"`
}

// The permission level applied to the principal. 
type GetAclResponsePermission string


const GetAclResponsePermissionRead GetAclResponsePermission = `READ`

const GetAclResponsePermissionWrite GetAclResponsePermission = `WRITE`

const GetAclResponsePermissionManage GetAclResponsePermission = `MANAGE`

type ListAclsRequest struct {
    // The name of the scope to fetch ACL information from. 
    Scope string ` url:"scope,omitempty"`
}


type ListAclsResponse struct {
    // The associated ACLs rule applied to principals in the given scope. 
    Items []AclItem `json:"items,omitempty"`
}


type ListScopesResponse struct {
    // The available secret scopes. 
    Scopes []SecretScope `json:"scopes,omitempty"`
}


type ListSecretsRequest struct {
    // The name of the scope to list secrets within. 
    Scope string ` url:"scope,omitempty"`
}


type ListSecretsResponse struct {
    // Metadata information of all secrets contained within the given scope. 
    Secrets []SecretMetadata `json:"secrets,omitempty"`
}


type PutAclRequest struct {
    // The permission level applied to the principal. 
    Permission PutAclRequestPermission `json:"permission"`
    // The principal in which the permission is applied. 
    Principal string `json:"principal"`
    // The name of the scope to apply permissions to. 
    Scope string `json:"scope"`
}

// The permission level applied to the principal. 
type PutAclRequestPermission string


const PutAclRequestPermissionRead PutAclRequestPermission = `READ`

const PutAclRequestPermissionWrite PutAclRequestPermission = `WRITE`

const PutAclRequestPermissionManage PutAclRequestPermission = `MANAGE`

type PutSecretRequest struct {
    // If specified, value will be stored as bytes. 
    BytesValue string `json:"bytes_value,omitempty"`
    // A unique name to identify the secret. 
    Key string `json:"key"`
    // The name of the scope to which the secret will be associated with. 
    Scope string `json:"scope"`
    // If specified, note that the value will be stored in UTF-8 (MB4) form. 
    StringValue string `json:"string_value,omitempty"`
}


type SecretMetadata struct {
    // A unique name to identify the secret. 
    Key string `json:"key,omitempty"`
    // The last updated timestamp (in milliseconds) for the secret. 
    LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
}


type SecretScope struct {
    // The type of secret scope backend. 
    BackendType SecretScopeBackendType `json:"backend_type,omitempty"`
    // The metadata for the secret scope if the type is ``AZURE_KEYVAULT`` 
    KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata `json:"keyvault_metadata,omitempty"`
    // A unique name to identify the secret scope. 
    Name string `json:"name,omitempty"`
}

// The type of secret scope backend. 
type SecretScopeBackendType string


const SecretScopeBackendTypeDatabricks SecretScopeBackendType = `DATABRICKS`

const SecretScopeBackendTypeAzureKeyvault SecretScopeBackendType = `AZURE_KEYVAULT`
