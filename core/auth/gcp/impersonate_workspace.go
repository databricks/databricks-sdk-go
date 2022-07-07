package gcp

import (
	"context"
	"log"
	"net/http"

	"github.com/databricks/sdk-go/core/auth/internal"
	"github.com/databricks/sdk-go/core/client"
	"google.golang.org/api/option"
)

type ImpersonateWorkspaceCredentials struct {
	GoogleServiceAccount string `name:"google_service_account" env:"DATABRICKS_GOOGLE_SERVICE_ACCOUNT"`

	// options used to enable unit testing mode for OIDC
	opts []option.ClientOption
}

func (c ImpersonateWorkspaceCredentials) Name() string {
	return "google-workspace"
}

func (c ImpersonateWorkspaceCredentials) Configure(ctx context.Context, conf *client.ApiClient) (func(*http.Request) error, error) {
	host := internal.Host(conf.Config.Host)
	if c.GoogleServiceAccount == "" || !host.IsGcp() || host.IsAccountsClient() {
		return nil, nil
	}
	inner, err := idTokenSource(ctx, conf.Config.Host, c.GoogleServiceAccount, c.opts...)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Using Google Default Application Credentials for Workspace")
	return internal.RefreshableVisitor(inner), nil
}
