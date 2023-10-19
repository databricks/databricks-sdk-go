package listing

import (
	"context"
	"errors"
	"testing"
)

// TestIterator tests the Iterator type.
func TestIterator(t *testing.T) {
	getNextPage := func(ctx context.Context, req int) ([]int, error) {
		if req >= 10 {
			return nil, nil
		}
		return []int{req, req + 1}, nil
	}
	getNextReq := func(resp []int) (int, ListingStatus) {
		if len(resp) == 0 {
			return 0, ListingStatusExhausted
		}
		return resp[len(resp)-1] + 1, ListingStatusCheckResult
	}
	getItems := func(resp []int) []int {
		return resp
	}
	it := NewIterator(0, getNextPage, getItems, getNextReq)
	for i := 0; i < 10; i++ {
		v, err := it.Next(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if v != i {
			t.Fatalf("expected %d, got %d", i, v)
		}
	}
	_, err := it.Next(context.Background())
	if !errors.Is(err, ErrNoMoreItems) {
		t.Fatalf("expected %v, got %v", ErrNoMoreItems, err)
	}
}
