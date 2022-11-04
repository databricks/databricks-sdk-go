// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package budgets

// all definitions in this file are in alphabetical order

type BudgetAlert struct {
	// List of email addresses to be notified when budget percentage is exceeded
	// in the given period.
	EmailNotifications []string `json:"email_notifications,omitempty"`
	// Percentage of the target amount used in the currect period that will
	// trigger a notification.
	MinPercentage int `json:"min_percentage,omitempty"`
}

// Budget configuration to be created.
type BudgetCreateRequest struct {
	Alerts []BudgetAlert `json:"alerts,omitempty"`
	// Optional end date of the budget.
	EndDate string `json:"end_date,omitempty"`
	// SQL-like filter expression with workspaceId, SKU and tag. Usage in your
	// account that matches this expression will be counted in this budget.
	//
	// Supported properties on left-hand side of comparison: * `workspaceId` -
	// the ID of the workspace * `sku` - SKU of the cluster, e.g.
	// `STANDARD_ALL_PURPOSE_COMPUTE` * `tag.tagName`, `tag.'tag name'` - tag of
	// the cluster
	//
	// Supported comparison operators: * `=` - equal * `!=` - not equal
	//
	// Supported logical operators: `AND`, `OR`.
	//
	// Examples: * `workspaceId=123 OR (sku='STANDARD_ALL_PURPOSE_COMPUTE' AND
	// tag.'my tag'='my value')` * `workspaceId!=456` *
	// `sku='STANDARD_ALL_PURPOSE_COMPUTE' OR sku='PREMIUM_ALL_PURPOSE_COMPUTE'`
	// * `tag.name1='value1' AND tag.name2='value2'`
	Filter string `json:"filter"`
	// Human-readable name of the budget.
	Name string `json:"name"`
	// Period length in years, months, weeks and/or days. Examples: `1 month`,
	// `30 days`, `1 year, 2 months, 1 week, 2 days`
	Period string `json:"period"`
	// Start date of the budget period calculation.
	StartDate string `json:"start_date"`
	// Target amount of the budget per period in USD.
	TargetAmount string `json:"target_amount"`
}

// SQL-like filter expression with workspaceId, SKU and tag. Usage in your
// account that matches this expression will be counted in this budget.
//
// Supported properties on left-hand side of comparison: * `workspaceId` - the
// ID of the workspace * `sku` - SKU of the cluster, e.g.
// `STANDARD_ALL_PURPOSE_COMPUTE` * `tag.tagName`, `tag.'tag name'` - tag of the
// cluster
//
// Supported comparison operators: * `=` - equal * `!=` - not equal
//
// Supported logical operators: `AND`, `OR`.
//
// Examples: * `workspaceId=123 OR (sku='STANDARD_ALL_PURPOSE_COMPUTE' AND
// tag.'my tag'='my value')` * `workspaceId!=456` *
// `sku='STANDARD_ALL_PURPOSE_COMPUTE' OR sku='PREMIUM_ALL_PURPOSE_COMPUTE'` *
// `tag.name1='value1' AND tag.name2='value2'`

// List of budgets.
type BudgetList struct {
	Budgets []BudgetWithStatus `json:"budgets,omitempty"`
}

// Period length in years, months, weeks and/or days. Examples: `1 month`, `30
// days`, `1 year, 2 months, 1 week, 2 days`

// Budget configuration with daily status.
type BudgetWithStatus struct {
	Alerts []BudgetAlert `json:"alerts,omitempty"`

	BudgetId string `json:"budget_id,omitempty"`

	CreationTime string `json:"creation_time,omitempty"`
	// Optional end date of the budget.
	EndDate string `json:"end_date,omitempty"`
	// SQL-like filter expression with workspaceId, SKU and tag. Usage in your
	// account that matches this expression will be counted in this budget.
	//
	// Supported properties on left-hand side of comparison: * `workspaceId` -
	// the ID of the workspace * `sku` - SKU of the cluster, e.g.
	// `STANDARD_ALL_PURPOSE_COMPUTE` * `tag.tagName`, `tag.'tag name'` - tag of
	// the cluster
	//
	// Supported comparison operators: * `=` - equal * `!=` - not equal
	//
	// Supported logical operators: `AND`, `OR`.
	//
	// Examples: * `workspaceId=123 OR (sku='STANDARD_ALL_PURPOSE_COMPUTE' AND
	// tag.'my tag'='my value')` * `workspaceId!=456` *
	// `sku='STANDARD_ALL_PURPOSE_COMPUTE' OR sku='PREMIUM_ALL_PURPOSE_COMPUTE'`
	// * `tag.name1='value1' AND tag.name2='value2'`
	Filter string `json:"filter,omitempty"`
	// Human-readable name of the budget.
	Name string `json:"name,omitempty"`
	// Period length in years, months, weeks and/or days. Examples: `1 month`,
	// `30 days`, `1 year, 2 months, 1 week, 2 days`
	Period string `json:"period,omitempty"`
	// Start date of the budget period calculation.
	StartDate string `json:"start_date,omitempty"`
	// Amount used in the budget for each day (noncumulative).
	StatusDaily []BudgetWithStatusStatusDailyItem `json:"status_daily,omitempty"`
	// Target amount of the budget per period in USD.
	TargetAmount string `json:"target_amount,omitempty"`

	UpdateTime string `json:"update_time,omitempty"`
}

type BudgetWithStatusStatusDailyItem struct {
	// Amount used in this day in USD.
	Amount string `json:"amount,omitempty"`

	Date string `json:"date,omitempty"`
}

type CreateBudgetRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Budget configuration to be created.
	Budget BudgetCreateRequest `json:"budget"`
}

type DeleteBudgetRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Budget ID
	BudgetId string `json:"-" path:"budget_id"`
}

type GetAllBudgetsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetBudgetRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Budget ID
	BudgetId string `json:"-" path:"budget_id"`
}

type ModifyBudgetRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Budget configuration to be created.
	Budget BudgetCreateRequest `json:"budget"`
	// Budget ID
	BudgetId string `json:"-" path:"budget_id"`
}
