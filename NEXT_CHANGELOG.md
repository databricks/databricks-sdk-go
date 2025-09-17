# NEXT CHANGELOG

## Release v0.84.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Added `JavaDependencies` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
* Added `FollowupQuestions` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment).
* Added `EffectiveCapacity` field for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Added `CreateTime` and `UpdateTime` fields for [tags.TagPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#TagPolicy).
* Added `TableDeltaUniformIcebergForeignDeltasharing` enum value for [catalog.SecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableKind).
* Added `InternalCatalogPathOverlapException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* [Breaking] Removed `DefaultDataSecurityMode` and `EffectiveDefaultDataSecurityMode` fields for [settingsv2.Setting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#Setting).
* Added [a.AccountGroupsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccountGroupsV2API) account-level service, [a.AccountServicePrincipalsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccountServicePrincipalsV2API) account-level service, [a.AccountUsersV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccountUsersV2API) account-level service, [w.GroupsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GroupsV2API) workspace-level service, [w.ServicePrincipalsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ServicePrincipalsV2API) workspace-level service and [w.UsersV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersV2API) workspace-level service.
* Added `NetsuiteJarPath` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Added `WorkdayReportParameters` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Added `InternalCatalogMissingUcPathException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `ListShares` method for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service.
* [Breaking] Removed `List` method for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service.
