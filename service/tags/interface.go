// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tags

import (
	"context"
)

// The Tag Policy API allows you to manage tag policies in Databricks.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type TagPoliciesService interface {

	// Creates a new tag policy.
	CreateTagPolicy(ctx context.Context, request CreateTagPolicyRequest) (*TagPolicy, error)

	// Deletes a tag policy by its key.
	DeleteTagPolicy(ctx context.Context, request DeleteTagPolicyRequest) error

	// Gets a single tag policy by its key.
	GetTagPolicy(ctx context.Context, request GetTagPolicyRequest) (*TagPolicy, error)

	// Lists all tag policies in the account.
	ListTagPolicies(ctx context.Context, request ListTagPoliciesRequest) (*ListTagPoliciesResponse, error)

	// Updates an existing tag policy.
	UpdateTagPolicy(ctx context.Context, request UpdateTagPolicyRequest) (*TagPolicy, error)
}
