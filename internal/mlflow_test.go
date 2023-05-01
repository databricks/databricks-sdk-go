package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
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

func deleteWorkaround(t *testing.T, w *databricks.WorkspaceClient, path, key, value string) {
	body, err := json.Marshal(map[string]string{key: value})
	require.NoError(t, err)
	req, err := http.NewRequest("DELETE", w.Config.Host+path, bytes.NewBuffer(body))
	require.NoError(t, err)
	err = w.Config.Authenticate(req)
	require.NoError(t, err)
	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, 200, res.StatusCode)
}

// disable once platform is fixed on mlflow side
var deleteWorkaroundEnabled = false

func TestAccModels(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		deleteModel(t, w, ctx, created)
	})

	err = w.ModelRegistry.UpdateModel(ctx, ml.UpdateModelRequest{
		Name:        created.RegisteredModel.Name,
		Description: RandomName("comment "),
	})
	require.NoError(t, err)

	all, err := w.ModelRegistry.ListModelsAll(ctx, ml.ListModelsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func deleteModel(t *testing.T, w *databricks.WorkspaceClient, ctx context.Context, created *ml.CreateModelResponse) {
	if deleteWorkaroundEnabled {
		deleteWorkaround(t, w, "/api/2.0/mlflow/registered-models/delete", "name", created.RegisteredModel.Name)
	} else {
		err := w.ModelRegistry.DeleteModel(ctx, ml.DeleteModelRequest{
			Name: created.RegisteredModel.Name,
		})
		require.NoError(t, err)
	}
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
		if deleteWorkaroundEnabled {
			deleteWorkaround(t, w, "/api/2.0/mlflow/registry-webhooks/delete", "id", created.Webhook.Id)
			return
		}
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

	if !deleteWorkaroundEnabled {
		// we cleanup in  the deleteModel for now.
		time.Sleep(10 * time.Second)
		t.Cleanup(func() {
			err := w.ModelRegistry.DeleteModelVersion(ctx, ml.DeleteModelVersionRequest{
				// TODO: OpenAPI: x-databricks-sdk-inline
				Name:    created.ModelVersion.Name,
				Version: created.ModelVersion.Version,
			})
			require.NoError(t, err)
		})
	}

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
		if deleteWorkaroundEnabled {
			return
		}
		err := w.ModelRegistry.DeleteModelVersion(ctx, ml.DeleteModelVersionRequest{
			// TODO: OpenAPI: x-databricks-sdk-inline
			Name:    mv.ModelVersion.Name,
			Version: mv.ModelVersion.Version,
		})
		require.NoError(t, err)
	})

	created, err := w.ModelRegistry.CreateComment(ctx, ml.CreateComment{
		Comment: RandomName("comment "),
		Name:    model.RegisteredModel.Name,
		Version: mv.ModelVersion.Version,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		// TODO: OpenAPI: x-databricks-sdk-inline
		if deleteWorkaroundEnabled {
			deleteWorkaround(t, w, "/api/2.0/mlflow/comments/delete", "id", created.Comment.Id)
			return
		}
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
