package gcp

import (
	"context"
	"log"
	"net/http"

	"github.com/databricks/sdk-go/core/auth/internal"
	"github.com/databricks/sdk-go/core/client"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

type ImpersonateAccountsCredentials struct {
	GoogleServiceAccount string `name:"google_service_account" env:"DATABRICKS_GOOGLE_SERVICE_ACCOUNT"`

	// options used to enable unit testing mode for OIDC
	opts []option.ClientOption
}

func (c ImpersonateAccountsCredentials) Name() string {
	return "google-accounts"
}

func (c ImpersonateAccountsCredentials) Configure(ctx context.Context, conf *client.ApiClient) (func(*http.Request) error, error) {
	host := internal.Host(conf.Config.Host)
	if c.GoogleServiceAccount == "" || !host.IsGcp() || !host.IsAccountsClient() {
		return nil, nil
	}
	inner, err := idTokenSource(ctx, conf.Config.Host, c.GoogleServiceAccount, c.opts...)
	if err != nil {
		return nil, err
	}
	// source for generateAccessToken
	platform, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: c.GoogleServiceAccount,
		Scopes: []string{
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/compute",
		},
	}, c.opts...)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Using Google Default Application Credentials for Accounts API")
	return internal.ServiceToServiceVisitor(inner, platform, "X-Databricks-GCP-SA-Access-Token"), nil
}
