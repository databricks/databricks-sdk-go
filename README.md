Databricks SDK for Go
---

Initial commit includes porting of Core Authentication layer.

## High level auth flow

```mermaid
sequenceDiagram
    ClustersAPI->>+DatabricksClient: GET .../clusters/list
    DatabricksClient->>+AuthRegistry: lazy auth
    AuthRegistry-->>+FirstCredentials: try configure
    FirstCredentials-->>-AuthRegistry: try next
    AuthRegistry->>+NextCredentials: try configure
    NextCredentials->>RequestVisitor: configured auth
    NextCredentials->>-AuthRegistry: authenticated
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

    Credentials --> "0..1" RequestVisitor: creates
    class Credentials {
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

    AzureSpnCredentials --> CredentialsRegistry: combines
    AzureSpnCredentials --|> Credentials: implements
    AzureSpnCredentials --> AzureSpnVisitor: creates
    AzureSpnVisitor --|> RefreshableVisitor: implements
    class AzureSpnCredentials {
        ClientID string
        SecretID string
        TenantID string
    }

    AzureCliCredentials --> CredentialsRegistry: combines
    AzureCliCredentials --|> Credentials: implements
    AzureCliCredentials --> AzureCliVisitor: creates
    AzureCliVisitor --|> RefreshableVisitor: implements

    GoogleCredentials --> CredentialsRegistry: combines
    GoogleCredentials --|> Credentials: implements
    GoogleCredentials --> GoogleVisitor: creates
    GoogleVisitor --|> RefreshableVisitor: implements
    class GoogleCredentials {
        ServiceAccount
    }
    
    DatabricksOauthCredentials --> CredentialsRegistry: combines
    DatabricksOauthCredentials --|> Credentials: implements
    DatabricksOauthCredentials --> DatabricksOauthVisitor: creates
    DatabricksOauthVisitor --|> RefreshableVisitor: implements
    class DatabricksOauthCredentials {
        []Scopes
    }
    
    StaticVisitor --|> RequestVisitor: implements
    class StaticVisitor

    DatabricksCliCredentials --> CredentialsRegistry: combines
    DatabricksCliCredentials --|> Credentials: implements
    DatabricksCliCredentials --> StaticVisitor: creates
    class DatabricksCliCredentials {
        Profile
    }

    PatCredentials --> CredentialsRegistry: combines
    PatCredentials --|> Credentials: implements
    PatCredentials --> StaticVisitor: creates
    class PatCredentials {
        Token
    }

    PatCredentials --> CredentialsRegistry: combines
    BasicCredentials --|> Credentials: implements
    BasicCredentials --> StaticVisitor: creates
    class BasicCredentials {
        Username
        Password
    }

    DatabricksClient --* CredentialsRegistry: (just pretty render)
    CredentialsRegistry --|> RequestVisitor: implements
    class CredentialsRegistry {
        - authMutex
    }
```

TODO:
---

- [ ] Azure MSI Auth ported
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
