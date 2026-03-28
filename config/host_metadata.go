package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/httpclient"
)

// HostMetadata holds the parsed response from the /.well-known/databricks-config discovery endpoint.
type HostMetadata struct {
	// OIDCEndpoint is the OIDC discovery URL for this host. For account hosts,
	// this may contain an {account_id} placeholder that callers must substitute.
	OIDCEndpoint string `json:"oidc_endpoint"`

	// AccountID is the Databricks account ID associated with this host, if available.
	AccountID string `json:"account_id"`

	// WorkspaceID is the Databricks workspace ID associated with this host, if available.
	WorkspaceID string `json:"workspace_id"`

	// Cloud is the cloud provider for this Databricks deployment (AWS, Azure, or GCP).
	Cloud environment.Cloud `json:"cloud"`

	// DefaultOIDCAudience is the default OIDC audience for token requests.
	// For workspace hosts: "https://<workspace_host>/oidc/v1/token"
	// For account/unified hosts: the resolved account ID.
	DefaultOIDCAudience string `json:"default_oidc_audience"`
}

// HostMetadataResolver, when set on [Config], overrides the default HTTP fetch
// of /.well-known/databricks-config during config resolution. The function
// receives the canonical host and returns metadata (or nil, nil if unavailable).
// This allows callers to provide cached metadata without the SDK making an HTTP call.
type HostMetadataResolver func(ctx context.Context, host string) (*HostMetadata, error)

// getHostMetadata fetches the raw Databricks well-known configuration from
// {host}/.well-known/databricks-config. The returned HostMetadata contains
// raw values with no substitution (e.g., {account_id} placeholders are left
// as-is). Callers are responsible for interpreting the result.
func getHostMetadata(ctx context.Context, host string, client *httpclient.ApiClient) (*HostMetadata, error) {
	discoveryURL := host + "/.well-known/databricks-config"
	var meta HostMetadata
	if err := client.Do(ctx, "GET", discoveryURL, httpclient.WithResponseUnmarshal(&meta)); err != nil {
		return nil, fmt.Errorf("fetching host metadata from %q: %w", discoveryURL, err)
	}
	return &meta, nil
}
