// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AlertsClient struct {
	AlertsInterface
}

func NewAlertsClient(cfg *config.Config) (*AlertsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AlertsClient{
		AlertsInterface: NewAlerts(apiClient),
	}, nil
}

type AlertsLegacyClient struct {
	AlertsLegacyInterface
}

func NewAlertsLegacyClient(cfg *config.Config) (*AlertsLegacyClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AlertsLegacyClient{
		AlertsLegacyInterface: NewAlertsLegacy(apiClient),
	}, nil
}

type DashboardWidgetsClient struct {
	DashboardWidgetsInterface
}

func NewDashboardWidgetsClient(cfg *config.Config) (*DashboardWidgetsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DashboardWidgetsClient{
		DashboardWidgetsInterface: NewDashboardWidgets(apiClient),
	}, nil
}

type DashboardsClient struct {
	DashboardsInterface
}

func NewDashboardsClient(cfg *config.Config) (*DashboardsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DashboardsClient{
		DashboardsInterface: NewDashboards(apiClient),
	}, nil
}

type DataSourcesClient struct {
	DataSourcesInterface
}

func NewDataSourcesClient(cfg *config.Config) (*DataSourcesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DataSourcesClient{
		DataSourcesInterface: NewDataSources(apiClient),
	}, nil
}

type DbsqlPermissionsClient struct {
	DbsqlPermissionsInterface
}

func NewDbsqlPermissionsClient(cfg *config.Config) (*DbsqlPermissionsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DbsqlPermissionsClient{
		DbsqlPermissionsInterface: NewDbsqlPermissions(apiClient),
	}, nil
}

type QueriesClient struct {
	QueriesInterface
}

func NewQueriesClient(cfg *config.Config) (*QueriesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueriesClient{
		QueriesInterface: NewQueries(apiClient),
	}, nil
}

type QueriesLegacyClient struct {
	QueriesLegacyInterface
}

func NewQueriesLegacyClient(cfg *config.Config) (*QueriesLegacyClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueriesLegacyClient{
		QueriesLegacyInterface: NewQueriesLegacy(apiClient),
	}, nil
}

type QueryHistoryClient struct {
	QueryHistoryInterface
}

func NewQueryHistoryClient(cfg *config.Config) (*QueryHistoryClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryHistoryClient{
		QueryHistoryInterface: NewQueryHistory(apiClient),
	}, nil
}

type QueryVisualizationsClient struct {
	QueryVisualizationsInterface
}

func NewQueryVisualizationsClient(cfg *config.Config) (*QueryVisualizationsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryVisualizationsClient{
		QueryVisualizationsInterface: NewQueryVisualizations(apiClient),
	}, nil
}

type QueryVisualizationsLegacyClient struct {
	QueryVisualizationsLegacyInterface
}

func NewQueryVisualizationsLegacyClient(cfg *config.Config) (*QueryVisualizationsLegacyClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryVisualizationsLegacyClient{
		QueryVisualizationsLegacyInterface: NewQueryVisualizationsLegacy(apiClient),
	}, nil
}

type RedashConfigClient struct {
	RedashConfigInterface
}

func NewRedashConfigClient(cfg *config.Config) (*RedashConfigClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RedashConfigClient{
		RedashConfigInterface: NewRedashConfig(apiClient),
	}, nil
}

type StatementExecutionClient struct {
	StatementExecutionInterface
}

func NewStatementExecutionClient(cfg *config.Config) (*StatementExecutionClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &StatementExecutionClient{
		StatementExecutionInterface: NewStatementExecution(apiClient),
	}, nil
}

type WarehousesClient struct {
	WarehousesInterface
}

func NewWarehousesClient(cfg *config.Config) (*WarehousesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &WarehousesClient{
		WarehousesInterface: NewWarehouses(apiClient),
	}, nil
}
