package integration

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/catalog/v2"
	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccConnections(t *testing.T) {
	ctx := ucwsTest(t)

	ConnectionsAPI, err := catalog.NewConnectionsClient(nil)
	require.NoError(t, err)
	connCreate, err := ConnectionsAPI.Create(ctx, catalog.CreateConnection{
		Comment:        "Go SDK Acceptance Test Connection",
		ConnectionType: catalog.ConnectionTypeDatabricks,
		Name:           RandomName("go-sdk-connection-"),
		Options: map[string]string{
			"host":                fmt.Sprintf("%s-fake-workspace.cloud.databricks.com", RandomName("go-sdk-connection-")),
			"httpPath":            fmt.Sprintf("/sql/1.0/warehouses/%s", RandomName("go-sdk-connection-")),
			"personalAccessToken": RandomName("go-sdk-connection-"),
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := ConnectionsAPI.Delete(ctx, catalog.DeleteConnectionRequest{Name: connCreate.Name})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	connUpdate, err := ConnectionsAPI.Update(ctx, catalog.UpdateConnection{
		Name: connCreate.Name,
		Options: map[string]string{
			"host":                fmt.Sprintf("%s-fake-workspace.cloud.databricks.com", RandomName("go-sdk-connection-")),
			"httpPath":            fmt.Sprintf("/sql/1.0/warehouses/%s", RandomName("go-sdk-connection-")),
			"personalAccessToken": RandomName("go-sdk-connection-"),
		},
	})
	require.NoError(t, err)

	conn, err := ConnectionsAPI.Get(ctx, catalog.GetConnectionRequest{Name: connUpdate.Name})
	require.NoError(t, err)
	assert.Equal(t, conn.Options, connUpdate.Options)

	connList, err := ConnectionsAPI.ListAll(ctx, catalog.ListConnectionsRequest{})
	require.NoError(t, err)
	assert.True(t, len(connList) >= 1)
}
