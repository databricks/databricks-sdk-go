package logger

import (
	"context"
)

type LevelLogger struct {
	Level      Level
	BaseLogger SimpleLogger
}

func (l *LevelLogger) Enabled(_ context.Context, level Level) bool {
	return level >= l.Level
}

func (l *LevelLogger) Tracef(ctx context.Context, format string, v ...any) {
	l.BaseLogger.Tracef(ctx, format, v...)
}

func (l *LevelLogger) Debugf(ctx context.Context, format string, v ...any) {
	l.BaseLogger.Debugf(ctx, format, v...)
}

func (l *LevelLogger) Infof(ctx context.Context, format string, v ...any) {
	l.BaseLogger.Infof(ctx, format, v...)
}

func (l *LevelLogger) Warnf(ctx context.Context, format string, v ...any) {
	l.BaseLogger.Warnf(ctx, format, v...)
}

func (l *LevelLogger) Errorf(ctx context.Context, format string, v ...any) {
	l.BaseLogger.Errorf(ctx, format, v...)
}
