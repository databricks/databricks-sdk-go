package databricks

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/databricks/databricks-sdk-go/internal"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

type GoogleCredentials struct {
}

func (c GoogleCredentials) Name() string {
	return "google-credentials"
}

func (c GoogleCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.GoogleCredentials == "" || !cfg.IsGcp() {
		return nil, nil
	}
	json, err := readCredentials(cfg.GoogleCredentials)
	if err != nil {
		err = fmt.Errorf("could not read GoogleCredentials. "+
			"Make sure the file exists, or the JSON content is valid: %w", err)
		return nil, err
	}
	inner, err := idtoken.NewTokenSource(ctx, cfg.Host, option.WithCredentialsJSON(json))
	if err != nil {
		return nil, fmt.Errorf("could not obtain OIDC token from JSON: %w", err)
	}
	// Obtain token source for creating Google Cloud Platform token.
	creds, err := google.CredentialsFromJSON(ctx, json,
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/compute")
	if err != nil {
		return nil, fmt.Errorf("could not obtain OAuth2 token from JSON: %w", err)
	}
	logger.Infof("Using Google Credentials")
	return internal.ServiceToServiceVisitor(inner, creds.TokenSource, "X-Databricks-GCP-SA-Access-Token"), nil
}

// Reads credentials as JSON. Credentials can be either a path to JSON file,
// or actual JSON string.
func readCredentials(credentials string) ([]byte, error) {
	// Try to read credentials as file path.
	if _, err := os.Stat(credentials); err == nil {
		jsonContents, err := ioutil.ReadFile(credentials)
		if err != nil {
			return jsonContents, err
		}
		return jsonContents, nil
	}
	// Assume that credential is actually JSON string.
	return []byte(credentials), nil
}

type GoogleDefaultCredentials struct {
	// options used to enable unit testing mode for OIDC
	opts []option.ClientOption
}

func (c GoogleDefaultCredentials) Name() string {
	return "google-id"
}

func (c GoogleDefaultCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.GoogleServiceAccount == "" || !cfg.IsGcp() {
		return nil, nil
	}
	inner, err := c.idTokenSource(ctx, cfg.Host, cfg.GoogleServiceAccount, c.opts...)
	if err != nil {
		return nil, err
	}
	if !cfg.IsAccountClient() {
		logger.Infof("Using Google Default Application Credentials for Workspace")
		return internal.RefreshableVisitor(inner), nil
	}
	// source for generateAccessToken
	platform, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: cfg.GoogleServiceAccount,
		Scopes: []string{
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/compute",
		},
	}, c.opts...)
	if err != nil {
		return nil, err
	}
	logger.Infof("Using Google Default Application Credentials for Accounts API")
	return internal.ServiceToServiceVisitor(inner, platform, "X-Databricks-GCP-SA-Access-Token"), nil
}

func (c GoogleDefaultCredentials) idTokenSource(ctx context.Context, host, serviceAccount string,
	opts ...option.ClientOption) (oauth2.TokenSource, error) {
	ts, err := impersonate.IDTokenSource(ctx, impersonate.IDTokenConfig{
		Audience:        host,
		TargetPrincipal: serviceAccount,
		IncludeEmail:    true,
	}, opts...)
	if err != nil {
		err = fmt.Errorf("could not obtain OIDC token. %w Running 'gcloud auth application-default login' may help", err)
		return nil, err
	}
	return ts, err
}
