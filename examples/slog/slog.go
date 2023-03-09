package main

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/exp/slog"
)

// slogAdapter makes an slog.Logger usable with the Databricks SDK.
type slogAdapter struct {
	*slog.Logger
}

func (s *slogAdapter) Enabled(level logger.Level) bool {
	switch level {
	case logger.LevelTrace:
		return s.Logger.Enabled(nil, slog.LevelDebug)
	case logger.LevelDebug:
		return s.Logger.Enabled(nil, slog.LevelDebug)
	case logger.LevelInfo:
		return s.Logger.Enabled(nil, slog.LevelInfo)
	case logger.LevelWarn:
		return s.Logger.Enabled(nil, slog.LevelWarn)
	case logger.LevelError:
		return s.Logger.Enabled(nil, slog.LevelError)
	default:
		return true
	}
}

func (s *slogAdapter) Tracef(format string, v ...any) {
	s.Debug(fmt.Sprintf(format, v...))
}

func (s *slogAdapter) Debugf(format string, v ...any) {
	s.Debug(fmt.Sprintf(format, v...))
}

func (s *slogAdapter) Infof(format string, v ...any) {
	s.Info(fmt.Sprintf(format, v...))
}

func (s *slogAdapter) Warnf(format string, v ...any) {
	s.Warn(fmt.Sprintf(format, v...))
}

func (s *slogAdapter) Errorf(format string, v ...any) {
	s.Error(fmt.Sprintf(format, v...), nil)
}
