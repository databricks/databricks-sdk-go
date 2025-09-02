// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tags

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just TagAssignments API methods
type tagAssignmentsImpl struct {
	client *client.DatabricksClient
}

func (a *tagAssignmentsImpl) CreateTagAssignment(ctx context.Context, request CreateTagAssignmentRequest) (*TagAssignment, error) {
	var tagAssignment TagAssignment
	path := "/api/2.0/entity-tag-assignments"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.TagAssignment, &tagAssignment)
	return &tagAssignment, err
}

func (a *tagAssignmentsImpl) DeleteTagAssignment(ctx context.Context, request DeleteTagAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.0/entity-tag-assignments/%v/%v/tags/%v", request.EntityType, request.EntityId, request.TagKey)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *tagAssignmentsImpl) GetTagAssignment(ctx context.Context, request GetTagAssignmentRequest) (*TagAssignment, error) {
	var tagAssignment TagAssignment
	path := fmt.Sprintf("/api/2.0/entity-tag-assignments/%v/%v/tags/%v", request.EntityType, request.EntityId, request.TagKey)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &tagAssignment)
	return &tagAssignment, err
}

// List the tag assignments for an entity
func (a *tagAssignmentsImpl) ListTagAssignments(ctx context.Context, request ListTagAssignmentsRequest) listing.Iterator[TagAssignment] {

	getNextPage := func(ctx context.Context, req ListTagAssignmentsRequest) (*ListTagAssignmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListTagAssignments(ctx, req)
	}
	getItems := func(resp *ListTagAssignmentsResponse) []TagAssignment {
		return resp.TagAssignments
	}
	getNextReq := func(resp *ListTagAssignmentsResponse) *ListTagAssignmentsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List the tag assignments for an entity
func (a *tagAssignmentsImpl) ListTagAssignmentsAll(ctx context.Context, request ListTagAssignmentsRequest) ([]TagAssignment, error) {
	iterator := a.ListTagAssignments(ctx, request)
	return listing.ToSlice[TagAssignment](ctx, iterator)
}

func (a *tagAssignmentsImpl) internalListTagAssignments(ctx context.Context, request ListTagAssignmentsRequest) (*ListTagAssignmentsResponse, error) {
	var listTagAssignmentsResponse ListTagAssignmentsResponse
	path := fmt.Sprintf("/api/2.0/entity-tag-assignments/%v/%v", request.EntityType, request.EntityId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTagAssignmentsResponse)
	return &listTagAssignmentsResponse, err
}

func (a *tagAssignmentsImpl) UpdateTagAssignment(ctx context.Context, request UpdateTagAssignmentRequest) (*TagAssignment, error) {
	var tagAssignment TagAssignment
	path := fmt.Sprintf("/api/2.0/entity-tag-assignments/%v/%v/tags/%v", request.EntityType, request.EntityId, request.TagKey)
	queryParams := make(map[string]any)
	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.TagAssignment, &tagAssignment)
	return &tagAssignment, err
}

// unexported type that holds implementations of just TagPolicies API methods
type tagPoliciesImpl struct {
	client *client.DatabricksClient
}

func (a *tagPoliciesImpl) CreateTagPolicy(ctx context.Context, request CreateTagPolicyRequest) (*TagPolicy, error) {
	var tagPolicy TagPolicy
	path := "/api/2.1/tag-policies"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.TagPolicy, &tagPolicy)
	return &tagPolicy, err
}

func (a *tagPoliciesImpl) DeleteTagPolicy(ctx context.Context, request DeleteTagPolicyRequest) error {
	path := fmt.Sprintf("/api/2.1/tag-policies/%v", request.TagKey)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *tagPoliciesImpl) GetTagPolicy(ctx context.Context, request GetTagPolicyRequest) (*TagPolicy, error) {
	var tagPolicy TagPolicy
	path := fmt.Sprintf("/api/2.1/tag-policies/%v", request.TagKey)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &tagPolicy)
	return &tagPolicy, err
}

// Lists the tag policies for all governed tags in the account.
func (a *tagPoliciesImpl) ListTagPolicies(ctx context.Context, request ListTagPoliciesRequest) listing.Iterator[TagPolicy] {

	getNextPage := func(ctx context.Context, req ListTagPoliciesRequest) (*ListTagPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListTagPolicies(ctx, req)
	}
	getItems := func(resp *ListTagPoliciesResponse) []TagPolicy {
		return resp.TagPolicies
	}
	getNextReq := func(resp *ListTagPoliciesResponse) *ListTagPoliciesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Lists the tag policies for all governed tags in the account.
func (a *tagPoliciesImpl) ListTagPoliciesAll(ctx context.Context, request ListTagPoliciesRequest) ([]TagPolicy, error) {
	iterator := a.ListTagPolicies(ctx, request)
	return listing.ToSlice[TagPolicy](ctx, iterator)
}

func (a *tagPoliciesImpl) internalListTagPolicies(ctx context.Context, request ListTagPoliciesRequest) (*ListTagPoliciesResponse, error) {
	var listTagPoliciesResponse ListTagPoliciesResponse
	path := "/api/2.1/tag-policies"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTagPoliciesResponse)
	return &listTagPoliciesResponse, err
}

func (a *tagPoliciesImpl) UpdateTagPolicy(ctx context.Context, request UpdateTagPolicyRequest) (*TagPolicy, error) {
	var tagPolicy TagPolicy
	path := fmt.Sprintf("/api/2.1/tag-policies/%v", request.TagKey)
	queryParams := make(map[string]any)
	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.TagPolicy, &tagPolicy)
	return &tagPolicy, err
}
