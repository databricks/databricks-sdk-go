package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatNames(t *testing.T) {
	for _, v := range []struct {
		service, method, flat string
	}{
		{"secrets", "delete_acl", "delete_secret_acl"},
		{"secrets", "list_scopes", "list_secret_scopes"},
		{"workspace_conf", "get_status", "get_workspace_conf_status"},
		{"statement_execution", "execute_statement", "execute_statement"},
		{"statement_execution", "get_statement_result_chunk_n", "get_statement_execution_result_chunk_n"},

		// TBD:
		{"current_user", "me", "me"},
		{"warehouses", "get_workspace_warehouse_config", "get_warehouse_workspace_config"},
		{"libraries", "install", "install_cluster_library"},
		{"policy_families", "get", "get_cluster_policy_family"},
		{"workspace", "get_status", "get_notebook_status"},
		{"model_registry", "create_comment", "create_model_comment"},
		{"token_management", "create_obo_token", "create_obo_token"},
	} {
		m := &Method{
			Named: Named{
				Name: v.method,
			},
			Service: &Service{
				Named: Named{
					Name: v.service,
				},
			},
		}
		assert.Equal(t, v.flat, m.AsFlat().SnakeName())
	}
}

func TestCmdletNames(t *testing.T) {
	for _, v := range []struct {
		service, method, cmdlet string
	}{
		{"secrets", "delete_acl", "Delete-DatabricksSecretAcl"},
		{"secrets", "list_scopes", "List-DatabricksSecretScopes"},
		{"workspace_conf", "get_status", "Get-DatabricksWorkspaceConfStatus"},

		// TBD:
		{"current_user", "me", "Me-Databricks"},
		{"warehouses", "get_workspace_warehouse_config", "Get-DatabricksWarehouseWorkspaceConfig"},
		{"libraries", "install", "Install-DatabricksClusterLibrary"},
		{"policy_families", "get", "Get-DatabricksClusterPolicyFamily"},
		{"workspace", "get_status", "Get-DatabricksNotebookStatus"},
		{"model_registry", "create_comment", "Create-DatabricksModelComment"},
		{"token_management", "create_obo_token", "Create-DatabricksOboToken"},
	} {
		m := &Method{
			Named: Named{
				Name: v.method,
			},
			Service: &Service{
				Named: Named{
					Name: v.service,
				},
			},
		}
		assert.Equal(t, v.cmdlet, m.CmdletName())
	}
}
