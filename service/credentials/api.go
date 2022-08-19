package credentials

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// New creates ClustersAPI instance from provider meta
func New(cfg *databricks.Config) *CredentialsAPI {
	return &CredentialsAPI{
		client: client.New(cfg),
	}
}

type CredentialsAPI struct {
	client *client.DatabricksClient
}

// Create creates a set of Credentials for the cross account role
func (a CredentialsAPI) Create(ctx context.Context, credentialsName string, roleArn string) (Credentials, error) {
	var creds Credentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", a.client.Config.AccountID)
	err := a.client.Post(ctx, credentialsAPIPath, Credentials{
		CredentialsName: credentialsName,
		AwsCredentials: &AwsCredentials{
			StsRole: &StsRole{
				RoleArn: roleArn,
			},
		},
	}, &creds)
	return creds, err
}

// Read returns the credentials object along with metadata
func (a CredentialsAPI) Read(ctx context.Context, credentialsID string) (Credentials, error) {
	var creds Credentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", a.client.Config.AccountID, credentialsID)
	err := a.client.Get(ctx, credentialsAPIPath, nil, &creds)
	return creds, err
}

// Delete deletes the credentials object given a credentials id
func (a CredentialsAPI) Delete(ctx context.Context, credentialsID string) error {
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", a.client.Config.AccountID, credentialsID)
	return a.client.Delete(ctx, credentialsAPIPath, nil)
}

// List lists all the available credentials object in the account
func (a CredentialsAPI) List(ctx context.Context) ([]Credentials, error) {
	var credsList []Credentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", a.client.Config.AccountID)
	err := a.client.Get(ctx, credentialsAPIPath, nil, &credsList)
	return credsList, err
}
