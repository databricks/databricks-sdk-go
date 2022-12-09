// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

// all definitions in this file are in alphabetical order

type CreatePolicy struct {
	// Policy definition document expressed in Databricks Cluster Policy
	// Definition Language.
	Definition string `json:"definition"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string `json:"name"`
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
	Definition string `json:"definition"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string `json:"name"`
	// The ID of the policy to update.
	PolicyId string `json:"policy_id"`
}

// Get entity
type Get struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId string `json:"-" url:"policy_id,omitempty"`
}

type ListPoliciesResponse struct {
	// List of policies.
	Policies []Policy `json:"policies,omitempty"`
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
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string `json:"name,omitempty"`
	// Canonical unique identifier for the Cluster Policy.
	PolicyId string `json:"policy_id,omitempty"`
}
