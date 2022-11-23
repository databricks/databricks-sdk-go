// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package libraries

import (
	"context"
	"net/http"
)

type databricksClient interface {
	Do(ctx context.Context, method string, path string, request any, response any) error
	ConfiguredAccountID() string
	IsAws() bool
}

// unexported type that holds implementations of just Libraries API methods
type librariesImpl struct {
	client databricksClient
}

func (a *librariesImpl) AllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error) {
	var listAllClusterLibraryStatusesResponse ListAllClusterLibraryStatusesResponse
	path := "/api/2.0/libraries/all-cluster-statuses"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listAllClusterLibraryStatusesResponse)
	return &listAllClusterLibraryStatusesResponse, err
}

func (a *librariesImpl) ClusterStatus(ctx context.Context, request ClusterStatusRequest) (*ClusterLibraryStatuses, error) {
	var clusterLibraryStatuses ClusterLibraryStatuses
	path := "/api/2.0/libraries/cluster-status"
	err := a.client.Do(ctx, http.MethodGet, path, request, &clusterLibraryStatuses)
	return &clusterLibraryStatuses, err
}

func (a *librariesImpl) Install(ctx context.Context, request InstallLibraries) error {
	path := "/api/2.0/libraries/install"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *librariesImpl) Uninstall(ctx context.Context, request UninstallLibraries) error {
	path := "/api/2.0/libraries/uninstall"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}
