package common

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type oneTimeCloser struct {
	r        io.Reader
	isClosed bool
}

func newOneTimeCloser(r io.Reader) *oneTimeCloser {
	return &oneTimeCloser{r: r}
}

func (c *oneTimeCloser) Read(p []byte) (n int, err error) {
	if c.isClosed {
		return 0, io.ErrClosedPipe
	}
	return c.r.Read(p)
}

func (c *oneTimeCloser) Close() error {
	if c.isClosed {
		return io.ErrClosedPipe
	}
	c.isClosed = true
	return nil
}

func TestNewResponseWrapperBodyCanBeConsumed(t *testing.T) {
	ctx := context.Background()
	inner, err := http.NewRequestWithContext(ctx, "GET", "abc", nil)
	inner.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	req, err := NewRequestBody("abc123")
	assert.NoError(t, err)
	resp := &http.Response{
		Body:    newOneTimeCloser(strings.NewReader("Response Body")),
		Request: inner,
		Header:  make(http.Header),
	}
	resp.Header.Set("Content-Type", "application/json")

	wrapper, err := NewResponseWrapper(resp, req)
	assert.NoError(t, err)

	bs, err := io.ReadAll(wrapper.Response.Body)
	assert.NoError(t, err)
	// The response body should be available both in the Body field of the
	// *http.Response as well as the DebugBytes field of the ResponseWrapper,
	// because this is a non-streaming response.
	assert.Equal(t, "Response Body", string(bs))
	assert.Equal(t, "Response Body", string(wrapper.DebugBytes))
}
