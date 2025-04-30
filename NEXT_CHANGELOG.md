# NEXT CHANGELOG

## Release v0.66.0

### New Features and Improvements

### Bug Fixes
* Tolerate trailing slashes in hostnames in `Config`.

### Documentation

### Internal Changes

### API Changes
* Added [w.AlertsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsV2API) workspace-level service.
* Added `UpdateNccAzurePrivateEndpointRulePublic` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
* Added `CreatedAt`, `CreatedBy` and `MetastoreId` fields for [catalog.SetArtifactAllowlist](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetArtifactAllowlist).
* [Breaking] Added `NetworkConnectivityConfig` field for [settings.CreateNetworkConnectivityConfigRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateNetworkConnectivityConfigRequest).
* [Breaking] Added `PrivateEndpointRule` field for [settings.CreatePrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRuleRequest).
* Added `DomainNames` field for [settings.NccAzurePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRule).
* Added `AutoResolveDisplayName` field for [sql.CreateAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateAlertRequest).
* Added `AutoResolveDisplayName` field for [sql.CreateQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryRequest).
* Added `CreateCleanRoom`, `ExecuteCleanRoomTask` and `ModifyCleanRoom` enum values for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).
* Added `DnsResolutionError` and `GcpDeniedByOrgPolicy` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `Expired` enum value for [settings.NccAzurePrivateEndpointRuleConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleConnectionState).
* [Breaking] Changed `CreateNetworkConnectivityConfiguration` and `CreatePrivateEndpointRule` methods for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service with new required argument order.
* [Breaking] Changed `WorkloadSize` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput) to type `string`.
* [Breaking] Changed `GroupId` field for [settings.NccAzurePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRule) to type `string`.
* [Breaking] Changed `TargetServices` field for [settings.NccAzureServiceEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzureServiceEndpointRule) to type [settings.EgressResourceTypeList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressResourceTypeList).
* [Breaking] Removed `Name` and `Region` fields for [settings.CreateNetworkConnectivityConfigRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateNetworkConnectivityConfigRequest).
* [Breaking] Removed `GroupId` and `ResourceId` fields for [settings.CreatePrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRuleRequest).
* [Breaking] Removed `Large`, `Medium` and `Small` enum values for [serving.ServedModelInputWorkloadSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadSize).
* [Breaking] Removed `Blob`, `Dfs`, `MysqlServer` and `SqlServer` enum values for [settings.NccAzurePrivateEndpointRuleGroupId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleGroupId).
* [Breaking] Field `AppDeployment` of `CreateAppDeploymentRequest` is changed from `*AppDeployment` to `AppDeployment`.
* [Breaking] Field `App` of `CreateAppRequest` is changed from `*App` to `App`.
* [Breaking] Field `App` of `UpdateAppRequest` is changed from `*App` to `App`.
* [Breaking] Field `BudgetPolicy` of `UpdateBudgetPolicyRequest` is changed from `*BudgetPolicy` to `BudgetPolicy`.
* [Breaking] Field `OnlineTable` of `CreateOnlineTableRequest` is changed from `*OnlineTable` to `OnlineTable`.
* [Breaking] Field `CleanRoomAsset` of `CreateCleanRoomAssetRequest` is changed from `*CleanRoomAsset` to `CleanRoomAsset`.
* [Breaking] Field `CleanRoom` of `CreateCleanRoomRequest` is changed from `*CleanRoom` to `CleanRoom`.
* [Breaking] Field `CleanRoomAsset` of `UpdateCleanRoomAssetRequest` is changed from `*CleanRoomAsset` to `CleanRoomAsset`.
* [Breaking] Field `Dashboard` of `CreateDashboardRequest` is changed from `*Dashboard` to `Dashboard`.
* [Breaking] Field `Schedule` of `CreateScheduleRequest` is changed from `*Schedule` to `Schedule`.
* [Breaking] Field `Subscription` of `CreateSubscriptionRequest` is changed from `*Subscription` to `Subscription`.
* [Breaking] Field `Dashboard` of `UpdateDashboardRequest` is changed from `*Dashboard` to `Dashboard`.
