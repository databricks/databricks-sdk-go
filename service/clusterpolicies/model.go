// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

import "fmt"

// all definitions in this file are in alphabetical order

type CreatePolicy struct {
	// Policy definition document expressed in Databricks Cluster Policy
	// Definition Language.
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	Description string `json:"description,omitempty"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser int64 `json:"max_clusters_per_user,omitempty"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string `json:"name"`
	// Policy definition JSON document expressed in Databricks Policy Definition
	// Language. The JSON document must be passed as a string and cannot be
	// embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId string `json:"policy_family_id,omitempty"`
}

type CreatePolicyResponse struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId string `json:"policy_id,omitempty"`
}

type DeletePolicy struct {
	// The ID of the policy to delete.
	PolicyId string `json:"policy_id"`
}

type EditPolicy struct {
	// Policy definition document expressed in Databricks Cluster Policy
	// Definition Language.
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	Description string `json:"description,omitempty"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser int64 `json:"max_clusters_per_user,omitempty"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string `json:"name"`
	// Policy definition JSON document expressed in Databricks Policy Definition
	// Language. The JSON document must be passed as a string and cannot be
	// embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId string `json:"policy_family_id,omitempty"`
	// The ID of the policy to update.
	PolicyId string `json:"policy_id"`
}

// Get entity
type Get struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId string `json:"-" url:"policy_id"`
}

type GetPolicyFamilyRequest struct {
	PolicyFamilyId string `json:"-" url:"-"`
}

// Get a cluster policy
type List struct {
	// The cluster policy attribute to sort by. * `POLICY_CREATION_TIME` - Sort
	// result list by policy creation time. * `POLICY_NAME` - Sort result list
	// by policy name.
	SortColumn ListSortColumn `json:"-" url:"sort_column,omitempty"`
	// The order in which the policies get listed. * `DESC` - Sort result list
	// in descending order. * `ASC` - Sort result list in ascending order.
	SortOrder ListSortOrder `json:"-" url:"sort_order,omitempty"`
}

type ListPoliciesResponse struct {
	// List of policies.
	Policies []Policy `json:"policies,omitempty"`
}

type ListPolicyFamiliesRequest struct {
	// The max number of policy families to return.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// A token that can be used to get the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
}

type ListPolicyFamiliesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of policy families.
	PolicyFamilies []PolicyFamily `json:"policy_families"`
}

type ListSortColumn string

const ListSortColumnPolicyCreationTime ListSortColumn = `POLICY_CREATION_TIME`

const ListSortColumnPolicyName ListSortColumn = `POLICY_NAME`

// String representation for [fmt.Print]
func (lsc *ListSortColumn) String() string {
	return string(*lsc)
}

// Set raw string value and validate it against allowed values
func (lsc *ListSortColumn) Set(v string) error {
	switch v {
	case `POLICY_CREATION_TIME`, `POLICY_NAME`:
		*lsc = ListSortColumn(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "POLICY_CREATION_TIME", "POLICY_NAME"`, v)
	}
}

// Type always returns ListSortColumn to satisfy [pflag.Value] interface
func (lsc *ListSortColumn) Type() string {
	return "ListSortColumn"
}

type ListSortOrder string

const ListSortOrderAsc ListSortOrder = `ASC`

const ListSortOrderDesc ListSortOrder = `DESC`

// String representation for [fmt.Print]
func (lso *ListSortOrder) String() string {
	return string(*lso)
}

// Set raw string value and validate it against allowed values
func (lso *ListSortOrder) Set(v string) error {
	switch v {
	case `ASC`, `DESC`:
		*lso = ListSortOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASC", "DESC"`, v)
	}
}

// Type always returns ListSortOrder to satisfy [pflag.Value] interface
func (lso *ListSortOrder) Type() string {
	return "ListSortOrder"
}

type Policy struct {
	// Creation time. The timestamp (in millisecond) when this Cluster Policy
	// was created.
	CreatedAtTimestamp int64 `json:"created_at_timestamp,omitempty"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Policy definition document expressed in Databricks Cluster Policy
	// Definition Language.
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	Description string `json:"description,omitempty"`
	// If true, policy is a default policy created and managed by Databricks.
	// Default policies cannot be deleted, and their policy families cannot be
	// changed.
	IsDefault bool `json:"is_default,omitempty"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser int64 `json:"max_clusters_per_user,omitempty"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string `json:"name,omitempty"`
	// Policy definition JSON document expressed in Databricks Policy Definition
	// Language. The JSON document must be passed as a string and cannot be
	// embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty"`
	// ID of the policy family.
	PolicyFamilyId string `json:"policy_family_id,omitempty"`
	// Canonical unique identifier for the Cluster Policy.
	PolicyId string `json:"policy_id,omitempty"`
}

type PolicyFamily struct {
	// Policy definition document expressed in Databricks Cluster Policy
	// Definition Language.
	Definition string `json:"definition"`
	// Human-readable description of the purpose of the policy family.
	Description string `json:"description"`
	// Name of the policy family.
	Name string `json:"name"`
	// ID of the policy family.
	PolicyFamilyId string `json:"policy_family_id"`
}
