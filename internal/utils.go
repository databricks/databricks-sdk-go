package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
)

func CreateTestPythonFile(t *testing.T, ctx context.Context, apiClient *client.DatabricksClient, testDir string, testId string, content string) string {
	workspaceService := workspace.New(apiClient)
	err := workspaceService.Mkdirs(ctx, workspace.MkdirsRequest{
		Path: testDir,
	})
	assert.NoError(t, err)
	err = workspaceService.Import(ctx, workspace.ImportRequest{
		Content:   base64.StdEncoding.Strict().EncodeToString([]byte(content)),
		Format:    "SOURCE",
		Language:  "PYTHON",
		Overwrite: true,
		Path:      fmt.Sprintf("%s/test-%s.py", testDir, testId),
	})
	assert.NoError(t, err)
	_, err = workspaceService.GetStatus(ctx, workspace.GetStatusRequest{
		Path: fmt.Sprintf("%s/test-%s.py", testDir, testId),
	})
	assert.NoError(t, err)
	return fmt.Sprintf("%s/test-%s.py", testDir, testId)
}

func StartDefaultTestCluster(t *testing.T, ctx context.Context, apiClient *client.DatabricksClient) {
	clusterService := clusters.New(apiClient)
	err := retries.Wait(ctx, 10*time.Minute, func() *retries.Err {
		clusterDetails, err := clusterService.GetCluster(ctx, clusters.GetClusterRequest{
			ClusterId: GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID"),
		})
		if err != nil {
			return retries.Halt(err)
		}
		if clusterDetails.State == clusters.GetClusterResponseStateRunning {
			return nil
		}
		if clusterDetails.State == clusters.GetClusterResponseStateTerminated {
			err = clusterService.StartCluster(ctx, clusters.StartClusterRequest{
				ClusterId: GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID"),
			})
			if err != nil {
				return retries.Halt(err)
			}
		}
		return retries.Continue(fmt.Errorf("%s", clusterDetails.State))
	})
	assert.NoError(t, err)
}
