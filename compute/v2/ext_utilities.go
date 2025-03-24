package compute

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/databricks/log"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
)

// getOrCreateClusterMutex guards "mounting" cluster creation to prevent multiple
// redundant instances created at the same name. Compute package private property.
// https://github.com/databricks/terraform-provider-databricks/issues/445
var getOrCreateClusterMutex sync.Mutex

func (c *ClusterDetails) IsRunningOrResizing() bool {
	return c.State == StateRunning || c.State == StateResizing
}

// use mutex around starting a cluster by ID
var mu sync.Mutex

func (a *ClustersClient) isErrFailedToReach(err error) bool {
	if err == nil {
		return false
	}
	// TODO: get a bit better handling of these
	return strings.HasPrefix(err.Error(), "failed to reach")
}

func waitGetClusterTerminated(ctx context.Context, c *ClustersClient, clusterId string, timeout time.Duration) (*ClusterDetails, error) {
	return retries.Poll[ClusterDetails](ctx, timeout, func() (*ClusterDetails, *retries.Err) {
		clusterDetails, err := c.Get(ctx, GetClusterRequest{
			ClusterId: clusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := clusterDetails.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case StateTerminated: // target state
			return clusterDetails, nil
		case StateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				StateTerminated, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

func waitGetClusterRunning(ctx context.Context, c *ClustersClient, clusterId string, timeout time.Duration) (*ClusterDetails, error) {
	return retries.Poll[ClusterDetails](ctx, timeout, func() (*ClusterDetails, *retries.Err) {
		clusterDetails, err := c.Get(ctx, GetClusterRequest{
			ClusterId: clusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := clusterDetails.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case StateRunning: // target state
			return clusterDetails, nil
		case StateError, StateTerminated:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				StateTerminated, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

func (a *ClustersClient) startClusterIfNeeded(ctx context.Context, clusterId string, timeout time.Duration) error {
	info, err := a.GetByClusterId(ctx, clusterId)
	if err != nil {
		return fmt.Errorf("get cluster info: %w", err)
	}
	switch info.State {
	case StateRunning:
		return nil
	case StateTerminating:
		_, err = waitGetClusterTerminated(ctx, a, clusterId, timeout)
		if err != nil {
			return fmt.Errorf("terminating -> terminated: %w", err)
		}
		waiter, err := a.Start(ctx, StartCluster{
			ClusterId: clusterId,
		})
		waiter.WaitUntilDone(ctx, nil)
		if err != nil {
			return fmt.Errorf("terminated -> terminating -> start: %w", err)
		}
		return nil
	case StateTerminated:
		waiter, err := a.Start(ctx, StartCluster{
			ClusterId: clusterId,
		})
		waiter.WaitUntilDone(ctx, nil)
		if err != nil {
			return fmt.Errorf("terminated -> start: %w", err)
		}
		return nil
	case StatePending, StateResizing, StateRestarting:
		_, err = waitGetClusterRunning(ctx, a, clusterId, timeout)
		if err != nil {
			return fmt.Errorf("wait: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("cluster %s is in %s state: %s", info.ClusterName, info.State, info.StateMessage)
	}
}

func (a *ClustersClient) EnsureClusterIsRunning(ctx context.Context, clusterId string) error {
	mu.Lock()
	defer mu.Unlock()
	timeout := 20 * time.Minute
	return retries.Wait(ctx, timeout, func() *retries.Err {
		err := a.startClusterIfNeeded(ctx, clusterId, timeout)
		var apiErr *apierr.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode == "INVALID_STATE" {
			log.DebugContext(ctx, "Cluster was started by other process: %s Retrying.", apiErr.Message)
			return retries.Continue(apiErr)
		}
		if a.isErrFailedToReach(err) {
			return retries.Continue(err)
		}
		if err != nil {
			return retries.Halt(err)
		}
		return nil
	})
}

// TODO: This method currently contains cloud-specific logic (AWS vs others) that requires
// the client to be aware of the cloud provider. However, we don't want client to contain
// information about the cloud. We're temporarily including this method, but need to
// decide on a better approach for handling these mixins before the SDK-Mod release.
//
// GetOrCreateRunningCluster creates an autoterminating cluster if it doesn't exist
func (a *ClustersClient) GetOrCreateRunningCluster(ctx context.Context, name string, custom ...CreateCluster) (c *ClusterDetails, err error) {
	getOrCreateClusterMutex.Lock()
	defer getOrCreateClusterMutex.Unlock()
	if len(custom) > 1 {
		err = fmt.Errorf("you can only specify 1 custom cluster conf, not %d", len(custom))
		return
	}
	clusters, err := a.ListAll(ctx, ListClustersRequest{})
	if err != nil {
		return
	}
	for _, cl := range clusters {
		if cl.ClusterName != name {
			continue
		}
		log.InfoContext(ctx, "Found reusable cluster '%s'", name)
		if cl.IsRunningOrResizing() {
			return &cl, nil
		}
		waiter, err := a.Start(ctx, StartCluster{
			ClusterId: cl.ClusterId,
		})
		if err != nil {
			log.InfoContext(ctx, "Cluster %s cannot be started, creating an autoterminating cluster: %s", name, err)
			break
		}
		started, err := waiter.WaitUntilDone(ctx, nil)
		if err != nil {
			log.InfoContext(ctx, "Cluster %s cannot be started, creating an autoterminating cluster: %s", name, err)
			break
		}
		return started, nil
	}
	nodeTypes, err := a.ListNodeTypes(ctx)
	if err != nil {
		return nil, fmt.Errorf("list node types: %w", err)
	}
	smallestNodeType, err := nodeTypes.Smallest(NodeTypeRequest{
		LocalDisk: true,
	})
	if err != nil {
		return nil, err
	}
	log.InfoContext(ctx, "Creating an autoterminating cluster with node type %s", smallestNodeType)
	versions, err := a.SparkVersions(ctx)
	if err != nil {
		return nil, fmt.Errorf("list spark versions: %w", err)
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
	if a.clustersImpl.client.IsAws() {
		r.AwsAttributes = &AwsAttributes{
			Availability: "SPOT",
		}
	}
	if len(custom) == 1 {
		r = custom[0]
	}
	waiter, err := a.Create(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("create cluster: %w", err)
	}
	return waiter.WaitUntilDone(ctx, nil)
}
