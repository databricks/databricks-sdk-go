// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// all definitions in this file are in alphabetical order

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type"`
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

type CreateNetworkConnectivityConfigRequest struct {
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name string `json:"name"`
	// The Azure region for this network connectivity configuration. Only
	// workspaces in the same Azure region can be attached to this network
	// connectivity configuration.
	Region string `json:"region"`
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	ApplicationId string `json:"application_id"`
	// Comment that describes the purpose of the token.
	Comment string `json:"comment,omitempty"`
	// The number of seconds before the token expires.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateOboTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
	// Value of the token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateOboTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePrivateEndpointRuleRequest struct {
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId CreatePrivateEndpointRuleRequestGroupId `json:"group_id"`
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// The Azure resource ID of the target resource.
	ResourceId string `json:"resource_id"`
}

// The sub-resource type (group ID) of the target resource. Note that to connect
// to workspace root storage (root DBFS), you need two endpoints, one for `blob`
// and one for `dfs`.
type CreatePrivateEndpointRuleRequestGroupId string

const CreatePrivateEndpointRuleRequestGroupIdBlob CreatePrivateEndpointRuleRequestGroupId = `blob`

const CreatePrivateEndpointRuleRequestGroupIdDfs CreatePrivateEndpointRuleRequestGroupId = `dfs`

const CreatePrivateEndpointRuleRequestGroupIdMysqlServer CreatePrivateEndpointRuleRequestGroupId = `mysqlServer`

const CreatePrivateEndpointRuleRequestGroupIdSqlServer CreatePrivateEndpointRuleRequestGroupId = `sqlServer`

// String representation for [fmt.Print]
func (f *CreatePrivateEndpointRuleRequestGroupId) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreatePrivateEndpointRuleRequestGroupId) Set(v string) error {
	switch v {
	case `blob`, `dfs`, `mysqlServer`, `sqlServer`:
		*f = CreatePrivateEndpointRuleRequestGroupId(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "blob", "dfs", "mysqlServer", "sqlServer"`, v)
	}
}

// Type always returns CreatePrivateEndpointRuleRequestGroupId to satisfy [pflag.Value] interface
func (f *CreatePrivateEndpointRuleRequestGroupId) Type() string {
	return "CreatePrivateEndpointRuleRequestGroupId"
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	Comment string `json:"comment,omitempty"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo *PublicTokenInfo `json:"token_info,omitempty"`
	// The value of the new token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This represents the setting configuration for the default namespace in the
// Databricks workspace. Setting the default catalog for the workspace
// determines the catalog that is used when queries do not reference a fully
// qualified 3 level name. For example, if the default catalog is set to
// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the object
// 'retail_prod.default.myTable' (the schema 'default' is always assumed). This
// setting requires a restart of clusters and SQL warehouses to take effect.
// Additionally, the default namespace only applies when using Unity
// Catalog-enabled compute.
type DefaultNamespaceSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	Namespace StringMessage `json:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DefaultNamespaceSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DefaultNamespaceSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

// Delete the default namespace setting
type DeleteDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteDefaultNamespaceSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

// Delete Personal Compute setting
type DeletePersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeletePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeletePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeletePersonalComputeSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
}

// Delete the restrict workspace admins setting
type DeleteRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteRestrictWorkspaceAdminsSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken struct {
	// The requested token.
	Credential string `json:"credential,omitempty"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime int64 `json:"credentialEolTime,omitempty"`
	// User ID of the user that owns this token.
	OwnerId int64 `json:"ownerId,omitempty"`
	// The scopes of access granted in the token.
	Scopes []string `json:"scopes,omitempty"`
	// The type of this exchange token
	TokenType TokenType `json:"tokenType,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ExchangeToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	PartitionId PartitionId `json:"partitionId"`
	// Array of scopes for the token request.
	Scopes []string `json:"scopes"`
	// A list of token types being requested
	TokenType []TokenType `json:"tokenType"`
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {
	Values []ExchangeToken `json:"values,omitempty"`
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

// Get the default namespace setting
type GetDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

type GetIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

// Get Personal Compute setting
type GetPersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a private endpoint rule
type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
}

// Get the restrict workspace admins setting
type GetRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Check configuration status
type GetStatusRequest struct {
	Keys string `json:"-" url:"keys"`
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []TokenPermissionsDescription `json:"permission_levels,omitempty"`
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
}

// Definition of an IP Access list
type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	AddressCount int `json:"address_count,omitempty"`
	// Creation timestamp in milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// User ID of the user who created this list.
	CreatedBy int64 `json:"created_by,omitempty"`
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId string `json:"list_id,omitempty"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type,omitempty"`
	// Update timestamp in milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// User ID of the user who updated this list.
	UpdatedBy int64 `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *IpAccessListInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IpAccessListInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

type ListNccAzurePrivateEndpointRulesResponse struct {
	Items []NccAzurePrivateEndpointRule `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListNccAzurePrivateEndpointRulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNccAzurePrivateEndpointRulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListNetworkConnectivityConfigurationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkConnectivityConfigurationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListNetworkConnectivityConfigurationsResponse struct {
	Items []NetworkConnectivityConfiguration `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListNetworkConnectivityConfigurationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkConnectivityConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListPrivateEndpointRulesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPrivateEndpointRulesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos []PublicTokenInfo `json:"token_infos,omitempty"`
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById int64 `json:"-" url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"-" url:"created_by_username,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListTokenManagementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTokenManagementRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

// Type of IP access list. Valid values are as follows and are case-sensitive:
//
// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block list.
// Exclude this IP or range. IP addresses in the block list are excluded even if
// they are included in an allow list.
type ListType string

// An allow list. Include this IP or range.
const ListTypeAllow ListType = `ALLOW`

// A block list. Exclude this IP or range. IP addresses in the block list are
// excluded even if they are included in an allow list.
const ListTypeBlock ListType = `BLOCK`

// String representation for [fmt.Print]
func (f *ListType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListType) Set(v string) error {
	switch v {
	case `ALLOW`, `BLOCK`:
		*f = ListType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW", "BLOCK"`, v)
	}
}

// Type always returns ListType to satisfy [pflag.Value] interface
func (f *ListType) Type() string {
	return "ListType"
}

type NccAzurePrivateEndpointRule struct {
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is `ESTABLISHED`. Remember
	// that you must approve new endpoints on your resources in the Azure portal
	// before they take effect.
	//
	// The possible values are: - INIT: (deprecated) The endpoint has been
	// created and pending approval. - PENDING: The endpoint has been created
	// and pending approval. - ESTABLISHED: The endpoint has been approved and
	// is ready to use in your serverless compute resources. - REJECTED:
	// Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up.
	ConnectionState NccAzurePrivateEndpointRuleConnectionState `json:"connection_state,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Whether this private endpoint is deactivated.
	Deactivated bool `json:"deactivated,omitempty"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt int64 `json:"deactivated_at,omitempty"`
	// The name of the Azure private endpoint resource.
	EndpointName string `json:"endpoint_name,omitempty"`
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId NccAzurePrivateEndpointRuleGroupId `json:"group_id,omitempty"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The Azure resource ID of the target resource.
	ResourceId string `json:"resource_id,omitempty"`
	// The ID of a private endpoint rule.
	RuleId string `json:"rule_id,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *NccAzurePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NccAzurePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The current status of this private endpoint. The private endpoint rules are
// effective only if the connection state is `ESTABLISHED`. Remember that you
// must approve new endpoints on your resources in the Azure portal before they
// take effect.
//
// The possible values are: - INIT: (deprecated) The endpoint has been created
// and pending approval. - PENDING: The endpoint has been created and pending
// approval. - ESTABLISHED: The endpoint has been approved and is ready to use
// in your serverless compute resources. - REJECTED: Connection was rejected by
// the private link resource owner. - DISCONNECTED: Connection was removed by
// the private link resource owner, the private endpoint becomes informative and
// should be deleted for clean-up.
type NccAzurePrivateEndpointRuleConnectionState string

const NccAzurePrivateEndpointRuleConnectionStateDisconnected NccAzurePrivateEndpointRuleConnectionState = `DISCONNECTED`

const NccAzurePrivateEndpointRuleConnectionStateEstablished NccAzurePrivateEndpointRuleConnectionState = `ESTABLISHED`

const NccAzurePrivateEndpointRuleConnectionStateInit NccAzurePrivateEndpointRuleConnectionState = `INIT`

const NccAzurePrivateEndpointRuleConnectionStatePending NccAzurePrivateEndpointRuleConnectionState = `PENDING`

const NccAzurePrivateEndpointRuleConnectionStateRejected NccAzurePrivateEndpointRuleConnectionState = `REJECTED`

// String representation for [fmt.Print]
func (f *NccAzurePrivateEndpointRuleConnectionState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NccAzurePrivateEndpointRuleConnectionState) Set(v string) error {
	switch v {
	case `DISCONNECTED`, `ESTABLISHED`, `INIT`, `PENDING`, `REJECTED`:
		*f = NccAzurePrivateEndpointRuleConnectionState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISCONNECTED", "ESTABLISHED", "INIT", "PENDING", "REJECTED"`, v)
	}
}

// Type always returns NccAzurePrivateEndpointRuleConnectionState to satisfy [pflag.Value] interface
func (f *NccAzurePrivateEndpointRuleConnectionState) Type() string {
	return "NccAzurePrivateEndpointRuleConnectionState"
}

// The sub-resource type (group ID) of the target resource. Note that to connect
// to workspace root storage (root DBFS), you need two endpoints, one for `blob`
// and one for `dfs`.
type NccAzurePrivateEndpointRuleGroupId string

const NccAzurePrivateEndpointRuleGroupIdBlob NccAzurePrivateEndpointRuleGroupId = `blob`

const NccAzurePrivateEndpointRuleGroupIdDfs NccAzurePrivateEndpointRuleGroupId = `dfs`

const NccAzurePrivateEndpointRuleGroupIdMysqlServer NccAzurePrivateEndpointRuleGroupId = `mysqlServer`

const NccAzurePrivateEndpointRuleGroupIdSqlServer NccAzurePrivateEndpointRuleGroupId = `sqlServer`

// String representation for [fmt.Print]
func (f *NccAzurePrivateEndpointRuleGroupId) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NccAzurePrivateEndpointRuleGroupId) Set(v string) error {
	switch v {
	case `blob`, `dfs`, `mysqlServer`, `sqlServer`:
		*f = NccAzurePrivateEndpointRuleGroupId(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "blob", "dfs", "mysqlServer", "sqlServer"`, v)
	}
}

// Type always returns NccAzurePrivateEndpointRuleGroupId to satisfy [pflag.Value] interface
func (f *NccAzurePrivateEndpointRuleGroupId) Type() string {
	return "NccAzurePrivateEndpointRuleGroupId"
}

// The stable Azure service endpoints. You can configure the firewall of your
// Azure resources to allow traffic from your Databricks serverless compute
// resources.
type NccAzureServiceEndpointRule struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets []string `json:"subnets,omitempty"`
	// The Azure region in which this service endpoint rule applies.
	TargetRegion string `json:"target_region,omitempty"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices []string `json:"target_services,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *NccAzureServiceEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NccAzureServiceEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network connectivity rules that apply to network traffic from your
// serverless compute resources.
type NccEgressConfig struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules *NccEgressDefaultRules `json:"default_rules,omitempty"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules *NccEgressTargetRules `json:"target_rules,omitempty"`
}

// The network connectivity rules that are applied by default without resource
// specific configurations. You can find the stable network information of your
// serverless compute resources here.
type NccEgressDefaultRules struct {
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	AzureServiceEndpointRule *NccAzureServiceEndpointRule `json:"azure_service_endpoint_rule,omitempty"`
}

// The network connectivity rules that configured for each destinations. These
// rules override default rules.
type NccEgressTargetRules struct {
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRule `json:"azure_private_endpoint_rules,omitempty"`
}

type NetworkConnectivityConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig *NccEgressConfig `json:"egress_config,omitempty"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name string `json:"name,omitempty"`
	// Databricks network connectivity configuration ID.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The Azure region for this network connectivity configuration. Only
	// workspaces in the same Azure region can be attached to this network
	// connectivity configuration.
	Region string `json:"region,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *NetworkConnectivityConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkConnectivityConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId int64 `json:"workspaceId,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PartitionId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PartitionId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PersonalComputeMessage struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value PersonalComputeMessageEnum `json:"value"`
}

// ON: Grants all users in all workspaces access to the Personal Compute default
// policy, allowing all users to create single-machine compute resources.
// DELEGATE: Moves access control for the Personal Compute default policy to
// individual workspaces and requires a workspace’s users or groups to be
// added to the ACLs of that workspace’s Personal Compute default policy
// before they will be able to create compute resources through that policy.
type PersonalComputeMessageEnum string

const PersonalComputeMessageEnumDelegate PersonalComputeMessageEnum = `DELEGATE`

const PersonalComputeMessageEnumOn PersonalComputeMessageEnum = `ON`

// String representation for [fmt.Print]
func (f *PersonalComputeMessageEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PersonalComputeMessageEnum) Set(v string) error {
	switch v {
	case `DELEGATE`, `ON`:
		*f = PersonalComputeMessageEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELEGATE", "ON"`, v)
	}
}

// Type always returns PersonalComputeMessageEnum to satisfy [pflag.Value] interface
func (f *PersonalComputeMessageEnum) Type() string {
	return "PersonalComputeMessageEnum"
}

type PersonalComputeSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	PersonalCompute PersonalComputeMessage `json:"personal_compute"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PersonalComputeSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PersonalComputeSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublicTokenInfo struct {
	// Comment the token was created with, if applicable.
	Comment string `json:"comment,omitempty"`
	// Server time (in epoch milliseconds) when the token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// The ID of this token.
	TokenId string `json:"token_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PublicTokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublicTokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to replace an IP access list.
type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type"`
}

type RestrictWorkspaceAdminsMessage struct {
	Status RestrictWorkspaceAdminsMessageStatus `json:"status"`
}

type RestrictWorkspaceAdminsMessageStatus string

const RestrictWorkspaceAdminsMessageStatusAllowAll RestrictWorkspaceAdminsMessageStatus = `ALLOW_ALL`

const RestrictWorkspaceAdminsMessageStatusRestrictTokensAndJobRunAs RestrictWorkspaceAdminsMessageStatus = `RESTRICT_TOKENS_AND_JOB_RUN_AS`

const RestrictWorkspaceAdminsMessageStatusStatusUnspecified RestrictWorkspaceAdminsMessageStatus = `STATUS_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *RestrictWorkspaceAdminsMessageStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RestrictWorkspaceAdminsMessageStatus) Set(v string) error {
	switch v {
	case `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`, `STATUS_UNSPECIFIED`:
		*f = RestrictWorkspaceAdminsMessageStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL", "RESTRICT_TOKENS_AND_JOB_RUN_AS", "STATUS_UNSPECIFIED"`, v)
	}
}

// Type always returns RestrictWorkspaceAdminsMessageStatus to satisfy [pflag.Value] interface
func (f *RestrictWorkspaceAdminsMessageStatus) Type() string {
	return "RestrictWorkspaceAdminsMessageStatus"
}

type RestrictWorkspaceAdminsSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessage `json:"restrict_workspace_admins"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RestrictWorkspaceAdminsSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestrictWorkspaceAdminsSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId string `json:"token_id"`
}

type StringMessage struct {
	// Represents a generic string value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *StringMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StringMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions []TokenPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment string `json:"comment,omitempty"`
	// User ID of the user that created the token.
	CreatedById int64 `json:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"created_by_username,omitempty"`
	// Timestamp when the token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Timestamp when the token expires.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// User ID of the user that owns the token.
	OwnerId int64 `json:"owner_id,omitempty"`
	// ID of the token.
	TokenId string `json:"token_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type TokenPermissionLevel string

const TokenPermissionLevelCanUse TokenPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *TokenPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_USE`:
		*f = TokenPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_USE"`, v)
	}
}

// Type always returns TokenPermissionLevel to satisfy [pflag.Value] interface
func (f *TokenPermissionLevel) Type() string {
	return "TokenPermissionLevel"
}

type TokenPermissions struct {
	AccessControlList []TokenAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsRequest struct {
	AccessControlList []TokenAccessControlRequest `json:"access_control_list,omitempty"`
}

// The type of token request. As of now, only `AZURE_ACTIVE_DIRECTORY_TOKEN` is
// supported.
type TokenType string

const TokenTypeAzureActiveDirectoryToken TokenType = `AZURE_ACTIVE_DIRECTORY_TOKEN`

// String representation for [fmt.Print]
func (f *TokenType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenType) Set(v string) error {
	switch v {
	case `AZURE_ACTIVE_DIRECTORY_TOKEN`:
		*f = TokenType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AZURE_ACTIVE_DIRECTORY_TOKEN"`, v)
	}
}

// Type always returns TokenType to satisfy [pflag.Value] interface
func (f *TokenType) Type() string {
	return "TokenType"
}

// Details required to update a setting.
type UpdateDefaultNamespaceSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `json:"field_mask"`
	// This represents the setting configuration for the default namespace in
	// the Databricks workspace. Setting the default catalog for the workspace
	// determines the catalog that is used when queries do not reference a fully
	// qualified 3 level name. For example, if the default catalog is set to
	// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the
	// object 'retail_prod.default.myTable' (the schema 'default' is always
	// assumed). This setting requires a restart of clusters and SQL warehouses
	// to take effect. Additionally, the default namespace only applies when
	// using Unity Catalog-enabled compute.
	Setting DefaultNamespaceSetting `json:"setting"`
}

// Details required to update an IP access list.
type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateIpAccessList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateIpAccessList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `json:"field_mask"`

	Setting PersonalComputeSetting `json:"setting"`
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	FieldMask string `json:"field_mask"`

	Setting RestrictWorkspaceAdminsSetting `json:"setting"`
}

type WorkspaceConf map[string]string
