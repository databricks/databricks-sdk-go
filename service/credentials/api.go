// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package credentials

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewCredentialConfigurations(client *client.DatabricksClient) CredentialConfigurationsService {
	return &CredentialConfigurationsAPI{
		client: client,
	}
}

type CredentialConfigurationsAPI struct {
	client *client.DatabricksClient
}

// Create credential configuration
//
// Creates a Databricks credential configuration that represents cloud
// cross-account credentials for a specified account. Databricks uses this to
// set up network infrastructure properly to host Databricks clusters. For your
// AWS IAM role, you need to trust the External ID (the Databricks Account API
// account ID) in the returned credential object, and configure the required
// access policy.
//
// Save the response's `credentials_id` field, which is the ID for your new
// credential configuration object.
//
// For information about how to create a new workspace with this API, see
// [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html)
func (a *CredentialConfigurationsAPI) CreateCredentialConfig(ctx context.Context, request CreateCredentialRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", request.AccountId)
	err := a.client.Post(ctx, path, request, &credential)
	return &credential, err
}

// Delete credential configuration
//
// Deletes a Databricks credential configuration object for an account, both
// specified by ID. You cannot delete a credential that is associated with any
// workspace.
func (a *CredentialConfigurationsAPI) DeleteCredentialConfig(ctx context.Context, request DeleteCredentialConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", request.AccountId, request.CredentialsId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete credential configuration
//
// Deletes a Databricks credential configuration object for an account, both
// specified by ID. You cannot delete a credential that is associated with any
// workspace.
func (a *CredentialConfigurationsAPI) DeleteCredentialConfigByAccountIdAndCredentialsId(ctx context.Context, accountId string, credentialsId string) error {
	return a.DeleteCredentialConfig(ctx, DeleteCredentialConfigRequest{
		AccountId:     accountId,
		CredentialsId: credentialsId,
	})
}

// Get credential configuration
//
// Gets a Databricks credential configuration object for an account, both
// specified by ID.
func (a *CredentialConfigurationsAPI) GetCredentialConfig(ctx context.Context, request GetCredentialConfigRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", request.AccountId, request.CredentialsId)
	err := a.client.Get(ctx, path, request, &credential)
	return &credential, err
}

// Get credential configuration
//
// Gets a Databricks credential configuration object for an account, both
// specified by ID.
func (a *CredentialConfigurationsAPI) GetCredentialConfigByAccountIdAndCredentialsId(ctx context.Context, accountId string, credentialsId string) (*Credential, error) {
	return a.GetCredentialConfig(ctx, GetCredentialConfigRequest{
		AccountId:     accountId,
		CredentialsId: credentialsId,
	})
}

// Get all credential configurations
//
// Gets all Databricks credential configurations associated with an account
// specified by ID.
func (a *CredentialConfigurationsAPI) ListCredentials(ctx context.Context, request ListCredentialsRequest) ([]Credential, error) {
	var credentialList []Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", request.AccountId)
	err := a.client.Get(ctx, path, request, &credentialList)
	return credentialList, err
}

func (a *CredentialConfigurationsAPI) CredentialCredentialsNameToCredentialsIdMap(ctx context.Context, request ListCredentialsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListCredentials(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.CredentialsName] = v.CredentialsId
	}
	return mapping, nil
}

func (a *CredentialConfigurationsAPI) GetCredentialByCredentialsName(ctx context.Context, name string) (*Credential, error) {
	result, err := a.ListCredentials(ctx, ListCredentialsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.CredentialsName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Credential named '%s' does not exist", name)
}

// Get all credential configurations
//
// Gets all Databricks credential configurations associated with an account
// specified by ID.
func (a *CredentialConfigurationsAPI) ListCredentialsByAccountId(ctx context.Context, accountId string) ([]Credential, error) {
	return a.ListCredentials(ctx, ListCredentialsRequest{
		AccountId: accountId,
	})
}
