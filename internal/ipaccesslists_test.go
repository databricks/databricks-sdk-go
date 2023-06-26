package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccIpAccessLists(t *testing.T) {
	ctx, w := qa.WorkspaceTest(t)

	created, err := w.IpAccessLists.Create(ctx, settings.CreateIpAccessList{
		Label:       qa.RandomName("go-sdk-"),
		IpAddresses: []string{"1.0.0.0/16"},
		ListType:    settings.ListTypeBlock,
	})
	require.NoError(t, err)

	defer w.IpAccessLists.DeleteByIpAccessListId(ctx, created.IpAccessList.ListId)

	err = w.IpAccessLists.Replace(ctx, settings.ReplaceIpAccessList{
		IpAccessListId: created.IpAccessList.ListId,
		Label:          qa.RandomName("go-sdk-updated-"),
		IpAddresses:    []string{"1.0.0.0/24"},
		ListType:       settings.ListTypeBlock,
		Enabled:        false,
	})
	require.NoError(t, err)

	byId, err := w.IpAccessLists.GetByIpAccessListId(ctx, created.IpAccessList.ListId)
	require.NoError(t, err)

	all, err := w.IpAccessLists.ListAll(ctx)
	require.NoError(t, err)

	names, err := w.IpAccessLists.IpAccessListInfoLabelToListIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.IpAccessList.ListId, names[byId.IpAccessList.Label])

	byName, err := w.IpAccessLists.GetByLabel(ctx, byId.IpAccessList.Label)
	require.NoError(t, err)
	assert.Equal(t, byId.IpAccessList.ListId, byName.ListId)
}
