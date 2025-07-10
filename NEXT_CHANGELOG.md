# NEXT CHANGELOG

## Release v0.75.0

### New Features and Improvements

### Bug Fixes

- Default values for `ApiClient` are now properly set in the type constructor rather than in the config constructor. 

### Documentation

### Internal Changes

### API Changes
* Added [w.ServicePrincipalSecretsProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ServicePrincipalSecretsProxyAPI) workspace-level service.
* Added [w.DefaultWarehouseId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DefaultWarehouseIdAPI) workspace-level service.
* Added `Database` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Added `ReadReplicaCount` field for [ml.OnlineStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStore).
* Added `PageSize` field for [oauth2.ListServicePrincipalSecretsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListServicePrincipalSecretsRequest).
* Added `ProjectedRemainingTaskTotalTimeMs`, `RemainingTaskCount`, `RunnableTasks` and `WorkToBeDone` fields for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.CreateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsRequest).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.CreateCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsResponse).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.CredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CredentialInfo).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.GetCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetCredentialsResponse).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.UpdateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateCredentialsRequest).
* Added `Databricks` enum value for [catalog.SystemType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemType).
* Added `DriverDnsResolutionFailure` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `Confluence` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `DeltaIcebergTable` enum value for [sharing.TableInternalAttributesSharedTableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#TableInternalAttributesSharedTableType).
* [Breaking] Changed `Delete` method for [w.TableConstraints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableConstraintsAPI) workspace-level service to return `any`.
* [Breaking] Changed `Lifetime` field for [oauth2.CreateServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretRequest) to type `string`.
* [Breaking] Changed `ServicePrincipalId` field for [oauth2.CreateServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretRequest) to type `string`.
* [Breaking] Changed `ExpireTime` field for [oauth2.CreateServicePrincipalSecretResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretResponse) to type `string`.
* [Breaking] Changed `ServicePrincipalId` field for [oauth2.DeleteServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#DeleteServicePrincipalSecretRequest) to type `string`.
* [Breaking] Changed `ServicePrincipalId` field for [oauth2.ListServicePrincipalSecretsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListServicePrincipalSecretsRequest) to type `string`.
* [Breaking] Changed `ExpireTime` field for [oauth2.SecretInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#SecretInfo) to type `string`.
* [Breaking] Changed `Calls` field for [serving.AiGatewayRateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimit) to no longer be required.
* Changed `Calls` field for [serving.AiGatewayRateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimit) to no longer be required.
