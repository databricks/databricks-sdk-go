// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/cleanrooms/cleanroomspb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just CleanRoomAssetRevisions API methods
type cleanRoomAssetRevisionsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomAssetRevisionsImpl) Get(ctx context.Context, request GetCleanRoomAssetRevisionRequest) (*CleanRoomAsset, error) {
	requestPb, pbErr := GetCleanRoomAssetRevisionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAssetPb cleanroomspb.CleanRoomAssetPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v/revisions/%v", request.CleanRoomName, request.AssetType, request.Name, request.Etag)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&cleanRoomAssetPb,
	)
	resp, err := CleanRoomAssetFromPb(&cleanRoomAssetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List revisions for an asset
func (a *cleanRoomAssetRevisionsImpl) List(ctx context.Context, request ListCleanRoomAssetRevisionsRequest) listing.Iterator[CleanRoomAsset] {

	getNextPage := func(ctx context.Context, req ListCleanRoomAssetRevisionsRequest) (*ListCleanRoomAssetRevisionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomAssetRevisionsResponse) []CleanRoomAsset {
		return resp.Revisions
	}
	getNextReq := func(resp *ListCleanRoomAssetRevisionsResponse) *ListCleanRoomAssetRevisionsRequest {
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

// List revisions for an asset
func (a *cleanRoomAssetRevisionsImpl) ListAll(ctx context.Context, request ListCleanRoomAssetRevisionsRequest) ([]CleanRoomAsset, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomAsset](ctx, iterator)
}

func (a *cleanRoomAssetRevisionsImpl) internalList(ctx context.Context, request ListCleanRoomAssetRevisionsRequest) (*ListCleanRoomAssetRevisionsResponse, error) {
	requestPb, pbErr := ListCleanRoomAssetRevisionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCleanRoomAssetRevisionsResponsePb cleanroomspb.ListCleanRoomAssetRevisionsResponsePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v/revisions", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listCleanRoomAssetRevisionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListCleanRoomAssetRevisionsResponseFromPb(&listCleanRoomAssetRevisionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CleanRoomAssets API methods
type cleanRoomAssetsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomAssetsImpl) Create(ctx context.Context, request CreateCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	requestPb, pbErr := CreateCleanRoomAssetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAssetPb cleanroomspb.CleanRoomAssetPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.Asset,
		&cleanRoomAssetPb,
	)
	resp, err := CleanRoomAssetFromPb(&cleanRoomAssetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomAssetsImpl) CreateCleanRoomAssetReview(ctx context.Context, request CreateCleanRoomAssetReviewRequest) (*CreateCleanRoomAssetReviewResponse, error) {
	requestPb, pbErr := CreateCleanRoomAssetReviewRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createCleanRoomAssetReviewResponsePb cleanroomspb.CreateCleanRoomAssetReviewResponsePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v/reviews", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&createCleanRoomAssetReviewResponsePb,
	)
	resp, err := CreateCleanRoomAssetReviewResponseFromPb(&createCleanRoomAssetReviewResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomAssetsImpl) Delete(ctx context.Context, request DeleteCleanRoomAssetRequest) error {
	requestPb, pbErr := DeleteCleanRoomAssetRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *cleanRoomAssetsImpl) Get(ctx context.Context, request GetCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	requestPb, pbErr := GetCleanRoomAssetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAssetPb cleanroomspb.CleanRoomAssetPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&cleanRoomAssetPb,
	)
	resp, err := CleanRoomAssetFromPb(&cleanRoomAssetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List assets.
func (a *cleanRoomAssetsImpl) List(ctx context.Context, request ListCleanRoomAssetsRequest) listing.Iterator[CleanRoomAsset] {

	getNextPage := func(ctx context.Context, req ListCleanRoomAssetsRequest) (*ListCleanRoomAssetsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomAssetsResponse) []CleanRoomAsset {
		return resp.Assets
	}
	getNextReq := func(resp *ListCleanRoomAssetsResponse) *ListCleanRoomAssetsRequest {
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

// List assets.
func (a *cleanRoomAssetsImpl) ListAll(ctx context.Context, request ListCleanRoomAssetsRequest) ([]CleanRoomAsset, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomAsset](ctx, iterator)
}

func (a *cleanRoomAssetsImpl) internalList(ctx context.Context, request ListCleanRoomAssetsRequest) (*ListCleanRoomAssetsResponse, error) {
	requestPb, pbErr := ListCleanRoomAssetsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCleanRoomAssetsResponsePb cleanroomspb.ListCleanRoomAssetsResponsePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listCleanRoomAssetsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListCleanRoomAssetsResponseFromPb(&listCleanRoomAssetsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomAssetsImpl) Update(ctx context.Context, request UpdateCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	requestPb, pbErr := UpdateCleanRoomAssetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAssetPb cleanroomspb.CleanRoomAssetPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.Asset,
		&cleanRoomAssetPb,
	)
	resp, err := CleanRoomAssetFromPb(&cleanRoomAssetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CleanRoomAutoApprovalRules API methods
type cleanRoomAutoApprovalRulesImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomAutoApprovalRulesImpl) Create(ctx context.Context, request CreateCleanRoomAutoApprovalRuleRequest) (*CleanRoomAutoApprovalRule, error) {
	requestPb, pbErr := CreateCleanRoomAutoApprovalRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAutoApprovalRulePb cleanroomspb.CleanRoomAutoApprovalRulePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&cleanRoomAutoApprovalRulePb,
	)
	resp, err := CleanRoomAutoApprovalRuleFromPb(&cleanRoomAutoApprovalRulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomAutoApprovalRulesImpl) Delete(ctx context.Context, request DeleteCleanRoomAutoApprovalRuleRequest) error {
	requestPb, pbErr := DeleteCleanRoomAutoApprovalRuleRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules/%v", request.CleanRoomName, request.RuleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *cleanRoomAutoApprovalRulesImpl) Get(ctx context.Context, request GetCleanRoomAutoApprovalRuleRequest) (*CleanRoomAutoApprovalRule, error) {
	requestPb, pbErr := GetCleanRoomAutoApprovalRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAutoApprovalRulePb cleanroomspb.CleanRoomAutoApprovalRulePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules/%v", request.CleanRoomName, request.RuleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&cleanRoomAutoApprovalRulePb,
	)
	resp, err := CleanRoomAutoApprovalRuleFromPb(&cleanRoomAutoApprovalRulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List all auto-approval rules for the caller
func (a *cleanRoomAutoApprovalRulesImpl) List(ctx context.Context, request ListCleanRoomAutoApprovalRulesRequest) listing.Iterator[CleanRoomAutoApprovalRule] {

	getNextPage := func(ctx context.Context, req ListCleanRoomAutoApprovalRulesRequest) (*ListCleanRoomAutoApprovalRulesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomAutoApprovalRulesResponse) []CleanRoomAutoApprovalRule {
		return resp.Rules
	}
	getNextReq := func(resp *ListCleanRoomAutoApprovalRulesResponse) *ListCleanRoomAutoApprovalRulesRequest {
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

// List all auto-approval rules for the caller
func (a *cleanRoomAutoApprovalRulesImpl) ListAll(ctx context.Context, request ListCleanRoomAutoApprovalRulesRequest) ([]CleanRoomAutoApprovalRule, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomAutoApprovalRule](ctx, iterator)
}

func (a *cleanRoomAutoApprovalRulesImpl) internalList(ctx context.Context, request ListCleanRoomAutoApprovalRulesRequest) (*ListCleanRoomAutoApprovalRulesResponse, error) {
	requestPb, pbErr := ListCleanRoomAutoApprovalRulesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCleanRoomAutoApprovalRulesResponsePb cleanroomspb.ListCleanRoomAutoApprovalRulesResponsePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listCleanRoomAutoApprovalRulesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListCleanRoomAutoApprovalRulesResponseFromPb(&listCleanRoomAutoApprovalRulesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomAutoApprovalRulesImpl) Update(ctx context.Context, request UpdateCleanRoomAutoApprovalRuleRequest) (*CleanRoomAutoApprovalRule, error) {
	requestPb, pbErr := UpdateCleanRoomAutoApprovalRuleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAutoApprovalRulePb cleanroomspb.CleanRoomAutoApprovalRulePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/auto-approval-rules/%v", request.CleanRoomName, request.RuleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.AutoApprovalRule,
		&cleanRoomAutoApprovalRulePb,
	)
	resp, err := CleanRoomAutoApprovalRuleFromPb(&cleanRoomAutoApprovalRulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CleanRoomTaskRuns API methods
type cleanRoomTaskRunsImpl struct {
	client *client.DatabricksClient
}

// List all the historical notebook task runs in a clean room.
func (a *cleanRoomTaskRunsImpl) List(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) listing.Iterator[CleanRoomNotebookTaskRun] {

	getNextPage := func(ctx context.Context, req ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomNotebookTaskRunsResponse) []CleanRoomNotebookTaskRun {
		return resp.Runs
	}
	getNextReq := func(resp *ListCleanRoomNotebookTaskRunsResponse) *ListCleanRoomNotebookTaskRunsRequest {
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

// List all the historical notebook task runs in a clean room.
func (a *cleanRoomTaskRunsImpl) ListAll(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) ([]CleanRoomNotebookTaskRun, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomNotebookTaskRun](ctx, iterator)
}

func (a *cleanRoomTaskRunsImpl) internalList(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error) {
	requestPb, pbErr := ListCleanRoomNotebookTaskRunsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCleanRoomNotebookTaskRunsResponsePb cleanroomspb.ListCleanRoomNotebookTaskRunsResponsePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/runs", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listCleanRoomNotebookTaskRunsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListCleanRoomNotebookTaskRunsResponseFromPb(&listCleanRoomNotebookTaskRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CleanRooms API methods
type cleanRoomsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomsImpl) Create(ctx context.Context, request CreateCleanRoomRequest) (*CleanRoom, error) {
	requestPb, pbErr := CreateCleanRoomRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomPb cleanroomspb.CleanRoomPb
	path := "/api/2.0/clean-rooms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.CleanRoom,
		&cleanRoomPb,
	)
	resp, err := CleanRoomFromPb(&cleanRoomPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomsImpl) CreateOutputCatalog(ctx context.Context, request CreateCleanRoomOutputCatalogRequest) (*CreateCleanRoomOutputCatalogResponse, error) {
	requestPb, pbErr := CreateCleanRoomOutputCatalogRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createCleanRoomOutputCatalogResponsePb cleanroomspb.CreateCleanRoomOutputCatalogResponsePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/output-catalogs", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.OutputCatalog,
		&createCleanRoomOutputCatalogResponsePb,
	)
	resp, err := CreateCleanRoomOutputCatalogResponseFromPb(&createCleanRoomOutputCatalogResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomsImpl) Delete(ctx context.Context, request DeleteCleanRoomRequest) error {
	requestPb, pbErr := DeleteCleanRoomRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *cleanRoomsImpl) Get(ctx context.Context, request GetCleanRoomRequest) (*CleanRoom, error) {
	requestPb, pbErr := GetCleanRoomRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomPb cleanroomspb.CleanRoomPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&cleanRoomPb,
	)
	resp, err := CleanRoomFromPb(&cleanRoomPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Get a list of all clean rooms of the metastore. Only clean rooms the caller
// has access to are returned.
func (a *cleanRoomsImpl) List(ctx context.Context, request ListCleanRoomsRequest) listing.Iterator[CleanRoom] {

	getNextPage := func(ctx context.Context, req ListCleanRoomsRequest) (*ListCleanRoomsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomsResponse) []CleanRoom {
		return resp.CleanRooms
	}
	getNextReq := func(resp *ListCleanRoomsResponse) *ListCleanRoomsRequest {
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

// Get a list of all clean rooms of the metastore. Only clean rooms the caller
// has access to are returned.
func (a *cleanRoomsImpl) ListAll(ctx context.Context, request ListCleanRoomsRequest) ([]CleanRoom, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoom](ctx, iterator)
}

func (a *cleanRoomsImpl) internalList(ctx context.Context, request ListCleanRoomsRequest) (*ListCleanRoomsResponse, error) {
	requestPb, pbErr := ListCleanRoomsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCleanRoomsResponsePb cleanroomspb.ListCleanRoomsResponsePb
	path := "/api/2.0/clean-rooms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listCleanRoomsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListCleanRoomsResponseFromPb(&listCleanRoomsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomsImpl) Update(ctx context.Context, request UpdateCleanRoomRequest) (*CleanRoom, error) {
	requestPb, pbErr := UpdateCleanRoomRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomPb cleanroomspb.CleanRoomPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&cleanRoomPb,
	)
	resp, err := CleanRoomFromPb(&cleanRoomPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
