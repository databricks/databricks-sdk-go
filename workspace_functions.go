package databricks

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

// CurrentWorkspaceID returns the workspace ID of the workspace that this client is
// connected to.
func (w *WorkspaceClient) CurrentWorkspaceID(ctx context.Context) (int64, error) {
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
