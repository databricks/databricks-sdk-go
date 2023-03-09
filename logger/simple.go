package logger

import (
	"context"
	"log"
)

type SimpleLogger struct{}

func (l *SimpleLogger) Enabled(_ context.Context, level Level) bool {
	return true
}

func (l *SimpleLogger) Tracef(_ context.Context, format string, v ...any) {
	log.Printf("[TRACE] "+format, v...)
}

func (l *SimpleLogger) Debugf(_ context.Context, format string, v ...any) {
	log.Printf("[DEBUG] "+format, v...)
}

func (l *SimpleLogger) Infof(_ context.Context, format string, v ...any) {
	log.Printf("[INFO] "+format, v...)
}

func (l *SimpleLogger) Warnf(_ context.Context, format string, v ...any) {
	log.Printf("[WARN] "+format, v...)
}

func (l *SimpleLogger) Errorf(_ context.Context, format string, v ...any) {
	log.Printf("[ERROR] "+format, v...)
}
