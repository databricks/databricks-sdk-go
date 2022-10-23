Databricks SDK for Go
---

# Features

Databricks SDK for Go comes with a number of features to accelerate the development on top of the Databricks Lakehouse.

## Zero conf with unified auth

If you use conventional environment variables, the only thing required to start working with Databricks Workspaces is the following snippet:

```go
w := workspaces.New()
w./*press TAB for autocompletion*/
```

This means, that if you run [Databricks Terraform Provider](https://registry.terraform.io/providers/databrickslabs/databricks/latest), CLI, or applications build with Databricks SDK for other langauges, most likely they'll all interoperate nicely together. The conventional name for the variable that holds workspace-level client of Databricks SDK for Go is `w`, the shorthand for `workspace`.

## Long-running operations

**Stability:** _Experimental_

More than 20 methods across different Databricks APIs are the long-running operations for managing things like clusters, command execution, jobs, libraries, delta pipelines, and SQL warehouses. Let's take the example of Clusters API: once you create a cluster, you receive a cluster ID, the cluster is in the `PENDING` state, while Databricks takes care of provisioning virtual machines from the cloud provider in the background. But the cluster is only usable in `RUNNING` state. The other example is the API for running a job or repairing the run: right after the start, the run is in `PENDING` state, though the job is considered to be finished only when it's `TERMINATED` or `SKIPPED`. And of course you'd want to know the error message, when the long-running operation times out or why things fail. And sometimes you want to configure a custom timeout, other than the default 20 minutes. 

To hide all of the integration specific complexity from the end user, Databricks SDK for GoLang provides a high-level API for _triggering_ the long-running operations and _waiting_ for the releated entities to reach the right state or return back the error message about the problem in case of failure. All long-running operations have the `XxxAndWait` name pattern, where `Xxx` is the operation name. All these generated methods return information about the relevant entity once the operation is finished. It's possible to configure a custom timeout to `XxxAndWait` by providing a functional option argument constructed by `XxxTimeout(time.Duration)` package-level function.

In the following example, `CreateAndWait` will return `ClusterInfo` only once the cluster is in the `RUNNING` state, otherwise it'll timeout in 10 minutes:

```go
clusterInfo, err = w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
    ClusterName:            "Created cluster",
    SparkVersion:           latestLTS,
    NodeTypeId:             smallestWithDisk,
    AutoterminationMinutes: 10,
    NumWorkers:             1,
}, clusters.CreateTimeout(10*time.Minute))
```

### Command execution on clusters

**Stability:** _Experimental_

You can run Python, Scala, R, or SQL code on running interactive Databricks clusters and get the results back. All supplied code gets leading whitespace removed, so that you could easily embed Python code into Go applications. This high-level wrapper came from the Databricks Terraform provider, where it was tested for 2+ years for use-cases like [DBFS mounts](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/mount) or [sql permissions](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/sql_permissions). This interface hides the intricate complexity of all internal APIs involved to simplify the unit-testing experience for command execution. We're not recommending to use lower-level interfaces for command execution. Execution timeout is 20 minutes and canno be overriden for the sake of interface simplicity, meaning that you should only use this API if you have some relatively executions to perform. Please use jobs in case your commands have to run longer than 20 minutes. Or [database/sql driver](https://github.com/databricks/databricks-sql-go) in case your workload type is purely for business intelligence.

```go
res := w.CommandExecutor.Execute(ctx, clusterId, "python", "print(1)")
if res.Failed() {
    return fmt.Errorf("command failed: %w", res.Err())
}
println(res.Text())
// Out: 1
```

### Cluster library management

**Stability:** _Beta_

You can install or uninstall libraries on running Databricks clusters. `UpdateAndWait` follows all conventions of [long-running operations](#long-running-operations) and wraps `Install` and `Uninstall` operations, followed by checking for installation status of the cluster, exposing error messages back in a simplified way. This high-level wrapper came from the Databricks Terraform provider, where it was tested for 2+ years in [databricks_cluster](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/cluster) and [databricks_library](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/library) resources. We recommend to use `UpdateAndWait` as the only API to for cluster library management.

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

You can track the intermediate state of the long-running operation while waiting to reach the correct state by supplying `func(i *retries.Info[Zzz])` functional option, where `Zzz` is the return type of `XxxAndWait` method:

```go
clusterInfo, err = w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
    // ...
}, func(i *retries.Info[clusters.ClusterInfo]) {
    updateIntermediateState(i.Info.StateMessage)
})
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
