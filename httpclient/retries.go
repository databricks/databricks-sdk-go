package httpclient

import (
	"github.com/databricks/databricks-sdk-go/retries"
)

func makeRetrier[T any](c ClientConfig) retries.Retrier[T] {
	return retries.New[T](
		retries.WithTimeout(c.RetryTimeout),
		retries.WithBackoffFunc(c.RetryBackoff),
	)
}
