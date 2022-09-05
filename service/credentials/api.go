// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package credentials

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// These APIs manage credential configurations for this workspace. Databricks
// needs access to a cross-account service IAM role in your AWS account so that
// Databricks can deploy clusters in the appropriate VPC for the new workspace.
// A credential configuration encapsulates this role information, and its ID is
// used when creating a new workspace.
type CredentialsService interface {
	// Create a Databricks credential configuration that represents cloud
	// cross-account credentials for a specified account. Databricks uses this
	// to set up network infrastructure properly to host Databricks clusters.
	// For your AWS IAM role, you need to trust the External ID (the Databricks
	// Account API account ID) in the returned credential object, and configure
	// the required access policy. Save the response&#39;s `credentials_id` field,
	// which is the ID for your new credential configuration object. For
	// detailed instructions of creating a new workspace with this API, see
	// [Create a new workspace using the Account
	// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html)
	CreateCredentialConfig(ctx context.Context, createCredentialRequest CreateCredentialRequest) (*Credential, error)
	// Delete a Databricks credential configuration object for an account, both
	// specified by ID. You cannot delete a credential that is associated with
	// any workspace.
	DeleteCredentialConfig(ctx context.Context, deleteCredentialConfigRequest DeleteCredentialConfigRequest) error
	// Get a Databricks credential configuration object for an account, both
	// specified by ID.
	GetCredentialConfig(ctx context.Context, getCredentialConfigRequest GetCredentialConfigRequest) (*Credential, error)
	// Get all Databricks credential configurations associated with an account
	// specified by ID.
	ListCredentials(ctx context.Context, listCredentialsRequest ListCredentialsRequest) (*ListCredentialsResponse, error)
	GetCredentialConfigByAccountIdAndCredentialsId(ctx context.Context, accountId string, credentialsId string) (*Credential, error)
	DeleteCredentialConfigByAccountIdAndCredentialsId(ctx context.Context, accountId string, credentialsId string) error
	ListCredentialsByAccountId(ctx context.Context, accountId string) (*ListCredentialsResponse, error)
}

func New(client *client.DatabricksClient) CredentialsService {
	return &CredentialsAPI{
		client: client,
	}
}

type CredentialsAPI struct {
	client *client.DatabricksClient
}

// Create a Databricks credential configuration that represents cloud
// cross-account credentials for a specified account. Databricks uses this to
// set up network infrastructure properly to host Databricks clusters. For your
// AWS IAM role, you need to trust the External ID (the Databricks Account API
// account ID) in the returned credential object, and configure the required
// access policy. Save the response&#39;s `credentials_id` field, which is the ID
// for your new credential configuration object. For detailed instructions of
// creating a new workspace with this API, see [Create a new workspace using the
// Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html)
func (a *CredentialsAPI) CreateCredentialConfig(ctx context.Context, request CreateCredentialRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/accounts/%v/credentials", request.AccountId)
	err := a.client.Post(ctx, path, request, &credential)
	return &credential, err
}

// Delete a Databricks credential configuration object for an account, both
// specified by ID. You cannot delete a credential that is associated with any
// workspace.
func (a *CredentialsAPI) DeleteCredentialConfig(ctx context.Context, request DeleteCredentialConfigRequest) error {
	path := fmt.Sprintf("/accounts/%v/credentials/%v", request.AccountId, request.CredentialsId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get a Databricks credential configuration object for an account, both
// specified by ID.
func (a *CredentialsAPI) GetCredentialConfig(ctx context.Context, request GetCredentialConfigRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/accounts/%v/credentials/%v", request.AccountId, request.CredentialsId)
	err := a.client.Get(ctx, path, request, &credential)
	return &credential, err
}

// Get all Databricks credential configurations associated with an account
// specified by ID.
func (a *CredentialsAPI) ListCredentials(ctx context.Context, request ListCredentialsRequest) (*ListCredentialsResponse, error) {
	var listCredentialsResponse ListCredentialsResponse
	path := fmt.Sprintf("/accounts/%v/credentials", request.AccountId)
	err := a.client.Get(ctx, path, request, &listCredentialsResponse)
	return &listCredentialsResponse, err
}

func (a *CredentialsAPI) GetCredentialConfigByAccountIdAndCredentialsId(ctx context.Context, accountId string, credentialsId string) (*Credential, error) {
	return a.GetCredentialConfig(ctx, GetCredentialConfigRequest{
		AccountId:     accountId,
		CredentialsId: credentialsId,
	})
}

func (a *CredentialsAPI) DeleteCredentialConfigByAccountIdAndCredentialsId(ctx context.Context, accountId string, credentialsId string) error {
	return a.DeleteCredentialConfig(ctx, DeleteCredentialConfigRequest{
		AccountId:     accountId,
		CredentialsId: credentialsId,
	})
}

func (a *CredentialsAPI) ListCredentialsByAccountId(ctx context.Context, accountId string) (*ListCredentialsResponse, error) {
	return a.ListCredentials(ctx, ListCredentialsRequest{
		AccountId: accountId,
	})
}
