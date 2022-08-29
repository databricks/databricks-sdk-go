// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// These APIs manage budget configuration including notifications for exceeding
// a budget for a period. They can also retrieve the status of each budget. 
type BudgetsService interface {
    // Create a new budget in this account.
    CreateBudget(ctx context.Context, createBudgetRequest CreateBudgetRequest) (*BudgetWithStatus, error)
    // Delete budget specified by its UUID.
    DeleteBudget(ctx context.Context, deleteBudgetRequest DeleteBudgetRequest) error
    // Get all budgets associated with this account, including non-cumulative
    // status for each day the budget is configured for.
    GetAllBudgets(ctx context.Context, getAllBudgetsRequest GetAllBudgetsRequest) (*BudgetList, error)
    // Get budget specified by its UUID, including non-cumulative status for
    // each day the budget is configured for.
    GetBudget(ctx context.Context, getBudgetRequest GetBudgetRequest) (*BudgetWithStatus, error)
    // Modify a budget in this account. Budget properties will be fully
    // overwritten.
    ModifyBudget(ctx context.Context, modifyBudgetRequest ModifyBudgetRequest) error
	GetBudgetByAccountIdAndBudgetId(ctx context.Context, accountId string, budgetId string) (*BudgetWithStatus, error)
	DeleteBudgetByAccountIdAndBudgetId(ctx context.Context, accountId string, budgetId string) error
	GetAllBudgetsByAccountId(ctx context.Context, accountId string) (*BudgetList, error)
}

func New(client *client.DatabricksClient) BudgetsService {
	return &BudgetsAPI{
		client: client,
	}
}

type BudgetsAPI struct {
	client *client.DatabricksClient
}

// Create a new budget in this account.
func (a *BudgetsAPI) CreateBudget(ctx context.Context, request CreateBudgetRequest) (*BudgetWithStatus, error) {
	var budgetWithStatus BudgetWithStatus
	path := fmt.Sprintf("/accounts/%v/budget", request.AccountId)
	err := a.client.Post(ctx, path, request, &budgetWithStatus)
	return &budgetWithStatus, err
}

// Delete budget specified by its UUID.
func (a *BudgetsAPI) DeleteBudget(ctx context.Context, request DeleteBudgetRequest) error {
	path := fmt.Sprintf("/accounts/%v/budget/%v", request.AccountId, request.BudgetId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get all budgets associated with this account, including non-cumulative status
// for each day the budget is configured for.
func (a *BudgetsAPI) GetAllBudgets(ctx context.Context, request GetAllBudgetsRequest) (*BudgetList, error) {
	var budgetList BudgetList
	path := fmt.Sprintf("/accounts/%v/budget", request.AccountId)
	err := a.client.Get(ctx, path, request, &budgetList)
	return &budgetList, err
}

// Get budget specified by its UUID, including non-cumulative status for each
// day the budget is configured for.
func (a *BudgetsAPI) GetBudget(ctx context.Context, request GetBudgetRequest) (*BudgetWithStatus, error) {
	var budgetWithStatus BudgetWithStatus
	path := fmt.Sprintf("/accounts/%v/budget/%v", request.AccountId, request.BudgetId)
	err := a.client.Get(ctx, path, request, &budgetWithStatus)
	return &budgetWithStatus, err
}

// Modify a budget in this account. Budget properties will be fully overwritten.
func (a *BudgetsAPI) ModifyBudget(ctx context.Context, request ModifyBudgetRequest) error {
	path := fmt.Sprintf("/accounts/%v/budget/%v", request.AccountId, request.BudgetId)
	err := a.client.Patch(ctx, path, request)
	return err
}


func (a *BudgetsAPI) GetBudgetByAccountIdAndBudgetId(ctx context.Context, accountId string, budgetId string) (*BudgetWithStatus, error) {
	return a.GetBudget(ctx, GetBudgetRequest{
		AccountId: accountId,
		BudgetId: budgetId,
	})
}

func (a *BudgetsAPI) DeleteBudgetByAccountIdAndBudgetId(ctx context.Context, accountId string, budgetId string) error {
	return a.DeleteBudget(ctx, DeleteBudgetRequest{
		AccountId: accountId,
		BudgetId: budgetId,
	})
}

func (a *BudgetsAPI) GetAllBudgetsByAccountId(ctx context.Context, accountId string) (*BudgetList, error) {
	return a.GetAllBudgets(ctx, GetAllBudgetsRequest{
		AccountId: accountId,
	})
}
