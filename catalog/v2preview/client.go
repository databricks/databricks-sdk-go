// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalogpreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AccountMetastoreAssignmentsPreviewClient struct {
	AccountMetastoreAssignmentsPreviewInterface

	Config *config.Config
}

func NewAccountMetastoreAssignmentsPreviewClient(cfg *config.Config) (*AccountMetastoreAssignmentsPreviewClient, error) {
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

	return &AccountMetastoreAssignmentsPreviewClient{
		Config: cfg,
		AccountMetastoreAssignmentsPreviewInterface: NewAccountMetastoreAssignmentsPreview(apiClient),
	}, nil
}

type AccountMetastoresPreviewClient struct {
	AccountMetastoresPreviewInterface

	Config *config.Config
}

func NewAccountMetastoresPreviewClient(cfg *config.Config) (*AccountMetastoresPreviewClient, error) {
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

	return &AccountMetastoresPreviewClient{
		Config:                            cfg,
		AccountMetastoresPreviewInterface: NewAccountMetastoresPreview(apiClient),
	}, nil
}

type AccountStorageCredentialsPreviewClient struct {
	AccountStorageCredentialsPreviewInterface

	Config *config.Config
}

func NewAccountStorageCredentialsPreviewClient(cfg *config.Config) (*AccountStorageCredentialsPreviewClient, error) {
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

	return &AccountStorageCredentialsPreviewClient{
		Config: cfg,
		AccountStorageCredentialsPreviewInterface: NewAccountStorageCredentialsPreview(apiClient),
	}, nil
}

type ArtifactAllowlistsPreviewClient struct {
	ArtifactAllowlistsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewArtifactAllowlistsPreviewClient(cfg *config.Config) (*ArtifactAllowlistsPreviewClient, error) {
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

	return &ArtifactAllowlistsPreviewClient{
		Config:                             cfg,
		apiClient:                          apiClient,
		ArtifactAllowlistsPreviewInterface: NewArtifactAllowlistsPreview(databricksClient),
	}, nil
}

type CatalogsPreviewClient struct {
	CatalogsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewCatalogsPreviewClient(cfg *config.Config) (*CatalogsPreviewClient, error) {
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

	return &CatalogsPreviewClient{
		Config:                   cfg,
		apiClient:                apiClient,
		CatalogsPreviewInterface: NewCatalogsPreview(databricksClient),
	}, nil
}

type ConnectionsPreviewClient struct {
	ConnectionsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewConnectionsPreviewClient(cfg *config.Config) (*ConnectionsPreviewClient, error) {
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

	return &ConnectionsPreviewClient{
		Config:                      cfg,
		apiClient:                   apiClient,
		ConnectionsPreviewInterface: NewConnectionsPreview(databricksClient),
	}, nil
}

type CredentialsPreviewClient struct {
	CredentialsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewCredentialsPreviewClient(cfg *config.Config) (*CredentialsPreviewClient, error) {
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

	return &CredentialsPreviewClient{
		Config:                      cfg,
		apiClient:                   apiClient,
		CredentialsPreviewInterface: NewCredentialsPreview(databricksClient),
	}, nil
}

type ExternalLocationsPreviewClient struct {
	ExternalLocationsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewExternalLocationsPreviewClient(cfg *config.Config) (*ExternalLocationsPreviewClient, error) {
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

	return &ExternalLocationsPreviewClient{
		Config:                            cfg,
		apiClient:                         apiClient,
		ExternalLocationsPreviewInterface: NewExternalLocationsPreview(databricksClient),
	}, nil
}

type FunctionsPreviewClient struct {
	FunctionsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewFunctionsPreviewClient(cfg *config.Config) (*FunctionsPreviewClient, error) {
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

	return &FunctionsPreviewClient{
		Config:                    cfg,
		apiClient:                 apiClient,
		FunctionsPreviewInterface: NewFunctionsPreview(databricksClient),
	}, nil
}

type GrantsPreviewClient struct {
	GrantsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewGrantsPreviewClient(cfg *config.Config) (*GrantsPreviewClient, error) {
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

	return &GrantsPreviewClient{
		Config:                 cfg,
		apiClient:              apiClient,
		GrantsPreviewInterface: NewGrantsPreview(databricksClient),
	}, nil
}

type MetastoresPreviewClient struct {
	MetastoresPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewMetastoresPreviewClient(cfg *config.Config) (*MetastoresPreviewClient, error) {
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

	return &MetastoresPreviewClient{
		Config:                     cfg,
		apiClient:                  apiClient,
		MetastoresPreviewInterface: NewMetastoresPreview(databricksClient),
	}, nil
}

type ModelVersionsPreviewClient struct {
	ModelVersionsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewModelVersionsPreviewClient(cfg *config.Config) (*ModelVersionsPreviewClient, error) {
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

	return &ModelVersionsPreviewClient{
		Config:                        cfg,
		apiClient:                     apiClient,
		ModelVersionsPreviewInterface: NewModelVersionsPreview(databricksClient),
	}, nil
}

type OnlineTablesPreviewClient struct {
	OnlineTablesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewOnlineTablesPreviewClient(cfg *config.Config) (*OnlineTablesPreviewClient, error) {
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

	return &OnlineTablesPreviewClient{
		Config:                       cfg,
		apiClient:                    apiClient,
		OnlineTablesPreviewInterface: NewOnlineTablesPreview(databricksClient),
	}, nil
}

type QualityMonitorsPreviewClient struct {
	QualityMonitorsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewQualityMonitorsPreviewClient(cfg *config.Config) (*QualityMonitorsPreviewClient, error) {
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

	return &QualityMonitorsPreviewClient{
		Config:                          cfg,
		apiClient:                       apiClient,
		QualityMonitorsPreviewInterface: NewQualityMonitorsPreview(databricksClient),
	}, nil
}

type RegisteredModelsPreviewClient struct {
	RegisteredModelsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewRegisteredModelsPreviewClient(cfg *config.Config) (*RegisteredModelsPreviewClient, error) {
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

	return &RegisteredModelsPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		RegisteredModelsPreviewInterface: NewRegisteredModelsPreview(databricksClient),
	}, nil
}

type ResourceQuotasPreviewClient struct {
	ResourceQuotasPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewResourceQuotasPreviewClient(cfg *config.Config) (*ResourceQuotasPreviewClient, error) {
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

	return &ResourceQuotasPreviewClient{
		Config:                         cfg,
		apiClient:                      apiClient,
		ResourceQuotasPreviewInterface: NewResourceQuotasPreview(databricksClient),
	}, nil
}

type SchemasPreviewClient struct {
	SchemasPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewSchemasPreviewClient(cfg *config.Config) (*SchemasPreviewClient, error) {
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

	return &SchemasPreviewClient{
		Config:                  cfg,
		apiClient:               apiClient,
		SchemasPreviewInterface: NewSchemasPreview(databricksClient),
	}, nil
}

type StorageCredentialsPreviewClient struct {
	StorageCredentialsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewStorageCredentialsPreviewClient(cfg *config.Config) (*StorageCredentialsPreviewClient, error) {
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

	return &StorageCredentialsPreviewClient{
		Config:                             cfg,
		apiClient:                          apiClient,
		StorageCredentialsPreviewInterface: NewStorageCredentialsPreview(databricksClient),
	}, nil
}

type SystemSchemasPreviewClient struct {
	SystemSchemasPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewSystemSchemasPreviewClient(cfg *config.Config) (*SystemSchemasPreviewClient, error) {
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

	return &SystemSchemasPreviewClient{
		Config:                        cfg,
		apiClient:                     apiClient,
		SystemSchemasPreviewInterface: NewSystemSchemasPreview(databricksClient),
	}, nil
}

type TableConstraintsPreviewClient struct {
	TableConstraintsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewTableConstraintsPreviewClient(cfg *config.Config) (*TableConstraintsPreviewClient, error) {
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

	return &TableConstraintsPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		TableConstraintsPreviewInterface: NewTableConstraintsPreview(databricksClient),
	}, nil
}

type TablesPreviewClient struct {
	TablesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewTablesPreviewClient(cfg *config.Config) (*TablesPreviewClient, error) {
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

	return &TablesPreviewClient{
		Config:                 cfg,
		apiClient:              apiClient,
		TablesPreviewInterface: NewTablesPreview(databricksClient),
	}, nil
}

type TemporaryTableCredentialsPreviewClient struct {
	TemporaryTableCredentialsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewTemporaryTableCredentialsPreviewClient(cfg *config.Config) (*TemporaryTableCredentialsPreviewClient, error) {
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

	return &TemporaryTableCredentialsPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		TemporaryTableCredentialsPreviewInterface: NewTemporaryTableCredentialsPreview(databricksClient),
	}, nil
}

type VolumesPreviewClient struct {
	VolumesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewVolumesPreviewClient(cfg *config.Config) (*VolumesPreviewClient, error) {
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

	return &VolumesPreviewClient{
		Config:                  cfg,
		apiClient:               apiClient,
		VolumesPreviewInterface: NewVolumesPreview(databricksClient),
	}, nil
}

type WorkspaceBindingsPreviewClient struct {
	WorkspaceBindingsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewWorkspaceBindingsPreviewClient(cfg *config.Config) (*WorkspaceBindingsPreviewClient, error) {
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

	return &WorkspaceBindingsPreviewClient{
		Config:                            cfg,
		apiClient:                         apiClient,
		WorkspaceBindingsPreviewInterface: NewWorkspaceBindingsPreview(databricksClient),
	}, nil
}
