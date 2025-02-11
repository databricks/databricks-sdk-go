// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billingpreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type BillableUsagePreviewClient struct {
	BillableUsagePreviewInterface

	Config *config.Config
}

func NewBillableUsagePreviewClient(cfg *config.Config) (*BillableUsagePreviewClient, error) {
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

	return &BillableUsagePreviewClient{
		Config:                        cfg,
		BillableUsagePreviewInterface: NewBillableUsagePreview(apiClient),
	}, nil
}

type BudgetPolicyPreviewClient struct {
	BudgetPolicyPreviewInterface

	Config *config.Config
}

func NewBudgetPolicyPreviewClient(cfg *config.Config) (*BudgetPolicyPreviewClient, error) {
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

	return &BudgetPolicyPreviewClient{
		Config:                       cfg,
		BudgetPolicyPreviewInterface: NewBudgetPolicyPreview(apiClient),
	}, nil
}

type BudgetsPreviewClient struct {
	BudgetsPreviewInterface

	Config *config.Config
}

func NewBudgetsPreviewClient(cfg *config.Config) (*BudgetsPreviewClient, error) {
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

	return &BudgetsPreviewClient{
		Config:                  cfg,
		BudgetsPreviewInterface: NewBudgetsPreview(apiClient),
	}, nil
}

type LogDeliveryPreviewClient struct {
	LogDeliveryPreviewInterface

	Config *config.Config
}

func NewLogDeliveryPreviewClient(cfg *config.Config) (*LogDeliveryPreviewClient, error) {
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

	return &LogDeliveryPreviewClient{
		Config:                      cfg,
		LogDeliveryPreviewInterface: NewLogDeliveryPreview(apiClient),
	}, nil
}

type UsageDashboardsPreviewClient struct {
	UsageDashboardsPreviewInterface

	Config *config.Config
}

func NewUsageDashboardsPreviewClient(cfg *config.Config) (*UsageDashboardsPreviewClient, error) {
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

	return &UsageDashboardsPreviewClient{
		Config:                          cfg,
		UsageDashboardsPreviewInterface: NewUsageDashboardsPreview(apiClient),
	}, nil
}
