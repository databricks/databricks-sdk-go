// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package transitionrequests

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

type TransitionrequestsService interface {
	// Approve model version stage transition request.
	ApproveTransitionRequest(ctx context.Context, approveTransitionRequestRequest ApproveTransitionRequestRequest) (*ApproveTransitionRequestResponse, error)
	// Make a model version stage transition request.
	CreateTransitionRequest(ctx context.Context, createTransitionRequestRequest CreateTransitionRequestRequest) (*CreateTransitionRequestResponse, error)
	// Cancel model version stage transition request.
	DeleteTransitionRequest(ctx context.Context, deleteTransitionRequestRequest DeleteTransitionRequestRequest) error
	// Get all open stage transition requests for the model version.
	GetTransitionRequests(ctx context.Context, getTransitionRequestsRequest GetTransitionRequestsRequest) (*GetTransitionRequestsResponse, error)
	// Reject model version stage transition request.
	RejectTransitionRequest(ctx context.Context, rejectTransitionRequestRequest RejectTransitionRequestRequest) (*RejectTransitionRequestResponse, error)
}

func New(client *client.DatabricksClient) TransitionrequestsService {
	return &TransitionrequestsAPI{
		client: client,
	}
}

type TransitionrequestsAPI struct {
	client *client.DatabricksClient
}

// Approve model version stage transition request.
func (a *TransitionrequestsAPI) ApproveTransitionRequest(ctx context.Context, request ApproveTransitionRequestRequest) (*ApproveTransitionRequestResponse, error) {
	var approveTransitionRequestResponse ApproveTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/approve"
	err := a.client.Post(ctx, path, request, &approveTransitionRequestResponse)
	return &approveTransitionRequestResponse, err
}

// Make a model version stage transition request.
func (a *TransitionrequestsAPI) CreateTransitionRequest(ctx context.Context, request CreateTransitionRequestRequest) (*CreateTransitionRequestResponse, error) {
	var createTransitionRequestResponse CreateTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/create"
	err := a.client.Post(ctx, path, request, &createTransitionRequestResponse)
	return &createTransitionRequestResponse, err
}

// Cancel model version stage transition request.
func (a *TransitionrequestsAPI) DeleteTransitionRequest(ctx context.Context, request DeleteTransitionRequestRequest) error {
	path := "/api/2.0/mlflow/transition-requests/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get all open stage transition requests for the model version.
func (a *TransitionrequestsAPI) GetTransitionRequests(ctx context.Context, request GetTransitionRequestsRequest) (*GetTransitionRequestsResponse, error) {
	var getTransitionRequestsResponse GetTransitionRequestsResponse
	path := "/api/2.0/mlflow/transition-requests/list"
	err := a.client.Get(ctx, path, request, &getTransitionRequestsResponse)
	return &getTransitionRequestsResponse, err
}

// Reject model version stage transition request.
func (a *TransitionrequestsAPI) RejectTransitionRequest(ctx context.Context, request RejectTransitionRequestRequest) (*RejectTransitionRequestResponse, error) {
	var rejectTransitionRequestResponse RejectTransitionRequestResponse
	path := "/api/2.0/mlflow/transition-requests/reject"
	err := a.client.Post(ctx, path, request, &rejectTransitionRequestResponse)
	return &rejectTransitionRequestResponse, err
}
