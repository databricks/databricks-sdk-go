// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AccountMetastoreAssignmentsClient struct {
	AccountMetastoreAssignmentsInterface

	cfg *config.Config
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
		cfg:                                  cfg,
		AccountMetastoreAssignmentsInterface: NewAccountMetastoreAssignments(apiClient),
	}, nil
}

type AccountMetastoresClient struct {
	AccountMetastoresInterface

	cfg *config.Config
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
		cfg:                        cfg,
		AccountMetastoresInterface: NewAccountMetastores(apiClient),
	}, nil
}

type AccountStorageCredentialsClient struct {
	AccountStorageCredentialsInterface

	cfg *config.Config
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
		cfg:                                cfg,
		AccountStorageCredentialsInterface: NewAccountStorageCredentials(apiClient),
	}, nil
}

type ArtifactAllowlistsClient struct {
	ArtifactAllowlistsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                         cfg,
		apiClient:                   apiClient,
		ArtifactAllowlistsInterface: NewArtifactAllowlists(databricksClient),
	}, nil
}

type CatalogsClient struct {
	CatalogsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:               cfg,
		apiClient:         apiClient,
		CatalogsInterface: NewCatalogs(databricksClient),
	}, nil
}

type ConnectionsClient struct {
	ConnectionsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                  cfg,
		apiClient:            apiClient,
		ConnectionsInterface: NewConnections(databricksClient),
	}, nil
}

type CredentialsClient struct {
	CredentialsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                  cfg,
		apiClient:            apiClient,
		CredentialsInterface: NewCredentials(databricksClient),
	}, nil
}

type ExternalLocationsClient struct {
	ExternalLocationsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                        cfg,
		apiClient:                  apiClient,
		ExternalLocationsInterface: NewExternalLocations(databricksClient),
	}, nil
}

type FunctionsClient struct {
	FunctionsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                cfg,
		apiClient:          apiClient,
		FunctionsInterface: NewFunctions(databricksClient),
	}, nil
}

type GrantsClient struct {
	GrantsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:             cfg,
		apiClient:       apiClient,
		GrantsInterface: NewGrants(databricksClient),
	}, nil
}

type MetastoresClient struct {
	MetastoresInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                 cfg,
		apiClient:           apiClient,
		MetastoresInterface: NewMetastores(databricksClient),
	}, nil
}

type ModelVersionsClient struct {
	ModelVersionsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                    cfg,
		apiClient:              apiClient,
		ModelVersionsInterface: NewModelVersions(databricksClient),
	}, nil
}

type OnlineTablesClient struct {
	OnlineTablesInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                   cfg,
		apiClient:             apiClient,
		OnlineTablesInterface: NewOnlineTables(databricksClient),
	}, nil
}

type QualityMonitorsClient struct {
	QualityMonitorsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                      cfg,
		apiClient:                apiClient,
		QualityMonitorsInterface: NewQualityMonitors(databricksClient),
	}, nil
}

type RegisteredModelsClient struct {
	RegisteredModelsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		RegisteredModelsInterface: NewRegisteredModels(databricksClient),
	}, nil
}

type ResourceQuotasClient struct {
	ResourceQuotasInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                     cfg,
		apiClient:               apiClient,
		ResourceQuotasInterface: NewResourceQuotas(databricksClient),
	}, nil
}

type SchemasClient struct {
	SchemasInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:              cfg,
		apiClient:        apiClient,
		SchemasInterface: NewSchemas(databricksClient),
	}, nil
}

type StorageCredentialsClient struct {
	StorageCredentialsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                         cfg,
		apiClient:                   apiClient,
		StorageCredentialsInterface: NewStorageCredentials(databricksClient),
	}, nil
}

type SystemSchemasClient struct {
	SystemSchemasInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                    cfg,
		apiClient:              apiClient,
		SystemSchemasInterface: NewSystemSchemas(databricksClient),
	}, nil
}

type TableConstraintsClient struct {
	TableConstraintsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		TableConstraintsInterface: NewTableConstraints(databricksClient),
	}, nil
}

type TablesClient struct {
	TablesInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:             cfg,
		apiClient:       apiClient,
		TablesInterface: NewTables(databricksClient),
	}, nil
}

type TemporaryTableCredentialsClient struct {
	TemporaryTableCredentialsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                                cfg,
		apiClient:                          apiClient,
		TemporaryTableCredentialsInterface: NewTemporaryTableCredentials(databricksClient),
	}, nil
}

type VolumesClient struct {
	VolumesInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:              cfg,
		apiClient:        apiClient,
		VolumesInterface: NewVolumes(databricksClient),
	}, nil
}

type WorkspaceBindingsClient struct {
	WorkspaceBindingsInterface
	cfg       *config.Config
	apiClient *httpclient.ApiClient
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
		cfg:                        cfg,
		apiClient:                  apiClient,
		WorkspaceBindingsInterface: NewWorkspaceBindings(databricksClient),
	}, nil
}
