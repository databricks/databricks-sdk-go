package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/mlflow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccExperiments(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.Experiments.Create(ctx, mlflow.CreateExperiment{
		Name: RandomName("/tmp/go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Experiments.DeleteByExperimentId(ctx, created.ExperimentId)
		require.NoError(t, err)
	})

	err = w.Experiments.Update(ctx, mlflow.UpdateExperiment{
		NewName:      RandomName("/tmp/go-sdk-"),
		ExperimentId: created.ExperimentId,
	})
	require.NoError(t, err)

	_, err = w.Experiments.GetByExperimentId(ctx, created.ExperimentId)
	require.NoError(t, err)

	all, err := w.Experiments.ListAll(ctx, mlflow.ListExperimentsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestAccMLflowRuns(t *testing.T) {
	ctx, w := workspaceTest(t)

	experiment, err := w.Experiments.Create(ctx, mlflow.CreateExperiment{
		Name: RandomName("/tmp/go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Experiments.DeleteByExperimentId(ctx, experiment.ExperimentId)
		require.NoError(t, err)
	})

	created, err := w.MLflowRuns.Create(ctx, mlflow.CreateRun{
		ExperimentId: experiment.ExperimentId,
		Tags: []mlflow.RunTag{
			{
				Key:   "foo",
				Value: "bar",
			},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		// TODO: OpenAPI: fix mapping
		err := w.MLflowRuns.DeleteByRunId(ctx, created.Run.Info.RunId)
		require.NoError(t, err)
	})

	_, err = w.MLflowRuns.Update(ctx, mlflow.UpdateRun{
		RunId:  created.Run.Info.RunId,
		Status: mlflow.UpdateRunStatusKilled,
	})
	require.NoError(t, err)
}

func TestAccRegisteredModels(t *testing.T) {
	t.Skip("fixme")
	t.Skip("Missing required field: name?...")
	ctx, w := workspaceTest(t)

	created, err := w.RegisteredModels.Create(ctx, mlflow.CreateRegisteredModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		// TODO: OpenAPI: x-databricks-sdk-inline
		// FIXME: Missing required field: name?..
		err := w.RegisteredModels.DeleteByName(ctx, created.RegisteredModel.Name)
		require.NoError(t, err)
	})

	err = w.RegisteredModels.Update(ctx, mlflow.UpdateRegisteredModelRequest{
		Name:        created.RegisteredModel.Name,
		Description: RandomName("comment "),
	})
	require.NoError(t, err)

	_, err = w.RegisteredModels.GetByName(ctx, created.RegisteredModel.Name)
	require.NoError(t, err)

	all, err := w.RegisteredModels.ListAll(ctx, mlflow.ListRegisteredModelsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestAccRegistryWebhooks(t *testing.T) {
	t.Skip("incorrect OpenAPI spec")
	ctx, w := workspaceTest(t)

	// TODO: incomplete ID?...
	created, err := w.RegistryWebhooks.Create(ctx, mlflow.CreateRegistryWebhook{
		Description: RandomName("comment "),
		Events:      []mlflow.RegistryWebhookEvent{mlflow.RegistryWebhookEventModelVersionCreated},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.RegistryWebhooks.DeleteById(ctx, created.Comment.Comment)
		require.NoError(t, err)
	})
	err = w.RegistryWebhooks.Update(ctx, mlflow.UpdateRegistryWebhook{
		Id: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	all, err := w.RegistryWebhooks.ListAll(ctx, mlflow.ListRegistryWebhooksRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestAccModelVersions(t *testing.T) {
	t.Skip("fixme")
	ctx, w := workspaceTest(t)

	model, err := w.RegisteredModels.Create(ctx, mlflow.CreateRegisteredModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		// TODO: OpenAPI: x-databricks-sdk-inline
		// FIXME: Missing required field: name?..
		err := w.RegisteredModels.DeleteByName(ctx, model.RegisteredModel.Name)
		require.NoError(t, err)
	})

	created, err := w.ModelVersions.Create(ctx, mlflow.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.ModelVersions.Delete(ctx, mlflow.DeleteModelVersionRequest{
			// TODO: OpenAPI: x-databricks-sdk-inline
			Name:    created.ModelVersion.Name,
			Version: created.ModelVersion.Version,
		})
		require.NoError(t, err)
	})

	err = w.ModelVersions.Update(ctx, mlflow.UpdateModelVersionRequest{
		Description: RandomName("description "),
		Name:        created.ModelVersion.Name,
		Version:     created.ModelVersion.Version,
	})
	require.NoError(t, err)
}

func TestAccModelVersionComments(t *testing.T) {
	t.Skip("fixme")
	ctx, w := workspaceTest(t)

	model, err := w.RegisteredModels.Create(ctx, mlflow.CreateRegisteredModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		// TODO: OpenAPI: x-databricks-sdk-inline
		// FIXME: Missing required field: name?..
		err := w.RegisteredModels.DeleteByName(ctx, model.RegisteredModel.Name)
		require.NoError(t, err)
	})

	mv, err := w.ModelVersions.Create(ctx, mlflow.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.ModelVersions.Delete(ctx, mlflow.DeleteModelVersionRequest{
			// TODO: OpenAPI: x-databricks-sdk-inline
			Name:    mv.ModelVersion.Name,
			Version: mv.ModelVersion.Version,
		})
		require.NoError(t, err)
	})

	created, err := w.ModelVersionComments.Create(ctx, mlflow.CreateComment{
		Comment: RandomName("comment "),
		Name:    model.RegisteredModel.Name,
		Version: mv.ModelVersion.Version,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		// TODO: OpenAPI: x-databricks-sdk-inline
		// FIXME: missing ID?..
		err := w.ModelVersionComments.DeleteById(ctx, created.Comment.UserId)
		require.NoError(t, err)
	})
	_, err = w.ModelVersionComments.Update(ctx, mlflow.UpdateComment{
		Comment: RandomName("updated "),
		Id:      created.Comment.UserId,
	})
	require.NoError(t, err)
}

func TestAccTransitionRequests(t *testing.T) {
	t.Skip("fixme")
	ctx, w := workspaceTest(t)

	model, err := w.RegisteredModels.Create(ctx, mlflow.CreateRegisteredModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		// TODO: OpenAPI: x-databricks-sdk-inline
		// FIXME: Missing required field: name?..
		err := w.RegisteredModels.DeleteByName(ctx, model.RegisteredModel.Name)
		require.NoError(t, err)
	})

	version := model.RegisteredModel.LatestVersions[0].Version
	_, err = w.TransitionRequests.Create(ctx, mlflow.CreateTransitionRequest{
		Name:    model.RegisteredModel.Name,
		Stage:   mlflow.StageProduction,
		Version: version,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.TransitionRequests.Delete(ctx, mlflow.DeleteTransitionRequestRequest{
			Name:    model.RegisteredModel.Name,
			Version: version,
			Stage:   string(mlflow.StageProduction), // TODO: OpenAPI: make enum
		})
		require.NoError(t, err)
	})

	all, err := w.TransitionRequests.ListAll(ctx, mlflow.ListTransitionRequestsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}
