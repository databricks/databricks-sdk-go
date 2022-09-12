// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package libraries

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewLibraries(client *client.DatabricksClient) LibrariesService {
	return &LibrariesAPI{
		client: client,
	}
}

type LibrariesAPI struct {
	client *client.DatabricksClient
}

// Get the status of all libraries on all clusters. A status will be available
// for all libraries installed on this cluster via the API or the libraries UI
// as well as libraries set to be installed on all clusters via the libraries
// UI. If a library has been set to be installed on all clusters,
// “is_library_for_all_clusters“ will be “true“, even if the library was
// also installed on this specific cluster.. An example response: .. code:: {
// "statuses": [ { "cluster_id": "11203-my-cluster", "library_statuses": [ {
// "library": { "jar": "dbfs:/mnt/libraries/library.jar" }, "status":
// "INSTALLING", "messages": [], "is_library_for_all_clusters": false } ] }, {
// "cluster_id": "20131-my-other-cluster", "library_statuses": [ { "library": {
// "egg": "dbfs:/mnt/libraries/library.egg" }, "status": "ERROR", "messages":
// ["Could not download library"], "is_library_for_all_clusters": false } ] } ]
// }
func (a *LibrariesAPI) AllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error) {
	var listAllClusterLibraryStatusesResponse ListAllClusterLibraryStatusesResponse
	path := "/api/2.0/libraries/all-cluster-statuses"
	err := a.client.Get(ctx, path, nil, &listAllClusterLibraryStatusesResponse)
	return &listAllClusterLibraryStatusesResponse, err
}

// Get the status of libraries on a cluster. A status will be available for all
// libraries installed on this cluster via the API or the libraries UI as well
// as libraries set to be installed on all clusters via the libraries UI. If a
// library has been set to be installed on all clusters,
// “is_library_for_all_clusters“ will be “true“, even if the library was was
// also installed on this specific cluster. The order of returned libraries will
// be as follows. 1) Libraries set to be installed on this cluster will be
// returned first. Within this group, the order will be order in which the
// libraries were added to the cluster. 2) Libraries set to be installed on all
// clusters are returned next. Within this group there is no order guarantee. 3)
// Libraries that were previously requested on this cluster or on all clusters,
// but now marked for removal. Within this group there is no order guarantee. An
// example request: .. code::
// /libraries/cluster-status?cluster_id=11203-my-cluster And response: .. code::
// { "cluster_id": "11203-my-cluster", "library_statuses": [ { "library": {
// "jar": "dbfs:/mnt/libraries/library.jar" }, "status": "INSTALLED",
// "messages": [], "is_library_for_all_clusters": false }, { "library": {
// "pypi": { "package": "beautifulsoup4" }, }, "status": "INSTALLING",
// "messages": ["Successfully resolved package from PyPI"],
// "is_library_for_all_clusters": false }, { "library": { "cran": { "package":
// "ada", "repo": "http://cran.us.r-project.org" }, }, "status": "FAILED",
// "messages": ["R package installation is not supported on this spark
// version.\nPlease upgrade to Runtime 3.2 or higher"],
// "is_library_for_all_clusters": false } ] }
func (a *LibrariesAPI) ClusterStatus(ctx context.Context, request ClusterStatusRequest) (*ClusterStatusResponse, error) {
	var clusterStatusResponse ClusterStatusResponse
	path := "/api/2.0/libraries/cluster-status"
	err := a.client.Get(ctx, path, request, &clusterStatusResponse)
	return &clusterStatusResponse, err
}

// Get the status of libraries on a cluster. A status will be available for all
// libraries installed on this cluster via the API or the libraries UI as well
// as libraries set to be installed on all clusters via the libraries UI. If a
// library has been set to be installed on all clusters,
// “is_library_for_all_clusters“ will be “true“, even if the library was was
// also installed on this specific cluster. The order of returned libraries will
// be as follows. 1) Libraries set to be installed on this cluster will be
// returned first. Within this group, the order will be order in which the
// libraries were added to the cluster. 2) Libraries set to be installed on all
// clusters are returned next. Within this group there is no order guarantee. 3)
// Libraries that were previously requested on this cluster or on all clusters,
// but now marked for removal. Within this group there is no order guarantee. An
// example request: .. code::
// /libraries/cluster-status?cluster_id=11203-my-cluster And response: .. code::
// { "cluster_id": "11203-my-cluster", "library_statuses": [ { "library": {
// "jar": "dbfs:/mnt/libraries/library.jar" }, "status": "INSTALLED",
// "messages": [], "is_library_for_all_clusters": false }, { "library": {
// "pypi": { "package": "beautifulsoup4" }, }, "status": "INSTALLING",
// "messages": ["Successfully resolved package from PyPI"],
// "is_library_for_all_clusters": false }, { "library": { "cran": { "package":
// "ada", "repo": "http://cran.us.r-project.org" }, }, "status": "FAILED",
// "messages": ["R package installation is not supported on this spark
// version.\nPlease upgrade to Runtime 3.2 or higher"],
// "is_library_for_all_clusters": false } ] }
func (a *LibrariesAPI) ClusterStatusByClusterId(ctx context.Context, clusterId string) (*ClusterStatusResponse, error) {
	return a.ClusterStatus(ctx, ClusterStatusRequest{
		ClusterId: clusterId,
	})
}

// Add libraries to be installed on a cluster. The installation is asynchronous
// - it happens in the background after the completion of this request. Note
// that the actual set of libraries to be installed on a cluster is the union of
// the libraries specified via this method and the libraries set to be installed
// on all clusters via the libraries UI. Note that CRAN libraries can only be
// installed on clusters running Databricks Runtime 3.2 or higher. An example
// request: .. code:: { "cluster_id": "10201-my-cluster", "libraries": [ {
// "jar": "dbfs:/mnt/libraries/library.jar" }, { "egg":
// "dbfs:/mnt/libraries/library.egg" }, { "whl":
// "dbfs:/mnt/libraries/library.whl" }, { "maven": { "coordinates":
// "org.jsoup:jsoup:1.7.2", "exclusions": ["slf4j:slf4j"] } }, { "pypi": {
// "package": "simplejson", "repo": "http://my-pypi-mirror.com" } }, { "cran": {
// "package: "ada", "repo": "http://cran.us.r-project.org" } } ] }
func (a *LibrariesAPI) Install(ctx context.Context, request InstallLibrariesRequest) error {
	path := "/api/2.0/libraries/install"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Set libraries to be uninstalled on a cluster. The libraries won't be
// uninstalled until the cluster is restarted. Uninstalling libraries that are
// not installed on the cluster will have no impact but is not an error. An
// example request: .. code:: { "cluster_id": "10201-my-cluster", "libraries": [
// { "jar": "dbfs:/mnt/libraries/library.jar" }, { "cran": "ada" } ] }
func (a *LibrariesAPI) Uninstall(ctx context.Context, request UninstallLibrariesRequest) error {
	path := "/api/2.0/libraries/uninstall"
	err := a.client.Post(ctx, path, request, nil)
	return err
}
