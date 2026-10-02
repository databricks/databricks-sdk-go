# NEXT CHANGELOG

## Release v0.184.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add [mason](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/mason) package.
* Add [w.Mason](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/mason#MasonAPI) workspace-level service.
* Add `AiDecide` method for [w.AiFunctions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions#AiFunctionsAPI) workspace-level service.
* Add `CreateSkill`, `DeleteSkill`, `FinalizeSkill`, `GetSkill`, `ListSkills` and `UpdateSkill` methods for [w.AiGateway](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AiGatewayAPI) workspace-level service.
* Add `ListCommands` method for [w.Sandbox](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#SandboxAPI) workspace-level service.
* Add `ServiceCredential` field for [catalog.ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig).
* Add `SecretReference` field for [catalog.ModelProviderServiceConfigProviderSecret](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigProviderSecret).
* Add `ProjectEnvironment` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
* Add `EnvironmentVariables` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
* Add `EnvironmentVariables` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
* Add `OnMaintenanceComplete` and `OnMaintenanceStart` fields for [jobs.JobEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEmailNotifications).
* Add `EnvironmentVariables` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
* Add `EnvironmentVariables` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
* Add `EnvironmentVariablesKey` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `EnvironmentVariables` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
* Add `EnvironmentVariablesKey` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `EnvironmentVariablesKey` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Add `OnMaintenanceComplete` and `OnMaintenanceStart` fields for [jobs.TaskEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TaskEmailNotifications).
* Add `OnMaintenanceComplete` and `OnMaintenanceStart` fields for [jobs.WebhookNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotifications).
* Add `BudgetPolicyId` and `Tags` fields for [ml.BackfillFeaturesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#BackfillFeaturesRequest).
* Add `JobId` and `PipelineId` fields for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `BudgetPolicyId` and `Tags` fields for [ml.PurgeFeatureEntitiesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PurgeFeatureEntitiesRequest).
* Add `Environment` field for [sandbox.SandboxSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#SandboxSpec).
* Add `TableForeignDeltaDeltasharing` enum value for [catalog.SecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableKind).
* [Breaking] Change `CommandPath` field for [jobs.DeploymentSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DeploymentSpec) to no longer be required.
* Change `CommandPath` field for [jobs.DeploymentSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DeploymentSpec) to no longer be required.
* [Breaking] Change `ApiSecretRef` field for [ml.SchemaRegistryConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SchemaRegistryConfig) to no longer be required.
* Change `ApiSecretRef` field for [ml.SchemaRegistryConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SchemaRegistryConfig) to no longer be required.