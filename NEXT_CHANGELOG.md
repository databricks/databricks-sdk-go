# NEXT CHANGELOG

## Release v0.108.0

### Breaking Changes

### New Features and Improvements
* Use profile name as OAuth token cache key instead of host URL, so two profiles on the same host store separate tokens.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add [networking](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/networking) package.
* Add [a.Endpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/networking#EndpointsAPI) account-level service.
* Add `FilterCondition` and `Transformations` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource).
* Add `BudgetPolicyId` and `CustomTags` fields for [postgres.ProjectSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectSpec).
* Add `BudgetPolicyId` and `CustomTags` fields for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `CreateSpace`, `DeleteSpace`, `GetSpace`, `GetSpaceOperation`, `ListSpaces` and `UpdateSpace` methods for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsAPI) workspace-level service.
* Add `Space` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `Space` field for [apps.ListAppsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ListAppsRequest).
* [Breaking] Remove `FilterCondition` and `Transformations` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource).
