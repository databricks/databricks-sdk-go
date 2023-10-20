package listing_test

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	t.Run("basic iteration", func(t *testing.T) {
		reqs := []string{"page1", "page2", "page3"}
		pages := map[string]map[string][]int{
			"page1": {"page": {1, 2}},
			"page2": {"page": {3, 4}},
			"page3": {"page": {5, 6}},
		}
		nextReq := "page1"

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
		reqs := []string{"page1", "page2", "page3"}
		pages := map[string]map[string][]int{
			"page1": {"page": {1, 2}},
			"page2": {"page": {}},
			"page3": {"page": {3, 4}},
		}
		nextReq := "page1"

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

		for i := 1; i <= 4; i++ {
			item, err := iterator.Next(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, i, item)
		}

		_, err := iterator.Next(context.Background())
		assert.ErrorIs(t, err, listing.ErrNoMoreItems)
	})
}

func TestDedupeIterator(t *testing.T) {
	t.Run("basic iteration", func(t *testing.T) {
		reqs := []string{"page1", "page2", "page3"}
		pages := map[string][]int{
			"page1": {1, 2, 2},
			"page2": {3, 4, 4},
			"page3": {5, 5, 6},
		}
		nextReq := "page1"

		iterator := listing.NewIterator(&nextReq,
			func(ctx context.Context, req string) ([]int, error) {
				return pages[req], nil
			},
			func(resp []int) []int {
				return resp
			},
			func(resp []int) *string {
				for i, r := range reqs {
					if r == nextReq && i+1 < len(reqs) {
						nextReq = reqs[i+1]
						return &nextReq
					}
				}
				return nil
			},
		)

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
		reqs := []string{"page1", "page2", "page3"}
		pages := map[string][]int{
			"page1": {1, 2, 2},
			"page2": {3, 4, 4},
			"page3": {5, 5, 6},
		}
		nextReq := "page1"

		iterator := listing.NewIterator(&nextReq,
			func(ctx context.Context, req string) ([]int, error) {
				return pages[req], nil
			},
			func(resp []int) []int {
				return resp
			},
			func(resp []int) *string {
				for i, r := range reqs {
					if r == nextReq && i+1 < len(reqs) {
						nextReq = reqs[i+1]
						return &nextReq
					}
				}
				return nil
			},
		)

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

	})
}
