// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package disasterrecovery

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateFailoverGroupRequest struct {
	// The failover group to create.
	FailoverGroup FailoverGroup `json:"failover_group"`
	// Client-provided identifier for the failover group. Used to construct the
	// resource name as {parent}/failover-groups/{failover_group_id}.
	FailoverGroupId string `json:"-" url:"failover_group_id"`
	// The parent resource. Format: accounts/{account_id}.
	Parent string `json:"-" url:"-"`
	// When true, validates the request without creating the failover group.
	ValidateOnly bool `json:"-" url:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateFailoverGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateFailoverGroupRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateStableUrlRequest struct {
	// The parent resource. Format: accounts/{account_id}.
	Parent string `json:"-" url:"-"`
	// The stable URL to create.
	StableUrl StableUrl `json:"stable_url"`
	// Client-provided identifier for the stable URL. Used to construct the
	// resource name as {parent}/stable-urls/{stable_url_id}.
	StableUrlId string `json:"-" url:"stable_url_id"`
	// When true, validates the request without creating the stable URL.
	ValidateOnly bool `json:"-" url:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateStableUrlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateStableUrlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteFailoverGroupRequest struct {
	// Opaque version string for optimistic locking. If provided, must match the
	// current etag. If omitted, the delete proceeds without an etag check.
	Etag string `json:"-" url:"etag,omitempty"`
	// The fully qualified resource name of the failover group to delete.
	// Format: accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteFailoverGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteFailoverGroupRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteStableUrlRequest struct {
	// The fully qualified resource name. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}.
	Name string `json:"-" url:"-"`
}

// Request to failover a failover group to a new primary region.
type FailoverFailoverGroupRequest struct {
	// Opaque version string for optimistic locking. If provided, must match the
	// current etag. If omitted, the failover proceeds regardless of current
	// state.
	Etag string `json:"etag,omitempty"`
	// The type of failover to perform.
	FailoverType FailoverFailoverGroupRequestFailoverType `json:"failover_type"`
	// The fully qualified resource name of the failover group to failover.
	// Format: accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"-" url:"-"`
	// The target primary region. Must be one of the derived regions and
	// different from the current effective_primary_region. Serves as an
	// idempotency check.
	TargetPrimaryRegion string `json:"target_primary_region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FailoverFailoverGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailoverFailoverGroupRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of failover to perform.
type FailoverFailoverGroupRequestFailoverType string

const FailoverFailoverGroupRequestFailoverTypeForced FailoverFailoverGroupRequestFailoverType = `FORCED`

// String representation for [fmt.Print]
func (f *FailoverFailoverGroupRequestFailoverType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FailoverFailoverGroupRequestFailoverType) Set(v string) error {
	switch v {
	case `FORCED`:
		*f = FailoverFailoverGroupRequestFailoverType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FORCED"`, v)
	}
}

// Values returns all possible values for FailoverFailoverGroupRequestFailoverType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FailoverFailoverGroupRequestFailoverType) Values() []FailoverFailoverGroupRequestFailoverType {
	return []FailoverFailoverGroupRequestFailoverType{
		FailoverFailoverGroupRequestFailoverTypeForced,
	}
}

// Type always returns FailoverFailoverGroupRequestFailoverType to satisfy [pflag.Value] interface
func (f *FailoverFailoverGroupRequestFailoverType) Type() string {
	return "FailoverFailoverGroupRequestFailoverType"
}

// A failover group manages disaster recovery across workspace sets,
// coordinating UCDR and CPDR replication.
type FailoverGroup struct {
	// Time at which this failover group was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Current effective primary region. Replication flows FROM workspaces in
	// this region. Changes after a successful failover.
	EffectivePrimaryRegion string `json:"effective_primary_region,omitempty"`
	// Opaque version string for optimistic locking. Server-generated, returned
	// in responses. Must be provided on Update requests to prevent concurrent
	// modifications.
	Etag string `json:"etag,omitempty"`
	// Initial primary region. Used only in Create requests to set the starting
	// primary region. Not returned in responses.
	InitialPrimaryRegion string `json:"initial_primary_region"`
	// Fully qualified resource name in the format
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"name,omitempty"`
	// List of all regions participating in this failover group.
	Regions []string `json:"regions"`
	// The latest point in time to which data has been replicated.
	ReplicationPoint *time.Time `json:"replication_point,omitempty"`
	// Aggregate state of the failover group.
	State FailoverGroupState `json:"state,omitempty"`
	// Unity Catalog replication configuration.
	UnityCatalogAssets *UcReplicationConfig `json:"unity_catalog_assets,omitempty"`
	// Time at which this failover group was last modified.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Workspace sets, each containing workspaces that replicate to each other.
	WorkspaceSets []WorkspaceSet `json:"workspace_sets"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FailoverGroup) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FailoverGroup) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The aggregate state of a FailoverGroup.
type FailoverGroupState string

const FailoverGroupStateActive FailoverGroupState = `ACTIVE`

const FailoverGroupStateCreating FailoverGroupState = `CREATING`

const FailoverGroupStateCreationFailed FailoverGroupState = `CREATION_FAILED`

const FailoverGroupStateDeleting FailoverGroupState = `DELETING`

const FailoverGroupStateDeletionFailed FailoverGroupState = `DELETION_FAILED`

const FailoverGroupStateFailingOver FailoverGroupState = `FAILING_OVER`

const FailoverGroupStateFailoverFailed FailoverGroupState = `FAILOVER_FAILED`

const FailoverGroupStateInitialReplication FailoverGroupState = `INITIAL_REPLICATION`

// String representation for [fmt.Print]
func (f *FailoverGroupState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FailoverGroupState) Set(v string) error {
	switch v {
	case `ACTIVE`, `CREATING`, `CREATION_FAILED`, `DELETING`, `DELETION_FAILED`, `FAILING_OVER`, `FAILOVER_FAILED`, `INITIAL_REPLICATION`:
		*f = FailoverGroupState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "CREATING", "CREATION_FAILED", "DELETING", "DELETION_FAILED", "FAILING_OVER", "FAILOVER_FAILED", "INITIAL_REPLICATION"`, v)
	}
}

// Values returns all possible values for FailoverGroupState.
//
// There is no guarantee on the order of the values in the slice.
func (f *FailoverGroupState) Values() []FailoverGroupState {
	return []FailoverGroupState{
		FailoverGroupStateActive,
		FailoverGroupStateCreating,
		FailoverGroupStateCreationFailed,
		FailoverGroupStateDeleting,
		FailoverGroupStateDeletionFailed,
		FailoverGroupStateFailingOver,
		FailoverGroupStateFailoverFailed,
		FailoverGroupStateInitialReplication,
	}
}

// Type always returns FailoverGroupState to satisfy [pflag.Value] interface
func (f *FailoverGroupState) Type() string {
	return "FailoverGroupState"
}

type GetFailoverGroupRequest struct {
	// The fully qualified resource name of the failover group. Format:
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"-" url:"-"`
}

type GetStableUrlRequest struct {
	// The fully qualified resource name. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}.
	Name string `json:"-" url:"-"`
}

type ListFailoverGroupsRequest struct {
	// Maximum number of failover groups to return per page. Default: 50,
	// maximum: 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token received from a previous ListFailoverGroups call. Provide this
	// to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent resource. Format: accounts/{account_id}.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFailoverGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFailoverGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for listing failover groups.
type ListFailoverGroupsResponse struct {
	// The failover groups for this account.
	FailoverGroups []FailoverGroup `json:"failover_groups,omitempty"`
	// A token that can be sent as page_token to retrieve the next page. If
	// omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFailoverGroupsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFailoverGroupsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListStableUrlsRequest struct {
	// Maximum number of stable URLs to return per page. Default: 50, maximum:
	// 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token received from a previous ListStableUrls call. Provide this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The parent resource. Format: accounts/{account_id}.
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListStableUrlsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStableUrlsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for listing stable URLs.
type ListStableUrlsResponse struct {
	// A token that can be sent as page_token to retrieve the next page. If
	// omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The stable URLs for this account.
	StableUrls []StableUrl `json:"stable_urls,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListStableUrlsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListStableUrlsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A location mapping identified by a name, with URIs per region. The system
// derives replication direction from effective_primary_region.
type LocationMapping struct {
	// Resource name for this location.
	Name string `json:"name"`
	// URI for each region. Each entry maps a region name to a storage URI.
	UriByRegion []LocationMappingEntry `json:"uri_by_region"`
}

// A single entry in a location mapping, mapping a region to a storage URI. Used
// instead of map<string, string> for proto2 compatibility.
type LocationMappingEntry struct {
	// The region name.
	Region string `json:"region"`
	// The storage URI for this region.
	Uri string `json:"uri"`
}

// A stable URL provides a failover-aware endpoint for accessing a workspace.
// Its lifecycle is independent of any failover group.
type StableUrl struct {
	// The workspace this stable URL is initially bound to. Used only in Create
	// requests to associate the stable URL with a workspace. Not returned in
	// responses. Mirrors FailoverGroup.initial_primary_region semantics.
	InitialWorkspaceId string `json:"initial_workspace_id"`
	// Fully qualified resource name. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}.
	Name string `json:"name,omitempty"`
	// The stable URL endpoint. Generated by the backend on creation and
	// immutable thereafter. For non-Private-Link workspaces this is
	// `https://<spog_host>/?c=<connection_id>`. For Private-Link workspaces
	// this is the per-connection hostname.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StableUrl) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StableUrl) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Unity Catalog catalog to replicate.
type UcCatalog struct {
	// The name of the UC catalog to replicate.
	Name string `json:"name"`
}

// Unity Catalog replication configuration (top-level, not per-set).
type UcReplicationConfig struct {
	// UC catalogs to replicate.
	Catalogs []UcCatalog `json:"catalogs"`
	// The workspace set whose workspaces will be used for data replication of
	// all UC catalogs' underlying storage.
	DataReplicationWorkspaceSet string `json:"data_replication_workspace_set"`
	// Location mappings - storage URI per region for each location.
	LocationMappings []LocationMapping `json:"location_mappings,omitempty"`
}

type UpdateFailoverGroupRequest struct {
	// The failover group with updated fields. The name field identifies the
	// resource and is populated from the URL path.
	FailoverGroup FailoverGroup `json:"failover_group"`
	// Fully qualified resource name in the format
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name string `json:"-" url:"-"`
	// Comma-separated list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

// A set of workspaces that replicate to each other across regions.
type WorkspaceSet struct {
	// Resource name for this workspace set.
	Name string `json:"name"`
	// Whether to enable control plane DR (notebooks, jobs, clusters, etc.) for
	// this set. Requires all workspaces in the set to be Mission Critical tier.
	ReplicateWorkspaceAssets bool `json:"replicate_workspace_assets"`
	// Resource names of stable URLs associated with this workspace set. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}. The referenced stable
	// URLs must already exist (via CreateStableUrl).
	StableUrlNames []string `json:"stable_url_names,omitempty"`
	// Workspace IDs in this set. The system derives and validates regions. EA:
	// exactly 2 workspaces (one per region).
	WorkspaceIds []string `json:"workspace_ids"`
}
