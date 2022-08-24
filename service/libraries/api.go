// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package libraries

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)


type LibrariesService interface {
    // Get the status of libraries on a cluster. A status will be available for 
    // all libraries installed on this cluster via the API or the libraries UI 
    // as well as libraries set to be installed on all clusters via the 
    // libraries UI. If a library has been set to be installed on all clusters, 
    // ``is_library_for_all_clusters`` will be ``true``, even if the library 
    // was was also installed on this specific cluster. The order of returned 
    // libraries will be as follows. 1) Libraries set to be installed on this 
    // cluster will be returned first. Within this group, the order will be 
    // order in which the libraries were added to the cluster. 2) Libraries set 
    // to be installed on all clusters are returned next. Within this group 
    // there is no order guarantee. 3) Libraries that were previously requested 
    // on this cluster or on all clusters, but now marked for removal. Within 
    // this group there is no order guarantee. An example request: .. code:: 
    // /libraries/cluster-status?cluster_id=11203-my-cluster And response: .. 
    // code:: { &#34;cluster_id&#34;: &#34;11203-my-cluster&#34;, &#34;library_statuses&#34;: [ { 
    // &#34;library&#34;: { &#34;jar&#34;: &#34;dbfs:/mnt/libraries/library.jar&#34; }, &#34;status&#34;: 
    // &#34;INSTALLED&#34;, &#34;messages&#34;: [], &#34;is_library_for_all_clusters&#34;: false }, { 
    // &#34;library&#34;: { &#34;pypi&#34;: { &#34;package&#34;: &#34;beautifulsoup4&#34; }, }, &#34;status&#34;: 
    // &#34;INSTALLING&#34;, &#34;messages&#34;: [&#34;Successfully resolved package from PyPI&#34;], 
    // &#34;is_library_for_all_clusters&#34;: false }, { &#34;library&#34;: { &#34;cran&#34;: { 
    // &#34;package&#34;: &#34;ada&#34;, &#34;repo&#34;: &#34;http://cran.us.r-project.org&#34; }, }, &#34;status&#34;: 
    // &#34;FAILED&#34;, &#34;messages&#34;: [&#34;R package installation is not supported on this 
    // spark version.\nPlease upgrade to Runtime 3.2 or higher&#34;], 
    // &#34;is_library_for_all_clusters&#34;: false } ] } 
    ClusterStatus(ctx context.Context, clusterStatusRequest ClusterStatusRequest) (*ClusterStatusResponse, error)
    // Add libraries to be installed on a cluster. The installation is 
    // asynchronous - it happens in the background after the completion of this 
    // request. Note that the actual set of libraries to be installed on a 
    // cluster is the union of the libraries specified via this method and the 
    // libraries set to be installed on all clusters via the libraries UI. Note 
    // that CRAN libraries can only be installed on clusters running Databricks 
    // Runtime 3.2 or higher. An example request: .. code:: { &#34;cluster_id&#34;: 
    // &#34;10201-my-cluster&#34;, &#34;libraries&#34;: [ { &#34;jar&#34;: 
    // &#34;dbfs:/mnt/libraries/library.jar&#34; }, { &#34;egg&#34;: 
    // &#34;dbfs:/mnt/libraries/library.egg&#34; }, { &#34;whl&#34;: 
    // &#34;dbfs:/mnt/libraries/library.whl&#34; }, { &#34;maven&#34;: { &#34;coordinates&#34;: 
    // &#34;org.jsoup:jsoup:1.7.2&#34;, &#34;exclusions&#34;: [&#34;slf4j:slf4j&#34;] } }, { &#34;pypi&#34;: { 
    // &#34;package&#34;: &#34;simplejson&#34;, &#34;repo&#34;: &#34;http://my-pypi-mirror.com&#34; } }, { 
    // &#34;cran&#34;: { &#34;package: &#34;ada&#34;, &#34;repo&#34;: &#34;http://cran.us.r-project.org&#34; } } ] 
    // } 
    InstallLibraries(ctx context.Context, installLibrariesRequest InstallLibrariesRequest) error
    // Get the status of all libraries on all clusters. A status will be 
    // available for all libraries installed on this cluster via the API or the 
    // libraries UI as well as libraries set to be installed on all clusters 
    // via the libraries UI. If a library has been set to be installed on all 
    // clusters, ``is_library_for_all_clusters`` will be ``true``, even if the 
    // library was also installed on this specific cluster.. An example 
    // response: .. code:: { &#34;statuses&#34;: [ { &#34;cluster_id&#34;: &#34;11203-my-cluster&#34;, 
    // &#34;library_statuses&#34;: [ { &#34;library&#34;: { &#34;jar&#34;: 
    // &#34;dbfs:/mnt/libraries/library.jar&#34; }, &#34;status&#34;: &#34;INSTALLING&#34;, &#34;messages&#34;: 
    // [], &#34;is_library_for_all_clusters&#34;: false } ] }, { &#34;cluster_id&#34;: 
    // &#34;20131-my-other-cluster&#34;, &#34;library_statuses&#34;: [ { &#34;library&#34;: { &#34;egg&#34;: 
    // &#34;dbfs:/mnt/libraries/library.egg&#34; }, &#34;status&#34;: &#34;ERROR&#34;, &#34;messages&#34;: 
    // [&#34;Could not download library&#34;], &#34;is_library_for_all_clusters&#34;: false } ] 
    // } ] } 
    ListAllClusterLibraryStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error)
    // Set libraries to be uninstalled on a cluster. The libraries won&#39;t be 
    // uninstalled until the cluster is restarted. Uninstalling libraries that 
    // are not installed on the cluster will have no impact but is not an 
    // error. An example request: .. code:: { &#34;cluster_id&#34;: &#34;10201-my-cluster&#34;, 
    // &#34;libraries&#34;: [ { &#34;jar&#34;: &#34;dbfs:/mnt/libraries/library.jar&#34; }, { &#34;cran&#34;: 
    // &#34;ada&#34; } ] } 
    UninstallLibraries(ctx context.Context, uninstallLibrariesRequest UninstallLibrariesRequest) error
}

func New(client *client.DatabricksClient) LibrariesService {
	return &LibrariesAPI{
		client: client,
	}
}

type LibrariesAPI struct {
	client *client.DatabricksClient
}

// Get the status of libraries on a cluster. A status will be available for all 
// libraries installed on this cluster via the API or the libraries UI as well 
// as libraries set to be installed on all clusters via the libraries UI. If a 
// library has been set to be installed on all clusters, 
// ``is_library_for_all_clusters`` will be ``true``, even if the library was 
// was also installed on this specific cluster. The order of returned libraries 
// will be as follows. 1) Libraries set to be installed on this cluster will be 
// returned first. Within this group, the order will be order in which the 
// libraries were added to the cluster. 2) Libraries set to be installed on all 
// clusters are returned next. Within this group there is no order guarantee. 
// 3) Libraries that were previously requested on this cluster or on all 
// clusters, but now marked for removal. Within this group there is no order 
// guarantee. An example request: .. code:: 
// /libraries/cluster-status?cluster_id=11203-my-cluster And response: .. 
// code:: { &#34;cluster_id&#34;: &#34;11203-my-cluster&#34;, &#34;library_statuses&#34;: [ { 
// &#34;library&#34;: { &#34;jar&#34;: &#34;dbfs:/mnt/libraries/library.jar&#34; }, &#34;status&#34;: 
// &#34;INSTALLED&#34;, &#34;messages&#34;: [], &#34;is_library_for_all_clusters&#34;: false }, { 
// &#34;library&#34;: { &#34;pypi&#34;: { &#34;package&#34;: &#34;beautifulsoup4&#34; }, }, &#34;status&#34;: 
// &#34;INSTALLING&#34;, &#34;messages&#34;: [&#34;Successfully resolved package from PyPI&#34;], 
// &#34;is_library_for_all_clusters&#34;: false }, { &#34;library&#34;: { &#34;cran&#34;: { &#34;package&#34;: 
// &#34;ada&#34;, &#34;repo&#34;: &#34;http://cran.us.r-project.org&#34; }, }, &#34;status&#34;: &#34;FAILED&#34;, 
// &#34;messages&#34;: [&#34;R package installation is not supported on this spark 
// version.\nPlease upgrade to Runtime 3.2 or higher&#34;], 
// &#34;is_library_for_all_clusters&#34;: false } ] } 
func (a *LibrariesAPI) ClusterStatus(ctx context.Context, request ClusterStatusRequest) (*ClusterStatusResponse, error) {
	var clusterStatusResponse ClusterStatusResponse
	path := "/api/2.0/libraries/cluster-status"
	err := a.client.Get(ctx, path, request, &clusterStatusResponse)
	return &clusterStatusResponse, err
}

// Add libraries to be installed on a cluster. The installation is asynchronous 
// - it happens in the background after the completion of this request. Note 
// that the actual set of libraries to be installed on a cluster is the union 
// of the libraries specified via this method and the libraries set to be 
// installed on all clusters via the libraries UI. Note that CRAN libraries can 
// only be installed on clusters running Databricks Runtime 3.2 or higher. An 
// example request: .. code:: { &#34;cluster_id&#34;: &#34;10201-my-cluster&#34;, &#34;libraries&#34;: 
// [ { &#34;jar&#34;: &#34;dbfs:/mnt/libraries/library.jar&#34; }, { &#34;egg&#34;: 
// &#34;dbfs:/mnt/libraries/library.egg&#34; }, { &#34;whl&#34;: 
// &#34;dbfs:/mnt/libraries/library.whl&#34; }, { &#34;maven&#34;: { &#34;coordinates&#34;: 
// &#34;org.jsoup:jsoup:1.7.2&#34;, &#34;exclusions&#34;: [&#34;slf4j:slf4j&#34;] } }, { &#34;pypi&#34;: { 
// &#34;package&#34;: &#34;simplejson&#34;, &#34;repo&#34;: &#34;http://my-pypi-mirror.com&#34; } }, { &#34;cran&#34;: 
// { &#34;package: &#34;ada&#34;, &#34;repo&#34;: &#34;http://cran.us.r-project.org&#34; } } ] } 
func (a *LibrariesAPI) InstallLibraries(ctx context.Context, request InstallLibrariesRequest) error {
	path := "/api/2.0/libraries/install"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Get the status of all libraries on all clusters. A status will be available 
// for all libraries installed on this cluster via the API or the libraries UI 
// as well as libraries set to be installed on all clusters via the libraries 
// UI. If a library has been set to be installed on all clusters, 
// ``is_library_for_all_clusters`` will be ``true``, even if the library was 
// also installed on this specific cluster.. An example response: .. code:: { 
// &#34;statuses&#34;: [ { &#34;cluster_id&#34;: &#34;11203-my-cluster&#34;, &#34;library_statuses&#34;: [ { 
// &#34;library&#34;: { &#34;jar&#34;: &#34;dbfs:/mnt/libraries/library.jar&#34; }, &#34;status&#34;: 
// &#34;INSTALLING&#34;, &#34;messages&#34;: [], &#34;is_library_for_all_clusters&#34;: false } ] }, { 
// &#34;cluster_id&#34;: &#34;20131-my-other-cluster&#34;, &#34;library_statuses&#34;: [ { &#34;library&#34;: { 
// &#34;egg&#34;: &#34;dbfs:/mnt/libraries/library.egg&#34; }, &#34;status&#34;: &#34;ERROR&#34;, &#34;messages&#34;: 
// [&#34;Could not download library&#34;], &#34;is_library_for_all_clusters&#34;: false } ] } ] 
// } 
func (a *LibrariesAPI) ListAllClusterLibraryStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error) {
	var listAllClusterLibraryStatusesResponse ListAllClusterLibraryStatusesResponse
	path := "/api/2.0/libraries/all-cluster-statuses"
	err := a.client.Get(ctx, path, nil, &listAllClusterLibraryStatusesResponse)
	return &listAllClusterLibraryStatusesResponse, err
}

// Set libraries to be uninstalled on a cluster. The libraries won&#39;t be 
// uninstalled until the cluster is restarted. Uninstalling libraries that are 
// not installed on the cluster will have no impact but is not an error. An 
// example request: .. code:: { &#34;cluster_id&#34;: &#34;10201-my-cluster&#34;, &#34;libraries&#34;: 
// [ { &#34;jar&#34;: &#34;dbfs:/mnt/libraries/library.jar&#34; }, { &#34;cran&#34;: &#34;ada&#34; } ] } 
func (a *LibrariesAPI) UninstallLibraries(ctx context.Context, request UninstallLibrariesRequest) error {
	path := "/api/2.0/libraries/uninstall"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

