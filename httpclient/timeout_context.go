package httpclient

import (
	"context"
	"io"
	"sync"
	"time"
)

type timeoutContext struct {
	ctx    context.Context
	cancel context.CancelCauseFunc

	// Timeout is constant.
	// Deadline is updated when Mark function is called.
	timeout  time.Duration
	deadline time.Time

	// Protect against concurrent deadline reads/writes.
	lock sync.Mutex
}

type TimeoutTicker interface {
	Tick()
	Cancel()
}

func newTimeoutContext(ctx context.Context, timeout time.Duration) (context.Context, TimeoutTicker) {
	ctx, cancel := context.WithCancelCause(ctx)
	t := &timeoutContext{
		ctx:      ctx,
		cancel:   cancel,
		timeout:  timeout,
		deadline: time.Now().Add(timeout),
	}

	// Start goroutine to cancel the context when the deadline is reached.
	go t.run()
	return ctx, t
}

// Tick updates the deadline to the current time plus the timeout.
func (t *timeoutContext) Tick() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.deadline = time.Now().Add(t.timeout)
}

// Cancel cancels the context.
func (t *timeoutContext) Cancel() {
	t.cancel(context.Canceled)
}

// Deadline returns the current deadline.
func (t *timeoutContext) Deadline() time.Time {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.deadline
}

func (t *timeoutContext) run() {
	for {
		ttl := time.Until(t.Deadline())
		if ttl <= 0 {
			t.cancel(context.DeadlineExceeded)
			return
		}
		select {
		case <-time.After(ttl):
			// Check if the deadline has been updated
			continue
		case <-t.ctx.Done():
			return
		}
	}
}

type timeoutTickReader struct {
	t TimeoutTicker
	r io.Reader
}

func (t timeoutTickReader) Read(p []byte) (n int, err error) {
	defer t.t.Tick()
	return t.r.Read(p)
}

// tickingReader returns a reader that ticks the timeout ticker on each read.
// This is used when reading the request body and writing it to the server.
func tickingReader(t TimeoutTicker, r io.Reader) io.Reader {
	if r == nil {
		return nil
	}
	return timeoutTickReader{t, r}
}

type timeoutTickReadCloser struct {
	t TimeoutTicker
	r io.ReadCloser
}

// tickingReadCloser returns a read closer that ticks the timeout ticker on each read.
// It also cancels the ticker when the read closer is closed.
// This is used specifically for reading the response body.
func tickingReadCloser(t TimeoutTicker, r io.ReadCloser) io.ReadCloser {
	if r == nil {
		return nil
	}
	return timeoutTickReadCloser{t, r}
}

func (t timeoutTickReadCloser) Read(p []byte) (n int, err error) {
	defer t.t.Tick()
	return t.r.Read(p)
}

func (t timeoutTickReadCloser) Close() error {
	t.t.Cancel()
	return t.r.Close()
}
