# Databricks SDK for Go (Beta)

The Databricks SDK for Go includes functionality to accelerate development for the Databricks Lakehouse. It covers all Databricks public REST API operations. The SDK's internal HTTP client is robust and handles failures on different levels by performing intelligent retries.

## Getting started

During the Beta period, you must clone and then reference this repository locally, as follows. After the Beta period, you will be able to download and reference the Databricks SDK for Go package by using familiar commands such as `go get` and `go install`.

1. On your local development machine with Go already [installed](https://go.dev/doc/install), clone this repository into your existing Go code project's current working directory, by running the `git clone` command:

   ```bash
   git clone git@github.com:databricks/databricks-sdk-go.git
   ```
   
   This command creates a folder named `databricks-sdk-go` within your project. This folder contains a local copy of the Databricks SDK for Go source code.
   
2. Create a `go.mod` file to track your Go code's dependencies by running the `go mod init` command, for example:

   ```bash
   go mod init sample
   ```
   
3. Take a dependency on the Databricks SDK for Go package by running the `go mod edit -require` command:

   ```bash
   go mod edit -require github.com/databricks/databricks-sdk-go@v0.0.0
   ```
   
4. Redirect the dependency to your locally downloaded copy of the Databricks SDK for Go by running the `go mod edit -replace` command:

   ```bash
   go mod edit -replace github.com/databricks/databricks-sdk-go@v0.0.0=./databricks-sdk-go/ 
   ```

   Your `go.mod` file should now look like this:
   
   ```go
   module sample

   go 1.18

   require github.com/databricks/databricks-sdk-go v0.0.0

   replace github.com/databricks/databricks-sdk-go v0.0.0 => ./databricks-sdk-go/
   ```

5. Within your project, create a Go code file that imports the Databricks SDK for Go, for example a file named `main.go` with the following contents:

   ```go
   package main

   import (
     "context"
     "fmt"
     
     "github.com/databricks/databricks-sdk-go/workspaces"
   )

   func main() {
     const path = "/"
     w := workspaces.New()

     resp, err := w.Workspace.GetStatusByPath(context.Background(), path)

     if err != nil {
       panic(err)
     }

     fmt.Printf("The path '%s' is a '%s'.\n", path, resp.ObjectType)
   }
   ```

5. Add any misssing module dependencies by running the `go mod tidy` command:

   ```bash
   go mod tiny
   ```
   
6. Grab copies of all packages needed to support builds and tests of packages in your `main` module, by running the `go mod vendor` command:

   ```bash
   go mod vendor
   ```

7. Set up Databricks authentication on your local development machine, if you have not done so already. For details, see the next section, "Authentication."
8. Run your Go code file, assuming a file named `main.go`, by running the `go run` command:

   ```bash
   go run main.go
   ```
   
   Assuming the preceding example code is run, the output is: 
   
   ```bash
   The path '/' is a 'DIRECTORY'.   
   ```

## Authentication

If you use a Databricks configuration profiles file or Databricks-specific environment variables for Databricks authentication, the only code required to start working with a Databricks workspace is the following code snippet:

```go
w := workspaces.New()
w./*press TAB for autocompletion*/
```

The conventional name for the variable that holds the workspace-level client of the Databricks SDK for Go is `w`, which is shorthand for `workspace`.

### Default authentication flow

If you run the [Databricks Terraform Provider](https://registry.terraform.io/providers/databrickslabs/databricks/latest), the [Databricks CLI](https://docs.databricks.com/dev-tools/cli/index.html), or applications that target the Databricks SDKs for other langauges, most likely they will all interoperate nicely together. By default, the Databricks SDK for Go checks the following to perform Databricks authentication, in the following order:

1. Read any direct fields that are hard-coded into `databricks.Config`.
2. Read locally-set Databricks-specific environment variables for authentication into `*databricks.Config`.
3. Read the contents of the local `.databrickscfg` file's `DEFAULT` profile from its default location (`~` for Linux or macOS, and `%USERPROFILE%` for Windows) into `*databricks.Config`.
4. If any direct fields are NOT hard-coded in `databricks.Config`, AND Databricks-specific environment variables exist AND a `.databrickscfg` file exists in the default location with a `DEFAULT` profile, then the Databricks SDK for Go stops running and returns a "multiple authentication credentials exist" error.
5. Otherwise, the Databricks SDK for Go performs Databricks authentication by using the information in `*databricks.Config`.

Depending on the type of authentication, the Databricks SDK for Go uses the following information. Presented are the `*databricks.Config` fields, their descriptions, any corresponding environment variables, and any corresponding `.databrickscfg` file fields, respectively.

### Databricks native authentication

 * `host` _(string)_: The Databricks host URL for either the Databricks workspace endpoint or the Databricks accounts endpoint. Environment: `DATABRICKS_HOST`. `.databrickscfg` file: `host`.
 * `account_id` _(string)_: The Databricks account ID for the Databricks accounts endpoint. Only has effect when `host` is either `https://accounts.cloud.databricks.com/` _(AWS)_, `https://accounts.azuredatabricks.net/` _(Azure)_, or `https://accounts.gcp.databricks.com/` _(GCP)_.  Environment: `DATABRICKS_ACCOUNT_ID`. `.databrickscfg` file: `host`.
 * `token` _(string)_: The Databricks personal access token (PAT) or Azure Active Directory (Azure AD) token. Environment: `DATABRICKS_TOKEN`. `.databrickscfg` file: `token`.
 * `username` _(string)_: The Databricks username part of basic authentication. Only possible when `host` is `*.cloud.databricks.com` _(AWS)_. Environment: `DATABRICKS_USERNAME`. `.databrickscfg` file: `username`.
 * `password` _(string)_: The Databricks password part of basic authentication. Only possible when `host` is `*.cloud.databricks.com` _(AWS)_. Environment: `DATABRICKS_PASSWORD`. `.databrickscfg` file: `password`.

### Google Cloud Platform native authentication

The Databricks SDK for Go picks up an OAuth token in the scope of the Google Default Application Credentials (DAC) flow. This means that if you have run `gcloud auth application-default login` on your development machine, or launch the application on the compute, that is allowed to impersonate the Google Cloud service account specified in `google_service_account`. Authentication should then work out of the box.  There are no `.databrickscfg` file equivalents for the following values.

 * `google_service_account` _(string)_: The Google Compute Platform (GCP) service account e-mail used for impersonation in the Default Application Credentials Flow that does not require a password. Environment: `DATABRICKS_GOOGLE_SERVICE_ACCOUNT`.
 * `google_credentials` _(string)_: GCP Service Account Credentials JSON or the location of these credentials on the local filesustem. Environment: `GOOGLE_CREDENTIALS`.

### Azure native authentication

The Databricks SDK for Go picks up an Azure CLI token, if you've previously authenticated as an Azure user by running `az login` on your machine. If you need to authenticate as Azure Active Directory service principal, you nyst specify `azure_workspace_resource_id` and `azure_use_msi`; or `azure_tenant_id`, `azure_client_id`, and `azure_client_secret`. There are no `.databrickscfg` file equivalents for the following values.

 * `azure_workspace_resource_id` _(string)_: The Azure Resource Manager ID for the Azure Databricks workspace, which is exchanged for a Host. Environment: `DATABRICKS_AZURE_RESOURCE_ID`.
 * `azure_use_msi` _(string)_: `true` to use Azure Managed Service Identity passwordless authentication flow for service principals. Environment: `ARM_USE_MSI`. _This feature is not yet implemented in the Databricks SDK for Go._
 * `azure_client_secret` _(string)_: The Azure Active Directory service principal secret. Environment: `ARM_CLIENT_SECRET`.
 * `azure_client_id` _(string)_: The Azure Active Directory service principal Application ID. Environment: `ARM_CLIENT_ID`.
 * `azure_tenant_id` _(string)_: The Azure Active Directory tenant ID. Environment: `ARM_TENANT_ID`.
 * `azure_environment` _(string)_: the Azure environment type (such as Public, UsGov, China, and Germany) for a specific set of API endpoints. Defaults to `PUBLIC`. Environment: `ARM_ENVIRONMENT`.

### Overriding `.databrickscfg` 
 * `profile` _(string)_: A connection profile specified within `.databrickscfg` to use instead of `DEFAULT`. Environment: `DATABRICKS_CONFIG_PROFILE`.
 * `config_file` _(string)_: A non-default location of the Databricks CLI credentials file. Environment: `DATABRICKS_CONFIG_FILE`.

### Additional authentication configuration options
 * `auth_type` _(string)_: When multiple auth attributes are available in the environment, use the auth type specified by this argument. This argument also holds the currently selected auth.
 * `http_timeout_seconds` _(int)_: Number of seconds for HTTP timeout. Default is _60_.
 * `retry_timeout_seconds` _(int)_: Number of seconds to keep retrying HTTP requests. Default is _300 (5 minutes)_.
 * `debug_truncate_bytes` _(int)_: Truncate JSON fields in debug logs above this limit. Default is 96. Environment: `DATABRICKS_DEBUG_TRUNCATE_BYTES`.
 * `debug_headers` _(bool)_: `true` to debug HTTP headers of requests made by the application. Default is `false`, as headers contain sensitive data, such as access tokens. Environment: `DATABRICKS_DEBUG_HEADERS`.
 * `rate_limit` _(int)_: Maximum number of requests per second made to Databricks REST API. Environment: `DATABRICKS_RATE_LIMIT`.

### Custom credentials provider

In some cases, you may want to have deeper control over authentication to Databricks. This can be achieved by creating your own credentials provider that returns an HTTP request visitor:

```go
type CustomCredentials struct {}

func (c *CustomCredentials) Name() string {
	return "custom"
}

func (c *CustomCredentials) Configure(ctx context.Context, cfg *databricks.Config) (func(*http.Request) error, error) {
	return func(r *http.Request) error {
		token := "..."
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}, nil
}

func main() {
	w := workspaces.New(&databricks.Config{
		Credentials: &CustomCredentials{},
	})
    // ..
}
```

## Long-running operations

**Stability:** [Experimental](https://docs.databricks.com/release-notes/release-types.html)

More than 20 methods across different Databricks APIs are long-running operations for managing things like clusters, command execution, jobs, libraries, Delta Live Tables pipelines, and Databricks SQL warehouses. For example, in the Clusters API, once you create a cluster, you receive a cluster ID, and the cluster is in the `PENDING` state while Databricks takes care of provisioning virtual machines from the cloud provider in the background. But the cluster is only usable in the `RUNNING` state. Another example is the API for running a job or repairing the run: right after the run starts, the run is in the `PENDING` state, though the job is considered to be finished only when it is in the `TERMINATED` or `SKIPPED` states. And of course you. would want to know the error message when the long-running operation times out or why things fail. And sometimes you want to configure a custom timeout other than the default of 20 minutes. 

To hide all of the integration-specific complexity from the end user, Databricks SDK for Go provides a high-level API for _triggering_ the long-running operations and _waiting_ for the releated entities to reach the right state or return back the error message about the problem in case of failure. All long-running operations have the `XxxAndWait` name pattern, where `Xxx` is the operation name. All these generated methods return information about the relevant entity once the operation is finished. It is possible to configure a custom timeout to `XxxAndWait` by providing a functional option argument constructed by `retries.Timeout[Zzz](time.Duration)` function, where `Zzz` is the result type of `XxxAndWait`.

In the following example, `CreateAndWait` returns `ClusterInfo` only once the cluster is in the `RUNNING` state, otherwise it will timeout in 10 minutes:

```go
clusterInfo, err = w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
    ClusterName:            "Created cluster",
    SparkVersion:           latestLTS,
    NodeTypeId:             smallestWithDisk,
    AutoterminationMinutes: 10,
    NumWorkers:             1,
}, retries.Timeout[clusters.ClusterInfo](10*time.Minute))
```

### Command execution on clusters

**Stability:** [Experimental](https://docs.databricks.com/release-notes/release-types.html)

You can run Python, Scala, R, or SQL code on running interactive Databricks clusters and get the results back. All supplied code gets leading whitespace removed, so that you could easily embed Python code into Go applications. This high-level wrapper comes from the Databricks Terraform provider, where it was tested for over 2 years for use cases such as [DBFS mounts](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mount) and [SQL permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_permissions). This interface hides the intricate complexity of all internal APIs involved to simplify the unit-testing experience for command execution. Databricks does not recommending that you use lower-level interfaces for command execution. The execution timeout is 20 minutes and cannot be overriden for the sake of interface simplicity, meaning that you should only use this API if you have some relatively complex executions to perform. Please use jobs in case your commands must run longer than 20 minutes. Or use the [Databricks SQL Driver for Go](https://github.com/databricks/databricks-sql-go) in case your workload type is purely for business intelligence.

```go
res := w.CommandExecutor.Execute(ctx, clusterId, "python", "print(1)")
if res.Failed() {
    return fmt.Errorf("command failed: %w", res.Err())
}
println(res.Text())
// Out: 1
```

### Cluster library management

**Stability:** [Beta](https://docs.databricks.com/release-notes/release-types.html)

You can install or uninstall libraries on running Databricks clusters. `UpdateAndWait` follows all conventions of [long-running operations](#long-running-operations) and wraps `Install` and `Uninstall` operations, followed by checking for the installation status of the cluster, exposing error messages back in a simplified way. This high-level wrapper came from the Databricks Terraform provider, where it was tested for over 2 years in the [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) and [databricks_library](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/library) resources. Databricks recommends that you use `UpdateAndWait` as the only API for cluster library management.

```go
err = w.Libraries.UpdateAndWait(ctx, libraries.Update{
    ClusterId: clusterId,
    Install: []libraries.Library{
        {
            Pypi: &libraries.PythonPyPiLibrary{
                Package: "dbl-tempo",
            },
        },
    },
})
```

### Advanced usage

**Stability:** _Experimental_

You can track the intermediate state of a long-running operation while waiting to reach the correct state by supplying the `func(i *retries.Info[Zzz])` functional option, where `Zzz` is the return type of the `XxxAndWait` method:

```go
clusterInfo, err = w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
    // ...
}, func(i *retries.Info[clusters.ClusterInfo]) {
    updateIntermediateState(i.Info.StateMessage)
})
```

## Paginated responses

**Stability:** [Experimental](https://docs.databricks.com/release-notes/release-types.html)

On the platform side, some Databricks APIs have result pagination, and some of them do not. Some APIs follow the offset-plus-limit pagination, some start their offsets from 0 and some from 1, some use the cursor-based iteration, and others just return all results in a single response. The Databricks SDK for Go hides this intricate complexity and generates a more high-level interface for retrieving all results of a certain entity type. The naming pattern is `XxxAll`, where `Xxx` is the name of the method to retrieve a single page of results.

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

**Stability:** [Beta](https://docs.databricks.com/release-notes/release-types.html)

The Databricks SDK for Go provides selector methods that make developing multi-cloud applications easier and just rely on characteristics of the virtual machine, such as the number of cores or availability of local disks or always picking up the latest Databricks Runtime for the interactive cluster or per-job cluster.

```go
// Fetch the list of spark runtime versions.
sparkVersions, err := w.Clusters.SparkVersions(ctx)
if err != nil {
    return err
}

// Select the latest LTS version.
latestLTS, err := sparkVersions.Select(clusters.SparkVersionRequest{
    Latest:          true,
    LongTermSupport: true,
})
if err != nil {
    return err
}

// Fetch the list of available node types.
nodeTypes, err := w.Clusters.ListNodeTypes(ctx)
if err != nil {
    return err
}

// Select the smallest node type ID.
smallestWithDisk, err := nodeTypes.Smallest(clusters.NodeTypeRequest{
    LocalDisk: true,
})
if err != nil {
    return err
}

// Create the cluster and wait for it to start properly.
runningCluster, err := w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
    ClusterName:            clusterName,
    SparkVersion:           latestLTS,
    NodeTypeId:             smallestWithDisk,
    AutoterminationMinutes: 15,
    NumWorkers:             1,
})
```

## `io.Reader` integration for DBFS

**Stability:** [Beta](https://docs.databricks.com/release-notes/release-types.html)

Use the higher-level `w.Dbfs.Open` and `w.Dbfs.Overwrite` methods to work with remote files through the `io.Reader` interface. Internally, these methods wrap the low-level intricacies of working with Databricks REST APIs, providing a convenient interface to you as a developer.

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
