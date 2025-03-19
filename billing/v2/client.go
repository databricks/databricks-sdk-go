// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type BillableUsageClient struct {
	BillableUsageInterface
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
		BillableUsageInterface: NewBillableUsage(apiClient.ApiClient()),
	}, nil
}

type BudgetPolicyClient struct {
	BudgetPolicyInterface
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
		BudgetPolicyInterface: NewBudgetPolicy(apiClient.ApiClient()),
	}, nil
}

type BudgetsClient struct {
	BudgetsInterface
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
		BudgetsInterface: NewBudgets(apiClient.ApiClient()),
	}, nil
}

type LogDeliveryClient struct {
	LogDeliveryInterface
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
		LogDeliveryInterface: NewLogDelivery(apiClient.ApiClient()),
	}, nil
}

type UsageDashboardsClient struct {
	UsageDashboardsInterface
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
		UsageDashboardsInterface: NewUsageDashboards(apiClient.ApiClient()),
	}, nil
}
