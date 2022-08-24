// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package datasharing

// all definitions in this file are in alphabetical order

type CreateProviderRequest struct {
    // [Create,Update:IGN] Whether this provider is successfully activated by 
    // the data provider. This field is only present when the authentication 
    // type is DATABRICKS. 
    ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType CreateProviderRequestAuthenticationType `json:"authentication_type,omitempty"`
    // [Create,Update:OPT] Description about the provider. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of Provider creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create,Update:REQ] The name of the Provider. 
    Name string `json:"name,omitempty"`
    // [Create,Update:IGN] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
    // [Create,Update:OPT] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
    // [Create,Update:IGN] The server-generated one-time sharing code. This 
    // field is only present when the authentication type is DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create,Update:IGN] Username of user who last modified Share. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type CreateProviderRequestAuthenticationType string


const CreateProviderRequestAuthenticationTypeUnknown CreateProviderRequestAuthenticationType = `UNKNOWN`

const CreateProviderRequestAuthenticationTypeToken CreateProviderRequestAuthenticationType = `TOKEN`

const CreateProviderRequestAuthenticationTypeDatabricks CreateProviderRequestAuthenticationType = `DATABRICKS`

type CreateProviderResponse struct {
    // [Create,Update:IGN] Whether this provider is successfully activated by 
    // the data provider. This field is only present when the authentication 
    // type is DATABRICKS. 
    ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType CreateProviderResponseAuthenticationType `json:"authentication_type,omitempty"`
    // [Create,Update:OPT] Description about the provider. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of Provider creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create,Update:REQ] The name of the Provider. 
    Name string `json:"name,omitempty"`
    // [Create,Update:IGN] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
    // [Create,Update:OPT] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
    // [Create,Update:IGN] The server-generated one-time sharing code. This 
    // field is only present when the authentication type is DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create,Update:IGN] Username of user who last modified Share. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type CreateProviderResponseAuthenticationType string


const CreateProviderResponseAuthenticationTypeUnknown CreateProviderResponseAuthenticationType = `UNKNOWN`

const CreateProviderResponseAuthenticationTypeToken CreateProviderResponseAuthenticationType = `TOKEN`

const CreateProviderResponseAuthenticationTypeDatabricks CreateProviderResponseAuthenticationType = `DATABRICKS`

type CreateRecipientRequest struct {
    // [Create:IGN,Update:IGN] A boolean status field showing whether the 
    // Recipient&#39;s activation URL has been exercised or not. 
    Activated bool `json:"activated,omitempty"`
    // [Create:IGN,Update:IGN] Full activation url to retrieve the access 
    // token. It will be empty if the token is already retrieved. 
    ActivationUrl string `json:"activation_url,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType CreateRecipientRequestAuthenticationType `json:"authentication_type,omitempty"`
    // [Create:OPT,Update:OPT] Description about the recipient. 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN,Update:IGN] Time at which this Recipient was created, in 
    // epoch milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:OPT,Update:OPT] IP Access List 
    IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
    // [Create:REQ,Update:OPT] Name of Recipient. 
    Name string `json:"name,omitempty"`
    // [Create:OPT,Update:IGN] The one-time sharing code provided by the data 
    // recipient. This field is only present when the authentication type is 
    // DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create:IGN,Update:IGN] Recipient Tokens This field is only present when 
    // the authentication type is TOKEN. 
    Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
    // [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in 
    // epoch milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient updater. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type CreateRecipientRequestAuthenticationType string


const CreateRecipientRequestAuthenticationTypeUnknown CreateRecipientRequestAuthenticationType = `UNKNOWN`

const CreateRecipientRequestAuthenticationTypeToken CreateRecipientRequestAuthenticationType = `TOKEN`

const CreateRecipientRequestAuthenticationTypeDatabricks CreateRecipientRequestAuthenticationType = `DATABRICKS`

type CreateRecipientResponse struct {
    // [Create:IGN,Update:IGN] A boolean status field showing whether the 
    // Recipient&#39;s activation URL has been exercised or not. 
    Activated bool `json:"activated,omitempty"`
    // [Create:IGN,Update:IGN] Full activation url to retrieve the access 
    // token. It will be empty if the token is already retrieved. 
    ActivationUrl string `json:"activation_url,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType CreateRecipientResponseAuthenticationType `json:"authentication_type,omitempty"`
    // [Create:OPT,Update:OPT] Description about the recipient. 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN,Update:IGN] Time at which this Recipient was created, in 
    // epoch milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:OPT,Update:OPT] IP Access List 
    IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
    // [Create:REQ,Update:OPT] Name of Recipient. 
    Name string `json:"name,omitempty"`
    // [Create:OPT,Update:IGN] The one-time sharing code provided by the data 
    // recipient. This field is only present when the authentication type is 
    // DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create:IGN,Update:IGN] Recipient Tokens This field is only present when 
    // the authentication type is TOKEN. 
    Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
    // [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in 
    // epoch milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient updater. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type CreateRecipientResponseAuthenticationType string


const CreateRecipientResponseAuthenticationTypeUnknown CreateRecipientResponseAuthenticationType = `UNKNOWN`

const CreateRecipientResponseAuthenticationTypeToken CreateRecipientResponseAuthenticationType = `TOKEN`

const CreateRecipientResponseAuthenticationTypeDatabricks CreateRecipientResponseAuthenticationType = `DATABRICKS`

type CreateShareRequest struct {
    // [Create: OPT] comment when creating the share 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN] Time at which this Share was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN] Username of Share creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:REQ] Name of the Share. 
    Name string `json:"name,omitempty"`
    // [Create: IGN] A list of shared data objects within the Share. 
    Objects []SharedDataObject `json:"objects,omitempty"`
}


type CreateShareResponse struct {
    // [Create: OPT] comment when creating the share 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN] Time at which this Share was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN] Username of Share creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:REQ] Name of the Share. 
    Name string `json:"name,omitempty"`
    // [Create: IGN] A list of shared data objects within the Share. 
    Objects []SharedDataObject `json:"objects,omitempty"`
}


type DeleteProviderRequest struct {
    // Required. Name of the provider. 
    NameArg string ` path:"name_arg"`
}


type DeleteRecipientRequest struct {
    // Required. Name of the recipient. 
    Name string ` path:"name"`
}


type DeleteShareRequest struct {
    
    Name string ` path:"name"`
}


type GetActivationUrlInfoRequest struct {
    // Required. The one time activation url. It also accepts activation token. 
    ActivationUrl string ` path:"activation_url"`
}


type GetProviderRequest struct {
    // Required. Name of the provider. 
    NameArg string ` path:"name_arg"`
}


type GetProviderResponse struct {
    // [Create,Update:IGN] Whether this provider is successfully activated by 
    // the data provider. This field is only present when the authentication 
    // type is DATABRICKS. 
    ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType GetProviderResponseAuthenticationType `json:"authentication_type,omitempty"`
    // [Create,Update:OPT] Description about the provider. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of Provider creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create,Update:REQ] The name of the Provider. 
    Name string `json:"name,omitempty"`
    // [Create,Update:IGN] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
    // [Create,Update:OPT] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
    // [Create,Update:IGN] The server-generated one-time sharing code. This 
    // field is only present when the authentication type is DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create,Update:IGN] Username of user who last modified Share. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type GetProviderResponseAuthenticationType string


const GetProviderResponseAuthenticationTypeUnknown GetProviderResponseAuthenticationType = `UNKNOWN`

const GetProviderResponseAuthenticationTypeToken GetProviderResponseAuthenticationType = `TOKEN`

const GetProviderResponseAuthenticationTypeDatabricks GetProviderResponseAuthenticationType = `DATABRICKS`

type GetRecipientRequest struct {
    // Required. Name of the recipient. 
    Name string ` path:"name"`
}


type GetRecipientResponse struct {
    // [Create:IGN,Update:IGN] A boolean status field showing whether the 
    // Recipient&#39;s activation URL has been exercised or not. 
    Activated bool `json:"activated,omitempty"`
    // [Create:IGN,Update:IGN] Full activation url to retrieve the access 
    // token. It will be empty if the token is already retrieved. 
    ActivationUrl string `json:"activation_url,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType GetRecipientResponseAuthenticationType `json:"authentication_type,omitempty"`
    // [Create:OPT,Update:OPT] Description about the recipient. 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN,Update:IGN] Time at which this Recipient was created, in 
    // epoch milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:OPT,Update:OPT] IP Access List 
    IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
    // [Create:REQ,Update:OPT] Name of Recipient. 
    Name string `json:"name,omitempty"`
    // [Create:OPT,Update:IGN] The one-time sharing code provided by the data 
    // recipient. This field is only present when the authentication type is 
    // DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create:IGN,Update:IGN] Recipient Tokens This field is only present when 
    // the authentication type is TOKEN. 
    Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
    // [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in 
    // epoch milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient updater. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type GetRecipientResponseAuthenticationType string


const GetRecipientResponseAuthenticationTypeUnknown GetRecipientResponseAuthenticationType = `UNKNOWN`

const GetRecipientResponseAuthenticationTypeToken GetRecipientResponseAuthenticationType = `TOKEN`

const GetRecipientResponseAuthenticationTypeDatabricks GetRecipientResponseAuthenticationType = `DATABRICKS`

type GetRecipientSharePermissionsRequest struct {
    // Required. The name of the Recipient. 
    Name string ` path:"name"`
}


type GetRecipientSharePermissionsResponse struct {
    
    PermissionsOut []ShareToPrivilegeAssignment `json:"permissions_out,omitempty"`
}


type GetSharePermissionsRequest struct {
    // Required. The name of the Share. 
    Name string ` path:"name"`
}


type GetSharePermissionsResponse struct {
    // Note to self (acain): Unfortunately, neither json_inline nor json_map 
    // work here. 
    PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}


type GetShareRequest struct {
    
    IncludeSharedData bool ` query:"include_shared_data,omitempty"`
    
    Name string ` path:"name"`
}


type GetShareResponse struct {
    // [Create: OPT] comment when creating the share 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN] Time at which this Share was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN] Username of Share creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:REQ] Name of the Share. 
    Name string `json:"name,omitempty"`
    // [Create: IGN] A list of shared data objects within the Share. 
    Objects []SharedDataObject `json:"objects,omitempty"`
}


type IpAccessList struct {
    // Allowed IP Addresses in CIDR notation. Limit of 100. 
    AllowedIpAddresses []string `json:"allowed_ip_addresses,omitempty"`
}


type ListProviderSharesRequest struct {
    // Required. Name of the provider in which to list shares. 
    ProviderNameArg string ` path:"provider_name_arg"`
}


type ListProviderSharesResponse struct {
    
    Shares []ProviderShare `json:"shares,omitempty"`
}


type ListProvidersResponse struct {
    
    Providers []ProviderInfo `json:"providers,omitempty"`
}


type ListRecipientsResponse struct {
    
    Recipients []RecipientInfo `json:"recipients,omitempty"`
}


type ListSharesResponse struct {
    
    Shares []ShareInfo `json:"shares,omitempty"`
}


type Partition struct {
    
    Values []PartitionValue `json:"values,omitempty"`
}


type PartitionValue struct {
    // The name of the partition column. 
    Name string `json:"name,omitempty"`
    // The operator to apply for the value. 
    Op PartitionValueOp `json:"op,omitempty"`
    // The value of the partition column. When this value is not set, it means 
    // `null` value. 
    Value string `json:"value,omitempty"`
}

// The operator to apply for the value. 
type PartitionValueOp string


const PartitionValueOpEqual PartitionValueOp = `EQUAL`

const PartitionValueOpLike PartitionValueOp = `LIKE`

type PermissionsChange struct {
    // The set of privileges to add. 
    Add []PermissionsChangeAddItem `json:"add,omitempty"`
    // The principal whose privileges we are changing. 
    Principal string `json:"principal,omitempty"`
    // The set of privileges to remove. 
    Remove []PermissionsChangeRemoveItem `json:"remove,omitempty"`
}


type PermissionsChangeAddItem string


const PermissionsChangeAddItemUnknownPrivilege PermissionsChangeAddItem = `UNKNOWN_PRIVILEGE`

const PermissionsChangeAddItemSelect PermissionsChangeAddItem = `SELECT`

const PermissionsChangeAddItemCreate PermissionsChangeAddItem = `CREATE`

const PermissionsChangeAddItemModify PermissionsChangeAddItem = `MODIFY`

const PermissionsChangeAddItemUsage PermissionsChangeAddItem = `USAGE`

const PermissionsChangeAddItemReadFiles PermissionsChangeAddItem = `READ_FILES`

const PermissionsChangeAddItemWriteFiles PermissionsChangeAddItem = `WRITE_FILES`

const PermissionsChangeAddItemCreateTable PermissionsChangeAddItem = `CREATE_TABLE`

const PermissionsChangeAddItemCreateMount PermissionsChangeAddItem = `CREATE_MOUNT`

type PermissionsChangeRemoveItem string


const PermissionsChangeRemoveItemUnknownPrivilege PermissionsChangeRemoveItem = `UNKNOWN_PRIVILEGE`

const PermissionsChangeRemoveItemSelect PermissionsChangeRemoveItem = `SELECT`

const PermissionsChangeRemoveItemCreate PermissionsChangeRemoveItem = `CREATE`

const PermissionsChangeRemoveItemModify PermissionsChangeRemoveItem = `MODIFY`

const PermissionsChangeRemoveItemUsage PermissionsChangeRemoveItem = `USAGE`

const PermissionsChangeRemoveItemReadFiles PermissionsChangeRemoveItem = `READ_FILES`

const PermissionsChangeRemoveItemWriteFiles PermissionsChangeRemoveItem = `WRITE_FILES`

const PermissionsChangeRemoveItemCreateTable PermissionsChangeRemoveItem = `CREATE_TABLE`

const PermissionsChangeRemoveItemCreateMount PermissionsChangeRemoveItem = `CREATE_MOUNT`

type PrivilegeAssignment struct {
    // The principal (user email address or group name). 
    Principal string `json:"principal,omitempty"`
    // The privileges assigned to the principal. 
    Privileges []PrivilegeAssignmentPrivilegesItem `json:"privileges,omitempty"`
}


type PrivilegeAssignmentPrivilegesItem string


const PrivilegeAssignmentPrivilegesItemUnknownPrivilege PrivilegeAssignmentPrivilegesItem = `UNKNOWN_PRIVILEGE`

const PrivilegeAssignmentPrivilegesItemSelect PrivilegeAssignmentPrivilegesItem = `SELECT`

const PrivilegeAssignmentPrivilegesItemCreate PrivilegeAssignmentPrivilegesItem = `CREATE`

const PrivilegeAssignmentPrivilegesItemModify PrivilegeAssignmentPrivilegesItem = `MODIFY`

const PrivilegeAssignmentPrivilegesItemUsage PrivilegeAssignmentPrivilegesItem = `USAGE`

const PrivilegeAssignmentPrivilegesItemReadFiles PrivilegeAssignmentPrivilegesItem = `READ_FILES`

const PrivilegeAssignmentPrivilegesItemWriteFiles PrivilegeAssignmentPrivilegesItem = `WRITE_FILES`

const PrivilegeAssignmentPrivilegesItemCreateTable PrivilegeAssignmentPrivilegesItem = `CREATE_TABLE`

const PrivilegeAssignmentPrivilegesItemCreateMount PrivilegeAssignmentPrivilegesItem = `CREATE_MOUNT`

type ProviderInfo struct {
    // [Create,Update:IGN] Whether this provider is successfully activated by 
    // the data provider. This field is only present when the authentication 
    // type is DATABRICKS. 
    ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType ProviderInfoAuthenticationType `json:"authentication_type,omitempty"`
    // [Create,Update:OPT] Description about the provider. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of Provider creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create,Update:REQ] The name of the Provider. 
    Name string `json:"name,omitempty"`
    // [Create,Update:IGN] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
    // [Create,Update:OPT] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
    // [Create,Update:IGN] The server-generated one-time sharing code. This 
    // field is only present when the authentication type is DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create,Update:IGN] Username of user who last modified Share. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type ProviderInfoAuthenticationType string


const ProviderInfoAuthenticationTypeUnknown ProviderInfoAuthenticationType = `UNKNOWN`

const ProviderInfoAuthenticationTypeToken ProviderInfoAuthenticationType = `TOKEN`

const ProviderInfoAuthenticationTypeDatabricks ProviderInfoAuthenticationType = `DATABRICKS`

type ProviderShare struct {
    // The name of the Provider Share. 
    Name string `json:"name,omitempty"`
}


type RecipientInfo struct {
    // [Create:IGN,Update:IGN] A boolean status field showing whether the 
    // Recipient&#39;s activation URL has been exercised or not. 
    Activated bool `json:"activated,omitempty"`
    // [Create:IGN,Update:IGN] Full activation url to retrieve the access 
    // token. It will be empty if the token is already retrieved. 
    ActivationUrl string `json:"activation_url,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType RecipientInfoAuthenticationType `json:"authentication_type,omitempty"`
    // [Create:OPT,Update:OPT] Description about the recipient. 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN,Update:IGN] Time at which this Recipient was created, in 
    // epoch milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:OPT,Update:OPT] IP Access List 
    IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
    // [Create:REQ,Update:OPT] Name of Recipient. 
    Name string `json:"name,omitempty"`
    // [Create:OPT,Update:IGN] The one-time sharing code provided by the data 
    // recipient. This field is only present when the authentication type is 
    // DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create:IGN,Update:IGN] Recipient Tokens This field is only present when 
    // the authentication type is TOKEN. 
    Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
    // [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in 
    // epoch milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient updater. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type RecipientInfoAuthenticationType string


const RecipientInfoAuthenticationTypeUnknown RecipientInfoAuthenticationType = `UNKNOWN`

const RecipientInfoAuthenticationTypeToken RecipientInfoAuthenticationType = `TOKEN`

const RecipientInfoAuthenticationTypeDatabricks RecipientInfoAuthenticationType = `DATABRICKS`

type RecipientProfile struct {
    
    BearerToken string `json:"bearer_token,omitempty"`
    
    Endpoint string `json:"endpoint,omitempty"`
    
    ShareCredentialsVersion int `json:"share_credentials_version,omitempty"`
}


type RecipientTokenInfo struct {
    // Full activation url to retrieve the access token. It will be empty if 
    // the token is already retrieved. 
    ActivationUrl string `json:"activation_url,omitempty"`
    // Time at which this Recipient Token was created, in epoch milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // Username of Recipient Token creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // Expiration timestamp of the token in epoch milliseconds. 
    ExpirationTime int64 `json:"expiration_time,omitempty"`
    // Unique id of the Recipient Token. 
    Id string `json:"id,omitempty"`
    // Time at which this Recipient Token was updated, in epoch milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // Username of Recipient Token updater. 
    UpdatedBy string `json:"updated_by,omitempty"`
}


type RetrieveTokenRequest struct {
    // Required. The one time activation url. It also accepts activation token. 
    ActivationUrl string ` path:"activation_url"`
}


type RetrieveTokenResponse struct {
    
    BearerToken string `json:"bearerToken,omitempty"`
    
    Endpoint string `json:"endpoint,omitempty"`
    
    ExpirationTime string `json:"expirationTime,omitempty"`
    // These field names must follow the delta sharing protocol. 
    ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`
}


type RotateRecipientTokenRequest struct {
    // Required. This will set the expiration_time of existing token only to a 
    // smaller timestamp, it cannot extend the expiration_time. Use 0 to expire 
    // the existing token immediately, negative number will return an error. 
    ExistingTokenExpireInSeconds int64 `json:"existing_token_expire_in_seconds,omitempty"`
    // Required. The name of the Recipient. 
    Name string ` path:"name"`
}


type RotateRecipientTokenResponse struct {
    // [Create:IGN,Update:IGN] A boolean status field showing whether the 
    // Recipient&#39;s activation URL has been exercised or not. 
    Activated bool `json:"activated,omitempty"`
    // [Create:IGN,Update:IGN] Full activation url to retrieve the access 
    // token. It will be empty if the token is already retrieved. 
    ActivationUrl string `json:"activation_url,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType RotateRecipientTokenResponseAuthenticationType `json:"authentication_type,omitempty"`
    // [Create:OPT,Update:OPT] Description about the recipient. 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN,Update:IGN] Time at which this Recipient was created, in 
    // epoch milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:OPT,Update:OPT] IP Access List 
    IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
    // [Create:REQ,Update:OPT] Name of Recipient. 
    Name string `json:"name,omitempty"`
    // [Create:OPT,Update:IGN] The one-time sharing code provided by the data 
    // recipient. This field is only present when the authentication type is 
    // DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create:IGN,Update:IGN] Recipient Tokens This field is only present when 
    // the authentication type is TOKEN. 
    Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
    // [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in 
    // epoch milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient updater. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type RotateRecipientTokenResponseAuthenticationType string


const RotateRecipientTokenResponseAuthenticationTypeUnknown RotateRecipientTokenResponseAuthenticationType = `UNKNOWN`

const RotateRecipientTokenResponseAuthenticationTypeToken RotateRecipientTokenResponseAuthenticationType = `TOKEN`

const RotateRecipientTokenResponseAuthenticationTypeDatabricks RotateRecipientTokenResponseAuthenticationType = `DATABRICKS`

type ShareInfo struct {
    // [Create: OPT] comment when creating the share 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN] Time at which this Share was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN] Username of Share creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:REQ] Name of the Share. 
    Name string `json:"name,omitempty"`
    // [Create: IGN] A list of shared data objects within the Share. 
    Objects []SharedDataObject `json:"objects,omitempty"`
}


type ShareToPrivilegeAssignment struct {
    // The privileges assigned to the principal. 
    PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
    // The share name. 
    ShareName string `json:"share_name,omitempty"`
}


type SharedDataObject struct {
    // The time when this data object is added to the Share, in epoch 
    // milliseconds. Output only field. [Update:IGN] 
    AddedAt int64 `json:"added_at,omitempty"`
    // Username of the sharer. Output only field. [Update:IGN] 
    AddedBy string `json:"added_by,omitempty"`
    // A user-provided comment when adding the data object to the share. 
    // [Update:OPT] 
    Comment string `json:"comment,omitempty"`
    // The type of the data object. Output only field. [Update:IGN] 
    DataObjectType string `json:"data_object_type,omitempty"`
    // A fully qualified name that uniquely identifies a data object. For 
    // example, a table&#39;s fully qualified name is in the format of 
    // `&lt;catalog&gt;.&lt;schema&gt;.&lt;table&gt;`. [Update:REQ] 
    Name string `json:"name,omitempty"`
    
    Partitions []Partition `json:"partitions,omitempty"`
    // A user-provided new name for the data object within the share. If this 
    // new name is not not provided, the object&#39;s original name will be used as 
    // the `shared_as` name. The `shared_as` name must be unique within a 
    // Share. For tables, the new name must follow the format of 
    // `&lt;schema&gt;.&lt;table&gt;`. [Update:OPT] 
    SharedAs string `json:"shared_as,omitempty"`
}


type SharedDataObjectUpdate struct {
    
    Action SharedDataObjectUpdateAction `json:"action,omitempty"`
    // The data object that is being updated (added / removed). 
    DataObject *SharedDataObject `json:"data_object,omitempty"`
}


type SharedDataObjectUpdateAction string


const SharedDataObjectUpdateActionAdd SharedDataObjectUpdateAction = `ADD`

const SharedDataObjectUpdateActionRemove SharedDataObjectUpdateAction = `REMOVE`

type UpdateProviderRequest struct {
    // [Create,Update:IGN] Whether this provider is successfully activated by 
    // the data provider. This field is only present when the authentication 
    // type is DATABRICKS. 
    ActivatedByProvider bool `json:"activated_by_provider,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType UpdateProviderRequestAuthenticationType `json:"authentication_type,omitempty"`
    // [Create,Update:OPT] Description about the provider. 
    Comment string `json:"comment,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create,Update:IGN] Username of Provider creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create,Update:REQ] The name of the Provider. 
    Name string `json:"name,omitempty"`
    // Required. Name of the provider. 
    NameArg string ` path:"name_arg"`
    // [Create,Update:IGN] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
    // [Create,Update:OPT] This field is only present when the authentication 
    // type is TOKEN. 
    RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
    // [Create,Update:IGN] The server-generated one-time sharing code. This 
    // field is only present when the authentication type is DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create,Update:IGN] Time at which this Provider was created, in epoch 
    // milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create,Update:IGN] Username of user who last modified Share. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type UpdateProviderRequestAuthenticationType string


const UpdateProviderRequestAuthenticationTypeUnknown UpdateProviderRequestAuthenticationType = `UNKNOWN`

const UpdateProviderRequestAuthenticationTypeToken UpdateProviderRequestAuthenticationType = `TOKEN`

const UpdateProviderRequestAuthenticationTypeDatabricks UpdateProviderRequestAuthenticationType = `DATABRICKS`

type UpdateRecipientRequest struct {
    // [Create:IGN,Update:IGN] A boolean status field showing whether the 
    // Recipient&#39;s activation URL has been exercised or not. 
    Activated bool `json:"activated,omitempty"`
    // [Create:IGN,Update:IGN] Full activation url to retrieve the access 
    // token. It will be empty if the token is already retrieved. 
    ActivationUrl string `json:"activation_url,omitempty"`
    // [Create:REQ,Update:IGN] The delta sharing authentication type. 
    AuthenticationType UpdateRecipientRequestAuthenticationType `json:"authentication_type,omitempty"`
    // [Create:OPT,Update:OPT] Description about the recipient. 
    Comment string `json:"comment,omitempty"`
    // [Create:IGN,Update:IGN] Time at which this Recipient was created, in 
    // epoch milliseconds. 
    CreatedAt int64 `json:"created_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient creator. 
    CreatedBy string `json:"created_by,omitempty"`
    // [Create:OPT,Update:OPT] IP Access List 
    IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
    // [Create:REQ,Update:OPT] Name of Recipient. 
    Name string `json:"name,omitempty"`
    // Required. Name of the recipient. 
    NameArg string ` path:"name_arg"`
    // [Create:OPT,Update:IGN] The one-time sharing code provided by the data 
    // recipient. This field is only present when the authentication type is 
    // DATABRICKS. 
    SharingCode string `json:"sharing_code,omitempty"`
    // [Create:IGN,Update:IGN] Recipient Tokens This field is only present when 
    // the authentication type is TOKEN. 
    Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
    // [Create:IGN,Update:IGN] Time at which thie Recipient was updated, in 
    // epoch milliseconds. 
    UpdatedAt int64 `json:"updated_at,omitempty"`
    // [Create:IGN,Update:IGN] Username of Recipient updater. 
    UpdatedBy string `json:"updated_by,omitempty"`
}

// [Create:REQ,Update:IGN] The delta sharing authentication type. 
type UpdateRecipientRequestAuthenticationType string


const UpdateRecipientRequestAuthenticationTypeUnknown UpdateRecipientRequestAuthenticationType = `UNKNOWN`

const UpdateRecipientRequestAuthenticationTypeToken UpdateRecipientRequestAuthenticationType = `TOKEN`

const UpdateRecipientRequestAuthenticationTypeDatabricks UpdateRecipientRequestAuthenticationType = `DATABRICKS`

type UpdateSharePermissionsRequest struct {
    
    Changes []PermissionsChange `json:"changes,omitempty"`
    // Required. The name of the Share. 
    Name string ` path:"name"`
}


type UpdateShareRequest struct {
    
    Name string ` path:"name"`
    
    Updates []SharedDataObjectUpdate `json:"updates,omitempty"`
}

