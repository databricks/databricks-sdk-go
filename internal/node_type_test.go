package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/stretchr/testify/assert"
)

func TestNodeType(t *testing.T) {
	lst := clusters.ListNodeTypesResponse{
		NodeTypes: []clusters.NodeType{
			{
				NodeTypeId:     "m-fleet.xlarge",
				InstanceTypeId: "m-fleet.xlarge",
				MemoryMb:       16384,
				NumCores:       4,
			},
			{
				NodeTypeId: "Random_05",
				MemoryMb:   1024,
				NumCores:   32,
				NodeInstanceType: &clusters.NodeInstanceType{
					LocalDisks:      3,
					LocalDiskSizeGb: 100,
				},
			},
			{
				NodeTypeId: "Standard_L80s_v2",
				MemoryMb:   655360,
				NumCores:   80,
				NodeInstanceType: &clusters.NodeInstanceType{
					LocalDisks:      2,
					InstanceTypeId:  "Standard_L80s_v2",
					LocalDiskSizeGb: 160,
					LocalNvmeDisks:  1,
				},
			},
			{
				NodeTypeId: "Random_01",
				MemoryMb:   8192,
				NumCores:   8,
				NodeInstanceType: &clusters.NodeInstanceType{
					InstanceTypeId: "_",
				},
			},
			{
				NodeTypeId: "Random_02",
				MemoryMb:   8192,
				NumCores:   8,
				NumGpus:    2,
				NodeInstanceType: &clusters.NodeInstanceType{
					InstanceTypeId: "_",
				},
			},
			{
				NodeTypeId: "Random_03",
				MemoryMb:   8192,
				NumCores:   8,
				NumGpus:    1,
				NodeInstanceType: &clusters.NodeInstanceType{
					InstanceTypeId:      "_",
					LocalNvmeDisks:      15,
					LocalNvmeDiskSizeGb: 235,
				},
			},
			{
				NodeTypeId:   "Random_04",
				MemoryMb:     32000,
				NumCores:     32,
				IsDeprecated: true,
				NodeInstanceType: &clusters.NodeInstanceType{
					LocalDisks:      2,
					LocalDiskSizeGb: 20,
				},
			},
			{
				NodeTypeId: "Standard_F4s",
				MemoryMb:   8192,
				NumCores:   4,
				NodeInstanceType: &clusters.NodeInstanceType{
					LocalDisks:      1,
					LocalDiskSizeGb: 16,
					LocalNvmeDisks:  0,
				},
			},
		},
	}
	nt, err := lst.Smallest(clusters.NodeTypeRequest{LocalDiskMinSize: 200, MinMemoryGB: 8, MinCores: 8, MinGPUs: 1})
	assert.NoError(t, err)
	assert.Equal(t, "Random_03", nt)
}

func TestNodeTypeCategory(t *testing.T) {
	lst := clusters.ListNodeTypesResponse{
		NodeTypes: []clusters.NodeType{
			{
				NodeTypeId:     "Random_05",
				InstanceTypeId: "Random_05",
				MemoryMb:       1024,
				NumCores:       32,
				NodeInstanceType: &clusters.NodeInstanceType{
					LocalDisks:      3,
					LocalDiskSizeGb: 100,
				},
			},
			{
				NodeTypeId:     "Random_01",
				InstanceTypeId: "Random_01",
				MemoryMb:       8192,
				NumCores:       8,
				NodeInstanceType: &clusters.NodeInstanceType{
					InstanceTypeId: "_",
				},
				Category: "Memory Optimized",
			},
			{
				NodeTypeId:     "Random_02",
				InstanceTypeId: "Random_02",
				MemoryMb:       8192,
				NumCores:       8,
				NumGpus:        2,
				Category:       "Storage Optimized",
			},
		},
	}
	nt, err := lst.Smallest(clusters.NodeTypeRequest{Category: "Storage optimized"})
	assert.NoError(t, err)
	assert.Equal(t, "Random_02", nt)
}

func TestNodeTypeCategoryNotAvailable(t *testing.T) {
	lst := clusters.ListNodeTypesResponse{
		NodeTypes: []clusters.NodeType{
			{
				NodeTypeId:     "Random_05",
				InstanceTypeId: "Random_05",
				MemoryMb:       1024,
				NumCores:       32,
				NodeInstanceType: &clusters.NodeInstanceType{
					LocalDisks:      3,
					LocalDiskSizeGb: 100,
				},
			},
			{
				NodeTypeId:     "Random_01",
				InstanceTypeId: "Random_01",
				MemoryMb:       8192,
				NumCores:       8,
				NodeInstanceType: &clusters.NodeInstanceType{
					InstanceTypeId: "_",
				},
				NodeInfo: &clusters.CloudProviderNodeInfo{
					Status: []clusters.CloudProviderNodeStatus{
						clusters.CloudProviderNodeStatusNotavailableinregion,
						clusters.CloudProviderNodeStatusNotenabledonsubscription},
				},
				Category: "Storage Optimized",
			},
			{
				NodeTypeId:     "Random_02",
				InstanceTypeId: "Random_02",
				MemoryMb:       8192,
				NumCores:       8,
				NumGpus:        2,
				Category:       "Storage Optimized",
			},
		},
	}
	nt, err := lst.Smallest(clusters.NodeTypeRequest{Category: "Storage optimized"})
	assert.NoError(t, err)
	assert.Equal(t, "Random_02", nt)
}

func TestNodeTypeFleet(t *testing.T) {
	lst := clusters.ListNodeTypesResponse{
		NodeTypes: []clusters.NodeType{
			{
				NodeTypeId:     "Random_05",
				InstanceTypeId: "Random_05",
				MemoryMb:       1024,
				NumCores:       4,
			},
			{
				NodeTypeId:     "m-fleet.xlarge",
				InstanceTypeId: "m-fleet.xlarge",
				MemoryMb:       16384,
				NumCores:       4,
			},
			{
				NodeTypeId:     "m-fleet.2xlarge",
				InstanceTypeId: "m-fleet.2xlarge",
				MemoryMb:       32768,
				NumCores:       8,
			},
		},
	}
	nt, err := lst.Smallest(clusters.NodeTypeRequest{Fleet: true, MinCores: 8})
	assert.NoError(t, err)
	assert.Equal(t, "m-fleet.2xlarge", nt)
}

func TestNodeTypeEmptyList(t *testing.T) {
	lst := clusters.ListNodeTypesResponse{
		NodeTypes: []clusters.NodeType{},
	}
	_, err := lst.Smallest(clusters.NodeTypeRequest{Fleet: true})
	assert.ErrorContains(t, err, "cannot determine smallest node type with empty response")
}
