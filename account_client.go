// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package databricks

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/config"
)

type AccountClient struct {
	Config *config.Config
}

var ErrNotAccountClient = errors.New("invalid Databricks Account configuration")

// NewAccountClient creates new Databricks SDK client for Accounts or returns
// error in case configuration is wrong
func NewAccountClient(c ...*Config) (*AccountClient, error) {
	var cfg *config.Config
	if len(c) == 1 {
		// first config
		cfg = (*config.Config)(c[0])
	} else {
		// default config
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, ErrNotAccountClient
	}

	return &AccountClient{
		Config: cfg,
	}, nil
}
