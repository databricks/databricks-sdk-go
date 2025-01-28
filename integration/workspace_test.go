package integration

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/ml/v2"
	"github.com/databricks/databricks-sdk-go/workspace/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func myNotebookPath(t *testing.T, cfg *config.Config) string {
	ctx := context.Background()
	workspaceAPI, err := workspace.NewWorkspaceClient(cfg)
	require.NoError(t, err)
	testDir := filepath.Join("/Users", me(t, cfg).UserName, ".sdk", "notebooks", RandomName("t-"))
	notebook := filepath.Join(testDir, RandomName("n-"))

	err = workspaceAPI.MkdirsByPath(ctx, testDir)
	require.NoError(t, err)
	t.Cleanup(func() {
		err = workspaceAPI.Delete(ctx, workspace.Delete{
			Path:      testDir,
			Recursive: true,
		})
		require.NoError(t, err)
	})

	return notebook
}

// TODO: Enable this testing after testing infrastructure is moved to OAuth
func TestGetOAuthToken(t *testing.T) {
	t.Skip("This test only works with OAuth credentials. Automatic testing uses PAT tokens")
	// Create an Endpoint that can be used to query
	ctx := workspaceTest(t)
	// For manual testing using Databricks CLI credentials
	// w = databricks.Must(databricks.NewWorkspaceClient())
	ModelRegistryApi, err := ml.NewModelRegistryClient(nil)
	created, err := ModelRegistryApi.CreateModel(ctx, ml.CreateModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		ModelRegistryApi.DeleteModel(ctx, ml.DeleteModelRequest{
			Name: created.RegisteredModel.Name,
		})
		require.NoError(t, err)
	})

	model, err := ModelRegistryApi.GetModel(ctx, ml.GetModelRequest{
		Name: created.RegisteredModel.Name,
	})
	require.NoError(t, err)

	err = ModelRegistryApi.UpdateModel(ctx, ml.UpdateModelRequest{
		Name:        model.RegisteredModelDatabricks.Name,
		Description: RandomName("comment "),
	})
	require.NoError(t, err)

	all, err := ModelRegistryApi.ListModelsAll(ctx, ml.ListModelsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)

	// Enable once the API returns the DataplaneInfo
	// endpoint, err := w.ServingEndpoints.Create(ctx, serving.CreateServingEndpoint{
	// 	Name: RandomName("go-sdk-"),
	// 	Config: serving.EndpointCoreConfigInput{
	// 		ServedModels: []serving.ServedModelInput{
	// 			{
	// 				ModelName:          model.RegisteredModelDatabricks.Name,
	// 				ModelVersion:       model.RegisteredModelDatabricks.LatestVersions[0].Version,
	// 				Name:               RandomName("go-sdk-"),
	// 				ScaleToZeroEnabled: true,
	// 				WorkloadSize:       serving.ServedModelInputWorkloadSizeSmall,
	// 			},
	// 		},
	// 	},
	// })
	// require.NoError(t, err)
	// t.Cleanup(func() {
	// 	err := w.ServingEndpoints.Delete(ctx, serving.DeleteServingEndpointRequest{
	// 		Name: endpoint.Name,
	// 	})
	// 	require.NoError(t, err)
	// })
	// readyEndpoint, err := endpoint.Get()
	// require.NoError(t, err)

	// _, err = w.GetOAuthToken(readyEndpoint.DataplaneInfo.AuthorizationDetails)
	// require.NoError(t, err)
	// // TODO: Query DataPlane API with the token
}

func TestAccWorkspaceIntegration(t *testing.T) {
	ctx := workspaceTest(t)
	WorkspaceAPI, err := workspace.NewWorkspaceClient(nil)
	require.NoError(t, err)
	notebook := myNotebookPath(t, WorkspaceAPI.Config)

	// Import the test notebook
	err = WorkspaceAPI.Import(ctx, workspace.Import{
		Path:      notebook,
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguagePython,
		Content:   base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")),
		Overwrite: true,
	})
	require.NoError(t, err)

	// Get test notebook status
	getStatusResponse, err := WorkspaceAPI.GetStatusByPath(ctx, notebook)
	require.NoError(t, err)
	assert.True(t, getStatusResponse.Language == workspace.LanguagePython)
	assert.True(t, getStatusResponse.Path == notebook)
	assert.True(t, getStatusResponse.ObjectType == workspace.ObjectTypeNotebook)

	// Export the notebook and assert the contents
	exportResponse, err := WorkspaceAPI.Export(ctx, workspace.ExportRequest{
		Format: workspace.ExportFormatSource,
		Path:   notebook,
	})
	require.NoError(t, err)
	assert.True(t, exportResponse.Content == base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")))

	// Assert the test notebook is present in test dir using list api
	objects, err := WorkspaceAPI.ListAll(ctx, workspace.ListWorkspaceRequest{
		Path: filepath.Dir(notebook),
	})
	require.NoError(t, err)

	paths, err := WorkspaceAPI.ObjectInfoPathToObjectIdMap(ctx, workspace.ListWorkspaceRequest{
		Path: filepath.Dir(notebook),
	})
	require.NoError(t, err)
	assert.Equal(t, len(objects), len(paths))
	assert.Contains(t, paths, notebook)
}

func TestAccWorkspaceUploadNotebookWithFileExtensionNoTranspile(t *testing.T) {
	// TODO: remove NoTranspile suffix once other languages get Upload/Donwload features
	ctx := workspaceTest(t)

	WorkspaceAPI, err := workspace.NewWorkspaceClient(nil)
	require.NoError(t, err)
	notebookPath := filepath.Join("/Users", me(t, WorkspaceAPI.Config).UserName, RandomName("notebook-")+".py")

	err = WorkspaceAPI.Upload(ctx, notebookPath, strings.NewReader("print(1)"))
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = WorkspaceAPI.Delete(ctx, workspace.Delete{
			Path: notebookPath,
		})
		require.NoError(t, err)
	})

	info, err := WorkspaceAPI.GetStatusByPath(ctx, notebookPath)
	assert.NoError(t, err)
	assert.Equal(t, workspace.LanguagePython, info.Language)
	assert.Equal(t, workspace.ObjectTypeNotebook, info.ObjectType)

	contents, err := WorkspaceAPI.ReadFile(ctx, notebookPath)
	assert.NoError(t, err)

	assert.Equal(t, "# Databricks notebook source\nprint(1)", string(contents))
}

func TestAccWorkspaceUploadNotebookWithFileNoExtensionNoTranspile(t *testing.T) {
	// TODO: remove NoTranspile suffix once other languages get Upload/Donwload features
	ctx := workspaceTest(t)

	WorkspaceAPI, err := workspace.NewWorkspaceClient(nil)
	require.NoError(t, err)
	notebookPath := filepath.Join("/Users", me(t, WorkspaceAPI.Config).UserName, RandomName("notebook-"))

	err = WorkspaceAPI.Upload(ctx, notebookPath, strings.NewReader("print(1)"),
		workspace.UploadLanguage(workspace.LanguagePython))
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = WorkspaceAPI.Delete(ctx, workspace.Delete{
			Path: notebookPath,
		})
		require.NoError(t, err)
	})

	info, err := WorkspaceAPI.GetStatusByPath(ctx, notebookPath)
	assert.NoError(t, err)
	assert.Equal(t, workspace.LanguagePython, info.Language)
	assert.Equal(t, workspace.ObjectTypeNotebook, info.ObjectType)

	contents, err := WorkspaceAPI.ReadFile(ctx, notebookPath)
	assert.NoError(t, err)

	assert.Equal(t, "# Databricks notebook source\nprint(1)", string(contents))
}

func TestAccWorkspaceUploadFileNoTranspile(t *testing.T) {
	// TODO: remove NoTranspile suffix once other languages get Upload/Donwload features
	ctx := workspaceTest(t)

	WorkspaceAPI, err := workspace.NewWorkspaceClient(nil)
	require.NoError(t, err)
	txtPath := filepath.Join("/Users", me(t, WorkspaceAPI.Config).UserName, RandomName("txt-"))

	err = WorkspaceAPI.Upload(ctx, txtPath, strings.NewReader("print(1)"),
		workspace.UploadFormat(workspace.ImportFormatAuto))
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = WorkspaceAPI.Delete(ctx, workspace.Delete{
			Path: txtPath,
		})
		require.NoError(t, err)
	})

	info, err := WorkspaceAPI.GetStatusByPath(ctx, txtPath)
	assert.NoError(t, err)
	assert.Equal(t, workspace.ObjectTypeFile, info.ObjectType)

	contents, err := WorkspaceAPI.ReadFile(ctx, txtPath)
	assert.NoError(t, err)

	assert.Equal(t, "print(1)", string(contents))
}

func TestAccWorkspaceRecursiveListNoTranspile(t *testing.T) {
	ctx := workspaceTest(t)
	WorkspaceAPI, err := workspace.NewWorkspaceClient(nil)
	require.NoError(t, err)
	notebook := myNotebookPath(t, WorkspaceAPI.Config)

	// Import the test notebook
	err = WorkspaceAPI.Upload(ctx, notebook+".py", strings.NewReader("print(1)"),
		workspace.UploadOverwrite())
	require.NoError(t, err)

	allMyNotebooks, err := WorkspaceAPI.RecursiveList(ctx, filepath.Join("/Users", me(t, WorkspaceAPI.Config).UserName, ".sdk"))
	require.NoError(t, err)
	assert.True(t, len(allMyNotebooks) >= 1)
}
