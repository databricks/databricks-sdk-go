// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package repos

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just Repos API methods
type reposImpl struct {
	client *client.DatabricksClient
}

func (a *reposImpl) Create(ctx context.Context, request CreateRepo) (*RepoInfo, error) {
	var repoInfo RepoInfo
	path := "/api/2.0/repos"
	err := a.client.Post(ctx, path, request, &repoInfo)
	return &repoInfo, err
}

func (a *reposImpl) Delete(ctx context.Context, request DeleteRequest) error {
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *reposImpl) Get(ctx context.Context, request GetRequest) (*RepoInfo, error) {
	var repoInfo RepoInfo
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	err := a.client.Get(ctx, path, request, &repoInfo)
	return &repoInfo, err
}

func (a *reposImpl) List(ctx context.Context, request ListRequest) (*ListReposResponse, error) {
	var listReposResponse ListReposResponse
	path := "/api/2.0/repos"
	err := a.client.Get(ctx, path, request, &listReposResponse)
	return &listReposResponse, err
}

func (a *reposImpl) Update(ctx context.Context, request UpdateRepo) error {
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	err := a.client.Patch(ctx, path, request)
	return err
}
