// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just ClusterPolicies API methods
type clusterPoliciesImpl struct {
	client *client.DatabricksClient
}

func (a *clusterPoliciesImpl) Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error) {

	requestPb, pbErr := createPolicyToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createPolicyResponsePb createPolicyResponsePb
	path := "/api/2.0/policies/clusters/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createPolicyResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := createPolicyResponseFromPb(&createPolicyResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clusterPoliciesImpl) Delete(ctx context.Context, request DeletePolicy) error {

	requestPb, pbErr := deletePolicyToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deletePolicyResponsePb deletePolicyResponsePb
	path := "/api/2.0/policies/clusters/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &deletePolicyResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clusterPoliciesImpl) Edit(ctx context.Context, request EditPolicy) error {

	requestPb, pbErr := editPolicyToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var editPolicyResponsePb editPolicyResponsePb
	path := "/api/2.0/policies/clusters/edit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &editPolicyResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clusterPoliciesImpl) Get(ctx context.Context, request GetClusterPolicyRequest) (*Policy, error) {

	requestPb, pbErr := getClusterPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var policyPb policyPb
	path := "/api/2.0/policies/clusters/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &policyPb)
	if err != nil {
		return nil, err
	}
	resp, err := policyFromPb(&policyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clusterPoliciesImpl) GetPermissionLevels(ctx context.Context, request GetClusterPolicyPermissionLevelsRequest) (*GetClusterPolicyPermissionLevelsResponse, error) {

	requestPb, pbErr := getClusterPolicyPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getClusterPolicyPermissionLevelsResponsePb getClusterPolicyPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v/permissionLevels", requestPb.ClusterPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getClusterPolicyPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getClusterPolicyPermissionLevelsResponseFromPb(&getClusterPolicyPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clusterPoliciesImpl) GetPermissions(ctx context.Context, request GetClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {

	requestPb, pbErr := getClusterPolicyPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var clusterPolicyPermissionsPb clusterPolicyPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", requestPb.ClusterPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &clusterPolicyPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := clusterPolicyPermissionsFromPb(&clusterPolicyPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List cluster policies.
//
// Returns a list of policies accessible by the requesting user.
func (a *clusterPoliciesImpl) List(ctx context.Context, request ListClusterPoliciesRequest) listing.Iterator[Policy] {

	getNextPage := func(ctx context.Context, req ListClusterPoliciesRequest) (*ListPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListPoliciesResponse) []Policy {
		return resp.Policies
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List cluster policies.
//
// Returns a list of policies accessible by the requesting user.
func (a *clusterPoliciesImpl) ListAll(ctx context.Context, request ListClusterPoliciesRequest) ([]Policy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Policy](ctx, iterator)
}

func (a *clusterPoliciesImpl) internalList(ctx context.Context, request ListClusterPoliciesRequest) (*ListPoliciesResponse, error) {

	requestPb, pbErr := listClusterPoliciesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listPoliciesResponsePb listPoliciesResponsePb
	path := "/api/2.0/policies/clusters/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listPoliciesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listPoliciesResponseFromPb(&listPoliciesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clusterPoliciesImpl) SetPermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {

	requestPb, pbErr := clusterPolicyPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var clusterPolicyPermissionsPb clusterPolicyPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", requestPb.ClusterPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, (*requestPb), &clusterPolicyPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := clusterPolicyPermissionsFromPb(&clusterPolicyPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clusterPoliciesImpl) UpdatePermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {

	requestPb, pbErr := clusterPolicyPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var clusterPolicyPermissionsPb clusterPolicyPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", requestPb.ClusterPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, (*requestPb), &clusterPolicyPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := clusterPolicyPermissionsFromPb(&clusterPolicyPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Clusters API methods
type clustersImpl struct {
	client *client.DatabricksClient
}

func (a *clustersImpl) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {

	requestPb, pbErr := changeClusterOwnerToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var changeClusterOwnerResponsePb changeClusterOwnerResponsePb
	path := "/api/2.1/clusters/change-owner"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &changeClusterOwnerResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clustersImpl) Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error) {

	requestPb, pbErr := createClusterToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createClusterResponsePb createClusterResponsePb
	path := "/api/2.1/clusters/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createClusterResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := createClusterResponseFromPb(&createClusterResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clustersImpl) Delete(ctx context.Context, request DeleteCluster) error {

	requestPb, pbErr := deleteClusterToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteClusterResponsePb deleteClusterResponsePb
	path := "/api/2.1/clusters/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &deleteClusterResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clustersImpl) Edit(ctx context.Context, request EditCluster) error {

	requestPb, pbErr := editClusterToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var editClusterResponsePb editClusterResponsePb
	path := "/api/2.1/clusters/edit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &editClusterResponsePb)
	if err != nil {
		return err
	}

	return err
}

// List cluster activity events.
//
// Retrieves a list of events about the activity of a cluster. This API is
// paginated. If there are more events to read, the response includes all the
// parameters necessary to request the next page of events.
func (a *clustersImpl) Events(ctx context.Context, request GetEvents) listing.Iterator[ClusterEvent] {

	getNextPage := func(ctx context.Context, req GetEvents) (*GetEventsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalEvents(ctx, req)
	}
	getItems := func(resp *GetEventsResponse) []ClusterEvent {
		return resp.Events
	}
	getNextReq := func(resp *GetEventsResponse) *GetEvents {
		if resp.NextPage == nil {
			return nil
		}
		request = *resp.NextPage

		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List cluster activity events.
//
// Retrieves a list of events about the activity of a cluster. This API is
// paginated. If there are more events to read, the response includes all the
// parameters necessary to request the next page of events.
func (a *clustersImpl) EventsAll(ctx context.Context, request GetEvents) ([]ClusterEvent, error) {
	iterator := a.Events(ctx, request)
	return listing.ToSliceN[ClusterEvent, int64](ctx, iterator, request.Limit)

}

func (a *clustersImpl) internalEvents(ctx context.Context, request GetEvents) (*GetEventsResponse, error) {

	requestPb, pbErr := getEventsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getEventsResponsePb getEventsResponsePb
	path := "/api/2.1/clusters/events"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &getEventsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getEventsResponseFromPb(&getEventsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clustersImpl) Get(ctx context.Context, request GetClusterRequest) (*ClusterDetails, error) {

	requestPb, pbErr := getClusterRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var clusterDetailsPb clusterDetailsPb
	path := "/api/2.1/clusters/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &clusterDetailsPb)
	if err != nil {
		return nil, err
	}
	resp, err := clusterDetailsFromPb(&clusterDetailsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clustersImpl) GetPermissionLevels(ctx context.Context, request GetClusterPermissionLevelsRequest) (*GetClusterPermissionLevelsResponse, error) {

	requestPb, pbErr := getClusterPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getClusterPermissionLevelsResponsePb getClusterPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v/permissionLevels", requestPb.ClusterId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getClusterPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getClusterPermissionLevelsResponseFromPb(&getClusterPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clustersImpl) GetPermissions(ctx context.Context, request GetClusterPermissionsRequest) (*ClusterPermissions, error) {

	requestPb, pbErr := getClusterPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var clusterPermissionsPb clusterPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", requestPb.ClusterId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &clusterPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := clusterPermissionsFromPb(&clusterPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List clusters.
//
// Return information about all pinned and active clusters, and all clusters
// terminated within the last 30 days. Clusters terminated prior to this period
// are not included.
func (a *clustersImpl) List(ctx context.Context, request ListClustersRequest) listing.Iterator[ClusterDetails] {

	getNextPage := func(ctx context.Context, req ListClustersRequest) (*ListClustersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListClustersResponse) []ClusterDetails {
		return resp.Clusters
	}
	getNextReq := func(resp *ListClustersResponse) *ListClustersRequest {
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

// List clusters.
//
// Return information about all pinned and active clusters, and all clusters
// terminated within the last 30 days. Clusters terminated prior to this period
// are not included.
func (a *clustersImpl) ListAll(ctx context.Context, request ListClustersRequest) ([]ClusterDetails, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ClusterDetails](ctx, iterator)
}

func (a *clustersImpl) internalList(ctx context.Context, request ListClustersRequest) (*ListClustersResponse, error) {

	requestPb, pbErr := listClustersRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listClustersResponsePb listClustersResponsePb
	path := "/api/2.1/clusters/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listClustersResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listClustersResponseFromPb(&listClustersResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clustersImpl) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {

	var listNodeTypesResponsePb listNodeTypesResponsePb
	path := "/api/2.1/clusters/list-node-types"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listNodeTypesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listNodeTypesResponseFromPb(&listNodeTypesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clustersImpl) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {

	var listAvailableZonesResponsePb listAvailableZonesResponsePb
	path := "/api/2.1/clusters/list-zones"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listAvailableZonesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listAvailableZonesResponseFromPb(&listAvailableZonesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clustersImpl) PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error {

	requestPb, pbErr := permanentDeleteClusterToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var permanentDeleteClusterResponsePb permanentDeleteClusterResponsePb
	path := "/api/2.1/clusters/permanent-delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &permanentDeleteClusterResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clustersImpl) Pin(ctx context.Context, request PinCluster) error {

	requestPb, pbErr := pinClusterToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var pinClusterResponsePb pinClusterResponsePb
	path := "/api/2.1/clusters/pin"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &pinClusterResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clustersImpl) Resize(ctx context.Context, request ResizeCluster) error {

	requestPb, pbErr := resizeClusterToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var resizeClusterResponsePb resizeClusterResponsePb
	path := "/api/2.1/clusters/resize"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &resizeClusterResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clustersImpl) Restart(ctx context.Context, request RestartCluster) error {

	requestPb, pbErr := restartClusterToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var restartClusterResponsePb restartClusterResponsePb
	path := "/api/2.1/clusters/restart"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &restartClusterResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clustersImpl) SetPermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error) {

	requestPb, pbErr := clusterPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var clusterPermissionsPb clusterPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", requestPb.ClusterId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, (*requestPb), &clusterPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := clusterPermissionsFromPb(&clusterPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clustersImpl) SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error) {

	var getSparkVersionsResponsePb getSparkVersionsResponsePb
	path := "/api/2.1/clusters/spark-versions"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getSparkVersionsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getSparkVersionsResponseFromPb(&getSparkVersionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *clustersImpl) Start(ctx context.Context, request StartCluster) error {

	requestPb, pbErr := startClusterToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var startClusterResponsePb startClusterResponsePb
	path := "/api/2.1/clusters/start"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &startClusterResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clustersImpl) Unpin(ctx context.Context, request UnpinCluster) error {

	requestPb, pbErr := unpinClusterToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var unpinClusterResponsePb unpinClusterResponsePb
	path := "/api/2.1/clusters/unpin"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &unpinClusterResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clustersImpl) Update(ctx context.Context, request UpdateCluster) error {

	requestPb, pbErr := updateClusterToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateClusterResponsePb updateClusterResponsePb
	path := "/api/2.1/clusters/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &updateClusterResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *clustersImpl) UpdatePermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error) {

	requestPb, pbErr := clusterPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var clusterPermissionsPb clusterPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", requestPb.ClusterId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, (*requestPb), &clusterPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := clusterPermissionsFromPb(&clusterPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CommandExecution API methods
type commandExecutionImpl struct {
	client *client.DatabricksClient
}

func (a *commandExecutionImpl) Cancel(ctx context.Context, request CancelCommand) error {

	requestPb, pbErr := cancelCommandToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var cancelResponsePb cancelResponsePb
	path := "/api/1.2/commands/cancel"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &cancelResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *commandExecutionImpl) CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error) {

	requestPb, pbErr := commandStatusRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var commandStatusResponsePb commandStatusResponsePb
	path := "/api/1.2/commands/status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &commandStatusResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := commandStatusResponseFromPb(&commandStatusResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *commandExecutionImpl) ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error) {

	requestPb, pbErr := contextStatusRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var contextStatusResponsePb contextStatusResponsePb
	path := "/api/1.2/contexts/status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &contextStatusResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := contextStatusResponseFromPb(&contextStatusResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *commandExecutionImpl) Create(ctx context.Context, request CreateContext) (*Created, error) {

	requestPb, pbErr := createContextToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createdPb createdPb
	path := "/api/1.2/contexts/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createdPb)
	if err != nil {
		return nil, err
	}
	resp, err := createdFromPb(&createdPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *commandExecutionImpl) Destroy(ctx context.Context, request DestroyContext) error {

	requestPb, pbErr := destroyContextToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var destroyResponsePb destroyResponsePb
	path := "/api/1.2/contexts/destroy"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &destroyResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *commandExecutionImpl) Execute(ctx context.Context, request Command) (*Created, error) {

	requestPb, pbErr := commandToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createdPb createdPb
	path := "/api/1.2/commands/execute"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createdPb)
	if err != nil {
		return nil, err
	}
	resp, err := createdFromPb(&createdPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just GlobalInitScripts API methods
type globalInitScriptsImpl struct {
	client *client.DatabricksClient
}

func (a *globalInitScriptsImpl) Create(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateResponse, error) {

	requestPb, pbErr := globalInitScriptCreateRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createResponsePb createResponsePb
	path := "/api/2.0/global-init-scripts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := createResponseFromPb(&createResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *globalInitScriptsImpl) Delete(ctx context.Context, request DeleteGlobalInitScriptRequest) error {

	requestPb, pbErr := deleteGlobalInitScriptRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", requestPb.ScriptId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, (*requestPb), &deleteResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *globalInitScriptsImpl) Get(ctx context.Context, request GetGlobalInitScriptRequest) (*GlobalInitScriptDetailsWithContent, error) {

	requestPb, pbErr := getGlobalInitScriptRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var globalInitScriptDetailsWithContentPb globalInitScriptDetailsWithContentPb
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", requestPb.ScriptId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &globalInitScriptDetailsWithContentPb)
	if err != nil {
		return nil, err
	}
	resp, err := globalInitScriptDetailsWithContentFromPb(&globalInitScriptDetailsWithContentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Get init scripts.
//
// Get a list of all global init scripts for this workspace. This returns all
// properties for each script but **not** the script contents. To retrieve the
// contents of a script, use the [get a global init
// script](:method:globalinitscripts/get) operation.
func (a *globalInitScriptsImpl) List(ctx context.Context) listing.Iterator[GlobalInitScriptDetails] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListGlobalInitScriptsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListGlobalInitScriptsResponse) []GlobalInitScriptDetails {
		return resp.Scripts
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get init scripts.
//
// Get a list of all global init scripts for this workspace. This returns all
// properties for each script but **not** the script contents. To retrieve the
// contents of a script, use the [get a global init
// script](:method:globalinitscripts/get) operation.
func (a *globalInitScriptsImpl) ListAll(ctx context.Context) ([]GlobalInitScriptDetails, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[GlobalInitScriptDetails](ctx, iterator)
}

func (a *globalInitScriptsImpl) internalList(ctx context.Context) (*ListGlobalInitScriptsResponse, error) {

	var listGlobalInitScriptsResponsePb listGlobalInitScriptsResponsePb
	path := "/api/2.0/global-init-scripts"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listGlobalInitScriptsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listGlobalInitScriptsResponseFromPb(&listGlobalInitScriptsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *globalInitScriptsImpl) Update(ctx context.Context, request GlobalInitScriptUpdateRequest) error {

	requestPb, pbErr := globalInitScriptUpdateRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateResponsePb updateResponsePb
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", requestPb.ScriptId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, (*requestPb), &updateResponsePb)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just InstancePools API methods
type instancePoolsImpl struct {
	client *client.DatabricksClient
}

func (a *instancePoolsImpl) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {

	requestPb, pbErr := createInstancePoolToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createInstancePoolResponsePb createInstancePoolResponsePb
	path := "/api/2.0/instance-pools/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &createInstancePoolResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := createInstancePoolResponseFromPb(&createInstancePoolResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *instancePoolsImpl) Delete(ctx context.Context, request DeleteInstancePool) error {

	requestPb, pbErr := deleteInstancePoolToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteInstancePoolResponsePb deleteInstancePoolResponsePb
	path := "/api/2.0/instance-pools/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &deleteInstancePoolResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *instancePoolsImpl) Edit(ctx context.Context, request EditInstancePool) error {

	requestPb, pbErr := editInstancePoolToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var editInstancePoolResponsePb editInstancePoolResponsePb
	path := "/api/2.0/instance-pools/edit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &editInstancePoolResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *instancePoolsImpl) Get(ctx context.Context, request GetInstancePoolRequest) (*GetInstancePool, error) {

	requestPb, pbErr := getInstancePoolRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getInstancePoolPb getInstancePoolPb
	path := "/api/2.0/instance-pools/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getInstancePoolPb)
	if err != nil {
		return nil, err
	}
	resp, err := getInstancePoolFromPb(&getInstancePoolPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *instancePoolsImpl) GetPermissionLevels(ctx context.Context, request GetInstancePoolPermissionLevelsRequest) (*GetInstancePoolPermissionLevelsResponse, error) {

	requestPb, pbErr := getInstancePoolPermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getInstancePoolPermissionLevelsResponsePb getInstancePoolPermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v/permissionLevels", requestPb.InstancePoolId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getInstancePoolPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getInstancePoolPermissionLevelsResponseFromPb(&getInstancePoolPermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *instancePoolsImpl) GetPermissions(ctx context.Context, request GetInstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {

	requestPb, pbErr := getInstancePoolPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var instancePoolPermissionsPb instancePoolPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", requestPb.InstancePoolId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &instancePoolPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := instancePoolPermissionsFromPb(&instancePoolPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List instance pool info.
//
// Gets a list of instance pools with their statistics.
func (a *instancePoolsImpl) List(ctx context.Context) listing.Iterator[InstancePoolAndStats] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListInstancePools, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListInstancePools) []InstancePoolAndStats {
		return resp.InstancePools
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List instance pool info.
//
// Gets a list of instance pools with their statistics.
func (a *instancePoolsImpl) ListAll(ctx context.Context) ([]InstancePoolAndStats, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[InstancePoolAndStats](ctx, iterator)
}

func (a *instancePoolsImpl) internalList(ctx context.Context) (*ListInstancePools, error) {

	var listInstancePoolsPb listInstancePoolsPb
	path := "/api/2.0/instance-pools/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listInstancePoolsPb)
	if err != nil {
		return nil, err
	}
	resp, err := listInstancePoolsFromPb(&listInstancePoolsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *instancePoolsImpl) SetPermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {

	requestPb, pbErr := instancePoolPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var instancePoolPermissionsPb instancePoolPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", requestPb.InstancePoolId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, (*requestPb), &instancePoolPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := instancePoolPermissionsFromPb(&instancePoolPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *instancePoolsImpl) UpdatePermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {

	requestPb, pbErr := instancePoolPermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var instancePoolPermissionsPb instancePoolPermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", requestPb.InstancePoolId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, (*requestPb), &instancePoolPermissionsPb)
	if err != nil {
		return nil, err
	}
	resp, err := instancePoolPermissionsFromPb(&instancePoolPermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just InstanceProfiles API methods
type instanceProfilesImpl struct {
	client *client.DatabricksClient
}

func (a *instanceProfilesImpl) Add(ctx context.Context, request AddInstanceProfile) error {

	requestPb, pbErr := addInstanceProfileToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var addResponsePb addResponsePb
	path := "/api/2.0/instance-profiles/add"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &addResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *instanceProfilesImpl) Edit(ctx context.Context, request InstanceProfile) error {

	requestPb, pbErr := instanceProfileToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var editResponsePb editResponsePb
	path := "/api/2.0/instance-profiles/edit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &editResponsePb)
	if err != nil {
		return err
	}

	return err
}

// List available instance profiles.
//
// List the instance profiles that the calling user can use to launch a cluster.
//
// This API is available to all users.
func (a *instanceProfilesImpl) List(ctx context.Context) listing.Iterator[InstanceProfile] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListInstanceProfilesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListInstanceProfilesResponse) []InstanceProfile {
		return resp.InstanceProfiles
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List available instance profiles.
//
// List the instance profiles that the calling user can use to launch a cluster.
//
// This API is available to all users.
func (a *instanceProfilesImpl) ListAll(ctx context.Context) ([]InstanceProfile, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[InstanceProfile](ctx, iterator)
}

func (a *instanceProfilesImpl) internalList(ctx context.Context) (*ListInstanceProfilesResponse, error) {

	var listInstanceProfilesResponsePb listInstanceProfilesResponsePb
	path := "/api/2.0/instance-profiles/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listInstanceProfilesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listInstanceProfilesResponseFromPb(&listInstanceProfilesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *instanceProfilesImpl) Remove(ctx context.Context, request RemoveInstanceProfile) error {

	requestPb, pbErr := removeInstanceProfileToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var removeResponsePb removeResponsePb
	path := "/api/2.0/instance-profiles/remove"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &removeResponsePb)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just Libraries API methods
type librariesImpl struct {
	client *client.DatabricksClient
}

// Get all statuses.
//
// Get the status of all libraries on all clusters. A status is returned for all
// libraries installed on this cluster via the API or the libraries UI.
func (a *librariesImpl) AllClusterStatuses(ctx context.Context) listing.Iterator[ClusterLibraryStatuses] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListAllClusterLibraryStatusesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalAllClusterStatuses(ctx)
	}
	getItems := func(resp *ListAllClusterLibraryStatusesResponse) []ClusterLibraryStatuses {
		return resp.Statuses
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get all statuses.
//
// Get the status of all libraries on all clusters. A status is returned for all
// libraries installed on this cluster via the API or the libraries UI.
func (a *librariesImpl) AllClusterStatusesAll(ctx context.Context) ([]ClusterLibraryStatuses, error) {
	iterator := a.AllClusterStatuses(ctx)
	return listing.ToSlice[ClusterLibraryStatuses](ctx, iterator)
}

func (a *librariesImpl) internalAllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error) {

	var listAllClusterLibraryStatusesResponsePb listAllClusterLibraryStatusesResponsePb
	path := "/api/2.0/libraries/all-cluster-statuses"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listAllClusterLibraryStatusesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listAllClusterLibraryStatusesResponseFromPb(&listAllClusterLibraryStatusesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Get status.
//
// Get the status of libraries on a cluster. A status is returned for all
// libraries installed on this cluster via the API or the libraries UI. The
// order of returned libraries is as follows: 1. Libraries set to be installed
// on this cluster, in the order that the libraries were added to the cluster,
// are returned first. 2. Libraries that were previously requested to be
// installed on this cluster or, but are now marked for removal, in no
// particular order, are returned last.
func (a *librariesImpl) ClusterStatus(ctx context.Context, request ClusterStatus) listing.Iterator[LibraryFullStatus] {

	getNextPage := func(ctx context.Context, req ClusterStatus) (*ClusterLibraryStatuses, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalClusterStatus(ctx, req)
	}
	getItems := func(resp *ClusterLibraryStatuses) []LibraryFullStatus {
		return resp.LibraryStatuses
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get status.
//
// Get the status of libraries on a cluster. A status is returned for all
// libraries installed on this cluster via the API or the libraries UI. The
// order of returned libraries is as follows: 1. Libraries set to be installed
// on this cluster, in the order that the libraries were added to the cluster,
// are returned first. 2. Libraries that were previously requested to be
// installed on this cluster or, but are now marked for removal, in no
// particular order, are returned last.
func (a *librariesImpl) ClusterStatusAll(ctx context.Context, request ClusterStatus) ([]LibraryFullStatus, error) {
	iterator := a.ClusterStatus(ctx, request)
	return listing.ToSlice[LibraryFullStatus](ctx, iterator)
}

func (a *librariesImpl) internalClusterStatus(ctx context.Context, request ClusterStatus) (*ClusterLibraryStatuses, error) {

	requestPb, pbErr := clusterStatusToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var clusterLibraryStatusesPb clusterLibraryStatusesPb
	path := "/api/2.0/libraries/cluster-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &clusterLibraryStatusesPb)
	if err != nil {
		return nil, err
	}
	resp, err := clusterLibraryStatusesFromPb(&clusterLibraryStatusesPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *librariesImpl) Install(ctx context.Context, request InstallLibraries) error {

	requestPb, pbErr := installLibrariesToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var installLibrariesResponsePb installLibrariesResponsePb
	path := "/api/2.0/libraries/install"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &installLibrariesResponsePb)
	if err != nil {
		return err
	}

	return err
}

func (a *librariesImpl) Uninstall(ctx context.Context, request UninstallLibraries) error {

	requestPb, pbErr := uninstallLibrariesToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var uninstallLibrariesResponsePb uninstallLibrariesResponsePb
	path := "/api/2.0/libraries/uninstall"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &uninstallLibrariesResponsePb)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just PolicyComplianceForClusters API methods
type policyComplianceForClustersImpl struct {
	client *client.DatabricksClient
}

func (a *policyComplianceForClustersImpl) EnforceCompliance(ctx context.Context, request EnforceClusterComplianceRequest) (*EnforceClusterComplianceResponse, error) {

	requestPb, pbErr := enforceClusterComplianceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var enforceClusterComplianceResponsePb enforceClusterComplianceResponsePb
	path := "/api/2.0/policies/clusters/enforce-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, (*requestPb), &enforceClusterComplianceResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := enforceClusterComplianceResponseFromPb(&enforceClusterComplianceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *policyComplianceForClustersImpl) GetCompliance(ctx context.Context, request GetClusterComplianceRequest) (*GetClusterComplianceResponse, error) {

	requestPb, pbErr := getClusterComplianceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getClusterComplianceResponsePb getClusterComplianceResponsePb
	path := "/api/2.0/policies/clusters/get-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &getClusterComplianceResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := getClusterComplianceResponseFromPb(&getClusterComplianceResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List cluster policy compliance.
//
// Returns the policy compliance status of all clusters that use a given policy.
// Clusters could be out of compliance if their policy was updated after the
// cluster was last edited.
func (a *policyComplianceForClustersImpl) ListCompliance(ctx context.Context, request ListClusterCompliancesRequest) listing.Iterator[ClusterCompliance] {

	getNextPage := func(ctx context.Context, req ListClusterCompliancesRequest) (*ListClusterCompliancesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCompliance(ctx, req)
	}
	getItems := func(resp *ListClusterCompliancesResponse) []ClusterCompliance {
		return resp.Clusters
	}
	getNextReq := func(resp *ListClusterCompliancesResponse) *ListClusterCompliancesRequest {
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

// List cluster policy compliance.
//
// Returns the policy compliance status of all clusters that use a given policy.
// Clusters could be out of compliance if their policy was updated after the
// cluster was last edited.
func (a *policyComplianceForClustersImpl) ListComplianceAll(ctx context.Context, request ListClusterCompliancesRequest) ([]ClusterCompliance, error) {
	iterator := a.ListCompliance(ctx, request)
	return listing.ToSlice[ClusterCompliance](ctx, iterator)
}

func (a *policyComplianceForClustersImpl) internalListCompliance(ctx context.Context, request ListClusterCompliancesRequest) (*ListClusterCompliancesResponse, error) {

	requestPb, pbErr := listClusterCompliancesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listClusterCompliancesResponsePb listClusterCompliancesResponsePb
	path := "/api/2.0/policies/clusters/list-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listClusterCompliancesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listClusterCompliancesResponseFromPb(&listClusterCompliancesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just PolicyFamilies API methods
type policyFamiliesImpl struct {
	client *client.DatabricksClient
}

func (a *policyFamiliesImpl) Get(ctx context.Context, request GetPolicyFamilyRequest) (*PolicyFamily, error) {

	requestPb, pbErr := getPolicyFamilyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var policyFamilyPb policyFamilyPb
	path := fmt.Sprintf("/api/2.0/policy-families/%v", requestPb.PolicyFamilyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &policyFamilyPb)
	if err != nil {
		return nil, err
	}
	resp, err := policyFamilyFromPb(&policyFamilyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List policy families.
//
// Returns the list of policy definition types available to use at their latest
// version. This API is paginated.
func (a *policyFamiliesImpl) List(ctx context.Context, request ListPolicyFamiliesRequest) listing.Iterator[PolicyFamily] {

	getNextPage := func(ctx context.Context, req ListPolicyFamiliesRequest) (*ListPolicyFamiliesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListPolicyFamiliesResponse) []PolicyFamily {
		return resp.PolicyFamilies
	}
	getNextReq := func(resp *ListPolicyFamiliesResponse) *ListPolicyFamiliesRequest {
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

// List policy families.
//
// Returns the list of policy definition types available to use at their latest
// version. This API is paginated.
func (a *policyFamiliesImpl) ListAll(ctx context.Context, request ListPolicyFamiliesRequest) ([]PolicyFamily, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PolicyFamily](ctx, iterator)
}

func (a *policyFamiliesImpl) internalList(ctx context.Context, request ListPolicyFamiliesRequest) (*ListPolicyFamiliesResponse, error) {

	requestPb, pbErr := listPolicyFamiliesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listPolicyFamiliesResponsePb listPolicyFamiliesResponsePb
	path := "/api/2.0/policy-families"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, (*requestPb), &listPolicyFamiliesResponsePb)
	if err != nil {
		return nil, err
	}
	resp, err := listPolicyFamiliesResponseFromPb(&listPolicyFamiliesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
