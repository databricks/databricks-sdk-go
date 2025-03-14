// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AlertsClient struct {
	AlertsInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AlertsClient{
		apiClient:       databricksClient.ApiClient(),
		AlertsInterface: NewAlerts(databricksClient),
	}, nil
}

type AlertsLegacyClient struct {
	AlertsLegacyInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AlertsLegacyClient{
		apiClient:             databricksClient.ApiClient(),
		AlertsLegacyInterface: NewAlertsLegacy(databricksClient),
	}, nil
}

type DashboardWidgetsClient struct {
	DashboardWidgetsInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DashboardWidgetsClient{
		apiClient:                 databricksClient.ApiClient(),
		DashboardWidgetsInterface: NewDashboardWidgets(databricksClient),
	}, nil
}

type DashboardsClient struct {
	DashboardsInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DashboardsClient{
		apiClient:           databricksClient.ApiClient(),
		DashboardsInterface: NewDashboards(databricksClient),
	}, nil
}

type DataSourcesClient struct {
	DataSourcesInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DataSourcesClient{
		apiClient:            databricksClient.ApiClient(),
		DataSourcesInterface: NewDataSources(databricksClient),
	}, nil
}

type DbsqlPermissionsClient struct {
	DbsqlPermissionsInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DbsqlPermissionsClient{
		apiClient:                 databricksClient.ApiClient(),
		DbsqlPermissionsInterface: NewDbsqlPermissions(databricksClient),
	}, nil
}

type QueriesClient struct {
	QueriesInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueriesClient{
		apiClient:        databricksClient.ApiClient(),
		QueriesInterface: NewQueries(databricksClient),
	}, nil
}

type QueriesLegacyClient struct {
	QueriesLegacyInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueriesLegacyClient{
		apiClient:              databricksClient.ApiClient(),
		QueriesLegacyInterface: NewQueriesLegacy(databricksClient),
	}, nil
}

type QueryHistoryClient struct {
	QueryHistoryInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryHistoryClient{
		apiClient:             databricksClient.ApiClient(),
		QueryHistoryInterface: NewQueryHistory(databricksClient),
	}, nil
}

type QueryVisualizationsClient struct {
	QueryVisualizationsInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryVisualizationsClient{
		apiClient:                    databricksClient.ApiClient(),
		QueryVisualizationsInterface: NewQueryVisualizations(databricksClient),
	}, nil
}

type QueryVisualizationsLegacyClient struct {
	QueryVisualizationsLegacyInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryVisualizationsLegacyClient{
		apiClient:                          databricksClient.ApiClient(),
		QueryVisualizationsLegacyInterface: NewQueryVisualizationsLegacy(databricksClient),
	}, nil
}

type RedashConfigClient struct {
	RedashConfigInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RedashConfigClient{
		apiClient:             databricksClient.ApiClient(),
		RedashConfigInterface: NewRedashConfig(databricksClient),
	}, nil
}

type StatementExecutionClient struct {
	StatementExecutionInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &StatementExecutionClient{
		apiClient:                   databricksClient.ApiClient(),
		StatementExecutionInterface: NewStatementExecution(databricksClient),
	}, nil
}

type WarehousesClient struct {
	WarehousesInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &WarehousesClient{
		apiClient:           databricksClient.ApiClient(),
		WarehousesInterface: NewWarehouses(databricksClient),
	}, nil
}
