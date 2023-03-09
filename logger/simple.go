package logger

import "log"

type SimpleLogger struct{}

func (l *SimpleLogger) Enabled(level Level) bool {
	return true
}

func (l *SimpleLogger) Tracef(format string, v ...any) {
	log.Printf("[TRACE] "+format, v...)
}

func (l *SimpleLogger) Debugf(format string, v ...any) {
	log.Printf("[DEBUG] "+format, v...)
}

func (l *SimpleLogger) Infof(format string, v ...any) {
	log.Printf("[INFO] "+format, v...)
}

func (l *SimpleLogger) Warnf(format string, v ...any) {
	log.Printf("[WARN] "+format, v...)
}

func (l *SimpleLogger) Errorf(format string, v ...any) {
	log.Printf("[ERROR] "+format, v...)
}
