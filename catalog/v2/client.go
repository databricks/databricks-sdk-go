// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AccountMetastoreAssignmentsClient struct {
	AccountMetastoreAssignmentsInterface
}

func NewAccountMetastoreAssignmentsClient(cfg *config.Config) (*AccountMetastoreAssignmentsClient, error) {
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

	return &AccountMetastoreAssignmentsClient{
		AccountMetastoreAssignmentsInterface: NewAccountMetastoreAssignments(apiClient.ApiClient()),
	}, nil
}

type AccountMetastoresClient struct {
	AccountMetastoresInterface
}

func NewAccountMetastoresClient(cfg *config.Config) (*AccountMetastoresClient, error) {
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

	return &AccountMetastoresClient{
		AccountMetastoresInterface: NewAccountMetastores(apiClient.ApiClient()),
	}, nil
}

type AccountStorageCredentialsClient struct {
	AccountStorageCredentialsInterface
}

func NewAccountStorageCredentialsClient(cfg *config.Config) (*AccountStorageCredentialsClient, error) {
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

	return &AccountStorageCredentialsClient{
		AccountStorageCredentialsInterface: NewAccountStorageCredentials(apiClient.ApiClient()),
	}, nil
}

type ArtifactAllowlistsClient struct {
	ArtifactAllowlistsInterface
}

func NewArtifactAllowlistsClient(cfg *config.Config) (*ArtifactAllowlistsClient, error) {
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

	return &ArtifactAllowlistsClient{
		ArtifactAllowlistsInterface: NewArtifactAllowlists(apiClient.ApiClient()),
	}, nil
}

type CatalogsClient struct {
	CatalogsInterface
}

func NewCatalogsClient(cfg *config.Config) (*CatalogsClient, error) {
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

	return &CatalogsClient{
		CatalogsInterface: NewCatalogs(apiClient.ApiClient()),
	}, nil
}

type ConnectionsClient struct {
	ConnectionsInterface
}

func NewConnectionsClient(cfg *config.Config) (*ConnectionsClient, error) {
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

	return &ConnectionsClient{
		ConnectionsInterface: NewConnections(apiClient.ApiClient()),
	}, nil
}

type CredentialsClient struct {
	CredentialsInterface
}

func NewCredentialsClient(cfg *config.Config) (*CredentialsClient, error) {
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

	return &CredentialsClient{
		CredentialsInterface: NewCredentials(apiClient.ApiClient()),
	}, nil
}

type ExternalLocationsClient struct {
	ExternalLocationsInterface
}

func NewExternalLocationsClient(cfg *config.Config) (*ExternalLocationsClient, error) {
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

	return &ExternalLocationsClient{
		ExternalLocationsInterface: NewExternalLocations(apiClient.ApiClient()),
	}, nil
}

type FunctionsClient struct {
	FunctionsInterface
}

func NewFunctionsClient(cfg *config.Config) (*FunctionsClient, error) {
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

	return &FunctionsClient{
		FunctionsInterface: NewFunctions(apiClient.ApiClient()),
	}, nil
}

type GrantsClient struct {
	GrantsInterface
}

func NewGrantsClient(cfg *config.Config) (*GrantsClient, error) {
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

	return &GrantsClient{
		GrantsInterface: NewGrants(apiClient.ApiClient()),
	}, nil
}

type MetastoresClient struct {
	MetastoresInterface
}

func NewMetastoresClient(cfg *config.Config) (*MetastoresClient, error) {
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

	return &MetastoresClient{
		MetastoresInterface: NewMetastores(apiClient.ApiClient()),
	}, nil
}

type ModelVersionsClient struct {
	ModelVersionsInterface
}

func NewModelVersionsClient(cfg *config.Config) (*ModelVersionsClient, error) {
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

	return &ModelVersionsClient{
		ModelVersionsInterface: NewModelVersions(apiClient.ApiClient()),
	}, nil
}

type OnlineTablesClient struct {
	OnlineTablesInterface
}

func NewOnlineTablesClient(cfg *config.Config) (*OnlineTablesClient, error) {
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

	return &OnlineTablesClient{
		OnlineTablesInterface: NewOnlineTables(apiClient.ApiClient()),
	}, nil
}

type QualityMonitorsClient struct {
	QualityMonitorsInterface
}

func NewQualityMonitorsClient(cfg *config.Config) (*QualityMonitorsClient, error) {
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

	return &QualityMonitorsClient{
		QualityMonitorsInterface: NewQualityMonitors(apiClient.ApiClient()),
	}, nil
}

type RegisteredModelsClient struct {
	RegisteredModelsInterface
}

func NewRegisteredModelsClient(cfg *config.Config) (*RegisteredModelsClient, error) {
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

	return &RegisteredModelsClient{
		RegisteredModelsInterface: NewRegisteredModels(apiClient.ApiClient()),
	}, nil
}

type ResourceQuotasClient struct {
	ResourceQuotasInterface
}

func NewResourceQuotasClient(cfg *config.Config) (*ResourceQuotasClient, error) {
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

	return &ResourceQuotasClient{
		ResourceQuotasInterface: NewResourceQuotas(apiClient.ApiClient()),
	}, nil
}

type SchemasClient struct {
	SchemasInterface
}

func NewSchemasClient(cfg *config.Config) (*SchemasClient, error) {
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

	return &SchemasClient{
		SchemasInterface: NewSchemas(apiClient.ApiClient()),
	}, nil
}

type StorageCredentialsClient struct {
	StorageCredentialsInterface
}

func NewStorageCredentialsClient(cfg *config.Config) (*StorageCredentialsClient, error) {
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

	return &StorageCredentialsClient{
		StorageCredentialsInterface: NewStorageCredentials(apiClient.ApiClient()),
	}, nil
}

type SystemSchemasClient struct {
	SystemSchemasInterface
}

func NewSystemSchemasClient(cfg *config.Config) (*SystemSchemasClient, error) {
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

	return &SystemSchemasClient{
		SystemSchemasInterface: NewSystemSchemas(apiClient.ApiClient()),
	}, nil
}

type TableConstraintsClient struct {
	TableConstraintsInterface
}

func NewTableConstraintsClient(cfg *config.Config) (*TableConstraintsClient, error) {
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

	return &TableConstraintsClient{
		TableConstraintsInterface: NewTableConstraints(apiClient.ApiClient()),
	}, nil
}

type TablesClient struct {
	TablesInterface
}

func NewTablesClient(cfg *config.Config) (*TablesClient, error) {
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

	return &TablesClient{
		TablesInterface: NewTables(apiClient.ApiClient()),
	}, nil
}

type TemporaryTableCredentialsClient struct {
	TemporaryTableCredentialsInterface
}

func NewTemporaryTableCredentialsClient(cfg *config.Config) (*TemporaryTableCredentialsClient, error) {
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

	return &TemporaryTableCredentialsClient{
		TemporaryTableCredentialsInterface: NewTemporaryTableCredentials(apiClient.ApiClient()),
	}, nil
}

type VolumesClient struct {
	VolumesInterface
}

func NewVolumesClient(cfg *config.Config) (*VolumesClient, error) {
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

	return &VolumesClient{
		VolumesInterface: NewVolumes(apiClient.ApiClient()),
	}, nil
}

type WorkspaceBindingsClient struct {
	WorkspaceBindingsInterface
}

func NewWorkspaceBindingsClient(cfg *config.Config) (*WorkspaceBindingsClient, error) {
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

	return &WorkspaceBindingsClient{
		WorkspaceBindingsInterface: NewWorkspaceBindings(apiClient.ApiClient()),
	}, nil
}
