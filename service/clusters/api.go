package clusters

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/sdk-go/databricks"
	"github.com/databricks/sdk-go/databricks/apierr"
	"github.com/databricks/sdk-go/databricks/client"
	"github.com/databricks/sdk-go/databricks/logger"
	"github.com/databricks/sdk-go/retries"
)

// ClustersAPI is a struct that contains the Databricks api client to perform queries
type ClustersAPI struct {
	client *client.DatabricksClient
}

// New creates ClustersAPI instance from provider meta
func New(c ...*databricks.Config) *ClustersAPI {
	return &ClustersAPI{
		client: client.New(c...),
	}
}

func (a *ClustersAPI) defaultTimeout() time.Duration {
	return 30 * time.Minute
}

// Create creates a new Spark cluster and waits till it's running
func (a *ClustersAPI) Create(ctx context.Context, cluster Cluster) (info ClusterInfo, err error) {
	var ci ClusterID
	err = a.client.Post(ctx, "/clusters/create", cluster, &ci)
	if err != nil {
		return
	}
	info, err = a.waitForClusterStatus(ctx, ci.ClusterID, ClusterStateRunning)
	if err != nil {
		// https://github.com/databricks/terraform-provider-databricks/issues/383
		logger.Errorf("Cleaning up created cluster, that failed to start: %s", err.Error())
		deleteErr := a.PermanentDelete(ctx, ci.ClusterID)
		if deleteErr != nil {
			logger.Errorf("Failed : %s", deleteErr.Error())
			err = deleteErr
		}
	}
	return
}

// Edit edits the configuration of a cluster to match the provided attributes and size
func (a *ClustersAPI) Edit(ctx context.Context, cluster Cluster) (info ClusterInfo, err error) {
	info, err = a.Get(ctx, cluster.ClusterID)
	if err != nil {
		return info, err
	}
	switch info.State {
	case ClusterStateRunning, ClusterStateTerminated:
		// it's already running or terminated, so we're safe to edit
		break
	case ClusterStatePending, ClusterStateResizing, ClusterStateRestarting:
		// let's wait tiny bit, so we return RUNNING cluster info
		info, err = a.waitForClusterStatus(ctx, info.ClusterID, ClusterStateRunning)
		if err != nil {
			return info, err
		}
	case ClusterStateTerminating:
		// let it finish terminating, so it's safe to edit.
		// TERMINATED cluster info will be returned this way
		info, err = a.waitForClusterStatus(ctx, info.ClusterID, ClusterStateTerminated)
		if err != nil {
			return info, err
		}
	case ClusterStateError, ClusterStateUnknown:
		// we don't know what to do, so return error
		return info, fmt.Errorf("unexpected state: %#v", info.StateMessage)
	}
	err = a.client.Post(ctx, "/clusters/edit", cluster, nil)
	if err != nil {
		return info, err
	}
	if info.IsRunningOrResizing() {
		// so if cluster was running, we'll start and wait again
		return a.StartAndGetInfo(ctx, info.ClusterID)
	}
	// only State / ClusterID properties will be valid in this return
	return info, err
}

// ListZones returns the zones info sent by the cloud service provider
func (a *ClustersAPI) ListZones(ctx context.Context) (ZonesInfo, error) {
	var zonesInfo ZonesInfo
	err := a.client.Get(ctx, "/clusters/list-zones", nil, &zonesInfo)
	return zonesInfo, err
}

// Start a terminated Spark cluster given its ID and wait till it's running
func (a *ClustersAPI) Start(ctx context.Context, clusterID string) error {
	_, err := a.StartAndGetInfo(ctx, clusterID)
	return err
}

// StartAndGetInfo starts cluster and returns info
func (a *ClustersAPI) StartAndGetInfo(ctx context.Context, clusterID string) (ClusterInfo, error) {
	info, err := a.Get(ctx, clusterID)
	if err != nil {
		return info, err
	}
	switch info.State {
	case ClusterStateRunning:
		// it's already running, so we're good to return
		return info, nil
	case ClusterStatePending, ClusterStateResizing, ClusterStateRestarting:
		// let's wait tiny bit, so we return RUNNING cluster info
		return a.waitForClusterStatus(ctx, info.ClusterID, ClusterStateRunning)
	case ClusterStateTerminating:
		// let it finish terminating, so it's safe to start again.
		// TERMINATED cluster info will be returned this way
		info, err = a.waitForClusterStatus(ctx, info.ClusterID, ClusterStateTerminated)
		if err != nil {
			return info, err
		}
	case ClusterStateError, ClusterStateUnknown:
		// most likely we can start error'ed cluster again...
		logger.Errorf("Cluster %s: %s", info.State, info.StateMessage)
	}
	err = a.client.Post(ctx, "/clusters/start", ClusterID{ClusterID: clusterID}, nil)
	if err != nil {
		if !strings.Contains(err.Error(),
			fmt.Sprintf("Cluster %s is in unexpected state Pending.", clusterID)) {
			return info, err
		}
	}
	return a.waitForClusterStatus(ctx, clusterID, ClusterStateRunning)
}

func (a *ClustersAPI) waitForClusterStatus(ctx context.Context, clusterID string, desired ClusterState) (result ClusterInfo, err error) {
	return result, retries.Wait(ctx, a.defaultTimeout(), func() *retries.Err {
		clusterInfo, err := a.Get(ctx, clusterID)
		if apierr.IsMissing(err) {
			logger.Infof("Cluster %s not found. Retrying", clusterID)
			return retries.Continue(err)
		}
		if err != nil {
			return retries.Halt(err)
		}
		result = clusterInfo
		logger.Debugf("Cluster %s is %s: %s", clusterID, clusterInfo.State, clusterInfo.StateMessage)
		if clusterInfo.State == desired {
			return nil
		}
		if !clusterInfo.State.CanReach(desired) {
			docLink := "https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterstate"
			details := ""
			if clusterInfo.TerminationReason != nil {
				details = fmt.Sprintf(", Termination info: code: %s, type: %s, parameters: %v",
					clusterInfo.TerminationReason.Code, clusterInfo.TerminationReason.Type,
					clusterInfo.TerminationReason.Parameters)
			}
			return retries.Continues(fmt.Sprintf(
				"%s is not able to transition from %s to %s: %s%s. Please see %s for more details",
				clusterID, clusterInfo.State, desired, clusterInfo.StateMessage, details, docLink))
		}
		return retries.Continues(fmt.Sprintf("%s is %s, but has to be %s", clusterID, clusterInfo.State, desired))
	})
}

// Terminate terminates a Spark cluster given its ID
func (a *ClustersAPI) Terminate(ctx context.Context, clusterID string) error {
	err := a.client.Post(ctx, "/clusters/delete", ClusterID{ClusterID: clusterID}, nil)
	if err != nil {
		return err
	}
	_, err = a.waitForClusterStatus(ctx, clusterID, ClusterStateTerminated)
	return err
}

// PermanentDelete permanently delete a cluster
func (a *ClustersAPI) PermanentDelete(ctx context.Context, clusterID string) error {
	err := a.Terminate(ctx, clusterID)
	if err != nil {
		return err
	}
	r := ClusterID{ClusterID: clusterID}
	err = a.client.Post(ctx, "/clusters/permanent-delete", r, nil)
	if err == nil {
		return nil
	}
	if !strings.Contains(err.Error(), "unpin the cluster first") {
		return err
	}
	// unpin cluster if it's pinned
	err = a.Unpin(ctx, clusterID)
	if err != nil {
		return err
	}
	// and try removing it again
	return a.client.Post(ctx, "/clusters/permanent-delete", r, nil)
}

// Get retrieves the information for a cluster given its identifier
func (a *ClustersAPI) Get(ctx context.Context, clusterID string) (ci ClusterInfo, err error) {
	err = wrapMissingClusterError(a.client.Get(ctx, "/clusters/get",
		ClusterID{ClusterID: clusterID}, &ci), clusterID)
	return
}

// Pin ensure that an interactive cluster configuration is retained even after a cluster has been terminated for more than 30 days
func (a *ClustersAPI) Pin(ctx context.Context, clusterID string) error {
	return a.client.Post(ctx, "/clusters/pin", ClusterID{ClusterID: clusterID}, nil)
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a *ClustersAPI) Unpin(ctx context.Context, clusterID string) error {
	return a.client.Post(ctx, "/clusters/unpin", ClusterID{ClusterID: clusterID}, nil)
}

// Events - only using Cluster ID string to get all events
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#events
func (a *ClustersAPI) Events(ctx context.Context, eventsRequest EventsRequest) ([]ClusterEvent, error) {
	var eventsResponse EventsResponse
	err := a.client.Post(ctx, "/clusters/events", eventsRequest, &eventsResponse)
	if err != nil {
		return nil, err
	}

	totalCount := int(eventsResponse.TotalCount)
	if (eventsRequest.MaxItems) > 0 && (eventsRequest.MaxItems < uint(totalCount)) {
		totalCount = int(eventsRequest.MaxItems)
	}
	events := make([]ClusterEvent, totalCount)
	if totalCount == 0 {
		return events, nil
	}
	startPos := 0
	curPos := len(eventsResponse.Events)
	copy(events[startPos:curPos], eventsResponse.Events)
	for curPos < totalCount && eventsResponse.NextPage != nil {
		err := a.client.Post(ctx, "/clusters/events", eventsResponse.NextPage, &eventsResponse)
		if err != nil {
			return nil, err
		}
		startPos = curPos
		curLen := len(eventsResponse.Events)
		restItems := totalCount - startPos
		if restItems < curLen {
			curLen = restItems
		}
		curPos += curLen
		copy(events[startPos:curPos], eventsResponse.Events[0:curLen])
	}

	return events[0:curPos], err
}

// List return information about all pinned clusters, currently active clusters,
// up to 70 of the most recently terminated interactive clusters in the past 30 days,
// and up to 30 of the most recently terminated job clusters in the past 30 days
func (a *ClustersAPI) List(ctx context.Context) ([]ClusterInfo, error) {
	var clusterList ClusterList
	err := a.client.Get(ctx, "/clusters/list", nil, &clusterList)
	return clusterList.Clusters, err
}

// ListNodeTypes returns a sorted list of supported Spark node types
func (a ClustersAPI) ListNodeTypes(ctx context.Context) (l NodeTypeList, err error) {
	err = a.client.Get(ctx, "/clusters/list-node-types", nil, &l)
	return
}

// ListSparkVersions returns smallest (or default) node type id given the criteria
func (a *ClustersAPI) ListSparkVersions(ctx context.Context) (SparkVersionsList, error) {
	var sparkVersions SparkVersionsList
	err := a.client.Get(ctx, "/clusters/spark-versions", nil, &sparkVersions)
	return sparkVersions, err
}
