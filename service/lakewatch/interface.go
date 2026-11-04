// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package lakewatch

import (
	"context"
)

// Manage Lakewatch: security-analytics data sources, detection rules, and
// per-catalog configuration.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LakewatchService interface {

	// Creates a new datasource.
	CreateDatasource(ctx context.Context, request CreateDatasourceRequest) (*Datasource, error)

	// Creates a new rule.
	CreateRule(ctx context.Context, request CreateRuleRequest) (*Rule, error)

	// Deletes a datasource by its resource name.
	DeleteDatasource(ctx context.Context, request DeleteDatasourceRequest) error

	// Deletes a rule by its resource name.
	DeleteRule(ctx context.Context, request DeleteRuleRequest) error

	// Returns the alert delivery configuration for a catalog.
	GetAlertConfig(ctx context.Context, request GetAlertConfigRequest) (*AlertConfig, error)

	// Returns the common Lakewatch configuration for a catalog.
	GetCommonConfig(ctx context.Context, request GetCommonConfigRequest) (*CommonConfig, error)

	// Returns a single datasource by its resource name.
	GetDatasource(ctx context.Context, request GetDatasourceRequest) (*Datasource, error)

	// Returns the datasource ingestion configuration for a catalog.
	GetDatasourceConfig(ctx context.Context, request GetDatasourceConfigRequest) (*DatasourceConfig, error)

	// Returns a single rule by its resource name.
	GetRule(ctx context.Context, request GetRuleRequest) (*Rule, error)

	// Returns the detection-rule configuration for a catalog.
	GetRuleConfig(ctx context.Context, request GetRuleConfigRequest) (*RuleConfig, error)

	// Lists the datasources in a Unity Catalog catalog, optionally narrowed to
	// one schema or filtered by name. Results are paginated. Datasources that
	// are still being created or deleted are not listed.
	ListDatasources(ctx context.Context, request ListDatasourcesRequest) (*ListDatasourcesResponse, error)

	// Lists rules, with optional filtering and pagination.
	ListRules(ctx context.Context, request ListRulesRequest) (*ListRulesResponse, error)

	// Updates the alert delivery configuration for a catalog.
	UpdateAlertConfig(ctx context.Context, request UpdateAlertConfigRequest) (*AlertConfig, error)

	// Updates the common Lakewatch configuration for a catalog.
	UpdateCommonConfig(ctx context.Context, request UpdateCommonConfigRequest) (*CommonConfig, error)

	// Updates an existing datasource.
	UpdateDatasource(ctx context.Context, request UpdateDatasourceRequest) (*Datasource, error)

	// Updates the datasource ingestion configuration for a catalog.
	UpdateDatasourceConfig(ctx context.Context, request UpdateDatasourceConfigRequest) (*DatasourceConfig, error)

	// Updates an existing rule.
	UpdateRule(ctx context.Context, request UpdateRuleRequest) (*Rule, error)

	// Updates the detection-rule configuration for a catalog.
	UpdateRuleConfig(ctx context.Context, request UpdateRuleConfigRequest) (*RuleConfig, error)
}
