// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package budgets

// all definitions in this file are in alphabetical order

type BudgetAlert struct {
	// List of email addresses to be notified when budget percentage is exceeded
	// in the given period
	EmailNotifications []string `json:"email_notifications,omitempty"`
	// Percentage of the target amount used in the currect period that will
	// trigger a notification
	MinPercentage int `json:"min_percentage,omitempty"`
}

// Budget configuration to be created
type BudgetCreateRequest struct {
	Alerts []BudgetAlert `json:"alerts,omitempty"`
	// Optional end date of the budget
	EndDate string `json:"end_date,omitempty"`

	Filter string `json:"filter"`
	// Human-readable name of the budget
	Name string `json:"name"`

	Period string `json:"period"`
	// Start date of the budget period calculation
	StartDate string `json:"start_date"`
	// Target amount of the budget per period in USD
	TargetAmount string `json:"target_amount"`
}

// SQL-like filter expression with workspaceId, SKU and tag. Usage in your
// account that matches this expression will be counted in this budget.
// Supported properties on left-hand side of comparison: * `workspaceId` - the
// ID of the workspace * `sku` - SKU of the cluster, e.g.
// `STANDARD_ALL_PURPOSE_COMPUTE` * `tag.tagName`, `tag.&#39;tag name&#39;` - tag of the
// cluster Supported comparison operators: * `=` - equal * `!=` - not equal
// Supported logical operators: `AND`, `OR`. Examples: * `workspaceId=123 OR
// (sku=&#39;STANDARD_ALL_PURPOSE_COMPUTE&#39; AND tag.&#39;my tag&#39;=&#39;my value&#39;)` *
// `workspaceId!=456` * `sku=&#39;STANDARD_ALL_PURPOSE_COMPUTE&#39; OR
// sku=&#39;PREMIUM_ALL_PURPOSE_COMPUTE&#39;` * `tag.name1=&#39;value1&#39; AND
// tag.name2=&#39;value2&#39;`

// List of Budgets
type BudgetList struct {
	Budgets []BudgetWithStatus `json:"budgets,omitempty"`
}

// Period length in years, months, weeks and/or days. Examples: `1 month`, `30
// days`, `1 year, 2 months, 1 week, 2 days`

// Budget configuration with daily status
type BudgetWithStatus struct {
	Alerts []BudgetAlert `json:"alerts,omitempty"`

	BudgetId string `json:"budget_id,omitempty"`

	CreationTime string `json:"creation_time,omitempty"`
	// Optional end date of the budget
	EndDate string `json:"end_date,omitempty"`

	Filter string `json:"filter,omitempty"`
	// Human-readable name of the budget
	Name string `json:"name,omitempty"`

	Period string `json:"period,omitempty"`
	// Start date of the budget period calculation
	StartDate string `json:"start_date,omitempty"`
	// Amount used in the budget for each day (non-cumulative)
	StatusDaily []BudgetWithStatusStatusDailyItem `json:"status_daily,omitempty"`
	// Target amount of the budget per period in USD
	TargetAmount string `json:"target_amount,omitempty"`

	UpdateTime string `json:"update_time,omitempty"`
}

type BudgetWithStatusStatusDailyItem struct {
	// Amount used in this day in USD
	Amount string `json:"amount,omitempty"`

	Date string `json:"date,omitempty"`
}

type CreateBudgetRequest struct {
	// Databricks account ID. Your account must be on the E2 version of the
	// platform or on a select custom plan that allows multiple workspaces per
	// account.
	AccountId string ` path:"account_id"`

	Budget BudgetCreateRequest `json:"budget"`
}

type DeleteBudgetRequest struct {
	// Databricks account ID. Your account must be on the E2 version of the
	// platform or on a select custom plan that allows multiple workspaces per
	// account.
	AccountId string ` path:"account_id"`
	// Budget ID
	BudgetId string ` path:"budget_id"`
}

type GetAllBudgetsRequest struct {
	// Databricks account ID. Your account must be on the E2 version of the
	// platform or on a select custom plan that allows multiple workspaces per
	// account.
	AccountId string ` path:"account_id"`
}

type GetBudgetRequest struct {
	// Databricks account ID. Your account must be on the E2 version of the
	// platform or on a select custom plan that allows multiple workspaces per
	// account.
	AccountId string ` path:"account_id"`
	// Budget ID
	BudgetId string ` path:"budget_id"`
}

type ModifyBudgetRequest struct {
	// Databricks account ID. Your account must be on the E2 version of the
	// platform or on a select custom plan that allows multiple workspaces per
	// account.
	AccountId string ` path:"account_id"`

	Budget BudgetCreateRequest `json:"budget"`
	// Budget ID
	BudgetId string ` path:"budget_id"`
}
