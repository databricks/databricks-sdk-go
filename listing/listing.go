package listing

import (
	"context"
	"errors"
)

type Iterator[T any] interface {
	HasNext(context.Context) bool
	Next(context.Context) (T, error)
}

func ToSlice[T any](ctx context.Context, it Iterator[T]) ([]T, error) {
	var items []T
	for it.HasNext(ctx) {
		item, err := it.Next(ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// Use struct{} for Req to indicate one-shot iterator.
type IteratorImpl[Req, Resp, T any] struct {
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
) *IteratorImpl[Req, Resp, T] {
	return &IteratorImpl[Req, Resp, T]{
		nextReq:     nextReq,
		getNextPage: getNextPage,
		getItems:    getItems,
		getNextReq:  getNextReq,
	}
}

func (i *IteratorImpl[Req, Resp, T]) loadNextPageIfNeeded(ctx context.Context) error {
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

func (i *IteratorImpl[Req, Resp, T]) Next(ctx context.Context) (T, error) {
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

func (i *IteratorImpl[Req, Resp, T]) HasNext(ctx context.Context) bool {
	err := i.loadNextPageIfNeeded(ctx)
	if err != nil {
		return true
	}
	return i.currentPageIdx < len(i.currentPage)
}

type DedupeIteratorImpl[T any, Id comparable] struct {
	it      Iterator[T]
	getId   func(T) Id
	seen    map[Id]struct{}
	current *T
}

func NewDedupeIterator[T any, Id comparable](it Iterator[T], getId func(T) Id) *DedupeIteratorImpl[T, Id] {
	return &DedupeIteratorImpl[T, Id]{
		it:    it,
		getId: getId,
		seen:  make(map[Id]struct{}),
	}
}

func (i *DedupeIteratorImpl[T, Id]) Next(ctx context.Context) (T, error) {
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
		id := i.getId(t)
		if _, ok := i.seen[id]; !ok {
			i.seen[id] = struct{}{}
			return t, nil
		}
	}
}

func (i *DedupeIteratorImpl[T, Id]) HasNext(ctx context.Context) bool {
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
		id := i.getId(t)
		if _, ok := i.seen[id]; !ok {
			i.current = &t
			i.seen[id] = struct{}{}
			return true
		}
	}
}
