package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/authconv"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/impersonate"
)

// gcpServiceAccountAccessTokenHeader carries a Google Cloud OAuth access token
// (cloud-platform scope) that Databricks uses to provision GCP resources on the
// caller's behalf. It is a passthrough credential, independent of the identity
// in the Authorization header.
const gcpServiceAccountAccessTokenHeader = "X-Databricks-GCP-SA-Access-Token"

// GcpM2mCredentials authenticates the request identity with a Databricks OAuth
// service-principal (M2M) token, and additionally attaches a Google Cloud
// access token in the X-Databricks-GCP-SA-Access-Token header so Databricks can
// provision GCP resources on the caller's behalf.
//
// This is the supported way to call GCP account-level provisioning APIs when
// SSO is enabled: identity comes from a Databricks-governed service principal
// (Authorization header), while the Google access token supplies only the GCP
// resource credential.
//
// It must be selected explicitly via auth_type = "oauth-m2m-gcp". The mode
// combines the "oauth" (client_id/client_secret) and "google"
// (google_credentials/google_service_account) config groups, which the default
// single-auth-method conflict check rejects unless an auth type is set.
type GcpM2mCredentials struct {
	// googleTokenSource, when non-nil, is used as the Google Cloud access token
	// source instead of deriving one from the config. Test-only seam.
	googleTokenSource auth.TokenSource
}

func (c GcpM2mCredentials) Name() string {
	return "oauth-m2m-gcp"
}

func (c GcpM2mCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if !cfg.IsGcp() || cfg.ClientID == "" || cfg.ClientSecret == "" {
		return nil, nil
	}
	if c.googleTokenSource == nil && cfg.GoogleCredentials == "" && cfg.GoogleServiceAccount == "" {
		return nil, nil
	}

	primary, err := databricksOAuthTokenSource(ctx, cfg)
	if err != nil {
		return nil, err
	}

	secondary := c.googleTokenSource
	if secondary == nil {
		secondary, err = googleAccessTokenSource(ctx, cfg)
		if err != nil {
			return nil, err
		}
	}

	logger.Infof(ctx, "Using Databricks OAuth (M2M) with GCP service account access token passthrough")

	// Google token sources cache internally; disable async refresh to avoid
	// redundant work. The Databricks primary refreshes synchronously on expiry.
	opts := append(cacheOptions(cfg), auth.WithAsyncRefresh(false))
	// secondaryOptional is false: this mode exists to attach the GCP access
	// token, so a failure to obtain it must fail the request rather than
	// silently drop the header.
	visitor := serviceToServiceVisitor(primary, secondary, gcpServiceAccountAccessTokenHeader, false, opts...)
	return newVisitorOAuthCredentials(visitor, auth.NewCachedTokenSource(primary, opts...)), nil
}

// googleAccessTokenSource returns a token source for a Google Cloud OAuth
// access token (cloud-platform scope). It prefers a service-account JSON key
// (google_credentials) and falls back to service-account impersonation
// (google_service_account).
func googleAccessTokenSource(ctx context.Context, cfg *Config) (auth.TokenSource, error) {
	scopes := []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/compute",
	}
	switch {
	case cfg.GoogleCredentials != "":
		jsonBytes, err := readCredentials(cfg.GoogleCredentials)
		if err != nil {
			return nil, fmt.Errorf("could not read GoogleCredentials. "+
				"Make sure the file exists, or the JSON content is valid: %w", err)
		}
		creds, err := google.CredentialsFromJSON(ctx, jsonBytes, scopes...)
		if err != nil {
			return nil, fmt.Errorf("could not obtain GCP access token from JSON: %w", err)
		}
		return authconv.AuthTokenSource(creds.TokenSource), nil
	case cfg.GoogleServiceAccount != "":
		platform, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
			TargetPrincipal: cfg.GoogleServiceAccount,
			Scopes:          scopes,
		})
		if err != nil {
			return nil, fmt.Errorf("could not create GCP SA access token source: %w", err)
		}
		return authconv.AuthTokenSource(platform), nil
	default:
		return nil, fmt.Errorf("oauth-m2m-gcp requires google_credentials or google_service_account to be set")
	}
}
