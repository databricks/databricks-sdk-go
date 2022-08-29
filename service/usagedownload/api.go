// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package usagedownload

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// This API allows you to download billable usage logs for the specified account
// and date range. This feature works with all account types.
type UsagedownloadService interface {
    // Return billable usage logs in CSV format for the specified account and
    // date range. See [CSV file
    // schema](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html#csv-file-schema)
    // for the data schema. Note that this method may take multiple seconds to
    // complete.
    DownloadBillableUsage(ctx context.Context, downloadBillableUsageRequest DownloadBillableUsageRequest) error
	DownloadBillableUsageByAccountId(ctx context.Context, accountId string) error
}

func New(client *client.DatabricksClient) UsagedownloadService {
	return &UsagedownloadAPI{
		client: client,
	}
}

type UsagedownloadAPI struct {
	client *client.DatabricksClient
}

// Return billable usage logs in CSV format for the specified account and date
// range. See [CSV file
// schema](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html#csv-file-schema)
// for the data schema. Note that this method may take multiple seconds to
// complete.
func (a *UsagedownloadAPI) DownloadBillableUsage(ctx context.Context, request DownloadBillableUsageRequest) error {
	path := fmt.Sprintf("/accounts/%v/usage/download", request.AccountId)
	err := a.client.Get(ctx, path, request, nil)
	return err
}


func (a *UsagedownloadAPI) DownloadBillableUsageByAccountId(ctx context.Context, accountId string) error {
	return a.DownloadBillableUsage(ctx, DownloadBillableUsageRequest{
		AccountId: accountId,
	})
}
