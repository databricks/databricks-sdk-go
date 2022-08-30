// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package modelregistrywebhooks

import (
	"context"
	

	"github.com/databricks/databricks-sdk-go/databricks/client"
)


type ModelregistrywebhooksService interface {
    // This endpoint is in Public Preview. Create a registry webhook.
    CreateRegistryWebhook(ctx context.Context, createRegistryWebhookRequest CreateRegistryWebhookRequest) (*CreateRegistryWebhookResponse, error)
    // This endpoint is in Public Preview. Delete a registry webhook.
    DeleteRegistryWebhook(ctx context.Context, deleteRegistryWebhookRequest DeleteRegistryWebhookRequest) error
    // This endpoint is in Public Preview. List registry webhooks.
    ListRegistryWebhooks(ctx context.Context, listRegistryWebhooksRequest ListRegistryWebhooksRequest) (*ListRegistryWebhooksResponse, error)
    // This endpoint is in Public Preview. Test a registry webhook.
    TestRegistryWebhook(ctx context.Context, testRegistryWebhookRequest TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error)
    // This endpoint is in Public Preview. Update a registry webhook.
    UpdateRegistryWebhook(ctx context.Context, updateRegistryWebhookRequest UpdateRegistryWebhookRequest) error
}

func New(client *client.DatabricksClient) ModelregistrywebhooksService {
	return &ModelregistrywebhooksAPI{
		client: client,
	}
}

type ModelregistrywebhooksAPI struct {
	client *client.DatabricksClient
}

// This endpoint is in Public Preview. Create a registry webhook.
func (a *ModelregistrywebhooksAPI) CreateRegistryWebhook(ctx context.Context, request CreateRegistryWebhookRequest) (*CreateRegistryWebhookResponse, error) {
	var createRegistryWebhookResponse CreateRegistryWebhookResponse
	path := "/api/2.0/mlflow/registry-webhooks/create"
	err := a.client.Post(ctx, path, request, &createRegistryWebhookResponse)
	return &createRegistryWebhookResponse, err
}

// This endpoint is in Public Preview. Delete a registry webhook.
func (a *ModelregistrywebhooksAPI) DeleteRegistryWebhook(ctx context.Context, request DeleteRegistryWebhookRequest) error {
	path := "/api/2.0/mlflow/registry-webhooks/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

// This endpoint is in Public Preview. List registry webhooks.
func (a *ModelregistrywebhooksAPI) ListRegistryWebhooks(ctx context.Context, request ListRegistryWebhooksRequest) (*ListRegistryWebhooksResponse, error) {
	var listRegistryWebhooksResponse ListRegistryWebhooksResponse
	path := "/api/2.0/mlflow/registry-webhooks/list"
	err := a.client.Get(ctx, path, request, &listRegistryWebhooksResponse)
	return &listRegistryWebhooksResponse, err
}

// This endpoint is in Public Preview. Test a registry webhook.
func (a *ModelregistrywebhooksAPI) TestRegistryWebhook(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error) {
	var testRegistryWebhookResponse TestRegistryWebhookResponse
	path := "/api/2.0/mlflow/registry-webhooks/test"
	err := a.client.Post(ctx, path, request, &testRegistryWebhookResponse)
	return &testRegistryWebhookResponse, err
}

// This endpoint is in Public Preview. Update a registry webhook.
func (a *ModelregistrywebhooksAPI) UpdateRegistryWebhook(ctx context.Context, request UpdateRegistryWebhookRequest) error {
	path := "/api/2.0/mlflow/registry-webhooks/update"
	err := a.client.Patch(ctx, path, request)
	return err
}

