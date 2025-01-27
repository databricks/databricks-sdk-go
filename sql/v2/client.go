// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AlertsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Alerts AlertsInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Alerts:    NewAlerts(databricksClient),
	}, nil
}

type AlertsLegacyClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	AlertsLegacy AlertsLegacyInterface
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
		cfg:          cfg,
		apiClient:    apiClient,
		AlertsLegacy: NewAlertsLegacy(databricksClient),
	}, nil
}

type DashboardWidgetsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	DashboardWidgets DashboardWidgetsInterface
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
		cfg:              cfg,
		apiClient:        apiClient,
		DashboardWidgets: NewDashboardWidgets(databricksClient),
	}, nil
}

type DashboardsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Dashboards DashboardsInterface
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
		cfg:        cfg,
		apiClient:  apiClient,
		Dashboards: NewDashboards(databricksClient),
	}, nil
}

type DataSourcesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	DataSources DataSourcesInterface
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
		cfg:         cfg,
		apiClient:   apiClient,
		DataSources: NewDataSources(databricksClient),
	}, nil
}

type DbsqlPermissionsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	DbsqlPermissions DbsqlPermissionsInterface
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
		cfg:              cfg,
		apiClient:        apiClient,
		DbsqlPermissions: NewDbsqlPermissions(databricksClient),
	}, nil
}

type QueriesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Queries QueriesInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Queries:   NewQueries(databricksClient),
	}, nil
}

type QueriesLegacyClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	QueriesLegacy QueriesLegacyInterface
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
		cfg:           cfg,
		apiClient:     apiClient,
		QueriesLegacy: NewQueriesLegacy(databricksClient),
	}, nil
}

type QueryHistoryClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	QueryHistory QueryHistoryInterface
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
		cfg:          cfg,
		apiClient:    apiClient,
		QueryHistory: NewQueryHistory(databricksClient),
	}, nil
}

type QueryVisualizationsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	QueryVisualizations QueryVisualizationsInterface
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
		cfg:                 cfg,
		apiClient:           apiClient,
		QueryVisualizations: NewQueryVisualizations(databricksClient),
	}, nil
}

type QueryVisualizationsLegacyClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	QueryVisualizationsLegacy QueryVisualizationsLegacyInterface
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		QueryVisualizationsLegacy: NewQueryVisualizationsLegacy(databricksClient),
	}, nil
}

type StatementExecutionClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	StatementExecution StatementExecutionInterface
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
		cfg:                cfg,
		apiClient:          apiClient,
		StatementExecution: NewStatementExecution(databricksClient),
	}, nil
}

type WarehousesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Warehouses WarehousesInterface
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
		cfg:        cfg,
		apiClient:  apiClient,
		Warehouses: NewWarehouses(databricksClient),
	}, nil
}
