package httpclient

import (
	"github.com/databricks/databricks-sdk-go/retries"
)

func makeRetrier[T any](c ClientConfig) retries.Retrier[T] {
	return retries.New[T](
		retries.WithTimeout(c.RetryTimeout),
		retries.WithRetryFunc(retries.DefaultShouldRetry),
		retries.WithBackoffFunc(c.RetryBackoff),
	)
}
