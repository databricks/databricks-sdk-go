package internal

import (
	"context"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccExperiments(t *testing.T) {
	ctx, w := workspaceTest(t)

	experiment, err := w.Experiments.CreateExperiment(ctx, ml.CreateExperiment{
		Name: RandomName("/tmp/go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Experiments.DeleteExperiment(ctx, ml.DeleteExperiment{
			ExperimentId: experiment.ExperimentId,
		})
		require.NoError(t, err)
	})

	err = w.Experiments.UpdateExperiment(ctx, ml.UpdateExperiment{
		NewName:      RandomName("/tmp/go-sdk-"),
		ExperimentId: experiment.ExperimentId,
	})
	require.NoError(t, err)

	_, err = w.Experiments.GetExperiment(ctx, ml.GetExperimentRequest{
		ExperimentId: experiment.ExperimentId,
	})
	require.NoError(t, err)

	all, err := w.Experiments.ListExperimentsAll(ctx, ml.ListExperimentsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestAccMLflowRuns(t *testing.T) {
	ctx, w := workspaceTest(t)

	experiment, err := w.Experiments.CreateExperiment(ctx, ml.CreateExperiment{
		Name: RandomName("/tmp/go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Experiments.DeleteExperiment(ctx, ml.DeleteExperiment{
			ExperimentId: experiment.ExperimentId,
		})
		require.NoError(t, err)
	})

	created, err := w.Experiments.CreateRun(ctx, ml.CreateRun{
		ExperimentId: experiment.ExperimentId,
		Tags: []ml.RunTag{
			{
				Key:   "foo",
				Value: "bar",
			},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		// TODO: OpenAPI: fix mapping
		err := w.Experiments.DeleteRun(ctx, ml.DeleteRun{
			RunId: created.Run.Info.RunId,
		})
		require.NoError(t, err)
	})

	_, err = w.Experiments.UpdateRun(ctx, ml.UpdateRun{
		RunId:  created.Run.Info.RunId,
		Status: ml.UpdateRunStatusKilled,
	})
	require.NoError(t, err)
}

func TestAccModels(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		deleteModel(t, w, ctx, created)
	})

	model, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{
		Name: created.RegisteredModel.Name,
	})
	require.NoError(t, err)

	err = w.ModelRegistry.UpdateModel(ctx, ml.UpdateModelRequest{
		Name:        model.RegisteredModelDatabricks.Name,
		Description: RandomName("comment "),
	})
	require.NoError(t, err)

	all, err := w.ModelRegistry.ListModelsAll(ctx, ml.ListModelsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func deleteModel(t *testing.T, w *databricks.WorkspaceClient, ctx context.Context, created *ml.CreateModelResponse) {
	err := w.ModelRegistry.DeleteModel(ctx, ml.DeleteModelRequest{
		Name: created.RegisteredModel.Name,
	})
	require.NoError(t, err)
}

func TestAccRegistryWebhooks(t *testing.T) {
	ctx, w := workspaceTest(t)
	created, err := w.ModelRegistry.CreateWebhook(ctx, ml.CreateRegistryWebhook{
		Description: RandomName("comment "),
		Events:      []ml.RegistryWebhookEvent{ml.RegistryWebhookEventModelVersionCreated},
		HttpUrlSpec: &ml.HttpUrlSpec{
			Url: w.Config.CanonicalHostName(),
		},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.ModelRegistry.DeleteWebhook(ctx, ml.DeleteWebhookRequest{
			Id: created.Webhook.Id,
		})
		require.NoError(t, err)
	})
	err = w.ModelRegistry.UpdateWebhook(ctx, ml.UpdateRegistryWebhook{
		Id:          created.Webhook.Id,
		Description: RandomName("updated "),
	})
	require.NoError(t, err)

	all, err := w.ModelRegistry.ListWebhooksAll(ctx, ml.ListWebhooksRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func deleteModelVersion(t *testing.T, w *databricks.WorkspaceClient, ctx context.Context, created *ml.CreateModelVersionResponse) {
	// Only delete if model version doesn't have "PENDING_REGISTRATION" status else it could lead to test failure due to race condition
	startTime := time.Now()
	for {
		currentModelVersion, err := w.ModelRegistry.GetModelVersion(ctx, ml.GetModelVersionRequest{
			Name:    created.ModelVersion.Name,
			Version: created.ModelVersion.Version,
		})
		require.NoError(t, err)

		// TODO: x-databricks-wait to be added to all relevant MLflow APIs
		if currentModelVersion.ModelVersion.Status != ml.ModelVersionStatusPendingRegistration {
			break
		}
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(startTime).Seconds()
		if elapsedTime > 10 {
			t.Fatalf("Timeout: Model version still has PENDING_REGISTERATION status after 10 seconds")
		}
		time.Sleep(2 * time.Second)
	}

	err := w.ModelRegistry.DeleteModelVersion(ctx, ml.DeleteModelVersionRequest{
		// TODO: OpenAPI: x-databricks-sdk-inline
		Name:    created.ModelVersion.Name,
		Version: created.ModelVersion.Version,
	})
	require.NoError(t, err)
}

func TestAccModelVersions(t *testing.T) {
	ctx, w := workspaceTest(t)

	model, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		deleteModel(t, w, ctx, model)
	})

	created, err := w.ModelRegistry.CreateModelVersion(ctx, ml.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		deleteModelVersion(t, w, ctx, created)
	})

	err = w.ModelRegistry.UpdateModelVersion(ctx, ml.UpdateModelVersionRequest{
		Description: RandomName("description "),
		Name:        created.ModelVersion.Name,
		Version:     created.ModelVersion.Version,
	})
	require.NoError(t, err)
}

func TestAccModelVersionComments(t *testing.T) {
	ctx, w := workspaceTest(t)

	model, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		deleteModel(t, w, ctx, model)
	})

	mv, err := w.ModelRegistry.CreateModelVersion(ctx, ml.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		deleteModelVersion(t, w, ctx, mv)
	})

	created, err := w.ModelRegistry.CreateComment(ctx, ml.CreateComment{
		Comment: RandomName("comment "),
		Name:    mv.ModelVersion.Name,
		Version: mv.ModelVersion.Version,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		// TODO: OpenAPI: x-databricks-sdk-inline
		err := w.ModelRegistry.DeleteComment(ctx, ml.DeleteCommentRequest{
			Id: created.Comment.Id,
		})
		require.NoError(t, err)
	})
	_, err = w.ModelRegistry.UpdateComment(ctx, ml.UpdateComment{
		Comment: RandomName("updated "),
		Id:      created.Comment.Id,
	})
	require.NoError(t, err)
}
