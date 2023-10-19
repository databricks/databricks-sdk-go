package listing

import (
	"context"
	"errors"
)

// Use struct{} for Req to indicate one-shot iterator.
type Iterator[Req, Resp, T any] struct {
	nextReq        Req
	getNextPage    func(context.Context, Req) (Resp, error)
	getItems       func(Resp) []T
	getNextReq     func(Resp) Req
	currentPage    []T
	currentPageIdx int
	isExhausted    bool
}

var ErrNoMoreItems = errors.New("no more items")

func NewIterator[Req, Resp, T any](
	nextReq Req,
	getNextPage func(context.Context, Req) (Resp, error),
	getItems func(Resp) []T,
	getNextReq func(Resp) Req,
) *Iterator[Req, Resp, T] {
	return &Iterator[Req, Resp, T]{
		nextReq:     nextReq,
		getNextPage: getNextPage,
		getItems:    getItems,
		getNextReq:  getNextReq,
	}
}

func (i *Iterator[Req, Resp, T]) loadNextPageIfNeeded(ctx context.Context) (bool, error) {
	if i.currentPageIdx < len(i.currentPage) {
		return true, nil
	}

	if i.isExhausted {
		return false, ErrNoMoreItems
	}
	resp, err := i.getNextPage(ctx, i.nextReq)
	if err != nil {
		return false, err
	}
	items := i.getItems(resp)
	i.currentPage = items
	i.currentPageIdx = 0
	i.nextReq = i.getNextReq(resp)
	if !i.isExhausted && len(items) == 0 {
		i.isExhausted = true
	}
	return !i.isExhausted, err
}

func (i *Iterator[Req, Resp, T]) Next(ctx context.Context) (T, error) {
	var t T
	ok, err := i.loadNextPageIfNeeded(ctx)
	if err != nil {
		return t, err
	}
	if !ok {
		return t, ErrNoMoreItems
	}
	item := i.currentPage[i.currentPageIdx]
	i.currentPageIdx++
	return item, nil
}

func (i *Iterator[Req, Resp, T]) HasNext(ctx context.Context) (bool, error) {
	return i.loadNextPageIfNeeded(ctx)
}
