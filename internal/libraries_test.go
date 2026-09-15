package internal

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/require"
)

func TestLibrariesUpdateAndWait(t *testing.T) {
	const clusterID = "test-cluster"
	library := compute.Library{
		Pypi: &compute.PythonPyPiLibrary{
			Package: "test-package",
		},
	}
	installRequest := compute.InstallLibraries{
		ClusterId: clusterID,
		Libraries: []compute.Library{library},
	}
	uninstallRequest := compute.UninstallLibraries{
		ClusterId: clusterID,
		Libraries: []compute.Library{library},
	}

	qa.HTTPFixtures{
		{
			Method:          http.MethodPost,
			Resource:        "/api/2.0/libraries/install",
			ExpectedRequest: installRequest,
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=" + clusterID,
			Response: compute.ClusterLibraryStatuses{
				ClusterId: clusterID,
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Library: &library,
						Status:  compute.LibraryInstallStatusPending,
					},
				},
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=" + clusterID,
			Response: compute.ClusterLibraryStatuses{
				ClusterId: clusterID,
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Library: &library,
						Status:  compute.LibraryInstallStatusInstalled,
					},
				},
			},
		},
		{
			Method:          http.MethodPost,
			Resource:        "/api/2.0/libraries/uninstall",
			ExpectedRequest: uninstallRequest,
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/libraries/cluster-status?cluster_id=" + clusterID,
			Response: compute.ClusterLibraryStatuses{
				ClusterId: clusterID,
				LibraryStatuses: []compute.LibraryFullStatus{
					{
						Library: &library,
						Status:  compute.LibraryInstallStatusUninstallOnRestart,
					},
				},
			},
		},
	}.ApplyClient(t, func(ctx context.Context, apiClient *client.DatabricksClient) {
		libraries := compute.NewLibraries(apiClient)
		timeout := retries.Timeout[compute.ClusterLibraryStatuses](5 * time.Second)

		err := libraries.UpdateAndWait(ctx, compute.Update{
			ClusterId: clusterID,
			Install:   []compute.Library{library},
		}, timeout)
		require.NoError(t, err)

		err = libraries.UpdateAndWait(ctx, compute.Update{
			ClusterId: clusterID,
			Uninstall: []compute.Library{library},
		}, timeout)
		require.NoError(t, err)
	})
}
