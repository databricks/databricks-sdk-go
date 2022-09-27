package databricks

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/databricks-sdk-go/databricks/internal"
	"github.com/databricks/databricks-sdk-go/databricks/logger"
)

type DatabricksOAuthM2M struct {
}

func (c DatabricksOAuthM2M) Name() string {
	return "oauth-m2m"
}

func (c DatabricksOAuthM2M) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.ClientID == "" || cfg.ClientSecret == "" {
		return nil, nil
	}
	logger.Infof("Generating Databricks OAuth token for Service Principal (%s)", cfg.ClientID)
	ts := (&clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		AuthStyle:    oauth2.AuthStyleInHeader,
		TokenURL:     "https://login.staging.databricks.com/oauth2/aus28jhkb8qmokpLr1d7/v1/token",
		Scopes:       []string{"all-apis"},
	}).TokenSource(ctx)
	return internal.RefreshableVisitor(ts), nil
}
