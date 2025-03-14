package integration

import (
	"context"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/ml/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccExperiments(t *testing.T) {
	ctx := workspaceTest(t)
	cfg := &config.Config{}
	ExperimentsAPI, err := ml.NewExperimentsClient(cfg)
	require.NoError(t, err)
	experiment, err := ExperimentsAPI.CreateExperiment(ctx, ml.CreateExperiment{
		Name: RandomName("/tmp/go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := ExperimentsAPI.DeleteExperiment(ctx, ml.DeleteExperiment{
			ExperimentId: experiment.ExperimentId,
		})
		require.NoError(t, err)
	})

	err = ExperimentsAPI.UpdateExperiment(ctx, ml.UpdateExperiment{
		NewName:      RandomName("/tmp/go-sdk-"),
		ExperimentId: experiment.ExperimentId,
	})
	require.NoError(t, err)

	_, err = ExperimentsAPI.GetExperiment(ctx, ml.GetExperimentRequest{
		ExperimentId: experiment.ExperimentId,
	})
	require.NoError(t, err)

	all, err := ExperimentsAPI.ListExperimentsAll(ctx, ml.ListExperimentsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestAccMLflowRuns(t *testing.T) {
	ctx := workspaceTest(t)

	ExperimentsAPI, err := ml.NewExperimentsClient(nil)
	require.NoError(t, err)
	experiment, err := ExperimentsAPI.CreateExperiment(ctx, ml.CreateExperiment{
		Name: RandomName("/tmp/go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := ExperimentsAPI.DeleteExperiment(ctx, ml.DeleteExperiment{
			ExperimentId: experiment.ExperimentId,
		})
		require.NoError(t, err)
	})

	created, err := ExperimentsAPI.CreateRun(ctx, ml.CreateRun{
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
		err := ExperimentsAPI.DeleteRun(ctx, ml.DeleteRun{
			RunId: created.Run.Info.RunId,
		})
		require.NoError(t, err)
	})

	_, err = ExperimentsAPI.UpdateRun(ctx, ml.UpdateRun{
		RunId:  created.Run.Info.RunId,
		Status: ml.UpdateRunStatusKilled,
	})
	require.NoError(t, err)
}

func TestAccModels(t *testing.T) {
	ctx := workspaceTest(t)

	ModelRegistryAPI, err := ml.NewModelRegistryClient(nil)
	require.NoError(t, err)
	created, err := ModelRegistryAPI.CreateModel(ctx, ml.CreateModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		deleteModel(t, ModelRegistryAPI, ctx, created)
	})

	model, err := ModelRegistryAPI.GetModel(ctx, ml.GetModelRequest{
		Name: created.RegisteredModel.Name,
	})
	require.NoError(t, err)

	err = ModelRegistryAPI.UpdateModel(ctx, ml.UpdateModelRequest{
		Name:        model.RegisteredModelDatabricks.Name,
		Description: RandomName("comment "),
	})
	require.NoError(t, err)

	all, err := ModelRegistryAPI.ListModelsAll(ctx, ml.ListModelsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func deleteModel(t *testing.T, ModelRegistryAPI *ml.ModelRegistryClient, ctx context.Context, created *ml.CreateModelResponse) {
	err := ModelRegistryAPI.DeleteModel(ctx, ml.DeleteModelRequest{
		Name: created.RegisteredModel.Name,
	})
	require.NoError(t, err)
}

func TestAccRegistryWebhooks(t *testing.T) {
	ctx := workspaceTest(t)
	cfg := &config.Config{}
	ModelRegistryAPI, err := ml.NewModelRegistryClient(nil)
	require.NoError(t, err)
	created, err := ModelRegistryAPI.CreateWebhook(ctx, ml.CreateRegistryWebhook{
		Description: RandomName("comment "),
		Events:      []ml.RegistryWebhookEvent{ml.RegistryWebhookEventModelVersionCreated},
		HttpUrlSpec: &ml.HttpUrlSpec{
			Url: cfg.CanonicalHostName(),
		},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := ModelRegistryAPI.DeleteWebhook(ctx, ml.DeleteWebhookRequest{
			Id: created.Webhook.Id,
		})
		require.NoError(t, err)
	})
	err = ModelRegistryAPI.UpdateWebhook(ctx, ml.UpdateRegistryWebhook{
		Id:          created.Webhook.Id,
		Description: RandomName("updated "),
	})
	require.NoError(t, err)

	all, err := ModelRegistryAPI.ListWebhooksAll(ctx, ml.ListWebhooksRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func deleteModelVersion(t *testing.T, ModelRegistryAPI *ml.ModelRegistryClient, ctx context.Context, created *ml.CreateModelVersionResponse) {
	// Only delete if model version doesn't have "PENDING_REGISTRATION" status else it could lead to test failure due to race condition
	startTime := time.Now()
	for {
		currentModelVersion, err := ModelRegistryAPI.GetModelVersion(ctx, ml.GetModelVersionRequest{
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

	err := ModelRegistryAPI.DeleteModelVersion(ctx, ml.DeleteModelVersionRequest{
		// TODO: OpenAPI: x-databricks-sdk-inline
		Name:    created.ModelVersion.Name,
		Version: created.ModelVersion.Version,
	})
	require.NoError(t, err)
}

func TestAccModelVersions(t *testing.T) {
	ctx := workspaceTest(t)

	ModelRegistryAPI, err := ml.NewModelRegistryClient(nil)
	require.NoError(t, err)
	model, err := ModelRegistryAPI.CreateModel(ctx, ml.CreateModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		deleteModel(t, ModelRegistryAPI, ctx, model)
	})

	created, err := ModelRegistryAPI.CreateModelVersion(ctx, ml.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		deleteModelVersion(t, ModelRegistryAPI, ctx, created)
	})

	err = ModelRegistryAPI.UpdateModelVersion(ctx, ml.UpdateModelVersionRequest{
		Description: RandomName("description "),
		Name:        created.ModelVersion.Name,
		Version:     created.ModelVersion.Version,
	})
	require.NoError(t, err)
}

func TestAccModelVersionComments(t *testing.T) {
	ctx := workspaceTest(t)

	ModelRegistryAPI, err := ml.NewModelRegistryClient(nil)
	require.NoError(t, err)
	model, err := ModelRegistryAPI.CreateModel(ctx, ml.CreateModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		deleteModel(t, ModelRegistryAPI, ctx, model)
	})

	mv, err := ModelRegistryAPI.CreateModelVersion(ctx, ml.CreateModelVersionRequest{
		Name:   model.RegisteredModel.Name,
		Source: "dbfs:/tmp",
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		deleteModelVersion(t, ModelRegistryAPI, ctx, mv)
	})

	created, err := ModelRegistryAPI.CreateComment(ctx, ml.CreateComment{
		Comment: RandomName("comment "),
		Name:    mv.ModelVersion.Name,
		Version: mv.ModelVersion.Version,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		// TODO: OpenAPI: x-databricks-sdk-inline
		err := ModelRegistryAPI.DeleteComment(ctx, ml.DeleteCommentRequest{
			Id: created.Comment.Id,
		})
		require.NoError(t, err)
	})
	_, err = ModelRegistryAPI.UpdateComment(ctx, ml.UpdateComment{
		Comment: RandomName("updated "),
		Id:      created.Comment.Id,
	})
	require.NoError(t, err)
}
