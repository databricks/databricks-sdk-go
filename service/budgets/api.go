// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewBudgets(client *client.DatabricksClient) BudgetsService {
	return &BudgetsAPI{
		client: client,
	}
}

type BudgetsAPI struct {
	client *client.DatabricksClient
}

// Create a new budget
//
// Creates a new budget in the specified account.
func (a *BudgetsAPI) CreateBudget(ctx context.Context, request CreateBudgetRequest) (*BudgetWithStatus, error) {
	var budgetWithStatus BudgetWithStatus
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget", request.AccountId)
	err := a.client.Post(ctx, path, request, &budgetWithStatus)
	return &budgetWithStatus, err
}

// Delete budget
//
// Deletes the budget specified by its UUID.
func (a *BudgetsAPI) DeleteBudget(ctx context.Context, request DeleteBudgetRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", request.AccountId, request.BudgetId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete budget
//
// Deletes the budget specified by its UUID.
func (a *BudgetsAPI) DeleteBudgetByAccountIdAndBudgetId(ctx context.Context, accountId string, budgetId string) error {
	return a.DeleteBudget(ctx, DeleteBudgetRequest{
		AccountId: accountId,
		BudgetId:  budgetId,
	})
}

// Get all budgets
//
// Gets all budgets associated with this account, including noncumulative status
// for each day that the budget is configured to include.
func (a *BudgetsAPI) GetAllBudgets(ctx context.Context, request GetAllBudgetsRequest) (*BudgetList, error) {
	var budgetList BudgetList
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget", request.AccountId)
	err := a.client.Get(ctx, path, request, &budgetList)
	return &budgetList, err
}

// Get all budgets
//
// Gets all budgets associated with this account, including noncumulative status
// for each day that the budget is configured to include.
func (a *BudgetsAPI) GetAllBudgetsByAccountId(ctx context.Context, accountId string) (*BudgetList, error) {
	return a.GetAllBudgets(ctx, GetAllBudgetsRequest{
		AccountId: accountId,
	})
}

// Get budget and its status
//
// Gets the budget specified by its UUID, including noncumulative status for
// each day that the budget is configured to include.
func (a *BudgetsAPI) GetBudget(ctx context.Context, request GetBudgetRequest) (*BudgetWithStatus, error) {
	var budgetWithStatus BudgetWithStatus
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", request.AccountId, request.BudgetId)
	err := a.client.Get(ctx, path, request, &budgetWithStatus)
	return &budgetWithStatus, err
}

// Get budget and its status
//
// Gets the budget specified by its UUID, including noncumulative status for
// each day that the budget is configured to include.
func (a *BudgetsAPI) GetBudgetByAccountIdAndBudgetId(ctx context.Context, accountId string, budgetId string) (*BudgetWithStatus, error) {
	return a.GetBudget(ctx, GetBudgetRequest{
		AccountId: accountId,
		BudgetId:  budgetId,
	})
}

// Modify budget
//
// Modifies a budget in this account. Budget properties are completely
// overwritten.
func (a *BudgetsAPI) ModifyBudget(ctx context.Context, request ModifyBudgetRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", request.AccountId, request.BudgetId)
	err := a.client.Patch(ctx, path, request)
	return err
}
