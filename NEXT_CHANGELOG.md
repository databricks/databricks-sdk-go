# NEXT CHANGELOG

## Release v0.83.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [iamv2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2) package.
* Added [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service and [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service.
* Added `Feedback` field for [dashboards.GenieMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieMessage).
* Added `Disabled` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `AuxiliaryManagedLocation` field for [sharing.TableInternalAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#TableInternalAttributes).
* Added `Alerts` field for [sql.ListAlertsV2Response](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsV2Response).
* Added `NoActivatedK8s` and `UsagePolicyEntitlementDenied` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `ForeignCatalog` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `ForeignIcebergTable` enum value for [sharing.TableInternalAttributesSharedTableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#TableInternalAttributesSharedTableType).
* [Breaking] Removed `Disabled` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).