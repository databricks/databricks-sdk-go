package internal

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// captureHeadersTransport snapshots the request and response headers for a
// single matching URL path and lets every other request through unchanged.
type captureHeadersTransport struct {
	inner http.RoundTripper
	path  string

	reqHdr  http.Header
	respHdr http.Header
}

func (t *captureHeadersTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	matched := req.URL.Path == t.path
	if matched {
		t.reqHdr = req.Header.Clone()
	}
	resp, err := t.inner.RoundTrip(req)
	if matched && resp != nil {
		t.respHdr = resp.Header.Clone()
	}
	return resp, err
}

// TestAccUnifiedHostSendsWorkspaceIdHeader verifies, against a real unified
// (SPOG) host, that:
//  1. the SDK sends the new X-Databricks-Workspace-Id routing header on
//     workspace-scoped requests, and no longer sends the legacy
//     X-Databricks-Org-Id, and
//  2. the server still echoes the workspace ID on the legacy
//     X-Databricks-Org-Id response header — the response-side migration has
//     not yet happened, which is why CurrentWorkspaceID still reads the
//     legacy name in workspace_functions.go.
//
// Uses OAuth M2M auth (ClientID/ClientSecret) and probes the SCIM Me endpoint
// as the workspace-scoped call. A transport-level header capture is hooked in
// to verify which routing header actually crosses the wire.
//
// Runs against the workspace test environment (TEST_ENVIRONMENT_TYPE in
// {WORKSPACE, UC_WORKSPACE}). Requires UNIFIED_HOST, THIS_WORKSPACE_ID,
// TEST_ACCOUNT_ID, DATABRICKS_CLIENT_ID, and DATABRICKS_CLIENT_SECRET in the
// environment.
func TestAccUnifiedHostSendsWorkspaceIdHeader(t *testing.T) {
	loadDebugEnvIfRunsFromIDE(t, "workspace")
	if envType := testEnvironmentType(); envType != "" && envType != "WORKSPACE" && envType != "UC_WORKSPACE" {
		skipf(t)("Skipping workspace test: TEST_ENVIRONMENT_TYPE=%s", envType)
	}

	const apiPath = "/api/2.0/preview/scim/v2/Me"
	transport := &captureHeadersTransport{
		inner: http.DefaultTransport,
		path:  apiPath,
	}
	workspaceID := GetEnvOrSkipTest(t, "THIS_WORKSPACE_ID")
	cfg := &config.Config{
		Host:               GetEnvOrSkipTest(t, "UNIFIED_HOST"),
		AccountID:          GetEnvOrSkipTest(t, "TEST_ACCOUNT_ID"),
		ClientID:           GetEnvOrSkipTest(t, "DATABRICKS_CLIENT_ID"),
		ClientSecret:       GetEnvOrSkipTest(t, "DATABRICKS_CLIENT_SECRET"),
		WorkspaceID:        workspaceID,
		HTTPTransport:      transport,
		HTTPTimeoutSeconds: 60,
	}
	if err := cfg.EnsureResolved(); err != nil {
		skipf(t)("config resolve: %s", err)
	}
	t.Parallel()

	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient((*databricks.Config)(cfg)))

	// (1) workspace-scoped probe against the unified host: must succeed.
	me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
	require.NoError(t, err, "workspace-scoped SCIM Me should pass through the unified host gateway with the new routing header")
	assert.NotEmpty(t, me.UserName, "Me should return a non-empty UserName")

	// (2) On the request side, the SDK must send the new header and not the legacy one.
	require.NotNil(t, transport.reqHdr, "transport did not observe a request to %s", apiPath)
	assert.Equal(t, workspaceID, transport.reqHdr.Get("X-Databricks-Workspace-Id"),
		"SDK must send X-Databricks-Workspace-Id with Config.WorkspaceID")
	assert.Empty(t, transport.reqHdr.Get("X-Databricks-Org-Id"),
		"SDK must no longer send the legacy X-Databricks-Org-Id request header")

	// (3) On the response side, the server still echoes the legacy header name during the migration.
	require.NotNil(t, transport.respHdr, "transport did not observe a response from %s", apiPath)
	assert.NotEmpty(t, transport.respHdr.Get("X-Databricks-Org-Id"),
		"server is expected to still echo the legacy X-Databricks-Org-Id response header during the migration")
}
