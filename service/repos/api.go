// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package repos

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// The Repos API allows users to manage their git repos. 
type ReposService interface {
    // Creates a repo in the workspace and links it to the remote Git repo 
    // specified. Note that repos created programmatically must be linked to a 
    // remote Git repo, unlike repos created in the browser. 
    CreateRepo(ctx context.Context, createRepoRequest CreateRepoRequest) (*GetRepoResponse, error)
    // Deletes the specified repo 
    DeleteRepo(ctx context.Context, deleteRepoRequest DeleteRepoRequest) error
    // Returns the repo with the given repo ID. 
    GetRepo(ctx context.Context, getRepoRequest GetRepoRequest) (*GetRepoResponse, error)
    // Returns repos that the calling user has Manage permissions on. Results 
    // are paginated with each page containing twenty repos. 
    ListRepos(ctx context.Context, listReposRequest ListReposRequest) (*GetReposResponse, error)
    // Updates the repo to a different branch or tag, or updates the repo to 
    // the latest commit on the same branch. 
    UpdateRepo(ctx context.Context, updateRepoRequest UpdateRepoRequest) error
	GetRepoByRepoId(ctx context.Context, repoId string) (*GetRepoResponse, error)
	DeleteRepoByRepoId(ctx context.Context, repoId string) error
}

func New(client *client.DatabricksClient) ReposService {
	return &ReposAPI{
		client: client,
	}
}

type ReposAPI struct {
	client *client.DatabricksClient
}

// Creates a repo in the workspace and links it to the remote Git repo 
// specified. Note that repos created programmatically must be linked to a 
// remote Git repo, unlike repos created in the browser. 
func (a *ReposAPI) CreateRepo(ctx context.Context, request CreateRepoRequest) (*GetRepoResponse, error) {
	var getRepoResponse GetRepoResponse
	path := "/repos"
	err := a.client.Post(ctx, path, request, &getRepoResponse)
	return &getRepoResponse, err
}

// Deletes the specified repo 
func (a *ReposAPI) DeleteRepo(ctx context.Context, request DeleteRepoRequest) error {
	path := fmt.Sprintf("/repos/%v", request.RepoId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Returns the repo with the given repo ID. 
func (a *ReposAPI) GetRepo(ctx context.Context, request GetRepoRequest) (*GetRepoResponse, error) {
	var getRepoResponse GetRepoResponse
	path := fmt.Sprintf("/repos/%v", request.RepoId)
	err := a.client.Get(ctx, path, request, &getRepoResponse)
	return &getRepoResponse, err
}

// Returns repos that the calling user has Manage permissions on. Results are 
// paginated with each page containing twenty repos. 
func (a *ReposAPI) ListRepos(ctx context.Context, request ListReposRequest) (*GetReposResponse, error) {
	var getReposResponse GetReposResponse
	path := "/repos"
	err := a.client.Get(ctx, path, request, &getReposResponse)
	return &getReposResponse, err
}

// Updates the repo to a different branch or tag, or updates the repo to the 
// latest commit on the same branch. 
func (a *ReposAPI) UpdateRepo(ctx context.Context, request UpdateRepoRequest) error {
	path := fmt.Sprintf("/repos/%v", request.RepoId)
	err := a.client.Patch(ctx, path, request)
	return err
}


func (a *ReposAPI) GetRepoByRepoId(ctx context.Context, repoId string) (*GetRepoResponse, error) {
	return a.GetRepo(ctx, GetRepoRequest{
		RepoId: repoId,
	})
}

func (a *ReposAPI) DeleteRepoByRepoId(ctx context.Context, repoId string) error {
	return a.DeleteRepo(ctx, DeleteRepoRequest{
		RepoId: repoId,
	})
}
