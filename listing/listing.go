package listing

import (
	"context"
	"errors"
)

// Use struct{} for Req to indicate one-shot iterator.
type Iterator[Req, Resp, T any] struct {
	// nextReq is the request to be used to fetch the next page. If nil, then
	// there is no next page to fetch.
	nextReq *Req

	// getNextPage fetches the next page of items, returning the deserialized
	// response and error.
	getNextPage func(context.Context, Req) (Resp, error)

	// getItems selects the items being iterated over from the response.
	getItems func(Resp) []T

	// getNextReq is used to get the next request to be used in the next page.
	// If nil, then there is no next page to fetch.
	getNextReq func(Resp) *Req

	// currentPage is the current page of items.
	currentPage []T

	// currentPageIdx is the index of the next item from currentPage to return.
	currentPageIdx int

	// lastErr is the last error returned by getNextPage.
	lastErr error
}

var ErrNoMoreItems = errors.New("no more items")

func NewIterator[Req, Resp, T any](
	nextReq *Req,
	getNextPage func(context.Context, Req) (Resp, error),
	getItems func(Resp) []T,
	getNextReq func(Resp) *Req,
) *Iterator[Req, Resp, T] {
	return &Iterator[Req, Resp, T]{
		nextReq:     nextReq,
		getNextPage: getNextPage,
		getItems:    getItems,
		getNextReq:  getNextReq,
	}
}

func (i *Iterator[Req, Resp, T]) loadNextPageIfNeeded(ctx context.Context) error {
	if i.currentPageIdx < len(i.currentPage) {
		return nil
	}
	if i.nextReq == nil {
		i.currentPage = nil
		i.currentPageIdx = 0
		return nil
	}
	if i.lastErr != nil {
		return i.lastErr
	}

	// Keep loading pages while we have a next request and the current page is
	// empty.
	i.currentPage = nil
	for i.nextReq != nil && len(i.currentPage) == 0 {
		resp, err := i.getNextPage(ctx, *i.nextReq)
		i.lastErr = err
		if err != nil {
			return err
		}
		items := i.getItems(resp)
		i.currentPage = items
		i.currentPageIdx = 0
		i.nextReq = i.getNextReq(resp)
	}
	return nil
}

func (i *Iterator[Req, Resp, T]) Next(ctx context.Context) (T, error) {
	var t T
	err := i.loadNextPageIfNeeded(ctx)
	if err != nil {
		return t, err
	}
	if i.currentPageIdx >= len(i.currentPage) {
		return t, ErrNoMoreItems
	}
	item := i.currentPage[i.currentPageIdx]
	i.currentPageIdx++
	return item, nil
}

func (i *Iterator[Req, Resp, T]) HasNext(ctx context.Context) bool {
	err := i.loadNextPageIfNeeded(ctx)
	if err != nil {
		return true
	}
	return i.currentPageIdx < len(i.currentPage)
}

type DedupeIterator[Req, Resp any, T comparable] struct {
	it      *Iterator[Req, Resp, T]
	seen    map[T]struct{}
	current *T
}

func NewDedupeIterator[Req, Resp any, T comparable](
	it *Iterator[Req, Resp, T]) *DedupeIterator[Req, Resp, T] {
	return &DedupeIterator[Req, Resp, T]{
		it:   it,
		seen: make(map[T]struct{}),
	}
}

func (i *DedupeIterator[Req, Resp, T]) Next(ctx context.Context) (T, error) {
	if i.current != nil {
		t := *i.current
		i.current = nil
		return t, nil
	}
	for {
		t, err := i.it.Next(ctx)
		if err != nil {
			return t, err
		}
		if _, ok := i.seen[t]; !ok {
			i.seen[t] = struct{}{}
			return t, nil
		}
	}
}

func (i *DedupeIterator[Req, Resp, T]) HasNext(ctx context.Context) bool {
	if i.current != nil {
		return true
	}
	for {
		t, err := i.it.Next(ctx)
		if errors.Is(err, ErrNoMoreItems) {
			return false
		}
		if err != nil {
			return true
		}
		if _, ok := i.seen[t]; !ok {
			i.current = &t
			return true
		}
	}
}
