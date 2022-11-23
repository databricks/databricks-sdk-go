package logger

import "log"

var DefaultLogger Logger = SimpleLogger{}

func Tracef(format string, v ...interface{}) {
	DefaultLogger.Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	DefaultLogger.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	DefaultLogger.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	DefaultLogger.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	DefaultLogger.Errorf(format, v...)
}

type Logger interface {
	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type SimpleLogger struct{}

func (l SimpleLogger) Tracef(format string, v ...interface{}) {
	log.Printf("[TRACE] "+format, v...)
}

func (l SimpleLogger) Debugf(format string, v ...interface{}) {
	log.Printf("[DEBUG] "+format, v...)
}

func (l SimpleLogger) Infof(format string, v ...interface{}) {
	log.Printf("[INFO] "+format, v...)
}

func (l SimpleLogger) Warnf(format string, v ...interface{}) {
	log.Printf("[WARN] "+format, v...)
}

func (l SimpleLogger) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR] "+format, v...)
}
