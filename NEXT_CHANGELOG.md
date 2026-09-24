# NEXT CHANGELOG

## Release v0.183.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add [mason](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/mason) package.
* Add [w.Mason](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/mason#MasonAPI) workspace-level service.
* Add `ServiceCredential` field for [catalog.ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig).
* Add `SecretReference` field for [catalog.ModelProviderServiceConfigProviderSecret](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigProviderSecret).
* Add `OnMaintenanceComplete` and `OnMaintenanceStart` fields for [jobs.JobEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEmailNotifications).
* Add `OnMaintenanceComplete` and `OnMaintenanceStart` fields for [jobs.TaskEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TaskEmailNotifications).
* Add `OnMaintenanceComplete` and `OnMaintenanceStart` fields for [jobs.WebhookNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotifications).
* Add `BudgetPolicyId` and `Tags` fields for [ml.BackfillFeaturesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#BackfillFeaturesRequest).
* Add `BudgetPolicyId` and `Tags` fields for [ml.PurgeFeatureEntitiesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PurgeFeatureEntitiesRequest).