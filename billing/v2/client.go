// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type BillableUsageClient struct {
	cfg *config.Config

	BillableUsage BillableUsageInterface
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
		cfg:           cfg,
		BillableUsage: NewBillableUsage(apiClient),
	}, nil
}

type BudgetsClient struct {
	cfg *config.Config

	Budgets BudgetsInterface
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
		cfg:     cfg,
		Budgets: NewBudgets(apiClient),
	}, nil
}

type LogDeliveryClient struct {
	cfg *config.Config

	LogDelivery LogDeliveryInterface
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
		cfg:         cfg,
		LogDelivery: NewLogDelivery(apiClient),
	}, nil
}

type UsageDashboardsClient struct {
	cfg *config.Config

	UsageDashboards UsageDashboardsInterface
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
		cfg:             cfg,
		UsageDashboards: NewUsageDashboards(apiClient),
	}, nil
}
