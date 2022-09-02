package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
)

func createTestPythonFile(t *testing.T, ctx context.Context, apiClient *client.DatabricksClient, testDir string, testId string, content string) string {
	workspaceService := workspace.NewWorkspace(apiClient)
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
