package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccProvidersNoTranspile(t *testing.T) {
	t.Skip("fixme")
	ctx, w := ucwsTest(t)

	created, err := w.Providers.Create(ctx, sharing.CreateProvider{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Providers.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	_, err = w.Providers.Update(ctx, sharing.UpdateProvider{})
	require.NoError(t, err)

	_, err = w.Providers.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Providers.ListAll(ctx, sharing.ListProvidersRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccRecipients(t *testing.T) {
	ctx, w := ucwsTest(t)

	created, err := w.Recipients.Create(ctx, sharing.CreateRecipient{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Recipients.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	err = w.Recipients.Update(ctx, sharing.UpdateRecipient{
		Name:    created.Name,
		Comment: RandomName("comment "),
	})
	require.NoError(t, err)

	_, err = w.Recipients.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Recipients.ListAll(ctx, sharing.ListRecipientsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccSharesNoTranspile(t *testing.T) {
	t.Skip("expand with catalogs")
	ctx, w := ucwsTest(t)

	created, err := w.Shares.Create(ctx, sharing.CreateShare{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Shares.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	_, err = w.Shares.Update(ctx, sharing.UpdateShare{
		Name: RandomName("go-sdk-"),
		Updates: []sharing.SharedDataObjectUpdate{
			{
				Action: sharing.SharedDataObjectUpdateActionAdd,
				DataObject: &sharing.SharedDataObject{
					Name: RandomName("go-sdk-"),
				},
			},
		},
	})
	require.NoError(t, err)

	_, err = w.Shares.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Shares.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}
