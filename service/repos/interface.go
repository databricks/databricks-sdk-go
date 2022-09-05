// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package repos

import (
	"context"
)

// The Repos API allows users to manage their git repos.
type ReposService interface {
	// Creates a repo in the workspace and links it to the remote Git repo
	// specified. Note that repos created programmatically must be linked to a
	// remote Git repo, unlike repos created in the browser.
	Create(ctx context.Context, createRepo CreateRepo) (*RepoInfo, error)

	// Deletes the specified repo
	Delete(ctx context.Context, deleteRequest DeleteRequest) error

	DeleteByRepoId(ctx context.Context, repoId string) error
	// Returns the repo with the given repo ID.
	Get(ctx context.Context, getRequest GetRequest) (*RepoInfo, error)

	GetByRepoId(ctx context.Context, repoId string) (*RepoInfo, error)
	// Returns repos that the calling user has Manage permissions on. Results
	// are paginated with each page containing twenty repos.
	List(ctx context.Context, listRequest ListRequest) (*ListReposResponse, error)

	// Updates the repo to a different branch or tag, or updates the repo to the
	// latest commit on the same branch.
	Update(ctx context.Context, updateRepo UpdateRepo) error
}
