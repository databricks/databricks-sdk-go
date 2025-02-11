// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sqlpreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AlertsLegacyPreviewClient struct {
	AlertsLegacyPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewAlertsLegacyPreviewClient(cfg *config.Config) (*AlertsLegacyPreviewClient, error) {
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

	return &AlertsLegacyPreviewClient{
		Config:                       cfg,
		apiClient:                    apiClient,
		AlertsLegacyPreviewInterface: NewAlertsLegacyPreview(databricksClient),
	}, nil
}

type AlertsPreviewClient struct {
	AlertsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewAlertsPreviewClient(cfg *config.Config) (*AlertsPreviewClient, error) {
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

	return &AlertsPreviewClient{
		Config:                 cfg,
		apiClient:              apiClient,
		AlertsPreviewInterface: NewAlertsPreview(databricksClient),
	}, nil
}

type DashboardWidgetsPreviewClient struct {
	DashboardWidgetsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewDashboardWidgetsPreviewClient(cfg *config.Config) (*DashboardWidgetsPreviewClient, error) {
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

	return &DashboardWidgetsPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		DashboardWidgetsPreviewInterface: NewDashboardWidgetsPreview(databricksClient),
	}, nil
}

type DashboardsPreviewClient struct {
	DashboardsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewDashboardsPreviewClient(cfg *config.Config) (*DashboardsPreviewClient, error) {
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

	return &DashboardsPreviewClient{
		Config:                     cfg,
		apiClient:                  apiClient,
		DashboardsPreviewInterface: NewDashboardsPreview(databricksClient),
	}, nil
}

type DataSourcesPreviewClient struct {
	DataSourcesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewDataSourcesPreviewClient(cfg *config.Config) (*DataSourcesPreviewClient, error) {
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

	return &DataSourcesPreviewClient{
		Config:                      cfg,
		apiClient:                   apiClient,
		DataSourcesPreviewInterface: NewDataSourcesPreview(databricksClient),
	}, nil
}

type DbsqlPermissionsPreviewClient struct {
	DbsqlPermissionsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewDbsqlPermissionsPreviewClient(cfg *config.Config) (*DbsqlPermissionsPreviewClient, error) {
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

	return &DbsqlPermissionsPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		DbsqlPermissionsPreviewInterface: NewDbsqlPermissionsPreview(databricksClient),
	}, nil
}

type QueriesLegacyPreviewClient struct {
	QueriesLegacyPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewQueriesLegacyPreviewClient(cfg *config.Config) (*QueriesLegacyPreviewClient, error) {
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

	return &QueriesLegacyPreviewClient{
		Config:                        cfg,
		apiClient:                     apiClient,
		QueriesLegacyPreviewInterface: NewQueriesLegacyPreview(databricksClient),
	}, nil
}

type QueriesPreviewClient struct {
	QueriesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewQueriesPreviewClient(cfg *config.Config) (*QueriesPreviewClient, error) {
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

	return &QueriesPreviewClient{
		Config:                  cfg,
		apiClient:               apiClient,
		QueriesPreviewInterface: NewQueriesPreview(databricksClient),
	}, nil
}

type QueryHistoryPreviewClient struct {
	QueryHistoryPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewQueryHistoryPreviewClient(cfg *config.Config) (*QueryHistoryPreviewClient, error) {
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

	return &QueryHistoryPreviewClient{
		Config:                       cfg,
		apiClient:                    apiClient,
		QueryHistoryPreviewInterface: NewQueryHistoryPreview(databricksClient),
	}, nil
}

type QueryVisualizationsLegacyPreviewClient struct {
	QueryVisualizationsLegacyPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewQueryVisualizationsLegacyPreviewClient(cfg *config.Config) (*QueryVisualizationsLegacyPreviewClient, error) {
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

	return &QueryVisualizationsLegacyPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		QueryVisualizationsLegacyPreviewInterface: NewQueryVisualizationsLegacyPreview(databricksClient),
	}, nil
}

type QueryVisualizationsPreviewClient struct {
	QueryVisualizationsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewQueryVisualizationsPreviewClient(cfg *config.Config) (*QueryVisualizationsPreviewClient, error) {
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

	return &QueryVisualizationsPreviewClient{
		Config:                              cfg,
		apiClient:                           apiClient,
		QueryVisualizationsPreviewInterface: NewQueryVisualizationsPreview(databricksClient),
	}, nil
}

type RedashConfigPreviewClient struct {
	RedashConfigPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewRedashConfigPreviewClient(cfg *config.Config) (*RedashConfigPreviewClient, error) {
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

	return &RedashConfigPreviewClient{
		Config:                       cfg,
		apiClient:                    apiClient,
		RedashConfigPreviewInterface: NewRedashConfigPreview(databricksClient),
	}, nil
}

type StatementExecutionPreviewClient struct {
	StatementExecutionPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewStatementExecutionPreviewClient(cfg *config.Config) (*StatementExecutionPreviewClient, error) {
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

	return &StatementExecutionPreviewClient{
		Config:                             cfg,
		apiClient:                          apiClient,
		StatementExecutionPreviewInterface: NewStatementExecutionPreview(databricksClient),
	}, nil
}

type WarehousesPreviewClient struct {
	WarehousesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewWarehousesPreviewClient(cfg *config.Config) (*WarehousesPreviewClient, error) {
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

	return &WarehousesPreviewClient{
		Config:                     cfg,
		apiClient:                  apiClient,
		WarehousesPreviewInterface: NewWarehousesPreview(databricksClient),
	}, nil
}
