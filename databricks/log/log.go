// Package log provides the Databricks Logger interface and a default logger
// implementation.
//
// The Logger interface is designed to log messages at different levels of
// severity. When creating a logger, you can set a minimum log level, meaning
// any messages below this level will be ignored.
//
//	logger := log.New(log.LevelInfo) // drops trace and debug logs
//	log.SetDefaultLogger(logger)
//
// You can replace the default logger by implementing the Logger interface and
// assigning your custom logger.
//
//	type MyLogger struct {}
//	func (m *MyLogger) Log(ctx context.Context, level log.Level, format string, a ...any) {
//	    fmt.Printf("[%s] %s\n", level, fmt.Sprintf(format, a...))
//	}
//	log.SetDefaultLogger(&MyLogger{})
//
// You can retrieve the current default logger using the DefaultLogger function.
package log

import (
	"context"
	"fmt"
	"io"
	"os"
)

// Level represents the importance or severity of a log event. The higher the
// level, the more important or severe that event is.
type Level int

const (
	LevelTrace Level = -8
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger is the interface that wraps the Log method.
type Logger interface {
	Log(ctx context.Context, level Level, format string, a ...any)

	// TODO(renaud.hartert): Think about a potential alternative to exposing
	// this method.
	MinLevel() Level
}

type logger struct {
	level Level     // log level below that level will be dropped
	out   io.Writer // output writer for testing
}

// New creates a new logger with the given minimum log level. The returned
// logger will only log events that are at or above the minimum log level.
func New(level Level) Logger {
	return &logger{
		level: level,
		out:   os.Stdout,
	}
}

func (l *logger) Log(_ context.Context, level Level, format string, a ...any) {
	if l.level <= level {
		fmt.Fprintf(l.out, "[%s] %s\n", level.String(), fmt.Sprintf(format, a...))
	}
}

func (l *logger) MinLevel() Level {
	return l.level
}

var defaultLogger Logger = New(LevelInfo)

// DefaultLogger returns the default logger.
func DefaultLogger() Logger {
	return defaultLogger
}

// SetDefaultLogger sets the default logger. Any log events that are logged
// after setting the default logger will be logged using the new logger.
func SetDefaultLogger(logger Logger) {
	defaultLogger = logger
}

// Trace logs a message at the trace level.
func Trace(format string, v ...any) {
	TraceContext(context.Background(), format, v...)
}

// TraceContext logs a message at the trace level with context.
func TraceContext(ctx context.Context, format string, v ...any) {
	defaultLogger.Log(ctx, LevelTrace, format, v...)
}

// Debug logs a message at the debug level.
func Debug(format string, v ...any) {
	DebugContext(context.Background(), format, v...)
}

// DebugContext logs a message at the debug level with context.
func DebugContext(ctx context.Context, format string, v ...any) {
	defaultLogger.Log(ctx, LevelDebug, format, v...)
}

// Info logs a message at the info level.
func Info(format string, v ...any) {
	InfoContext(context.Background(), format, v...)
}

// InfoContext logs a message at the info level with context.
func InfoContext(ctx context.Context, format string, v ...any) {
	defaultLogger.Log(ctx, LevelInfo, format, v...)
}

// Warning logs a message at the warn level.
func Warning(format string, v ...any) {
	WarningContext(context.Background(), format, v...)
}

// WarningContext logs a message at the warn level with context.
func WarningContext(ctx context.Context, format string, v ...any) {
	defaultLogger.Log(ctx, LevelWarn, format, v...)
}

// Error logs a message at the error level.
func Error(format string, v ...any) {
	ErrorContext(context.Background(), format, v...)
}

// ErrorContext logs a message at the error level with context.
func ErrorContext(ctx context.Context, format string, v ...any) {
	defaultLogger.Log(ctx, LevelError, format, v...)
}
