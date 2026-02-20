# NEXT CHANGELOG

## Release v0.111.0

### Breaking Changes

### New Features and Improvements
* Add `DefaultCredentialStrategyProvider` package-level variable to allow 
  overriding the default authentication strategy for all `Config` instances.

### Bug Fixes

### Documentation

### Internal Changes

* Add `NewCredentialsChain` to allow internal tools to configure a custom
  credentials chain that differs from the SDK's default chain. This is not
  intended for end-user consumption.
* Export OIDC credential types (`GitHubOIDCCredentials`, `AzureDevOpsOIDCCredentials`,
  `EnvOIDCCredentials`, `FileOIDCCredentials`) so they can be used as building
  blocks in custom credentials chains.
* Make clusters creation in AWS not depend on cloud parsed from host. Changed default AWS availability for auto-created utility clusters from SPOT to SPOT_WITH_FALLBACK (API default).

### API Changes
