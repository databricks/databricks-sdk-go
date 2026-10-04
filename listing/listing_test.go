package listing_test

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/listing"
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

func TestToSliceN_fetchesOnlyRequiredPages(t *testing.T) {
	testCases := []struct {
		name          string
		pages         [][]int
		limit         int64
		want          []int
		wantRequests  int
		wantRemaining []int
	}{
		{
			name: "within page", pages: [][]int{{1, 2}, {3}}, limit: 1,
			want: []int{1}, wantRequests: 1, wantRemaining: []int{2, 3},
		},
		{
			name: "first page boundary", pages: [][]int{{1, 2}, {3}}, limit: 2,
			want: []int{1, 2}, wantRequests: 1, wantRemaining: []int{3},
		},
		{
			name: "later page boundary", pages: [][]int{{1}, {2}, {3}}, limit: 2,
			want: []int{1, 2}, wantRequests: 2, wantRemaining: []int{3},
		},
		{
			name: "empty page after limit", pages: [][]int{{1}, {}, {2}}, limit: 1,
			want: []int{1}, wantRequests: 1, wantRemaining: []int{2},
		},
		{
			name: "empty page before limit", pages: [][]int{{}, {1}, {2}}, limit: 1,
			want: []int{1}, wantRequests: 2, wantRemaining: []int{2},
		},
		{
			name: "zero limit", pages: [][]int{{1}, {2}}, limit: 0,
			want: []int{1, 2}, wantRequests: 2,
		},
		{
			name: "limit exceeds total", pages: [][]int{{1}, {2}}, limit: 3,
			want: []int{1, 2}, wantRequests: 2,
		},
		{
			name: "negative limit", pages: [][]int{{1}, {2}}, limit: -1,
			wantRemaining: []int{1, 2},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestCount := 0
			firstPage := 0
			iterator := listing.NewIterator(&firstPage,
				func(ctx context.Context, page int) (int, error) {
					requestCount++
					return page, nil
				},
				func(page int) []int { return tc.pages[page] },
				func(page int) *int {
					nextPage := page + 1
					if nextPage == len(tc.pages) {
						return nil
					}
					return &nextPage
				},
			)

			got, err := listing.ToSliceN(context.Background(), iterator, tc.limit)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantRequests, requestCount)

			remaining, err := listing.ToSlice(context.Background(), iterator)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantRemaining, remaining)
		})
	}
}

func TestToSliceN_doesNotCacheErrorBeyondLimit(t *testing.T) {
	pageErr := errors.New("next page failed")
	allowNextPage := false
	firstPage := 0
	nextPage := 1
	iterator := listing.NewIterator(&firstPage,
		func(ctx context.Context, page int) (int, error) {
			if page == nextPage && !allowNextPage {
				return 0, pageErr
			}
			return page, nil
		},
		func(page int) []int { return []int{page + 1} },
		func(page int) *int {
			if page == firstPage {
				return &nextPage
			}
			return nil
		},
	)

	got, err := listing.ToSliceN(context.Background(), iterator, 1)
	assert.NoError(t, err)
	assert.Equal(t, []int{1}, got)

	allowNextPage = true
	remaining, err := listing.ToSlice(context.Background(), iterator)
	assert.NoError(t, err)
	assert.Equal(t, []int{2}, remaining)
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

func TestLimitIterator(t *testing.T) {
	t.Run("caps results when limit less than total", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
			{req: "page3", page: map[string][]int{"page": {5, 6}}},
		}
		iterator := makeIterator(rrs)
		limited := listing.NewLimitIterator[int](iterator, 3)

		items, err := listing.ToSlice(context.Background(), limited)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, items)
	})

	t.Run("limit greater than total returns all items", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
		}
		iterator := makeIterator(rrs)
		limited := listing.NewLimitIterator[int](iterator, 100)

		items, err := listing.ToSlice(context.Background(), limited)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4}, items)
	})

	t.Run("limit of one", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
		}
		iterator := makeIterator(rrs)
		limited := listing.NewLimitIterator[int](iterator, 1)

		items, err := listing.ToSlice(context.Background(), limited)
		assert.NoError(t, err)
		assert.Equal(t, []int{1}, items)
	})

	t.Run("zero limit is no-op", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
		}
		iterator := makeIterator(rrs)
		limited := listing.NewLimitIterator[int](iterator, 0)

		items, err := listing.ToSlice(context.Background(), limited)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4}, items)
	})

	t.Run("negative limit is no-op", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
		}
		iterator := makeIterator(rrs)
		limited := listing.NewLimitIterator[int](iterator, -5)

		items, err := listing.ToSlice(context.Background(), limited)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4}, items)
	})

	t.Run("Next returns ErrNoMoreItems after limit exhausted", func(t *testing.T) {
		iterator := listing.SliceIterator[int]([]int{1, 2, 3, 4, 5})
		limited := listing.NewLimitIterator[int](&iterator, 2)

		item1, err := limited.Next(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 1, item1)

		item2, err := limited.Next(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 2, item2)

		_, err = limited.Next(context.Background())
		assert.ErrorIs(t, err, listing.ErrNoMoreItems)
	})

	t.Run("error propagation from inner iterator", func(t *testing.T) {
		expectedErr := errors.New("some error")
		iterator := listing.NewIterator[struct{}, []int, int](&struct{}{},
			func(ctx context.Context, req struct{}) ([]int, error) {
				return nil, expectedErr
			},
			nil,
			nil,
		)
		limited := listing.NewLimitIterator[int](iterator, 5)

		assert.True(t, limited.HasNext(context.Background()))
		_, err := limited.Next(context.Background())
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("standard HasNext/Next loop", func(t *testing.T) {
		rrs := []requestResponse{
			{req: "page1", page: map[string][]int{"page": {1, 2}}},
			{req: "page2", page: map[string][]int{"page": {3, 4}}},
			{req: "page3", page: map[string][]int{"page": {5, 6}}},
		}
		iterator := makeIterator(rrs)
		limited := listing.NewLimitIterator[int](iterator, 4)

		values := make([]int, 0)
		for limited.HasNext(context.Background()) {
			v, err := limited.Next(context.Background())
			assert.NoError(t, err)
			values = append(values, v)
		}
		assert.Equal(t, []int{1, 2, 3, 4}, values)
	})

	t.Run("HasNext is idempotent", func(t *testing.T) {
		iterator := listing.SliceIterator[int]([]int{1, 2, 3})
		limited := listing.NewLimitIterator[int](&iterator, 2)

		assert.True(t, limited.HasNext(context.Background()))
		assert.True(t, limited.HasNext(context.Background()))

		v, err := limited.Next(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 1, v)
	})
}
