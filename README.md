Databricks SDK for Go
---

# Features

Databricks SDK for Go comes with a number of features to accelerate the development on top of the Databricks Lakehouse. It covers 100% of Databricks Public REST API surface. Internal HTTP client is robust and handles failures on different levels by performing intelligent retries.

## Configuration

If you use conventional environment variables, the only thing required to start working with Databricks Workspaces is the following snippet:

```go
w := workspaces.New()
w./*press TAB for autocompletion*/
```

The conventional name for the variable that holds workspace-level client of Databricks SDK for Go is `w`, the shorthand for `workspace`.

### Default Authentication FLow

If you run [Databricks Terraform Provider](https://registry.terraform.io/providers/databrickslabs/databricks/latest), Databricks CLI, or applications build with Databricks SDK for other langauges, most likely they'll all interoperate nicely together. By default, the following steps are performed to authenticate the client:

1. Read any direct fields configured in `databricks.Config`.
2. Environment variables are loaded into `*databricks.Config`.
3. Contents of `~/.databrickscfg` is loaded into `*databricks.Config`.
4. `databricks.Config` asserts that exactly one authentication method is configured.
5. Authentication credentials are configured.

The purpose of `~/.databrickscfg` configuration file is hold connectivity profiles with possibly clear-text credentials to Databricks Workspaces or Databricks Accounts. Almost all entries from this configuration file can be set through environment variables. The same configuration file can be read via the official Databricks GoLang SDK and Databricks Python SDK. Legacy Databricks CLI supports reading only `host`, `token`, `username`, and `password` configuration options.

### Databricks Native Authentication
 * `host` _(string)_: Databricks host for either workspace endpoint or Accounts endpoint. Environment: `DATABRICKS_HOST`.
 * `account_id` _(string)_: Databricks Account ID for Accounts endpoint. Only has effect when `host` is either `https://accounts.cloud.databricks.com/`, `https://accounts.azuredatabricks.net/`, or `https://accounts.gcp.databricks.com/`.  Environment: `DATABRICKS_ACCOUNT_ID`.
 * `token` _(string)_: Personal Access Token (PAT). Environment: `DATABRICKS_TOKEN`.
 * `username` _(string)_: Username part of basic authentication. Only possible when `host` is `*.cloud.databricks.com` _(AWS)_. Environment: `DATABRICKS_USERNAME`.
 * `password` _(string)_: Password part of basic authentication. Only possible when `host` is `*.cloud.databricks.com` _(AWS)_. Environment: `DATABRICKS_PASSWORD`.

### Google Cloud Platform Native Authentication

Databricks SDK for Go picks up OAuth token in scope of Google Default Application Credentials (DAC) flow. This means that if you have run `gcloud auth application-default login` on your development machine, or launch the application on the compute, that is allowed to impersonate Service Account specified in `google_service_account`, authentication should work our of the box.

 * `google_service_account` _(string)_: Google Compute Platform (GCP) Service Account e-mail used for impersonation in the Default Application Credentials Flow that does not require a password. Environment: `DATABRICKS_GOOGLE_SERVICE_ACCOUNT`
 * `google_credentials` _(string)_: GCP Service Account Credentials JSON or the location of these credentials on the local filesustem. Environment: `GOOGLE_CREDENTIALS`.

### Azure Native Authentication

Databricks SDK for Go picks up Azure CLI token, if you've previously authenticated as User Principal by running `az login` on your machine. If you need to authenticate as Azure Active Directory Service Principal, you have to specify `azure_workspace_resource_id` and `azure_use_msi` or `azure_tenant_id`, `azure_client_id`, and `azure_client_secret`.

 * `azure_workspace_resource_id` _(string)_: Azure Resource Manager ID for Azure Databricks workspace, which is exhanged for a Host. Environment: `DATABRICKS_AZURE_RESOURCE_ID`.
 * `azure_use_msi` _(string)_: Instruct to use Azure Managed Service Identity passwordless authentication flow for Service Principals. Environment: `ARM_USE_MSI`. _This feature is not yet implemented in Databricks SDK for Go._
 * `azure_client_secret` _(string)_: Azure Active Directory Service Principal secret. Environment: `ARM_CLIENT_SECRET`.
 * `azure_client_id` _(string)_: Azure Active Directory Service Principal Application ID. Environment: `ARM_CLIENT_ID`
 * `azure_tenant_id` _(string)_: Azure Active Directory Tenant ID. Environment: `ARM_TENANT_ID`
 * `azure_environment` _(string)_: Azure Environment (Public, UsGov, China, Germany) has specific set of API endpoints. Defaults to `PUBLIC`. Environment: `ARM_ENVIRONMENT`.

### Overriding `~/.databrickscfg` 
 * `profile` _(string)_: Connection profile specified within `~/.databrickscfg`. Environment: `DATABRICKS_CONFIG_PROFILE`.
 * `config_file` _(string)_: Location of the Databricks CLI credentials file. By default, it is located in `~/.databrickscfg`. Environment: `DATABRICKS_CONFIG_FILE`.

### Other configuration
 * `auth_type` _(string)_: When multiple auth attributes are available in the environment, use the auth type specified by this argument. This argument also holds currently selected auth.
 * `http_timeout_seconds` _(int)_: Number of seconds for HTTP timeout. Default is _60_.
 * `retry_timeout_seconds` _(int)_: Number of seconds to keep retrying HTTP requests. Default is _300 (5 minutes)_.
 * `debug_truncate_bytes` _(int)_: Truncate JSON fields in debug logs above this limit. Default is 96. Environment: `DATABRICKS_DEBUG_TRUNCATE_BYTES`
 * `debug_headers` _(bool)_: Debug HTTP headers of requests made by the application. Default is false, as headers contain sensitive data, like tokens. Environment: `DATABRICKS_DEBUG_HEADERS`.
 * `rate_limit` _(int)_: Maximum number of requests per second made to Databricks REST API. Environment: `DATABRICKS_RATE_LIMIT`

## Paginated responses

**Stability:** _Experimental_

On the platform side, some Databricks APIs have result pagination and some of them don't. Some APIs follow the offset+limit pagination, some start their offsets from 0, some from 1, some use the cursor-based iteration, and others just return all results in the single response. Databricks SDK for Go hides this intricate complexity and generates a more high-level interface for retrieving all results of a certain entity type. Naming pattern is `XxxAll`, where `Xxx` is the name of the method to retrieve a single page of results.

```go
all, err := w.Repos.ListAll(ctx, repos.ListRequest{})
if err != nil {
    return fmt.Errorf("list repos: %w", err)
}
for _, repo := range all {
    println(repo.Path)
}
```

## Node type and Databricks Runtime selectors

**Stability:** _Beta_

Databricks SDK for Go provides selector methods, that make developing multi-cloud applications easier and just rely on characteristics of the virtual machine, like number of cores or availability of local disks, or always picking up the latest Databricks Runtime for the interactive cluster or per-job cluster.

```go
// Fetch list of spark runtime versions
sparkVersions, err := w.Clusters.SparkVersions(ctx)
if err != nil {
    return err
}

// Select the latest LTS version
latestLTS, err := sparkVersions.Select(clusters.SparkVersionRequest{
    Latest:          true,
    LongTermSupport: true,
})
if err != nil {
    return err
}

// Fetch list of available node types
nodeTypes, err := w.Clusters.ListNodeTypes(ctx)
if err != nil {
    return err
}

// Select the smallest node type id
smallestWithDisk, err := nodeTypes.Smallest(clusters.NodeTypeRequest{
    LocalDisk: true,
})
if err != nil {
    return err
}

// Create cluster and wait for it to start properly
runningCluster, err := w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
    ClusterName:            clusterName,
    SparkVersion:           latestLTS,
    NodeTypeId:             smallestWithDisk,
    AutoterminationMinutes: 15,
    NumWorkers:             1,
})
```

## `io.Reader` integration for DBFS

**Stability:** _Beta_

Please use higher-level `w.Dbfs.Open` and `w.Dbfs.Overwrite` to work with remote files via `io.Reader` interface. Internally, these methods wrap the low-level intricacies of working with Databricks REST APIs, providing a convenient interface to you as a developer.

```go
upload, _ := os.Open("/path/to/local/file.ext")
_ = w.Dbfs.Overwrite(ctx, "/path/to/remote/file", upload)

download, _ := os.Create("/path/to/local")
remote, _ := w.Dbfs.Open(ctx, "/path/to/remote")
_ = io.Copy(download, remote)
```

---

Work in progress

TODO:
---

- [ ] Azure MSI Auth ported
- [ ] Try pulling up packages for Azure and Google
- [ ] Pass tests for CommonEnvironmentClient
- [ ] CommandFactory should be done better
- [ ] Mention contributors from Terraform provider side
- [ ] Support Databricks OAuth
- [ ] Record configFixture{} compliance test JSON
- [ ] Activate custom path visitor for clients where AccountID is set for Accounts API
- [ ] Provide error explanation callback, so that terraform plugin could include documentation, based on context.

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
