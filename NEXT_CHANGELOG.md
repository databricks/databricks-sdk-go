# NEXT CHANGELOG

## Release v0.169.0

### Breaking Changes

### New Features and Improvements

* Add the Pi coding agent (`PI_CODING_AGENT`) to AI agent detection in the User-Agent header, so CLI usage driven by Pi is attributed to `pi`.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `KinesisStreamConfig` field for [ml.StreamSourceConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#StreamSourceConfig).
* Add `Mode` field for [pipelines.UpdateInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfo).
* [Breaking] Change `State` field for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to type `string`.
* [Breaking] Change `State` field for [bundledeployments.Resource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Resource) to type `string`.
* [Breaking] Remove `AwsAccessKeyId` and `AwsSecretAccessKey` fields for [catalog.ModelProviderServiceConfigAmazonBedrockProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigAmazonBedrockProviderDirectConfig).
* [Breaking] Remove `ClientId`, `ClientSecret` and `TenantId` fields for [catalog.ModelProviderServiceConfigAzureOpenAiProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigAzureOpenAiProviderDirectConfig).
* [Breaking] Remove `ClientId`, `ClientSecret` and `TenantId` fields for [catalog.ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig).