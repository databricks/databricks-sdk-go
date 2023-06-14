package internal

import (
	"encoding/base64"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccPipelines(t *testing.T) {
	ctx, w := workspaceTest(t)
	notebookPath := myNotebookPath(t, w)

	err := w.Workspace.Import(ctx, workspace.Import{
		Content:   base64.StdEncoding.EncodeToString([]byte(dltNotebook())),
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguageSql,
		Overwrite: true,
		Path:      notebookPath,
	})
	require.NoError(t, err)

	created, err := w.Pipelines.Create(ctx, pipelines.CreatePipeline{
		Continuous: false,
		Name:       RandomName("go-sdk-"),
		Libraries: []pipelines.PipelineLibrary{
			{
				Notebook: &pipelines.NotebookLibrary{
					Path: notebookPath,
				},
			},
		},
		Clusters: []pipelines.PipelineCluster{
			{
				InstancePoolId: GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
				Label:          "default",
				NumWorkers:     1,
				CustomTags: map[string]string{
					"cluster_type": "default",
				},
			},
		},
	})
	require.NoError(t, err)

	events, err := w.Pipelines.ListPipelineEventsAll(ctx, pipelines.ListPipelineEventsRequest{
		PipelineId: created.PipelineId,
	})
	require.NoError(t, err)
	// atleast created event should have been emitted
	assert.Greater(t, len(events), 0)

	defer w.Pipelines.DeleteByPipelineId(ctx, created.PipelineId)

	err = w.Pipelines.Update(ctx, pipelines.EditPipeline{
		PipelineId: created.PipelineId,
		Name:       RandomName("go-sdk-updated-"),
		Libraries: []pipelines.PipelineLibrary{
			{
				Notebook: &pipelines.NotebookLibrary{
					Path: notebookPath,
				},
			},
		},
		Clusters: []pipelines.PipelineCluster{
			{
				InstancePoolId: GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
				Label:          "default",
				NumWorkers:     1,
				CustomTags: map[string]string{
					"cluster_type": "default",
				},
			},
		},
	})
	require.NoError(t, err)

	byId, err := w.Pipelines.GetByPipelineId(ctx, created.PipelineId)
	require.NoError(t, err)

	all, err := w.Pipelines.ListPipelinesAll(ctx, pipelines.ListPipelinesRequest{})
	require.NoError(t, err)

	names, err := w.Pipelines.PipelineStateInfoNameToPipelineIdMap(ctx, pipelines.ListPipelinesRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.PipelineId, names[byId.Name])

	byName, err := w.Pipelines.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byName.PipelineId, byId.PipelineId)
}

// taken from Databricks Terraform Provider acceptance tests
func dltNotebook() string {
	return `CREATE LIVE TABLE clickstream_raw AS 
	SELECT * FROM json.` + "`/databricks-datasets/wikipedia-datasets/data-001/clickstream/raw-uncompressed-json/2015_2_clickstream.json`" + `
	
	-- COMMAND ----------
	
	CREATE LIVE TABLE clickstream_clean(
	  CONSTRAINT valid_current_page EXPECT (current_page_id IS NOT NULL and current_page_title IS NOT NULL),
	  CONSTRAINT valid_count EXPECT (click_count > 0) ON VIOLATION FAIL UPDATE
	) TBLPROPERTIES ("quality" = "silver")
	AS SELECT
	  CAST (curr_id AS INT) AS current_page_id,
	  curr_title AS current_page_title,
	  CAST(n AS INT) AS click_count,
	  CAST (prev_id AS INT) AS previous_page_id,
	  prev_title AS previous_page_title
	FROM live.clickstream_raw
	
	-- COMMAND ----------
	
	CREATE LIVE TABLE top_spark_referers TBLPROPERTIES ("quality" = "gold")
	AS SELECT
	  previous_page_title as referrer,
	  click_count
	FROM live.clickstream_clean
	WHERE current_page_title = 'Apache_Spark'
	ORDER BY click_count DESC
	LIMIT 10`
}
