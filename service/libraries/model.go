// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package libraries

// all definitions in this file are in alphabetical order

type ClusterLibraryStatuses struct {
	// Unique identifier for the cluster.
	ClusterId string `json:"cluster_id,omitempty"`
	// Status of all libraries on the cluster.
	LibraryStatuses []LibraryFullStatus `json:"library_statuses,omitempty"`
}

// Get status
type ClusterStatus struct {
	// Unique identifier of the cluster whose status should be retrieved.
	ClusterId string `json:"-" url:"cluster_id,omitempty"`
}

type InstallLibraries struct {
	// Unique identifier for the cluster on which to install these libraries.
	ClusterId string `json:"cluster_id"`
	// The libraries to install.
	Libraries []Library `json:"libraries,omitempty"`
}

type Library struct {
	// Specification of a CRAN library to be installed as part of the library
	Cran *RCranLibrary `json:"cran,omitempty"`
	// URI of the egg to be installed. Currently only DBFS and S3 URIs are
	// supported. For example: `{ "egg": "dbfs:/my/egg" }` or `{ "egg":
	// "s3://my-bucket/egg" }`. If S3 is used, please make sure the cluster has
	// read access on the library. You may need to launch the cluster with an
	// IAM role to access the S3 URI.
	Egg string `json:"egg,omitempty"`
	// URI of the jar to be installed. Currently only DBFS and S3 URIs are
	// supported. For example: `{ "jar": "dbfs:/mnt/databricks/library.jar" }`
	// or `{ "jar": "s3://my-bucket/library.jar" }`. If S3 is used, please make
	// sure the cluster has read access on the library. You may need to launch
	// the cluster with an IAM role to access the S3 URI.
	Jar string `json:"jar,omitempty"`
	// Specification of a maven library to be installed. For example: `{
	// "coordinates": "org.jsoup:jsoup:1.7.2" }`
	Maven *MavenLibrary `json:"maven,omitempty"`
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	Pypi *PythonPyPiLibrary `json:"pypi,omitempty"`
	// URI of the wheel to be installed. For example: `{ "whl": "dbfs:/my/whl"
	// }` or `{ "whl": "s3://my-bucket/whl" }`. If S3 is used, please make sure
	// the cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Whl string `json:"whl,omitempty"`
}

type LibraryFullStatus struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	IsLibraryForAllClusters bool `json:"is_library_for_all_clusters,omitempty"`
	// Unique identifier for the library.
	Library *Library `json:"library,omitempty"`
	// All the info and warning messages that have occurred so far for this
	// library.
	Messages []string `json:"messages,omitempty"`
	// Status of installing the library on the cluster.
	Status LibraryFullStatusStatus `json:"status,omitempty"`
}

// Status of installing the library on the cluster.
type LibraryFullStatusStatus string

const LibraryFullStatusStatusFailed LibraryFullStatusStatus = `FAILED`

const LibraryFullStatusStatusInstalled LibraryFullStatusStatus = `INSTALLED`

const LibraryFullStatusStatusInstalling LibraryFullStatusStatus = `INSTALLING`

const LibraryFullStatusStatusPending LibraryFullStatusStatus = `PENDING`

const LibraryFullStatusStatusResolving LibraryFullStatusStatus = `RESOLVING`

const LibraryFullStatusStatusSkipped LibraryFullStatusStatus = `SKIPPED`

const LibraryFullStatusStatusUninstallOnRestart LibraryFullStatusStatus = `UNINSTALL_ON_RESTART`

type ListAllClusterLibraryStatusesResponse struct {
	// A list of cluster statuses.
	Statuses []ClusterLibraryStatuses `json:"statuses,omitempty"`
}

type MavenLibrary struct {
	// Gradle-style maven coordinates. For example: "org.jsoup:jsoup:1.7.2".
	Coordinates string `json:"coordinates"`
	// List of dependences to exclude. For example: `["slf4j:slf4j",
	// "*:hadoop-client"]`.
	//
	// Maven dependency exclusions:
	// https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html.
	Exclusions []string `json:"exclusions,omitempty"`
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	Repo string `json:"repo,omitempty"`
}

type PythonPyPiLibrary struct {
	// The name of the pypi package to install. An optional exact version
	// specification is also supported. Examples: "simplejson" and
	// "simplejson==3.8.0".
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	Repo string `json:"repo,omitempty"`
}

type RCranLibrary struct {
	// The name of the CRAN package to install.
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	Repo string `json:"repo,omitempty"`
}

type UninstallLibraries struct {
	// Unique identifier for the cluster on which to uninstall these libraries.
	ClusterId string `json:"cluster_id"`
	// The libraries to uninstall.
	Libraries []Library `json:"libraries,omitempty"`
}
