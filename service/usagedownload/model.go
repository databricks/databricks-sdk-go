// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package usagedownload

// all definitions in this file are in alphabetical order

type DownloadBillableUsageRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	EndMonth string ` url:"end_month,omitempty"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	PersonalData bool ` url:"personal_data,omitempty"`
	// Format: `YYYY-MM`. First month to return billable usage logs for. This
	// field is required.
	StartMonth string ` url:"start_month,omitempty"`
}

// Format specification for month in the format `YYYY-MM`. This is used to
// specify billable usage `start_month` and `end_month` properties. Note that
// billable usage logs are unavailable before March 2019 (`2019-03`).
