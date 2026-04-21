package databricks

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

// CurrentWorkspaceID returns the workspace ID of the workspace that this client is
// connected to.
//
// If Config.WorkspaceID is already set (from the databrickscfg profile, the
// DATABRICKS_WORKSPACE_ID env var, host metadata, or a ?o= query param extracted
// by the caller), it is returned without an API round-trip. Otherwise the ID is
// fetched from the X-Databricks-Org-Id response header on /api/2.0/preview/scim/v2/Me.
func (w *WorkspaceClient) CurrentWorkspaceID(ctx context.Context) (int64, error) {
	if w.Config != nil && w.Config.WorkspaceID != "" {
		return strconv.ParseInt(w.Config.WorkspaceID, 10, 64)
	}
	var workspaceIdStr string
	err := w.apiClient.Do(ctx, "GET", "/api/2.0/preview/scim/v2/Me", httpclient.WithResponseHeader("X-Databricks-Org-Id", &workspaceIdStr))
	if err != nil {
		return 0, err
	}
	workspaceId, err := strconv.ParseInt(workspaceIdStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return workspaceId, nil
}
