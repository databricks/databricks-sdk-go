Databricks SDK for Go
---

Initial commit includes porting of Core Authentication layer.

## High level auth flow

```mermaid
sequenceDiagram
    ClustersAPI->>+DatabricksClient: GET .../clusters/list
    DatabricksClient->>+databricks.Config: Authenticate(HttpRequest)

    databricks.Config-->>+DefaultCredentials: Configure(databricks.Config)
    DefaultCredentials-->>+FirstCredentials: try configure
    FirstCredentials-->>-DefaultCredentials: try next
    DefaultCredentials->>+NextCredentials: try configure
    NextCredentials->>RequestVisitor: configured auth
    NextCredentials->>-DefaultCredentials: authenticated
    DefaultCredentials->>-databricks.Config: set AuthType & request visitor

    databricks.Config->>+RequestVisitor: visit HTTP request
    RequestVisitor-->>+IdentityProvider: ensure fresh token
    IdentityProvider-->>-RequestVisitor: access token
    RequestVisitor->>-databricks.Config: added HTTP headers
    
    databricks.Config->>-DatabricksClient: added HTTP headers

    DatabricksClient->>+API: authenticated request
    API->>-DatabricksClient: JSON payload
    DatabricksClient->>-ClustersAPI: ClustersList or error
```

## Almost ideal design

```mermaid
classDiagram
    Credentials "0..1" <-- Config: Configure(self)
    RequestVisitor --* Config: configured auth
    class Config {
        * Host string
        * AzureResourceID string
        * AzureEnvironment string
        Credentials: DefaultCredentials

        Authenticate(HttpRequest) error
    }

    Config --* DatabricksClient
    class DatabricksClient {
        Config
        - retryPolicy
        
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
        Configure(Config) RequestVisitor
    }

    class RequestVisitor {
        <<interface>>
        Visit(HttpRequest) error
    }

    AzureSpnCredentials --* authProviders
    AzureSpnCredentials ..|> Credentials
    class AzureSpnCredentials {
        ClientID string
        SecretID string
        TenantID string
    }

    AzureCliCredentials --* authProviders
    AzureCliCredentials ..|> Credentials

    GoogleCredentials --* authProviders
    GoogleCredentials ..|> Credentials
    class GoogleCredentials {
        ServiceAccount
    }
    
    DatabricksOauthCredentials --* authProviders
    DatabricksOauthCredentials ..|> Credentials
    class DatabricksOauthCredentials {
        []Scopes
    }

    PatCredentials --* authProviders
    PatCredentials ..|> Credentials
    class PatCredentials {
        Token
    }

    BasicCredentials --* authProviders
    BasicCredentials ..|> Credentials
    class BasicCredentials {
        Username
        Password
    }

    DatabricksCliCredentials --* authProviders
    DatabricksCliCredentials ..|> Credentials
    class DatabricksCliCredentials {
        Profile
    }

    authProviders --> DefaultCredentials: for reach ConfigAttributes()
    DefaultCredentials <-- DatabricksCliCredentials: set explicit
    DefaultCredentials ..|> Credentials
    class DefaultCredentials {
        - explicit map[string]string
        - skip CredentialsName
    }

    class authProviders {
        ConfigAttributes()
    }
```


TODO:
---

- [ ] change origin to the databricks-sdk-go
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
