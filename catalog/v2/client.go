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
		AccountMetastoreAssignmentsInterface: NewAccountMetastoreAssignments(apiClient),
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
		AccountMetastoresInterface: NewAccountMetastores(apiClient),
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
		AccountStorageCredentialsInterface: NewAccountStorageCredentials(apiClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ArtifactAllowlistsClient{
		ArtifactAllowlistsInterface: NewArtifactAllowlists(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CatalogsClient{
		CatalogsInterface: NewCatalogs(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ConnectionsClient{
		ConnectionsInterface: NewConnections(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CredentialsClient{
		CredentialsInterface: NewCredentials(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ExternalLocationsClient{
		ExternalLocationsInterface: NewExternalLocations(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &FunctionsClient{
		FunctionsInterface: NewFunctions(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GrantsClient{
		GrantsInterface: NewGrants(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &MetastoresClient{
		MetastoresInterface: NewMetastores(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ModelVersionsClient{
		ModelVersionsInterface: NewModelVersions(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &OnlineTablesClient{
		OnlineTablesInterface: NewOnlineTables(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QualityMonitorsClient{
		QualityMonitorsInterface: NewQualityMonitors(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RegisteredModelsClient{
		RegisteredModelsInterface: NewRegisteredModels(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ResourceQuotasClient{
		ResourceQuotasInterface: NewResourceQuotas(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &SchemasClient{
		SchemasInterface: NewSchemas(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &StorageCredentialsClient{
		StorageCredentialsInterface: NewStorageCredentials(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &SystemSchemasClient{
		SystemSchemasInterface: NewSystemSchemas(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &TableConstraintsClient{
		TableConstraintsInterface: NewTableConstraints(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &TablesClient{
		TablesInterface: NewTables(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &TemporaryTableCredentialsClient{
		TemporaryTableCredentialsInterface: NewTemporaryTableCredentials(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &VolumesClient{
		VolumesInterface: NewVolumes(databricksClient),
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &WorkspaceBindingsClient{
		WorkspaceBindingsInterface: NewWorkspaceBindings(databricksClient),
	}, nil
}
