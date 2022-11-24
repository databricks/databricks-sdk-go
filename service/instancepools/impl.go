// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instancepools

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just InstancePools API methods
type instancePoolsImpl struct {
	client *client.DatabricksClient
}

func (a *instancePoolsImpl) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {
	var createInstancePoolResponse CreateInstancePoolResponse
	path := "/api/2.0/instance-pools/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createInstancePoolResponse)
	return &createInstancePoolResponse, err
}

func (a *instancePoolsImpl) Delete(ctx context.Context, request DeleteInstancePool) error {
	path := "/api/2.0/instance-pools/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *instancePoolsImpl) Edit(ctx context.Context, request EditInstancePool) error {
	path := "/api/2.0/instance-pools/edit"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *instancePoolsImpl) Get(ctx context.Context, request GetRequest) (*GetInstancePool, error) {
	var getInstancePool GetInstancePool
	path := "/api/2.0/instance-pools/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getInstancePool)
	return &getInstancePool, err
}

func (a *instancePoolsImpl) List(ctx context.Context) (*ListInstancePools, error) {
	var listInstancePools ListInstancePools
	path := "/api/2.0/instance-pools/list"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listInstancePools)
	return &listInstancePools, err
}
