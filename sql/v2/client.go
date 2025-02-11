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
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &AlertsClient{
		Config:          cfg,
		apiClient:       apiClient,
		AlertsInterface: NewAlerts(databricksClient),
	}, nil
}

type AlertsLegacyClient struct {
	AlertsLegacyInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &AlertsLegacyClient{
		Config:                cfg,
		apiClient:             apiClient,
		AlertsLegacyInterface: NewAlertsLegacy(databricksClient),
	}, nil
}

type DashboardWidgetsClient struct {
	DashboardWidgetsInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &DashboardWidgetsClient{
		Config:                    cfg,
		apiClient:                 apiClient,
		DashboardWidgetsInterface: NewDashboardWidgets(databricksClient),
	}, nil
}

type DashboardsClient struct {
	DashboardsInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &DashboardsClient{
		Config:              cfg,
		apiClient:           apiClient,
		DashboardsInterface: NewDashboards(databricksClient),
	}, nil
}

type DataSourcesClient struct {
	DataSourcesInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &DataSourcesClient{
		Config:               cfg,
		apiClient:            apiClient,
		DataSourcesInterface: NewDataSources(databricksClient),
	}, nil
}

type DbsqlPermissionsClient struct {
	DbsqlPermissionsInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &DbsqlPermissionsClient{
		Config:                    cfg,
		apiClient:                 apiClient,
		DbsqlPermissionsInterface: NewDbsqlPermissions(databricksClient),
	}, nil
}

type QueriesClient struct {
	QueriesInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &QueriesClient{
		Config:           cfg,
		apiClient:        apiClient,
		QueriesInterface: NewQueries(databricksClient),
	}, nil
}

type QueriesLegacyClient struct {
	QueriesLegacyInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &QueriesLegacyClient{
		Config:                 cfg,
		apiClient:              apiClient,
		QueriesLegacyInterface: NewQueriesLegacy(databricksClient),
	}, nil
}

type QueryHistoryClient struct {
	QueryHistoryInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &QueryHistoryClient{
		Config:                cfg,
		apiClient:             apiClient,
		QueryHistoryInterface: NewQueryHistory(databricksClient),
	}, nil
}

type QueryVisualizationsClient struct {
	QueryVisualizationsInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &QueryVisualizationsClient{
		Config:                       cfg,
		apiClient:                    apiClient,
		QueryVisualizationsInterface: NewQueryVisualizations(databricksClient),
	}, nil
}

type QueryVisualizationsLegacyClient struct {
	QueryVisualizationsLegacyInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &QueryVisualizationsLegacyClient{
		Config:                             cfg,
		apiClient:                          apiClient,
		QueryVisualizationsLegacyInterface: NewQueryVisualizationsLegacy(databricksClient),
	}, nil
}

type RedashConfigClient struct {
	RedashConfigInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &RedashConfigClient{
		Config:                cfg,
		apiClient:             apiClient,
		RedashConfigInterface: NewRedashConfig(databricksClient),
	}, nil
}

type StatementExecutionClient struct {
	StatementExecutionInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &StatementExecutionClient{
		Config:                      cfg,
		apiClient:                   apiClient,
		StatementExecutionInterface: NewStatementExecution(databricksClient),
	}, nil
}

type WarehousesClient struct {
	WarehousesInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &WarehousesClient{
		Config:              cfg,
		apiClient:           apiClient,
		WarehousesInterface: NewWarehouses(databricksClient),
	}, nil
}
