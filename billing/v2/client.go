// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

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
