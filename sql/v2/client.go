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
	cfg       *config.Config
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
		cfg:             cfg,
		apiClient:       apiClient,
		AlertsInterface: NewAlerts(databricksClient),
	}, nil
}

type AlertsLegacyClient struct {
	AlertsLegacyInterface
	cfg       *config.Config
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
		cfg:                   cfg,
		apiClient:             apiClient,
		AlertsLegacyInterface: NewAlertsLegacy(databricksClient),
	}, nil
}

type DashboardWidgetsClient struct {
	DashboardWidgetsInterface
	cfg       *config.Config
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		DashboardWidgetsInterface: NewDashboardWidgets(databricksClient),
	}, nil
}

type DashboardsClient struct {
	DashboardsInterface
	cfg       *config.Config
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
		cfg:                 cfg,
		apiClient:           apiClient,
		DashboardsInterface: NewDashboards(databricksClient),
	}, nil
}

type DataSourcesClient struct {
	DataSourcesInterface
	cfg       *config.Config
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
		cfg:                  cfg,
		apiClient:            apiClient,
		DataSourcesInterface: NewDataSources(databricksClient),
	}, nil
}

type DbsqlPermissionsClient struct {
	DbsqlPermissionsInterface
	cfg       *config.Config
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		DbsqlPermissionsInterface: NewDbsqlPermissions(databricksClient),
	}, nil
}

type QueriesClient struct {
	QueriesInterface
	cfg       *config.Config
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
		cfg:              cfg,
		apiClient:        apiClient,
		QueriesInterface: NewQueries(databricksClient),
	}, nil
}

type QueriesLegacyClient struct {
	QueriesLegacyInterface
	cfg       *config.Config
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
		cfg:                    cfg,
		apiClient:              apiClient,
		QueriesLegacyInterface: NewQueriesLegacy(databricksClient),
	}, nil
}

type QueryHistoryClient struct {
	QueryHistoryInterface
	cfg       *config.Config
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
		cfg:                   cfg,
		apiClient:             apiClient,
		QueryHistoryInterface: NewQueryHistory(databricksClient),
	}, nil
}

type QueryVisualizationsClient struct {
	QueryVisualizationsInterface
	cfg       *config.Config
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
		cfg:                          cfg,
		apiClient:                    apiClient,
		QueryVisualizationsInterface: NewQueryVisualizations(databricksClient),
	}, nil
}

type QueryVisualizationsLegacyClient struct {
	QueryVisualizationsLegacyInterface
	cfg       *config.Config
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
		cfg:                                cfg,
		apiClient:                          apiClient,
		QueryVisualizationsLegacyInterface: NewQueryVisualizationsLegacy(databricksClient),
	}, nil
}

type StatementExecutionClient struct {
	StatementExecutionInterface
	cfg       *config.Config
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
		cfg:                         cfg,
		apiClient:                   apiClient,
		StatementExecutionInterface: NewStatementExecution(databricksClient),
	}, nil
}

type WarehousesClient struct {
	WarehousesInterface
	cfg       *config.Config
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
		cfg:                 cfg,
		apiClient:           apiClient,
		WarehousesInterface: NewWarehouses(databricksClient),
	}, nil
}
