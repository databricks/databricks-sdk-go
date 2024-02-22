// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Lakeview API methods
type lakeviewImpl struct {
	client *client.DatabricksClient
}

func (a *lakeviewImpl) Publish(ctx context.Context, request PublishRequest) error {
	var publishResponse PublishResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &publishResponse)
	return err
}
