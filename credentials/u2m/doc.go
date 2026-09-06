// Package u2m provides Databricks OAuth endpoint discovery and legacy
// interactive user-to-machine authentication APIs.
//
// Interactive U2M authentication and token persistence are implemented by
// the Databricks CLI. The interactive types remain available for source
// compatibility but are deprecated. OAuthEndpointSupplier,
// BasicOAuthEndpointSupplier, OAuthAuthorizationServer, and
// ErrOAuthNotSupported remain supported for SDK authentication flows.
package u2m
