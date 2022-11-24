// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusters

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Clusters API methods
type clustersImpl struct {
	client *client.DatabricksClient
}

func (a *clustersImpl) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {
	path := "/api/2.0/clusters/change-owner"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error) {
	var createClusterResponse CreateClusterResponse
	path := "/api/2.0/clusters/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createClusterResponse)
	return &createClusterResponse, err
}

func (a *clustersImpl) Delete(ctx context.Context, request DeleteCluster) error {
	path := "/api/2.0/clusters/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Edit(ctx context.Context, request EditCluster) error {
	path := "/api/2.0/clusters/edit"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Events(ctx context.Context, request GetEvents) (*GetEventsResponse, error) {
	var getEventsResponse GetEventsResponse
	path := "/api/2.0/clusters/events"
	err := a.client.Do(ctx, http.MethodPost, path, request, &getEventsResponse)
	return &getEventsResponse, err
}

func (a *clustersImpl) Get(ctx context.Context, request GetRequest) (*ClusterInfo, error) {
	var clusterInfo ClusterInfo
	path := "/api/2.0/clusters/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &clusterInfo)
	return &clusterInfo, err
}

func (a *clustersImpl) List(ctx context.Context, request ListRequest) (*ListClustersResponse, error) {
	var listClustersResponse ListClustersResponse
	path := "/api/2.0/clusters/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listClustersResponse)
	return &listClustersResponse, err
}

func (a *clustersImpl) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {
	var listNodeTypesResponse ListNodeTypesResponse
	path := "/api/2.0/clusters/list-node-types"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listNodeTypesResponse)
	return &listNodeTypesResponse, err
}

func (a *clustersImpl) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {
	var listAvailableZonesResponse ListAvailableZonesResponse
	path := "/api/2.0/clusters/list-zones"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listAvailableZonesResponse)
	return &listAvailableZonesResponse, err
}

func (a *clustersImpl) PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error {
	path := "/api/2.0/clusters/permanent-delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Pin(ctx context.Context, request PinCluster) error {
	path := "/api/2.0/clusters/pin"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Resize(ctx context.Context, request ResizeCluster) error {
	path := "/api/2.0/clusters/resize"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Restart(ctx context.Context, request RestartCluster) error {
	path := "/api/2.0/clusters/restart"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error) {
	var getSparkVersionsResponse GetSparkVersionsResponse
	path := "/api/2.0/clusters/spark-versions"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getSparkVersionsResponse)
	return &getSparkVersionsResponse, err
}

func (a *clustersImpl) Start(ctx context.Context, request StartCluster) error {
	path := "/api/2.0/clusters/start"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Unpin(ctx context.Context, request UnpinCluster) error {
	path := "/api/2.0/clusters/unpin"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

// unexported type that holds implementations of just InstanceProfiles API methods
type instanceProfilesImpl struct {
	client *client.DatabricksClient
}

func (a *instanceProfilesImpl) Add(ctx context.Context, request AddInstanceProfile) error {
	path := "/api/2.0/instance-profiles/add"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *instanceProfilesImpl) Edit(ctx context.Context, request InstanceProfile) error {
	path := "/api/2.0/instance-profiles/edit"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *instanceProfilesImpl) List(ctx context.Context) (*ListInstanceProfilesResponse, error) {
	var listInstanceProfilesResponse ListInstanceProfilesResponse
	path := "/api/2.0/instance-profiles/list"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listInstanceProfilesResponse)
	return &listInstanceProfilesResponse, err
}

func (a *instanceProfilesImpl) Remove(ctx context.Context, request RemoveInstanceProfile) error {
	path := "/api/2.0/instance-profiles/remove"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}
