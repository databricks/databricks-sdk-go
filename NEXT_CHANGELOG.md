# NEXT CHANGELOG

## Release v0.105.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `BaseEnvironment` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
* Add `Filters` field for [jobs.DashboardTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DashboardTask).
* Add `Id` field for [sharing.CreateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CreateRecipient).
* Add `Id` field for [sharing.RecipientInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientInfo).
* Add `Id` field for [sharing.UpdateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateRecipient).
* Add `QueryTags` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest).
* Add `QueryTags` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
* Add `SswsToken` enum value for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).
* Add `UcVolumeMisconfigured` enum value for [compute.EventType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventType).
* [Breaking] Change `Username` field for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User) to no longer be required.