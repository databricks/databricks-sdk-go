Databricks SDK for Go
---

Initial commit includes porting of Core Authentication layer.

## High level auth flow

```mermaid
sequenceDiagram
    ClustersAPI->>+DatabricksClient: GET .../clusters/list
    DatabricksClient->>+AuthRegistry: lazy auth
    AuthRegistry-->>+FirstAuthorizer: try configure
    FirstAuthorizer-->>-AuthRegistry: try next
    AuthRegistry->>+NextAuthorizer: try configure
    NextAuthorizer->>RequestVisitor: configured auth
    NextAuthorizer->>-AuthRegistry: authenticated
    AuthRegistry->>+RequestVisitor: visit HTTP request
    RequestVisitor-->>+IdentityProvider: ensure fresh token
    IdentityProvider-->>-RequestVisitor: access token
    RequestVisitor->>-AuthRegistry: added HTTP headers
    AuthRegistry->>-DatabricksClient: added HTTP headers

    DatabricksClient->>+API: authenticated request
    API->>-DatabricksClient: JSON payload
    DatabricksClient->>-ClustersAPI: ClustersList or error
```

## Almost ideal design

```mermaid
classDiagram
    class DatabricksClient {
        - requestVisitor
        - retryPolicy
        - authDiscovery
        Get(path, query) T
        Post(path, body) T
        Put(path, body) T
        Patch(path, body) T
        Delete(path, query) T
    }

    Authorizer --> "0..1" RequestVisitor: creates
    class Authorizer {
        <<interface>>
        Name() string
        TryConfigure() RequestVisitor
    }

    class RequestVisitor {
        <<interface>>
        Visit(HttpRequest) error
    }

    RefreshableVisitor ..> RequestVisitor: extends
    class RefreshableVisitor {
        <<interface>>
        - state Token 
        EnsureFresh() error
    }

    RefreshableVisitor --* Token: contains
    class Token {
        - accessToken JWT
        - refreshToken JWT
        - expiresOn Time
        IsExpired() bool
    }

    AzureSpnAuthorizer --> AuthorizerRegistry: combines
    AzureSpnAuthorizer --|> Authorizer: implements
    AzureSpnAuthorizer --> AzureSpnVisitor: creates
    AzureSpnVisitor --|> RefreshableVisitor: implements
    class AzureSpnAuthorizer {
        ClientID string
        SecretID string
        TenantID string
    }

    AzureCliAuthorizer --> AuthorizerRegistry: combines
    AzureCliAuthorizer --|> Authorizer: implements
    AzureCliAuthorizer --> AzureCliVisitor: creates
    AzureCliVisitor --|> RefreshableVisitor: implements

    GoogleAuthorizer --> AuthorizerRegistry: combines
    GoogleAuthorizer --|> Authorizer: implements
    GoogleAuthorizer --> GoogleVisitor: creates
    GoogleVisitor --|> RefreshableVisitor: implements
    class GoogleAuthorizer {
        ServiceAccount
    }
    
    DatabricksOauthAuthorizer --> AuthorizerRegistry: combines
    DatabricksOauthAuthorizer --|> Authorizer: implements
    DatabricksOauthAuthorizer --> DatabricksOauthVisitor: creates
    DatabricksOauthVisitor --|> RefreshableVisitor: implements
    class DatabricksOauthAuthorizer {
        []Scopes
    }
    
    StaticVisitor --|> RequestVisitor: implements
    class StaticVisitor

    DatabricksCliAuthorizer --> AuthorizerRegistry: combines
    DatabricksCliAuthorizer --|> Authorizer: implements
    DatabricksCliAuthorizer --> StaticVisitor: creates
    class DatabricksCliAuthorizer {
        Profile
    }

    PatAuthorizer --> AuthorizerRegistry: combines
    PatAuthorizer --|> Authorizer: implements
    PatAuthorizer --> StaticVisitor: creates
    class PatAuthorizer {
        Token
    }

    PatAuthorizer --> AuthorizerRegistry: combines
    BasicAuthorizer --|> Authorizer: implements
    BasicAuthorizer --> StaticVisitor: creates
    class BasicAuthorizer {
        Username
        Password
    }

    DatabricksClient --* AuthorizerRegistry: (just pretty render)
    AuthorizerRegistry --|> RequestVisitor: implements
    class AuthorizerRegistry {
        - authMutex
    }
```

TODO:
---

- [ ] Try pulling up packages for Azure and Google
- [ ] Pass tests for CommonEnvironmentClient
- [ ] CommandFactory should be done better
- [ ] Propagate optional key-value pairs to user agent
- [ ] Propagate OS as default key-value pair to user agent
- [ ] Get rid of Terraform SDK leftovers
- [ ] Mention contributors from Terraform provider side
- [ ] Support Databricks OAuth
- [ ] Record configFixture{} compliance test JSON
- [ ] Activate custom path visitor for clients where AccountID is set for Accounts API
- [ ] Provide error explanation callback, so that terraform plugin could include documentation, based on context.
- [ ] figure out if we should replace go-retryablehttp (MPL 2.0) with own impl, or something close to vendor/cloud.google.com/go/compute/metadata/retry.go

Mono-packages & spirit:

- [ ] single shipping vehicle per language
- [ ] Databricks Runtime should eventually include Python SDK by default
- [ ] python sdk -> SDK
- [ ] discoverability
- [ ] clean delineation
- [ ] convenience
- [ ] door for opening packages later
- [ ] easy explanation of purpose
- [ ] pulling in dependencies
- [ ] make it very clear
- [ ] add OWNERS approval on go.mod/setup.py changes (well... dependabot?...)
