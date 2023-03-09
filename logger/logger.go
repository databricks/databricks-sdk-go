package logger

import (
	"context"
)

var DefaultLogger Logger = &SimpleLogger{}

// Level maps to the logging levels in [Logger].
type Level int

const (
	LevelTrace = -8
	LevelDebug = -4
	LevelInfo  = 0
	LevelWarn  = 4
	LevelError = 8
)

type Logger interface {
	// Enabled returns if logging at the specified level is enabled.
	// This can be used to avoid computating an expensive value if it
	// won't be logged anyway.
	Enabled(level Level) bool

	// Tracef logs a formatted string at [LevelTrace].
	Tracef(format string, v ...any)

	// Debugf logs a formatted string at [LevelDebug].
	Debugf(format string, v ...any)

	// Infof logs a formatted string at [LevelInfo].
	Infof(format string, v ...any)

	// Warnf logs a formatted string at [LevelWarn].
	Warnf(format string, v ...any)

	// Errorf logs a formatted string at [LevelError].
	Errorf(format string, v ...any)
}

// Get returns either the logger configured on the context,
// or the global logger if one isn't defined.
func Get(ctx context.Context) Logger {
	logger, ok := FromContext(ctx)
	if !ok {
		logger = DefaultLogger
	}
	return logger
}

func Tracef(ctx context.Context, format string, v ...any) {
	Get(ctx).Tracef(format, v...)
}

func Debugf(ctx context.Context, format string, v ...any) {
	Get(ctx).Debugf(format, v...)
}

func Infof(ctx context.Context, format string, v ...any) {
	Get(ctx).Infof(format, v...)
}

func Warnf(ctx context.Context, format string, v ...any) {
	Get(ctx).Warnf(format, v...)
}

func Errorf(ctx context.Context, format string, v ...any) {
	Get(ctx).Errorf(format, v...)
}
