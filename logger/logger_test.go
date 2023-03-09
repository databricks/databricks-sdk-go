package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	t.Cleanup(func() {
		DefaultLogger = &SimpleLogger{}
	})

	t1 := &SimpleLogger{}
	t2 := &SimpleLogger{}
	DefaultLogger = t1

	logger := Get(context.Background())
	assert.Equal(t, logger, t1)

	ctx := NewContext(context.Background(), t2)
	logger = Get(ctx)
	assert.Equal(t, logger, t2)
}
