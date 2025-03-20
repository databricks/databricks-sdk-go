// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// This API allows you to download billable usage logs for the specified account
// and date range. This feature works with all account types.
type BillableUsageClient struct {
	BillableUsageAPI
}

func NewBillableUsageClient(cfg *config.Config) (*BillableUsageClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &BillableUsageClient{
		BillableUsageAPI: BillableUsageAPI{
			billableUsageImpl: billableUsageImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// A service serves REST API about Budget policies
type BudgetPolicyClient struct {
	BudgetPolicyAPI
}

func NewBudgetPolicyClient(cfg *config.Config) (*BudgetPolicyClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &BudgetPolicyClient{
		BudgetPolicyAPI: BudgetPolicyAPI{
			budgetPolicyImpl: budgetPolicyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs manage budget configurations for this account. Budgets enable you
// to monitor usage across your account. You can set up budgets to either track
// account-wide spending, or apply filters to track the spending of specific
// teams, projects, or workspaces.
type BudgetsClient struct {
	BudgetsAPI
}

func NewBudgetsClient(cfg *config.Config) (*BudgetsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &BudgetsClient{
		BudgetsAPI: BudgetsAPI{
			budgetsImpl: budgetsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs manage log delivery configurations for this account. The two
// supported log types for this API are _billable usage logs_ and _audit logs_.
// This feature is in Public Preview. This feature works with all account ID
// types.
//
// Log delivery works with all account types. However, if your account is on the
// E2 version of the platform or on a select custom plan that allows multiple
// workspaces per account, you can optionally configure different storage
// destinations for each workspace. Log delivery status is also provided to know
// the latest status of log delivery attempts. The high-level flow of billable
// usage delivery:
//
// 1. **Create storage**: In AWS, [create a new AWS S3 bucket] with a specific
// bucket policy. Using Databricks APIs, call the Account API to create a
// [storage configuration object](:method:Storage/Create) that uses the bucket
// name. 2. **Create credentials**: In AWS, create the appropriate AWS IAM role.
// For full details, including the required IAM role policies and trust
// relationship, see [Billable usage log delivery]. Using Databricks APIs, call
// the Account API to create a [credential configuration
// object](:method:Credentials/Create) that uses the IAM role"s ARN. 3. **Create
// log delivery configuration**: Using Databricks APIs, call the Account API to
// [create a log delivery configuration](:method:LogDelivery/Create) that uses
// the credential and storage configuration objects from previous steps. You can
// specify if the logs should include all events of that log type in your
// account (_Account level_ delivery) or only events for a specific set of
// workspaces (_workspace level_ delivery). Account level log delivery applies
// to all current and future workspaces plus account level logs, while workspace
// level log delivery solely delivers logs related to the specified workspaces.
// You can create multiple types of delivery configurations per account.
//
// For billable usage delivery: * For more information about billable usage
// logs, see [Billable usage log delivery]. For the CSV schema, see the [Usage
// page]. * The delivery location is
// `<bucket-name>/<prefix>/billable-usage/csv/`, where `<prefix>` is the name of
// the optional delivery path prefix you set up during log delivery
// configuration. Files are named
// `workspaceId=<workspace-id>-usageMonth=<month>.csv`. * All billable usage
// logs apply to specific workspaces (_workspace level_ logs). You can aggregate
// usage for your entire account by creating an _account level_ delivery
// configuration that delivers logs for all current and future workspaces in
// your account. * The files are delivered daily by overwriting the month's CSV
// file for each workspace.
//
// For audit log delivery: * For more information about about audit log
// delivery, see [Audit log delivery], which includes information about the used
// JSON schema. * The delivery location is
// `<bucket-name>/<delivery-path-prefix>/workspaceId=<workspaceId>/date=<yyyy-mm-dd>/auditlogs_<internal-id>.json`.
// Files may get overwritten with the same content multiple times to achieve
// exactly-once delivery. * If the audit log delivery configuration included
// specific workspace IDs, only _workspace-level_ audit logs for those
// workspaces are delivered. If the log delivery configuration applies to the
// entire account (_account level_ delivery configuration), the audit log
// delivery includes workspace-level audit logs for all workspaces in the
// account as well as account-level audit logs. See [Audit log delivery] for
// details. * Auditable events are typically available in logs within 15
// minutes.
//
// [Audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [Billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
// [Usage page]: https://docs.databricks.com/administration-guide/account-settings/usage.html
// [create a new AWS S3 bucket]: https://docs.databricks.com/administration-guide/account-api/aws-storage.html
type LogDeliveryClient struct {
	LogDeliveryAPI
}

func NewLogDeliveryClient(cfg *config.Config) (*LogDeliveryClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &LogDeliveryClient{
		LogDeliveryAPI: LogDeliveryAPI{
			logDeliveryImpl: logDeliveryImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs manage usage dashboards for this account. Usage dashboards enable
// you to gain insights into your usage with pre-built dashboards: visualize
// breakdowns, analyze tag attributions, and identify cost drivers.
type UsageDashboardsClient struct {
	UsageDashboardsAPI
}

func NewUsageDashboardsClient(cfg *config.Config) (*UsageDashboardsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &UsageDashboardsClient{
		UsageDashboardsAPI: UsageDashboardsAPI{
			usageDashboardsImpl: usageDashboardsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
