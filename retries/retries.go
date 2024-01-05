package retries

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/logger"
)

// Deprecated: use return types from non-*AndWait methods
type Info[T any] struct {
	Info    *T
	Timeout time.Duration
}

// Deprecated: use return types from non-*AndWait methods
type Option[T any] func(*Info[T])

// Deprecated: use return types from non-*AndWait methods
func Timeout[T any](dur time.Duration) Option[T] {
	return func(i *Info[T]) {
		i.Timeout = dur
	}
}

// Deprecated: use return types from non-*AndWait methods
func OnPoll[T any](callback func(*T)) Option[T] {
	return func(i *Info[T]) {
		if i.Info == nil {
			return
		}
		callback(i.Info)
	}
}

type Err struct {
	Err  error
	Halt bool
}

func (e *Err) Error() string {
	return e.Err.Error()
}

func Halt(err error) *Err {
	return &Err{err, true}
}

func Continue(err error) *Err {
	return &Err{err, false}
}

func Continues(msg string) *Err {
	return Continue(fmt.Errorf(msg))
}

func Continuef(format string, err error, args ...interface{}) *Err {
	wrapped := fmt.Errorf(format, append([]interface{}{err}, args...))
	return Continue(wrapped)
}

var maxWait = 10 * time.Second
var minJitter = 50 * time.Millisecond
var maxJitter = 750 * time.Millisecond

func Backoff(attempt int) time.Duration {
	wait := time.Duration(attempt) * time.Second
	if wait > maxWait {
		wait = maxWait
	}
	// add some random jitter
	rand.Seed(time.Now().UnixNano())
	jitter := rand.Intn(int(maxJitter)-int(minJitter)+1) + int(minJitter)
	wait += time.Duration(jitter)
	return wait
}

type ErrTimedOut struct {
	err error
}

func (et *ErrTimedOut) Error() string {
	return fmt.Sprintf("timed out: %s", et.err)
}

func (et *ErrTimedOut) Unwrap() error {
	return et.err
}

type RetryOption func(*RetryConfig)

type RetryConfig struct {
	timeout    time.Duration
	shouldHalt func(error) bool
}

func WithTimeout(timeout time.Duration) RetryOption {
	return func(rc *RetryConfig) {
		rc.timeout = timeout
	}
}

func OnErrors(on ...error) RetryOption {
	return func(rc *RetryConfig) {
		rc.shouldHalt = func(err error) bool {
			for _, e := range on {
				if errors.Is(err, e) {
					return true
				}
			}
			return false
		}
	}
}

func WithHalt(halt func(error) bool) RetryOption {
	return func(rc *RetryConfig) {
		rc.shouldHalt = halt
	}
}

type retrier[T any] struct {
	config RetryConfig
}

func New[T any](configOpts ...func(*RetryConfig)) *retrier[T] {
	config := RetryConfig{
		timeout: 20 * time.Minute,
	}
	for _, opt := range configOpts {
		opt(&config)
	}
	return &retrier[T]{config}
}

func (r *retrier[T]) Wait(ctx context.Context, fn func(ctx context.Context) error) error {
	_, err := r.Run(ctx, func(ctx context.Context) (*T, error) {
		return nil, fn(ctx)
	})
	return err
}

func (r *retrier[T]) Run(ctx context.Context, fn func(context.Context) (*T, error)) (*T, error) {
	ctx, cancel := context.WithTimeout(ctx, r.config.timeout)
	defer cancel()
	var attempt int
	var lastErr error
	for {
		attempt++
		entity, err := fn(ctx)
		if err == nil {
			return entity, nil
		}
		if r.config.shouldHalt != nil && r.config.shouldHalt(err) {
			return nil, err
		}
		lastErr = err
		wait := Backoff(attempt)
		timer := time.NewTimer(wait)
		logger.Tracef(ctx, "%s. Sleeping %s",
			strings.TrimSuffix(err.Error(), "."),
			wait.Round(time.Millisecond))
		select {
		// stop when either this or parent context times out
		case <-ctx.Done():
			timer.Stop()
			return nil, &ErrTimedOut{lastErr}
		case <-timer.C:
		}
	}
}

func isHalt(err error) bool {
	if err == nil {
		return true
	}
	e := err.(*Err)
	return e == nil || e.Halt
}

func Wait(ctx context.Context, timeout time.Duration, fn func() *Err) error {
	return New[struct{}](WithTimeout(timeout), WithHalt(isHalt)).Wait(ctx, func(_ context.Context) error {
		err := fn()
		if err != nil {
			return err
		}
		return nil
	})
}

func Poll[T any](ctx context.Context, timeout time.Duration, fn func() (*T, *Err)) (*T, error) {
	return New[T](WithTimeout(timeout), WithHalt(isHalt)).Run(ctx, func(_ context.Context) (*T, error) {
		res, err := fn()
		if err != nil {
			return res, err
		}
		return res, nil
	})
}
