package databricks

import (
	"testing"
)

func TestReal(t *testing.T) {
	// config := &Config{
	// 	AuthType: "azure-cli",
	// 	Profile:  "DEFAULT",
	// }
	// w := Must(NewWorkspaceClient(config))

	// clusterDetails := compute.ClusterSpec{
	// 	NumWorkers:   0,
	// 	ClusterName:  "HectorTestGoLangDefault",
	// 	SparkVersion: "10.4.x-scala2.12",
	// 	SparkConf: map[string]string{
	// 		"spark.databricks.delta.preview.enabled": "true",
	// 		"spark.databricks.cluster.profile":       "singleNode",
	// 		"spark.master":                           "local[*]",
	// 	},
	// 	AzureAttributes: &compute.AzureAttributes{
	// 		FirstOnDemand:   1,
	// 		Availability:    compute.AzureAvailabilityOnDemandAzure,
	// 		SpotBidMaxPrice: -1,
	// 	},
	// 	NodeTypeId:       "Standard_DS3_v2",
	// 	DriverNodeTypeId: "Standard_DS3_v2",
	// 	SshPublicKeys:    []string{},
	// 	CustomTags: map[string]string{
	// 		"ResourceClass": "SingleNode",
	// 	},
	// 	SparkEnvVars: map[string]string{
	// 		"PIP_EXTRA_INDEX_URL_SECRET": "{{secrets/shared/PIP_EXTRA_INDEX_URL_SECRET}}",
	// 		"PIP_EXTRA_INDEX_URL_USER":   "{{secrets/shared/PIP_EXTRA_INDEX_URL_USER}}",
	// 	},
	// 	EnableElasticDisk: true,
	// 	ClusterSource:     compute.ClusterSourceJob,
	// 	InitScripts:       []compute.InitScriptInfo{},
	// 	PolicyId:          "",
	// 	DataSecurityMode:  compute.DataSecurityModeLegacySingleUser,
	// 	RuntimeEngine:     compute.RuntimeEngineStandard,
	// }
	// job := jobs.CreateJob{
	// 	JobClusters: []jobs.JobCluster{{
	// 		JobClusterKey: "MyTestGCluster",
	// 		NewCluster:    &clusterDetails,
	// 	}},
	// }
	// createCluster, err := w.Jobs.Create(context.Background(), job)
	// if err != nil {
	// 	panic(err)
	// }
	// print("%v", createCluster)
}
