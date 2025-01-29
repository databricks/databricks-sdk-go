// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type BillableUsageClient struct {
	BillableUsageInterface

	Config *config.Config
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
		Config:                 cfg,
		BillableUsageInterface: NewBillableUsage(apiClient),
	}, nil
}

type BudgetPolicyClient struct {
	BudgetPolicyInterface

	Config *config.Config
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
		Config:                cfg,
		BudgetPolicyInterface: NewBudgetPolicy(apiClient),
	}, nil
}

type BudgetsClient struct {
	BudgetsInterface

	Config *config.Config
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
		Config:           cfg,
		BudgetsInterface: NewBudgets(apiClient),
	}, nil
}

type LogDeliveryClient struct {
	LogDeliveryInterface

	Config *config.Config
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
		Config:               cfg,
		LogDeliveryInterface: NewLogDelivery(apiClient),
	}, nil
}

type UsageDashboardsClient struct {
	UsageDashboardsInterface

	Config *config.Config
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
		Config:                   cfg,
		UsageDashboardsInterface: NewUsageDashboards(apiClient),
	}, nil
}
