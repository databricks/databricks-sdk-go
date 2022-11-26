package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/clusterpolicies"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccClusterPolicies(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.ClusterPolicies.Create(ctx, clusterpolicies.CreatePolicy{
		Name: RandomName("go-sdk-"),
		Definition: `{
			"spark_conf.spark.databricks.delta.preview.enabled": {
				"type": "fixed",
				"value": true
			}
		}`,
	})
	require.NoError(t, err)
	defer w.ClusterPolicies.DeleteByPolicyId(ctx, created.PolicyId)

	policy, err := w.ClusterPolicies.GetByPolicyId(ctx, created.PolicyId)
	require.NoError(t, err)

	byName, err := w.ClusterPolicies.GetPolicyByName(ctx, policy.Name)
	require.NoError(t, err)
	assert.Equal(t, policy.Name, byName.Name)

	names, err := w.ClusterPolicies.PolicyNameToPolicyIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, policy.PolicyId, names[policy.Name])

	err = w.ClusterPolicies.Edit(ctx, clusterpolicies.EditPolicy{
		PolicyId: policy.PolicyId,
		Name:     policy.Name,
		Definition: `{
			"spark_conf.spark.databricks.delta.preview.enabled": {
				"type": "fixed",
				"value": false
			}
		}`,
	})
	require.NoError(t, err)
}
