package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/ipaccesslists"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccIpAccessLists(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.IpAccessLists.CreateIpAccessList(ctx, ipaccesslists.CreateIpAccessListRequest{
		Label:       RandomName("go-sdk-"),
		IpAddresses: []string{"1.0.0.0/16"},
		ListType:    ipaccesslists.ListTypeBlock,
	})
	require.NoError(t, err)

	defer w.IpAccessLists.DeleteIpAccessListByIpAccessListId(ctx, created.ListId)

	err = w.IpAccessLists.ReplaceIpAccessList(ctx, ipaccesslists.ReplaceIpAccessListRequest{
		IpAccessListId: created.ListId,
		Label:          RandomName("go-sdk-updated-"),
		IpAddresses:    []string{"1.0.0.0/24"},
		ListType:       ipaccesslists.ListTypeBlock,
		Enabled:        false,
	})
	require.NoError(t, err)

	// TODO: OpenAPI: Fetch -> Get
	list, err := w.IpAccessLists.FetchIpAccessListByIpAccessListId(ctx, created.ListId)
	require.NoError(t, err)

	all, err := w.IpAccessLists.ListIpAccessListsAll(ctx)
	require.NoError(t, err)

	names, err := w.IpAccessLists.IpAccessListInfoLabelToListIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, list.ListId, names[list.Label])

	byName, err := w.IpAccessLists.GetIpAccessListInfoByLabel(ctx, list.Label)
	require.NoError(t, err)
	assert.Equal(t, list.ListId, byName.ListId)
}
