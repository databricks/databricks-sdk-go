// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package repos

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewRepos(client *client.DatabricksClient) ReposService {
	return &ReposAPI{
		client: client,
	}
}

type ReposAPI struct {
	client *client.DatabricksClient
}

// Create a repo
//
// Creates a repo in the workspace and links it to the remote Git repo
// specified. Note that repos created programmatically must be linked to a
// remote Git repo, unlike repos created in the browser.
func (a *ReposAPI) Create(ctx context.Context, request CreateRepo) (*RepoInfo, error) {
	var repoInfo RepoInfo
	path := "/api/2.0/repos"
	err := a.client.Post(ctx, path, request, &repoInfo)
	return &repoInfo, err
}

// Delete a repo
//
// Deletes the specified repo.
func (a *ReposAPI) Delete(ctx context.Context, request DeleteRequest) error {
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a repo
//
// Deletes the specified repo.
func (a *ReposAPI) DeleteByRepoId(ctx context.Context, repoId int64) error {
	return a.Delete(ctx, DeleteRequest{
		RepoId: repoId,
	})
}

// Get a repo
//
// Returns the repo with the given repo ID.
func (a *ReposAPI) Get(ctx context.Context, request GetRequest) (*RepoInfo, error) {
	var repoInfo RepoInfo
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	err := a.client.Get(ctx, path, request, &repoInfo)
	return &repoInfo, err
}

// Get a repo
//
// Returns the repo with the given repo ID.
func (a *ReposAPI) GetByRepoId(ctx context.Context, repoId int64) (*RepoInfo, error) {
	return a.Get(ctx, GetRequest{
		RepoId: repoId,
	})
}

// Get repos
//
// Returns repos that the calling user has Manage permissions on. Results are
// paginated with each page containing twenty repos.
func (a *ReposAPI) List(ctx context.Context, request ListRequest) (*ListReposResponse, error) {
	var listReposResponse ListReposResponse
	path := "/api/2.0/repos"
	err := a.client.Get(ctx, path, request, &listReposResponse)
	return &listReposResponse, err
}

// Update a repo
//
// Updates the repo to a different branch or tag, or updates the repo to the
// latest commit on the same branch.
func (a *ReposAPI) Update(ctx context.Context, request UpdateRepo) error {
	path := fmt.Sprintf("/api/2.0/repos/%v", request.RepoId)
	err := a.client.Patch(ctx, path, request)
	return err
}
