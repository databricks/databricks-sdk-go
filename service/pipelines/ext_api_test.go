package pipelines

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListPipelinesAll(t *testing.T) {
	ctx := context.Background()

	t.Run("list pipelines with pagination", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?",
				Response: ListPipelinesResponse{
					NextPageToken: "page-token-1",
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-1",
							PipelineId:      "00000000-0000-0000-0000-000000000001",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-2",
							PipelineId:      "00000000-0000-0000-0000-000000000002",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?page_token=page-token-1",
				Response: ListPipelinesResponse{
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-3",
							PipelineId:      "00000000-0000-0000-0000-000000000003",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-4",
							PipelineId:      "00000000-0000-0000-0000-000000000004",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockPipelinesImpl := &pipelinesImpl{
			client: client,
		}
		api := &PipelinesAPI{pipelinesImpl: *mockPipelinesImpl}

		allPipelines, err := api.ListPipelinesAll(ctx, ListPipelinesRequest{})

		require.NoError(t, err)
		assert.Equal(t, 4, len(allPipelines))
		assert.Equal(t, "pipeline-1", allPipelines[0].Name)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", allPipelines[0].PipelineId)
		assert.Equal(t, "pipeline-2", allPipelines[1].Name)
		assert.Equal(t, "00000000-0000-0000-0000-000000000002", allPipelines[1].PipelineId)
		assert.Equal(t, "pipeline-3", allPipelines[2].Name)
		assert.Equal(t, "00000000-0000-0000-0000-000000000003", allPipelines[2].PipelineId)
		assert.Equal(t, "pipeline-4", allPipelines[3].Name)
		assert.Equal(t, "00000000-0000-0000-0000-000000000004", allPipelines[3].PipelineId)
		assert.Equal(t, PipelineStateIdle, allPipelines[0].State)
	})
}

func TestPipelineStateInfoNameToPipelineIdMap(t *testing.T) {
	ctx := context.Background()

	t.Run("name to pipeline id map with pagination", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?",
				Response: ListPipelinesResponse{
					NextPageToken: "page-token-1",
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-1",
							PipelineId:      "00000000-0000-0000-0000-000000000001",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-2",
							PipelineId:      "00000000-0000-0000-0000-000000000002",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?page_token=page-token-1",
				Response: ListPipelinesResponse{
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-3",
							PipelineId:      "00000000-0000-0000-0000-000000000003",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "pipeline-4",
							PipelineId:      "00000000-0000-0000-0000-000000000004",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockPipelinesImpl := &pipelinesImpl{
			client: client,
		}
		api := &PipelinesAPI{pipelinesImpl: *mockPipelinesImpl}

		nameToIdMap, err := api.PipelineStateInfoNameToPipelineIdMap(ctx, ListPipelinesRequest{})

		require.NoError(t, err)
		assert.Equal(t, 4, len(nameToIdMap))
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", nameToIdMap["pipeline-1"])
		assert.Equal(t, "00000000-0000-0000-0000-000000000002", nameToIdMap["pipeline-2"])
		assert.Equal(t, "00000000-0000-0000-0000-000000000003", nameToIdMap["pipeline-3"])
		assert.Equal(t, "00000000-0000-0000-0000-000000000004", nameToIdMap["pipeline-4"])
	})

	t.Run("name to pipeline id map detects duplicates", func(t *testing.T) {
		var requestMocks qa.HTTPFixtures = []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/pipelines?",
				Response: ListPipelinesResponse{
					Statuses: []PipelineStateInfo{
						{
							CreatorUserName: "user-1",
							Name:            "duplicate-pipeline-name",
							PipelineId:      "00000000-0000-0000-0000-000000000001",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
						{
							CreatorUserName: "user-1",
							Name:            "duplicate-pipeline-name",
							PipelineId:      "00000000-0000-0000-0000-000000000002",
							RunAsUserName:   "user-1",
							State:           PipelineStateIdle,
						},
					},
				},
			},
		}
		client, server := requestMocks.Client(t)
		defer server.Close()

		mockPipelinesImpl := &pipelinesImpl{
			client: client,
		}
		api := &PipelinesAPI{pipelinesImpl: *mockPipelinesImpl}

		nameToIdMap, err := api.PipelineStateInfoNameToPipelineIdMap(ctx, ListPipelinesRequest{})

		require.Error(t, err)
		assert.Nil(t, nameToIdMap)
		assert.Contains(t, err.Error(), "duplicate .Name: duplicate-pipeline-name")
	})
}
