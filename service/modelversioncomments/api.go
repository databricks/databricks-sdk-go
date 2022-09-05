// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package modelversioncomments

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

type ModelversioncommentsService interface {
	// Make a comment on a model version. A comment can be submitted either by a
	// user or programmatically to display relevant information about the model.
	// For example, test results or deployment errors.
	CreateComment(ctx context.Context, createCommentRequest CreateCommentRequest) (*CreateCommentResponse, error)
	// Delete a comment on a model version.
	DeleteComment(ctx context.Context, deleteCommentRequest DeleteCommentRequest) error
	// Edit a comment on a model version.
	UpdateComment(ctx context.Context, updateCommentRequest UpdateCommentRequest) (*UpdateCommentResponse, error)
}

func New(client *client.DatabricksClient) ModelversioncommentsService {
	return &ModelversioncommentsAPI{
		client: client,
	}
}

type ModelversioncommentsAPI struct {
	client *client.DatabricksClient
}

// Make a comment on a model version. A comment can be submitted either by a
// user or programmatically to display relevant information about the model. For
// example, test results or deployment errors.
func (a *ModelversioncommentsAPI) CreateComment(ctx context.Context, request CreateCommentRequest) (*CreateCommentResponse, error) {
	var createCommentResponse CreateCommentResponse
	path := "/api/2.0/mlflow/comments/create"
	err := a.client.Post(ctx, path, request, &createCommentResponse)
	return &createCommentResponse, err
}

// Delete a comment on a model version.
func (a *ModelversioncommentsAPI) DeleteComment(ctx context.Context, request DeleteCommentRequest) error {
	path := "/api/2.0/mlflow/comments/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

// Edit a comment on a model version.
func (a *ModelversioncommentsAPI) UpdateComment(ctx context.Context, request UpdateCommentRequest) (*UpdateCommentResponse, error) {
	var updateCommentResponse UpdateCommentResponse
	path := "/api/2.0/mlflow/comments/update"
	err := a.client.Post(ctx, path, request, &updateCommentResponse)
	return &updateCommentResponse, err
}
