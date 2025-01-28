package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/compute/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccClusterPolicyFamilies(t *testing.T) {
	ctx := workspaceTest(t)

	PolicyFamiliesAPI, err := compute.NewPolicyFamiliesClient(nil)
	require.NoError(t, err)
	all, err := PolicyFamiliesAPI.ListAll(ctx, compute.ListPolicyFamiliesRequest{})
	require.NoError(t, err)

	firstFamily, err := PolicyFamiliesAPI.Get(ctx, compute.GetPolicyFamilyRequest{
		PolicyFamilyId: all[0].PolicyFamilyId,
	})
	require.NoError(t, err)
	assert.NotEqual(t, "", firstFamily.Name)
}

func TestAccClusterPolicies(t *testing.T) {
	ctx := workspaceTest(t)

	ClusterPoliciesAPI, err := compute.NewClusterPoliciesClient(nil)
	require.NoError(t, err)
	created, err := ClusterPoliciesAPI.Create(ctx, compute.CreatePolicy{
		Name: RandomName("go-sdk-"),
		Definition: `{
			"spark_conf.spark.databricks.delta.preview.enabled": {
				"type": "fixed",
				"value": true
			}
		}`,
	})
	require.NoError(t, err)
	defer ClusterPoliciesAPI.DeleteByPolicyId(ctx, created.PolicyId)

	policy, err := ClusterPoliciesAPI.GetByPolicyId(ctx, created.PolicyId)
	require.NoError(t, err)

	byName, err := ClusterPoliciesAPI.GetByName(ctx, policy.Name)
	require.NoError(t, err)
	assert.Equal(t, policy.PolicyId, byName.PolicyId)

	all, err := ClusterPoliciesAPI.ListAll(ctx, compute.ListClusterPoliciesRequest{})
	require.NoError(t, err)

	names, err := ClusterPoliciesAPI.PolicyNameToPolicyIdMap(ctx, compute.ListClusterPoliciesRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, policy.PolicyId, names[policy.Name])

	err = ClusterPoliciesAPI.Edit(ctx, compute.EditPolicy{
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
