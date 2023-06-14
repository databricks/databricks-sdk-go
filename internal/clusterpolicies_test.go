package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccClusterPolicyFamilies(t *testing.T) {
	ctx, w := workspaceTest(t)

	all, err := w.PolicyFamilies.ListAll(ctx, compute.ListPolicyFamiliesRequest{})
	require.NoError(t, err)

	firstFamily, err := w.PolicyFamilies.Get(ctx, compute.GetPolicyFamilyRequest{
		PolicyFamilyId: all[0].PolicyFamilyId,
	})
	require.NoError(t, err)
	assert.NotEqual(t, "", firstFamily.Name)
}

func TestAccClusterPolicies(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.ClusterPolicies.Create(ctx, compute.CreatePolicy{
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

	byName, err := w.ClusterPolicies.GetByName(ctx, policy.Name)
	require.NoError(t, err)
	assert.Equal(t, policy.PolicyId, byName.PolicyId)

	all, err := w.ClusterPolicies.ListAll(ctx, compute.ListClusterPoliciesRequest{})
	require.NoError(t, err)

	names, err := w.ClusterPolicies.PolicyNameToPolicyIdMap(ctx, compute.ListClusterPoliciesRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, policy.PolicyId, names[policy.Name])

	err = w.ClusterPolicies.Edit(ctx, compute.EditPolicy{
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
