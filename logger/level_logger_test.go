package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenInfoLevelThenDebugDisabled(t *testing.T) {
	t.Cleanup(func() {
		DefaultLogger = &SimpleLogger{}
	})

	infoLevelLogger := &LevelLogger{
		Level: LevelInfo,
	}
	debugEnabled := infoLevelLogger.Enabled(context.Background(), LevelDebug)
	assert.False(t, debugEnabled)
}

func TestWhenInfoLevelThenErrorEnabled(t *testing.T) {
	infoLevelLogger := &LevelLogger{
		Level: LevelInfo,
	}

	errorEnabled := infoLevelLogger.Enabled(context.Background(), LevelError)
	assert.True(t, errorEnabled)
}
