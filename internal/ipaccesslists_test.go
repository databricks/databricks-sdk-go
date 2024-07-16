package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/stretchr/testify/assert"
)

func TestAccIpAccessLists(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.IpAccessLists.Create(ctx, settings.CreateIpAccessList{
		Label:       RandomName("go-sdk-"),
		IpAddresses: []string{"1.0.0.0/16"},
		ListType:    settings.ListTypeBlock,
	})
	assert.NoError(t, err)

	defer w.IpAccessLists.DeleteByIpAccessListId(ctx, created.IpAccessList.ListId)

	err = w.IpAccessLists.Replace(ctx, settings.ReplaceIpAccessList{
		IpAccessListId: created.IpAccessList.ListId,
		Label:          RandomName("go-sdk-updated-"),
		IpAddresses:    []string{"1.0.0.0/24"},
		ListType:       settings.ListTypeBlock,
		Enabled:        false,
	})
	assert.NoError(t, err)

	byId, err := w.IpAccessLists.GetByIpAccessListId(ctx, created.IpAccessList.ListId)
	assert.NoError(t, err)

	all, err := w.IpAccessLists.ListAll(ctx)
	assert.NoError(t, err)

	names, err := w.IpAccessLists.IpAccessListInfoLabelToListIdMap(ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.IpAccessList.ListId, names[byId.IpAccessList.Label])

	byName, err := w.IpAccessLists.GetByLabel(ctx, byId.IpAccessList.Label)
	assert.NoError(t, err)
	assert.Equal(t, byId.IpAccessList.ListId, byName.ListId)
}
