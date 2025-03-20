// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// You can use cluster policies to control users' ability to configure clusters
// based on a set of rules. These rules specify which attributes or attribute
// values can be used during cluster creation. Cluster policies have ACLs that
// limit their use to specific users and groups.
//
// With cluster policies, you can: - Auto-install cluster libraries on the next
// restart by listing them in the policy's "libraries" field (Public Preview). -
// Limit users to creating clusters with the prescribed settings. - Simplify the
// user interface, enabling more users to create clusters, by fixing and hiding
// some fields. - Manage costs by setting limits on attributes that impact the
// hourly rate.
//
// Cluster policy permissions limit which policies a user can select in the
// Policy drop-down when the user creates a cluster: - A user who has
// unrestricted cluster create permission can select the Unrestricted policy and
// create fully-configurable clusters. - A user who has both unrestricted
// cluster create permission and access to cluster policies can select the
// Unrestricted policy and policies they have access to. - A user that has
// access to only cluster policies, can select the policies they have access to.
//
// If no policies exist in the workspace, the Policy drop-down doesn't appear.
// Only admin users can create, edit, and delete policies. Admin users also have
// access to all policies.
type ClusterPoliciesClient struct {
	ClusterPoliciesAPI
}

func NewClusterPoliciesClient(cfg *config.Config) (*ClusterPoliciesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ClusterPoliciesClient{
		ClusterPoliciesAPI: ClusterPoliciesAPI{
			clusterPoliciesImpl: clusterPoliciesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The Clusters API allows you to create, start, edit, list, terminate, and
// delete clusters.
//
// Databricks maps cluster node instance types to compute units known as DBUs.
// See the instance type pricing page for a list of the supported instance types
// and their corresponding DBUs.
//
// A Databricks cluster is a set of computation resources and configurations on
// which you run data engineering, data science, and data analytics workloads,
// such as production ETL pipelines, streaming analytics, ad-hoc analytics, and
// machine learning.
//
// You run these workloads as a set of commands in a notebook or as an automated
// job. Databricks makes a distinction between all-purpose clusters and job
// clusters. You use all-purpose clusters to analyze data collaboratively using
// interactive notebooks. You use job clusters to run fast and robust automated
// jobs.
//
// You can create an all-purpose cluster using the UI, CLI, or REST API. You can
// manually terminate and restart an all-purpose cluster. Multiple users can
// share such clusters to do collaborative interactive analysis.
//
// IMPORTANT: Databricks retains cluster configuration information for
// terminated clusters for 30 days. To keep an all-purpose cluster configuration
// even after it has been terminated for more than 30 days, an administrator can
// pin a cluster to the cluster list.
type ClustersClient struct {
	ClustersAPI
}

func NewClustersClient(cfg *config.Config) (*ClustersClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ClustersClient{
		ClustersAPI: ClustersAPI{
			clustersImpl: clustersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// This API allows execution of Python, Scala, SQL, or R commands on running
// Databricks Clusters. This API only supports (classic) all-purpose clusters.
// Serverless compute is not supported.
type CommandExecutionClient struct {
	CommandExecutionAPI
}

func NewCommandExecutionClient(cfg *config.Config) (*CommandExecutionClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CommandExecutionClient{
		CommandExecutionAPI: CommandExecutionAPI{
			commandExecutionImpl: commandExecutionImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The Global Init Scripts API enables Workspace administrators to configure
// global initialization scripts for their workspace. These scripts run on every
// node in every cluster in the workspace.
//
// **Important:** Existing clusters must be restarted to pick up any changes
// made to global init scripts. Global init scripts are run in order. If the
// init script returns with a bad exit code, the Apache Spark container fails to
// launch and init scripts with later position are skipped. If enough containers
// fail, the entire cluster fails with a `GLOBAL_INIT_SCRIPT_FAILURE` error
// code.
type GlobalInitScriptsClient struct {
	GlobalInitScriptsAPI
}

func NewGlobalInitScriptsClient(cfg *config.Config) (*GlobalInitScriptsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GlobalInitScriptsClient{
		GlobalInitScriptsAPI: GlobalInitScriptsAPI{
			globalInitScriptsImpl: globalInitScriptsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Instance Pools API are used to create, edit, delete and list instance pools
// by using ready-to-use cloud instances which reduces a cluster start and
// auto-scaling times.
//
// Databricks pools reduce cluster start and auto-scaling times by maintaining a
// set of idle, ready-to-use instances. When a cluster is attached to a pool,
// cluster nodes are created using the pool’s idle instances. If the pool has
// no idle instances, the pool expands by allocating a new instance from the
// instance provider in order to accommodate the cluster’s request. When a
// cluster releases an instance, it returns to the pool and is free for another
// cluster to use. Only clusters attached to a pool can use that pool’s idle
// instances.
//
// You can specify a different pool for the driver node and worker nodes, or use
// the same pool for both.
//
// Databricks does not charge DBUs while instances are idle in the pool.
// Instance provider billing does apply. See pricing.
type InstancePoolsClient struct {
	InstancePoolsAPI
}

func NewInstancePoolsClient(cfg *config.Config) (*InstancePoolsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &InstancePoolsClient{
		InstancePoolsAPI: InstancePoolsAPI{
			instancePoolsImpl: instancePoolsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The Instance Profiles API allows admins to add, list, and remove instance
// profiles that users can launch clusters with. Regular users can list the
// instance profiles available to them. See [Secure access to S3 buckets] using
// instance profiles for more information.
//
// [Secure access to S3 buckets]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/instance-profiles.html
type InstanceProfilesClient struct {
	InstanceProfilesAPI
}

func NewInstanceProfilesClient(cfg *config.Config) (*InstanceProfilesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &InstanceProfilesClient{
		InstanceProfilesAPI: InstanceProfilesAPI{
			instanceProfilesImpl: instanceProfilesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The Libraries API allows you to install and uninstall libraries and get the
// status of libraries on a cluster.
//
// To make third-party or custom code available to notebooks and jobs running on
// your clusters, you can install a library. Libraries can be written in Python,
// Java, Scala, and R. You can upload Python, Java, Scala and R libraries and
// point to external packages in PyPI, Maven, and CRAN repositories.
//
// Cluster libraries can be used by all notebooks running on a cluster. You can
// install a cluster library directly from a public repository such as PyPI or
// Maven, using a previously installed workspace library, or using an init
// script.
//
// When you uninstall a library from a cluster, the library is removed only when
// you restart the cluster. Until you restart the cluster, the status of the
// uninstalled library appears as Uninstall pending restart.
type LibrariesClient struct {
	LibrariesAPI
}

func NewLibrariesClient(cfg *config.Config) (*LibrariesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &LibrariesClient{
		LibrariesAPI: LibrariesAPI{
			librariesImpl: librariesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The policy compliance APIs allow you to view and manage the policy compliance
// status of clusters in your workspace.
//
// A cluster is compliant with its policy if its configuration satisfies all its
// policy rules. Clusters could be out of compliance if their policy was updated
// after the cluster was last edited.
//
// The get and list compliance APIs allow you to view the policy compliance
// status of a cluster. The enforce compliance API allows you to update a
// cluster to be compliant with the current version of its policy.
type PolicyComplianceForClustersClient struct {
	PolicyComplianceForClustersAPI
}

func NewPolicyComplianceForClustersClient(cfg *config.Config) (*PolicyComplianceForClustersClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PolicyComplianceForClustersClient{
		PolicyComplianceForClustersAPI: PolicyComplianceForClustersAPI{
			policyComplianceForClustersImpl: policyComplianceForClustersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// View available policy families. A policy family contains a policy definition
// providing best practices for configuring clusters for a particular use case.
//
// Databricks manages and provides policy families for several common cluster
// use cases. You cannot create, edit, or delete policy families.
//
// Policy families cannot be used directly to create clusters. Instead, you
// create cluster policies using a policy family. Cluster policies created using
// a policy family inherit the policy family's policy definition.
type PolicyFamiliesClient struct {
	PolicyFamiliesAPI
}

func NewPolicyFamiliesClient(cfg *config.Config) (*PolicyFamiliesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PolicyFamiliesClient{
		PolicyFamiliesAPI: PolicyFamiliesAPI{
			policyFamiliesImpl: policyFamiliesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
