package internal

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccConnections(t *testing.T) {
	ctx, w := ucwsTest(t)

	connCreate, err := w.Connections.Create(ctx, catalog.CreateConnection{
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
		err := w.Connections.Delete(ctx, catalog.DeleteConnectionRequest{Name: connCreate.Name})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	connUpdate, err := w.Connections.Update(ctx, catalog.UpdateConnection{
		Name: connCreate.Name,
		Options: map[string]string{
			"host":                fmt.Sprintf("%s-fake-workspace.cloud.databricks.com", RandomName("go-sdk-connection-")),
			"httpPath":            fmt.Sprintf("/sql/1.0/warehouses/%s", RandomName("go-sdk-connection-")),
			"personalAccessToken": RandomName("go-sdk-connection-"),
		},
	})
	require.NoError(t, err)

	conn, err := w.Connections.Get(ctx, catalog.GetConnectionRequest{Name: connUpdate.Name})
	require.NoError(t, err)
	assert.Equal(t, conn.Options, connUpdate.Options)

	connList, err := w.Connections.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(connList) >= 1)
}
