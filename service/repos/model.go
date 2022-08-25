// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package repos

// all definitions in this file are in alphabetical order

type CreateRepoRequest struct {
    
    Path string `json:"path,omitempty"`
    
    Provider string `json:"provider"`
    
    Url string `json:"url"`
}


type DeleteRepoRequest struct {
    // The ID for the corresponding repo to access. 
    RepoId string ` path:"repo_id"`
}


type GetRepoRequest struct {
    // The ID for the corresponding repo to access. 
    RepoId string ` path:"repo_id"`
}


type GetRepoResponse struct {
    
    Branch string `json:"branch,omitempty"`
    
    HeadCommitId string `json:"head_commit_id,omitempty"`
    
    Id int64 `json:"id,omitempty"`
    
    Path string `json:"path,omitempty"`
    
    Provider string `json:"provider,omitempty"`
    
    Url string `json:"url,omitempty"`
}


type GetReposResponse struct {
    
    NextPageToken string `json:"next_page_token,omitempty"`
    
    Repos []GetRepoResponse `json:"repos,omitempty"`
}


type ListReposRequest struct {
    // Token used to get the next page of results. If not specified, returns 
    // the first page of results as well as a next page token if there are more 
    // results. 
    NextPageToken string ` url:"next_page_token,omitempty"`
    // Filters repos that have paths starting with the given path prefix. 
    PathPrefix string ` url:"path_prefix,omitempty"`
}


type UpdateRepoRequest struct {
    // The ID for the corresponding repo to access. 
    RepoId string ` path:"repo_id"`
}

// Branch that the local version of the repo is checked out to. 

// SHA-1 hash representing the commit ID of the current HEAD of the repo. 

// ID of the repo object in the workspace. 

// Token that can be specified as a query parameter to the GET /repos endpoint 
// to retrieve the next page of results. 

// Desired path for the repo in the workspace. Must be in the format 
// /Repos/{folder}/{repo-name}. 

// Git provider. This field is case-insensitive. The available Git providers 
// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices, gitHubEnterprise, 
// bitbucketServer, gitLabEnterpriseEdition and awsCodeCommit. 

// URL of the Git repository to be linked. 

