package config

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/config/credentials"
)

// recordingStrategy is a test helper that records whether Configure was called.
type recordingStrategy struct {
	name   string
	called bool
}

func (r *recordingStrategy) Name() string { return r.name }
func (r *recordingStrategy) Configure(_ context.Context, _ *Config) (credentials.CredentialsProvider, error) {
	r.called = true
	return nil, nil
}

// TestCredentialsChain_CloudFiltering_SkipsOnCloudMismatch verifies that the
// chain skips a cloud-specific strategy in auto-detect mode when the detected
// cloud does not match the strategy's required cloud.
func TestCredentialsChain_CloudFiltering_SkipsOnCloudMismatch(t *testing.T) {
	azureStrategy := &recordingStrategy{name: "azure-cli"}
	chain := &credentialsChain{
		strategies: []CredentialsStrategy{azureStrategy},
		cloudRequirements: map[string]environment.Cloud{
			"azure-cli": environment.CloudAzure,
		},
	}

	// GCP host: azure-cli must be skipped in auto-detect mode.
	cfg := &Config{Host: "https://xyz.gcp.databricks.com/", resolved: true}
	chain.Configure(context.Background(), cfg) //nolint:errcheck

	if azureStrategy.called {
		t.Error("azure-cli strategy was called on GCP host, want it to be skipped in auto-detect mode")
	}
}

// TestCredentialsChain_CloudFiltering_BypassesOnExplicitAuthType verifies that
// the cloud filter is bypassed when AuthType is explicitly set, so that a user
// can request "azure-cli" even on a GCP host.
func TestCredentialsChain_CloudFiltering_BypassesOnExplicitAuthType(t *testing.T) {
	azureStrategy := &recordingStrategy{name: "azure-cli"}
	chain := &credentialsChain{
		strategies: []CredentialsStrategy{azureStrategy},
		cloudRequirements: map[string]environment.Cloud{
			"azure-cli": environment.CloudAzure,
		},
	}

	// GCP host but auth_type is explicitly set: cloud filter must be bypassed.
	cfg := &Config{Host: "https://xyz.gcp.databricks.com/", AuthType: "azure-cli", resolved: true}
	chain.Configure(context.Background(), cfg) //nolint:errcheck

	if !azureStrategy.called {
		t.Error("azure-cli strategy was not called despite explicit auth_type on GCP host, want bypass of cloud filter")
	}
}

func TestDefaultCredentialStrategy(t *testing.T) {
	original := DefaultCredentialStrategyProvider
	t.Cleanup(func() { DefaultCredentialStrategyProvider = original })

	want := &DefaultCredentials{}
	DefaultCredentialStrategyProvider = func() CredentialsStrategy {
		return want
	}

	cfg := &Config{
		Host: "https://example.databricks.com",
	}
	cfg.Authenticate(&http.Request{Header: http.Header{}})

	if cfg.Credentials != want {
		t.Errorf("Credentials: got %v, want %v", cfg.Credentials, want)
	}
}

func TestDefaultCredentials_Configure(t *testing.T) {
	testCases := []struct {
		desc     string
		authType string
		wantErr  string
	}{
		{
			desc:     "unknown auth type",
			authType: "unknown-auth-type-1337",
			wantErr:  "auth type \"unknown-auth-type-1337\" not found",
		},
		{
			desc:     "not valid auth",
			authType: "",
			wantErr:  "cannot configure default credentials",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			cfg := &Config{
				AuthType: tc.authType,
				resolved: true, // avoid calling EnsureResolved
			}

			dc := DefaultCredentials{}
			got, gotErr := dc.Configure(ctx, cfg)

			if got != nil {
				t.Errorf("DefaultCredentials.Configure: got %v, want nil", got)
			}
			if gotErr == nil {
				t.Errorf("DefaultCredentials.Configure: got error %v, want non-nil", gotErr)
			}
			if !strings.Contains(gotErr.Error(), tc.wantErr) {
				t.Errorf("DefaultCredentials.Configure: got error %v, want error containing %q", gotErr, tc.wantErr)
			}
		})
	}
}
