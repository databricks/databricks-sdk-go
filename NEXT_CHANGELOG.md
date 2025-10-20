# NEXT CHANGELOG

## Release v0.87.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `AbsoluteSessionLifetimeInMinutes` and `EnableSingleUseRefreshTokens` fields for [oauth2.TokenAccessPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#TokenAccessPolicy).
* Add `NetworkConnectivityConfigId` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
* Add `OauthMtls` enum value for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).
* Add `NetworkCheckNicFailureDueToMisconfig`, `NetworkCheckDnsServerFailureDueToMisconfig`, `NetworkCheckStorageFailureDueToMisconfig`, `NetworkCheckMetadataEndpointFailureDueToMisconfig`, `NetworkCheckControlPlaneFailureDueToMisconfig` and `NetworkCheckMultipleComponentsFailureDueToMisconfig` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `Creating` and `CreateFailed` enum values for [settings.NccPrivateEndpointRulePrivateLinkConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccPrivateEndpointRulePrivateLinkConnectionState).
* Add `NetworkCheckNicFailureDueToMisconfig`, `NetworkCheckDnsServerFailureDueToMisconfig`, `NetworkCheckStorageFailureDueToMisconfig`, `NetworkCheckMetadataEndpointFailureDueToMisconfig`, `NetworkCheckControlPlaneFailureDueToMisconfig` and `NetworkCheckMultipleComponentsFailureDueToMisconfig` enum values for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).
* [Breaking] Remove `Update` method for [w.RecipientFederationPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientFederationPoliciesAPI) workspace-level service.