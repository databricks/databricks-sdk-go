package retries

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/logger"
)

type Info[T any] struct {
	Info    *T
	Timeout time.Duration
}

type Option[T any] func(*Info[T])

func Timeout[T any](dur time.Duration) Option[T] {
	return func(i *Info[T]) {
		i.Timeout = dur
	}
}

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

type WaitFn func() *Err

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

func Wait(pctx context.Context, timeout time.Duration, fn WaitFn) error {
	ctx, cancel := context.WithTimeout(pctx, timeout)
	defer cancel()
	var attempt int
	var lastErr error
	for {
		attempt++
		res := fn()
		if res == nil {
			return nil
		}
		if res.Halt {
			return res.Err
		}
		lastErr = res.Err
		wait := Backoff(attempt)
		timer := time.NewTimer(wait)
		logger.Tracef(ctx, "%s. Sleeping %s",
			strings.TrimSuffix(res.Err.Error(), "."),
			wait.Round(time.Millisecond))
		select {
		// stop when either this or parent context times out
		case <-ctx.Done():
			timer.Stop()
			return fmt.Errorf("timed out: %w", lastErr)
		case <-timer.C:
		}
	}
}

func Poll[T any](pctx context.Context, timeout time.Duration, fn func() (*T, *Err)) (*T, error) {
	ctx, cancel := context.WithTimeout(pctx, timeout)
	defer cancel()
	var attempt int
	var lastErr error
	for {
		attempt++
		entity, err := fn()
		if err == nil {
			return entity, nil
		}
		if err.Halt {
			return nil, err.Err
		}
		lastErr = err.Err
		wait := Backoff(attempt)
		timer := time.NewTimer(wait)
		logger.Tracef(ctx, "%s. Sleeping %s",
			strings.TrimSuffix(err.Err.Error(), "."),
			wait.Round(time.Millisecond))
		select {
		// stop when either this or parent context times out
		case <-ctx.Done():
			timer.Stop()
			return nil, fmt.Errorf("timed out: %w", lastErr)
		case <-timer.C:
		}
	}
}
