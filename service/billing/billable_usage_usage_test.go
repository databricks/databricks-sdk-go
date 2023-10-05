// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package billing_test

import (
	"context"

	"github.com/databricks/databricks-sdk-go"

	"github.com/databricks/databricks-sdk-go/service/billing"
)

func ExampleBillableUsageAPI_Download_usageDownload() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	_, err = a.BillableUsage.Download(ctx, billing.DownloadRequest{
		StartMonth: "2022-01",
		EndMonth:   "2022-02",
	})
	if err != nil {
		panic(err)
	}

}
