package databricks

import (
	"context"
	"strconv"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

func (w *WorkspaceClient) CurrentWorkspaceId(ctx context.Context) (int64, error) {
	if w.cachedWorkspaceId != nil {
		return *w.cachedWorkspaceId, nil
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.cachedWorkspaceId != nil {
		return *w.cachedWorkspaceId, nil
	}
	var workspaceIdStr string
	err := w.apiClient.Client.Do(ctx, "GET", "/api/2.0/preview/scim/v2/Me", httpclient.WithResponseHeader("X-Databricks-Org-Id", &workspaceIdStr))
	if err != nil {
		return 0, err
	}
	workspaceId, err := strconv.ParseInt(workspaceIdStr, 10, 64)
	if err != nil {
		return 0, err
	}
	w.cachedWorkspaceId = &workspaceId
	return workspaceId, nil
}
