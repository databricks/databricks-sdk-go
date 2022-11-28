package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/ipaccesslists"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccIpAccessLists(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.IpAccessLists.Create(ctx, ipaccesslists.CreateIpAccessList{
		Label:       RandomName("go-sdk-"),
		IpAddresses: []string{"1.0.0.0/16"},
		ListType:    ipaccesslists.ListTypeBlock,
	})
	require.NoError(t, err)

	defer w.IpAccessLists.DeleteByIpAccessListId(ctx, created.IpAccessList.ListId)

	err = w.IpAccessLists.Replace(ctx, ipaccesslists.ReplaceIpAccessList{
		IpAccessListId: created.IpAccessList.ListId,
		Label:          RandomName("go-sdk-updated-"),
		IpAddresses:    []string{"1.0.0.0/24"},
		ListType:       ipaccesslists.ListTypeBlock,
		Enabled:        false,
	})
	require.NoError(t, err)

	list, err := w.IpAccessLists.GetByIpAccessListId(ctx, created.IpAccessList.ListId)
	require.NoError(t, err)

	all, err := w.IpAccessLists.ListAll(ctx)
	require.NoError(t, err)

	names, err := w.IpAccessLists.IpAccessListInfoLabelToListIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, list.IpAccessList.ListId, names[list.IpAccessList.Label])

	byName, err := w.IpAccessLists.GetIpAccessListInfoByLabel(ctx, list.IpAccessList.Label)
	require.NoError(t, err)
	assert.Equal(t, list.IpAccessList.ListId, byName.ListId)
}
