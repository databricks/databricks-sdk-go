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
	getNextReq := func(resp []int) *int {
		if len(resp) == 0 {
			return nil
		}
		x := resp[len(resp)-1] + 1
		return &x
	}
	getItems := func(resp []int) []int {
		return resp
	}
	x := 0
	it := NewIterator(&x, getNextPage, getItems, getNextReq)
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

// Tests to write:
// TestIterator_GetNextPageErrors tests that errors from getNextPage are propagated.

// TestIterator_GetNextReqCheckResult tests that the iterator stops if there are no more items.

// TestIterator_GetNextReqExhausted tests that the iterator stops even if some items are returned.

// TestIterator_GetNextReqNotExhausted tests that the iterator continues if there are more items until CheckResult with no entries or Exhausted.

// TestIterator_NextReturnsNextItem tests that Next returns the next item.

// TestIterator_NextErrNoMoreItems tests that Next returns ErrNoMoreItems when there are no more items.

// TestIterator_HasNextNoNext tests that HasNext returns false when there are no more items.

// TestIterator_HasNextNext tests that HasNext returns true when there are more items.
