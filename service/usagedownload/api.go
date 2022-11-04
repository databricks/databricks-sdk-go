// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package usagedownload

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewBillableUsageDownload(client *client.DatabricksClient) BillableUsageDownloadService {
	return &BillableUsageDownloadAPI{
		client: client,
	}
}

type BillableUsageDownloadAPI struct {
	client *client.DatabricksClient
}

// Return billable usage logs
//
// Returns billable usage logs in CSV format for the specified account and date
// range. For the data schema, see [CSV file
// schema](https://docs.databricks.com/administration-guide/account-settings/usage-analysis.html#schema).
// Note that this method might take multiple seconds to complete.
func (a *BillableUsageDownloadAPI) DownloadBillableUsage(ctx context.Context, request DownloadBillableUsageRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/usage/download", request.AccountId)
	err := a.client.Get(ctx, path, request, nil)
	return err
}

// Return billable usage logs
//
// Returns billable usage logs in CSV format for the specified account and date
// range. For the data schema, see [CSV file
// schema](https://docs.databricks.com/administration-guide/account-settings/usage-analysis.html#schema).
// Note that this method might take multiple seconds to complete.
func (a *BillableUsageDownloadAPI) DownloadBillableUsageByAccountId(ctx context.Context, accountId string) error {
	return a.DownloadBillableUsage(ctx, DownloadBillableUsageRequest{
		AccountId: accountId,
	})
}
