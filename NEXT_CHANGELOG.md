# NEXT CHANGELOG

## Release v0.175.0

### Breaking Changes

### New Features and Improvements

* Added the `oauth-m2m-gcp` authentication type, which authenticates with a Databricks OAuth service-principal token and passes a Google Cloud access token through the `X-Databricks-GCP-SA-Access-Token` header. This lets GCP account-level provisioning APIs be called with a Databricks-governed identity when SSO is enabled and Google ID token auth is disabled. It must be selected explicitly via `AuthType: "oauth-m2m-gcp"`.

### Bug Fixes

### Documentation

* Documented the `oauth-m2m-gcp` authentication type in the GCP section of the README.

### Internal Changes

### API Changes
