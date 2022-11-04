// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package repos

// all definitions in this file are in alphabetical order

type CreateRepo struct {
	// Desired path for the repo in the workspace. Must be in the format
	// /Repos/{folder}/{repo-name}.
	Path string `json:"path,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	Provider string `json:"provider"`
	// URL of the Git repository to be linked.
	Url string `json:"url"`
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the GET /repos
	// endpoint to retrieve the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Repos []RepoInfo `json:"repos,omitempty"`
}

type RepoInfo struct {
	// Branch that the local version of the repo is checked out to.
	Branch string `json:"branch,omitempty"`
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	HeadCommitId string `json:"head_commit_id,omitempty"`
	// ID of the repo object in the workspace.
	Id int64 `json:"id,omitempty"`
	// Desired path for the repo in the workspace. Must be in the format
	// /Repos/{folder}/{repo-name}.
	Path string `json:"path,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	Provider string `json:"provider,omitempty"`
	// URL of the Git repository to be linked.
	Url string `json:"url,omitempty"`
}

type UpdateRepo struct {
	// Branch that the local version of the repo is checked out to.
	Branch string `json:"branch,omitempty"`
	// The ID for the corresponding repo to access.
	RepoId int64 `json:"-" path:"repo_id"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	Tag string `json:"tag,omitempty"`
}

// Branch that the local version of the repo is checked out to.

type DeleteRequest struct {
	// The ID for the corresponding repo to access.
	RepoId int64 `json:"-" path:"repo_id"`
}

type GetRequest struct {
	// The ID for the corresponding repo to access.
	RepoId int64 `json:"-" path:"repo_id"`
}

// SHA-1 hash representing the commit ID of the current HEAD of the repo.

// ID of the repo object in the workspace.

type ListRequest struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	NextPageToken string `json:"-" url:"next_page_token,omitempty"`
	// Filters repos that have paths starting with the given path prefix.
	PathPrefix string `json:"-" url:"path_prefix,omitempty"`
}

// Token that can be specified as a query parameter to the GET /repos endpoint
// to retrieve the next page of results.

// Desired path for the repo in the workspace. Must be in the format
// /Repos/{folder}/{repo-name}.

// Git provider. This field is case-insensitive. The available Git providers are
// gitHub, bitbucketCloud, gitLab, azureDevOpsServices, gitHubEnterprise,
// bitbucketServer, gitLabEnterpriseEdition and awsCodeCommit.

// Tag that the local version of the repo is checked out to. Updating the repo
// to a tag puts the repo in a detached HEAD state. Before committing new
// changes, you must update the repo to a branch instead of the detached HEAD.

// URL of the Git repository to be linked.
