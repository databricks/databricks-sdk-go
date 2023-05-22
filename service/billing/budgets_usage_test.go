// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package billing_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/billing"
)

func ExampleBudgetsAPI_Create_budgets() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.Budgets.Create(ctx, billing.WrappedBudget{
		Budget: billing.Budget{
			Name:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Filter:       "tag.tagName = 'all'",
			Period:       "1 month",
			StartDate:    "2022-01-01",
			TargetAmount: "100",
			Alerts: []billing.BudgetAlert{billing.BudgetAlert{
				EmailNotifications: []string{"admin@example.com"},
				MinPercentage:      50,
			}},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = a.Budgets.DeleteByBudgetId(ctx, created.Budget.BudgetId)
	if err != nil {
		panic(err)
	}

}

func ExampleBudgetsAPI_Get_budgets() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.Budgets.Create(ctx, billing.WrappedBudget{
		Budget: billing.Budget{
			Name:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Filter:       "tag.tagName = 'all'",
			Period:       "1 month",
			StartDate:    "2022-01-01",
			TargetAmount: "100",
			Alerts: []billing.BudgetAlert{billing.BudgetAlert{
				EmailNotifications: []string{"admin@example.com"},
				MinPercentage:      50,
			}},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := a.Budgets.GetByBudgetId(ctx, created.Budget.BudgetId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = a.Budgets.DeleteByBudgetId(ctx, created.Budget.BudgetId)
	if err != nil {
		panic(err)
	}

}

func ExampleBudgetsAPI_ListAll_budgets() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	all, err := a.Budgets.ListAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleBudgetsAPI_Update_budgets() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.Budgets.Create(ctx, billing.WrappedBudget{
		Budget: billing.Budget{
			Name:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Filter:       "tag.tagName = 'all'",
			Period:       "1 month",
			StartDate:    "2022-01-01",
			TargetAmount: "100",
			Alerts: []billing.BudgetAlert{billing.BudgetAlert{
				EmailNotifications: []string{"admin@example.com"},
				MinPercentage:      50,
			}},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = a.Budgets.Update(ctx, billing.WrappedBudget{
		BudgetId: created.Budget.BudgetId,
		Budget: billing.Budget{
			Name:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Filter:       "tag.tagName = 'all'",
			Period:       "1 month",
			StartDate:    "2022-01-01",
			TargetAmount: "100",
			Alerts: []billing.BudgetAlert{billing.BudgetAlert{
				EmailNotifications: []string{"admin@example.com"},
				MinPercentage:      70,
			}},
		},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = a.Budgets.DeleteByBudgetId(ctx, created.Budget.BudgetId)
	if err != nil {
		panic(err)
	}

}
