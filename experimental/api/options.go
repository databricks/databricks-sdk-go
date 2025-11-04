package api

import (
	"time"
)

// Option is an option used by Execute to control the behavior of an API call.
// An Option essentially act as a convenient way to configure Options.
type Option interface {
	Apply(*Options) error
}

type optionFn func(*Options)

func (fn optionFn) Apply(opts *Options) error {
	fn(opts)
	return nil
}

// WithRetrier configures to use the given Retrier provider. If no retrier is
// provided, the API call is not retried.
func WithRetrier(r func() Retrier) Option {
	return optionFn(func(o *Options) {
		o.retrier = r
	})
}

// WithTimeout is a convenience option to set the timeout duration. This
// option is ignored if the Execute call is made with a context that already
// contains a timeout.
//
// The timeout covers the whole APICall execution; it is not a timeout on each
// intermediary API call.
func WithTimeout(t time.Duration) Option {
	return optionFn(func(o *Options) {
		o.timeout = t
	})
}

// WithLimiter configures to use the given Limiter provider.
//
// The limiter is used to potential rate limit the API call. If no limiter is
// provided, the API call is not rate limited.
func WithLimiter(l Limiter) Option {
	return optionFn(func(o *Options) {
		o.rateLimiter = l
	})
}

// Options to control the behavior of an API call.
type Options struct {
	// Provides a new Retrier to be used to execute an APICall. The function
	// is called for each APICall and must be thread-safe. The retrier must be
	// fresh within the context of an Execute call (e.g. no need to reset a
	// BackoffStrategy).
	retrier func() Retrier

	rateLimiter Limiter

	timeout time.Duration
}
