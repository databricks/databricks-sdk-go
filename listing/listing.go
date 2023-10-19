package listing

import (
	"context"
	"errors"
)

type ListingStatus int

const (
	// ListingStatusExhausted indicates that the iterator is exhausted and
	// there are no more items to be fetched.
	ListingStatusExhausted ListingStatus = iota
	// ListingStatusNotExhausted indicates that the iterator is not exhausted
	// and there are more items to be fetched.
	ListingStatusNotExhausted
	// ListingStatusCheckResult indicates that the iterator is exhausted if and
	// only if there are no results returned.
	ListingStatusCheckResult
)

// Use struct{} for Req to indicate one-shot iterator.
type Iterator[Req, Resp, T any] struct {
	// nextReq is the request to be used to fetch the next page.
	nextReq Req

	// getNextPage fetches the next page of items, returning the deserialized
	// response and error.
	getNextPage func(context.Context, Req) (Resp, error)

	// getItems selects the items being iterated over from the response.
	getItems func(Resp) []T

	// getNextReq is used to get the next request to be used in the next page.
	// The ListingStatus value indicates whether the client can stop iterating.
	getNextReq func(Resp) (Req, ListingStatus)

	// currentPage is the current page of items.
	currentPage []T

	// currentPageIdx is the index of the next item from currentPage to return.
	currentPageIdx int

	// isExhausted indicates whether there are no more items to be fetched.
	isExhausted bool
}

var ErrNoMoreItems = errors.New("no more items")

func NewIterator[Req, Resp, T any](
	nextReq Req,
	getNextPage func(context.Context, Req) (Resp, error),
	getItems func(Resp) []T,
	getNextReq func(Resp) (Req, ListingStatus),
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
	if i.isExhausted {
		return ErrNoMoreItems
	}
	resp, err := i.getNextPage(ctx, i.nextReq)
	if err != nil {
		return err
	}
	items := i.getItems(resp)
	i.currentPage = items
	i.currentPageIdx = 0
	var status ListingStatus
	if i.getNextReq == nil {
		status = ListingStatusExhausted
	} else {
		i.nextReq, status = i.getNextReq(resp)
	}
	if !i.isExhausted && (status == ListingStatusExhausted || (status == ListingStatusCheckResult && len(items) == 0)) {
		i.isExhausted = true
	}
	return err
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

func (i *Iterator[Req, Resp, T]) HasNext(ctx context.Context) (bool, error) {
	err := i.loadNextPageIfNeeded(ctx)
	if err != nil {
		return false, err
	}
	return i.currentPageIdx < len(i.currentPage), nil
}
