// TODO : Add the missing methods and implement the methods
// This file has not been completely shifted from the SDK-Beta
// as we still don't have the wait for state methods in the SDK-mod
package compute

import (
	"strings"
	"sync"
)

type clustersAPIUtilities interface {
}

// getOrCreateClusterMutex guards "mounting" cluster creation to prevent multiple
// redundant instances created at the same name. Compute package private property.
// https://github.com/databricks/terraform-provider-databricks/issues/445
var getOrCreateClusterMutex sync.Mutex

func (c *ClusterDetails) IsRunningOrResizing() bool {
	return c.State == StateRunning || c.State == StateResizing
}

// use mutex around starting a cluster by ID
var mu sync.Mutex

func (a *ClustersAPI) isErrFailedToReach(err error) bool {
	if err == nil {
		return false
	}
	// TODO: get a bit better handling of these
	return strings.HasPrefix(err.Error(), "failed to reach")
}
