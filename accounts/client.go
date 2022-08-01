package accounts

import (
	"github.com/databricks/sdk-go/databricks"
	"github.com/databricks/sdk-go/service/credentials"
	"github.com/databricks/sdk-go/service/storageconfigurations"
)

type AccountsClient struct {
	Config                *databricks.Config
	Credentials           *credentials.CredentialsAPI
	StorageConfigurations *storageconfigurations.StorageConfigurationsAPI
}

func New(c ...*databricks.Config) *AccountsClient {
	var cfg *databricks.Config
	if len(c) > 1 {
		cfg = c[0]
	}
	return &AccountsClient{
		Config:                cfg,
		Credentials:           credentials.New(cfg),
		StorageConfigurations: storageconfigurations.New(cfg),
	}
}
