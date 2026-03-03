package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

// hostMetadata holds the parsed response from the /.well-known/databricks-config discovery endpoint.
type hostMetadata struct {
	// OIDCEndpoint is the OIDC discovery URL for this host. For account hosts,
	// this may contain an {account_id} placeholder that callers must substitute.
	OIDCEndpoint string `json:"oidc_endpoint"`

	// AccountID is the Databricks account ID associated with this host, if available.
	AccountID string `json:"account_id"`

	// WorkspaceID is the Databricks workspace ID associated with this host, if available.
	WorkspaceID string `json:"workspace_id"`
}

// getHostMetadata fetches the raw Databricks well-known configuration from
// {host}/.well-known/databricks-config. The returned hostMetadata contains
// raw values with no substitution (e.g., {account_id} placeholders are left
// as-is). Callers are responsible for interpreting the result.
func getHostMetadata(ctx context.Context, host string, client *httpclient.ApiClient) (*hostMetadata, error) {
	discoveryURL := host + "/.well-known/databricks-config"
	var meta hostMetadata
	if err := client.Do(ctx, "GET", discoveryURL, httpclient.WithResponseUnmarshal(&meta)); err != nil {
		return nil, fmt.Errorf("fetching host metadata from %q: %w", discoveryURL, err)
	}
	return &meta, nil
}
