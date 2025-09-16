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
