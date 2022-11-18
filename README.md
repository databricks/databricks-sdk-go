# Databricks SDK for Go (Beta)

**Stability**: [Beta](https://docs.databricks.com/release-notes/release-types.html) (unless otherwise noted in the following sections)

The Databricks SDK for Go includes functionality to accelerate development with [Go](https://go.dev) for the Databricks Lakehouse. It covers all public [Databricks REST API](https://docs.databricks.com/dev-tools/api/index.html) operations. The SDK's internal HTTP client is robust and handles failures on different levels by performing intelligent retries.

## Contents

- [Getting started](#getting-started)
- [Authentication](#authentication)
- [Code examples](#code-examples)
- [Long running operations](#long-running-operations)
- [Paginated responses](#paginated-responses)
- [Node type and Databricks Runtime selectors](#node-type-and-databricks-runtime-selectors)
- [io.Reader integration for DBFS](#io-reader-integration-for-dbfs)

<a id="getting-started"/>

## Getting started

During the Beta period, you must clone and then reference this repository locally, as follows. After the Beta period, you will be able to download and reference the Databricks SDK for Go package by using familiar commands such as `go get` and `go install`.

1. On your local development machine with Go already [installed](https://go.dev/doc/install) and a Go code [project](https://go.dev/doc/tutorial/getting-started) active, clone this repository into your existing Go code project's current working directory, by running the `git clone` command:

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

5. Within your project, create a Go code file that imports the Databricks SDK for Go. The following example, in a file named `main.go` with the following contents, simply reports the object type that represents the root directory in your Databricks workspace (in this case, `DIRECTORY`):

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

6. Add any misssing module dependencies by running the `go mod tidy` command:

   ```bash
   go mod tidy
   ```
   
   **Note**: If you get the error `go: warning: "all" matched no packages`, you forgot to add the preceding Go code file that imports the Databricks SDK for Go. 
   
7. Grab copies of all packages needed to support builds and tests of packages in your `main` module, by running the `go mod vendor` command:

   ```bash
   go mod vendor
   ```

8. Set up Databricks authentication on your local development machine, if you have not done so already. For details, see the next section, "Authentication."
9. Run your Go code file, assuming a file named `main.go`, by running the `go run` command:

   ```bash
   go run main.go
   ```
   
   Assuming the preceding example code is run, the output is: 
   
   ```bash
   The path '/' is a 'DIRECTORY'.   
   ```

<a id="authentication"/>

## Authentication

If you use Databricks [configuration profiles](https://docs.databricks.com/dev-tools/auth.html#configuration-profiles) or Databricks-specific [environment variables](https://docs.databricks.com/dev-tools/auth.html#environment-variables) for [Databricks authentication](https://docs.databricks.com/dev-tools/auth.html), the only code required to start working with a Databricks workspace is the following code snippet, which instructs the Databricks SDK for Go to use its [default authentication flow](#default-authentication-flow):

```go
w := workspaces.New()
w./*press TAB for autocompletion*/
```

The conventional name for the variable that holds the workspace-level client of the Databricks SDK for Go is `w`, which is shorthand for `workspace`.

### In this section

- [Default authentication flow](#default-authentication-flow)
- [Databricks native authentication](#databricks-native-authentication)
- [Azure native authentication](#azure-native-authentication)
- [Google Cloud Platform native authentication](#google-cloud-platform-native-authentication)
- [Overriding .databrickscfg](#overriding-databrickscfg)
- [Additional authentication configuration options](#additional-authentication-configuration-options)
- [Custom credentials provider](#custom-credentials-provider)

<a id="default-authentication-flow"/>

### Default authentication flow

If you run the [Databricks Terraform Provider](https://registry.terraform.io/providers/databrickslabs/databricks/latest), the [Databricks CLI](https://docs.databricks.com/dev-tools/cli/index.html), or applications that target the Databricks SDKs for other langauges, most likely they will all interoperate nicely together. By default, the Databricks SDK for Go checks the following to perform [Databricks authentication](https://docs.databricks.com/dev-tools/auth.html), in the following order for each of its supported Databricks authentication methods:

1. Read any direct fields that are hard-coded into `*databricks.Config`.
2. Read locally-set Databricks-specific [environment variables](https://docs.databricks.com/dev-tools/auth.html#environment-variables) for authentication into `*databricks.Config`, as applicable for the current Databricks authentication method.
3. As applicable for the current Databricks authentication method, reads the contents of the local `.databrickscfg` file's `DEFAULT` [configuration profile](https://docs.databricks.com/dev-tools/auth.html#configuration-profiles) from its default location (`~` for Linux or macOS, and `%USERPROFILE%` for Windows) into `*databricks.Config`.
4. If any direct fields are NOT hard-coded in `databricks.Config`, AND Databricks-specific environment variables exist AND a `.databrickscfg` file exists in the default location with a `DEFAULT` profile, then the Databricks SDK for Go stops running and returns the error `panic: validate: more than one authorization method configured`. To fix this error, either remove the duplicated Databricks-specific environment variables or remove the `DEFAULT` profile from the `.databrickscfg` file.
5. Otherwise, the Databricks SDK for Go performs Databricks authentication by using the information in `*databricks.Config`. As applicable for the current Databricks authentication method, the SDK attempts to get related authentication information from the Azure CLI or Google Cloud CLI.

The Databricks SDK for Go uses the following Databricks authentication methods in the following order to attempt to authenticate with your Databricks account, workspace, or both. Once the SDK successfully authenticates, it skips the remaining methods:

1. [Databricks native authentication](#databricks-native-authentication), regardless of cloud provider. If this fails, the SDK will move on to attempt Azure native authentication.
2. [Azure native authentication](#azure-native-authentication). For Google Cloud Platform accounts and workspaces, the SDK will still attempt Azure native authentication; however, these attempts will of course fail, and the SDK will then move on to attempt Google Cloud Platform native authentication.
3. [Google Cloud Platform native authentication](#google-cloud-platform-native-authentication). For Azure accounts and workspaces, if step 2 fails, the SDK will still attempt Google Cloud Platform native authentication, but these attempts will of course fail as well.
4. If the SDK is still unsuccessful, it will return an error and the code stops running.

You can instruct the Databricks SDK for Go to move directly to a specific Databricks authentication method by setting the `AuthType` field in `*databricks.Config` as described in the following sections.

Depending on the Databricks authentication method, the Databricks SDK for Go uses the following information. Presented are the `*databricks.Config` fields, their descriptions, any corresponding environment variables, and any corresponding `.databrickscfg` file fields, respectively.

<a id="databricks-native-authentication"/>

### Databricks native authentication

By default, the Databricks SDK for Go initially tries Databricks personal access token (PAT) authentication (`AuthType: "pat"` in `*databricks.Config`). If the SDK is unsuccessful, it then tries Databricks basic (username/password) authentication (`AuthType: "basic"` in `*databricks.Config`).

 * `host` _(string)_: The Databricks host URL for either the Databricks workspace endpoint or the Databricks accounts endpoint. Environment: `DATABRICKS_HOST`. `.databrickscfg` file: `host`.
 * `account_id` _(string)_: The Databricks account ID for the Databricks accounts endpoint. Only has effect when `host` is either `https://accounts.cloud.databricks.com/` _(AWS)_, `https://accounts.azuredatabricks.net/` _(Azure)_, or `https://accounts.gcp.databricks.com/` _(GCP)_.  Environment: `DATABRICKS_ACCOUNT_ID`. `.databrickscfg` file: `host`.
 * `token` _(string)_: The Databricks personal access token (PAT) or Azure Active Directory (Azure AD) token. Environment: `DATABRICKS_TOKEN`. `.databrickscfg` file: `token`.
 * `username` _(string)_: The Databricks username part of basic authentication. Only possible when `host` is `*.cloud.databricks.com` _(AWS)_. Environment: `DATABRICKS_USERNAME`. `.databrickscfg` file: `username`.
 * `password` _(string)_: The Databricks password part of basic authentication. Only possible when `host` is `*.cloud.databricks.com` _(AWS)_. Environment: `DATABRICKS_PASSWORD`. `.databrickscfg` file: `password`.

For example, to print Databricks PAT authentication settings:

```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"

  "github.com/databricks/databricks-sdk-go/workspaces"
  // "github.com/databricks/databricks-sdk-go/databricks"
)

func main() {
  /*
   * Perform Databricks token authentication for a Databricks workspace.
   *
   * Choose from one of the following authentication options:
   */

  /*
   * Option 1: To use Databricks token authentication by default, uncomment
   * the following line of code and then run. This assumes you have already set
   * your Databricks workspace URL and token through one of:
   *
   * - A Databricks configuration profile named DEFAULT.
   * - The environment variables DATABRICKS_HOST and DATABRICKS_TOKEN.
   */

  // w := workspaces.New()

  /* 
   * Option 2: To ask the user at run time for the needed information,
   * uncomment the following code, uncomment the import
   * "github.com/databricks/databricks-sdk-go/databricks"
   * and then run.
   */

  // w := workspaces.New(&databricks.Config{
  //   AuthType: "pat",
  //   Host:     askFor("Databricks workspace URL:"),
  //   Token:    askFor("Access token:"),
  // })

  fmt.Printf("The workspace URL is '%s', and the token is '%s'.\n",
    w.Config.Host,
    w.Config.Token,
  )

  // Now call the Databricks workspace APIs as desired...
}

// For Option 2, ask the user for the missing information.
func askFor(prompt string) string {
  var s string
  r := bufio.NewReader(os.Stdin)
  for {
    fmt.Fprint(os.Stdout, prompt+" ")
    s, _ = r.ReadString('\n')
    if s != "" {
      break
    }
  }
  return strings.TrimSpace(s)
}
```

<a id="azure-native-authentication"/>

### Azure native authentication

By default, the Databricks SDK for Go first tries Azure client secret authentication (`AuthType: "azure-client-secret"` in `*databricks.Config`). If the SDK is unsuccessful, it then tries Azure CLI authentication (`AuthType: "azure-cli"` in `*databricks.Config`). See [Manage service principals](https://learn.microsoft.com/azure/databricks/administration-guide/users-groups/service-principals).

The Databricks SDK for Go picks up an Azure CLI token, if you've previously authenticated as an Azure user by running `az login` on your machine. See [Get Azure AD tokens for users by using the Azure CLI](https://learn.microsoft.com/azure/databricks/dev-tools/api/latest/aad/user-aad-token). If you need to authenticate as Azure Active Directory service principal, you must specify `azure_workspace_resource_id` and `azure_use_msi`; or `azure_tenant_id`, `azure_client_id`, and `azure_client_secret`. See [Add a service principal to your Azure Databricks account](https://learn.microsoft.com/azure/databricks/administration-guide/users-groups/service-principals#add-sp-account). There are no `.databrickscfg` file equivalents for the following values.

 * `azure_workspace_resource_id` _(string)_: The Azure Resource Manager ID for the Azure Databricks workspace, which is exchanged for a Databricks host URL. Environment: `DATABRICKS_AZURE_RESOURCE_ID`.
 * `azure_use_msi` _(string)_: `true` to use Azure Managed Service Identity passwordless authentication flow for service principals. Environment: `ARM_USE_MSI`. _This feature is not yet implemented in the Databricks SDK for Go._
 * `azure_client_secret` _(string)_: The Azure Active Directory service principal's client secret. Environment: `ARM_CLIENT_SECRET`.
 * `azure_client_id` _(string)_: The Azure Active Directory service principal's application ID. Environment: `ARM_CLIENT_ID`.
 * `azure_tenant_id` _(string)_: The Azure Active Directory service principal's tenant ID. Environment: `ARM_TENANT_ID`.
 * `azure_environment` _(string)_: the Azure environment type (such as Public, UsGov, China, and Germany) for a specific set of API endpoints. Defaults to `PUBLIC`. Environment: `ARM_ENVIRONMENT`.

For example, to print Azure client secret authentication settings:

```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"

  "github.com/databricks/databricks-sdk-go/databricks"
  "github.com/databricks/databricks-sdk-go/workspaces"
)

func main() {
  /*
   * Perform Azure client secret authentication for a Databricks workspace.
   * 
   * Choose from one of the following authentication options:
   */

  /*
   * Option 1: Uncomment the following line of code and then run. This
   * assumes you have already set the following environment variables:
   *
   * - DATABRICKS_AZURE_RESOURCE_ID
   * - ARM_TENANT_ID
   * - ARM_CLIENT_ID
   * - ARM_CLIENT_SECRET
   */

  // w := workspaces.New(&databricks.Config{AuthType: "azure-client-secret"})

  /* 
   * Option 2: To ask the user at run time for the needed information,
   * uncomment the following code and then run.
   */

  // w := workspaces.New(&databricks.Config{
  //   AuthType:          "azure-client-secret",
  //   AzureResourceID:   askFor("Azure resource ID for your Azure Databricks workspace:"),
  //   AzureTenantID:     askFor("Azure tenant ID for your Azure AD service principal:"),
  //   AzureClientID:     askFor("Azure client ID for your Azure AD service principal:"),
  //   AzureClientSecret: askFor("Azure client secret for your Azure AD service principal:"),
  // })

  fmt.Printf("The resource ID is '%s', the tenant ID is '%s', "+
    "the client ID is '%s', and the client secret is '%s'.",
    w.Config.AzureResourceID,
    w.Config.AzureTenantID,
    w.Config.AzureClientID,
    w.Config.AzureClientSecret,
  )

  // Now call the Databricks workspace APIs as desired...
}

// For Option 2, ask the user for the missing information.
func askFor(prompt string) string {
  var s string
  r := bufio.NewReader(os.Stdin)
  for {
    fmt.Fprint(os.Stdout, prompt+" ")
    s, _ = r.ReadString('\n')
    if s != "" {
      break
    }
  }
  return strings.TrimSpace(s)
}
```

<a id="google-cloud-platform-native-authentication"/>

### Google Cloud Platform native authentication

By default, the Databricks SDK for Go first tries Google Cloud Platform (GCP) ID authentication (`AuthType: "google-id"` in `*databricks.Config`). If the SDK is unsuccessful, it then tries GCP credentials authentication (`AuthType: "google-credentials"` in `*databricks.Config`).

The Databricks SDK for Go picks up an OAuth token in the scope of the Google Default Application Credentials (DAC) flow. This means that if you have run `gcloud auth application-default login` on your development machine, or launch the application on the compute, that is allowed to impersonate the Google Cloud service account specified in `google_service_account`. Authentication should then work out of the box. See [Creating and managing service accounts](https://cloud.google.com/iam/docs/creating-managing-service-accounts). There are no `.databrickscfg` file equivalents for the following values.

 * `google_service_account` _(string)_: The Google Cloud Platform (GCP) service account e-mail used for impersonation in the Default Application Credentials Flow that does not require a password. Environment: `DATABRICKS_GOOGLE_SERVICE_ACCOUNT`.
 * `google_credentials` _(string)_: GCP Service Account Credentials JSON or the location of these credentials on the local filesustem. Environment: `GOOGLE_CREDENTIALS`.

For example, to print GCP ID authentication settings:

```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"

  "github.com/databricks/databricks-sdk-go/databricks"
  "github.com/databricks/databricks-sdk-go/workspaces"
)

func main() {
  /*
   * Perform Google Cloud Platform ID authentication for a Databricks workspace.
   *
   * Choose from one of the following authentication options:
   */

  /*
   * Option 1: Uncomment the following line of code and then run. This
   * assumes you have already set the environment variables
   * DATABRICKS_HOST and GOOGLE_SERVICE_ACCOUNT.
   */

  // w := workspaces.New(&databricks.Config{AuthType: "google-id"})

  /*
   * Option 2: To ask the user at run time for the needed information,
   * uncomment the following code and then run.
   */

  // w := workspaces.New(&databricks.Config{
  //   AuthType:             "google-id",
  //   Host:                 askFor("Databricks workspace URL:"),
  //   GoogleServiceAccount: askFor("Google Cloud Platform service account's email address:"),
  // })

  fmt.Printf("The host is '%s', and the service account email is '%s'.",
    w.Config.Host,
    w.Config.GoogleServiceAccount,
  )

  // Now call the Databricks workspace APIs as desired...
}

// For Option 2, ask the user for the missing information.
func askFor(prompt string) string {
  var s string
  r := bufio.NewReader(os.Stdin)
  for {
    fmt.Fprint(os.Stdout, prompt+" ")
    s, _ = r.ReadString('\n')
    if s != "" {
      break
    }
  }
  return strings.TrimSpace(s)
}
```

<a id="overriding-databrickscfg"/>

### Overriding `.databrickscfg` 

For [Databricks native authentication](#databricks-native-authentication), you can override the default behavior in `*databricks.Config` for using `.databrickscfg` as follows:

 * `profile` _(string)_: A connection profile specified within `.databrickscfg` to use instead of `DEFAULT`. Environment: `DATABRICKS_CONFIG_PROFILE`.
 * `config_file` _(string)_: A non-default location of the Databricks CLI credentials file. Environment: `DATABRICKS_CONFIG_FILE`.

For example, to use a profile named `MYPROFILE` instead of `DEFAULT`:

```go
package main

import (
  "fmt"

  "github.com/databricks/databricks-sdk-go/databricks"
  "github.com/databricks/databricks-sdk-go/workspaces"
)

func main() {
  w := workspaces.New(&databricks.Config{
    AuthType: "pat",
    Profile:  "MYPROFILE",
  })

  fmt.Printf("Using authorization type '%s', the profile is '%s'.",
    w.Config.AuthType,
    w.Config.Profile,
  )

  // Now call the Databricks workspace APIs as desired...
}
```

<a id="additional-authentication-configuration-options"/>

### Additional authentication configuration options

For all authentication methods, you can override the default behavior in `*databricks.Config` as follows:

 * `auth_type` _(string)_: When multiple auth attributes are available in the environment, use the auth type specified by this argument. This argument also holds the currently selected auth.
 * `http_timeout_seconds` _(int)_: Number of seconds for HTTP timeout. Default is _60_.
 * `retry_timeout_seconds` _(int)_: Number of seconds to keep retrying HTTP requests. Default is _300 (5 minutes)_.
 * `debug_truncate_bytes` _(int)_: Truncate JSON fields in debug logs above this limit. Default is 96. Environment: `DATABRICKS_DEBUG_TRUNCATE_BYTES`.
 * `debug_headers` _(bool)_: `true` to debug HTTP headers of requests made by the application. Default is `false`, as headers contain sensitive data, such as access tokens. Environment: `DATABRICKS_DEBUG_HEADERS`.
 * `rate_limit` _(int)_: Maximum number of requests per second made to Databricks REST API. Environment: `DATABRICKS_RATE_LIMIT`.

For example, to turn on debug HTTP headers:

```go
package main

import (
  "fmt"

  "github.com/databricks/databricks-sdk-go/databricks"
  "github.com/databricks/databricks-sdk-go/workspaces"
)

func main() {
  w := workspaces.New(&databricks.Config{
    DebugHeaders: true,
  })

  fmt.Printf("Debug headers is '%t'.", w.Config.DebugHeaders,)
  
  // Now call the Databricks workspace APIs as desired...
}
```

<a id="custom-credentials-provider"/>

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

<a id="code-examples"/>

## Code examples

To find code examples that demonstrate how to call the Databricks SDK for Go, see the top-level [examples](/examples) folder within this repository.

<a id="long-running-operations"/>

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

### In this section

- [Command execution on clusters](#command-execution-on-clusters)
- [Cluster library management](#cluster-library-management)
- [Advanced usage](#advanced-usage)

<a id="command-execution-on-clusters"/>

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

<a id="cluster-library-management"/>

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

<a id="advanced-usage"/>

### Advanced usage

**Stability:** [Experimental](https://docs.databricks.com/release-notes/release-types.html)

You can track the intermediate state of a long-running operation while waiting to reach the correct state by supplying the `func(i *retries.Info[Zzz])` functional option, where `Zzz` is the return type of the `XxxAndWait` method:

```go
clusterInfo, err = w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
    // ...
}, func(i *retries.Info[clusters.ClusterInfo]) {
    updateIntermediateState(i.Info.StateMessage)
})
```

<a id="paginated-responses"/>

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

<a id="node-type-and-databricks-runtime-selectors"/>

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

<a id="io-reader-integration-for-dbfs"/>

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
