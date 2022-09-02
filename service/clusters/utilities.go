package clusters

import (
	"context"
	"fmt"
	"sync"

	"github.com/databricks/databricks-sdk-go/databricks/logger"
)

// getOrCreateClusterMutex guards "mounting" cluster creation to prevent multiple
// redundant instances created at the same name. Compute package private property.
// https://github.com/databricks/terraform-provider-databricks/issues/445
var getOrCreateClusterMutex sync.Mutex

func (ci *ClusterInfo) IsRunningOrResizing() bool {
	return ci.State == ClusterInfoStateRunning || ci.State == ClusterInfoStateResizing
}

// TODO: Consistency
// CreateClusterResponse -> ClusterInfo ?..
// CreateClusterRequest -> Cluster

// GetOrCreateRunningCluster creates an autoterminating cluster if it doesn't exist
func (a *ClustersAPI) GetOrCreateRunningCluster(ctx context.Context, name string, custom ...CreateCluster) (c *CreateClusterResponse, err error) {
	getOrCreateClusterMutex.Lock()
	defer getOrCreateClusterMutex.Unlock()
	if len(custom) > 1 {
		err = fmt.Errorf("you can only specify 1 custom cluster conf, not %d", len(custom))
		return
	}
	clusters, err := a.List(ctx, ListRequest{})
	if err != nil {
		return
	}
	for _, cl := range clusters.Clusters {
		if cl.ClusterName == name {
			logger.Infof("Found reusable cluster '%s'", name)
			clusterAvailable := true
			if !cl.IsRunningOrResizing() {
				err = a.Start(ctx, StartCluster{
					ClusterId: cl.ClusterId,
				})
				if err != nil {
					clusterAvailable = false
					logger.Infof("Cluster %s cannot be started, creating an autoterminating cluster", name)
				}
			}
			if clusterAvailable {
				return &CreateClusterResponse{
					ClusterId: cl.ClusterId,
				}, nil
			}
		}
	}
	nodeTypes, err := a.ListNodeTypes(ctx)
	if err != nil {
		return nil, fmt.Errorf("list node types: %w", err) // TODO: GENERATOR: prefix method name
	}
	smallestNodeType, err := nodeTypes.Smallest(NodeTypeRequest{
		LocalDisk: true,
	})
	if err != nil {
		return nil, err
	}
	logger.Infof("Creating an autoterminating cluster with node type %s", smallestNodeType)
	versions, err := a.SparkVersions(ctx)
	if err != nil {
		return nil, fmt.Errorf("list spark versions: %w", err) // TODO: GENERATOR: prefix method name
	}
	version, err := versions.Select(SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		return nil, err
	}
	r := CreateCluster{
		NumWorkers:             1,
		ClusterName:            name,
		SparkVersion:           version,
		NodeTypeId:             smallestNodeType,
		AutoterminationMinutes: 10,
	}
	if a.client.Config.IsAws() {
		r.AwsAttributes = &AwsAttributes{
			Availability: "SPOT",
		}
	}
	if len(custom) == 1 {
		r = custom[0]
	}
	return a.Create(ctx, r)
}
