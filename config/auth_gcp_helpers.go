package config

import (
	"context"

	"github.com/databricks/databricks-sdk-go/logger"
)

// requiresGcpSaAccessToken determines whether a GCP SA access token is needed
// for the X-Databricks-GCP-SA-Access-Token header. It uses host metadata to
// determine if the host is an account-level host (no workspace_id), and falls
// back to checking AccountID when metadata is unavailable.
func requiresGcpSaAccessToken(ctx context.Context, cfg *Config) bool {
	if cfg.Host == "" {
		return cfg.AccountID != ""
	}
	if err := cfg.EnsureResolved(); err != nil {
		return cfg.AccountID != ""
	}
	meta, err := getHostMetadata(ctx, cfg.CanonicalHostName(), cfg.refreshClient)
	if err != nil {
		logger.Debugf(ctx, "Failed to fetch host metadata for GCP SA check: %v", err)
		return cfg.AccountID != ""
	}
	return meta.WorkspaceID == ""
}
