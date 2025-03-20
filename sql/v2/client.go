// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AlertsClient struct {
	AlertsAPI
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
		AlertsAPI: AlertsAPI{
			alertsImpl: alertsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type AlertsLegacyClient struct {
	AlertsLegacyAPI
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
		AlertsLegacyAPI: AlertsLegacyAPI{
			alertsLegacyImpl: alertsLegacyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type DashboardWidgetsClient struct {
	DashboardWidgetsAPI
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
		DashboardWidgetsAPI: DashboardWidgetsAPI{
			dashboardWidgetsImpl: dashboardWidgetsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type DashboardsClient struct {
	DashboardsAPI
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
		DashboardsAPI: DashboardsAPI{
			dashboardsImpl: dashboardsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type DataSourcesClient struct {
	DataSourcesAPI
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
		DataSourcesAPI: DataSourcesAPI{
			dataSourcesImpl: dataSourcesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type DbsqlPermissionsClient struct {
	DbsqlPermissionsAPI
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
		DbsqlPermissionsAPI: DbsqlPermissionsAPI{
			dbsqlPermissionsImpl: dbsqlPermissionsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type QueriesClient struct {
	QueriesAPI
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
		QueriesAPI: QueriesAPI{
			queriesImpl: queriesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type QueriesLegacyClient struct {
	QueriesLegacyAPI
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
		QueriesLegacyAPI: QueriesLegacyAPI{
			queriesLegacyImpl: queriesLegacyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type QueryHistoryClient struct {
	QueryHistoryAPI
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
		QueryHistoryAPI: QueryHistoryAPI{
			queryHistoryImpl: queryHistoryImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type QueryVisualizationsClient struct {
	QueryVisualizationsAPI
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
		QueryVisualizationsAPI: QueryVisualizationsAPI{
			queryVisualizationsImpl: queryVisualizationsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type QueryVisualizationsLegacyClient struct {
	QueryVisualizationsLegacyAPI
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
		QueryVisualizationsLegacyAPI: QueryVisualizationsLegacyAPI{
			queryVisualizationsLegacyImpl: queryVisualizationsLegacyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type RedashConfigClient struct {
	RedashConfigAPI
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
		RedashConfigAPI: RedashConfigAPI{
			redashConfigImpl: redashConfigImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type StatementExecutionClient struct {
	StatementExecutionAPI
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
		StatementExecutionAPI: StatementExecutionAPI{
			statementExecutionImpl: statementExecutionImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type WarehousesClient struct {
	WarehousesAPI
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
		WarehousesAPI: WarehousesAPI{
			warehousesImpl: warehousesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
