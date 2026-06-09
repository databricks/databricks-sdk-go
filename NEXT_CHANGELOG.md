# NEXT CHANGELOG

## Release v0.142.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add [aisearch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aisearch) and [bundledeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments) packages.
* Add [w.AiSearch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aisearch#AiSearchAPI) workspace-level service.
* Add [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service.
* Add `RunningInstances` field for [apps.ApplicationStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ApplicationStatus).
* Add `CustomMaxRetentionHours` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
* Add `EnvironmentSettings` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
* Add `CustomMaxRetentionHours` field for [catalog.CreateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCatalog).
* Add `EnvironmentSettings` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
* Add `CustomMaxRetentionHours` field for [catalog.CreateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateSchema).
* Add `CustomMaxRetentionHours` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
* Add `CustomMaxRetentionHours` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
* Add `EnvironmentSettings` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
* Add `CustomMaxRetentionHours` field for [catalog.UpdateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSchema).
* Add `StreamSource` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `IngestionConfig` field for [ml.KafkaConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaConfig).
* Add `ClusteringColumns`, `EnableAutoClustering` and `TableProperties` fields for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Add `BranchId` field for [postgres.Branch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Branch).
* Add `CatalogId` field for [postgres.Catalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Catalog).
* Add `DatabaseId` field for [postgres.Database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Database).
* Add `EndpointId` field for [postgres.Endpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Endpoint).
* Add `ProjectId` field for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `RoleId` field for [postgres.Role](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Role).
* Add `SyncedTableId` field for [postgres.SyncedTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTable).
* Add `AllowedDatabricksDestinations` field for [settings.EgressNetworkPolicyNetworkAccessPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyNetworkAccessPolicy).
* Add `Facets`, `QueryColumns` and `SortColumns` fields for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).
* Add `FacetResult` field for [vectorsearch.QueryVectorIndexResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexResponse).
* Add `FacetColumnCount` and `FacetColumns` fields for [vectorsearch.ResultManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ResultManifest).
* Add `DangerouslyForceDiscardAll` field for [workspace.UpdateRepoRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateRepoRequest).
* [Breaking] Remove [bundle](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle) package.
* [Breaking] Remove [w.Bundle](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#BundleAPI) workspace-level service.