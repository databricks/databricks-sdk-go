// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type AccountMetastoreAssignmentsClient struct {
	AccountMetastoreAssignmentsAPI
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
		AccountMetastoreAssignmentsAPI: AccountMetastoreAssignmentsAPI{
			accountMetastoreAssignmentsImpl: accountMetastoreAssignmentsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type AccountMetastoresClient struct {
	AccountMetastoresAPI
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
		AccountMetastoresAPI: AccountMetastoresAPI{
			accountMetastoresImpl: accountMetastoresImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type AccountStorageCredentialsClient struct {
	AccountStorageCredentialsAPI
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
		AccountStorageCredentialsAPI: AccountStorageCredentialsAPI{
			accountStorageCredentialsImpl: accountStorageCredentialsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ArtifactAllowlistsClient struct {
	ArtifactAllowlistsAPI
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
		ArtifactAllowlistsAPI: ArtifactAllowlistsAPI{
			artifactAllowlistsImpl: artifactAllowlistsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type CatalogsClient struct {
	CatalogsAPI
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
		CatalogsAPI: CatalogsAPI{
			catalogsImpl: catalogsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ConnectionsClient struct {
	ConnectionsAPI
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
		ConnectionsAPI: ConnectionsAPI{
			connectionsImpl: connectionsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type CredentialsClient struct {
	CredentialsAPI
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
		CredentialsAPI: CredentialsAPI{
			credentialsImpl: credentialsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ExternalLocationsClient struct {
	ExternalLocationsAPI
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
		ExternalLocationsAPI: ExternalLocationsAPI{
			externalLocationsImpl: externalLocationsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type FunctionsClient struct {
	FunctionsAPI
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
		FunctionsAPI: FunctionsAPI{
			functionsImpl: functionsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type GrantsClient struct {
	GrantsAPI
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
		GrantsAPI: GrantsAPI{
			grantsImpl: grantsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type MetastoresClient struct {
	MetastoresAPI
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
		MetastoresAPI: MetastoresAPI{
			metastoresImpl: metastoresImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ModelVersionsClient struct {
	ModelVersionsAPI
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
		ModelVersionsAPI: ModelVersionsAPI{
			modelVersionsImpl: modelVersionsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type OnlineTablesClient struct {
	OnlineTablesAPI
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
		OnlineTablesAPI: OnlineTablesAPI{
			onlineTablesImpl: onlineTablesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type QualityMonitorsClient struct {
	QualityMonitorsAPI
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
		QualityMonitorsAPI: QualityMonitorsAPI{
			qualityMonitorsImpl: qualityMonitorsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type RegisteredModelsClient struct {
	RegisteredModelsAPI
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
		RegisteredModelsAPI: RegisteredModelsAPI{
			registeredModelsImpl: registeredModelsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ResourceQuotasClient struct {
	ResourceQuotasAPI
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
		ResourceQuotasAPI: ResourceQuotasAPI{
			resourceQuotasImpl: resourceQuotasImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type SchemasClient struct {
	SchemasAPI
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
		SchemasAPI: SchemasAPI{
			schemasImpl: schemasImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type StorageCredentialsClient struct {
	StorageCredentialsAPI
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
		StorageCredentialsAPI: StorageCredentialsAPI{
			storageCredentialsImpl: storageCredentialsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type SystemSchemasClient struct {
	SystemSchemasAPI
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
		SystemSchemasAPI: SystemSchemasAPI{
			systemSchemasImpl: systemSchemasImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type TableConstraintsClient struct {
	TableConstraintsAPI
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
		TableConstraintsAPI: TableConstraintsAPI{
			tableConstraintsImpl: tableConstraintsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type TablesClient struct {
	TablesAPI
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
		TablesAPI: TablesAPI{
			tablesImpl: tablesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type TemporaryTableCredentialsClient struct {
	TemporaryTableCredentialsAPI
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
		TemporaryTableCredentialsAPI: TemporaryTableCredentialsAPI{
			temporaryTableCredentialsImpl: temporaryTableCredentialsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type VolumesClient struct {
	VolumesAPI
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
		VolumesAPI: VolumesAPI{
			volumesImpl: volumesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type WorkspaceBindingsClient struct {
	WorkspaceBindingsAPI
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
		WorkspaceBindingsAPI: WorkspaceBindingsAPI{
			workspaceBindingsImpl: workspaceBindingsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
