package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/settings/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccIpAccessLists(t *testing.T) {
	ctx := workspaceTest(t)

	IpAccessListsAPI, err := settings.NewIpAccessListsClient(nil)
	require.NoError(t, err)
	created, err := IpAccessListsAPI.Create(ctx, settings.CreateIpAccessList{
		Label:       RandomName("go-sdk-"),
		IpAddresses: []string{"1.0.0.0/16"},
		ListType:    settings.ListTypeBlock,
	})
	require.NoError(t, err)

	defer IpAccessListsAPI.DeleteByIpAccessListId(ctx, created.IpAccessList.ListId)

	err = IpAccessListsAPI.Replace(ctx, settings.ReplaceIpAccessList{
		IpAccessListId: created.IpAccessList.ListId,
		Label:          RandomName("go-sdk-updated-"),
		IpAddresses:    []string{"1.0.0.0/24"},
		ListType:       settings.ListTypeBlock,
		Enabled:        false,
	})
	require.NoError(t, err)

	byId, err := IpAccessListsAPI.GetByIpAccessListId(ctx, created.IpAccessList.ListId)
	require.NoError(t, err)

	all, err := IpAccessListsAPI.ListAll(ctx)
	require.NoError(t, err)

	names, err := IpAccessListsAPI.IpAccessListInfoLabelToListIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.IpAccessList.ListId, names[byId.IpAccessList.Label])

	byName, err := IpAccessListsAPI.GetByLabel(ctx, byId.IpAccessList.Label)
	require.NoError(t, err)
	assert.Equal(t, byId.IpAccessList.ListId, byName.ListId)
}
