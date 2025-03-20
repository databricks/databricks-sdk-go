// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// These APIs manage metastore assignments to a workspace.
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

// These APIs manage Unity Catalog metastores for an account. A metastore
// contains catalogs that can be associated with workspaces
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

// These APIs manage storage credentials for a particular metastore.
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

// In Databricks Runtime 13.3 and above, you can add libraries and init scripts
// to the `allowlist` in UC so that users can leverage these artifacts on
// compute configured with shared access mode.
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

// A catalog is the first layer of Unity Catalog’s three-level namespace.
// It’s used to organize your data assets. Users can see all catalogs on which
// they have been assigned the USE_CATALOG data permission.
//
// In Unity Catalog, admins and data stewards manage users and their access to
// data centrally across all of the workspaces in a Databricks account. Users in
// different workspaces can share access to the same data, depending on
// privileges granted centrally in Unity Catalog.
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

// Connections allow for creating a connection to an external data source.
//
// A connection is an abstraction of an external data source that can be
// connected from Databricks Compute. Creating a connection object is the first
// step to managing external data sources within Unity Catalog, with the second
// step being creating a data object (catalog, schema, or table) using the
// connection. Data objects derived from a connection can be written to or read
// from similar to other Unity Catalog data objects based on cloud storage.
// Users may create different types of connections with each connection having a
// unique set of configuration options to support credential management and
// other settings.
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

// A credential represents an authentication and authorization mechanism for
// accessing services on your cloud tenant. Each credential is subject to Unity
// Catalog access-control policies that control which users and groups can
// access the credential.
//
// To create credentials, you must be a Databricks account admin or have the
// `CREATE SERVICE CREDENTIAL` privilege. The user who creates the credential
// can delegate ownership to another user or group to manage permissions on it.
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

// An external location is an object that combines a cloud storage path with a
// storage credential that authorizes access to the cloud storage path. Each
// external location is subject to Unity Catalog access-control policies that
// control which users and groups can access the credential. If a user does not
// have access to an external location in Unity Catalog, the request fails and
// Unity Catalog does not attempt to authenticate to your cloud tenant on the
// user’s behalf.
//
// Databricks recommends using external locations rather than using storage
// credentials directly.
//
// To create external locations, you must be a metastore admin or a user with
// the **CREATE_EXTERNAL_LOCATION** privilege.
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

// Functions implement User-Defined Functions (UDFs) in Unity Catalog.
//
// The function implementation can be any SQL expression or Query, and it can be
// invoked wherever a table reference is allowed in a query. In Unity Catalog, a
// function resides at the same level as a table, so it can be referenced with
// the form __catalog_name__.__schema_name__.__function_name__.
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

// In Unity Catalog, data is secure by default. Initially, users have no access
// to data in a metastore. Access can be granted by either a metastore admin,
// the owner of an object, or the owner of the catalog or schema that contains
// the object. Securable objects in Unity Catalog are hierarchical and
// privileges are inherited downward.
//
// Securable objects in Unity Catalog are hierarchical and privileges are
// inherited downward. This means that granting a privilege on the catalog
// automatically grants the privilege to all current and future objects within
// the catalog. Similarly, privileges granted on a schema are inherited by all
// current and future objects within that schema.
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

// A metastore is the top-level container of objects in Unity Catalog. It stores
// data assets (tables and views) and the permissions that govern access to
// them. Databricks account admins can create metastores and assign them to
// Databricks workspaces to control which workloads use each metastore. For a
// workspace to use Unity Catalog, it must have a Unity Catalog metastore
// attached.
//
// Each metastore is configured with a root storage location in a cloud storage
// account. This storage location is used for metadata and managed tables data.
//
// NOTE: This metastore is distinct from the metastore included in Databricks
// workspaces created before Unity Catalog was released. If your workspace
// includes a legacy Hive metastore, the data in that metastore is available in
// a catalog named hive_metastore.
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

// Databricks provides a hosted version of MLflow Model Registry in Unity
// Catalog. Models in Unity Catalog provide centralized access control,
// auditing, lineage, and discovery of ML models across Databricks workspaces.
//
// This API reference documents the REST endpoints for managing model versions
// in Unity Catalog. For more details, see the [registered models API
// docs](/api/workspace/registeredmodels).
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

// Online tables provide lower latency and higher QPS access to data from Delta
// tables.
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

// A monitor computes and monitors data or model quality metrics for a table
// over time. It generates metrics tables and a dashboard that you can use to
// monitor table health and set alerts.
//
// Most write operations require the user to be the owner of the table (or its
// parent schema or parent catalog). Viewing the dashboard, computed metrics, or
// monitor configuration only requires the user to have **SELECT** privileges on
// the table (along with **USE_SCHEMA** and **USE_CATALOG**).
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

// Databricks provides a hosted version of MLflow Model Registry in Unity
// Catalog. Models in Unity Catalog provide centralized access control,
// auditing, lineage, and discovery of ML models across Databricks workspaces.
//
// An MLflow registered model resides in the third layer of Unity Catalog’s
// three-level namespace. Registered models contain model versions, which
// correspond to actual ML models (MLflow models). Creating new model versions
// currently requires use of the MLflow Python client. Once model versions are
// created, you can load them for batch inference using MLflow Python client
// APIs, or deploy them for real-time serving using Databricks Model Serving.
//
// All operations on registered models and model versions require USE_CATALOG
// permissions on the enclosing catalog and USE_SCHEMA permissions on the
// enclosing schema. In addition, the following additional privileges are
// required for various operations:
//
// * To create a registered model, users must additionally have the CREATE_MODEL
// permission on the target schema. * To view registered model or model version
// metadata, model version data files, or invoke a model version, users must
// additionally have the EXECUTE permission on the registered model * To update
// registered model or model version tags, users must additionally have APPLY
// TAG permissions on the registered model * To update other registered model or
// model version metadata (comments, aliases) create a new model version, or
// update permissions on the registered model, users must be owners of the
// registered model.
//
// Note: The securable type for models is "FUNCTION". When using REST APIs (e.g.
// tagging, grants) that specify a securable type, use "FUNCTION" as the
// securable type.
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

// Unity Catalog enforces resource quotas on all securable objects, which limits
// the number of resources that can be created. Quotas are expressed in terms of
// a resource type and a parent (for example, tables per metastore or schemas
// per catalog). The resource quota APIs enable you to monitor your current
// usage and limits. For more information on resource quotas see the [Unity
// Catalog documentation].
//
// [Unity Catalog documentation]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#resource-quotas
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

// A schema (also called a database) is the second layer of Unity Catalog’s
// three-level namespace. A schema organizes tables, views and functions. To
// access (or list) a table or view in a schema, users must have the USE_SCHEMA
// data permission on the schema and its parent catalog, and they must have the
// SELECT permission on the table or view.
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

// A storage credential represents an authentication and authorization mechanism
// for accessing data stored on your cloud tenant. Each storage credential is
// subject to Unity Catalog access-control policies that control which users and
// groups can access the credential. If a user does not have access to a storage
// credential in Unity Catalog, the request fails and Unity Catalog does not
// attempt to authenticate to your cloud tenant on the user’s behalf.
//
// Databricks recommends using external locations rather than using storage
// credentials directly.
//
// To create storage credentials, you must be a Databricks account admin. The
// account admin who creates the storage credential can delegate ownership to
// another user or group to manage permissions on it.
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

// A system schema is a schema that lives within the system catalog. A system
// schema may contain information about customer usage of Unity Catalog such as
// audit-logs, billing-logs, lineage information, etc.
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

// Primary key and foreign key constraints encode relationships between fields
// in tables.
//
// Primary and foreign keys are informational only and are not enforced. Foreign
// keys must reference a primary key in another table. This primary key is the
// parent constraint of the foreign key and the table this primary key is on is
// the parent table of the foreign key. Similarly, the foreign key is the child
// constraint of its referenced primary key; the table of the foreign key is the
// child table of the primary key.
//
// You can declare primary keys and foreign keys as part of the table
// specification during table creation. You can also add or drop constraints on
// existing tables.
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

// A table resides in the third layer of Unity Catalog’s three-level
// namespace. It contains rows of data. To create a table, users must have
// CREATE_TABLE and USE_SCHEMA permissions on the schema, and they must have the
// USE_CATALOG permission on its parent catalog. To query a table, users must
// have the SELECT permission on the table, and they must have the USE_CATALOG
// permission on its parent catalog and the USE_SCHEMA permission on its parent
// schema.
//
// A table can be managed or external. From an API perspective, a __VIEW__ is a
// particular kind of table (rather than a managed or external table).
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

// Temporary Table Credentials refer to short-lived, downscoped credentials used
// to access cloud storage locationswhere table data is stored in Databricks.
// These credentials are employed to provide secure and time-limitedaccess to
// data in cloud environments such as AWS, Azure, and Google Cloud. Each cloud
// provider has its own typeof credentials: AWS uses temporary session tokens
// via AWS Security Token Service (STS), Azure utilizesShared Access Signatures
// (SAS) for its data storage services, and Google Cloud supports temporary
// credentialsthrough OAuth 2.0.Temporary table credentials ensure that data
// access is limited in scope and duration, reducing the risk ofunauthorized
// access or misuse. To use the temporary table credentials API, a metastore
// admin needs to enable the external_access_enabled flag (off by default) at
// the metastore level, and user needs to be granted the EXTERNAL USE SCHEMA
// permission at the schema level by catalog admin. Note that EXTERNAL USE
// SCHEMA is a schema level permission that can only be granted by catalog admin
// explicitly and is not included in schema ownership or ALL PRIVILEGES on the
// schema for security reason.
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

// Volumes are a Unity Catalog (UC) capability for accessing, storing,
// governing, organizing and processing files. Use cases include running machine
// learning on unstructured data such as image, audio, video, or PDF files,
// organizing data sets during the data exploration stages in data science,
// working with libraries that require access to the local file system on
// cluster machines, storing library and config files of arbitrary formats such
// as .whl or .txt centrally and providing secure access across workspaces to
// it, or transforming and querying non-tabular data files in ETL.
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

// A securable in Databricks can be configured as __OPEN__ or __ISOLATED__. An
// __OPEN__ securable can be accessed from any workspace, while an __ISOLATED__
// securable can only be accessed from a configured list of workspaces. This API
// allows you to configure (bind) securables to workspaces.
//
// NOTE: The __isolation_mode__ is configured for the securable itself (using
// its Update method) and the workspace bindings are only consulted when the
// securable's __isolation_mode__ is set to __ISOLATED__.
//
// A securable's workspace bindings can be configured by a metastore admin or
// the owner of the securable.
//
// The original path (/api/2.1/unity-catalog/workspace-bindings/catalogs/{name})
// is deprecated. Please use the new path
// (/api/2.1/unity-catalog/bindings/{securable_type}/{securable_name}) which
// introduces the ability to bind a securable in READ_ONLY mode (catalogs only).
//
// Securable types that support binding: - catalog - storage_credential -
// external_location
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
