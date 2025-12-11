# NEXT CHANGELOG

## Release v0.95.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `EffectiveUsagePolicyId` and `UsagePolicyId` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `ExternalAccessEnabled` field for [catalog.CreateAccountsMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateAccountsMetastore).
* Add `ExternalAccessEnabled` field for [catalog.CreateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastore).
* Add `ExternalAccessEnabled` field for [catalog.UpdateAccountsMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAccountsMetastore).
* Add `ExternalAccessEnabled` field for [catalog.UpdateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastore).
* Add `ControlPlaneConnectionFailure` and `ControlPlaneConnectionFailureDueToMisconfig` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `ControlPlaneConnectionFailure` and `ControlPlaneConnectionFailureDueToMisconfig` enum values for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).