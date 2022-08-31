// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

// all definitions in this file are in alphabetical order

type CreatePolicyRequest struct {
    // Policy definition document expressed in Databricks Cluster Policy
    // Definition Language.
    Definition string `json:"definition,omitempty"`
    // Cluster Policy name requested by the user. This has to be unique. Length
    // must be between 1 and 100 characters.
    Name string `json:"name,omitempty"`
}


type CreatePolicyResponse struct {
    
    PolicyId string `json:"policy_id,omitempty"`
}


type DeletePolicyRequest struct {
    // The ID of the policy to delete.
    PolicyId string `json:"policy_id"`
}


type EditPolicyRequest struct {
    // Policy definition document expressed in Databricks Cluster Policy
    // Definition Language.
    Definition string `json:"definition,omitempty"`
    // Cluster Policy name requested by the user. This has to be unique. Length
    // must be between 1 and 100 characters.
    Name string `json:"name,omitempty"`
    // The ID of the policy to update.
    PolicyId string `json:"policy_id"`
}


type GetPolicyRequest struct {
    // The policy ID about which to retrieve information.
    PolicyId string ` url:"policy_id,omitempty"`
}


type GetPolicyResponse struct {
    // Creation time. The timestamp (in millisecond) when this Cluster Policy
    // was created.
    CreatedAtTimestamp int64 `json:"created_at_timestamp,omitempty"`
    // Creator user name. The field won&#39;t be included in the response if the
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


type ListPoliciesResponse struct {
    
    Policies []Policy `json:"policies,omitempty"`
}


type Policy struct {
    // Creation time. The timestamp (in millisecond) when this Cluster Policy
    // was created.
    CreatedAtTimestamp int64 `json:"created_at_timestamp,omitempty"`
    // Creator user name. The field won&#39;t be included in the response if the
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

