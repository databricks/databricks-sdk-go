// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tags

import (
	"context"
)

// Manage tag assignments on workspace-scoped objects.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type TagAssignmentsService interface {

	// Create a tag assignment
	CreateTagAssignment(ctx context.Context, request CreateTagAssignmentRequest) (*TagAssignment, error)

	// Delete a tag assignment
	DeleteTagAssignment(ctx context.Context, request DeleteTagAssignmentRequest) error

	// Get a tag assignment
	GetTagAssignment(ctx context.Context, request GetTagAssignmentRequest) (*TagAssignment, error)

	// List the tag assignments for an entity
	ListTagAssignments(ctx context.Context, request ListTagAssignmentsRequest) (*ListTagAssignmentsResponse, error)

	// Update a tag assignment
	UpdateTagAssignment(ctx context.Context, request UpdateTagAssignmentRequest) (*TagAssignment, error)
}

// The Tag Policy API allows you to manage policies for governed tags in
// Databricks.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type TagPoliciesService interface {

	// Creates a new tag policy, making the associated tag key governed.
	CreateTagPolicy(ctx context.Context, request CreateTagPolicyRequest) (*TagPolicy, error)

	// Deletes a tag policy by its associated governed tag's key, leaving that
	// tag key ungoverned.
	DeleteTagPolicy(ctx context.Context, request DeleteTagPolicyRequest) error

	// Gets a single tag policy by its associated governed tag's key.
	GetTagPolicy(ctx context.Context, request GetTagPolicyRequest) (*TagPolicy, error)

	// Lists the tag policies for all governed tags in the account.
	ListTagPolicies(ctx context.Context, request ListTagPoliciesRequest) (*ListTagPoliciesResponse, error)

	// Updates an existing tag policy for a single governed tag.
	UpdateTagPolicy(ctx context.Context, request UpdateTagPolicyRequest) (*TagPolicy, error)
}
