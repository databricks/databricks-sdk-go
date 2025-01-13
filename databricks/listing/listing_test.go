package listing_test

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/stretchr/testify/assert"
)

type requestResponse struct {
	req  string
	page map[string][]int
}

func makeIterator(reqsAndPages []requestResponse) listing.Iterator[int] {
	reqs := make([]string, len(reqsAndPages))
	pages := make(map[string]map[string][]int)
	for i, reqAndPage := range reqsAndPages {
		reqs[i] = reqAndPage.req
		pages[reqAndPage.req] = reqAndPage.page
	}
	nextReq := reqs[0]
	iterator := listing.NewIterator(&nextReq,
		func(ctx context.Context, req string) (map[string][]int, error) {
			return pages[req], nil
		},
		func(resp map[string][]int) []int {
			return resp["page"]
		},
		func(resp map[string][]int) *string {
			for i, r := range reqs {
				if r == nextReq && i+1 < len(reqs) {
					nextReq = reqs[i+1]
					return &nextReq
				}
			}
			return nil
		},
	)
	return iterator
}

func TestIterator(t *testing.T) {
	t.Run("basic iteration", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
			{req: "page3", page: map[string][]int{"page": {5, 6}}},
		}
		iterator := makeIterator(rrs)

		for i := 1; iterator.HasNext(context.Background()); i++ {
			item, err := iterator.Next(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, i, item)
		}

		_, err := iterator.Next(context.Background())
		assert.ErrorIs(t, err, listing.ErrNoMoreItems)
	})

	t.Run("error propagation", func(t *testing.T) {
		expectedErr := errors.New("some error")
		iterator := listing.NewIterator[struct{}, []int, struct{}](&struct{}{},
			func(ctx context.Context, req struct{}) ([]int, error) {
				return nil, expectedErr
			},
			nil,
			nil,
		)

		_, err := iterator.Next(context.Background())
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("error propagation from HasNext", func(t *testing.T) {
		expectedErr := errors.New("some error")
		iterator := listing.NewIterator[struct{}, []int, struct{}](&struct{}{},
			func(ctx context.Context, req struct{}) ([]int, error) {
				return nil, expectedErr
			},
			nil,
			nil,
		)

		b := iterator.HasNext(context.Background())
		assert.True(t, b)
		// Error from HasNext is stored and returned on call to Next
		_, err := iterator.Next(context.Background())
		assert.ErrorIs(t, err, expectedErr)
		// Error is cached
		_, err2 := iterator.Next(context.Background())
		assert.Equal(t, err, err2)
	})

	t.Run("iteration with empty page in the middle", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {}}},
			{req: "page3", page: map[string][]int{"page": {3, 4}}},
		}
		iterator := makeIterator(rrs)

		for i := 1; i <= 4; i++ {
			item, err := iterator.Next(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, i, item)
		}

		_, err := iterator.Next(context.Background())
		assert.ErrorIs(t, err, listing.ErrNoMoreItems)
	})

	t.Run("iteration with empty page at the end", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
			{req: "page3", page: map[string][]int{"page": {}}},
		}
		iterator := makeIterator(rrs)

		for i := 1; i <= 4; i++ {
			item, err := iterator.Next(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, i, item)
		}

		_, err := iterator.Next(context.Background())
		assert.ErrorIs(t, err, listing.ErrNoMoreItems)
	})

	t.Run("ToSlice returns all items", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
			{req: "page3", page: map[string][]int{"page": {5, 6}}},
		}
		iterator := makeIterator(rrs)

		items, err := listing.ToSlice(context.Background(), iterator)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, items)
	})

	t.Run("ToSliceN returns the first N items", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
			{req: "page3", page: map[string][]int{"page": {5, 6}}},
		}
		iterator := makeIterator(rrs)

		items, err := listing.ToSliceN(context.Background(), iterator, 3)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, items)
	})
}

func TestDedupeIterator(t *testing.T) {
	t.Run("basic iteration", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4, 4}}},
			{req: "page3", page: map[string][]int{"page": {5, 5, 6}}},
		}
		iterator := makeIterator(rrs)
		dedupeIterator := listing.NewDedupeIterator[int, int](
			iterator, func(a int) int { return a })

		for i := 1; i <= 6; i++ {
			item, err := dedupeIterator.Next(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, i, item)
		}

		_, err := dedupeIterator.Next(context.Background())
		assert.ErrorIs(t, err, listing.ErrNoMoreItems)
	})

	t.Run("HasNext caches appropriately", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4, 4}}},
			{req: "page3", page: map[string][]int{"page": {5, 5, 6}}},
		}
		iterator := makeIterator(rrs)
		dedupeIterator := listing.NewDedupeIterator[int, int](
			iterator, func(a int) int { return a })

		values := make([]int, 0)
		for dedupeIterator.HasNext(context.Background()) {
			assert.True(t, dedupeIterator.HasNext(context.Background()))
			v, err := dedupeIterator.Next(context.Background())
			values = append(values, v)
			assert.NoError(t, err)
		}
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, values)
	})

	t.Run("HasNext returns true when there is an error", func(t *testing.T) {
		expectedErr := errors.New("some error")
		iterator := listing.NewIterator[struct{}, []int, int](&struct{}{},
			func(ctx context.Context, req struct{}) ([]int, error) {
				return nil, expectedErr
			},
			nil,
			nil,
		)
		dedupeIterator := listing.NewDedupeIterator[int, int](
			iterator, func(a int) int { return a })
		assert.True(t, dedupeIterator.HasNext(context.Background()))
		_, err := dedupeIterator.Next(context.Background())
		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestSliceIterator(t *testing.T) {
	t.Run("basic iteration", func(t *testing.T) {
		iterator := listing.SliceIterator[int]([]int{1, 2, 3, 4, 5})

		for i := 1; i <= 5; i++ {
			assert.True(t, iterator.HasNext(context.Background()))
			item, err := iterator.Next(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, i, item)
		}

		assert.False(t, iterator.HasNext(context.Background()))
		_, err := iterator.Next(context.Background())
		assert.ErrorIs(t, err, listing.ErrNoMoreItems)
	})

	t.Run("ToSlice returns all items", func(t *testing.T) {
		iterator := listing.SliceIterator[int]([]int{1, 2, 3, 4, 5})

		items, err := listing.ToSlice[int](context.Background(), &iterator)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, items)
	})

	t.Run("ToSliceN returns the first N items", func(t *testing.T) {
		iterator := listing.SliceIterator[int]([]int{1, 2, 3, 4, 5})

		items, err := listing.ToSliceN[int](context.Background(), &iterator, 3)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, items)
	})

}
