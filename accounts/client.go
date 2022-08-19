package accounts

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/credentials"
	"github.com/databricks/databricks-sdk-go/service/storageconfigurations"
)

type AccountsClient struct {
	Config                *databricks.Config
	Credentials           *credentials.CredentialsAPI
	StorageConfigurations *storageconfigurations.StorageConfigurationsAPI
}

func New(c ...*databricks.Config) *AccountsClient {
	var cfg *databricks.Config
	if len(c) == 1 {
		cfg = c[0]
	} else {
		cfg = &databricks.Config{}
	}
	return &AccountsClient{
		Config:                cfg,
		Credentials:           credentials.New(cfg),
		StorageConfigurations: storageconfigurations.New(cfg),
	}
}
