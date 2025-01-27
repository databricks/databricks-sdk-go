// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AccountMetastoreAssignmentsClient struct {
	cfg *config.Config

	AccountMetastoreAssignments AccountMetastoreAssignmentsInterface
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
		cfg:                         cfg,
		AccountMetastoreAssignments: NewAccountMetastoreAssignments(apiClient),
	}, nil
}

type AccountMetastoresClient struct {
	cfg *config.Config

	AccountMetastores AccountMetastoresInterface
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
		cfg:               cfg,
		AccountMetastores: NewAccountMetastores(apiClient),
	}, nil
}

type AccountStorageCredentialsClient struct {
	cfg *config.Config

	AccountStorageCredentials AccountStorageCredentialsInterface
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
		cfg:                       cfg,
		AccountStorageCredentials: NewAccountStorageCredentials(apiClient),
	}, nil
}

type ArtifactAllowlistsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ArtifactAllowlists ArtifactAllowlistsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &ArtifactAllowlistsClient{
		cfg:                cfg,
		apiClient:          apiClient,
		ArtifactAllowlists: NewArtifactAllowlists(databricksClient),
	}, nil
}

type CatalogsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Catalogs CatalogsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &CatalogsClient{
		cfg:       cfg,
		apiClient: apiClient,
		Catalogs:  NewCatalogs(databricksClient),
	}, nil
}

type ConnectionsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Connections ConnectionsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &ConnectionsClient{
		cfg:         cfg,
		apiClient:   apiClient,
		Connections: NewConnections(databricksClient),
	}, nil
}

type CredentialsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Credentials CredentialsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &CredentialsClient{
		cfg:         cfg,
		apiClient:   apiClient,
		Credentials: NewCredentials(databricksClient),
	}, nil
}

type ExternalLocationsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ExternalLocations ExternalLocationsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &ExternalLocationsClient{
		cfg:               cfg,
		apiClient:         apiClient,
		ExternalLocations: NewExternalLocations(databricksClient),
	}, nil
}

type FunctionsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Functions FunctionsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &FunctionsClient{
		cfg:       cfg,
		apiClient: apiClient,
		Functions: NewFunctions(databricksClient),
	}, nil
}

type GrantsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Grants GrantsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &GrantsClient{
		cfg:       cfg,
		apiClient: apiClient,
		Grants:    NewGrants(databricksClient),
	}, nil
}

type MetastoresClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Metastores MetastoresInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &MetastoresClient{
		cfg:        cfg,
		apiClient:  apiClient,
		Metastores: NewMetastores(databricksClient),
	}, nil
}

type ModelVersionsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ModelVersions ModelVersionsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &ModelVersionsClient{
		cfg:           cfg,
		apiClient:     apiClient,
		ModelVersions: NewModelVersions(databricksClient),
	}, nil
}

type OnlineTablesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	OnlineTables OnlineTablesInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &OnlineTablesClient{
		cfg:          cfg,
		apiClient:    apiClient,
		OnlineTables: NewOnlineTables(databricksClient),
	}, nil
}

type QualityMonitorsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	QualityMonitors QualityMonitorsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &QualityMonitorsClient{
		cfg:             cfg,
		apiClient:       apiClient,
		QualityMonitors: NewQualityMonitors(databricksClient),
	}, nil
}

type RegisteredModelsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	RegisteredModels RegisteredModelsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &RegisteredModelsClient{
		cfg:              cfg,
		apiClient:        apiClient,
		RegisteredModels: NewRegisteredModels(databricksClient),
	}, nil
}

type ResourceQuotasClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ResourceQuotas ResourceQuotasInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &ResourceQuotasClient{
		cfg:            cfg,
		apiClient:      apiClient,
		ResourceQuotas: NewResourceQuotas(databricksClient),
	}, nil
}

type SchemasClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Schemas SchemasInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &SchemasClient{
		cfg:       cfg,
		apiClient: apiClient,
		Schemas:   NewSchemas(databricksClient),
	}, nil
}

type StorageCredentialsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	StorageCredentials StorageCredentialsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &StorageCredentialsClient{
		cfg:                cfg,
		apiClient:          apiClient,
		StorageCredentials: NewStorageCredentials(databricksClient),
	}, nil
}

type SystemSchemasClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	SystemSchemas SystemSchemasInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &SystemSchemasClient{
		cfg:           cfg,
		apiClient:     apiClient,
		SystemSchemas: NewSystemSchemas(databricksClient),
	}, nil
}

type TableConstraintsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	TableConstraints TableConstraintsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &TableConstraintsClient{
		cfg:              cfg,
		apiClient:        apiClient,
		TableConstraints: NewTableConstraints(databricksClient),
	}, nil
}

type TablesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Tables TablesInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &TablesClient{
		cfg:       cfg,
		apiClient: apiClient,
		Tables:    NewTables(databricksClient),
	}, nil
}

type TemporaryTableCredentialsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	TemporaryTableCredentials TemporaryTableCredentialsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &TemporaryTableCredentialsClient{
		cfg:                       cfg,
		apiClient:                 apiClient,
		TemporaryTableCredentials: NewTemporaryTableCredentials(databricksClient),
	}, nil
}

type VolumesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Volumes VolumesInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &VolumesClient{
		cfg:       cfg,
		apiClient: apiClient,
		Volumes:   NewVolumes(databricksClient),
	}, nil
}

type WorkspaceBindingsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	WorkspaceBindings WorkspaceBindingsInterface
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &WorkspaceBindingsClient{
		cfg:               cfg,
		apiClient:         apiClient,
		WorkspaceBindings: NewWorkspaceBindings(databricksClient),
	}, nil
}
