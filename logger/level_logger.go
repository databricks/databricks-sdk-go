package logger

import (
	"context"
)

type LevelLogger struct {
	SimpleLogger
	Level Level
}

func (l *LevelLogger) Enabled(_ context.Context, level Level) bool {
	return level >= l.Level
}
