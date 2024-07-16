package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
)

func TestAccClusterPolicyFamilies(t *testing.T) {
	ctx, w := workspaceTest(t)

	all, err := w.PolicyFamilies.ListAll(ctx, compute.ListPolicyFamiliesRequest{})
	assert.NoError(t, err)

	firstFamily, err := w.PolicyFamilies.Get(ctx, compute.GetPolicyFamilyRequest{
		PolicyFamilyId: all[0].PolicyFamilyId,
	})
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	defer w.ClusterPolicies.DeleteByPolicyId(ctx, created.PolicyId)

	policy, err := w.ClusterPolicies.GetByPolicyId(ctx, created.PolicyId)
	assert.NoError(t, err)

	byName, err := w.ClusterPolicies.GetByName(ctx, policy.Name)
	assert.NoError(t, err)
	assert.Equal(t, policy.PolicyId, byName.PolicyId)

	all, err := w.ClusterPolicies.ListAll(ctx, compute.ListClusterPoliciesRequest{})
	assert.NoError(t, err)

	names, err := w.ClusterPolicies.PolicyNameToPolicyIdMap(ctx, compute.ListClusterPoliciesRequest{})
	assert.NoError(t, err)
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
	assert.NoError(t, err)
}
